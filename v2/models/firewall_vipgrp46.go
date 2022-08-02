package models

const FirewallVipgrp46Path = "firewall/vipgrp46/"

type FirewallVipgrp46 struct {
	Color    *int64                    `json:"color,omitempty"`
	Comments *string                   `json:"comments,omitempty"`
	Member   *[]FirewallVipgrp46Member `json:"member,omitempty"`
	Name     *string                   `json:"name,omitempty"`
	Uuid     *string                   `json:"uuid,omitempty"`
}

const FirewallVipgrp46MemberPath = "firewall/vipgrp46/member/"

type FirewallVipgrp46Member struct {
	Name *string `json:"name,omitempty"`
}
