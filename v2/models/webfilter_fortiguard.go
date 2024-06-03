package models

const WebfilterFortiguardPath = "webfilter/fortiguard/"

type WebfilterFortiguard struct {
	CacheMemPercent        *int64  `json:"cache-mem-percent,omitempty"`
	CacheMode              *string `json:"cache-mode,omitempty"`
	CachePrefixMatch       *string `json:"cache-prefix-match,omitempty"`
	ClosePorts             *string `json:"close-ports,omitempty"`
	EmbedImage             *string `json:"embed-image,omitempty"`
	OvrdAuthHttps          *string `json:"ovrd-auth-https,omitempty"`
	OvrdAuthPortHttp       *int64  `json:"ovrd-auth-port-http,omitempty"`
	OvrdAuthPortHttps      *int64  `json:"ovrd-auth-port-https,omitempty"`
	OvrdAuthPortHttpsFlow  *int64  `json:"ovrd-auth-port-https-flow,omitempty"`
	OvrdAuthPortWarning    *int64  `json:"ovrd-auth-port-warning,omitempty"`
	RequestPacketSizeLimit *int64  `json:"request-packet-size-limit,omitempty"`
	WarnAuthHttps          *string `json:"warn-auth-https,omitempty"`
}
