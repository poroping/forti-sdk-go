package models

const EmailfilterBwlPath = "emailfilter/bwl/"

type EmailfilterBwl struct {
	Comment *string                  `json:"comment,omitempty"`
	Entries *[]EmailfilterBwlEntries `json:"entries,omitempty"`
	Id      *int64                   `json:"id,omitempty"`
	Name    *string                  `json:"name,omitempty"`
}

const EmailfilterBwlEntriesPath = "emailfilter/bwl/entries/"

type EmailfilterBwlEntries struct {
	Action       *string `json:"action,omitempty"`
	AddrType     *string `json:"addr-type,omitempty"`
	EmailPattern *string `json:"email-pattern,omitempty"`
	Id           *int64  `json:"id,omitempty"`
	Ip4Subnet    *string `json:"ip4-subnet,omitempty"`
	Ip6Subnet    *string `json:"ip6-subnet,omitempty"`
	PatternType  *string `json:"pattern-type,omitempty"`
	Status       *string `json:"status,omitempty"`
	Type         *string `json:"type,omitempty"`
}
