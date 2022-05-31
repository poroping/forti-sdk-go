package models

const FirewallInternetServiceSldPath = "firewall/internet-service-sld/"

type FirewallInternetServiceSld struct {
	Id   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

//defaultfuncs
func (def *FirewallInternetServiceSld) defaults() {
	def.Id = "0"
	def.Name = ""
}
