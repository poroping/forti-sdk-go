package models

const LogTacacsaccounting2SettingPath = "log.tacacs+accounting2/setting/"

type LogTacacsaccounting2Setting struct {
	Server    *string `json:"server,omitempty"`
	ServerKey *string `json:"server-key,omitempty"`
	Status    *string `json:"status,omitempty"`
}
