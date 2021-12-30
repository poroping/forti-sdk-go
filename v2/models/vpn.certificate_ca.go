package models

const VpncertificateCaPath = "vpn/certificate/ca/"

type VpncertificateCa struct {
	AutoUpdateDays        *int64  `json:"auto-update-days,omitempty"`
	AutoUpdateDaysWarning *int64  `json:"auto-update-days-warning,omitempty"`
	Ca                    *string `json:"ca,omitempty"`
	CaIdentifier          *string `json:"ca-identifier,omitempty"`
	Name                  *string `json:"name,omitempty"`
	Range                 *string `json:"range,omitempty"`
	ScepUrl               *string `json:"scep-url,omitempty"`
	Source                *string `json:"source,omitempty"`
	SourceIp              *string `json:"source-ip,omitempty"`
	SslInspectionTrusted  *string `json:"ssl-inspection-trusted,omitempty"`
}
