package models

const FirewallVipgrpPath = "firewall/vipgrp/"

type FirewallVipgrp struct {
	Color     *int64                  `json:"color,omitempty"`
	Comments  *string                 `json:"comments,omitempty"`
	Interface *string                 `json:"interface,omitempty"`
	Member    *[]FirewallVipgrpMember `json:"member,omitempty"`
	Name      *string                 `json:"name,omitempty"`
	Uuid      *string                 `json:"uuid,omitempty"`
}

const FirewallVipgrpMemberPath = "firewall/vipgrp/member/"

type FirewallVipgrpMember struct {
	Name *string `json:"name,omitempty"`
}
