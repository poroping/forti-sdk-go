package models

const EmailfilterDnsblPath = "emailfilter/dnsbl/"

type EmailfilterDnsbl struct {
	Comment *string                    `json:"comment,omitempty"`
	Entries *[]EmailfilterDnsblEntries `json:"entries,omitempty"`
	Id      *int64                     `json:"id,omitempty"`
	Name    *string                    `json:"name,omitempty"`
}

type EmailfilterDnsblEntries struct {
	Action *string `json:"action,omitempty"`
	Id     *int64  `json:"id,omitempty"`
	Server *string `json:"server,omitempty"`
	Status *string `json:"status,omitempty"`
}
