package models

const UserNacPolicyPath = "user/nac-policy/"

type UserNacPolicy struct {
	Category         *string                     `json:"category,omitempty"`
	Description      *string                     `json:"description,omitempty"`
	EmsTag           *string                     `json:"ems-tag,omitempty"`
	Family           *string                     `json:"family,omitempty"`
	FirewallAddress  *string                     `json:"firewall-address,omitempty"`
	Host             *string                     `json:"host,omitempty"`
	HwVendor         *string                     `json:"hw-vendor,omitempty"`
	HwVersion        *string                     `json:"hw-version,omitempty"`
	Mac              *string                     `json:"mac,omitempty"`
	Name             *string                     `json:"name,omitempty"`
	Os               *string                     `json:"os,omitempty"`
	Src              *string                     `json:"src,omitempty"`
	SsidPolicy       *string                     `json:"ssid-policy,omitempty"`
	Status           *string                     `json:"status,omitempty"`
	SwVersion        *string                     `json:"sw-version,omitempty"`
	SwitchAutoAuth   *string                     `json:"switch-auto-auth,omitempty"`
	SwitchFortilink  *string                     `json:"switch-fortilink,omitempty"`
	SwitchGroup      *[]UserNacPolicySwitchGroup `json:"switch-group,omitempty"`
	SwitchMacPolicy  *string                     `json:"switch-mac-policy,omitempty"`
	SwitchPortPolicy *string                     `json:"switch-port-policy,omitempty"`
	SwitchScope      *[]UserNacPolicySwitchScope `json:"switch-scope,omitempty"`
	Type             *string                     `json:"type,omitempty"`
	User             *string                     `json:"user,omitempty"`
	UserGroup        *string                     `json:"user-group,omitempty"`
}

type UserNacPolicySwitchGroup struct {
	Name *string `json:"name,omitempty"`
}

type UserNacPolicySwitchScope struct {
	SwitchId *string `json:"switch-id,omitempty"`
}
