package models

const FirewallDoSPolicyPath = "firewall/DoS-policy/"

type FirewallDoSPolicy struct {
	Anomaly   *[]FirewallDoSPolicyAnomaly `json:"anomaly,omitempty"`
	Comments  *string                     `json:"comments,omitempty"`
	Dstaddr   *[]FirewallDoSPolicyDstaddr `json:"dstaddr,omitempty"`
	Interface *string                     `json:"interface,omitempty"`
	Name      *string                     `json:"name,omitempty"`
	Policyid  *int64                      `json:"policyid,omitempty"`
	Service   *[]FirewallDoSPolicyService `json:"service,omitempty"`
	Srcaddr   *[]FirewallDoSPolicySrcaddr `json:"srcaddr,omitempty"`
	Status    *string                     `json:"status,omitempty"`
}

type FirewallDoSPolicyAnomaly struct {
	Action                 *string `json:"action,omitempty"`
	Log                    *string `json:"log,omitempty"`
	Name                   *string `json:"name,omitempty"`
	Quarantine             *string `json:"quarantine,omitempty"`
	QuarantineExpiry       *string `json:"quarantine-expiry,omitempty"`
	QuarantineLog          *string `json:"quarantine-log,omitempty"`
	Status                 *string `json:"status,omitempty"`
	SynproxyTcpMss         *string `json:"synproxy-tcp-mss,omitempty"`
	SynproxyTcpSack        *string `json:"synproxy-tcp-sack,omitempty"`
	SynproxyTcpTimestamp   *string `json:"synproxy-tcp-timestamp,omitempty"`
	SynproxyTcpWindow      *string `json:"synproxy-tcp-window,omitempty"`
	SynproxyTcpWindowscale *string `json:"synproxy-tcp-windowscale,omitempty"`
	SynproxyTos            *string `json:"synproxy-tos,omitempty"`
	SynproxyTtl            *string `json:"synproxy-ttl,omitempty"`
	Threshold              *int64  `json:"threshold,omitempty"`
	Thresholddefault       *int64  `json:"threshold(default),omitempty"`
}

type FirewallDoSPolicyDstaddr struct {
	Name *string `json:"name,omitempty"`
}

type FirewallDoSPolicyService struct {
	Name *string `json:"name,omitempty"`
}

type FirewallDoSPolicySrcaddr struct {
	Name *string `json:"name,omitempty"`
}
