package models

const FirewallMulticastAddress6Path = "firewall/multicast-address6/"

type FirewallMulticastAddress6 struct {
	Color   *int64                             `json:"color,omitempty"`
	Comment *string                            `json:"comment,omitempty"`
	Ip6     *string                            `json:"ip6,omitempty"`
	Name    *string                            `json:"name,omitempty"`
	Tagging []FirewallMulticastAddress6Tagging `json:"tagging,omitempty"`
}

type FirewallMulticastAddress6Tagging struct {
	Category *string                                `json:"category,omitempty"`
	Name     *string                                `json:"name,omitempty"`
	Tags     []FirewallMulticastAddress6TaggingTags `json:"tags,omitempty"`
}

type FirewallMulticastAddress6TaggingTags struct {
	Name *string `json:"name,omitempty"`
}
