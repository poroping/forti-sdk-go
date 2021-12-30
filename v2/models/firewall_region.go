package models

const FirewallRegionPath = "firewall/region/"

type FirewallRegion struct {
	City  *[]FirewallRegionCity `json:"city,omitempty"`
	Fosid *int64                `json:"fosid,omitempty"`
	Name  *string               `json:"name,omitempty"`
}

type FirewallRegionCity struct {
	Id *int64 `json:"id,omitempty"`
}
