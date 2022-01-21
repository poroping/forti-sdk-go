package models

const FirewallVipgrp6Path = "firewall/vipgrp6/"

type FirewallVipgrp6 struct {
	Color    *int64                   `json:"color,omitempty"`
	Comments *string                  `json:"comments,omitempty"`
	Member   *[]FirewallVipgrp6Member `json:"member,omitempty"`
	Name     *string                  `json:"name,omitempty"`
	Uuid     *string                  `json:"uuid,omitempty"`
}

type FirewallVipgrp6Member struct {
	Name *string `json:"name,omitempty"`
}
