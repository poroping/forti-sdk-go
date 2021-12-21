package models

const LogdiskSettingPath = "log.disk/setting/"

type LogdiskSetting struct {
	Diskfull                   *string  `json:"diskfull,omitempty"`
	DlpArchiveQuota            *float64 `json:"dlp-archive-quota,omitempty"`
	FullFinalWarningThreshold  *float64 `json:"full-final-warning-threshold,omitempty"`
	FullFirstWarningThreshold  *float64 `json:"full-first-warning-threshold,omitempty"`
	FullSecondWarningThreshold *float64 `json:"full-second-warning-threshold,omitempty"`
	Interface                  *string  `json:"interface,omitempty"`
	InterfaceSelectMethod      *string  `json:"interface-select-method,omitempty"`
	IpsArchive                 *string  `json:"ips-archive,omitempty"`
	LogQuota                   *float64 `json:"log-quota,omitempty"`
	MaxLogFileSize             *float64 `json:"max-log-file-size,omitempty"`
	MaxPolicyPacketCaptureSize *float64 `json:"max-policy-packet-capture-size,omitempty"`
	MaximumLogAge              *float64 `json:"maximum-log-age,omitempty"`
	ReportQuota                *float64 `json:"report-quota,omitempty"`
	RollDay                    *string  `json:"roll-day,omitempty"`
	RollSchedule               *string  `json:"roll-schedule,omitempty"`
	RollTime                   *string  `json:"roll-time,omitempty"`
	SourceIp                   *string  `json:"source-ip,omitempty"`
	Status                     *string  `json:"status,omitempty"`
	Upload                     *string  `json:"upload,omitempty"`
	UploadDeleteFiles          *string  `json:"upload-delete-files,omitempty"`
	UploadDestination          *string  `json:"upload-destination,omitempty"`
	UploadSslConn              *string  `json:"upload-ssl-conn,omitempty"`
	Uploaddir                  *string  `json:"uploaddir,omitempty"`
	Uploadip                   *string  `json:"uploadip,omitempty"`
	Uploadpass                 *string  `json:"uploadpass,omitempty"`
	Uploadport                 *float64 `json:"uploadport,omitempty"`
	Uploadsched                *string  `json:"uploadsched,omitempty"`
	Uploadtime                 *string  `json:"uploadtime,omitempty"`
	Uploadtype                 *string  `json:"uploadtype,omitempty"`
	Uploaduser                 *string  `json:"uploaduser,omitempty"`
}
