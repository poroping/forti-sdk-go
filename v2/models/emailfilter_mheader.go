package models

const EmailfilterMheaderPath = "emailfilter/mheader/"

type EmailfilterMheader struct {
	Comment *string                     `json:"comment,omitempty"`
	Entries []EmailfilterMheaderEntries `json:"entries,omitempty"`
	Fosid   *int64                      `json:"fosid,omitempty"`
	Name    *string                     `json:"name,omitempty"`
}

type EmailfilterMheaderEntries struct {
	Action      *string `json:"action,omitempty"`
	Fieldbody   *string `json:"fieldbody,omitempty"`
	Fieldname   *string `json:"fieldname,omitempty"`
	Id          *int64  `json:"id,omitempty"`
	PatternType *string `json:"pattern-type,omitempty"`
	Status      *string `json:"status,omitempty"`
}
