package models

const LogSyslogd3OverrideSettingPath = "log.syslogd3/override-setting/"

type LogSyslogd3OverrideSetting struct {
	Certificate           *string                                      `json:"certificate,omitempty"`
	CustomFieldName       *[]LogSyslogd3OverrideSettingCustomFieldName `json:"custom-field-name,omitempty"`
	EncAlgorithm          *string                                      `json:"enc-algorithm,omitempty"`
	Facility              *string                                      `json:"facility,omitempty"`
	Format                *string                                      `json:"format,omitempty"`
	Interface             *string                                      `json:"interface,omitempty"`
	InterfaceSelectMethod *string                                      `json:"interface-select-method,omitempty"`
	MaxLogRate            *int64                                       `json:"max-log-rate,omitempty"`
	Mode                  *string                                      `json:"mode,omitempty"`
	Port                  *int64                                       `json:"port,omitempty"`
	Priority              *string                                      `json:"priority,omitempty"`
	Server                *string                                      `json:"server,omitempty"`
	SourceIp              *string                                      `json:"source-ip,omitempty"`
	SslMinProtoVersion    *string                                      `json:"ssl-min-proto-version,omitempty"`
	Status                *string                                      `json:"status,omitempty"`
}

const LogSyslogd3OverrideSettingCustomFieldNamePath = "log.syslogd3/override-setting/custom-field-name/"

type LogSyslogd3OverrideSettingCustomFieldName struct {
	Custom *string `json:"custom,omitempty"`
	Id     *int64  `json:"id,omitempty"`
	Name   *string `json:"name,omitempty"`
}
