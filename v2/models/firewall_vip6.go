package models

const FirewallVip6Path = "firewall/vip6/"

type FirewallVip6 struct {
	AddNat64Route                *string                              `json:"add-nat64-route,omitempty"`
	ArpReply                     *string                              `json:"arp-reply,omitempty"`
	Color                        *int64                               `json:"color,omitempty"`
	Comment                      *string                              `json:"comment,omitempty"`
	EmbeddedIpv4Address          *string                              `json:"embedded-ipv4-address,omitempty"`
	Extip                        *string                              `json:"extip,omitempty"`
	Extport                      *string                              `json:"extport,omitempty"`
	HttpCookieAge                *int64                               `json:"http-cookie-age,omitempty"`
	HttpCookieDomain             *string                              `json:"http-cookie-domain,omitempty"`
	HttpCookieDomainFromHost     *string                              `json:"http-cookie-domain-from-host,omitempty"`
	HttpCookieGeneration         *int64                               `json:"http-cookie-generation,omitempty"`
	HttpCookiePath               *string                              `json:"http-cookie-path,omitempty"`
	HttpCookieShare              *string                              `json:"http-cookie-share,omitempty"`
	HttpIpHeader                 *string                              `json:"http-ip-header,omitempty"`
	HttpIpHeaderName             *string                              `json:"http-ip-header-name,omitempty"`
	HttpMultiplex                *string                              `json:"http-multiplex,omitempty"`
	HttpRedirect                 *string                              `json:"http-redirect,omitempty"`
	HttpsCookieSecure            *string                              `json:"https-cookie-secure,omitempty"`
	Id                           *int64                               `json:"id,omitempty"`
	Ipv4Mappedip                 *string                              `json:"ipv4-mappedip,omitempty"`
	Ipv4Mappedport               *string                              `json:"ipv4-mappedport,omitempty"`
	LdbMethod                    *string                              `json:"ldb-method,omitempty"`
	Mappedip                     *string                              `json:"mappedip,omitempty"`
	Mappedport                   *string                              `json:"mappedport,omitempty"`
	MaxEmbryonicConnections      *int64                               `json:"max-embryonic-connections,omitempty"`
	Monitor                      *[]FirewallVip6Monitor               `json:"monitor,omitempty"`
	Name                         *string                              `json:"name,omitempty"`
	NatSourceVip                 *string                              `json:"nat-source-vip,omitempty"`
	Nat64                        *string                              `json:"nat64,omitempty"`
	Nat66                        *string                              `json:"nat66,omitempty"`
	OutlookWebAccess             *string                              `json:"outlook-web-access,omitempty"`
	Persistence                  *string                              `json:"persistence,omitempty"`
	Portforward                  *string                              `json:"portforward,omitempty"`
	Protocol                     *string                              `json:"protocol,omitempty"`
	Realservers                  *[]FirewallVip6Realservers           `json:"realservers,omitempty"`
	ServerType                   *string                              `json:"server-type,omitempty"`
	SrcFilter                    *[]FirewallVip6SrcFilter             `json:"src-filter,omitempty"`
	SslAcceptFfdheGroups         *string                              `json:"ssl-accept-ffdhe-groups,omitempty"`
	SslAlgorithm                 *string                              `json:"ssl-algorithm,omitempty"`
	SslCertificate               *string                              `json:"ssl-certificate,omitempty"`
	SslCipherSuites              *[]FirewallVip6SslCipherSuites       `json:"ssl-cipher-suites,omitempty"`
	SslClientFallback            *string                              `json:"ssl-client-fallback,omitempty"`
	SslClientRekeyCount          *int64                               `json:"ssl-client-rekey-count,omitempty"`
	SslClientRenegotiation       *string                              `json:"ssl-client-renegotiation,omitempty"`
	SslClientSessionStateMax     *int64                               `json:"ssl-client-session-state-max,omitempty"`
	SslClientSessionStateTimeout *int64                               `json:"ssl-client-session-state-timeout,omitempty"`
	SslClientSessionStateType    *string                              `json:"ssl-client-session-state-type,omitempty"`
	SslDhBits                    *string                              `json:"ssl-dh-bits,omitempty"`
	SslHpkp                      *string                              `json:"ssl-hpkp,omitempty"`
	SslHpkpAge                   *int64                               `json:"ssl-hpkp-age,omitempty"`
	SslHpkpBackup                *string                              `json:"ssl-hpkp-backup,omitempty"`
	SslHpkpIncludeSubdomains     *string                              `json:"ssl-hpkp-include-subdomains,omitempty"`
	SslHpkpPrimary               *string                              `json:"ssl-hpkp-primary,omitempty"`
	SslHpkpReportUri             *string                              `json:"ssl-hpkp-report-uri,omitempty"`
	SslHsts                      *string                              `json:"ssl-hsts,omitempty"`
	SslHstsAge                   *int64                               `json:"ssl-hsts-age,omitempty"`
	SslHstsIncludeSubdomains     *string                              `json:"ssl-hsts-include-subdomains,omitempty"`
	SslHttpLocationConversion    *string                              `json:"ssl-http-location-conversion,omitempty"`
	SslHttpMatchHost             *string                              `json:"ssl-http-match-host,omitempty"`
	SslMaxVersion                *string                              `json:"ssl-max-version,omitempty"`
	SslMinVersion                *string                              `json:"ssl-min-version,omitempty"`
	SslMode                      *string                              `json:"ssl-mode,omitempty"`
	SslPfs                       *string                              `json:"ssl-pfs,omitempty"`
	SslSendEmptyFrags            *string                              `json:"ssl-send-empty-frags,omitempty"`
	SslServerAlgorithm           *string                              `json:"ssl-server-algorithm,omitempty"`
	SslServerCipherSuites        *[]FirewallVip6SslServerCipherSuites `json:"ssl-server-cipher-suites,omitempty"`
	SslServerMaxVersion          *string                              `json:"ssl-server-max-version,omitempty"`
	SslServerMinVersion          *string                              `json:"ssl-server-min-version,omitempty"`
	SslServerSessionStateMax     *int64                               `json:"ssl-server-session-state-max,omitempty"`
	SslServerSessionStateTimeout *int64                               `json:"ssl-server-session-state-timeout,omitempty"`
	SslServerSessionStateType    *string                              `json:"ssl-server-session-state-type,omitempty"`
	Type                         *string                              `json:"type,omitempty"`
	Uuid                         *string                              `json:"uuid,omitempty"`
	WeblogicServer               *string                              `json:"weblogic-server,omitempty"`
	WebsphereServer              *string                              `json:"websphere-server,omitempty"`
}

