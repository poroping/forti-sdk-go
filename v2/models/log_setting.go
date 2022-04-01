package models

const LogSettingPath = "log/setting/"

type LogSetting struct {
	AnonymizationHash    *string                      `json:"anonymization-hash,omitempty"`
	BriefTrafficFormat   *string                      `json:"brief-traffic-format,omitempty"`
	CustomLogFields      *[]LogSettingCustomLogFields `json:"custom-log-fields,omitempty"`
	DaemonLog            *string                      `json:"daemon-log,omitempty"`
	ExpolicyImplicitLog  *string                      `json:"expolicy-implicit-log,omitempty"`
	FazOverride          *string                      `json:"faz-override,omitempty"`
	FortiviewWeeklyData  *string                      `json:"fortiview-weekly-data,omitempty"`
	FwpolicyImplicitLog  *string                      `json:"fwpolicy-implicit-log,omitempty"`
	Fwpolicy6ImplicitLog *string                      `json:"fwpolicy6-implicit-log,omitempty"`
	LocalInAllow         *string                      `json:"local-in-allow,omitempty"`
	LocalInDenyBroadcast *string                      `json:"local-in-deny-broadcast,omitempty"`
	LocalInDenyUnicast   *string                      `json:"local-in-deny-unicast,omitempty"`
	LocalOut             *string                      `json:"local-out,omitempty"`
	LocalOutIocDetection *string                      `json:"local-out-ioc-detection,omitempty"`
	LogInvalidPacket     *string                      `json:"log-invalid-packet,omitempty"`
	LogPolicyComment     *string                      `json:"log-policy-comment,omitempty"`
	LogPolicyName        *string                      `json:"log-policy-name,omitempty"`
	LogUserInUpper       *string                      `json:"log-user-in-upper,omitempty"`
	NeighborEvent        *string                      `json:"neighbor-event,omitempty"`
	ResolveIp            *string                      `json:"resolve-ip,omitempty"`
	ResolvePort          *string                      `json:"resolve-port,omitempty"`
	RestApiGet           *string                      `json:"rest-api-get,omitempty"`
	RestApiSet           *string                      `json:"rest-api-set,omitempty"`
	SyslogOverride       *string                      `json:"syslog-override,omitempty"`
	UserAnonymize        *string                      `json:"user-anonymize,omitempty"`
}

type LogSettingCustomLogFields struct {
	FieldId *string `json:"field-id,omitempty"`
}
