package models

const SystemIpsUrlfilterDnsPath = "system/ips-urlfilter-dns/"

type SystemIpsUrlfilterDns struct {
	Address        *string `json:"address,omitempty"`
	Ipv6Capability *string `json:"ipv6-capability,omitempty"`
	Status         *string `json:"status,omitempty"`
}
