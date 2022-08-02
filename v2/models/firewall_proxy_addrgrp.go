package models

const FirewallProxyAddrgrpPath = "firewall/proxy-addrgrp/"

type FirewallProxyAddrgrp struct {
	Color   *int64                         `json:"color,omitempty"`
	Comment *string                        `json:"comment,omitempty"`
	Member  *[]FirewallProxyAddrgrpMember  `json:"member,omitempty"`
	Name    *string                        `json:"name,omitempty"`
	Tagging *[]FirewallProxyAddrgrpTagging `json:"tagging,omitempty"`
	Type    *string                        `json:"type,omitempty"`
	Uuid    *string                        `json:"uuid,omitempty"`
}

const FirewallProxyAddrgrpMemberPath = "firewall/proxy-addrgrp/member/"

type FirewallProxyAddrgrpMember struct {
	Name *string `json:"name,omitempty"`
}

const FirewallProxyAddrgrpTaggingPath = "firewall/proxy-addrgrp/tagging/"

type FirewallProxyAddrgrpTagging struct {
	Category *string                            `json:"category,omitempty"`
	Name     *string                            `json:"name,omitempty"`
	Tags     *[]FirewallProxyAddrgrpTaggingTags `json:"tags,omitempty"`
}

const FirewallProxyAddrgrpTaggingTagsPath = "firewall/proxy-addrgrp/tagging/tags/"

type FirewallProxyAddrgrpTaggingTags struct {
	Name *string `json:"name,omitempty"`
}
