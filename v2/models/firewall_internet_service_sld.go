package models

const FirewallInternetServiceSldPath = "firewall/internet-service-sld/"

type FirewallInternetServiceSld struct {
	Id   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// Set FirewallInternetServiceSld values to defaults
func (def *FirewallInternetServiceSld) Defaults() {
	def.Id = intPtr(0)
	def.Name = stringPtr("")
}
