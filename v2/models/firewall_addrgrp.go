package models

const FirewallAddrgrpPath = "firewall/addrgrp/"

type FirewallAddrgrp struct {
	AllowRouting  *string                         `json:"allow-routing,omitempty"`
	Category      *string                         `json:"category,omitempty"`
	Color         *int64                          `json:"color,omitempty"`
	Comment       *string                         `json:"comment,omitempty"`
	Exclude       *string                         `json:"exclude,omitempty"`
	ExcludeMember *[]FirewallAddrgrpExcludeMember `json:"exclude-member,omitempty"`
	FabricObject  *string                         `json:"fabric-object,omitempty"`
	Member        *[]FirewallAddrgrpMember        `json:"member,omitempty"`
	Name          *string                         `json:"name,omitempty"`
	Tagging       *[]FirewallAddrgrpTagging       `json:"tagging,omitempty"`
	Type          *string                         `json:"type,omitempty"`
	Uuid          *string                         `json:"uuid,omitempty"`
	Visibility    *string                         `json:"visibility,omitempty"`
}

type FirewallAddrgrpExcludeMember struct {
	Name *string `json:"name,omitempty"`
}

type FirewallAddrgrpMember struct {
	Name *string `json:"name,omitempty"`
}

type FirewallAddrgrpTagging struct {
	Category *string                       `json:"category,omitempty"`
	Name     *string                       `json:"name,omitempty"`
	Tags     *[]FirewallAddrgrpTaggingTags `json:"tags,omitempty"`
}

type FirewallAddrgrpTaggingTags struct {
	Name *string `json:"name,omitempty"`
}
