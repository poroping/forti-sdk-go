package models

const FirewallInternetServiceDefinitionPath = "firewall/internet-service-definition/"

type FirewallInternetServiceDefinition struct {
	Entry *[]FirewallInternetServiceDefinitionEntry `json:"entry,omitempty"`
	Fosid *float64                                  `json:"fosid,omitempty"`
}

type FirewallInternetServiceDefinitionEntry struct {
	CategoryId *float64                                           `json:"category-id,omitempty"`
	Name       *string                                            `json:"name,omitempty"`
	PortRange  *[]FirewallInternetServiceDefinitionEntryPortRange `json:"port-range,omitempty"`
	Protocol   *float64                                           `json:"protocol,omitempty"`
	SeqNum     *float64                                           `json:"seq-num,omitempty"`
}

type FirewallInternetServiceDefinitionEntryPortRange struct {
	EndPort   *float64 `json:"end-port,omitempty"`
	Id        *float64 `json:"id,omitempty"`
	StartPort *float64 `json:"start-port,omitempty"`
}
