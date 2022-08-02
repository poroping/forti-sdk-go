package models

const FirewallInternetServiceExtensionPath = "firewall/internet-service-extension/"

type FirewallInternetServiceExtension struct {
	Comment      *string                                         `json:"comment,omitempty"`
	DisableEntry *[]FirewallInternetServiceExtensionDisableEntry `json:"disable-entry,omitempty"`
	Entry        *[]FirewallInternetServiceExtensionEntry        `json:"entry,omitempty"`
	Id           *int64                                          `json:"id,omitempty"`
}

const FirewallInternetServiceExtensionDisableEntryPath = "firewall/internet-service-extension/disable-entry/"

type FirewallInternetServiceExtensionDisableEntry struct {
	Id        *int64                                                   `json:"id,omitempty"`
	IpRange   *[]FirewallInternetServiceExtensionDisableEntryIpRange   `json:"ip-range,omitempty"`
	PortRange *[]FirewallInternetServiceExtensionDisableEntryPortRange `json:"port-range,omitempty"`
	Protocol  *int64                                                   `json:"protocol,omitempty"`
}

const FirewallInternetServiceExtensionDisableEntryIpRangePath = "firewall/internet-service-extension/disable-entry/ip-range/"

type FirewallInternetServiceExtensionDisableEntryIpRange struct {
	EndIp   *string `json:"end-ip,omitempty"`
	Id      *int64  `json:"id,omitempty"`
	StartIp *string `json:"start-ip,omitempty"`
}

const FirewallInternetServiceExtensionDisableEntryPortRangePath = "firewall/internet-service-extension/disable-entry/port-range/"

type FirewallInternetServiceExtensionDisableEntryPortRange struct {
	EndPort   *int64 `json:"end-port,omitempty"`
	Id        *int64 `json:"id,omitempty"`
	StartPort *int64 `json:"start-port,omitempty"`
}

const FirewallInternetServiceExtensionEntryPath = "firewall/internet-service-extension/entry/"

type FirewallInternetServiceExtensionEntry struct {
	Dst       *[]FirewallInternetServiceExtensionEntryDst       `json:"dst,omitempty"`
	Id        *int64                                            `json:"id,omitempty"`
	PortRange *[]FirewallInternetServiceExtensionEntryPortRange `json:"port-range,omitempty"`
	Protocol  *int64                                            `json:"protocol,omitempty"`
}

const FirewallInternetServiceExtensionEntryDstPath = "firewall/internet-service-extension/entry/dst/"

type FirewallInternetServiceExtensionEntryDst struct {
	Name *string `json:"name,omitempty"`
}

const FirewallInternetServiceExtensionEntryPortRangePath = "firewall/internet-service-extension/entry/port-range/"

type FirewallInternetServiceExtensionEntryPortRange struct {
	EndPort   *int64 `json:"end-port,omitempty"`
	Id        *int64 `json:"id,omitempty"`
	StartPort *int64 `json:"start-port,omitempty"`
}
