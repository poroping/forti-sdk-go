package models

const LogwebtrendsSettingPath = "log.webtrends/setting/"

type LogwebtrendsSetting struct {
	Server *string `json:"server,omitempty"`
	Status *string `json:"status,omitempty"`
}
