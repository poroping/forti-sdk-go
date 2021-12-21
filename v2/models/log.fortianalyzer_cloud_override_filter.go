package models

const LogfortianalyzerCloudOverrideFilterPath = "log.fortianalyzer-cloud/override-filter/"

type LogfortianalyzerCloudOverrideFilter struct {
	Anomaly          *string                                         `json:"anomaly,omitempty"`
	DlpArchive       *string                                         `json:"dlp-archive,omitempty"`
	Filter           *string                                         `json:"filter,omitempty"`
	FilterType       *string                                         `json:"filter-type,omitempty"`
	ForwardTraffic   *string                                         `json:"forward-traffic,omitempty"`
	FreeStyle        *[]LogfortianalyzerCloudOverrideFilterFreeStyle `json:"free-style,omitempty"`
	Gtp              *string                                         `json:"gtp,omitempty"`
	LocalTraffic     *string                                         `json:"local-traffic,omitempty"`
	MulticastTraffic *string                                         `json:"multicast-traffic,omitempty"`
	Severity         *string                                         `json:"severity,omitempty"`
	SnifferTraffic   *string                                         `json:"sniffer-traffic,omitempty"`
	Voip             *string                                         `json:"voip,omitempty"`
}

type LogfortianalyzerCloudOverrideFilterFreeStyle struct {
	Category   *string `json:"category,omitempty"`
	Filter     *string `json:"filter,omitempty"`
	FilterType *string `json:"filter-type,omitempty"`
	Id         *int64  `json:"id,omitempty"`
}
