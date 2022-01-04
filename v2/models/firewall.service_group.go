package models

const FirewallServiceGroupPath = "firewall.service/group/"

type FirewallServiceGroup struct {
	Color        *int64                        `json:"color,omitempty"`
	Comment      *string                       `json:"comment,omitempty"`
	FabricObject *string                       `json:"fabric-object,omitempty"`
	Member       *[]FirewallServiceGroupMember `json:"member,omitempty"`
	Name         *string                       `json:"name,omitempty"`
	Proxy        *string                       `json:"proxy,omitempty"`
}

type FirewallServiceGroupMember struct {
	Name *string `json:"name,omitempty"`
}
