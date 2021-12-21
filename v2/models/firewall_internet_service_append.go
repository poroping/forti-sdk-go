package models

const FirewallInternetServiceAppendPath = "firewall/internet-service-append/"

type FirewallInternetServiceAppend struct {
	AppendPort *float64 `json:"append-port,omitempty"`
	MatchPort  *float64 `json:"match-port,omitempty"`
}
