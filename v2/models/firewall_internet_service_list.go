package models

const FirewallInternetServiceListPath = "firewall/internet-service-list/"

type FirewallInternetServiceList struct {
	Fosid *int64  `json:"fosid,omitempty"`
	Name  *string `json:"name,omitempty"`
}
