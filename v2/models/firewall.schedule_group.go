package models

const FirewallScheduleGroupPath = "firewall.schedule/group/"

type FirewallScheduleGroup struct {
	Color        *int64                         `json:"color,omitempty"`
	FabricObject *string                        `json:"fabric-object,omitempty"`
	Member       *[]FirewallScheduleGroupMember `json:"member,omitempty"`
	Name         *string                        `json:"name,omitempty"`
}

type FirewallScheduleGroupMember struct {
	Name *string `json:"name,omitempty"`
}
