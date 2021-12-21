package models

const FirewallVipgrp64Path = "firewall/vipgrp64/"

type FirewallVipgrp64 struct {
	Color    *float64                  `json:"color,omitempty"`
	Comments *string                   `json:"comments,omitempty"`
	Member   *[]FirewallVipgrp64Member `json:"member,omitempty"`
	Name     *string                   `json:"name,omitempty"`
	Uuid     *string                   `json:"uuid,omitempty"`
}

type FirewallVipgrp64Member struct {
	Name *string `json:"name,omitempty"`
}
