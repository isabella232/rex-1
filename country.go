package rex

import "strings"

func CountryNormalizedAlpha2(country string) string {
	if country == "" {
		return "xx"
	}

	c := strings.ToLower(country)

	if len(c) == 3 {
		c = Countries[c]
		if c == "" {
			return "xx"
		}
	}

	if len(c) != 2 {
		return "xx"
	}

	return c
}

var Countries map[string]string = map[string]string{
	"and": "ad",
	"are": "ae",
	"afg": "af",
	"atg": "ag",
	"aia": "ai",
	"alb": "al",
	"arm": "am",
	"ago": "ao",
	"ata": "aq",
	"arg": "ar",
	"asm": "as",
	"aut": "at",
	"aus": "au",
	"abw": "aw",
	"ala": "ax",
	"aze": "az",
	"bih": "ba",
	"brb": "bb",
	"bgd": "bd",
	"bel": "be",
	"bfa": "bf",
	"bgr": "bg",
	"bhr": "bh",
	"bdi": "bi",
	"ben": "bj",
	"blm": "bl",
	"bmu": "bm",
	"brn": "bn",
	"bol": "bo",
	"bes": "bq",
	"bra": "br",
	"bhs": "bs",
	"btn": "bt",
	"bvt": "bv",
	"bwa": "bw",
	"blr": "by",
	"blz": "bz",
	"can": "ca",
	"cck": "cc",
	"cod": "cd",
	"caf": "cf",
	"cog": "cg",
	"che": "ch",
	"civ": "ci",
	"cok": "ck",
	"chl": "cl",
	"cmr": "cm",
	"chn": "cn",
	"col": "co",
	"cri": "cr",
	"cub": "cu",
	"cpv": "cv",
	"cuw": "cw",
	"cxr": "cx",
	"cyp": "cy",
	"cze": "cz",
	"deu": "de",
	"dji": "dj",
	"dnk": "dk",
	"dma": "dm",
	"dom": "do",
	"dza": "dz",
	"ecu": "ec",
	"est": "ee",
	"egy": "eg",
	"esh": "eh",
	"eri": "er",
	"esp": "es",
	"eth": "et",
	"fin": "fi",
	"fji": "fj",
	"flk": "fk",
	"fsm": "fm",
	"fro": "fo",
	"fra": "fr",
	"gab": "ga",
	"gbr": "gb",
	"grd": "gd",
	"geo": "ge",
	"guf": "gf",
	"ggy": "gg",
	"gha": "gh",
	"gib": "gi",
	"grl": "gl",
	"gmb": "gm",
	"gin": "gn",
	"glp": "gp",
	"gnq": "gq",
	"grc": "gr",
	"sgs": "gs",
	"gtm": "gt",
	"gum": "gu",
	"gnb": "gw",
	"guy": "gy",
	"hkg": "hk",
	"hmd": "hm",
	"hnd": "hn",
	"hrv": "hr",
	"hti": "ht",
	"hun": "hu",
	"idn": "id",
	"irl": "ie",
	"isr": "il",
	"imn": "im",
	"ind": "in",
	"iot": "io",
	"irq": "iq",
	"irn": "ir",
	"isl": "is",
	"ita": "it",
	"jey": "je",
	"jam": "jm",
	"jor": "jo",
	"jpn": "jp",
	"ken": "ke",
	"kgz": "kg",
	"khm": "kh",
	"kir": "ki",
	"com": "km",
	"kna": "kn",
	"prk": "kp",
	"kor": "kr",
	"xkx": "xk",
	"kwt": "kw",
	"cym": "ky",
	"kaz": "kz",
	"lao": "la",
	"lbn": "lb",
	"lca": "lc",
	"lie": "li",
	"lka": "lk",
	"lbr": "lr",
	"lso": "ls",
	"ltu": "lt",
	"lux": "lu",
	"lva": "lv",
	"lby": "ly",
	"mar": "ma",
	"mco": "mc",
	"mda": "md",
	"mne": "me",
	"maf": "mf",
	"mdg": "mg",
	"mhl": "mh",
	"mkd": "mk",
	"mli": "ml",
	"mmr": "mm",
	"mng": "mn",
	"mac": "mo",
	"mnp": "mp",
	"mtq": "mq",
	"mrt": "mr",
	"msr": "ms",
	"mlt": "mt",
	"mus": "mu",
	"mdv": "mv",
	"mwi": "mw",
	"mex": "mx",
	"mys": "my",
	"moz": "mz",
	"nam": "na",
	"ncl": "nc",
	"ner": "ne",
	"nfk": "nf",
	"nga": "ng",
	"nic": "ni",
	"nld": "nl",
	"nor": "no",
	"npl": "np",
	"nru": "nr",
	"niu": "nu",
	"nzl": "nz",
	"omn": "om",
	"pan": "pa",
	"per": "pe",
	"pyf": "pf",
	"png": "pg",
	"phl": "ph",
	"pak": "pk",
	"pol": "pl",
	"spm": "pm",
	"pcn": "pn",
	"pri": "pr",
	"pse": "ps",
	"prt": "pt",
	"plw": "pw",
	"pry": "py",
	"qat": "qa",
	"reu": "re",
	"rou": "ro",
	"srb": "rs",
	"rus": "ru",
	"rwa": "rw",
	"sau": "sa",
	"slb": "sb",
	"syc": "sc",
	"sdn": "sd",
	"ssd": "ss",
	"swe": "se",
	"sgp": "sg",
	"shn": "sh",
	"svn": "si",
	"sjm": "sj",
	"svk": "sk",
	"sle": "sl",
	"smr": "sm",
	"sen": "sn",
	"som": "so",
	"sur": "sr",
	"stp": "st",
	"slv": "sv",
	"sxm": "sx",
	"syr": "sy",
	"swz": "sz",
	"tca": "tc",
	"tcd": "td",
	"atf": "tf",
	"tgo": "tg",
	"tha": "th",
	"tjk": "tj",
	"tkl": "tk",
	"tls": "tl",
	"tkm": "tm",
	"tun": "tn",
	"ton": "to",
	"tur": "tr",
	"tto": "tt",
	"tuv": "tv",
	"twn": "tw",
	"tza": "tz",
	"ukr": "ua",
	"uga": "ug",
	"umi": "um",
	"usa": "us",
	"ury": "uy",
	"uzb": "uz",
	"vat": "va",
	"vct": "vc",
	"ven": "ve",
	"vgb": "vg",
	"vir": "vi",
	"vnm": "vn",
	"vut": "vu",
	"wlf": "wf",
	"wsm": "ws",
	"yem": "ye",
	"myt": "yt",
	"zaf": "za",
	"zmb": "zm",
	"zwe": "zw",
	"scg": "cs",
	"ant": "an",
}
