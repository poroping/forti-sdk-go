package models

const WebProxyProfilePath = "web-proxy/profile/"

type WebProxyProfile struct {
	HeaderClientIp             *string                   `json:"header-client-ip,omitempty"`
	HeaderFrontEndHttps        *string                   `json:"header-front-end-https,omitempty"`
	HeaderViaRequest           *string                   `json:"header-via-request,omitempty"`
	HeaderViaResponse          *string                   `json:"header-via-response,omitempty"`
	HeaderXAuthenticatedGroups *string                   `json:"header-x-authenticated-groups,omitempty"`
	HeaderXAuthenticatedUser   *string                   `json:"header-x-authenticated-user,omitempty"`
	HeaderXForwardedClientCert *string                   `json:"header-x-forwarded-client-cert,omitempty"`
	HeaderXForwardedFor        *string                   `json:"header-x-forwarded-for,omitempty"`
	Headers                    *[]WebProxyProfileHeaders `json:"headers,omitempty"`
	LogHeaderChange            *string                   `json:"log-header-change,omitempty"`
	Name                       *string                   `json:"name,omitempty"`
	StripEncoding              *string                   `json:"strip-encoding,omitempty"`
}

const WebProxyProfileHeadersPath = "web-proxy/profile/headers/"

type WebProxyProfileHeaders struct {
	Action         *string                           `json:"action,omitempty"`
	AddOption      *string                           `json:"add-option,omitempty"`
	Base64Encoding *string                           `json:"base64-encoding,omitempty"`
	Content        *string                           `json:"content,omitempty"`
	Dstaddr        *[]WebProxyProfileHeadersDstaddr  `json:"dstaddr,omitempty"`
	Dstaddr6       *[]WebProxyProfileHeadersDstaddr6 `json:"dstaddr6,omitempty"`
	Id             *int64                            `json:"id,omitempty"`
	Name           *string                           `json:"name,omitempty"`
	Protocol       *string                           `json:"protocol,omitempty"`
}

const WebProxyProfileHeadersDstaddrPath = "web-proxy/profile/headers/dstaddr/"

type WebProxyProfileHeadersDstaddr struct {
	Name *string `json:"name,omitempty"`
}

const WebProxyProfileHeadersDstaddr6Path = "web-proxy/profile/headers/dstaddr6/"

type WebProxyProfileHeadersDstaddr6 struct {
	Name *string `json:"name,omitempty"`
}
