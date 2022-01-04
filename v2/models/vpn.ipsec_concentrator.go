package models

const VpnIpsecConcentratorPath = "vpn.ipsec/concentrator/"

type VpnIpsecConcentrator struct {
	Id       *int64                        `json:"id,omitempty"`
	Member   *[]VpnIpsecConcentratorMember `json:"member,omitempty"`
	Name     *string                       `json:"name,omitempty"`
	SrcCheck *string                       `json:"src-check,omitempty"`
}

type VpnIpsecConcentratorMember struct {
	Name *string `json:"name,omitempty"`
}
