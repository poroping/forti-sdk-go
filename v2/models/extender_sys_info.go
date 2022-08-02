package models

const ExtenderSysInfoPath = "extender/sys-info/"

type ExtenderSysInfo struct {
	Sn *string `json:"<sn>,omitempty"`
}
