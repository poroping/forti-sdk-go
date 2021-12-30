package models

const FirewallDoSPolicy6Path = "firewall/DoS-policy6/"

type FirewallDoSPolicy6 struct {
	Anomaly   *[]FirewallDoSPolicy6Anomaly `json:"anomaly,omitempty"`
	Comments  *string                      `json:"comments,omitempty"`
	Dstaddr   *[]FirewallDoSPolicy6Dstaddr `json:"dstaddr,omitempty"`
	Interface *string                      `json:"interface,omitempty"`
	Name      *string                      `json:"name,omitempty"`
	Policyid  *int64                       `json:"policyid,omitempty"`
	Service   *[]FirewallDoSPolicy6Service `json:"service,omitempty"`
	Srcaddr   *[]FirewallDoSPolicy6Srcaddr `json:"srcaddr,omitempty"`
	Status    *string                      `json:"status,omitempty"`
}

type FirewallDoSPolicy6Anomaly struct {
	Action           *string `json:"action,omitempty"`
	Log              *string `json:"log,omitempty"`
	Name             *string `json:"name,omitempty"`
	Quarantine       *string `json:"quarantine,omitempty"`
	QuarantineExpiry *string `json:"quarantine-expiry,omitempty"`
	QuarantineLog    *string `json:"quarantine-log,omitempty"`
	Status           *string `json:"status,omitempty"`
	Threshold        *int64  `json:"threshold,omitempty"`
	Thresholddefault *int64  `json:"threshold(default),omitempty"`
}

type FirewallDoSPolicy6Dstaddr struct {
	Name *string `json:"name,omitempty"`
}

type FirewallDoSPolicy6Service struct {
	Name *string `json:"name,omitempty"`
}

type FirewallDoSPolicy6Srcaddr struct {
	Name *string `json:"name,omitempty"`
}
