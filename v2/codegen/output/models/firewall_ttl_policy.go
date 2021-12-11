package models

const FirewallTtlPolicyPath = "firewall/ttl-policy/"

type FirewallTtlPolicy struct {
	Action   *string                    `json:"action,omitempty"`
	Fosid    *int64                     `json:"fosid,omitempty"`
	Schedule *string                    `json:"schedule,omitempty"`
	Service  []FirewallTtlPolicyService `json:"service,omitempty"`
	Srcaddr  []FirewallTtlPolicySrcaddr `json:"srcaddr,omitempty"`
	Srcintf  *string                    `json:"srcintf,omitempty"`
	Status   *string                    `json:"status,omitempty"`
	Ttl      *string                    `json:"ttl,omitempty"`
}

type FirewallTtlPolicyService struct {
	Name *string `json:"name,omitempty"`
}

type FirewallTtlPolicySrcaddr struct {
	Name *string `json:"name,omitempty"`
}
