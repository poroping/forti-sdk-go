package models

const WebfilterIpsUrlfilterSetting6Path = "webfilter/ips-urlfilter-setting6/"

type WebfilterIpsUrlfilterSetting6 struct {
	Device    *string `json:"device,omitempty"`
	Distance  *int64  `json:"distance,omitempty"`
	Gateway6  *string `json:"gateway6,omitempty"`
	GeoFilter *string `json:"geo-filter,omitempty"`
}
