package models

const FirewallInternetServiceNamePath = "firewall/internet-service-name/"

type FirewallInternetServiceName struct {
	CityId            *int64  `json:"city-id,omitempty"`
	CountryId         *int64  `json:"country-id,omitempty"`
	InternetServiceId *int64  `json:"internet-service-id,omitempty"`
	Name              *string `json:"name,omitempty"`
	RegionId          *int64  `json:"region-id,omitempty"`
	Type              *string `json:"type,omitempty"`
}
