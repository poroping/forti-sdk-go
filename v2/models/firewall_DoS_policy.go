package models

const FirewallDosPolicyPath = "firewall/dos-policy/"

type FirewallDosPolicy struct {
	Anomaly   *[]FirewallDosPolicyAnomaly `json:"anomaly,omitempty"`
	Comments  *string                     `json:"comments,omitempty"`
	Dstaddr   *[]FirewallDosPolicyDstaddr `json:"dstaddr,omitempty"`
	Interface *string                     `json:"interface,omitempty"`
	Name      *string                     `json:"name,omitempty"`
	Policyid  *int64                      `json:"policyid,omitempty"`
	Service   *[]FirewallDosPolicyService `json:"service,omitempty"`
	Srcaddr   *[]FirewallDosPolicySrcaddr `json:"srcaddr,omitempty"`
	Status    *string                     `json:"status,omitempty"`
}

const FirewallDosPolicyAnomalyPath = "firewall/dos-policy/anomaly/"

type FirewallDosPolicyAnomaly struct {
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

const FirewallDosPolicyDstaddrPath = "firewall/dos-policy/dstaddr/"

type FirewallDosPolicyDstaddr struct {
	Name *string `json:"name,omitempty"`
}

const FirewallDosPolicyServicePath = "firewall/dos-policy/service/"

type FirewallDosPolicyService struct {
	Name *string `json:"name,omitempty"`
}

const FirewallDosPolicySrcaddrPath = "firewall/dos-policy/srcaddr/"

type FirewallDosPolicySrcaddr struct {
	Name *string `json:"name,omitempty"`
}
