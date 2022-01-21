package models

const SystemNdProxyPath = "system/nd-proxy/"

type SystemNdProxy struct {
	Member *[]SystemNdProxyMember `json:"member,omitempty"`
	Status *string                `json:"status,omitempty"`
}

type SystemNdProxyMember struct {
	InterfaceName *string `json:"interface-name,omitempty"`
}
