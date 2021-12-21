package models

const WanoptAuthGroupPath = "wanopt/auth-group/"

type WanoptAuthGroup struct {
	AuthMethod *string `json:"auth-method,omitempty"`
	Cert       *string `json:"cert,omitempty"`
	Name       *string `json:"name,omitempty"`
	Peer       *string `json:"peer,omitempty"`
	PeerAccept *string `json:"peer-accept,omitempty"`
	Psk        *string `json:"psk,omitempty"`
}
