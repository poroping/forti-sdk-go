package models

const FirewallVendorMacPath = "firewall/vendor-mac/"

type FirewallVendorMac struct {
	Id        *int64  `json:"id,omitempty"`
	MacNumber *int64  `json:"mac-number,omitempty"`
	Name      *string `json:"name,omitempty"`
	Obsolete  *int64  `json:"obsolete,omitempty"`
}

//defaultfuncs
func (def *FirewallVendorMac) defaults() {
	def.Id = "0"
	def.MacNumber = "0"
	def.Name = ""
	def.Obsolete = "0"
}
