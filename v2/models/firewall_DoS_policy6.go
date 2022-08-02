package models

const FirewallDosPolicy6Path = "firewall/dos-policy6/"

type FirewallDosPolicy6 struct {
	Anomaly   *[]FirewallDosPolicy6Anomaly `json:"anomaly,omitempty"`
	Comments  *string                      `json:"comments,omitempty"`
	Dstaddr   *[]FirewallDosPolicy6Dstaddr `json:"dstaddr,omitempty"`
	Interface *string                      `json:"interface,omitempty"`
	Name      *string                      `json:"name,omitempty"`
	Policyid  *int64                       `json:"policyid,omitempty"`
	Service   *[]FirewallDosPolicy6Service `json:"service,omitempty"`
	Srcaddr   *[]FirewallDosPolicy6Srcaddr `json:"srcaddr,omitempty"`
	Status    *string                      `json:"status,omitempty"`
}

const FirewallDosPolicy6AnomalyPath = "firewall/dos-policy6/anomaly/"

type FirewallDosPolicy6Anomaly struct {
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

const FirewallDosPolicy6DstaddrPath = "firewall/dos-policy6/dstaddr/"

type FirewallDosPolicy6Dstaddr struct {
	Name *string `json:"name,omitempty"`
}

const FirewallDosPolicy6ServicePath = "firewall/dos-policy6/service/"

type FirewallDosPolicy6Service struct {
	Name *string `json:"name,omitempty"`
}

const FirewallDosPolicy6SrcaddrPath = "firewall/dos-policy6/srcaddr/"

type FirewallDosPolicy6Srcaddr struct {
	Name *string `json:"name,omitempty"`
}
