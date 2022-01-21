package models

const FirewallLdbMonitorPath = "firewall/ldb-monitor/"

type FirewallLdbMonitor struct {
	DnsMatchIp       *string `json:"dns-match-ip,omitempty"`
	DnsProtocol      *string `json:"dns-protocol,omitempty"`
	DnsRequestDomain *string `json:"dns-request-domain,omitempty"`
	HttpGet          *string `json:"http-get,omitempty"`
	HttpMatch        *string `json:"http-match,omitempty"`
	HttpMaxRedirects *int64  `json:"http-max-redirects,omitempty"`
	Interval         *int64  `json:"interval,omitempty"`
	Name             *string `json:"name,omitempty"`
	Port             *int64  `json:"port,omitempty"`
	Retry            *int64  `json:"retry,omitempty"`
	SrcIp            *string `json:"src-ip,omitempty"`
	Timeout          *int64  `json:"timeout,omitempty"`
	Type             *string `json:"type,omitempty"`
}
