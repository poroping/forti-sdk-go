package models

const SwitchControllersecurityPolicy8021XPath = "switch-controller.security-policy/802-1X/"

type SwitchControllersecurityPolicy8021X struct {
	AuthFailVlan            *string                                         `json:"auth-fail-vlan,omitempty"`
	AuthFailVlanId          *string                                         `json:"auth-fail-vlan-id,omitempty"`
	AuthserverTimeoutPeriod *float64                                        `json:"authserver-timeout-period,omitempty"`
	AuthserverTimeoutVlan   *string                                         `json:"authserver-timeout-vlan,omitempty"`
	AuthserverTimeoutVlanid *string                                         `json:"authserver-timeout-vlanid,omitempty"`
	EapAutoUntaggedVlans    *string                                         `json:"eap-auto-untagged-vlans,omitempty"`
	EapPassthru             *string                                         `json:"eap-passthru,omitempty"`
	FramevidApply           *string                                         `json:"framevid-apply,omitempty"`
	GuestAuthDelay          *float64                                        `json:"guest-auth-delay,omitempty"`
	GuestVlan               *string                                         `json:"guest-vlan,omitempty"`
	GuestVlanId             *string                                         `json:"guest-vlan-id,omitempty"`
	MacAuthBypass           *string                                         `json:"mac-auth-bypass,omitempty"`
	Name                    *string                                         `json:"name,omitempty"`
	OpenAuth                *string                                         `json:"open-auth,omitempty"`
	PolicyType              *string                                         `json:"policy-type,omitempty"`
	RadiusTimeoutOverwrite  *string                                         `json:"radius-timeout-overwrite,omitempty"`
	SecurityMode            *string                                         `json:"security-mode,omitempty"`
	UserGroup               *[]SwitchControllersecurityPolicy8021XUserGroup `json:"user-group,omitempty"`
}

type SwitchControllersecurityPolicy8021XUserGroup struct {
	Name *string `json:"name,omitempty"`
}
