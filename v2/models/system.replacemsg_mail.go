package models

const SystemReplacemsgMailPath = "system.replacemsg/mail/"

type SystemReplacemsgMail struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}
