package models

const LogNullDeviceSettingPath = "log.null-device/setting/"

type LogNullDeviceSetting struct {
	Status *string `json:"status,omitempty"`
}
