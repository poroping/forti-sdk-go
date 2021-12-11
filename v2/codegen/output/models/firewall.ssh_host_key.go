package models

const FirewallsshHostKeyPath = "firewall.ssh/host-key/"

type FirewallsshHostKey struct {
	Hostname  *string `json:"hostname,omitempty"`
	Ip        *string `json:"ip,omitempty"`
	Name      *string `json:"name,omitempty"`
	Nid       *string `json:"nid,omitempty"`
	Port      *int64  `json:"port,omitempty"`
	PublicKey *string `json:"public-key,omitempty"`
	Status    *string `json:"status,omitempty"`
	Type      *string `json:"type,omitempty"`
	Usage     *string `json:"usage,omitempty"`
}
