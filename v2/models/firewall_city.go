package models

const FirewallCityPath = "firewall/city/"

type FirewallCity struct {
	Id   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// Set FirewallCity values to defaults
func (def *FirewallCity) Defaults() {
	def.Id = intPtr(0)
	def.Name = stringPtr("")
}
