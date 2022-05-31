package models

const FirewallInternetServiceIpblReasonPath = "firewall/internet-service-ipbl-reason/"

type FirewallInternetServiceIpblReason struct {
	Id   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

//defaultfuncs
func (def *FirewallInternetServiceIpblReason) defaults() {
	def.Id = "0"
	def.Name = ""
}
