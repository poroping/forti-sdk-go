package models

const FirewallLdbMonitorPath = "firewall/ldb-monitor/"

type FirewallLdbMonitor struct {
	DnsMatchIp       *string  `json:"dns-match-ip,omitempty"`
	DnsProtocol      *string  `json:"dns-protocol,omitempty"`
	DnsRequestDomain *string  `json:"dns-request-domain,omitempty"`
	HttpGet          *string  `json:"http-get,omitempty"`
	HttpMatch        *string  `json:"http-match,omitempty"`
	HttpMaxRedirects *float64 `json:"http-max-redirects,omitempty"`
	Interval         *float64 `json:"interval,omitempty"`
	Name             *string  `json:"name,omitempty"`
	Port             *float64 `json:"port,omitempty"`
	Retry            *float64 `json:"retry,omitempty"`
	SrcIp            *string  `json:"src-ip,omitempty"`
	Timeout          *float64 `json:"timeout,omitempty"`
	Type             *string  `json:"type,omitempty"`
}
