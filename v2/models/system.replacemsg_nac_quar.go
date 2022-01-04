package models

const SystemReplacemsgNacQuarPath = "system.replacemsg/nac-quar/"

type SystemReplacemsgNacQuar struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}
