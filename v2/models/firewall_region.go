package models

const FirewallRegionPath = "firewall/region/"

type FirewallRegion struct {
	City *[]FirewallRegionCity `json:"city,omitempty"`
	Id   *int64                `json:"id,omitempty"`
	Name *string               `json:"name,omitempty"`
}

type FirewallRegionCity struct {
	Id *int64 `json:"id,omitempty"`
}

//defaultfuncs
func (def *FirewallRegion) defaults() {
	def.Id = "0"
	def.Name = ""
}
