package models

const FirewallInternetServiceAdditionPath = "firewall/internet-service-addition/"

type FirewallInternetServiceAddition struct {
	Comment *string                                 `json:"comment,omitempty"`
	Entry   *[]FirewallInternetServiceAdditionEntry `json:"entry,omitempty"`
	Fosid   *float64                                `json:"fosid,omitempty"`
}

type FirewallInternetServiceAdditionEntry struct {
	Id        *float64                                         `json:"id,omitempty"`
	PortRange *[]FirewallInternetServiceAdditionEntryPortRange `json:"port-range,omitempty"`
	Protocol  *float64                                         `json:"protocol,omitempty"`
}

type FirewallInternetServiceAdditionEntryPortRange struct {
	EndPort   *float64 `json:"end-port,omitempty"`
	Id        *float64 `json:"id,omitempty"`
	StartPort *float64 `json:"start-port,omitempty"`
}
