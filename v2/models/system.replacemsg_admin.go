package models

const SystemreplacemsgAdminPath = "system/replacemsg/admin/"

type SystemreplacemsgAdmin struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}
