package models

const FirewallwildcardFqdnGroupPath = "firewall.wildcard-fqdn/group/"

type FirewallwildcardFqdnGroup struct {
	Color   *int64                            `json:"color,omitempty"`
	Comment *string                           `json:"comment,omitempty"`
	Member  []FirewallwildcardFqdnGroupMember `json:"member,omitempty"`
	Name    *string                           `json:"name,omitempty"`
	Uuid    *string                           `json:"uuid,omitempty"`
}

type FirewallwildcardFqdnGroupMember struct {
	Name *string `json:"name,omitempty"`
}
