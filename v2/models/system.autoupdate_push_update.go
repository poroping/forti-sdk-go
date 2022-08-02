package models

const SystemAutoupdatePushUpdatePath = "system.autoupdate/push-update/"

type SystemAutoupdatePushUpdate struct {
	Address  *string `json:"address,omitempty"`
	Override *string `json:"override,omitempty"`
	Port     *int64  `json:"port,omitempty"`
	Status   *string `json:"status,omitempty"`
}
