package models

const SystemFortimanagerPath = "system/fortimanager/"

type SystemFortimanager struct {
	CentralManagement                *string `json:"central-management,omitempty"`
	CentralMgmtAutoBackup            *string `json:"central-mgmt-auto-backup,omitempty"`
	CentralMgmtScheduleConfigRestore *string `json:"central-mgmt-schedule-config-restore,omitempty"`
	CentralMgmtScheduleScriptRestore *string `json:"central-mgmt-schedule-script-restore,omitempty"`
	Ip                               *string `json:"ip,omitempty"`
	Ipsec                            *string `json:"ipsec,omitempty"`
	Vdom                             *string `json:"vdom,omitempty"`
}
