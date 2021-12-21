package models

const FirewallProxyAddrgrpPath = "firewall/proxy-addrgrp/"

type FirewallProxyAddrgrp struct {
	Color      *float64                       `json:"color,omitempty"`
	Comment    *string                        `json:"comment,omitempty"`
	Member     *[]FirewallProxyAddrgrpMember  `json:"member,omitempty"`
	Name       *string                        `json:"name,omitempty"`
	Tagging    *[]FirewallProxyAddrgrpTagging `json:"tagging,omitempty"`
	Type       *string                        `json:"type,omitempty"`
	Uuid       *string                        `json:"uuid,omitempty"`
	Visibility *string                        `json:"visibility,omitempty"`
}

type FirewallProxyAddrgrpMember struct {
	Name *string `json:"name,omitempty"`
}

type FirewallProxyAddrgrpTagging struct {
	Category *string                            `json:"category,omitempty"`
	Name     *string                            `json:"name,omitempty"`
	Tags     *[]FirewallProxyAddrgrpTaggingTags `json:"tags,omitempty"`
}

type FirewallProxyAddrgrpTaggingTags struct {
	Name *string `json:"name,omitempty"`
}
