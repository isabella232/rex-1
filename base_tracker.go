package rex

import (
	"encoding/json"

	"time"
)

type BaseTracker struct {
	*EventMetadata
}

func NewBaseTracker(metadata *EventMetadata) *BaseTracker {
	return &BaseTracker{
		EventMetadata: metadata,
	}
}

func (self *BaseTracker) Encode(message interface{}) []byte {
	bytes, _ := json.Marshal(message)
	return bytes
}

func (self *BaseTracker) AddMetadata(e EventBase, full bool) {
	event := e.Base()

	if event.Timestamp == 0 {
		event.Timestamp = getCurrentTime()

		if full == true {
			event.Service = self.Service
			event.Environment = self.Environment
			event.Cluster = self.Cluster
			event.Host = self.Host
			event.Release = self.Release
		}
	}
}

func (self *BaseTracker) AddMetadataMap(event map[string]interface{}, full bool) {
	if event["timestamp"] == 0 {
		event["timestamp"] = getCurrentTime()

		if full == true {
			event["service"] = self.Service
			event["env"] = self.Environment
			event["cluster"] = self.Cluster
			event["host"] = self.Host
			event["release"] = self.Release
		}
	}
}
func getCurrentTime() float64 {
	return time.Now().UnixNano() / 1000000000
}
