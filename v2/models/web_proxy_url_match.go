package models

const WebProxyUrlMatchPath = "web-proxy/url-match/"

type WebProxyUrlMatch struct {
	CacheExemption *string `json:"cache-exemption,omitempty"`
	Comment        *string `json:"comment,omitempty"`
	ForwardServer  *string `json:"forward-server,omitempty"`
	Name           *string `json:"name,omitempty"`
	Status         *string `json:"status,omitempty"`
	UrlPattern     *string `json:"url-pattern,omitempty"`
}
