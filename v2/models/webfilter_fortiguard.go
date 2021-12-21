package models

const WebfilterFortiguardPath = "webfilter/fortiguard/"

type WebfilterFortiguard struct {
	CacheMemPercent        *float64 `json:"cache-mem-percent,omitempty"`
	CacheMode              *string  `json:"cache-mode,omitempty"`
	CachePrefixMatch       *string  `json:"cache-prefix-match,omitempty"`
	ClosePorts             *string  `json:"close-ports,omitempty"`
	OvrdAuthHttps          *string  `json:"ovrd-auth-https,omitempty"`
	OvrdAuthPortHttp       *float64 `json:"ovrd-auth-port-http,omitempty"`
	OvrdAuthPortHttps      *float64 `json:"ovrd-auth-port-https,omitempty"`
	OvrdAuthPortHttpsFlow  *float64 `json:"ovrd-auth-port-https-flow,omitempty"`
	OvrdAuthPortWarning    *float64 `json:"ovrd-auth-port-warning,omitempty"`
	RequestPacketSizeLimit *float64 `json:"request-packet-size-limit,omitempty"`
	WarnAuthHttps          *string  `json:"warn-auth-https,omitempty"`
}
