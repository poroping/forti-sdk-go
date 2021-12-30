package models

const FirewallVendorMacPath = "firewall/vendor-mac/"

type FirewallVendorMac struct {
	Fosid     *int64  `json:"fosid,omitempty"`
	MacNumber *int64  `json:"mac-number,omitempty"`
	Name      *string `json:"name,omitempty"`
	Obsolete  *int64  `json:"obsolete,omitempty"`
}
