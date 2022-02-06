package models

const SshFilterProfilePath = "ssh-filter/profile/"

type SshFilterProfile struct {
	Block             *string                          `json:"block,omitempty"`
	DefaultCommandLog *string                          `json:"default-command-log,omitempty"`
	FileFilter        *SshFilterProfileFileFilter      `json:"file-filter,omitempty"`
	Log               *string                          `json:"log,omitempty"`
	Name              *string                          `json:"name,omitempty"`
	ShellCommands     *[]SshFilterProfileShellCommands `json:"shell-commands,omitempty"`
}

type SshFilterProfileFileFilter struct {
	Entries             *[]SshFilterProfileFileFilterEntries `json:"entries,omitempty"`
	Log                 *string                              `json:"log,omitempty"`
	ScanArchiveContents *string                              `json:"scan-archive-contents,omitempty"`
	Status              *string                              `json:"status,omitempty"`
}

type SshFilterProfileFileFilterEntries struct {
	Action            *string                                      `json:"action,omitempty"`
	Comment           *string                                      `json:"comment,omitempty"`
	Direction         *string                                      `json:"direction,omitempty"`
	FileType          *[]SshFilterProfileFileFilterEntriesFileType `json:"file-type,omitempty"`
	Filter            *string                                      `json:"filter,omitempty"`
	PasswordProtected *string                                      `json:"password-protected,omitempty"`
}

type SshFilterProfileFileFilterEntriesFileType struct {
	Name *string `json:"name,omitempty"`
}

type SshFilterProfileShellCommands struct {
	Action   *string `json:"action,omitempty"`
	Alert    *string `json:"alert,omitempty"`
	Id       *int64  `json:"id,omitempty"`
	Log      *string `json:"log,omitempty"`
	Pattern  *string `json:"pattern,omitempty"`
	Severity *string `json:"severity,omitempty"`
	Type     *string `json:"type,omitempty"`
}
