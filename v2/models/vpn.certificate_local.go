package models

const VpnCertificateLocalPath = "vpn.certificate/local/"

type VpnCertificateLocal struct {
	AcmeCaUrl                 *string `json:"acme-ca-url,omitempty"`
	AcmeDomain                *string `json:"acme-domain,omitempty"`
	AcmeEmail                 *string `json:"acme-email,omitempty"`
	AcmeRenewWindow           *int64  `json:"acme-renew-window,omitempty"`
	AcmeRsaKeySize            *int64  `json:"acme-rsa-key-size,omitempty"`
	AutoRegenerateDays        *int64  `json:"auto-regenerate-days,omitempty"`
	AutoRegenerateDaysWarning *int64  `json:"auto-regenerate-days-warning,omitempty"`
	CaIdentifier              *string `json:"ca-identifier,omitempty"`
	Certificate               *string `json:"certificate,omitempty"`
	CmpPath                   *string `json:"cmp-path,omitempty"`
	CmpRegenerationMethod     *string `json:"cmp-regeneration-method,omitempty"`
	CmpServer                 *string `json:"cmp-server,omitempty"`
	CmpServerCert             *string `json:"cmp-server-cert,omitempty"`
	Comments                  *string `json:"comments,omitempty"`
	Csr                       *string `json:"csr,omitempty"`
	EnrollProtocol            *string `json:"enroll-protocol,omitempty"`
	IkeLocalid                *string `json:"ike-localid,omitempty"`
	IkeLocalidType            *string `json:"ike-localid-type,omitempty"`
	Name                      *string `json:"name,omitempty"`
	NameEncoding              *string `json:"name-encoding,omitempty"`
	Password                  *string `json:"password,omitempty"`
	PrivateKey                *string `json:"private-key,omitempty"`
	Range                     *string `json:"range,omitempty"`
	ScepPassword              *string `json:"scep-password,omitempty"`
	ScepUrl                   *string `json:"scep-url,omitempty"`
	Source                    *string `json:"source,omitempty"`
	SourceIp                  *string `json:"source-ip,omitempty"`
	State                     *string `json:"state,omitempty"`
}
