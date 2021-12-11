package models

const LogmemorySettingPath = "log.memory/setting/"

type LogmemorySetting struct {
	Status *string `json:"status,omitempty"`
}
