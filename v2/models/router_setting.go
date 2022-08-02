package models

const RouterSettingPath = "router/setting/"

type RouterSetting struct {
	Hostname   *string `json:"hostname,omitempty"`
	ShowFilter *string `json:"show-filter,omitempty"`
}
