package models

const VpnsslSettingsPath = "vpn.ssl/settings/"

type VpnsslSettings struct {
	Algorithm                  *string                            `json:"algorithm,omitempty"`
	AuthSessionCheckSourceIp   *string                            `json:"auth-session-check-source-ip,omitempty"`
	AuthTimeout                *int64                             `json:"auth-timeout,omitempty"`
	AuthenticationRule         []VpnsslSettingsAuthenticationRule `json:"authentication-rule,omitempty"`
	AutoTunnelStaticRoute      *string                            `json:"auto-tunnel-static-route,omitempty"`
	BannedCipher               *string                            `json:"banned-cipher,omitempty"`
	CheckReferer               *string                            `json:"check-referer,omitempty"`
	Ciphersuite                *string                            `json:"ciphersuite,omitempty"`
	ClientSigalgs              *string                            `json:"client-sigalgs,omitempty"`
	DefaultPortal              *string                            `json:"default-portal,omitempty"`
	DeflateCompressionLevel    *int64                             `json:"deflate-compression-level,omitempty"`
	DeflateMinDataSize         *int64                             `json:"deflate-min-data-size,omitempty"`
	DnsServer1                 *string                            `json:"dns-server1,omitempty"`
	DnsServer2                 *string                            `json:"dns-server2,omitempty"`
	DnsSuffix                  *string                            `json:"dns-suffix,omitempty"`
	DtlsHelloTimeout           *int64                             `json:"dtls-hello-timeout,omitempty"`
	DtlsMaxProtoVer            *string                            `json:"dtls-max-proto-ver,omitempty"`
	DtlsMinProtoVer            *string                            `json:"dtls-min-proto-ver,omitempty"`
	DtlsTunnel                 *string                            `json:"dtls-tunnel,omitempty"`
	DualStackMode              *string                            `json:"dual-stack-mode,omitempty"`
	Encode2fSequence           *string                            `json:"encode-2f-sequence,omitempty"`
	EncryptAndStorePassword    *string                            `json:"encrypt-and-store-password,omitempty"`
	ForceTwoFactorAuth         *string                            `json:"force-two-factor-auth,omitempty"`
	HeaderXForwardedFor        *string                            `json:"header-x-forwarded-for,omitempty"`
	HstsIncludeSubdomains      *string                            `json:"hsts-include-subdomains,omitempty"`
	HttpCompression            *string                            `json:"http-compression,omitempty"`
	HttpOnlyCookie             *string                            `json:"http-only-cookie,omitempty"`
	HttpRequestBodyTimeout     *int64                             `json:"http-request-body-timeout,omitempty"`
	HttpRequestHeaderTimeout   *int64                             `json:"http-request-header-timeout,omitempty"`
	HttpsRedirect              *string                            `json:"https-redirect,omitempty"`
	IdleTimeout                *int64                             `json:"idle-timeout,omitempty"`
	Ipv6DnsServer1             *string                            `json:"ipv6-dns-server1,omitempty"`
	Ipv6DnsServer2             *string                            `json:"ipv6-dns-server2,omitempty"`
	Ipv6WinsServer1            *string                            `json:"ipv6-wins-server1,omitempty"`
	Ipv6WinsServer2            *string                            `json:"ipv6-wins-server2,omitempty"`
	LoginAttemptLimit          *int64                             `json:"login-attempt-limit,omitempty"`
	LoginBlockTime             *int64                             `json:"login-block-time,omitempty"`
	LoginTimeout               *int64                             `json:"login-timeout,omitempty"`
	Port                       *int64                             `json:"port,omitempty"`
	PortPrecedence             *string                            `json:"port-precedence,omitempty"`
	Reqclientcert              *string                            `json:"reqclientcert,omitempty"`
	SamlRedirectPort           *int64                             `json:"saml-redirect-port,omitempty"`
	Servercert                 *string                            `json:"servercert,omitempty"`
	SourceAddress              []VpnsslSettingsSourceAddress      `json:"source-address,omitempty"`
	SourceAddressNegate        *string                            `json:"source-address-negate,omitempty"`
	SourceAddress6             []VpnsslSettingsSourceAddress6     `json:"source-address6,omitempty"`
	SourceAddress6Negate       *string                            `json:"source-address6-negate,omitempty"`
	SourceInterface            []VpnsslSettingsSourceInterface    `json:"source-interface,omitempty"`
	SslClientRenegotiation     *string                            `json:"ssl-client-renegotiation,omitempty"`
	SslInsertEmptyFragment     *string                            `json:"ssl-insert-empty-fragment,omitempty"`
	SslMaxProtoVer             *string                            `json:"ssl-max-proto-ver,omitempty"`
	SslMinProtoVer             *string                            `json:"ssl-min-proto-ver,omitempty"`
	Status                     *string                            `json:"status,omitempty"`
	TransformBackwardSlashes   *string                            `json:"transform-backward-slashes,omitempty"`
	TunnelAddrAssignedMethod   *string                            `json:"tunnel-addr-assigned-method,omitempty"`
	TunnelConnectWithoutReauth *string                            `json:"tunnel-connect-without-reauth,omitempty"`
	TunnelIpPools              []VpnsslSettingsTunnelIpPools      `json:"tunnel-ip-pools,omitempty"`
	TunnelIpv6Pools            []VpnsslSettingsTunnelIpv6Pools    `json:"tunnel-ipv6-pools,omitempty"`
	TunnelUserSessionTimeout   *int64                             `json:"tunnel-user-session-timeout,omitempty"`
	UnsafeLegacyRenegotiation  *string                            `json:"unsafe-legacy-renegotiation,omitempty"`
	UrlObscuration             *string                            `json:"url-obscuration,omitempty"`
	UserPeer                   *string                            `json:"user-peer,omitempty"`
	WinsServer1                *string                            `json:"wins-server1,omitempty"`
	WinsServer2                *string                            `json:"wins-server2,omitempty"`
	XContentTypeOptions        *string                            `json:"x-content-type-options,omitempty"`
}

