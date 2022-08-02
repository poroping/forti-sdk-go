package models

const DlpFilepatternPath = "dlp/filepattern/"

type DlpFilepattern struct {
	Comment *string                  `json:"comment,omitempty"`
	Entries *[]DlpFilepatternEntries `json:"entries,omitempty"`
	Id      *int64                   `json:"id,omitempty"`
	Name    *string                  `json:"name,omitempty"`
}

const DlpFilepatternEntriesPath = "dlp/filepattern/entries/"

type DlpFilepatternEntries struct {
	FileType   *string `json:"file-type,omitempty"`
	FilterType *string `json:"filter-type,omitempty"`
	Pattern    *string `json:"pattern,omitempty"`
}
