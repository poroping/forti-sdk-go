package models

const DlpFpDocSourcePath = "dlp/fp-doc-source/"

type DlpFpDocSource struct {
	Date               *float64 `json:"date,omitempty"`
	FilePath           *string  `json:"file-path,omitempty"`
	FilePattern        *string  `json:"file-pattern,omitempty"`
	KeepModified       *string  `json:"keep-modified,omitempty"`
	Name               *string  `json:"name,omitempty"`
	Password           *string  `json:"password,omitempty"`
	Period             *string  `json:"period,omitempty"`
	RemoveDeleted      *string  `json:"remove-deleted,omitempty"`
	ScanOnCreation     *string  `json:"scan-on-creation,omitempty"`
	ScanSubdirectories *string  `json:"scan-subdirectories,omitempty"`
	Sensitivity        *string  `json:"sensitivity,omitempty"`
	Server             *string  `json:"server,omitempty"`
	ServerType         *string  `json:"server-type,omitempty"`
	TodHour            *float64 `json:"tod-hour,omitempty"`
	TodMin             *float64 `json:"tod-min,omitempty"`
	Username           *string  `json:"username,omitempty"`
	Vdom               *string  `json:"vdom,omitempty"`
	Weekday            *string  `json:"weekday,omitempty"`
}
