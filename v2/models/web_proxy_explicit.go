package models

const WebProxyExplicitPath = "web-proxy/explicit/"

type WebProxyExplicit struct {
	FtpIncomingPort         *string                      `json:"ftp-incoming-port,omitempty"`
	FtpOverHttp             *string                      `json:"ftp-over-http,omitempty"`
	HttpConnectionMode      *string                      `json:"http-connection-mode,omitempty"`
	HttpIncomingPort        *string                      `json:"http-incoming-port,omitempty"`
	HttpsIncomingPort       *string                      `json:"https-incoming-port,omitempty"`
	HttpsReplacementMessage *string                      `json:"https-replacement-message,omitempty"`
	IncomingIp              *string                      `json:"incoming-ip,omitempty"`
	IncomingIp6             *string                      `json:"incoming-ip6,omitempty"`
	Ipv6Status              *string                      `json:"ipv6-status,omitempty"`
	MessageUponServerError  *string                      `json:"message-upon-server-error,omitempty"`
	OutgoingIp              *string                      `json:"outgoing-ip,omitempty"`
	OutgoingIp6             *string                      `json:"outgoing-ip6,omitempty"`
	PacFileData             *string                      `json:"pac-file-data,omitempty"`
	PacFileName             *string                      `json:"pac-file-name,omitempty"`
	PacFileServerPort       *string                      `json:"pac-file-server-port,omitempty"`
	PacFileServerStatus     *string                      `json:"pac-file-server-status,omitempty"`
	PacFileThroughHttps     *string                      `json:"pac-file-through-https,omitempty"`
	PacFileUrl              *string                      `json:"pac-file-url,omitempty"`
	PacPolicy               *[]WebProxyExplicitPacPolicy `json:"pac-policy,omitempty"`
	PrefDnsResult           *string                      `json:"pref-dns-result,omitempty"`
	Realm                   *string                      `json:"realm,omitempty"`
	SecDefaultAction        *string                      `json:"sec-default-action,omitempty"`
	Socks                   *string                      `json:"socks,omitempty"`
	SocksIncomingPort       *string                      `json:"socks-incoming-port,omitempty"`
	SslAlgorithm            *string                      `json:"ssl-algorithm,omitempty"`
	Status                  *string                      `json:"status,omitempty"`
	StrictGuest             *string                      `json:"strict-guest,omitempty"`
	TraceAuthNoRsp          *string                      `json:"trace-auth-no-rsp,omitempty"`
	UnknownHttpVersion      *string                      `json:"unknown-http-version,omitempty"`
}

type WebProxyExplicitPacPolicy struct {
	Comments    *string                              `json:"comments,omitempty"`
	Dstaddr     *[]WebProxyExplicitPacPolicyDstaddr  `json:"dstaddr,omitempty"`
	PacFileData *string                              `json:"pac-file-data,omitempty"`
	PacFileName *string                              `json:"pac-file-name,omitempty"`
	Policyid    *int64                               `json:"policyid,omitempty"`
	Srcaddr     *[]WebProxyExplicitPacPolicySrcaddr  `json:"srcaddr,omitempty"`
	Srcaddr6    *[]WebProxyExplicitPacPolicySrcaddr6 `json:"srcaddr6,omitempty"`
	Status      *string                              `json:"status,omitempty"`
}

type WebProxyExplicitPacPolicyDstaddr struct {
	Name *string `json:"name,omitempty"`
}

type WebProxyExplicitPacPolicySrcaddr struct {
	Name *string `json:"name,omitempty"`
}

type WebProxyExplicitPacPolicySrcaddr6 struct {
	Name *string `json:"name,omitempty"`
}
