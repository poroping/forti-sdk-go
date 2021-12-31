package models

const FirewallVendorMacPath = "firewall/vendor-mac/"

type FirewallVendorMac struct {
	Id        *int64  `json:"id,omitempty"`
	MacNumber *int64  `json:"mac-number,omitempty"`
	Name      *string `json:"name,omitempty"`
	Obsolete  *int64  `json:"obsolete,omitempty"`
}
