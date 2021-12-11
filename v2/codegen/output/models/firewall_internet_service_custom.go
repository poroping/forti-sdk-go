package models

const FirewallInternetServiceCustomPath = "firewall/internet-service-custom/"

type FirewallInternetServiceCustom struct {
	Comment    *string                              `json:"comment,omitempty"`
	Entry      []FirewallInternetServiceCustomEntry `json:"entry,omitempty"`
	Name       *string                              `json:"name,omitempty"`
	Reputation *int64                               `json:"reputation,omitempty"`
}

type FirewallInternetServiceCustomEntry struct {
	Dst       []FirewallInternetServiceCustomEntryDst       `json:"dst,omitempty"`
	Id        *int64                                        `json:"id,omitempty"`
	PortRange []FirewallInternetServiceCustomEntryPortRange `json:"port-range,omitempty"`
	Protocol  *int64                                        `json:"protocol,omitempty"`
}

type FirewallInternetServiceCustomEntryDst struct {
	Name *string `json:"name,omitempty"`
}

type FirewallInternetServiceCustomEntryPortRange struct {
	EndPort   *int64 `json:"end-port,omitempty"`
	Id        *int64 `json:"id,omitempty"`
	StartPort *int64 `json:"start-port,omitempty"`
}
