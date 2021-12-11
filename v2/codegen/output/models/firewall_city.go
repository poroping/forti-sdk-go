package models

const FirewallCityPath = "firewall/city/"

type FirewallCity struct {
	Fosid *int64  `json:"fosid,omitempty"`
	Name  *string `json:"name,omitempty"`
}
