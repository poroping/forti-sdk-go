package models

const SystemGeoipCountryPath = "system/geoip-country/"

type SystemGeoipCountry struct {
	Fosid *string `json:"fosid,omitempty"`
	Name  *string `json:"name,omitempty"`
}
