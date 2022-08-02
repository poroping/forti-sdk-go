package models

const SwitchControllerGlobalPath = "switch-controller/global/"

type SwitchControllerGlobal struct {
	BounceQuarantinedLink            *string                                   `json:"bounce-quarantined-link,omitempty"`
	CustomCommand                    *[]SwitchControllerGlobalCustomCommand    `json:"custom-command,omitempty"`
	DefaultVirtualSwitchVlan         *string                                   `json:"default-virtual-switch-vlan,omitempty"`
	DhcpServerAccessList             *string                                   `json:"dhcp-server-access-list,omitempty"`
	DisableDiscovery                 *[]SwitchControllerGlobalDisableDiscovery `json:"disable-discovery,omitempty"`
	FipsEnforce                      *string                                   `json:"fips-enforce,omitempty"`
	FirmwareProvisionOnAuthorization *string                                   `json:"firmware-provision-on-authorization,omitempty"`
	HttpsImagePush                   *string                                   `json:"https-image-push,omitempty"`
	LogMacLimitViolations            *string                                   `json:"log-mac-limit-violations,omitempty"`
	MacAgingInterval                 *int64                                    `json:"mac-aging-interval,omitempty"`
	MacEventLogging                  *string                                   `json:"mac-event-logging,omitempty"`
	MacRetentionPeriod               *int64                                    `json:"mac-retention-period,omitempty"`
	MacViolationTimer                *int64                                    `json:"mac-violation-timer,omitempty"`
	QuarantineMode                   *string                                   `json:"quarantine-mode,omitempty"`
	SnDnsResolution                  *string                                   `json:"sn-dns-resolution,omitempty"`
	UpdateUserDevice                 *string                                   `json:"update-user-device,omitempty"`
	VlanAllMode                      *string                                   `json:"vlan-all-mode,omitempty"`
	VlanOptimization                 *string                                   `json:"vlan-optimization,omitempty"`
}

const SwitchControllerGlobalCustomCommandPath = "switch-controller/global/custom-command/"

type SwitchControllerGlobalCustomCommand struct {
	CommandEntry *string `json:"command-entry,omitempty"`
	CommandName  *string `json:"command-name,omitempty"`
}

const SwitchControllerGlobalDisableDiscoveryPath = "switch-controller/global/disable-discovery/"

type SwitchControllerGlobalDisableDiscovery struct {
	Name *string `json:"name,omitempty"`
}
