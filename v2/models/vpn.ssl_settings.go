package models

const VpnSslSettingsPath = "vpn.ssl/settings/"

type VpnSslSettings struct {
	Algorithm                  *string                             `json:"algorithm,omitempty"`
	AuthSessionCheckSourceIp   *string                             `json:"auth-session-check-source-ip,omitempty"`
	AuthTimeout                *int64                              `json:"auth-timeout,omitempty"`
	AuthenticationRule         *[]VpnSslSettingsAuthenticationRule `json:"authentication-rule,omitempty"`
	AutoTunnelStaticRoute      *string                             `json:"auto-tunnel-static-route,omitempty"`
	BannedCipher               *string                             `json:"banned-cipher,omitempty"`
	BrowserLanguageDetection   *string                             `json:"browser-language-detection,omitempty"`
	CheckReferer               *string                             `json:"check-referer,omitempty"`
	Ciphersuite                *string                             `json:"ciphersuite,omitempty"`
	ClientSigalgs              *string                             `json:"client-sigalgs,omitempty"`
	DefaultPortal              *string                             `json:"default-portal,omitempty"`
	DeflateCompressionLevel    *int64                              `json:"deflate-compression-level,omitempty"`
	DeflateMinDataSize         *int64                              `json:"deflate-min-data-size,omitempty"`
	DnsServer1                 *string                             `json:"dns-server1,omitempty"`
	DnsServer2                 *string                             `json:"dns-server2,omitempty"`
	DnsSuffix                  *string                             `json:"dns-suffix,omitempty"`
	DtlsHelloTimeout           *int64                              `json:"dtls-hello-timeout,omitempty"`
	DtlsMaxProtoVer            *string                             `json:"dtls-max-proto-ver,omitempty"`
	DtlsMinProtoVer            *string                             `json:"dtls-min-proto-ver,omitempty"`
	DtlsTunnel                 *string                             `json:"dtls-tunnel,omitempty"`
	DualStackMode              *string                             `json:"dual-stack-mode,omitempty"`
	Encode2fSequence           *string                             `json:"encode-2f-sequence,omitempty"`
	EncryptAndStorePassword    *string                             `json:"encrypt-and-store-password,omitempty"`
	ForceTwoFactorAuth         *string                             `json:"force-two-factor-auth,omitempty"`
	HeaderXForwardedFor        *string                             `json:"header-x-forwarded-for,omitempty"`
	HstsIncludeSubdomains      *string                             `json:"hsts-include-subdomains,omitempty"`
	HttpCompression            *string                             `json:"http-compression,omitempty"`
	HttpOnlyCookie             *string                             `json:"http-only-cookie,omitempty"`
	HttpRequestBodyTimeout     *int64                              `json:"http-request-body-timeout,omitempty"`
	HttpRequestHeaderTimeout   *int64                              `json:"http-request-header-timeout,omitempty"`
	HttpsRedirect              *string                             `json:"https-redirect,omitempty"`
	IdleTimeout                *int64                              `json:"idle-timeout,omitempty"`
	Ipv6DnsServer1             *string                             `json:"ipv6-dns-server1,omitempty"`
	Ipv6DnsServer2             *string                             `json:"ipv6-dns-server2,omitempty"`
	Ipv6WinsServer1            *string                             `json:"ipv6-wins-server1,omitempty"`
	Ipv6WinsServer2            *string                             `json:"ipv6-wins-server2,omitempty"`
	LoginAttemptLimit          *int64                              `json:"login-attempt-limit,omitempty"`
	LoginBlockTime             *int64                              `json:"login-block-time,omitempty"`
	LoginTimeout               *int64                              `json:"login-timeout,omitempty"`
	Port                       *int64                              `json:"port,omitempty"`
	PortPrecedence             *string                             `json:"port-precedence,omitempty"`
	Reqclientcert              *string                             `json:"reqclientcert,omitempty"`
	RouteSourceInterface       *string                             `json:"route-source-interface,omitempty"`
	SamlRedirectPort           *int64                              `json:"saml-redirect-port,omitempty"`
	ServerHostname             *string                             `json:"server-hostname,omitempty"`
	Servercert                 *string                             `json:"servercert,omitempty"`
	SourceAddress              *[]VpnSslSettingsSourceAddress      `json:"source-address,omitempty"`
	SourceAddressNegate        *string                             `json:"source-address-negate,omitempty"`
	SourceAddress6             *[]VpnSslSettingsSourceAddress6     `json:"source-address6,omitempty"`
	SourceAddress6Negate       *string                             `json:"source-address6-negate,omitempty"`
	SourceInterface            *[]VpnSslSettingsSourceInterface    `json:"source-interface,omitempty"`
	SslClientRenegotiation     *string                             `json:"ssl-client-renegotiation,omitempty"`
	SslInsertEmptyFragment     *string                             `json:"ssl-insert-empty-fragment,omitempty"`
	SslMaxProtoVer             *string                             `json:"ssl-max-proto-ver,omitempty"`
	SslMinProtoVer             *string                             `json:"ssl-min-proto-ver,omitempty"`
	Status                     *string                             `json:"status,omitempty"`
	Tlsv10                     *string                             `json:"tlsv1-0,omitempty"`
	Tlsv11                     *string                             `json:"tlsv1-1,omitempty"`
	Tlsv12                     *string                             `json:"tlsv1-2,omitempty"`
	Tlsv13                     *string                             `json:"tlsv1-3,omitempty"`
	TransformBackwardSlashes   *string                             `json:"transform-backward-slashes,omitempty"`
	TunnelAddrAssignedMethod   *string                             `json:"tunnel-addr-assigned-method,omitempty"`
	TunnelConnectWithoutReauth *string                             `json:"tunnel-connect-without-reauth,omitempty"`
	TunnelIpPools              *[]VpnSslSettingsTunnelIpPools      `json:"tunnel-ip-pools,omitempty"`
	TunnelIpv6Pools            *[]VpnSslSettingsTunnelIpv6Pools    `json:"tunnel-ipv6-pools,omitempty"`
	TunnelUserSessionTimeout   *int64                              `json:"tunnel-user-session-timeout,omitempty"`
	UnsafeLegacyRenegotiation  *string                             `json:"unsafe-legacy-renegotiation,omitempty"`
	UrlObscuration             *string                             `json:"url-obscuration,omitempty"`
	UserPeer                   *string                             `json:"user-peer,omitempty"`
	WebModeSnat                *string                             `json:"web-mode-snat,omitempty"`
	WinsServer1                *string                             `json:"wins-server1,omitempty"`
	WinsServer2                *string                             `json:"wins-server2,omitempty"`
	XContentTypeOptions        *string                             `json:"x-content-type-options,omitempty"`
	ZtnaTrustedClient          *string                             `json:"ztna-trusted-client,omitempty"`
}

