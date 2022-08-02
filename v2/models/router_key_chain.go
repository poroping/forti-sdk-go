package models

const RouterKeyChainPath = "router/key-chain/"

type RouterKeyChain struct {
	Key  *[]RouterKeyChainKey `json:"key,omitempty"`
	Name *string              `json:"name,omitempty"`
}

const RouterKeyChainKeyPath = "router/key-chain/key/"

type RouterKeyChainKey struct {
	AcceptLifetime *string `json:"accept-lifetime,omitempty"`
	Algorithm      *string `json:"algorithm,omitempty"`
	Id             *string `json:"id,omitempty"`
	KeyString      *string `json:"key-string,omitempty"`
	SendLifetime   *string `json:"send-lifetime,omitempty"`
}
