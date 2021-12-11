package models

const FirewallsshLocalKeyPath = "firewall.ssh/local-key/"

type FirewallsshLocalKey struct {
	Name       *string `json:"name,omitempty"`
	Password   *string `json:"password,omitempty"`
	PrivateKey *string `json:"private-key,omitempty"`
	PublicKey  *string `json:"public-key,omitempty"`
	Source     *string `json:"source,omitempty"`
}
