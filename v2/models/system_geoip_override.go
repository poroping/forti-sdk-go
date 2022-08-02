package models

const SystemGeoipOverridePath = "system/geoip-override/"

type SystemGeoipOverride struct {
	CountryId   *string                        `json:"country-id,omitempty"`
	Description *string                        `json:"description,omitempty"`
	IpRange     *[]SystemGeoipOverrideIpRange  `json:"ip-range,omitempty"`
	Ip6Range    *[]SystemGeoipOverrideIp6Range `json:"ip6-range,omitempty"`
	Name        *string                        `json:"name,omitempty"`
}

const SystemGeoipOverrideIpRangePath = "system/geoip-override/ip-range/"

type SystemGeoipOverrideIpRange struct {
	EndIp   *string `json:"end-ip,omitempty"`
	Id      *int64  `json:"id,omitempty"`
	StartIp *string `json:"start-ip,omitempty"`
}

const SystemGeoipOverrideIp6RangePath = "system/geoip-override/ip6-range/"

type SystemGeoipOverrideIp6Range struct {
	EndIp   *string `json:"end-ip,omitempty"`
	Id      *int64  `json:"id,omitempty"`
	StartIp *string `json:"start-ip,omitempty"`
}
