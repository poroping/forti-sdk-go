package models

const FirewallSshSettingPath = "firewall.ssh/setting/"

type FirewallSshSetting struct {
	Caname              *string `json:"caname,omitempty"`
	HostTrustedChecking *string `json:"host-trusted-checking,omitempty"`
	HostkeyDsa1024      *string `json:"hostkey-dsa1024,omitempty"`
	HostkeyEcdsa256     *string `json:"hostkey-ecdsa256,omitempty"`
	HostkeyEcdsa384     *string `json:"hostkey-ecdsa384,omitempty"`
	HostkeyEcdsa521     *string `json:"hostkey-ecdsa521,omitempty"`
	HostkeyEd25519      *string `json:"hostkey-ed25519,omitempty"`
	HostkeyRsa2048      *string `json:"hostkey-rsa2048,omitempty"`
	UntrustedCaname     *string `json:"untrusted-caname,omitempty"`
}
