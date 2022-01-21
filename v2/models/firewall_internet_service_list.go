package models

const FirewallInternetServiceListPath = "firewall/internet-service-list/"

type FirewallInternetServiceList struct {
	Id   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}
