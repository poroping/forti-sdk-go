package models

const FirewallInternetServiceIpblReasonPath = "firewall/internet-service-ipbl-reason/"

type FirewallInternetServiceIpblReason struct {
	Fosid *float64 `json:"fosid,omitempty"`
	Name  *string  `json:"name,omitempty"`
}
