package models

const SystemVdomDnsPath = "system/vdom-dns/"

type SystemVdomDns struct {
	AltPrimary            *string                        `json:"alt-primary,omitempty"`
	AltSecondary          *string                        `json:"alt-secondary,omitempty"`
	DnsOverTls            *string                        `json:"dns-over-tls,omitempty"`
	Interface             *string                        `json:"interface,omitempty"`
	InterfaceSelectMethod *string                        `json:"interface-select-method,omitempty"`
	Ip6Primary            *string                        `json:"ip6-primary,omitempty"`
	Ip6Secondary          *string                        `json:"ip6-secondary,omitempty"`
	Primary               *string                        `json:"primary,omitempty"`
	Protocol              *string                        `json:"protocol,omitempty"`
	Secondary             *string                        `json:"secondary,omitempty"`
	ServerHostname        *[]SystemVdomDnsServerHostname `json:"server-hostname,omitempty"`
	ServerSelectMethod    *string                        `json:"server-select-method,omitempty"`
	SourceIp              *string                        `json:"source-ip,omitempty"`
	SslCertificate        *string                        `json:"ssl-certificate,omitempty"`
	VdomDns               *string                        `json:"vdom-dns,omitempty"`
}

const SystemVdomDnsServerHostnamePath = "system/vdom-dns/server-hostname/"

type SystemVdomDnsServerHostname struct {
	Hostname *string `json:"hostname,omitempty"`
}