type VpnSslSettingsAuthenticationRule struct {
	Auth                 *string                                            `json:"auth,omitempty"`
	Cipher               *string                                            `json:"cipher,omitempty"`
	ClientCert           *string                                            `json:"client-cert,omitempty"`
	Groups               *[]VpnSslSettingsAuthenticationRuleGroups          `json:"groups,omitempty"`
	Id                   *int64                                             `json:"id,omitempty"`
	Portal               *string                                            `json:"portal,omitempty"`
	Realm                *string                                            `json:"realm,omitempty"`
	SourceAddress        *[]VpnSslSettingsAuthenticationRuleSourceAddress   `json:"source-address,omitempty"`
	SourceAddressNegate  *string                                            `json:"source-address-negate,omitempty"`
	SourceAddress6       *[]VpnSslSettingsAuthenticationRuleSourceAddress6  `json:"source-address6,omitempty"`
	SourceAddress6Negate *string                                            `json:"source-address6-negate,omitempty"`
	SourceInterface      *[]VpnSslSettingsAuthenticationRuleSourceInterface `json:"source-interface,omitempty"`
	UserPeer             *string                                            `json:"user-peer,omitempty"`
	Users                *[]VpnSslSettingsAuthenticationRuleUsers           `json:"users,omitempty"`
}

type VpnSslSettingsAuthenticationRuleGroups struct {
	Name *string `json:"name,omitempty"`
}

type VpnSslSettingsAuthenticationRuleSourceAddress struct {
	Name *string `json:"name,omitempty"`
}

type VpnSslSettingsAuthenticationRuleSourceAddress6 struct {
	Name *string `json:"name,omitempty"`
}

type VpnSslSettingsAuthenticationRuleSourceInterface struct {
	Name *string `json:"name,omitempty"`
}

type VpnSslSettingsAuthenticationRuleUsers struct {
	Name *string `json:"name,omitempty"`
}

type VpnSslSettingsSourceAddress struct {
	Name *string `json:"name,omitempty"`
}

type VpnSslSettingsSourceAddress6 struct {
	Name *string `json:"name,omitempty"`
}

type VpnSslSettingsSourceInterface struct {
	Name *string `json:"name,omitempty"`
}

type VpnSslSettingsTunnelIpPools struct {
	Name *string `json:"name,omitempty"`
}

type VpnSslSettingsTunnelIpv6Pools struct {
	Name *string `json:"name,omitempty"`
}
