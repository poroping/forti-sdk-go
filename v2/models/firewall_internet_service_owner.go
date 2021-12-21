package models

const FirewallInternetServiceOwnerPath = "firewall/internet-service-owner/"

type FirewallInternetServiceOwner struct {
	Fosid *float64 `json:"fosid,omitempty"`
	Name  *string  `json:"name,omitempty"`
}
