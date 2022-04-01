package models

const EmailfilterBlockAllowListPath = "emailfilter/block-allow-list/"

type EmailfilterBlockAllowList struct {
	Comment *string                             `json:"comment,omitempty"`
	Entries *[]EmailfilterBlockAllowListEntries `json:"entries,omitempty"`
	Id      *int64                              `json:"id,omitempty"`
	Name    *string                             `json:"name,omitempty"`
}

type EmailfilterBlockAllowListEntries struct {
	Action       *string `json:"action,omitempty"`
	AddrType     *string `json:"addr-type,omitempty"`
	EmailPattern *string `json:"email-pattern,omitempty"`
	Id           *int64  `json:"id,omitempty"`
	Ip4Subnet    *string `json:"ip4-subnet,omitempty"`
	Ip6Subnet    *string `json:"ip6-subnet,omitempty"`
	Pattern      *string `json:"pattern,omitempty"`
	PatternType  *string `json:"pattern-type,omitempty"`
	Status       *string `json:"status,omitempty"`
	Type         *string `json:"type,omitempty"`
}
