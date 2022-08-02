package models

const FirewallDnstranslationPath = "firewall/dnstranslation/"

type FirewallDnstranslation struct {
	Dst     *string `json:"dst,omitempty"`
	Id      *int64  `json:"id,omitempty"`
	Netmask *string `json:"netmask,omitempty"`
	Src     *string `json:"src,omitempty"`
}
