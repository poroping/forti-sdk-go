package models

const LogmemoryGlobalSettingPath = "log.memory/global-setting/"

type LogmemoryGlobalSetting struct {
	FullFinalWarningThreshold  *int64 `json:"full-final-warning-threshold,omitempty"`
	FullFirstWarningThreshold  *int64 `json:"full-first-warning-threshold,omitempty"`
	FullSecondWarningThreshold *int64 `json:"full-second-warning-threshold,omitempty"`
	MaxSize                    *int64 `json:"max-size,omitempty"`
}
