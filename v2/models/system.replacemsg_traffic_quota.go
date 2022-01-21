package models

const SystemReplacemsgTrafficQuotaPath = "system.replacemsg/traffic-quota/"

type SystemReplacemsgTrafficQuota struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}
