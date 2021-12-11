package models

const FirewallIppoolPath = "firewall/ippool/"

type FirewallIppool struct {
	AddNat64Route       *string `json:"add-nat64-route,omitempty"`
	ArpIntf             *string `json:"arp-intf,omitempty"`
	ArpReply            *string `json:"arp-reply,omitempty"`
	AssociatedInterface *string `json:"associated-interface,omitempty"`
	BlockSize           *int64  `json:"block-size,omitempty"`
	Comments            *string `json:"comments,omitempty"`
	Endip               *string `json:"endip,omitempty"`
	Endport             *int64  `json:"endport,omitempty"`
	Name                *string `json:"name,omitempty"`
	Nat64               *string `json:"nat64,omitempty"`
	NumBlocksPerUser    *int64  `json:"num-blocks-per-user,omitempty"`
	PbaTimeout          *int64  `json:"pba-timeout,omitempty"`
	PermitAnyHost       *string `json:"permit-any-host,omitempty"`
	PortPerUser         *int64  `json:"port-per-user,omitempty"`
	SourceEndip         *string `json:"source-endip,omitempty"`
	SourceStartip       *string `json:"source-startip,omitempty"`
	Startip             *string `json:"startip,omitempty"`
	Startport           *int64  `json:"startport,omitempty"`
	Type                *string `json:"type,omitempty"`
}
