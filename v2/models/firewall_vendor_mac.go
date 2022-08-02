package models

const FirewallVendorMacPath = "firewall/vendor-mac/"

type FirewallVendorMac struct {
	Id        *int64  `json:"id,omitempty"`
	MacNumber *int64  `json:"mac-number,omitempty"`
	Name      *string `json:"name,omitempty"`
	Obsolete  *int64  `json:"obsolete,omitempty"`
}

// Set FirewallVendorMac values to defaults
func (def *FirewallVendorMac) Defaults() {
	def.Id = intPtr(0)
	def.MacNumber = intPtr(0)
	def.Name = stringPtr("")
	def.Obsolete = intPtr(0)
}
