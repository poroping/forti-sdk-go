package models

const LogFortiguardOverrideFilterPath = "log.fortiguard/override-filter/"

type LogFortiguardOverrideFilter struct {
	Anomaly          *string                                 `json:"anomaly,omitempty"`
	Filter           *string                                 `json:"filter,omitempty"`
	FilterType       *string                                 `json:"filter-type,omitempty"`
	ForwardTraffic   *string                                 `json:"forward-traffic,omitempty"`
	FreeStyle        *[]LogFortiguardOverrideFilterFreeStyle `json:"free-style,omitempty"`
	Gtp              *string                                 `json:"gtp,omitempty"`
	LocalTraffic     *string                                 `json:"local-traffic,omitempty"`
	MulticastTraffic *string                                 `json:"multicast-traffic,omitempty"`
	Severity         *string                                 `json:"severity,omitempty"`
	SnifferTraffic   *string                                 `json:"sniffer-traffic,omitempty"`
	Voip             *string                                 `json:"voip,omitempty"`
	ZtnaTraffic      *string                                 `json:"ztna-traffic,omitempty"`
}

const LogFortiguardOverrideFilterFreeStylePath = "log.fortiguard/override-filter/free-style/"

type LogFortiguardOverrideFilterFreeStyle struct {
	Category   *string `json:"category,omitempty"`
	Filter     *string `json:"filter,omitempty"`
	FilterType *string `json:"filter-type,omitempty"`
	Id         *int64  `json:"id,omitempty"`
}
