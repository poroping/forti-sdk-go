package models

const SystemReplacemsgUtmPath = "system.replacemsg/utm/"

type SystemReplacemsgUtm struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}
