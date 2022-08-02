package models

const SystemFmPath = "system/fm/"

type SystemFm struct {
	AutoBackup             *string `json:"auto-backup,omitempty"`
	Id                     *string `json:"id,omitempty"`
	Ip                     *string `json:"ip,omitempty"`
	Ipsec                  *string `json:"ipsec,omitempty"`
	ScheduledConfigRestore *string `json:"scheduled-config-restore,omitempty"`
	Status                 *string `json:"status,omitempty"`
	Vdom                   *string `json:"vdom,omitempty"`
}
