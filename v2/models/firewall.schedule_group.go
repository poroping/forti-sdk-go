package models

const FirewallscheduleGroupPath = "firewall.schedule/group/"

type FirewallscheduleGroup struct {
	Color        *int64                         `json:"color,omitempty"`
	FabricObject *string                        `json:"fabric-object,omitempty"`
	Member       *[]FirewallscheduleGroupMember `json:"member,omitempty"`
	Name         *string                        `json:"name,omitempty"`
}

type FirewallscheduleGroupMember struct {
	Name *string `json:"name,omitempty"`
}
