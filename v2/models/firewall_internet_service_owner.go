package models

const FirewallInternetServiceOwnerPath = "firewall/internet-service-owner/"

type FirewallInternetServiceOwner struct {
	Id   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// Set FirewallInternetServiceOwner values to defaults
func (def *FirewallInternetServiceOwner) Defaults() {
	def.Id = intPtr(0)
	def.Name = stringPtr("")
}
