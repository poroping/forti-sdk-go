package models

const FirewallInternetServiceBotnetPath = "firewall/internet-service-botnet/"

type FirewallInternetServiceBotnet struct {
	Fosid *int64  `json:"fosid,omitempty"`
	Name  *string `json:"name,omitempty"`
}
