package models

const FirewallProutePath = "firewall/proute/"

type FirewallProute struct {
	PolicyRouteId *string `json:"<policy route id>,omitempty"`
}
