package models

const FirewallserviceGroupPath = "firewall/service/group/"

type FirewallserviceGroup struct {
	Color        *int64                        `json:"color,omitempty"`
	Comment      *string                       `json:"comment,omitempty"`
	FabricObject *string                       `json:"fabric-object,omitempty"`
	Member       *[]FirewallserviceGroupMember `json:"member,omitempty"`
	Name         *string                       `json:"name,omitempty"`
	Proxy        *string                       `json:"proxy,omitempty"`
}

type FirewallserviceGroupMember struct {
	Name *string `json:"name,omitempty"`
}
