package models

const EmailfilterBwordPath = "emailfilter/bword/"

type EmailfilterBword struct {
	Comment *string                    `json:"comment,omitempty"`
	Entries *[]EmailfilterBwordEntries `json:"entries,omitempty"`
	Id      *int64                     `json:"id,omitempty"`
	Name    *string                    `json:"name,omitempty"`
}

const EmailfilterBwordEntriesPath = "emailfilter/bword/entries/"

type EmailfilterBwordEntries struct {
	Action      *string `json:"action,omitempty"`
	Id          *int64  `json:"id,omitempty"`
	Language    *string `json:"language,omitempty"`
	Pattern     *string `json:"pattern,omitempty"`
	PatternType *string `json:"pattern-type,omitempty"`
	Score       *int64  `json:"score,omitempty"`
	Status      *string `json:"status,omitempty"`
	Where       *string `json:"where,omitempty"`
}
