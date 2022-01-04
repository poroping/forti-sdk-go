package models

const VpnCertificateOcspServerPath = "vpn.certificate/ocsp-server/"

type VpnCertificateOcspServer struct {
	Cert          *string `json:"cert,omitempty"`
	Name          *string `json:"name,omitempty"`
	SecondaryCert *string `json:"secondary-cert,omitempty"`
	SecondaryUrl  *string `json:"secondary-url,omitempty"`
	SourceIp      *string `json:"source-ip,omitempty"`
	UnavailAction *string `json:"unavail-action,omitempty"`
	Url           *string `json:"url,omitempty"`
}
