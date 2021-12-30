package models

const FirewallInternetServiceDefinitionPath = "firewall/internet-service-definition/"

type FirewallInternetServiceDefinition struct {
	Entry *[]FirewallInternetServiceDefinitionEntry `json:"entry,omitempty"`
	Fosid *int64                                    `json:"fosid,omitempty"`
}

type FirewallInternetServiceDefinitionEntry struct {
	CategoryId *int64                                             `json:"category-id,omitempty"`
	Name       *string                                            `json:"name,omitempty"`
	PortRange  *[]FirewallInternetServiceDefinitionEntryPortRange `json:"port-range,omitempty"`
	Protocol   *int64                                             `json:"protocol,omitempty"`
	SeqNum     *int64                                             `json:"seq-num,omitempty"`
}

type FirewallInternetServiceDefinitionEntryPortRange struct {
	EndPort   *int64 `json:"end-port,omitempty"`
	Id        *int64 `json:"id,omitempty"`
	StartPort *int64 `json:"start-port,omitempty"`
}
