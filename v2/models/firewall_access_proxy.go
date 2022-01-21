package models

const FirewallAccessProxyPath = "firewall/access-proxy/"

type FirewallAccessProxy struct {
	ApiGateway               *[]FirewallAccessProxyApiGateway               `json:"api-gateway,omitempty"`
	ApiGateway6              *[]FirewallAccessProxyApiGateway6              `json:"api-gateway6,omitempty"`
	ClientCert               *string                                        `json:"client-cert,omitempty"`
	DecryptedTrafficMirror   *string                                        `json:"decrypted-traffic-mirror,omitempty"`
	EmptyCertAction          *string                                        `json:"empty-cert-action,omitempty"`
	LdbMethod                *string                                        `json:"ldb-method,omitempty"`
	LogBlockedTraffic        *string                                        `json:"log-blocked-traffic,omitempty"`
	Name                     *string                                        `json:"name,omitempty"`
	Realservers              *[]FirewallAccessProxyRealservers              `json:"realservers,omitempty"`
	ServerPubkeyAuth         *string                                        `json:"server-pubkey-auth,omitempty"`
	ServerPubkeyAuthSettings *[]FirewallAccessProxyServerPubkeyAuthSettings `json:"server-pubkey-auth-settings,omitempty"`
	Vip                      *string                                        `json:"vip,omitempty"`
}

type FirewallAccessProxyApiGateway struct {
	HttpCookieAge            *int64                                          `json:"http-cookie-age,omitempty"`
	HttpCookieDomain         *string                                         `json:"http-cookie-domain,omitempty"`
	HttpCookieDomainFromHost *string                                         `json:"http-cookie-domain-from-host,omitempty"`
	HttpCookieGeneration     *int64                                          `json:"http-cookie-generation,omitempty"`
	HttpCookiePath           *string                                         `json:"http-cookie-path,omitempty"`
	HttpCookieShare          *string                                         `json:"http-cookie-share,omitempty"`
	HttpsCookieSecure        *string                                         `json:"https-cookie-secure,omitempty"`
	Id                       *int64                                          `json:"id,omitempty"`
	LdbMethod                *string                                         `json:"ldb-method,omitempty"`
	Persistence              *string                                         `json:"persistence,omitempty"`
	Realservers              *[]FirewallAccessProxyApiGatewayRealservers     `json:"realservers,omitempty"`
	SamlRedirect             *string                                         `json:"saml-redirect,omitempty"`
	SamlServer               *string                                         `json:"saml-server,omitempty"`
	Service                  *string                                         `json:"service,omitempty"`
	SslAlgorithm             *string                                         `json:"ssl-algorithm,omitempty"`
	SslCipherSuites          *[]FirewallAccessProxyApiGatewaySslCipherSuites `json:"ssl-cipher-suites,omitempty"`
	SslDhBits                *string                                         `json:"ssl-dh-bits,omitempty"`
	SslMaxVersion            *string                                         `json:"ssl-max-version,omitempty"`
	SslMinVersion            *string                                         `json:"ssl-min-version,omitempty"`
	UrlMap                   *string                                         `json:"url-map,omitempty"`
	UrlMapType               *string                                         `json:"url-map-type,omitempty"`
	VirtualHost              *string                                         `json:"virtual-host,omitempty"`
}

type FirewallAccessProxyApiGatewayRealservers struct {
	AddrType             *string                                               `json:"addr-type,omitempty"`
	Address              *string                                               `json:"address,omitempty"`
	HealthCheck          *string                                               `json:"health-check,omitempty"`
	HealthCheckProto     *string                                               `json:"health-check-proto,omitempty"`
	HolddownInterval     *string                                               `json:"holddown-interval,omitempty"`
	HttpHost             *string                                               `json:"http-host,omitempty"`
	Id                   *int64                                                `json:"id,omitempty"`
	Ip                   *string                                               `json:"ip,omitempty"`
	Mappedport           *string                                               `json:"mappedport,omitempty"`
	Port                 *int64                                                `json:"port,omitempty"`
	SshClientCert        *string                                               `json:"ssh-client-cert,omitempty"`
	SshHostKey           *[]FirewallAccessProxyApiGatewayRealserversSshHostKey `json:"ssh-host-key,omitempty"`
	SshHostKeyValidation *string                                               `json:"ssh-host-key-validation,omitempty"`
	Status               *string                                               `json:"status,omitempty"`
	Type                 *string                                               `json:"type,omitempty"`
	Weight               *int64                                                `json:"weight,omitempty"`
}

type FirewallAccessProxyApiGatewayRealserversSshHostKey struct {
	Name *string `json:"name,omitempty"`
}

type FirewallAccessProxyApiGatewaySslCipherSuites struct {
	Cipher   *string `json:"cipher,omitempty"`
	Priority *int64  `json:"priority,omitempty"`
	Versions *string `json:"versions,omitempty"`
}

