package models

const FirewallRegionPath = "firewall/region/"

type FirewallRegion struct {
	City *[]FirewallRegionCity `json:"city,omitempty"`
	Id   *int64                `json:"id,omitempty"`
	Name *string               `json:"name,omitempty"`
}

const FirewallRegionCityPath = "firewall/region/city/"

type FirewallRegionCity struct {
	Id *int64 `json:"id,omitempty"`
}

// Set FirewallRegion values to defaults
func (def *FirewallRegion) Defaults() {
	def.Id = intPtr(0)
	def.Name = stringPtr("")
}
