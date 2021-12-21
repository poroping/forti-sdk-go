package models

const FirewallMulticastAddressPath = "firewall/multicast-address/"

type FirewallMulticastAddress struct {
	AssociatedInterface *string                            `json:"associated-interface,omitempty"`
	Color               *float64                           `json:"color,omitempty"`
	Comment             *string                            `json:"comment,omitempty"`
	EndIp               *string                            `json:"end-ip,omitempty"`
	Name                *string                            `json:"name,omitempty"`
	StartIp             *string                            `json:"start-ip,omitempty"`
	Subnet              *string                            `json:"subnet,omitempty"`
	Tagging             *[]FirewallMulticastAddressTagging `json:"tagging,omitempty"`
	Type                *string                            `json:"type,omitempty"`
	Visibility          *string                            `json:"visibility,omitempty"`
}

type FirewallMulticastAddressTagging struct {
	Category *string                                `json:"category,omitempty"`
	Name     *string                                `json:"name,omitempty"`
	Tags     *[]FirewallMulticastAddressTaggingTags `json:"tags,omitempty"`
}

type FirewallMulticastAddressTaggingTags struct {
	Name *string `json:"name,omitempty"`
}
