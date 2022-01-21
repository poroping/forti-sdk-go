package models

const FileFilterProfilePath = "file-filter/profile/"

type FileFilterProfile struct {
	Comment             *string                   `json:"comment,omitempty"`
	ExtendedLog         *string                   `json:"extended-log,omitempty"`
	FeatureSet          *string                   `json:"feature-set,omitempty"`
	Log                 *string                   `json:"log,omitempty"`
	Name                *string                   `json:"name,omitempty"`
	ReplacemsgGroup     *string                   `json:"replacemsg-group,omitempty"`
	Rules               *[]FileFilterProfileRules `json:"rules,omitempty"`
	ScanArchiveContents *string                   `json:"scan-archive-contents,omitempty"`
}

type FileFilterProfileRules struct {
	Action            *string                           `json:"action,omitempty"`
	Comment           *string                           `json:"comment,omitempty"`
	Direction         *string                           `json:"direction,omitempty"`
	FileType          *[]FileFilterProfileRulesFileType `json:"file-type,omitempty"`
	Name              *string                           `json:"name,omitempty"`
	PasswordProtected *string                           `json:"password-protected,omitempty"`
	Protocol          *string                           `json:"protocol,omitempty"`
}

type FileFilterProfileRulesFileType struct {
	Name *string `json:"name,omitempty"`
}
