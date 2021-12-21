package models

const FirewallInternetServiceListPath = "firewall/internet-service-list/"

type FirewallInternetServiceList struct {
	Fosid *float64 `json:"fosid,omitempty"`
	Name  *string  `json:"name,omitempty"`
}
