package models

const FirewallSslSshProfilePath = "firewall/ssl-ssh-profile/"

type FirewallSslSshProfile struct {
	Allowlist                    *string                            `json:"allowlist,omitempty"`
	BlockBlacklistedCertificates *string                            `json:"block-blacklisted-certificates,omitempty"`
	BlockBlocklistedCertificates *string                            `json:"block-blocklisted-certificates,omitempty"`
	Caname                       *string                            `json:"caname,omitempty"`
	Comment                      *string                            `json:"comment,omitempty"`
	Dot                          *[]FirewallSslSshProfileDot        `json:"dot,omitempty"`
	Ftps                         *[]FirewallSslSshProfileFtps       `json:"ftps,omitempty"`
	Https                        *[]FirewallSslSshProfileHttps      `json:"https,omitempty"`
	Imaps                        *[]FirewallSslSshProfileImaps      `json:"imaps,omitempty"`
	MapiOverHttps                *string                            `json:"mapi-over-https,omitempty"`
	Name                         *string                            `json:"name,omitempty"`
	Pop3s                        *[]FirewallSslSshProfilePop3s      `json:"pop3s,omitempty"`
	RpcOverHttps                 *string                            `json:"rpc-over-https,omitempty"`
	ServerCert                   *[]FirewallSslSshProfileServerCert `json:"server-cert,omitempty"`
	ServerCertMode               *string                            `json:"server-cert-mode,omitempty"`
	Smtps                        *[]FirewallSslSshProfileSmtps      `json:"smtps,omitempty"`
	Ssh                          *[]FirewallSslSshProfileSsh        `json:"ssh,omitempty"`
	Ssl                          *[]FirewallSslSshProfileSsl        `json:"ssl,omitempty"`
	SslAnomaliesLog              *string                            `json:"ssl-anomalies-log,omitempty"`
	SslAnomalyLog                *string                            `json:"ssl-anomaly-log,omitempty"`
	SslExempt                    *[]FirewallSslSshProfileSslExempt  `json:"ssl-exempt,omitempty"`
	SslExemptionLog              *string                            `json:"ssl-exemption-log,omitempty"`
	SslExemptionsLog             *string                            `json:"ssl-exemptions-log,omitempty"`
	SslHandshakeLog              *string                            `json:"ssl-handshake-log,omitempty"`
	SslNegotiationLog            *string                            `json:"ssl-negotiation-log,omitempty"`
	SslServer                    *[]FirewallSslSshProfileSslServer  `json:"ssl-server,omitempty"`
	SslServerCertLog             *string                            `json:"ssl-server-cert-log,omitempty"`
	SupportedAlpn                *string                            `json:"supported-alpn,omitempty"`
	UntrustedCaname              *string                            `json:"untrusted-caname,omitempty"`
	UseSslServer                 *string                            `json:"use-ssl-server,omitempty"`
	Whitelist                    *string                            `json:"whitelist,omitempty"`
}

type FirewallSslSshProfileDot struct {
	CertValidationFailure     *string `json:"cert-validation-failure,omitempty"`
	CertValidationTimeout     *string `json:"cert-validation-timeout,omitempty"`
	ClientCertificate         *string `json:"client-certificate,omitempty"`
	ExpiredServerCert         *string `json:"expired-server-cert,omitempty"`
	ProxyAfterTcpHandshake    *string `json:"proxy-after-tcp-handshake,omitempty"`
	RevokedServerCert         *string `json:"revoked-server-cert,omitempty"`
	SniServerCertCheck        *string `json:"sni-server-cert-check,omitempty"`
	Status                    *string `json:"status,omitempty"`
	UnsupportedSslCipher      *string `json:"unsupported-ssl-cipher,omitempty"`
	UnsupportedSslNegotiation *string `json:"unsupported-ssl-negotiation,omitempty"`
	UnsupportedSslVersion     *string `json:"unsupported-ssl-version,omitempty"`
	UntrustedServerCert       *string `json:"untrusted-server-cert,omitempty"`
}

type FirewallSslSshProfileFtps struct {
	CertValidationFailure     *string `json:"cert-validation-failure,omitempty"`
	CertValidationTimeout     *string `json:"cert-validation-timeout,omitempty"`
	ClientCertRequest         *string `json:"client-cert-request,omitempty"`
	ClientCertificate         *string `json:"client-certificate,omitempty"`
	ExpiredServerCert         *string `json:"expired-server-cert,omitempty"`
	InvalidServerCert         *string `json:"invalid-server-cert,omitempty"`
	Ports                     *int64  `json:"ports,omitempty"`
	RevokedServerCert         *string `json:"revoked-server-cert,omitempty"`
	SniServerCertCheck        *string `json:"sni-server-cert-check,omitempty"`
	Status                    *string `json:"status,omitempty"`
	UnsupportedSsl            *string `json:"unsupported-ssl,omitempty"`
	UnsupportedSslCipher      *string `json:"unsupported-ssl-cipher,omitempty"`
	UnsupportedSslNegotiation *string `json:"unsupported-ssl-negotiation,omitempty"`
	UnsupportedSslVersion     *string `json:"unsupported-ssl-version,omitempty"`
	UntrustedServerCert       *string `json:"untrusted-server-cert,omitempty"`
}

