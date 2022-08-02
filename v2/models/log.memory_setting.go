package models

const LogMemorySettingPath = "log.memory/setting/"

type LogMemorySetting struct {
	Status *string `json:"status,omitempty"`
}
