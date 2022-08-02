package models

const LogWebtrendsSettingPath = "log.webtrends/setting/"

type LogWebtrendsSetting struct {
	Server *string `json:"server,omitempty"`
	Status *string `json:"status,omitempty"`
}
