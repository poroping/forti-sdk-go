package models

const LogSyslogd4OverrideFilterPath = "log.syslogd4/override-filter/"

type LogSyslogd4OverrideFilter struct {
	Anomaly          *string                               `json:"anomaly,omitempty"`
	Filter           *string                               `json:"filter,omitempty"`
	FilterType       *string                               `json:"filter-type,omitempty"`
	ForwardTraffic   *string                               `json:"forward-traffic,omitempty"`
	FreeStyle        *[]LogSyslogd4OverrideFilterFreeStyle `json:"free-style,omitempty"`
	Gtp              *string                               `json:"gtp,omitempty"`
	LocalTraffic     *string                               `json:"local-traffic,omitempty"`
	MulticastTraffic *string                               `json:"multicast-traffic,omitempty"`
	Severity         *string                               `json:"severity,omitempty"`
	SnifferTraffic   *string                               `json:"sniffer-traffic,omitempty"`
	Voip             *string                               `json:"voip,omitempty"`
	ZtnaTraffic      *string                               `json:"ztna-traffic,omitempty"`
}

const LogSyslogd4OverrideFilterFreeStylePath = "log.syslogd4/override-filter/free-style/"

type LogSyslogd4OverrideFilterFreeStyle struct {
	Category   *string `json:"category,omitempty"`
	Filter     *string `json:"filter,omitempty"`
	FilterType *string `json:"filter-type,omitempty"`
	Id         *int64  `json:"id,omitempty"`
}
