package models

const FirewallIppool6Path = "firewall/ippool6/"

type FirewallIppool6 struct {
	AddNat46Route *string `json:"add-nat46-route,omitempty"`
	Comments      *string `json:"comments,omitempty"`
	Endip         *string `json:"endip,omitempty"`
	Name          *string `json:"name,omitempty"`
	Nat46         *string `json:"nat46,omitempty"`
	Startip       *string `json:"startip,omitempty"`
}
