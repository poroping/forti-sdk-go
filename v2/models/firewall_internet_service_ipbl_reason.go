package models

const FirewallInternetServiceIpblReasonPath = "firewall/internet-service-ipbl-reason/"

type FirewallInternetServiceIpblReason struct {
	Id   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// Set FirewallInternetServiceIpblReason values to defaults
func (def *FirewallInternetServiceIpblReason) Defaults() {
	def.Id = intPtr(0)
	def.Name = stringPtr("")
}
