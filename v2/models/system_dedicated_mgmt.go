package models

const SystemDedicatedMgmtPath = "system/dedicated-mgmt/"

type SystemDedicatedMgmt struct {
	DefaultGateway *string `json:"default-gateway,omitempty"`
	DhcpEndIp      *string `json:"dhcp-end-ip,omitempty"`
	DhcpNetmask    *string `json:"dhcp-netmask,omitempty"`
	DhcpServer     *string `json:"dhcp-server,omitempty"`
	DhcpStartIp    *string `json:"dhcp-start-ip,omitempty"`
	Interface      *string `json:"interface,omitempty"`
	Status         *string `json:"status,omitempty"`
}
