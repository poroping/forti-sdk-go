package models

const FirewallsslSettingPath = "firewall.ssl/setting/"

type FirewallsslSetting struct {
	AbbreviateHandshake    *string `json:"abbreviate-handshake,omitempty"`
	CertCacheCapacity      *int64  `json:"cert-cache-capacity,omitempty"`
	CertCacheTimeout       *int64  `json:"cert-cache-timeout,omitempty"`
	KxpQueueThreshold      *int64  `json:"kxp-queue-threshold,omitempty"`
	NoMatchingCipherAction *string `json:"no-matching-cipher-action,omitempty"`
	ProxyConnectTimeout    *int64  `json:"proxy-connect-timeout,omitempty"`
	SessionCacheCapacity   *int64  `json:"session-cache-capacity,omitempty"`
	SessionCacheTimeout    *int64  `json:"session-cache-timeout,omitempty"`
	SslDhBits              *string `json:"ssl-dh-bits,omitempty"`
	SslQueueThreshold      *int64  `json:"ssl-queue-threshold,omitempty"`
	SslSendEmptyFrags      *string `json:"ssl-send-empty-frags,omitempty"`
}
