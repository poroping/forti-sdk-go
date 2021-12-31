package models

const FirewallVipPath = "firewall/vip/"

type FirewallVip struct {
	AddNat46Route                *string                             `json:"add-nat46-route,omitempty"`
	ArpReply                     *string                             `json:"arp-reply,omitempty"`
	Color                        *int64                              `json:"color,omitempty"`
	Comment                      *string                             `json:"comment,omitempty"`
	DnsMappingTtl                *int64                              `json:"dns-mapping-ttl,omitempty"`
	Extaddr                      *[]FirewallVipExtaddr               `json:"extaddr,omitempty"`
	Extintf                      *string                             `json:"extintf,omitempty"`
	Extip                        *string                             `json:"extip,omitempty"`
	Extport                      *string                             `json:"extport,omitempty"`
	GratuitousArpInterval        *int64                              `json:"gratuitous-arp-interval,omitempty"`
	HttpCookieAge                *int64                              `json:"http-cookie-age,omitempty"`
	HttpCookieDomain             *string                             `json:"http-cookie-domain,omitempty"`
	HttpCookieDomainFromHost     *string                             `json:"http-cookie-domain-from-host,omitempty"`
	HttpCookieGeneration         *int64                              `json:"http-cookie-generation,omitempty"`
	HttpCookiePath               *string                             `json:"http-cookie-path,omitempty"`
	HttpCookieShare              *string                             `json:"http-cookie-share,omitempty"`
	HttpIpHeader                 *string                             `json:"http-ip-header,omitempty"`
	HttpIpHeaderName             *string                             `json:"http-ip-header-name,omitempty"`
	HttpMultiplex                *string                             `json:"http-multiplex,omitempty"`
	HttpRedirect                 *string                             `json:"http-redirect,omitempty"`
	HttpsCookieSecure            *string                             `json:"https-cookie-secure,omitempty"`
	Id                           *int64                              `json:"id,omitempty"`
	Ipv6Mappedip                 *string                             `json:"ipv6-mappedip,omitempty"`
	Ipv6Mappedport               *string                             `json:"ipv6-mappedport,omitempty"`
	LdbMethod                    *string                             `json:"ldb-method,omitempty"`
	MappedAddr                   *string                             `json:"mapped-addr,omitempty"`
	Mappedip                     *[]FirewallVipMappedip              `json:"mappedip,omitempty"`
	Mappedport                   *string                             `json:"mappedport,omitempty"`
	MaxEmbryonicConnections      *int64                              `json:"max-embryonic-connections,omitempty"`
	Monitor                      *[]FirewallVipMonitor               `json:"monitor,omitempty"`
	Name                         *string                             `json:"name,omitempty"`
	NatSourceVip                 *string                             `json:"nat-source-vip,omitempty"`
	Nat44                        *string                             `json:"nat44,omitempty"`
	Nat46                        *string                             `json:"nat46,omitempty"`
	OutlookWebAccess             *string                             `json:"outlook-web-access,omitempty"`
	Persistence                  *string                             `json:"persistence,omitempty"`
	Portforward                  *string                             `json:"portforward,omitempty"`
	PortmappingType              *string                             `json:"portmapping-type,omitempty"`
	Protocol                     *string                             `json:"protocol,omitempty"`
	Realservers                  *[]FirewallVipRealservers           `json:"realservers,omitempty"`
	ServerType                   *string                             `json:"server-type,omitempty"`
	Service                      *[]FirewallVipService               `json:"service,omitempty"`
	SrcFilter                    *[]FirewallVipSrcFilter             `json:"src-filter,omitempty"`
	SrcintfFilter                *[]FirewallVipSrcintfFilter         `json:"srcintf-filter,omitempty"`
	SslAlgorithm                 *string                             `json:"ssl-algorithm,omitempty"`
	SslCertificate               *string                             `json:"ssl-certificate,omitempty"`
	SslCipherSuites              *[]FirewallVipSslCipherSuites       `json:"ssl-cipher-suites,omitempty"`
	SslClientFallback            *string                             `json:"ssl-client-fallback,omitempty"`
	SslClientRekeyCount          *int64                              `json:"ssl-client-rekey-count,omitempty"`
	SslClientRenegotiation       *string                             `json:"ssl-client-renegotiation,omitempty"`
	SslClientSessionStateMax     *int64                              `json:"ssl-client-session-state-max,omitempty"`
	SslClientSessionStateTimeout *int64                              `json:"ssl-client-session-state-timeout,omitempty"`
	SslClientSessionStateType    *string                             `json:"ssl-client-session-state-type,omitempty"`
	SslDhBits                    *string                             `json:"ssl-dh-bits,omitempty"`
	SslHpkp                      *string                             `json:"ssl-hpkp,omitempty"`
	SslHpkpAge                   *int64                              `json:"ssl-hpkp-age,omitempty"`
	SslHpkpBackup                *string                             `json:"ssl-hpkp-backup,omitempty"`
	SslHpkpIncludeSubdomains     *string                             `json:"ssl-hpkp-include-subdomains,omitempty"`
	SslHpkpPrimary               *string                             `json:"ssl-hpkp-primary,omitempty"`
	SslHpkpReportUri             *string                             `json:"ssl-hpkp-report-uri,omitempty"`
	SslHsts                      *string                             `json:"ssl-hsts,omitempty"`
	SslHstsAge                   *int64                              `json:"ssl-hsts-age,omitempty"`
	SslHstsIncludeSubdomains     *string                             `json:"ssl-hsts-include-subdomains,omitempty"`
	SslHttpLocationConversion    *string                             `json:"ssl-http-location-conversion,omitempty"`
	SslHttpMatchHost             *string                             `json:"ssl-http-match-host,omitempty"`
	SslMaxVersion                *string                             `json:"ssl-max-version,omitempty"`
	SslMinVersion                *string                             `json:"ssl-min-version,omitempty"`
	SslMode                      *string                             `json:"ssl-mode,omitempty"`
	SslPfs                       *string                             `json:"ssl-pfs,omitempty"`
	SslSendEmptyFrags            *string                             `json:"ssl-send-empty-frags,omitempty"`
	SslServerAlgorithm           *string                             `json:"ssl-server-algorithm,omitempty"`
	SslServerCipherSuites        *[]FirewallVipSslServerCipherSuites `json:"ssl-server-cipher-suites,omitempty"`
	SslServerMaxVersion          *string                             `json:"ssl-server-max-version,omitempty"`
	SslServerMinVersion          *string                             `json:"ssl-server-min-version,omitempty"`
	SslServerSessionStateMax     *int64                              `json:"ssl-server-session-state-max,omitempty"`
	SslServerSessionStateTimeout *int64                              `json:"ssl-server-session-state-timeout,omitempty"`
	SslServerSessionStateType    *string                             `json:"ssl-server-session-state-type,omitempty"`
	Status                       *string                             `json:"status,omitempty"`
	Type                         *string                             `json:"type,omitempty"`
	Uuid                         *string                             `json:"uuid,omitempty"`
	WeblogicServer               *string                             `json:"weblogic-server,omitempty"`
	WebsphereServer              *string                             `json:"websphere-server,omitempty"`
}

