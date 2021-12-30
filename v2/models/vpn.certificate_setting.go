package models

const VpncertificateSettingPath = "vpn/certificate/setting/"

type VpncertificateSetting struct {
	CertnameDsa1024       *string                                 `json:"certname-dsa1024,omitempty"`
	CertnameDsa2048       *string                                 `json:"certname-dsa2048,omitempty"`
	CertnameEcdsa256      *string                                 `json:"certname-ecdsa256,omitempty"`
	CertnameEcdsa384      *string                                 `json:"certname-ecdsa384,omitempty"`
	CertnameEcdsa521      *string                                 `json:"certname-ecdsa521,omitempty"`
	CertnameEd25519       *string                                 `json:"certname-ed25519,omitempty"`
	CertnameEd448         *string                                 `json:"certname-ed448,omitempty"`
	CertnameRsa1024       *string                                 `json:"certname-rsa1024,omitempty"`
	CertnameRsa2048       *string                                 `json:"certname-rsa2048,omitempty"`
	CertnameRsa4096       *string                                 `json:"certname-rsa4096,omitempty"`
	CheckCaCert           *string                                 `json:"check-ca-cert,omitempty"`
	CheckCaChain          *string                                 `json:"check-ca-chain,omitempty"`
	CmpKeyUsageChecking   *string                                 `json:"cmp-key-usage-checking,omitempty"`
	CmpSaveExtraCerts     *string                                 `json:"cmp-save-extra-certs,omitempty"`
	CnAllowMulti          *string                                 `json:"cn-allow-multi,omitempty"`
	CnMatch               *string                                 `json:"cn-match,omitempty"`
	CrlVerification       *[]VpncertificateSettingCrlVerification `json:"crl-verification,omitempty"`
	Interface             *string                                 `json:"interface,omitempty"`
	InterfaceSelectMethod *string                                 `json:"interface-select-method,omitempty"`
	OcspDefaultServer     *string                                 `json:"ocsp-default-server,omitempty"`
	OcspOption            *string                                 `json:"ocsp-option,omitempty"`
	OcspStatus            *string                                 `json:"ocsp-status,omitempty"`
	SslMinProtoVersion    *string                                 `json:"ssl-min-proto-version,omitempty"`
	SslOcspSourceIp       *string                                 `json:"ssl-ocsp-source-ip,omitempty"`
	StrictCrlCheck        *string                                 `json:"strict-crl-check,omitempty"`
	StrictOcspCheck       *string                                 `json:"strict-ocsp-check,omitempty"`
	SubjectMatch          *string                                 `json:"subject-match,omitempty"`
	SubjectSet            *string                                 `json:"subject-set,omitempty"`
}

type VpncertificateSettingCrlVerification struct {
	ChainCrlAbsence *string `json:"chain-crl-absence,omitempty"`
	Expiry          *string `json:"expiry,omitempty"`
	LeafCrlAbsence  *string `json:"leaf-crl-absence,omitempty"`
}
