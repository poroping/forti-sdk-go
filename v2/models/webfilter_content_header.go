package models

const WebfilterContentHeaderPath = "webfilter/content-header/"

type WebfilterContentHeader struct {
	Comment *string                          `json:"comment,omitempty"`
	Entries *[]WebfilterContentHeaderEntries `json:"entries,omitempty"`
	Fosid   *float64                         `json:"fosid,omitempty"`
	Name    *string                          `json:"name,omitempty"`
}

type WebfilterContentHeaderEntries struct {
	Action   *string `json:"action,omitempty"`
	Category *string `json:"category,omitempty"`
	Pattern  *string `json:"pattern,omitempty"`
}