type FirewallAccessProxyApiGateway6 struct {
	HttpCookieAge            *int64                                           `json:"http-cookie-age,omitempty"`
	HttpCookieDomain         *string                                          `json:"http-cookie-domain,omitempty"`
	HttpCookieDomainFromHost *string                                          `json:"http-cookie-domain-from-host,omitempty"`
	HttpCookieGeneration     *int64                                           `json:"http-cookie-generation,omitempty"`
	HttpCookiePath           *string                                          `json:"http-cookie-path,omitempty"`
	HttpCookieShare          *string                                          `json:"http-cookie-share,omitempty"`
	HttpsCookieSecure        *string                                          `json:"https-cookie-secure,omitempty"`
	Id                       *int64                                           `json:"id,omitempty"`
	LdbMethod                *string                                          `json:"ldb-method,omitempty"`
	Persistence              *string                                          `json:"persistence,omitempty"`
	Realservers              *[]FirewallAccessProxyApiGateway6Realservers     `json:"realservers,omitempty"`
	SamlRedirect             *string                                          `json:"saml-redirect,omitempty"`
	SamlServer               *string                                          `json:"saml-server,omitempty"`
	Service                  *string                                          `json:"service,omitempty"`
	SslAlgorithm             *string                                          `json:"ssl-algorithm,omitempty"`
	SslCipherSuites          *[]FirewallAccessProxyApiGateway6SslCipherSuites `json:"ssl-cipher-suites,omitempty"`
	SslDhBits                *string                                          `json:"ssl-dh-bits,omitempty"`
	SslMaxVersion            *string                                          `json:"ssl-max-version,omitempty"`
	SslMinVersion            *string                                          `json:"ssl-min-version,omitempty"`
	UrlMap                   *string                                          `json:"url-map,omitempty"`
	UrlMapType               *string                                          `json:"url-map-type,omitempty"`
	VirtualHost              *string                                          `json:"virtual-host,omitempty"`
}

type FirewallAccessProxyApiGateway6Realservers struct {
	AddrType             *string                                                `json:"addr-type,omitempty"`
	Address              *string                                                `json:"address,omitempty"`
	HealthCheck          *string                                                `json:"health-check,omitempty"`
	HealthCheckProto     *string                                                `json:"health-check-proto,omitempty"`
	HolddownInterval     *string                                                `json:"holddown-interval,omitempty"`
	HttpHost             *string                                                `json:"http-host,omitempty"`
	Id                   *int64                                                 `json:"id,omitempty"`
	Ip                   *string                                                `json:"ip,omitempty"`
	Mappedport           *string                                                `json:"mappedport,omitempty"`
	Port                 *int64                                                 `json:"port,omitempty"`
	SshClientCert        *string                                                `json:"ssh-client-cert,omitempty"`
	SshHostKey           *[]FirewallAccessProxyApiGateway6RealserversSshHostKey `json:"ssh-host-key,omitempty"`
	SshHostKeyValidation *string                                                `json:"ssh-host-key-validation,omitempty"`
	Status               *string                                                `json:"status,omitempty"`
	Type                 *string                                                `json:"type,omitempty"`
	Weight               *int64                                                 `json:"weight,omitempty"`
}

type FirewallAccessProxyApiGateway6RealserversSshHostKey struct {
	Name *string `json:"name,omitempty"`
}

type FirewallAccessProxyApiGateway6SslCipherSuites struct {
	Cipher   *string `json:"cipher,omitempty"`
	Priority *int64  `json:"priority,omitempty"`
	Versions *string `json:"versions,omitempty"`
}

type FirewallAccessProxyRealservers struct {
	Id     *int64  `json:"id,omitempty"`
	Ip     *string `json:"ip,omitempty"`
	Port   *int64  `json:"port,omitempty"`
	Status *string `json:"status,omitempty"`
	Weight *int64  `json:"weight,omitempty"`
}

type FirewallAccessProxyServerPubkeyAuthSettings struct {
	AuthCa                *string                                                     `json:"auth-ca,omitempty"`
	CertExtension         *[]FirewallAccessProxyServerPubkeyAuthSettingsCertExtension `json:"cert-extension,omitempty"`
	PermitAgentForwarding *string                                                     `json:"permit-agent-forwarding,omitempty"`
	PermitPortForwarding  *string                                                     `json:"permit-port-forwarding,omitempty"`
	PermitPty             *string                                                     `json:"permit-pty,omitempty"`
	PermitUserRc          *string                                                     `json:"permit-user-rc,omitempty"`
	PermitX11Forwarding   *string                                                     `json:"permit-x11-forwarding,omitempty"`
	SourceAddress         *string                                                     `json:"source-address,omitempty"`
}

type FirewallAccessProxyServerPubkeyAuthSettingsCertExtension struct {
	Critical *string `json:"critical,omitempty"`
	Data     *string `json:"data,omitempty"`
	Name     *string `json:"name,omitempty"`
	Type     *string `json:"type,omitempty"`
}
