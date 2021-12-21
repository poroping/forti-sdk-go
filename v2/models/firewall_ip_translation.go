package models

const FirewallIpTranslationPath = "firewall/ip-translation/"

type FirewallIpTranslation struct {
	Endip      *string  `json:"endip,omitempty"`
	MapStartip *string  `json:"map-startip,omitempty"`
	Startip    *string  `json:"startip,omitempty"`
	Transid    *float64 `json:"transid,omitempty"`
	Type       *string  `json:"type,omitempty"`
}
