package models

const FirewallCountryPath = "firewall/country/"

type FirewallCountry struct {
	Id     *int64                   `json:"id,omitempty"`
	Name   *string                  `json:"name,omitempty"`
	Region *[]FirewallCountryRegion `json:"region,omitempty"`
}

const FirewallCountryRegionPath = "firewall/country/region/"

type FirewallCountryRegion struct {
	Id *int64 `json:"id,omitempty"`
}

// Set FirewallCountry values to defaults
func (def *FirewallCountry) Defaults() {
	def.Id = intPtr(0)
	def.Name = stringPtr("")
}
