package models

const SystemDhcpServerPath = "system.dhcp/server/"

type SystemDhcpServer struct {
	AutoConfiguration         *string                            `json:"auto-configuration,omitempty"`
	AutoManagedStatus         *string                            `json:"auto-managed-status,omitempty"`
	ConflictedIpTimeout       *int64                             `json:"conflicted-ip-timeout,omitempty"`
	DdnsAuth                  *string                            `json:"ddns-auth,omitempty"`
	DdnsKey                   *string                            `json:"ddns-key,omitempty"`
	DdnsKeyname               *string                            `json:"ddns-keyname,omitempty"`
	DdnsServerIp              *string                            `json:"ddns-server-ip,omitempty"`
	DdnsTtl                   *int64                             `json:"ddns-ttl,omitempty"`
	DdnsUpdate                *string                            `json:"ddns-update,omitempty"`
	DdnsUpdateOverride        *string                            `json:"ddns-update-override,omitempty"`
	DdnsZone                  *string                            `json:"ddns-zone,omitempty"`
	DefaultGateway            *string                            `json:"default-gateway,omitempty"`
	DhcpSettingsFromFortiipam *string                            `json:"dhcp-settings-from-fortiipam,omitempty"`
	DnsServer1                *string                            `json:"dns-server1,omitempty"`
	DnsServer2                *string                            `json:"dns-server2,omitempty"`
	DnsServer3                *string                            `json:"dns-server3,omitempty"`
	DnsServer4                *string                            `json:"dns-server4,omitempty"`
	DnsService                *string                            `json:"dns-service,omitempty"`
	Domain                    *string                            `json:"domain,omitempty"`
	ExcludeRange              *[]SystemDhcpServerExcludeRange    `json:"exclude-range,omitempty"`
	Filename                  *string                            `json:"filename,omitempty"`
	ForticlientOnNetStatus    *string                            `json:"forticlient-on-net-status,omitempty"`
	Id                        *int64                             `json:"id,omitempty"`
	Interface                 *string                            `json:"interface,omitempty"`
	IpMode                    *string                            `json:"ip-mode,omitempty"`
	IpRange                   *[]SystemDhcpServerIpRange         `json:"ip-range,omitempty"`
	IpsecLeaseHold            *int64                             `json:"ipsec-lease-hold,omitempty"`
	LeaseTime                 *int64                             `json:"lease-time,omitempty"`
	MacAclDefaultAction       *string                            `json:"mac-acl-default-action,omitempty"`
	Netmask                   *string                            `json:"netmask,omitempty"`
	NextServer                *string                            `json:"next-server,omitempty"`
	NtpServer1                *string                            `json:"ntp-server1,omitempty"`
	NtpServer2                *string                            `json:"ntp-server2,omitempty"`
	NtpServer3                *string                            `json:"ntp-server3,omitempty"`
	NtpService                *string                            `json:"ntp-service,omitempty"`
	Options                   *[]SystemDhcpServerOptions         `json:"options,omitempty"`
	ReservedAddress           *[]SystemDhcpServerReservedAddress `json:"reserved-address,omitempty"`
	ServerType                *string                            `json:"server-type,omitempty"`
	Status                    *string                            `json:"status,omitempty"`
	TftpServer                *[]SystemDhcpServerTftpServer      `json:"tftp-server,omitempty"`
	Timezone                  *string                            `json:"timezone,omitempty"`
	TimezoneOption            *string                            `json:"timezone-option,omitempty"`
	VciMatch                  *string                            `json:"vci-match,omitempty"`
	VciString                 *[]SystemDhcpServerVciString       `json:"vci-string,omitempty"`
	WifiAcService             *string                            `json:"wifi-ac-service,omitempty"`
	WifiAc1                   *string                            `json:"wifi-ac1,omitempty"`
	WifiAc2                   *string                            `json:"wifi-ac2,omitempty"`
	WifiAc3                   *string                            `json:"wifi-ac3,omitempty"`
	WinsServer1               *string                            `json:"wins-server1,omitempty"`
	WinsServer2               *string                            `json:"wins-server2,omitempty"`
}

type SystemDhcpServerExcludeRange struct {
	EndIp   *string `json:"end-ip,omitempty"`
	Id      *int64  `json:"id,omitempty"`
	StartIp *string `json:"start-ip,omitempty"`
}

type SystemDhcpServerIpRange struct {
	EndIp   *string `json:"end-ip,omitempty"`
	Id      *int64  `json:"id,omitempty"`
	StartIp *string `json:"start-ip,omitempty"`
}

type SystemDhcpServerOptions struct {
	Code  *int64  `json:"code,omitempty"`
	Id    *int64  `json:"id,omitempty"`
	Ip    *string `json:"ip,omitempty"`
	Type  *string `json:"type,omitempty"`
	Value *string `json:"value,omitempty"`
}

type SystemDhcpServerReservedAddress struct {
	Action        *string `json:"action,omitempty"`
	CircuitId     *string `json:"circuit-id,omitempty"`
	CircuitIdType *string `json:"circuit-id-type,omitempty"`
	Description   *string `json:"description,omitempty"`
	Id            *int64  `json:"id,omitempty"`
	Ip            *string `json:"ip,omitempty"`
	Mac           *string `json:"mac,omitempty"`
	RemoteId      *string `json:"remote-id,omitempty"`
	RemoteIdType  *string `json:"remote-id-type,omitempty"`
	Type          *string `json:"type,omitempty"`
}

type SystemDhcpServerTftpServer struct {
	TftpServer *string `json:"tftp-server,omitempty"`
}

type SystemDhcpServerVciString struct {
	VciString *string `json:"vci-string,omitempty"`
}
