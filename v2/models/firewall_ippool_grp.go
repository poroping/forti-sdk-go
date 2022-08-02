package models

const FirewallIppool_grpPath = "firewall/ippool_grp/"

type FirewallIppool_grp struct {
	Comments *string                     `json:"comments,omitempty"`
	Member   *[]FirewallIppool_grpMember `json:"member,omitempty"`
	Name     *string                     `json:"name,omitempty"`
	Uuid     *string                     `json:"uuid,omitempty"`
}

const FirewallIppool_grpMemberPath = "firewall/ippool_grp/member/"

type FirewallIppool_grpMember struct {
	Name *string `json:"name,omitempty"`
}
