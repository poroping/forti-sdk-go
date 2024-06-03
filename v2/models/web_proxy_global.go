package models

const WebProxyGlobalPath = "web-proxy/global/"

type WebProxyGlobal struct {
	FastPolicyMatch              *string                                `json:"fast-policy-match,omitempty"`
	ForwardProxyAuth             *string                                `json:"forward-proxy-auth,omitempty"`
	ForwardServerAffinityTimeout *int64                                 `json:"forward-server-affinity-timeout,omitempty"`
	LdapUserCache                *string                                `json:"ldap-user-cache,omitempty"`
	LearnClientIp                *string                                `json:"learn-client-ip,omitempty"`
	LearnClientIpFromHeader      *string                                `json:"learn-client-ip-from-header,omitempty"`
	LearnClientIpSrcaddr         *[]WebProxyGlobalLearnClientIpSrcaddr  `json:"learn-client-ip-srcaddr,omitempty"`
	LearnClientIpSrcaddr6        *[]WebProxyGlobalLearnClientIpSrcaddr6 `json:"learn-client-ip-srcaddr6,omitempty"`
	LogPolicyPending             *string                                `json:"log-policy-pending,omitempty"`
	MaxMessageLength             *int64                                 `json:"max-message-length,omitempty"`
	MaxRequestLength             *int64                                 `json:"max-request-length,omitempty"`
	MaxWafBodyCacheLength        *int64                                 `json:"max-waf-body-cache-length,omitempty"`
	PolicyCategoryDeepInspect    *string                                `json:"policy-category-deep-inspect,omitempty"`
	ProxyFqdn                    *string                                `json:"proxy-fqdn,omitempty"`
	SrcAffinityExemptAddr        *string                                `json:"src-affinity-exempt-addr,omitempty"`
	SrcAffinityExemptAddr6       *string                                `json:"src-affinity-exempt-addr6,omitempty"`
	SslCaCert                    *string                                `json:"ssl-ca-cert,omitempty"`
	SslCert                      *string                                `json:"ssl-cert,omitempty"`
	StrictWebCheck               *string                                `json:"strict-web-check,omitempty"`
	TunnelNonHttp                *string                                `json:"tunnel-non-http,omitempty"`
	UnknownHttpVersion           *string                                `json:"unknown-http-version,omitempty"`
	WebproxyProfile              *string                                `json:"webproxy-profile,omitempty"`
}

type WebProxyGlobalLearnClientIpSrcaddr struct {
	Name *string `json:"name,omitempty"`
}

type WebProxyGlobalLearnClientIpSrcaddr6 struct {
	Name *string `json:"name,omitempty"`
}
