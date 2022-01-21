package models

const LogFortiguardOverrideSettingPath = "log.fortiguard/override-setting/"

type LogFortiguardOverrideSetting struct {
	AccessConfig   *string `json:"access-config,omitempty"`
	MaxLogRate     *int64  `json:"max-log-rate,omitempty"`
	Override       *string `json:"override,omitempty"`
	Priority       *string `json:"priority,omitempty"`
	Status         *string `json:"status,omitempty"`
	UploadDay      *string `json:"upload-day,omitempty"`
	UploadInterval *string `json:"upload-interval,omitempty"`
	UploadOption   *string `json:"upload-option,omitempty"`
	UploadTime     *string `json:"upload-time,omitempty"`
}
