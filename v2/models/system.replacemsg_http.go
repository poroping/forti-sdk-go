package models

const SystemreplacemsgHttpPath = "system.replacemsg/http/"

type SystemreplacemsgHttp struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}
