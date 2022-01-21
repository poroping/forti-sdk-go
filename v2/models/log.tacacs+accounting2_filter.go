package models

const LogTacacsaccounting2FilterPath = "log.tacacs+accounting2/filter/"

type LogTacacsaccounting2Filter struct {
	CliCmdAudit       *string `json:"cli-cmd-audit,omitempty"`
	ConfigChangeAudit *string `json:"config-change-audit,omitempty"`
	LoginAudit        *string `json:"login-audit,omitempty"`
}
