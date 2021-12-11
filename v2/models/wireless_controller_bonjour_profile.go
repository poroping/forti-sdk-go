package models

const WirelessControllerBonjourProfilePath = "wireless-controller/bonjour-profile/"

type WirelessControllerBonjourProfile struct {
	Comment    *string                                      `json:"comment,omitempty"`
	Name       *string                                      `json:"name,omitempty"`
	PolicyList []WirelessControllerBonjourProfilePolicyList `json:"policy-list,omitempty"`
}

type WirelessControllerBonjourProfilePolicyList struct {
	Description *string `json:"description,omitempty"`
	FromVlan    *string `json:"from-vlan,omitempty"`
	PolicyId    *int64  `json:"policy-id,omitempty"`
	Services    *string `json:"services,omitempty"`
	ToVlan      *string `json:"to-vlan,omitempty"`
}
