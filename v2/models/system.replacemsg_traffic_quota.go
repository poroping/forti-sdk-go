package models

const SystemReplacemsgTrafficQuotaPath = "system.replacemsg/traffic-quota/"

type SystemReplacemsgTrafficQuota struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}

// Set SystemReplacemsgTrafficQuota values to defaults
func (def *SystemReplacemsgTrafficQuota) Defaults() {
	def.Buffer = stringPtr("<no value>")
	def.Format = stringPtr("")
	def.Header = stringPtr("")
	def.MsgType = stringPtr("")
}
