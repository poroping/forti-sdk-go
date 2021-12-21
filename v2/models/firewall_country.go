package models

const FirewallCountryPath = "firewall/country/"

type FirewallCountry struct {
	Fosid  *int64                   `json:"fosid,omitempty"`
	Name   *string                  `json:"name,omitempty"`
	Region *[]FirewallCountryRegion `json:"region,omitempty"`
}

type FirewallCountryRegion struct {
	Id *int64 `json:"id,omitempty"`
}