type FirewallVipExtaddr struct {
	Name *string `json:"name,omitempty"`
}

type FirewallVipMappedip struct {
	Range *string `json:"range,omitempty"`
}

type FirewallVipMonitor struct {
	Name *string `json:"name,omitempty"`
}

type FirewallVipRealservers struct {
	Address          *string                          `json:"address,omitempty"`
	ClientIp         *string                          `json:"client-ip,omitempty"`
	Healthcheck      *string                          `json:"healthcheck,omitempty"`
	HolddownInterval *int64                           `json:"holddown-interval,omitempty"`
	HttpHost         *string                          `json:"http-host,omitempty"`
	Id               *int64                           `json:"id,omitempty"`
	Ip               *string                          `json:"ip,omitempty"`
	MaxConnections   *int64                           `json:"max-connections,omitempty"`
	Monitor          *[]FirewallVipRealserversMonitor `json:"monitor,omitempty"`
	Port             *int64                           `json:"port,omitempty"`
	Status           *string                          `json:"status,omitempty"`
	Type             *string                          `json:"type,omitempty"`
	Weight           *int64                           `json:"weight,omitempty"`
}

type FirewallVipRealserversMonitor struct {
	Name *string `json:"name,omitempty"`
}

type FirewallVipService struct {
	Name *string `json:"name,omitempty"`
}

type FirewallVipSrcFilter struct {
	Range *string `json:"range,omitempty"`
}

type FirewallVipSrcintfFilter struct {
	InterfaceName *string `json:"interface-name,omitempty"`
}

type FirewallVipSslCipherSuites struct {
	Cipher   *string `json:"cipher,omitempty"`
	Priority *int64  `json:"priority,omitempty"`
	Versions *string `json:"versions,omitempty"`
}

type FirewallVipSslServerCipherSuites struct {
	Cipher   *string `json:"cipher,omitempty"`
	Priority *int64  `json:"priority,omitempty"`
	Versions *string `json:"versions,omitempty"`
}
