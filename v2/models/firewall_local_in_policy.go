package models

const FirewallLocalInPolicyPath = "firewall/local-in-policy/"

type FirewallLocalInPolicy struct {
	Action         *string                         `json:"action,omitempty"`
	Comments       *string                         `json:"comments,omitempty"`
	Dstaddr        *[]FirewallLocalInPolicyDstaddr `json:"dstaddr,omitempty"`
	DstaddrNegate  *string                         `json:"dstaddr-negate,omitempty"`
	HaMgmtIntfOnly *string                         `json:"ha-mgmt-intf-only,omitempty"`
	Intf           *string                         `json:"intf,omitempty"`
	Policyid       *int64                          `json:"policyid,omitempty"`
	Schedule       *string                         `json:"schedule,omitempty"`
	Service        *[]FirewallLocalInPolicyService `json:"service,omitempty"`
	ServiceNegate  *string                         `json:"service-negate,omitempty"`
	Srcaddr        *[]FirewallLocalInPolicySrcaddr `json:"srcaddr,omitempty"`
	SrcaddrNegate  *string                         `json:"srcaddr-negate,omitempty"`
	Status         *string                         `json:"status,omitempty"`
	Uuid           *string                         `json:"uuid,omitempty"`
}

const FirewallLocalInPolicyDstaddrPath = "firewall/local-in-policy/dstaddr/"

type FirewallLocalInPolicyDstaddr struct {
	Name *string `json:"name,omitempty"`
}

const FirewallLocalInPolicyServicePath = "firewall/local-in-policy/service/"

type FirewallLocalInPolicyService struct {
	Name *string `json:"name,omitempty"`
}

const FirewallLocalInPolicySrcaddrPath = "firewall/local-in-policy/srcaddr/"

type FirewallLocalInPolicySrcaddr struct {
	Name *string `json:"name,omitempty"`
}