type FirewallSslSshProfileHttps struct {
	CertProbeFailure          *string `json:"cert-probe-failure,omitempty"`
	CertValidationFailure     *string `json:"cert-validation-failure,omitempty"`
	CertValidationTimeout     *string `json:"cert-validation-timeout,omitempty"`
	ClientCertRequest         *string `json:"client-cert-request,omitempty"`
	ClientCertificate         *string `json:"client-certificate,omitempty"`
	ExpiredServerCert         *string `json:"expired-server-cert,omitempty"`
	InvalidServerCert         *string `json:"invalid-server-cert,omitempty"`
	Ports                     *int64  `json:"ports,omitempty"`
	ProxyAfterTcpHandshake    *string `json:"proxy-after-tcp-handshake,omitempty"`
	RevokedServerCert         *string `json:"revoked-server-cert,omitempty"`
	SniServerCertCheck        *string `json:"sni-server-cert-check,omitempty"`
	Status                    *string `json:"status,omitempty"`
	UnsupportedSsl            *string `json:"unsupported-ssl,omitempty"`
	UnsupportedSslCipher      *string `json:"unsupported-ssl-cipher,omitempty"`
	UnsupportedSslNegotiation *string `json:"unsupported-ssl-negotiation,omitempty"`
	UnsupportedSslVersion     *string `json:"unsupported-ssl-version,omitempty"`
	UntrustedServerCert       *string `json:"untrusted-server-cert,omitempty"`
}

type FirewallSslSshProfileImaps struct {
	CertValidationFailure     *string `json:"cert-validation-failure,omitempty"`
	CertValidationTimeout     *string `json:"cert-validation-timeout,omitempty"`
	ClientCertRequest         *string `json:"client-cert-request,omitempty"`
	ClientCertificate         *string `json:"client-certificate,omitempty"`
	ExpiredServerCert         *string `json:"expired-server-cert,omitempty"`
	InvalidServerCert         *string `json:"invalid-server-cert,omitempty"`
	Ports                     *int64  `json:"ports,omitempty"`
	ProxyAfterTcpHandshake    *string `json:"proxy-after-tcp-handshake,omitempty"`
	RevokedServerCert         *string `json:"revoked-server-cert,omitempty"`
	SniServerCertCheck        *string `json:"sni-server-cert-check,omitempty"`
	Status                    *string `json:"status,omitempty"`
	UnsupportedSsl            *string `json:"unsupported-ssl,omitempty"`
	UnsupportedSslCipher      *string `json:"unsupported-ssl-cipher,omitempty"`
	UnsupportedSslNegotiation *string `json:"unsupported-ssl-negotiation,omitempty"`
	UnsupportedSslVersion     *string `json:"unsupported-ssl-version,omitempty"`
	UntrustedServerCert       *string `json:"untrusted-server-cert,omitempty"`
}

type FirewallSslSshProfilePop3s struct {
	CertValidationFailure     *string `json:"cert-validation-failure,omitempty"`
	CertValidationTimeout     *string `json:"cert-validation-timeout,omitempty"`
	ClientCertRequest         *string `json:"client-cert-request,omitempty"`
	ClientCertificate         *string `json:"client-certificate,omitempty"`
	ExpiredServerCert         *string `json:"expired-server-cert,omitempty"`
	InvalidServerCert         *string `json:"invalid-server-cert,omitempty"`
	Ports                     *int64  `json:"ports,omitempty"`
	ProxyAfterTcpHandshake    *string `json:"proxy-after-tcp-handshake,omitempty"`
	RevokedServerCert         *string `json:"revoked-server-cert,omitempty"`
	SniServerCertCheck        *string `json:"sni-server-cert-check,omitempty"`
	Status                    *string `json:"status,omitempty"`
	UnsupportedSsl            *string `json:"unsupported-ssl,omitempty"`
	UnsupportedSslCipher      *string `json:"unsupported-ssl-cipher,omitempty"`
	UnsupportedSslNegotiation *string `json:"unsupported-ssl-negotiation,omitempty"`
	UnsupportedSslVersion     *string `json:"unsupported-ssl-version,omitempty"`
	UntrustedServerCert       *string `json:"untrusted-server-cert,omitempty"`
}

type FirewallSslSshProfileServerCert struct {
	Name *string `json:"name,omitempty"`
}

