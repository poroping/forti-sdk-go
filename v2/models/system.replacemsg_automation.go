package models

const SystemreplacemsgAutomationPath = "system/replacemsg/automation/"

type SystemreplacemsgAutomation struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}
