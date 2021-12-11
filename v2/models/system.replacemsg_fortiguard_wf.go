package models

const SystemreplacemsgFortiguardWfPath = "system.replacemsg/fortiguard-wf/"

type SystemreplacemsgFortiguardWf struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}
