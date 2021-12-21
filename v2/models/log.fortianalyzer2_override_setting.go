package models

const Logfortianalyzer2OverrideSettingPath = "log.fortianalyzer2/override-setting/"

type Logfortianalyzer2OverrideSetting struct {
	AccessConfig              *string                                   `json:"access-config,omitempty"`
	Certificate               *string                                   `json:"certificate,omitempty"`
	CertificateVerification   *string                                   `json:"certificate-verification,omitempty"`
	ConnTimeout               *float64                                  `json:"conn-timeout,omitempty"`
	EncAlgorithm              *string                                   `json:"enc-algorithm,omitempty"`
	HmacAlgorithm             *string                                   `json:"hmac-algorithm,omitempty"`
	Interface                 *string                                   `json:"interface,omitempty"`
	InterfaceSelectMethod     *string                                   `json:"interface-select-method,omitempty"`
	IpsArchive                *string                                   `json:"ips-archive,omitempty"`
	MaxLogRate                *float64                                  `json:"max-log-rate,omitempty"`
	MonitorFailureRetryPeriod *float64                                  `json:"monitor-failure-retry-period,omitempty"`
	MonitorKeepalivePeriod    *float64                                  `json:"monitor-keepalive-period,omitempty"`
	PresharedKey              *string                                   `json:"preshared-key,omitempty"`
	Priority                  *string                                   `json:"priority,omitempty"`
	Reliable                  *string                                   `json:"reliable,omitempty"`
	Serial                    *[]Logfortianalyzer2OverrideSettingSerial `json:"serial,omitempty"`
	Server                    *string                                   `json:"server,omitempty"`
	SourceIp                  *string                                   `json:"source-ip,omitempty"`
	SslMinProtoVersion        *string                                   `json:"ssl-min-proto-version,omitempty"`
	Status                    *string                                   `json:"status,omitempty"`
	UploadDay                 *string                                   `json:"upload-day,omitempty"`
	UploadInterval            *string                                   `json:"upload-interval,omitempty"`
	UploadOption              *string                                   `json:"upload-option,omitempty"`
	UploadTime                *string                                   `json:"upload-time,omitempty"`
	UseManagementVdom         *string                                   `json:"use-management-vdom,omitempty"`
}

type Logfortianalyzer2OverrideSettingSerial struct {
	Name *string `json:"name,omitempty"`
}
