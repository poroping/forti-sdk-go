package models

const LogTacacsaccountingSettingPath = "log.tacacs+accounting/setting/"

type LogTacacsaccountingSetting struct {
	Server    *string `json:"server,omitempty"`
	ServerKey *string `json:"server-key,omitempty"`
	Status    *string `json:"status,omitempty"`
}
