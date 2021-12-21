package models

const FirewallDnstranslationPath = "firewall/dnstranslation/"

type FirewallDnstranslation struct {
	Dst     *string `json:"dst,omitempty"`
	Fosid   *int64  `json:"fosid,omitempty"`
	Netmask *string `json:"netmask,omitempty"`
	Src     *string `json:"src,omitempty"`
}