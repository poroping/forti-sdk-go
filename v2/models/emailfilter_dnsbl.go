package models

const EmailfilterDnsblPath = "emailfilter/dnsbl/"

type EmailfilterDnsbl struct {
	Comment *string                    `json:"comment,omitempty"`
	Entries *[]EmailfilterDnsblEntries `json:"entries,omitempty"`
	Fosid   *float64                   `json:"fosid,omitempty"`
	Name    *string                    `json:"name,omitempty"`
}

type EmailfilterDnsblEntries struct {
	Action *string  `json:"action,omitempty"`
	Id     *float64 `json:"id,omitempty"`
	Server *string  `json:"server,omitempty"`
	Status *string  `json:"status,omitempty"`
}