type FirewallSslSshProfileSmtps struct {
	CertValidationFailure     *string `json:"cert-validation-failure,omitempty"`
	CertValidationTimeout     *string `json:"cert-validation-timeout,omitempty"`
	ClientCertRequest         *string `json:"client-cert-request,omitempty"`
	ClientCertificate         *string `json:"client-certificate,omitempty"`
	ExpiredServerCert         *string `json:"expired-server-cert,omitempty"`
	InvalidServerCert         *string `json:"invalid-server-cert,omitempty"`
	Ports                     *int64  `json:"ports,omitempty"`
	ProxyAfterTcpHandshake    *string `json:"proxy-after-tcp-handshake,omitempty"`
	RevokedServerCert         *string `json:"revoked-server-cert,omitempty"`
	SniServerCertCheck        *string `json:"sni-server-cert-check,omitempty"`
	Status                    *string `json:"status,omitempty"`
	UnsupportedSsl            *string `json:"unsupported-ssl,omitempty"`
	UnsupportedSslCipher      *string `json:"unsupported-ssl-cipher,omitempty"`
	UnsupportedSslNegotiation *string `json:"unsupported-ssl-negotiation,omitempty"`
	UnsupportedSslVersion     *string `json:"unsupported-ssl-version,omitempty"`
	UntrustedServerCert       *string `json:"untrusted-server-cert,omitempty"`
}

type FirewallSslSshProfileSsh struct {
	InspectAll             *string `json:"inspect-all,omitempty"`
	Ports                  *int64  `json:"ports,omitempty"`
	ProxyAfterTcpHandshake *string `json:"proxy-after-tcp-handshake,omitempty"`
	SshAlgorithm           *string `json:"ssh-algorithm,omitempty"`
	SshTunPolicyCheck      *string `json:"ssh-tun-policy-check,omitempty"`
	Status                 *string `json:"status,omitempty"`
	UnsupportedVersion     *string `json:"unsupported-version,omitempty"`
}

type FirewallSslSshProfileSsl struct {
	CertProbeFailure          *string `json:"cert-probe-failure,omitempty"`
	CertValidationFailure     *string `json:"cert-validation-failure,omitempty"`
	CertValidationTimeout     *string `json:"cert-validation-timeout,omitempty"`
	ClientCertRequest         *string `json:"client-cert-request,omitempty"`
	ClientCertificate         *string `json:"client-certificate,omitempty"`
	ExpiredServerCert         *string `json:"expired-server-cert,omitempty"`
	InspectAll                *string `json:"inspect-all,omitempty"`
	InvalidServerCert         *string `json:"invalid-server-cert,omitempty"`
	RevokedServerCert         *string `json:"revoked-server-cert,omitempty"`
	SniServerCertCheck        *string `json:"sni-server-cert-check,omitempty"`
	UnsupportedSsl            *string `json:"unsupported-ssl,omitempty"`
	UnsupportedSslCipher      *string `json:"unsupported-ssl-cipher,omitempty"`
	UnsupportedSslNegotiation *string `json:"unsupported-ssl-negotiation,omitempty"`
	UnsupportedSslVersion     *string `json:"unsupported-ssl-version,omitempty"`
	UntrustedServerCert       *string `json:"untrusted-server-cert,omitempty"`
}

type FirewallSslSshProfileSslExempt struct {
	Address            *string `json:"address,omitempty"`
	Address6           *string `json:"address6,omitempty"`
	FortiguardCategory *int64  `json:"fortiguard-category,omitempty"`
	Id                 *int64  `json:"id,omitempty"`
	Regex              *string `json:"regex,omitempty"`
	Type               *string `json:"type,omitempty"`
	WildcardFqdn       *string `json:"wildcard-fqdn,omitempty"`
}

type FirewallSslSshProfileSslServer struct {
	FtpsClientCertRequest     *string `json:"ftps-client-cert-request,omitempty"`
	FtpsClientCertificate     *string `json:"ftps-client-certificate,omitempty"`
	HttpsClientCertRequest    *string `json:"https-client-cert-request,omitempty"`
	HttpsClientCertificate    *string `json:"https-client-certificate,omitempty"`
	Id                        *int64  `json:"id,omitempty"`
	ImapsClientCertRequest    *string `json:"imaps-client-cert-request,omitempty"`
	ImapsClientCertificate    *string `json:"imaps-client-certificate,omitempty"`
	Ip                        *string `json:"ip,omitempty"`
	Pop3sClientCertRequest    *string `json:"pop3s-client-cert-request,omitempty"`
	Pop3sClientCertificate    *string `json:"pop3s-client-certificate,omitempty"`
	SmtpsClientCertRequest    *string `json:"smtps-client-cert-request,omitempty"`
	SmtpsClientCertificate    *string `json:"smtps-client-certificate,omitempty"`
	SslOtherClientCertRequest *string `json:"ssl-other-client-cert-request,omitempty"`
	SslOtherClientCertificate *string `json:"ssl-other-client-certificate,omitempty"`
}
