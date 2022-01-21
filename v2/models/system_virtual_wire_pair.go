package models

const SystemVirtualWirePairPath = "system/virtual-wire-pair/"

type SystemVirtualWirePair struct {
	Member       *[]SystemVirtualWirePairMember `json:"member,omitempty"`
	Name         *string                        `json:"name,omitempty"`
	VlanFilter   *string                        `json:"vlan-filter,omitempty"`
	WildcardVlan *string                        `json:"wildcard-vlan,omitempty"`
}

type SystemVirtualWirePairMember struct {
	InterfaceName *string `json:"interface-name,omitempty"`
}
