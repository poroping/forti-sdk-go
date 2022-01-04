package models

const SystemReplacemsgFortiguardWfPath = "system.replacemsg/fortiguard-wf/"

type SystemReplacemsgFortiguardWf struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}
