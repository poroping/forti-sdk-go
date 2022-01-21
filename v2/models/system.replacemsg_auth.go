package models

const SystemReplacemsgAuthPath = "system.replacemsg/auth/"

type SystemReplacemsgAuth struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}
