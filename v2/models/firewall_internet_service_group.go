package models

const FirewallInternetServiceGroupPath = "firewall/internet-service-group/"

type FirewallInternetServiceGroup struct {
	Comment   *string                               `json:"comment,omitempty"`
	Direction *string                               `json:"direction,omitempty"`
	Member    *[]FirewallInternetServiceGroupMember `json:"member,omitempty"`
	Name      *string                               `json:"name,omitempty"`
}

type FirewallInternetServiceGroupMember struct {
	Id   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}
