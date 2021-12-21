package models

const FirewallCityPath = "firewall/city/"

type FirewallCity struct {
	Fosid *float64 `json:"fosid,omitempty"`
	Name  *string  `json:"name,omitempty"`
}
