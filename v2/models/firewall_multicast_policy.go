package models

const FirewallMulticastPolicyPath = "firewall/multicast-policy/"

type FirewallMulticastPolicy struct {
	Action          *string                          `json:"action,omitempty"`
	AutoAsicOffload *string                          `json:"auto-asic-offload,omitempty"`
	Comments        *string                          `json:"comments,omitempty"`
	Dnat            *string                          `json:"dnat,omitempty"`
	Dstaddr         []FirewallMulticastPolicyDstaddr `json:"dstaddr,omitempty"`
	Dstintf         *string                          `json:"dstintf,omitempty"`
	EndPort         *int64                           `json:"end-port,omitempty"`
	Fosid           *int64                           `json:"fosid,omitempty"`
	Logtraffic      *string                          `json:"logtraffic,omitempty"`
	Name            *string                          `json:"name,omitempty"`
	Protocol        *int64                           `json:"protocol,omitempty"`
	Snat            *string                          `json:"snat,omitempty"`
	SnatIp          *string                          `json:"snat-ip,omitempty"`
	Srcaddr         []FirewallMulticastPolicySrcaddr `json:"srcaddr,omitempty"`
	Srcintf         *string                          `json:"srcintf,omitempty"`
	StartPort       *int64                           `json:"start-port,omitempty"`
	Status          *string                          `json:"status,omitempty"`
	Uuid            *string                          `json:"uuid,omitempty"`
}

type FirewallMulticastPolicyDstaddr struct {
	Name *string `json:"name,omitempty"`
}

type FirewallMulticastPolicySrcaddr struct {
	Name *string `json:"name,omitempty"`
}
