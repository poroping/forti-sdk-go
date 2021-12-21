package models

const SystemDdnsPath = "system/ddns/"

type SystemDdns struct {
	AddrType         *string                       `json:"addr-type,omitempty"`
	BoundIp          *string                       `json:"bound-ip,omitempty"`
	ClearText        *string                       `json:"clear-text,omitempty"`
	DdnsAuth         *string                       `json:"ddns-auth,omitempty"`
	DdnsDomain       *string                       `json:"ddns-domain,omitempty"`
	DdnsKey          *string                       `json:"ddns-key,omitempty"`
	DdnsKeyname      *string                       `json:"ddns-keyname,omitempty"`
	DdnsPassword     *string                       `json:"ddns-password,omitempty"`
	DdnsServer       *string                       `json:"ddns-server,omitempty"`
	DdnsServerAddr   *[]SystemDdnsDdnsServerAddr   `json:"ddns-server-addr,omitempty"`
	DdnsServerIp     *string                       `json:"ddns-server-ip,omitempty"`
	DdnsSn           *string                       `json:"ddns-sn,omitempty"`
	DdnsTtl          *float64                      `json:"ddns-ttl,omitempty"`
	DdnsUsername     *string                       `json:"ddns-username,omitempty"`
	DdnsZone         *string                       `json:"ddns-zone,omitempty"`
	Ddnsid           *float64                      `json:"ddnsid,omitempty"`
	MonitorInterface *[]SystemDdnsMonitorInterface `json:"monitor-interface,omitempty"`
	ServerType       *string                       `json:"server-type,omitempty"`
	SslCertificate   *string                       `json:"ssl-certificate,omitempty"`
	UpdateInterval   *float64                      `json:"update-interval,omitempty"`
	UsePublicIp      *string                       `json:"use-public-ip,omitempty"`
}

type SystemDdnsDdnsServerAddr struct {
	Addr *string `json:"addr,omitempty"`
}

type SystemDdnsMonitorInterface struct {
	InterfaceName *string `json:"interface-name,omitempty"`
}
