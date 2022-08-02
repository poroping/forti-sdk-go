package models

const FirewallInternetServiceIpblVendorPath = "firewall/internet-service-ipbl-vendor/"

type FirewallInternetServiceIpblVendor struct {
	Id   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// Set FirewallInternetServiceIpblVendor values to defaults
func (def *FirewallInternetServiceIpblVendor) Defaults() {
	def.Id = intPtr(0)
	def.Name = stringPtr("")
}
