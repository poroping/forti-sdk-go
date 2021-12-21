package models

const FirewallInternetServiceReputationPath = "firewall/internet-service-reputation/"

type FirewallInternetServiceReputation struct {
	Description *string  `json:"description,omitempty"`
	Fosid       *float64 `json:"fosid,omitempty"`
}
