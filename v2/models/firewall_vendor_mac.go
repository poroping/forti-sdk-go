package models

const FirewallVendorMacPath = "firewall/vendor-mac/"

type FirewallVendorMac struct {
	Fosid     *float64 `json:"fosid,omitempty"`
	MacNumber *float64 `json:"mac-number,omitempty"`
	Name      *string  `json:"name,omitempty"`
	Obsolete  *float64 `json:"obsolete,omitempty"`
}
