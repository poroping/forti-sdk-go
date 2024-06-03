package models

const LogTacacsaccounting2SettingPath = "log.tacacs+accounting2/setting/"

type LogTacacsaccounting2Setting struct {
	Interface             *string `json:"interface,omitempty"`
	InterfaceSelectMethod *string `json:"interface-select-method,omitempty"`
	Server                *string `json:"server,omitempty"`
	ServerKey             *string `json:"server-key,omitempty"`
	SourceIp              *string `json:"source-ip,omitempty"`
	Status                *string `json:"status,omitempty"`
}
