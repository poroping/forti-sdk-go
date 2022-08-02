package models

const FirewallMulticastAddress6Path = "firewall/multicast-address6/"

type FirewallMulticastAddress6 struct {
	Color   *int64                              `json:"color,omitempty"`
	Comment *string                             `json:"comment,omitempty"`
	Ip6     *string                             `json:"ip6,omitempty"`
	Name    *string                             `json:"name,omitempty"`
	Tagging *[]FirewallMulticastAddress6Tagging `json:"tagging,omitempty"`
}

const FirewallMulticastAddress6TaggingPath = "firewall/multicast-address6/tagging/"

type FirewallMulticastAddress6Tagging struct {
	Category *string                                 `json:"category,omitempty"`
	Name     *string                                 `json:"name,omitempty"`
	Tags     *[]FirewallMulticastAddress6TaggingTags `json:"tags,omitempty"`
}

const FirewallMulticastAddress6TaggingTagsPath = "firewall/multicast-address6/tagging/tags/"

type FirewallMulticastAddress6TaggingTags struct {
	Name *string `json:"name,omitempty"`
}
