package models

const FirewallipmacbindingSettingPath = "firewall.ipmacbinding/setting/"

type FirewallipmacbindingSetting struct {
	Bindthroughfw *string `json:"bindthroughfw,omitempty"`
	Bindtofw      *string `json:"bindtofw,omitempty"`
	Undefinedhost *string `json:"undefinedhost,omitempty"`
}
