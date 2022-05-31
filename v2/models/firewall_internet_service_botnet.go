package models

const FirewallInternetServiceBotnetPath = "firewall/internet-service-botnet/"

type FirewallInternetServiceBotnet struct {
	Id   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

//defaultfuncs
func (def *FirewallInternetServiceBotnet) defaults() {
	def.Id = "0"
	def.Name = ""
}
