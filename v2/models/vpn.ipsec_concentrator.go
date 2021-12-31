package models

const VpnipsecConcentratorPath = "vpn/ipsec/concentrator/"

type VpnipsecConcentrator struct {
	Id       *int64                        `json:"id,omitempty"`
	Member   *[]VpnipsecConcentratorMember `json:"member,omitempty"`
	Name     *string                       `json:"name,omitempty"`
	SrcCheck *string                       `json:"src-check,omitempty"`
}

type VpnipsecConcentratorMember struct {
	Name *string `json:"name,omitempty"`
}
