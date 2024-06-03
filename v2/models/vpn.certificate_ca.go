package models

const VpnCertificateCaPath = "vpn.certificate/ca/"

type VpnCertificateCa struct {
	AutoUpdateDays        *int64  `json:"auto-update-days,omitempty"`
	AutoUpdateDaysWarning *int64  `json:"auto-update-days-warning,omitempty"`
	Ca                    *string `json:"ca,omitempty"`
	CaIdentifier          *string `json:"ca-identifier,omitempty"`
	Details               *string `json:"details,omitempty"`
	Name                  *string `json:"name,omitempty"`
	Obsolete              *string `json:"obsolete,omitempty"`
	Range                 *string `json:"range,omitempty"`
	ScepUrl               *string `json:"scep-url,omitempty"`
	Source                *string `json:"source,omitempty"`
	SourceIp              *string `json:"source-ip,omitempty"`
	SslInspectionTrusted  *string `json:"ssl-inspection-trusted,omitempty"`
}
