package models

const SystemReplacemsgSpamPath = "system.replacemsg/spam/"

type SystemReplacemsgSpam struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}
