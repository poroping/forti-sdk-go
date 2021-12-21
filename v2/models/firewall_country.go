package models

const FirewallCountryPath = "firewall/country/"

type FirewallCountry struct {
	Fosid  *float64                 `json:"fosid,omitempty"`
	Name   *string                  `json:"name,omitempty"`
	Region *[]FirewallCountryRegion `json:"region,omitempty"`
}

type FirewallCountryRegion struct {
	Id *float64 `json:"id,omitempty"`
}
