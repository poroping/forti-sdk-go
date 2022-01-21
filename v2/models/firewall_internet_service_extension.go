package models

const FirewallInternetServiceExtensionPath = "firewall/internet-service-extension/"

type FirewallInternetServiceExtension struct {
	Comment      *string                                         `json:"comment,omitempty"`
	DisableEntry *[]FirewallInternetServiceExtensionDisableEntry `json:"disable-entry,omitempty"`
	Entry        *[]FirewallInternetServiceExtensionEntry        `json:"entry,omitempty"`
	Id           *int64                                          `json:"id,omitempty"`
}

type FirewallInternetServiceExtensionDisableEntry struct {
	Id        *int64                                                   `json:"id,omitempty"`
	IpRange   *[]FirewallInternetServiceExtensionDisableEntryIpRange   `json:"ip-range,omitempty"`
	PortRange *[]FirewallInternetServiceExtensionDisableEntryPortRange `json:"port-range,omitempty"`
	Protocol  *int64                                                   `json:"protocol,omitempty"`
}

type FirewallInternetServiceExtensionDisableEntryIpRange struct {
	EndIp   *string `json:"end-ip,omitempty"`
	Id      *int64  `json:"id,omitempty"`
	StartIp *string `json:"start-ip,omitempty"`
}

type FirewallInternetServiceExtensionDisableEntryPortRange struct {
	EndPort   *int64 `json:"end-port,omitempty"`
	Id        *int64 `json:"id,omitempty"`
	StartPort *int64 `json:"start-port,omitempty"`
}

type FirewallInternetServiceExtensionEntry struct {
	Dst       *[]FirewallInternetServiceExtensionEntryDst       `json:"dst,omitempty"`
	Id        *int64                                            `json:"id,omitempty"`
	PortRange *[]FirewallInternetServiceExtensionEntryPortRange `json:"port-range,omitempty"`
	Protocol  *int64                                            `json:"protocol,omitempty"`
}

type FirewallInternetServiceExtensionEntryDst struct {
	Name *string `json:"name,omitempty"`
}

type FirewallInternetServiceExtensionEntryPortRange struct {
	EndPort   *int64 `json:"end-port,omitempty"`
	Id        *int64 `json:"id,omitempty"`
	StartPort *int64 `json:"start-port,omitempty"`
}
