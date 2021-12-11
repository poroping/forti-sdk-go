package models

const Logtacacsaccounting2SettingPath = "log.tacacs+accounting2/setting/"

type Logtacacsaccounting2Setting struct {
	Server    *string `json:"server,omitempty"`
	ServerKey *string `json:"server-key,omitempty"`
	Status    *string `json:"status,omitempty"`
}
