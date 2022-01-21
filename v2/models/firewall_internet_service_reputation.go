package models

const FirewallInternetServiceReputationPath = "firewall/internet-service-reputation/"

type FirewallInternetServiceReputation struct {
	Description *string `json:"description,omitempty"`
	Id          *int64  `json:"id,omitempty"`
}
