package models

const LognullDeviceFilterPath = "log/null-device/filter/"

type LognullDeviceFilter struct {
	Anomaly          *string                         `json:"anomaly,omitempty"`
	Filter           *string                         `json:"filter,omitempty"`
	FilterType       *string                         `json:"filter-type,omitempty"`
	ForwardTraffic   *string                         `json:"forward-traffic,omitempty"`
	FreeStyle        *[]LognullDeviceFilterFreeStyle `json:"free-style,omitempty"`
	Gtp              *string                         `json:"gtp,omitempty"`
	LocalTraffic     *string                         `json:"local-traffic,omitempty"`
	MulticastTraffic *string                         `json:"multicast-traffic,omitempty"`
	Severity         *string                         `json:"severity,omitempty"`
	SnifferTraffic   *string                         `json:"sniffer-traffic,omitempty"`
	Voip             *string                         `json:"voip,omitempty"`
}

type LognullDeviceFilterFreeStyle struct {
	Category   *string `json:"category,omitempty"`
	Filter     *string `json:"filter,omitempty"`
	FilterType *string `json:"filter-type,omitempty"`
	Id         *int64  `json:"id,omitempty"`
}
