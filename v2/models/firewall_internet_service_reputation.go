package models

const FirewallInternetServiceReputationPath = "firewall/internet-service-reputation/"

type FirewallInternetServiceReputation struct {
	Description *string `json:"description,omitempty"`
	Id          *int64  `json:"id,omitempty"`
}

// Set FirewallInternetServiceReputation values to defaults
func (def *FirewallInternetServiceReputation) Defaults() {
	def.Description = stringPtr("")
	def.Id = intPtr(0)
}