type VpnsslSettingsAuthenticationRule struct {
	Auth                 *string                                           `json:"auth,omitempty"`
	Cipher               *string                                           `json:"cipher,omitempty"`
	ClientCert           *string                                           `json:"client-cert,omitempty"`
	Groups               []VpnsslSettingsAuthenticationRuleGroups          `json:"groups,omitempty"`
	Id                   *int64                                            `json:"id,omitempty"`
	Portal               *string                                           `json:"portal,omitempty"`
	Realm                *string                                           `json:"realm,omitempty"`
	SourceAddress        []VpnsslSettingsAuthenticationRuleSourceAddress   `json:"source-address,omitempty"`
	SourceAddressNegate  *string                                           `json:"source-address-negate,omitempty"`
	SourceAddress6       []VpnsslSettingsAuthenticationRuleSourceAddress6  `json:"source-address6,omitempty"`
	SourceAddress6Negate *string                                           `json:"source-address6-negate,omitempty"`
	SourceInterface      []VpnsslSettingsAuthenticationRuleSourceInterface `json:"source-interface,omitempty"`
	UserPeer             *string                                           `json:"user-peer,omitempty"`
	Users                []VpnsslSettingsAuthenticationRuleUsers           `json:"users,omitempty"`
}

type VpnsslSettingsAuthenticationRuleGroups struct {
	Name *string `json:"name,omitempty"`
}

type VpnsslSettingsAuthenticationRuleSourceAddress struct {
	Name *string `json:"name,omitempty"`
}

type VpnsslSettingsAuthenticationRuleSourceAddress6 struct {
	Name *string `json:"name,omitempty"`
}

type VpnsslSettingsAuthenticationRuleSourceInterface struct {
	Name *string `json:"name,omitempty"`
}

type VpnsslSettingsAuthenticationRuleUsers struct {
	Name *string `json:"name,omitempty"`
}

type VpnsslSettingsSourceAddress struct {
	Name *string `json:"name,omitempty"`
}

type VpnsslSettingsSourceAddress6 struct {
	Name *string `json:"name,omitempty"`
}

type VpnsslSettingsSourceInterface struct {
	Name *string `json:"name,omitempty"`
}

type VpnsslSettingsTunnelIpPools struct {
	Name *string `json:"name,omitempty"`
}

type VpnsslSettingsTunnelIpv6Pools struct {
	Name *string `json:"name,omitempty"`
}
