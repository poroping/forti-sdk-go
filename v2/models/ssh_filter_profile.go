package models

const SshFilterProfilePath = "ssh-filter/profile/"

type SshFilterProfile struct {
	Block             *string                         `json:"block,omitempty"`
	DefaultCommandLog *string                         `json:"default-command-log,omitempty"`
	Log               *string                         `json:"log,omitempty"`
	Name              *string                         `json:"name,omitempty"`
	ShellCommands     []SshFilterProfileShellCommands `json:"shell-commands,omitempty"`
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
