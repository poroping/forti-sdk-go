package models

const FirewallAccessProxy6Path = "firewall/access-proxy6/"

type FirewallAccessProxy6 struct {
	AddVhostDomainToDnsdb   *string                            `json:"add-vhost-domain-to-dnsdb,omitempty"`
	ApiGateway              *[]FirewallAccessProxy6ApiGateway  `json:"api-gateway,omitempty"`
	ApiGateway6             *[]FirewallAccessProxy6ApiGateway6 `json:"api-gateway6,omitempty"`
	AuthPortal              *string                            `json:"auth-portal,omitempty"`
	AuthVirtualHost         *string                            `json:"auth-virtual-host,omitempty"`
	ClientCert              *string                            `json:"client-cert,omitempty"`
	DecryptedTrafficMirror  *string                            `json:"decrypted-traffic-mirror,omitempty"`
	EmptyCertAction         *string                            `json:"empty-cert-action,omitempty"`
	HttpSupportedMaxVersion *string                            `json:"http-supported-max-version,omitempty"`
	LogBlockedTraffic       *string                            `json:"log-blocked-traffic,omitempty"`
	Name                    *string                            `json:"name,omitempty"`
	SvrPoolMultiplex        *string                            `json:"svr-pool-multiplex,omitempty"`
	SvrPoolServerMaxRequest *int64                             `json:"svr-pool-server-max-request,omitempty"`
	SvrPoolTtl              *int64                             `json:"svr-pool-ttl,omitempty"`
	UserAgentDetect         *string                            `json:"user-agent-detect,omitempty"`
	Vip                     *string                            `json:"vip,omitempty"`
}

type FirewallAccessProxy6ApiGateway struct {
	Application              *[]FirewallAccessProxy6ApiGatewayApplication     `json:"application,omitempty"`
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
	Realservers              *[]FirewallAccessProxy6ApiGatewayRealservers     `json:"realservers,omitempty"`
	SamlRedirect             *string                                          `json:"saml-redirect,omitempty"`
	SamlServer               *string                                          `json:"saml-server,omitempty"`
	Service                  *string                                          `json:"service,omitempty"`
	SslAlgorithm             *string                                          `json:"ssl-algorithm,omitempty"`
	SslCipherSuites          *[]FirewallAccessProxy6ApiGatewaySslCipherSuites `json:"ssl-cipher-suites,omitempty"`
	SslDhBits                *string                                          `json:"ssl-dh-bits,omitempty"`
	SslMaxVersion            *string                                          `json:"ssl-max-version,omitempty"`
	SslMinVersion            *string                                          `json:"ssl-min-version,omitempty"`
	SslRenegotiation         *string                                          `json:"ssl-renegotiation,omitempty"`
	SslVpnWebPortal          *string                                          `json:"ssl-vpn-web-portal,omitempty"`
	UrlMap                   *string                                          `json:"url-map,omitempty"`
	UrlMapType               *string                                          `json:"url-map-type,omitempty"`
	VirtualHost              *string                                          `json:"virtual-host,omitempty"`
}

type FirewallAccessProxy6ApiGatewayApplication struct {
	Name *string `json:"name,omitempty"`
}

type FirewallAccessProxy6ApiGatewayRealservers struct {
	AddrType             *string                                                `json:"addr-type,omitempty"`
	Address              *string                                                `json:"address,omitempty"`
	Domain               *string                                                `json:"domain,omitempty"`
	HealthCheck          *string                                                `json:"health-check,omitempty"`
	HealthCheckProto     *string                                                `json:"health-check-proto,omitempty"`
	HolddownInterval     *string                                                `json:"holddown-interval,omitempty"`
	HttpHost             *string                                                `json:"http-host,omitempty"`
	Id                   *int64                                                 `json:"id,omitempty"`
	Ip                   *string                                                `json:"ip,omitempty"`
	Mappedport           *string                                                `json:"mappedport,omitempty"`
	Port                 *int64                                                 `json:"port,omitempty"`
	SshClientCert        *string                                                `json:"ssh-client-cert,omitempty"`
	SshHostKey           *[]FirewallAccessProxy6ApiGatewayRealserversSshHostKey `json:"ssh-host-key,omitempty"`
	SshHostKeyValidation *string                                                `json:"ssh-host-key-validation,omitempty"`
	Status               *string                                                `json:"status,omitempty"`
	TranslateHost        *string                                                `json:"translate-host,omitempty"`
	Type                 *string                                                `json:"type,omitempty"`
	Weight               *int64                                                 `json:"weight,omitempty"`
}

