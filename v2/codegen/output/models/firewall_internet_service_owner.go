package models

const FirewallInternetServiceOwnerPath = "firewall/internet-service-owner/"

type FirewallInternetServiceOwner struct {
	Fosid *int64  `json:"fosid,omitempty"`
	Name  *string `json:"name,omitempty"`
}
