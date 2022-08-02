package models

const FirewallInternetServiceAdditionPath = "firewall/internet-service-addition/"

type FirewallInternetServiceAddition struct {
	Comment *string                                 `json:"comment,omitempty"`
	Entry   *[]FirewallInternetServiceAdditionEntry `json:"entry,omitempty"`
	Id      *int64                                  `json:"id,omitempty"`
}

const FirewallInternetServiceAdditionEntryPath = "firewall/internet-service-addition/entry/"

type FirewallInternetServiceAdditionEntry struct {
	Id        *int64                                           `json:"id,omitempty"`
	PortRange *[]FirewallInternetServiceAdditionEntryPortRange `json:"port-range,omitempty"`
	Protocol  *int64                                           `json:"protocol,omitempty"`
}

const FirewallInternetServiceAdditionEntryPortRangePath = "firewall/internet-service-addition/entry/port-range/"

type FirewallInternetServiceAdditionEntryPortRange struct {
	EndPort   *int64 `json:"end-port,omitempty"`
	Id        *int64 `json:"id,omitempty"`
	StartPort *int64 `json:"start-port,omitempty"`
}
