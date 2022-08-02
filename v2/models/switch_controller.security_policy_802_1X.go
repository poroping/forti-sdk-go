package models

const SwitchControllerSecurityPolicy8021XPath = "switch-controller.security-policy/802-1x/"

type SwitchControllerSecurityPolicy8021X struct {
	AuthFailVlan            *string                                         `json:"auth-fail-vlan,omitempty"`
	AuthFailVlanId          *string                                         `json:"auth-fail-vlan-id,omitempty"`
	AuthserverTimeoutPeriod *int64                                          `json:"authserver-timeout-period,omitempty"`
	AuthserverTimeoutVlan   *string                                         `json:"authserver-timeout-vlan,omitempty"`
	AuthserverTimeoutVlanid *string                                         `json:"authserver-timeout-vlanid,omitempty"`
	EapAutoUntaggedVlans    *string                                         `json:"eap-auto-untagged-vlans,omitempty"`
	EapPassthru             *string                                         `json:"eap-passthru,omitempty"`
	FramevidApply           *string                                         `json:"framevid-apply,omitempty"`
	GuestAuthDelay          *int64                                          `json:"guest-auth-delay,omitempty"`
	GuestVlan               *string                                         `json:"guest-vlan,omitempty"`
	GuestVlanId             *string                                         `json:"guest-vlan-id,omitempty"`
	MacAuthBypass           *string                                         `json:"mac-auth-bypass,omitempty"`
	Name                    *string                                         `json:"name,omitempty"`
	OpenAuth                *string                                         `json:"open-auth,omitempty"`
	PolicyType              *string                                         `json:"policy-type,omitempty"`
	RadiusTimeoutOverwrite  *string                                         `json:"radius-timeout-overwrite,omitempty"`
	SecurityMode            *string                                         `json:"security-mode,omitempty"`
	UserGroup               *[]SwitchControllerSecurityPolicy8021XUserGroup `json:"user-group,omitempty"`
}

const SwitchControllerSecurityPolicy8021XUserGroupPath = "switch-controller.security-policy/802-1x/user-group/"

type SwitchControllerSecurityPolicy8021XUserGroup struct {
	Name *string `json:"name,omitempty"`
}
