package models

const FirewallCityPath = "firewall/city/"

type FirewallCity struct {
	Id   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

//defaultfuncs
func (def *FirewallCity) defaults() {
	def.Id = "0"
	def.Name = ""
}
