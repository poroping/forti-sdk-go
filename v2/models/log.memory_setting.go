package models

const LogmemorySettingPath = "log/memory/setting/"

type LogmemorySetting struct {
	Diskfull *string `json:"diskfull,omitempty"`
	Status   *string `json:"status,omitempty"`
}
