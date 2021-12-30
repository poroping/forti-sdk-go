package models

const SystemreplacemsgAuthPath = "system/replacemsg/auth/"

type SystemreplacemsgAuth struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}
