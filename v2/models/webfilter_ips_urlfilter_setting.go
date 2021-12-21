package models

const WebfilterIpsUrlfilterSettingPath = "webfilter/ips-urlfilter-setting/"

type WebfilterIpsUrlfilterSetting struct {
	Device    *string  `json:"device,omitempty"`
	Distance  *float64 `json:"distance,omitempty"`
	Gateway   *string  `json:"gateway,omitempty"`
	GeoFilter *string  `json:"geo-filter,omitempty"`
}
