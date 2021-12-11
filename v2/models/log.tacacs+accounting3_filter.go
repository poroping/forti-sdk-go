package models

const Logtacacsaccounting3FilterPath = "log.tacacs+accounting3/filter/"

type Logtacacsaccounting3Filter struct {
	CliCmdAudit       *string `json:"cli-cmd-audit,omitempty"`
	ConfigChangeAudit *string `json:"config-change-audit,omitempty"`
	LoginAudit        *string `json:"login-audit,omitempty"`
}
