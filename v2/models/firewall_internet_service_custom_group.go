package models

const FirewallInternetServiceCustomGroupPath = "firewall/internet-service-custom-group/"

type FirewallInternetServiceCustomGroup struct {
	Comment *string                                    `json:"comment,omitempty"`
	Member  []FirewallInternetServiceCustomGroupMember `json:"member,omitempty"`
	Name    *string                                    `json:"name,omitempty"`
}

type FirewallInternetServiceCustomGroupMember struct {
	Name *string `json:"name,omitempty"`
}
