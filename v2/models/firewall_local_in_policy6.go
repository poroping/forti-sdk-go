package models

const FirewallLocalInPolicy6Path = "firewall/local-in-policy6/"

type FirewallLocalInPolicy6 struct {
	Action        *string                          `json:"action,omitempty"`
	Comments      *string                          `json:"comments,omitempty"`
	Dstaddr       *[]FirewallLocalInPolicy6Dstaddr `json:"dstaddr,omitempty"`
	DstaddrNegate *string                          `json:"dstaddr-negate,omitempty"`
	Intf          *string                          `json:"intf,omitempty"`
	Policyid      *float64                         `json:"policyid,omitempty"`
	Schedule      *string                          `json:"schedule,omitempty"`
	Service       *[]FirewallLocalInPolicy6Service `json:"service,omitempty"`
	ServiceNegate *string                          `json:"service-negate,omitempty"`
	Srcaddr       *[]FirewallLocalInPolicy6Srcaddr `json:"srcaddr,omitempty"`
	SrcaddrNegate *string                          `json:"srcaddr-negate,omitempty"`
	Status        *string                          `json:"status,omitempty"`
	Uuid          *string                          `json:"uuid,omitempty"`
}

type FirewallLocalInPolicy6Dstaddr struct {
	Name *string `json:"name,omitempty"`
}

type FirewallLocalInPolicy6Service struct {
	Name *string `json:"name,omitempty"`
}

type FirewallLocalInPolicy6Srcaddr struct {
	Name *string `json:"name,omitempty"`
}
