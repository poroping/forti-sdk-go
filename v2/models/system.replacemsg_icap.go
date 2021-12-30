package models

const SystemreplacemsgIcapPath = "system/replacemsg/icap/"

type SystemreplacemsgIcap struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}
