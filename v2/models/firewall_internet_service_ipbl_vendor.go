package models

const FirewallInternetServiceIpblVendorPath = "firewall/internet-service-ipbl-vendor/"

type FirewallInternetServiceIpblVendor struct {
	Id   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

//defaultfuncs
func (def *FirewallInternetServiceIpblVendor) defaults() {
	def.Id = "0"
	def.Name = ""
}
