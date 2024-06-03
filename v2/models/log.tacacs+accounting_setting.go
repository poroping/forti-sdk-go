package models

const LogTacacsaccountingSettingPath = "log.tacacs+accounting/setting/"

type LogTacacsaccountingSetting struct {
	Interface             *string `json:"interface,omitempty"`
	InterfaceSelectMethod *string `json:"interface-select-method,omitempty"`
	Server                *string `json:"server,omitempty"`
	ServerKey             *string `json:"server-key,omitempty"`
	SourceIp              *string `json:"source-ip,omitempty"`
	Status                *string `json:"status,omitempty"`
}
