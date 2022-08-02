package models

const WebfilterContentPath = "webfilter/content/"

type WebfilterContent struct {
	Comment *string                    `json:"comment,omitempty"`
	Entries *[]WebfilterContentEntries `json:"entries,omitempty"`
	Id      *int64                     `json:"id,omitempty"`
	Name    *string                    `json:"name,omitempty"`
}

const WebfilterContentEntriesPath = "webfilter/content/entries/"

type WebfilterContentEntries struct {
	Action      *string `json:"action,omitempty"`
	Lang        *string `json:"lang,omitempty"`
	Name        *string `json:"name,omitempty"`
	PatternType *string `json:"pattern-type,omitempty"`
	Score       *int64  `json:"score,omitempty"`
	Status      *string `json:"status,omitempty"`
}
