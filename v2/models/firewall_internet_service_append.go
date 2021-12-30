package models

const FirewallInternetServiceAppendPath = "firewall/internet-service-append/"

type FirewallInternetServiceAppend struct {
	AppendPort *int64 `json:"append-port,omitempty"`
	MatchPort  *int64 `json:"match-port,omitempty"`
}
