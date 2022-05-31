package models

const FirewallInternetServiceListPath = "firewall/internet-service-list/"

type FirewallInternetServiceList struct {
	Id   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

//defaultfuncs
func (def *FirewallInternetServiceList) defaults() {
	def.Id = "0"
	def.Name = ""
}
