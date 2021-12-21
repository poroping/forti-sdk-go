package models

const DlpFilepatternPath = "dlp/filepattern/"

type DlpFilepattern struct {
	Comment *string                  `json:"comment,omitempty"`
	Entries *[]DlpFilepatternEntries `json:"entries,omitempty"`
	Fosid   *float64                 `json:"fosid,omitempty"`
	Name    *string                  `json:"name,omitempty"`
}

type DlpFilepatternEntries struct {
	FileType   *string `json:"file-type,omitempty"`
	FilterType *string `json:"filter-type,omitempty"`
	Pattern    *string `json:"pattern,omitempty"`
}
