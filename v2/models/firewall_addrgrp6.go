package models

const FirewallAddrgrp6Path = "firewall/addrgrp6/"

type FirewallAddrgrp6 struct {
	Color        *float64                   `json:"color,omitempty"`
	Comment      *string                    `json:"comment,omitempty"`
	FabricObject *string                    `json:"fabric-object,omitempty"`
	Member       *[]FirewallAddrgrp6Member  `json:"member,omitempty"`
	Name         *string                    `json:"name,omitempty"`
	Tagging      *[]FirewallAddrgrp6Tagging `json:"tagging,omitempty"`
	Uuid         *string                    `json:"uuid,omitempty"`
	Visibility   *string                    `json:"visibility,omitempty"`
}

type FirewallAddrgrp6Member struct {
	Name *string `json:"name,omitempty"`
}

type FirewallAddrgrp6Tagging struct {
	Category *string                        `json:"category,omitempty"`
	Name     *string                        `json:"name,omitempty"`
	Tags     *[]FirewallAddrgrp6TaggingTags `json:"tags,omitempty"`
}

type FirewallAddrgrp6TaggingTags struct {
	Name *string `json:"name,omitempty"`
}
