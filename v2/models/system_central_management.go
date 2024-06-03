package models

const SystemCentralManagementPath = "system/central-management/"

type SystemCentralManagement struct {
	AllowMonitor                    *string                              `json:"allow-monitor,omitempty"`
	AllowPushConfiguration          *string                              `json:"allow-push-configuration,omitempty"`
	AllowPushFirmware               *string                              `json:"allow-push-firmware,omitempty"`
	AllowRemoteFirmwareUpgrade      *string                              `json:"allow-remote-firmware-upgrade,omitempty"`
	CaCert                          *string                              `json:"ca-cert,omitempty"`
	EncAlgorithm                    *string                              `json:"enc-algorithm,omitempty"`
	Fmg                             *string                              `json:"fmg,omitempty"`
	FmgSourceIp                     *string                              `json:"fmg-source-ip,omitempty"`
	FmgSourceIp6                    *string                              `json:"fmg-source-ip6,omitempty"`
	FmgUpdatePort                   *string                              `json:"fmg-update-port,omitempty"`
	FortigateCloudSsoDefaultProfile *string                              `json:"fortigate-cloud-sso-default-profile,omitempty"`
	IncludeDefaultServers           *string                              `json:"include-default-servers,omitempty"`
	Interface                       *string                              `json:"interface,omitempty"`
	InterfaceSelectMethod           *string                              `json:"interface-select-method,omitempty"`
	LocalCert                       *string                              `json:"local-cert,omitempty"`
	Mode                            *string                              `json:"mode,omitempty"`
	ScheduleConfigRestore           *string                              `json:"schedule-config-restore,omitempty"`
	ScheduleScriptRestore           *string                              `json:"schedule-script-restore,omitempty"`
	SerialNumber                    *string                              `json:"serial-number,omitempty"`
	ServerList                      *[]SystemCentralManagementServerList `json:"server-list,omitempty"`
	Type                            *string                              `json:"type,omitempty"`
	Vdom                            *string                              `json:"vdom,omitempty"`
}

type SystemCentralManagementServerList struct {
	AddrType       *string `json:"addr-type,omitempty"`
	Fqdn           *string `json:"fqdn,omitempty"`
	Id             *int64  `json:"id,omitempty"`
	ServerAddress  *string `json:"server-address,omitempty"`
	ServerAddress6 *string `json:"server-address6,omitempty"`
	ServerType     *string `json:"server-type,omitempty"`
}
