package models

const LogFortianalyzer2FilterPath = "log.fortianalyzer2/filter/"

type LogFortianalyzer2Filter struct {
	Anomaly          *string                             `json:"anomaly,omitempty"`
	DlpArchive       *string                             `json:"dlp-archive,omitempty"`
	Filter           *string                             `json:"filter,omitempty"`
	FilterType       *string                             `json:"filter-type,omitempty"`
	ForwardTraffic   *string                             `json:"forward-traffic,omitempty"`
	FreeStyle        *[]LogFortianalyzer2FilterFreeStyle `json:"free-style,omitempty"`
	Gtp              *string                             `json:"gtp,omitempty"`
	LocalTraffic     *string                             `json:"local-traffic,omitempty"`
	MulticastTraffic *string                             `json:"multicast-traffic,omitempty"`
	Severity         *string                             `json:"severity,omitempty"`
	SnifferTraffic   *string                             `json:"sniffer-traffic,omitempty"`
	Voip             *string                             `json:"voip,omitempty"`
	ZtnaTraffic      *string                             `json:"ztna-traffic,omitempty"`
}

const LogFortianalyzer2FilterFreeStylePath = "log.fortianalyzer2/filter/free-style/"

type LogFortianalyzer2FilterFreeStyle struct {
	Category   *string `json:"category,omitempty"`
	Filter     *string `json:"filter,omitempty"`
	FilterType *string `json:"filter-type,omitempty"`
	Id         *int64  `json:"id,omitempty"`
}
