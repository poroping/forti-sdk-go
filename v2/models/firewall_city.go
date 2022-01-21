package models

const FirewallCityPath = "firewall/city/"

type FirewallCity struct {
	Id   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}
