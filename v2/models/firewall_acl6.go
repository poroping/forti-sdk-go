package models

const FirewallAcl6Path = "firewall/acl6/"

type FirewallAcl6 struct {
	Comments  *string                `json:"comments,omitempty"`
	Dstaddr   *[]FirewallAcl6Dstaddr `json:"dstaddr,omitempty"`
	Interface *string                `json:"interface,omitempty"`
	Name      *string                `json:"name,omitempty"`
	Policyid  *int64                 `json:"policyid,omitempty"`
	Service   *[]FirewallAcl6Service `json:"service,omitempty"`
	Srcaddr   *[]FirewallAcl6Srcaddr `json:"srcaddr,omitempty"`
	Status    *string                `json:"status,omitempty"`
}

const FirewallAcl6DstaddrPath = "firewall/acl6/dstaddr/"

type FirewallAcl6Dstaddr struct {
	Name *string `json:"name,omitempty"`
}

const FirewallAcl6ServicePath = "firewall/acl6/service/"

type FirewallAcl6Service struct {
	Name *string `json:"name,omitempty"`
}

const FirewallAcl6SrcaddrPath = "firewall/acl6/srcaddr/"

type FirewallAcl6Srcaddr struct {
	Name *string `json:"name,omitempty"`
}
