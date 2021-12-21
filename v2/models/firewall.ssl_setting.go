package models

const FirewallsslSettingPath = "firewall.ssl/setting/"

type FirewallsslSetting struct {
	AbbreviateHandshake    *string  `json:"abbreviate-handshake,omitempty"`
	CertCacheCapacity      *float64 `json:"cert-cache-capacity,omitempty"`
	CertCacheTimeout       *float64 `json:"cert-cache-timeout,omitempty"`
	KxpQueueThreshold      *float64 `json:"kxp-queue-threshold,omitempty"`
	NoMatchingCipherAction *string  `json:"no-matching-cipher-action,omitempty"`
	ProxyConnectTimeout    *float64 `json:"proxy-connect-timeout,omitempty"`
	SessionCacheCapacity   *float64 `json:"session-cache-capacity,omitempty"`
	SessionCacheTimeout    *float64 `json:"session-cache-timeout,omitempty"`
	SslDhBits              *string  `json:"ssl-dh-bits,omitempty"`
	SslQueueThreshold      *float64 `json:"ssl-queue-threshold,omitempty"`
	SslSendEmptyFrags      *string  `json:"ssl-send-empty-frags,omitempty"`
}
