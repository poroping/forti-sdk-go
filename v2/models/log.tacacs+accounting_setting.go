package models

const LogtacacsaccountingSettingPath = "log/tacacs+accounting/setting/"

type LogtacacsaccountingSetting struct {
	Server    *string `json:"server,omitempty"`
	ServerKey *string `json:"server-key,omitempty"`
	Status    *string `json:"status,omitempty"`
}
