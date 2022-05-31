package models

const FirewallInternetServiceOwnerPath = "firewall/internet-service-owner/"

type FirewallInternetServiceOwner struct {
	Id   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

//defaultfuncs
func (def *FirewallInternetServiceOwner) defaults() {
	def.Id = "0"
	def.Name = ""
}
