package models

const Logsyslogd3FilterPath = "log.syslogd3/filter/"

type Logsyslogd3Filter struct {
	Anomaly          *string                      `json:"anomaly,omitempty"`
	ForwardTraffic   *string                      `json:"forward-traffic,omitempty"`
	FreeStyle        []Logsyslogd3FilterFreeStyle `json:"free-style,omitempty"`
	LocalTraffic     *string                      `json:"local-traffic,omitempty"`
	MulticastTraffic *string                      `json:"multicast-traffic,omitempty"`
	Severity         *string                      `json:"severity,omitempty"`
	SnifferTraffic   *string                      `json:"sniffer-traffic,omitempty"`
	Voip             *string                      `json:"voip,omitempty"`
}

type Logsyslogd3FilterFreeStyle struct {
	Category   *string `json:"category,omitempty"`
	Filter     *string `json:"filter,omitempty"`
	FilterType *string `json:"filter-type,omitempty"`
	Id         *int64  `json:"id,omitempty"`
}
