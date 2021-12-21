package models

const CertificateCaPath = "certificate/ca/"

type CertificateCa struct {
	AutoUpdateDays        *float64 `json:"auto-update-days,omitempty"`
	AutoUpdateDaysWarning *float64 `json:"auto-update-days-warning,omitempty"`
	Ca                    *string  `json:"ca,omitempty"`
	CaIdentifier          *string  `json:"ca-identifier,omitempty"`
	Name                  *string  `json:"name,omitempty"`
	Range                 *string  `json:"range,omitempty"`
	ScepUrl               *string  `json:"scep-url,omitempty"`
	Source                *string  `json:"source,omitempty"`
	SourceIp              *string  `json:"source-ip,omitempty"`
	SslInspectionTrusted  *string  `json:"ssl-inspection-trusted,omitempty"`
}
