package models

const FirewallInternetServiceExtensionPath = "firewall/internet-service-extension/"

type FirewallInternetServiceExtension struct {
	Comment      *string                                         `json:"comment,omitempty"`
	DisableEntry *[]FirewallInternetServiceExtensionDisableEntry `json:"disable-entry,omitempty"`
	Entry        *[]FirewallInternetServiceExtensionEntry        `json:"entry,omitempty"`
	Id           *int64                                          `json:"id,omitempty"`
}

type FirewallInternetServiceExtensionDisableEntry struct {
	AddrMode  *string                                                  `json:"addr-mode,omitempty"`
	Id        *int64                                                   `json:"id,omitempty"`
	IpRange   *[]FirewallInternetServiceExtensionDisableEntryIpRange   `json:"ip-range,omitempty"`
	Ip6Range  *[]FirewallInternetServiceExtensionDisableEntryIp6Range  `json:"ip6-range,omitempty"`
	PortRange *[]FirewallInternetServiceExtensionDisableEntryPortRange `json:"port-range,omitempty"`
	Protocol  *int64                                                   `json:"protocol,omitempty"`
}

type FirewallInternetServiceExtensionDisableEntryIpRange struct {
	EndIp   *string `json:"end-ip,omitempty"`
	Id      *int64  `json:"id,omitempty"`
	StartIp *string `json:"start-ip,omitempty"`
}

type FirewallInternetServiceExtensionDisableEntryIp6Range struct {
	EndIp6   *string `json:"end-ip6,omitempty"`
	Id       *int64  `json:"id,omitempty"`
	StartIp6 *string `json:"start-ip6,omitempty"`
}

type FirewallInternetServiceExtensionDisableEntryPortRange struct {
	EndPort   *int64 `json:"end-port,omitempty"`
	Id        *int64 `json:"id,omitempty"`
	StartPort *int64 `json:"start-port,omitempty"`
}

type FirewallInternetServiceExtensionEntry struct {
	AddrMode  *string                                           `json:"addr-mode,omitempty"`
	Dst       *[]FirewallInternetServiceExtensionEntryDst       `json:"dst,omitempty"`
	Dst6      *[]FirewallInternetServiceExtensionEntryDst6      `json:"dst6,omitempty"`
	Id        *int64                                            `json:"id,omitempty"`
	PortRange *[]FirewallInternetServiceExtensionEntryPortRange `json:"port-range,omitempty"`
	Protocol  *int64                                            `json:"protocol,omitempty"`
}

type FirewallInternetServiceExtensionEntryDst struct {
	Name *string `json:"name,omitempty"`
}

type FirewallInternetServiceExtensionEntryDst6 struct {
	Name *string `json:"name,omitempty"`
}

type FirewallInternetServiceExtensionEntryPortRange struct {
	EndPort   *int64 `json:"end-port,omitempty"`
	Id        *int64 `json:"id,omitempty"`
	StartPort *int64 `json:"start-port,omitempty"`
}
