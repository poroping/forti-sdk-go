package models

const FirewallTtlPolicyPath = "firewall/ttl-policy/"

type FirewallTtlPolicy struct {
	Action   *string                     `json:"action,omitempty"`
	Id       *int64                      `json:"id,omitempty"`
	Schedule *string                     `json:"schedule,omitempty"`
	Service  *[]FirewallTtlPolicyService `json:"service,omitempty"`
	Srcaddr  *[]FirewallTtlPolicySrcaddr `json:"srcaddr,omitempty"`
	Srcintf  *string                     `json:"srcintf,omitempty"`
	Status   *string                     `json:"status,omitempty"`
	Ttl      *string                     `json:"ttl,omitempty"`
}

const FirewallTtlPolicyServicePath = "firewall/ttl-policy/service/"

type FirewallTtlPolicyService struct {
	Name *string `json:"name,omitempty"`
}

const FirewallTtlPolicySrcaddrPath = "firewall/ttl-policy/srcaddr/"

type FirewallTtlPolicySrcaddr struct {
	Name *string `json:"name,omitempty"`
}
