package models

const WebProxyDebugUrlPath = "web-proxy/debug-url/"

type WebProxyDebugUrl struct {
	Exact      *string `json:"exact,omitempty"`
	Name       *string `json:"name,omitempty"`
	Status     *string `json:"status,omitempty"`
	UrlPattern *string `json:"url-pattern,omitempty"`
}
