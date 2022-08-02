package models

const WebfilterContentHeaderPath = "webfilter/content-header/"

type WebfilterContentHeader struct {
	Comment *string                          `json:"comment,omitempty"`
	Entries *[]WebfilterContentHeaderEntries `json:"entries,omitempty"`
	Id      *int64                           `json:"id,omitempty"`
	Name    *string                          `json:"name,omitempty"`
}

const WebfilterContentHeaderEntriesPath = "webfilter/content-header/entries/"

type WebfilterContentHeaderEntries struct {
	Action   *string `json:"action,omitempty"`
	Category *string `json:"category,omitempty"`
	Pattern  *string `json:"pattern,omitempty"`
}
