package models

const DnsfilterDomainFilterPath = "dnsfilter/domain-filter/"

type DnsfilterDomainFilter struct {
	Comment *string                         `json:"comment,omitempty"`
	Entries *[]DnsfilterDomainFilterEntries `json:"entries,omitempty"`
	Fosid   *int64                          `json:"fosid,omitempty"`
	Name    *string                         `json:"name,omitempty"`
}

type DnsfilterDomainFilterEntries struct {
	Action *string `json:"action,omitempty"`
	Domain *string `json:"domain,omitempty"`
	Id     *int64  `json:"id,omitempty"`
	Status *string `json:"status,omitempty"`
	Type   *string `json:"type,omitempty"`
}