type FirewallVip6Monitor struct {
	Name *string `json:"name,omitempty"`
}

type FirewallVip6Realservers struct {
	ClientIp         *string                           `json:"client-ip,omitempty"`
	Healthcheck      *string                           `json:"healthcheck,omitempty"`
	HolddownInterval *int64                            `json:"holddown-interval,omitempty"`
	HttpHost         *string                           `json:"http-host,omitempty"`
	Id               *int64                            `json:"id,omitempty"`
	Ip               *string                           `json:"ip,omitempty"`
	MaxConnections   *int64                            `json:"max-connections,omitempty"`
	Monitor          *[]FirewallVip6RealserversMonitor `json:"monitor,omitempty"`
	Port             *int64                            `json:"port,omitempty"`
	Status           *string                           `json:"status,omitempty"`
	Weight           *int64                            `json:"weight,omitempty"`
}

type FirewallVip6RealserversMonitor struct {
	Name *string `json:"name,omitempty"`
}

type FirewallVip6SrcFilter struct {
	Range *string `json:"range,omitempty"`
}

type FirewallVip6SslCipherSuites struct {
	Cipher   *string `json:"cipher,omitempty"`
	Priority *int64  `json:"priority,omitempty"`
	Versions *string `json:"versions,omitempty"`
}

type FirewallVip6SslServerCipherSuites struct {
	Cipher   *string `json:"cipher,omitempty"`
	Priority *int64  `json:"priority,omitempty"`
	Versions *string `json:"versions,omitempty"`
}
