package models

const VpnipsecForticlientPath = "vpn.ipsec/forticlient/"

type VpnipsecForticlient struct {
	Phase2name    *string `json:"phase2name,omitempty"`
	Realm         *string `json:"realm,omitempty"`
	Status        *string `json:"status,omitempty"`
	Usergroupname *string `json:"usergroupname,omitempty"`
}
