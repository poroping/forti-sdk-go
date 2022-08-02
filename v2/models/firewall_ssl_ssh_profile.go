package models

const FirewallSslSshProfilePath = "firewall/ssl-ssh-profile/"

type FirewallSslSshProfile struct {
	Allowlist                    *string                            `json:"allowlist,omitempty"`
	BlockBlacklistedCertificates *string                            `json:"block-blacklisted-certificates,omitempty"`
	BlockBlocklistedCertificates *string                            `json:"block-blocklisted-certificates,omitempty"`
	Caname                       *string                            `json:"caname,omitempty"`
	Comment                      *string                            `json:"comment,omitempty"`
	Dot                          *FirewallSslSshProfileDot          `json:"dot,omitempty"`
	Ftps                         *FirewallSslSshProfileFtps         `json:"ftps,omitempty"`
	Https                        *FirewallSslSshProfileHttps        `json:"https,omitempty"`
	Imaps                        *FirewallSslSshProfileImaps        `json:"imaps,omitempty"`
	MapiOverHttps                *string                            `json:"mapi-over-https,omitempty"`
	Name                         *string                            `json:"name,omitempty"`
	Pop3s                        *FirewallSslSshProfilePop3s        `json:"pop3s,omitempty"`
	RpcOverHttps                 *string                            `json:"rpc-over-https,omitempty"`
	ServerCert                   *[]FirewallSslSshProfileServerCert `json:"server-cert,omitempty"`
	ServerCertMode               *string                            `json:"server-cert-mode,omitempty"`
	Smtps                        *FirewallSslSshProfileSmtps        `json:"smtps,omitempty"`
	Ssh                          *FirewallSslSshProfileSsh          `json:"ssh,omitempty"`
	Ssl                          *FirewallSslSshProfileSsl          `json:"ssl,omitempty"`
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

const FirewallSslSshProfileDotPath = "firewall/ssl-ssh-profile/dot/"

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

const FirewallSslSshProfileFtpsPath = "firewall/ssl-ssh-profile/ftps/"

type FirewallSslSshProfileFtps struct {
	CertValidationFailure     *string `json:"cert-validation-failure,omitempty"`
	CertValidationTimeout     *string `json:"cert-validation-timeout,omitempty"`
	ClientCertificate         *string `json:"client-certificate,omitempty"`
	ExpiredServerCert         *string `json:"expired-server-cert,omitempty"`
	MinAllowedSslVersion      *string `json:"min-allowed-ssl-version,omitempty"`
	Ports                     *int64  `json:"ports,omitempty"`
	RevokedServerCert         *string `json:"revoked-server-cert,omitempty"`
	SniServerCertCheck        *string `json:"sni-server-cert-check,omitempty"`
	Status                    *string `json:"status,omitempty"`
	UnsupportedSslCipher      *string `json:"unsupported-ssl-cipher,omitempty"`
	UnsupportedSslNegotiation *string `json:"unsupported-ssl-negotiation,omitempty"`
	UnsupportedSslVersion     *string `json:"unsupported-ssl-version,omitempty"`
	UntrustedServerCert       *string `json:"untrusted-server-cert,omitempty"`
}

const FirewallSslSshProfileHttpsPath = "firewall/ssl-ssh-profile/https/"

type FirewallSslSshProfileHttps struct {
	CertProbeFailure          *string `json:"cert-probe-failure,omitempty"`
	CertValidationFailure     *string `json:"cert-validation-failure,omitempty"`
	CertValidationTimeout     *string `json:"cert-validation-timeout,omitempty"`
	ClientCertificate         *string `json:"client-certificate,omitempty"`
	ExpiredServerCert         *string `json:"expired-server-cert,omitempty"`
	MinAllowedSslVersion      *string `json:"min-allowed-ssl-version,omitempty"`
	Ports                     *int64  `json:"ports,omitempty"`
	ProxyAfterTcpHandshake    *string `json:"proxy-after-tcp-handshake,omitempty"`
	RevokedServerCert         *string `json:"revoked-server-cert,omitempty"`
	SniServerCertCheck        *string `json:"sni-server-cert-check,omitempty"`
	Status                    *string `json:"status,omitempty"`
	UnsupportedSslCipher      *string `json:"unsupported-ssl-cipher,omitempty"`
	UnsupportedSslNegotiation *string `json:"unsupported-ssl-negotiation,omitempty"`
	UnsupportedSslVersion     *string `json:"unsupported-ssl-version,omitempty"`
	UntrustedServerCert       *string `json:"untrusted-server-cert,omitempty"`
}

const FirewallSslSshProfileImapsPath = "firewall/ssl-ssh-profile/imaps/"

type FirewallSslSshProfileImaps struct {
	CertValidationFailure     *string `json:"cert-validation-failure,omitempty"`
	CertValidationTimeout     *string `json:"cert-validation-timeout,omitempty"`
	ClientCertificate         *string `json:"client-certificate,omitempty"`
	ExpiredServerCert         *string `json:"expired-server-cert,omitempty"`
	Ports                     *int64  `json:"ports,omitempty"`
	ProxyAfterTcpHandshake    *string `json:"proxy-after-tcp-handshake,omitempty"`
	RevokedServerCert         *string `json:"revoked-server-cert,omitempty"`
	SniServerCertCheck        *string `json:"sni-server-cert-check,omitempty"`
	Status                    *string `json:"status,omitempty"`
	UnsupportedSslCipher      *string `json:"unsupported-ssl-cipher,omitempty"`
	UnsupportedSslNegotiation *string `json:"unsupported-ssl-negotiation,omitempty"`
	UnsupportedSslVersion     *string `json:"unsupported-ssl-version,omitempty"`
	UntrustedServerCert       *string `json:"untrusted-server-cert,omitempty"`
}

const FirewallSslSshProfilePop3sPath = "firewall/ssl-ssh-profile/pop3s/"

type FirewallSslSshProfilePop3s struct {
	CertValidationFailure     *string `json:"cert-validation-failure,omitempty"`
	CertValidationTimeout     *string `json:"cert-validation-timeout,omitempty"`
	ClientCertificate         *string `json:"client-certificate,omitempty"`
	ExpiredServerCert         *string `json:"expired-server-cert,omitempty"`
	Ports                     *int64  `json:"ports,omitempty"`
	ProxyAfterTcpHandshake    *string `json:"proxy-after-tcp-handshake,omitempty"`
	RevokedServerCert         *string `json:"revoked-server-cert,omitempty"`
	SniServerCertCheck        *string `json:"sni-server-cert-check,omitempty"`
	Status                    *string `json:"status,omitempty"`
	UnsupportedSslCipher      *string `json:"unsupported-ssl-cipher,omitempty"`
	UnsupportedSslNegotiation *string `json:"unsupported-ssl-negotiation,omitempty"`
	UnsupportedSslVersion     *string `json:"unsupported-ssl-version,omitempty"`
	UntrustedServerCert       *string `json:"untrusted-server-cert,omitempty"`
}

const FirewallSslSshProfileServerCertPath = "firewall/ssl-ssh-profile/server-cert/"

type FirewallSslSshProfileServerCert struct {
	Name *string `json:"name,omitempty"`
}

const FirewallSslSshProfileSmtpsPath = "firewall/ssl-ssh-profile/smtps/"

type FirewallSslSshProfileSmtps struct {
	CertValidationFailure     *string `json:"cert-validation-failure,omitempty"`
	CertValidationTimeout     *string `json:"cert-validation-timeout,omitempty"`
	ClientCertificate         *string `json:"client-certificate,omitempty"`
	ExpiredServerCert         *string `json:"expired-server-cert,omitempty"`
	Ports                     *int64  `json:"ports,omitempty"`
	ProxyAfterTcpHandshake    *string `json:"proxy-after-tcp-handshake,omitempty"`
	RevokedServerCert         *string `json:"revoked-server-cert,omitempty"`
	SniServerCertCheck        *string `json:"sni-server-cert-check,omitempty"`
	Status                    *string `json:"status,omitempty"`
	UnsupportedSslCipher      *string `json:"unsupported-ssl-cipher,omitempty"`
	UnsupportedSslNegotiation *string `json:"unsupported-ssl-negotiation,omitempty"`
	UnsupportedSslVersion     *string `json:"unsupported-ssl-version,omitempty"`
	UntrustedServerCert       *string `json:"untrusted-server-cert,omitempty"`
}

const FirewallSslSshProfileSshPath = "firewall/ssl-ssh-profile/ssh/"

type FirewallSslSshProfileSsh struct {
	InspectAll             *string `json:"inspect-all,omitempty"`
	Ports                  *int64  `json:"ports,omitempty"`
	ProxyAfterTcpHandshake *string `json:"proxy-after-tcp-handshake,omitempty"`
	SshAlgorithm           *string `json:"ssh-algorithm,omitempty"`
	SshTunPolicyCheck      *string `json:"ssh-tun-policy-check,omitempty"`
	Status                 *string `json:"status,omitempty"`
	UnsupportedVersion     *string `json:"unsupported-version,omitempty"`
}

const FirewallSslSshProfileSslPath = "firewall/ssl-ssh-profile/ssl/"

type FirewallSslSshProfileSsl struct {
	CertProbeFailure          *string `json:"cert-probe-failure,omitempty"`
	CertValidationFailure     *string `json:"cert-validation-failure,omitempty"`
	CertValidationTimeout     *string `json:"cert-validation-timeout,omitempty"`
	ClientCertificate         *string `json:"client-certificate,omitempty"`
	ExpiredServerCert         *string `json:"expired-server-cert,omitempty"`
	InspectAll                *string `json:"inspect-all,omitempty"`
	MinAllowedSslVersion      *string `json:"min-allowed-ssl-version,omitempty"`
	RevokedServerCert         *string `json:"revoked-server-cert,omitempty"`
	SniServerCertCheck        *string `json:"sni-server-cert-check,omitempty"`
	UnsupportedSslCipher      *string `json:"unsupported-ssl-cipher,omitempty"`
	UnsupportedSslNegotiation *string `json:"unsupported-ssl-negotiation,omitempty"`
	UnsupportedSslVersion     *string `json:"unsupported-ssl-version,omitempty"`
	UntrustedServerCert       *string `json:"untrusted-server-cert,omitempty"`
}

const FirewallSslSshProfileSslExemptPath = "firewall/ssl-ssh-profile/ssl-exempt/"

type FirewallSslSshProfileSslExempt struct {
	Address            *string `json:"address,omitempty"`
	Address6           *string `json:"address6,omitempty"`
	FortiguardCategory *int64  `json:"fortiguard-category,omitempty"`
	Id                 *int64  `json:"id,omitempty"`
	Regex              *string `json:"regex,omitempty"`
	Type               *string `json:"type,omitempty"`
	WildcardFqdn       *string `json:"wildcard-fqdn,omitempty"`
}

const FirewallSslSshProfileSslServerPath = "firewall/ssl-ssh-profile/ssl-server/"

type FirewallSslSshProfileSslServer struct {
	FtpsClientCertificate     *string `json:"ftps-client-certificate,omitempty"`
	HttpsClientCertificate    *string `json:"https-client-certificate,omitempty"`
	Id                        *int64  `json:"id,omitempty"`
	ImapsClientCertificate    *string `json:"imaps-client-certificate,omitempty"`
	Ip                        *string `json:"ip,omitempty"`
	Pop3sClientCertificate    *string `json:"pop3s-client-certificate,omitempty"`
	SmtpsClientCertificate    *string `json:"smtps-client-certificate,omitempty"`
	SslOtherClientCertificate *string `json:"ssl-other-client-certificate,omitempty"`
}
