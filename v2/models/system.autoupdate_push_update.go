package models

const SystemautoupdatePushUpdatePath = "system.autoupdate/push-update/"

type SystemautoupdatePushUpdate struct {
	Address  *string  `json:"address,omitempty"`
	Override *string  `json:"override,omitempty"`
	Port     *float64 `json:"port,omitempty"`
	Status   *string  `json:"status,omitempty"`
}
