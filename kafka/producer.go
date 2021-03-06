package kafka

import (
	"errors"
	"fmt"
	"time"

	"github.com/Shopify/sarama"
	"github.com/juju/loggo"
	"github.com/rcrowley/go-metrics"
)

type ProduceErrorCallback func(*sarama.ProduceError)

type Producer struct {
	*sarama.Producer
	config   *sarama.ProducerConfig
	callback ProduceErrorCallback
	quit     chan bool
	done     chan bool
	log      loggo.Logger
	messages metrics.Timer
	errors   metrics.Timer
}

func (client *Client) NewProducer(name string, config *sarama.ProducerConfig, cb ProduceErrorCallback) (*Producer, error) {
	name = fmt.Sprintf("kafka.producer.%s.%s", client.GetId(), name)

	self := &Producer{
		config:   config,
		callback: cb,
		quit:     make(chan bool),
		done:     make(chan bool),
		log:      loggo.GetLogger(name),
		messages: metrics.NewRegisteredTimer(name+".messages", metrics.DefaultRegistry),
		errors:   metrics.NewRegisteredTimer(name+".errors", metrics.DefaultRegistry),
	}

	if self.config == nil {
		self.log.Infof("using default producer config")
		self.config = sarama.NewProducerConfig()
	}

	producer, err := sarama.NewProducer(client.Client, self.config)
	if err != nil {
		self.log.Errorf("failed to create producer: %s", err)
		return nil, err
	}

	self.Producer = producer
	go self.Start()

	return self, nil
}

func (client *Client) NewFastProducer(cb ProduceErrorCallback) (*Producer, error) {
	config := sarama.NewProducerConfig()
	config.AckSuccesses = false
	config.RequiredAcks = sarama.NoResponse
	config.FlushMsgCount = 10000
	config.FlushFrequency = 10 * time.Millisecond
	return client.NewProducer("fast", config, cb)
}

func (client *Client) NewSafeProducer() (*Producer, error) {
	config := sarama.NewProducerConfig()
	config.AckSuccesses = true
	config.RequiredAcks = sarama.WaitForAll
	config.Timeout = 20 * time.Millisecond
	return client.NewProducer("safe", config, nil)
}

func (self *Producer) Start() {
	if self.config.AckSuccesses == true {
		return
	}

	for {
		select {
		case err, ok := <-self.Errors():
			if ok && self.callback != nil {
				self.callback(err)
			}
		case <-self.quit:
			close(self.done)
			return
		}
	}
}

func (self *Producer) Shutdown() {
	if self.config.AckSuccesses == false {
		self.log.Infof("shutting down producer run loop")
		close(self.quit)
		self.log.Infof("waiting for run loop to finish")
		<-self.done
	}
	self.log.Infof("closing producer")
	self.Close()
	self.log.Infof("shutdown done")
}

func (self *Producer) Message(topic string, value []byte, timeout time.Duration) (msg *sarama.MessageToSend, err error) {
	start := time.Now()

	if value == nil || len(value) < 1 {
		self.errors.Update(time.Since(start))
		return msg, errors.New("empty message")
	}

	msg = &sarama.MessageToSend{
		Topic: string(topic),
		Value: sarama.ByteEncoder(value),
	}

	defer func() {
		// we might have tried to write to a closed channel during shutdown
		// recover and return the error to the caller
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("unknown error: %v", r)
			}
			self.errors.Update(time.Since(start))
		}
	}()

	select {
	case self.Input() <- msg:
		// fall through
	case <-time.After(timeout):
		self.errors.Update(time.Since(start))
		return msg, errors.New("input timed out")
	}

	if self.config.AckSuccesses == false {
		self.messages.Update(time.Since(start))
		return msg, nil
	}

	select {
	case err := <-self.Errors():
		self.errors.Update(time.Since(start))
		return msg, err.Err
	case <-self.Successes():
		self.messages.Update(time.Since(start))
	case <-time.After(self.config.Timeout):
		self.errors.Update(time.Since(start))
		return msg, errors.New("ack timed out")
	}

	return msg, nil
}
