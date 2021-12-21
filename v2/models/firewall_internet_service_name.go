package models

const FirewallInternetServiceNamePath = "firewall/internet-service-name/"

type FirewallInternetServiceName struct {
	CityId            *float64 `json:"city-id,omitempty"`
	CountryId         *float64 `json:"country-id,omitempty"`
	InternetServiceId *float64 `json:"internet-service-id,omitempty"`
	Name              *string  `json:"name,omitempty"`
	RegionId          *float64 `json:"region-id,omitempty"`
	Type              *string  `json:"type,omitempty"`
}
