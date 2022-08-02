package models

const LogTacacsaccounting3SettingPath = "log.tacacs+accounting3/setting/"

type LogTacacsaccounting3Setting struct {
	Server    *string `json:"server,omitempty"`
	ServerKey *string `json:"server-key,omitempty"`
	Status    *string `json:"status,omitempty"`
}
