package models

const WebfilterContentPath = "webfilter/content/"

type WebfilterContent struct {
	Comment *string                    `json:"comment,omitempty"`
	Entries *[]WebfilterContentEntries `json:"entries,omitempty"`
	Fosid   *int64                     `json:"fosid,omitempty"`
	Name    *string                    `json:"name,omitempty"`
}

type WebfilterContentEntries struct {
	Action      *string `json:"action,omitempty"`
	Lang        *string `json:"lang,omitempty"`
	Name        *string `json:"name,omitempty"`
	PatternType *string `json:"pattern-type,omitempty"`
	Score       *int64  `json:"score,omitempty"`
	Status      *string `json:"status,omitempty"`
}
