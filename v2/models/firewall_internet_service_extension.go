package models

const FirewallInternetServiceExtensionPath = "firewall/internet-service-extension/"

type FirewallInternetServiceExtension struct {
	Comment      *string                                         `json:"comment,omitempty"`
	DisableEntry *[]FirewallInternetServiceExtensionDisableEntry `json:"disable-entry,omitempty"`
	Entry        *[]FirewallInternetServiceExtensionEntry        `json:"entry,omitempty"`
	Fosid        *float64                                        `json:"fosid,omitempty"`
}

type FirewallInternetServiceExtensionDisableEntry struct {
	Id        *float64                                                 `json:"id,omitempty"`
	IpRange   *[]FirewallInternetServiceExtensionDisableEntryIpRange   `json:"ip-range,omitempty"`
	PortRange *[]FirewallInternetServiceExtensionDisableEntryPortRange `json:"port-range,omitempty"`
	Protocol  *float64                                                 `json:"protocol,omitempty"`
}

type FirewallInternetServiceExtensionDisableEntryIpRange struct {
	EndIp   *string  `json:"end-ip,omitempty"`
	Id      *float64 `json:"id,omitempty"`
	StartIp *string  `json:"start-ip,omitempty"`
}

type FirewallInternetServiceExtensionDisableEntryPortRange struct {
	EndPort   *float64 `json:"end-port,omitempty"`
	Id        *float64 `json:"id,omitempty"`
	StartPort *float64 `json:"start-port,omitempty"`
}

type FirewallInternetServiceExtensionEntry struct {
	Dst       *[]FirewallInternetServiceExtensionEntryDst       `json:"dst,omitempty"`
	Id        *float64                                          `json:"id,omitempty"`
	PortRange *[]FirewallInternetServiceExtensionEntryPortRange `json:"port-range,omitempty"`
	Protocol  *float64                                          `json:"protocol,omitempty"`
}

type FirewallInternetServiceExtensionEntryDst struct {
	Name *string `json:"name,omitempty"`
}

type FirewallInternetServiceExtensionEntryPortRange struct {
	EndPort   *float64 `json:"end-port,omitempty"`
	Id        *float64 `json:"id,omitempty"`
	StartPort *float64 `json:"start-port,omitempty"`
}
