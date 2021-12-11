package models

const FirewallInternetServiceReputationPath = "firewall/internet-service-reputation/"

type FirewallInternetServiceReputation struct {
	Description *string `json:"description,omitempty"`
	Fosid       *int64  `json:"fosid,omitempty"`
}
