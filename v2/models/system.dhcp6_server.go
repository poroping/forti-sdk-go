package models

const Systemdhcp6ServerPath = "system/dhcp6/server/"

type Systemdhcp6Server struct {
	DelegatedPrefixIaid *int64                          `json:"delegated-prefix-iaid,omitempty"`
	DnsSearchList       *string                         `json:"dns-search-list,omitempty"`
	DnsServer1          *string                         `json:"dns-server1,omitempty"`
	DnsServer2          *string                         `json:"dns-server2,omitempty"`
	DnsServer3          *string                         `json:"dns-server3,omitempty"`
	DnsServer4          *string                         `json:"dns-server4,omitempty"`
	DnsService          *string                         `json:"dns-service,omitempty"`
	Domain              *string                         `json:"domain,omitempty"`
	Id                  *int64                          `json:"id,omitempty"`
	Interface           *string                         `json:"interface,omitempty"`
	IpMode              *string                         `json:"ip-mode,omitempty"`
	IpRange             *[]Systemdhcp6ServerIpRange     `json:"ip-range,omitempty"`
	LeaseTime           *int64                          `json:"lease-time,omitempty"`
	Option1             *string                         `json:"option1,omitempty"`
	Option2             *string                         `json:"option2,omitempty"`
	Option3             *string                         `json:"option3,omitempty"`
	PrefixMode          *string                         `json:"prefix-mode,omitempty"`
	PrefixRange         *[]Systemdhcp6ServerPrefixRange `json:"prefix-range,omitempty"`
	RapidCommit         *string                         `json:"rapid-commit,omitempty"`
	Status              *string                         `json:"status,omitempty"`
	Subnet              *string                         `json:"subnet,omitempty"`
	UpstreamInterface   *string                         `json:"upstream-interface,omitempty"`
}

type Systemdhcp6ServerIpRange struct {
	EndIp   *string `json:"end-ip,omitempty"`
	Id      *int64  `json:"id,omitempty"`
	StartIp *string `json:"start-ip,omitempty"`
}

type Systemdhcp6ServerPrefixRange struct {
	EndPrefix    *string `json:"end-prefix,omitempty"`
	Id           *int64  `json:"id,omitempty"`
	PrefixLength *int64  `json:"prefix-length,omitempty"`
	StartPrefix  *string `json:"start-prefix,omitempty"`
}
