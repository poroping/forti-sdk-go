package models

const LogTacacsaccountingFilterPath = "log.tacacs+accounting/filter/"

type LogTacacsaccountingFilter struct {
	CliCmdAudit       *string `json:"cli-cmd-audit,omitempty"`
	ConfigChangeAudit *string `json:"config-change-audit,omitempty"`
	LoginAudit        *string `json:"login-audit,omitempty"`
}
