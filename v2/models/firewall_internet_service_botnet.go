package models

const FirewallInternetServiceBotnetPath = "firewall/internet-service-botnet/"

type FirewallInternetServiceBotnet struct {
	Id   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// Set FirewallInternetServiceBotnet values to defaults
func (def *FirewallInternetServiceBotnet) Defaults() {
	def.Id = intPtr(0)
	def.Name = stringPtr("")
}