type FirewallAccessProxy6ApiGatewayRealserversSshHostKey struct {
	Name *string `json:"name,omitempty"`
}

type FirewallAccessProxy6ApiGatewaySslCipherSuites struct {
	Cipher   *string `json:"cipher,omitempty"`
	Priority *int64  `json:"priority,omitempty"`
	Versions *string `json:"versions,omitempty"`
}

type FirewallAccessProxy6ApiGateway6 struct {
	Application              *[]FirewallAccessProxy6ApiGateway6Application     `json:"application,omitempty"`
	HttpCookieAge            *int64                                            `json:"http-cookie-age,omitempty"`
	HttpCookieDomain         *string                                           `json:"http-cookie-domain,omitempty"`
	HttpCookieDomainFromHost *string                                           `json:"http-cookie-domain-from-host,omitempty"`
	HttpCookieGeneration     *int64                                            `json:"http-cookie-generation,omitempty"`
	HttpCookiePath           *string                                           `json:"http-cookie-path,omitempty"`
	HttpCookieShare          *string                                           `json:"http-cookie-share,omitempty"`
	HttpsCookieSecure        *string                                           `json:"https-cookie-secure,omitempty"`
	Id                       *int64                                            `json:"id,omitempty"`
	LdbMethod                *string                                           `json:"ldb-method,omitempty"`
	Persistence              *string                                           `json:"persistence,omitempty"`
	Realservers              *[]FirewallAccessProxy6ApiGateway6Realservers     `json:"realservers,omitempty"`
	SamlRedirect             *string                                           `json:"saml-redirect,omitempty"`
	SamlServer               *string                                           `json:"saml-server,omitempty"`
	Service                  *string                                           `json:"service,omitempty"`
	SslAlgorithm             *string                                           `json:"ssl-algorithm,omitempty"`
	SslCipherSuites          *[]FirewallAccessProxy6ApiGateway6SslCipherSuites `json:"ssl-cipher-suites,omitempty"`
	SslDhBits                *string                                           `json:"ssl-dh-bits,omitempty"`
	SslMaxVersion            *string                                           `json:"ssl-max-version,omitempty"`
	SslMinVersion            *string                                           `json:"ssl-min-version,omitempty"`
	SslRenegotiation         *string                                           `json:"ssl-renegotiation,omitempty"`
	SslVpnWebPortal          *string                                           `json:"ssl-vpn-web-portal,omitempty"`
	UrlMap                   *string                                           `json:"url-map,omitempty"`
	UrlMapType               *string                                           `json:"url-map-type,omitempty"`
	VirtualHost              *string                                           `json:"virtual-host,omitempty"`
}

type FirewallAccessProxy6ApiGateway6Application struct {
	Name *string `json:"name,omitempty"`
}

type FirewallAccessProxy6ApiGateway6Realservers struct {
	AddrType             *string                                                 `json:"addr-type,omitempty"`
	Address              *string                                                 `json:"address,omitempty"`
	Domain               *string                                                 `json:"domain,omitempty"`
	HealthCheck          *string                                                 `json:"health-check,omitempty"`
	HealthCheckProto     *string                                                 `json:"health-check-proto,omitempty"`
	HolddownInterval     *string                                                 `json:"holddown-interval,omitempty"`
	HttpHost             *string                                                 `json:"http-host,omitempty"`
	Id                   *int64                                                  `json:"id,omitempty"`
	Ip                   *string                                                 `json:"ip,omitempty"`
	Mappedport           *string                                                 `json:"mappedport,omitempty"`
	Port                 *int64                                                  `json:"port,omitempty"`
	SshClientCert        *string                                                 `json:"ssh-client-cert,omitempty"`
	SshHostKey           *[]FirewallAccessProxy6ApiGateway6RealserversSshHostKey `json:"ssh-host-key,omitempty"`
	SshHostKeyValidation *string                                                 `json:"ssh-host-key-validation,omitempty"`
	Status               *string                                                 `json:"status,omitempty"`
	TranslateHost        *string                                                 `json:"translate-host,omitempty"`
	Type                 *string                                                 `json:"type,omitempty"`
	Weight               *int64                                                  `json:"weight,omitempty"`
}

type FirewallAccessProxy6ApiGateway6RealserversSshHostKey struct {
	Name *string `json:"name,omitempty"`
}

type FirewallAccessProxy6ApiGateway6SslCipherSuites struct {
	Cipher   *string `json:"cipher,omitempty"`
	Priority *int64  `json:"priority,omitempty"`
	Versions *string `json:"versions,omitempty"`
}
