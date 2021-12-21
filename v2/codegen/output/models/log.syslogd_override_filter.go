package models

const LogsyslogdOverrideFilterPath = "log.syslogd/override-filter/"

type LogsyslogdOverrideFilter struct {
	Anomaly          *string                             `json:"anomaly,omitempty"`
	ForwardTraffic   *string                             `json:"forward-traffic,omitempty"`
	FreeStyle        []LogsyslogdOverrideFilterFreeStyle `json:"free-style,omitempty"`
	LocalTraffic     *string                             `json:"local-traffic,omitempty"`
	MulticastTraffic *string                             `json:"multicast-traffic,omitempty"`
	Severity         *string                             `json:"severity,omitempty"`
	SnifferTraffic   *string                             `json:"sniffer-traffic,omitempty"`
	Voip             *string                             `json:"voip,omitempty"`
}

type LogsyslogdOverrideFilterFreeStyle struct {
	Category   *string `json:"category,omitempty"`
	Filter     *string `json:"filter,omitempty"`
	FilterType *string `json:"filter-type,omitempty"`
	Id         *int64  `json:"id,omitempty"`
}