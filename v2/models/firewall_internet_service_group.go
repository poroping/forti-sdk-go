package models

const FirewallInternetServiceGroupPath = "firewall/internet-service-group/"

type FirewallInternetServiceGroup struct {
	Comment   *string                               `json:"comment,omitempty"`
	Direction *string                               `json:"direction,omitempty"`
	Member    *[]FirewallInternetServiceGroupMember `json:"member,omitempty"`
	Name      *string                               `json:"name,omitempty"`
}

const FirewallInternetServiceGroupMemberPath = "firewall/internet-service-group/member/"

type FirewallInternetServiceGroupMember struct {
	Name *string `json:"name,omitempty"`
}
