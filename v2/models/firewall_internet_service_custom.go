package models

const FirewallInternetServiceCustomPath = "firewall/internet-service-custom/"

type FirewallInternetServiceCustom struct {
	Comment    *string                               `json:"comment,omitempty"`
	Entry      *[]FirewallInternetServiceCustomEntry `json:"entry,omitempty"`
	Name       *string                               `json:"name,omitempty"`
	Reputation *float64                              `json:"reputation,omitempty"`
}

type FirewallInternetServiceCustomEntry struct {
	Dst       *[]FirewallInternetServiceCustomEntryDst       `json:"dst,omitempty"`
	Id        *float64                                       `json:"id,omitempty"`
	PortRange *[]FirewallInternetServiceCustomEntryPortRange `json:"port-range,omitempty"`
	Protocol  *float64                                       `json:"protocol,omitempty"`
}

type FirewallInternetServiceCustomEntryDst struct {
	Name *string `json:"name,omitempty"`
}

type FirewallInternetServiceCustomEntryPortRange struct {
	EndPort   *float64 `json:"end-port,omitempty"`
	Id        *float64 `json:"id,omitempty"`
	StartPort *float64 `json:"start-port,omitempty"`
}
