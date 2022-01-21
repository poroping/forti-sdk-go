package models

const LogMemorySettingPath = "log.memory/setting/"

type LogMemorySetting struct {
	Diskfull *string `json:"diskfull,omitempty"`
	Status   *string `json:"status,omitempty"`
}
