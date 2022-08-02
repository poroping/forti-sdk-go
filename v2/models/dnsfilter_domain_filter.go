package models

const DnsfilterDomainFilterPath = "dnsfilter/domain-filter/"

type DnsfilterDomainFilter struct {
	Comment *string                         `json:"comment,omitempty"`
	Entries *[]DnsfilterDomainFilterEntries `json:"entries,omitempty"`
	Id      *int64                          `json:"id,omitempty"`
	Name    *string                         `json:"name,omitempty"`
}

const DnsfilterDomainFilterEntriesPath = "dnsfilter/domain-filter/entries/"

type DnsfilterDomainFilterEntries struct {
	Action *string `json:"action,omitempty"`
	Domain *string `json:"domain,omitempty"`
	Id     *int64  `json:"id,omitempty"`
	Status *string `json:"status,omitempty"`
	Type   *string `json:"type,omitempty"`
}
