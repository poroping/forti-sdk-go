package models

const FirewallIpmacbindingSettingPath = "firewall.ipmacbinding/setting/"

type FirewallIpmacbindingSetting struct {
	Bindthroughfw *string `json:"bindthroughfw,omitempty"`
	Bindtofw      *string `json:"bindtofw,omitempty"`
	Undefinedhost *string `json:"undefinedhost,omitempty"`
}
