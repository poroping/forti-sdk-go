package models

const SystemDnsServerPath = "system/dns-server/"

type SystemDnsServer struct {
	DnsfilterProfile *string `json:"dnsfilter-profile,omitempty"`
	Doh              *string `json:"doh,omitempty"`
	Mode             *string `json:"mode,omitempty"`
	Name             *string `json:"name,omitempty"`
}
