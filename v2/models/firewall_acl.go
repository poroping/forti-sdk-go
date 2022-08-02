package models

const FirewallAclPath = "firewall/acl/"

type FirewallAcl struct {
	Comments  *string               `json:"comments,omitempty"`
	Dstaddr   *[]FirewallAclDstaddr `json:"dstaddr,omitempty"`
	Interface *string               `json:"interface,omitempty"`
	Name      *string               `json:"name,omitempty"`
	Policyid  *int64                `json:"policyid,omitempty"`
	Service   *[]FirewallAclService `json:"service,omitempty"`
	Srcaddr   *[]FirewallAclSrcaddr `json:"srcaddr,omitempty"`
	Status    *string               `json:"status,omitempty"`
}

const FirewallAclDstaddrPath = "firewall/acl/dstaddr/"

type FirewallAclDstaddr struct {
	Name *string `json:"name,omitempty"`
}

const FirewallAclServicePath = "firewall/acl/service/"

type FirewallAclService struct {
	Name *string `json:"name,omitempty"`
}

const FirewallAclSrcaddrPath = "firewall/acl/srcaddr/"

type FirewallAclSrcaddr struct {
	Name *string `json:"name,omitempty"`
}
