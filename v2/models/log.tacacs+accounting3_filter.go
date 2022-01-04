package models

const LogTacacsaccounting3FilterPath = "log.tacacs+accounting3/filter/"

type LogTacacsaccounting3Filter struct {
	CliCmdAudit       *string `json:"cli-cmd-audit,omitempty"`
	ConfigChangeAudit *string `json:"config-change-audit,omitempty"`
	LoginAudit        *string `json:"login-audit,omitempty"`
}
