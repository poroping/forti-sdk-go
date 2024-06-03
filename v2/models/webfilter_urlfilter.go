package models

const WebfilterUrlfilterPath = "webfilter/urlfilter/"

type WebfilterUrlfilter struct {
	Comment            *string                      `json:"comment,omitempty"`
	Entries            *[]WebfilterUrlfilterEntries `json:"entries,omitempty"`
	Id                 *int64                       `json:"id,omitempty"`
	IpAddrBlock        *string                      `json:"ip-addr-block,omitempty"`
	Ip4MappedIp6       *string                      `json:"ip4-mapped-ip6,omitempty"`
	Name               *string                      `json:"name,omitempty"`
	OneArmIpsUrlfilter *string                      `json:"one-arm-ips-urlfilter,omitempty"`
}

type WebfilterUrlfilterEntries struct {
	Action           *string `json:"action,omitempty"`
	AntiphishAction  *string `json:"antiphish-action,omitempty"`
	DnsAddressFamily *string `json:"dns-address-family,omitempty"`
	Exempt           *string `json:"exempt,omitempty"`
	Id               *int64  `json:"id,omitempty"`
	ReferrerHost     *string `json:"referrer-host,omitempty"`
	Status           *string `json:"status,omitempty"`
	Type             *string `json:"type,omitempty"`
	Url              *string `json:"url,omitempty"`
	WebProxyProfile  *string `json:"web-proxy-profile,omitempty"`
}
