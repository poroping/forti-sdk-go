package models

const LognullDeviceSettingPath = "log.null-device/setting/"

type LognullDeviceSetting struct {
	Status *string `json:"status,omitempty"`
}
