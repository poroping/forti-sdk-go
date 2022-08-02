package models

const FirewallAddrgrp6Path = "firewall/addrgrp6/"

type FirewallAddrgrp6 struct {
	Color        *int64                     `json:"color,omitempty"`
	Comment      *string                    `json:"comment,omitempty"`
	FabricObject *string                    `json:"fabric-object,omitempty"`
	Member       *[]FirewallAddrgrp6Member  `json:"member,omitempty"`
	Name         *string                    `json:"name,omitempty"`
	Tagging      *[]FirewallAddrgrp6Tagging `json:"tagging,omitempty"`
	Uuid         *string                    `json:"uuid,omitempty"`
}

const FirewallAddrgrp6MemberPath = "firewall/addrgrp6/member/"

type FirewallAddrgrp6Member struct {
	Name *string `json:"name,omitempty"`
}

const FirewallAddrgrp6TaggingPath = "firewall/addrgrp6/tagging/"

type FirewallAddrgrp6Tagging struct {
	Category *string                        `json:"category,omitempty"`
	Name     *string                        `json:"name,omitempty"`
	Tags     *[]FirewallAddrgrp6TaggingTags `json:"tags,omitempty"`
}

const FirewallAddrgrp6TaggingTagsPath = "firewall/addrgrp6/tagging/tags/"

type FirewallAddrgrp6TaggingTags struct {
	Name *string `json:"name,omitempty"`
}
