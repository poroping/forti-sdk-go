package models

const SystemDnsPath = "system/dns/"

type SystemDns struct {
	AltPrimary             *string                    `json:"alt-primary,omitempty"`
	AltSecondary           *string                    `json:"alt-secondary,omitempty"`
	CacheNotfoundResponses *string                    `json:"cache-notfound-responses,omitempty"`
	DnsCacheLimit          *int64                     `json:"dns-cache-limit,omitempty"`
	DnsCacheTtl            *int64                     `json:"dns-cache-ttl,omitempty"`
	DnsOverTls             *string                    `json:"dns-over-tls,omitempty"`
	Domain                 *[]SystemDnsDomain         `json:"domain,omitempty"`
	Interface              *string                    `json:"interface,omitempty"`
	InterfaceSelectMethod  *string                    `json:"interface-select-method,omitempty"`
	Ip6Primary             *string                    `json:"ip6-primary,omitempty"`
	Ip6Secondary           *string                    `json:"ip6-secondary,omitempty"`
	Log                    *string                    `json:"log,omitempty"`
	Primary                *string                    `json:"primary,omitempty"`
	Protocol               *string                    `json:"protocol,omitempty"`
	Retry                  *int64                     `json:"retry,omitempty"`
	Secondary              *string                    `json:"secondary,omitempty"`
	ServerHostname         *[]SystemDnsServerHostname `json:"server-hostname,omitempty"`
	ServerSelectMethod     *string                    `json:"server-select-method,omitempty"`
	SourceIp               *string                    `json:"source-ip,omitempty"`
	SslCertificate         *string                    `json:"ssl-certificate,omitempty"`
	Timeout                *int64                     `json:"timeout,omitempty"`
}

type SystemDnsDomain struct {
	Domain *string `json:"domain,omitempty"`
}

type SystemDnsServerHostname struct {
	Hostname *string `json:"hostname,omitempty"`
}
