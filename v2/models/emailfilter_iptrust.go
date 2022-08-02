package models

const EmailfilterIptrustPath = "emailfilter/iptrust/"

type EmailfilterIptrust struct {
	Comment *string                      `json:"comment,omitempty"`
	Entries *[]EmailfilterIptrustEntries `json:"entries,omitempty"`
	Id      *int64                       `json:"id,omitempty"`
	Name    *string                      `json:"name,omitempty"`
}

const EmailfilterIptrustEntriesPath = "emailfilter/iptrust/entries/"

type EmailfilterIptrustEntries struct {
	AddrType  *string `json:"addr-type,omitempty"`
	Id        *int64  `json:"id,omitempty"`
	Ip4Subnet *string `json:"ip4-subnet,omitempty"`
	Ip6Subnet *string `json:"ip6-subnet,omitempty"`
	Status    *string `json:"status,omitempty"`
}
