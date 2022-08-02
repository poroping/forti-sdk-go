package models

const FirewallMulticastAddressPath = "firewall/multicast-address/"

type FirewallMulticastAddress struct {
	AssociatedInterface *string                            `json:"associated-interface,omitempty"`
	Color               *int64                             `json:"color,omitempty"`
	Comment             *string                            `json:"comment,omitempty"`
	EndIp               *string                            `json:"end-ip,omitempty"`
	Name                *string                            `json:"name,omitempty"`
	StartIp             *string                            `json:"start-ip,omitempty"`
	Subnet              *string                            `json:"subnet,omitempty"`
	Tagging             *[]FirewallMulticastAddressTagging `json:"tagging,omitempty"`
	Type                *string                            `json:"type,omitempty"`
}

const FirewallMulticastAddressTaggingPath = "firewall/multicast-address/tagging/"

type FirewallMulticastAddressTagging struct {
	Category *string                                `json:"category,omitempty"`
	Name     *string                                `json:"name,omitempty"`
	Tags     *[]FirewallMulticastAddressTaggingTags `json:"tags,omitempty"`
}

const FirewallMulticastAddressTaggingTagsPath = "firewall/multicast-address/tagging/tags/"

type FirewallMulticastAddressTaggingTags struct {
	Name *string `json:"name,omitempty"`
}
