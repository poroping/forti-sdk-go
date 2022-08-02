package models

const FirewallWildcardFqdnCustomPath = "firewall.wildcard-fqdn/custom/"

type FirewallWildcardFqdnCustom struct {
	Color        *int64  `json:"color,omitempty"`
	Comment      *string `json:"comment,omitempty"`
	Name         *string `json:"name,omitempty"`
	Uuid         *string `json:"uuid,omitempty"`
	WildcardFqdn *string `json:"wildcard-fqdn,omitempty"`
}
