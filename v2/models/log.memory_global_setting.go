package models

const LogmemoryGlobalSettingPath = "log.memory/global-setting/"

type LogmemoryGlobalSetting struct {
	FullFinalWarningThreshold  *float64 `json:"full-final-warning-threshold,omitempty"`
	FullFirstWarningThreshold  *float64 `json:"full-first-warning-threshold,omitempty"`
	FullSecondWarningThreshold *float64 `json:"full-second-warning-threshold,omitempty"`
	MaxSize                    *float64 `json:"max-size,omitempty"`
}
