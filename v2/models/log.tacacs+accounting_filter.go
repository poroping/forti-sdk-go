package models

const LogtacacsaccountingFilterPath = "log/tacacs+accounting/filter/"

type LogtacacsaccountingFilter struct {
	CliCmdAudit       *string `json:"cli-cmd-audit,omitempty"`
	ConfigChangeAudit *string `json:"config-change-audit,omitempty"`
	LoginAudit        *string `json:"login-audit,omitempty"`
}
