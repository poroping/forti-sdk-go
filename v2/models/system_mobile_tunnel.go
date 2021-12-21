package models

const SystemMobileTunnelPath = "system/mobile-tunnel/"

type SystemMobileTunnel struct {
	HashAlgorithm    *string                      `json:"hash-algorithm,omitempty"`
	HomeAddress      *string                      `json:"home-address,omitempty"`
	HomeAgent        *string                      `json:"home-agent,omitempty"`
	Lifetime         *float64                     `json:"lifetime,omitempty"`
	NMhaeKey         *string                      `json:"n-mhae-key,omitempty"`
	NMhaeKeyType     *string                      `json:"n-mhae-key-type,omitempty"`
	NMhaeSpi         *float64                     `json:"n-mhae-spi,omitempty"`
	Name             *string                      `json:"name,omitempty"`
	Network          *[]SystemMobileTunnelNetwork `json:"network,omitempty"`
	RegInterval      *float64                     `json:"reg-interval,omitempty"`
	RegRetry         *float64                     `json:"reg-retry,omitempty"`
	RenewInterval    *float64                     `json:"renew-interval,omitempty"`
	RoamingInterface *string                      `json:"roaming-interface,omitempty"`
	Status           *string                      `json:"status,omitempty"`
	TunnelMode       *string                      `json:"tunnel-mode,omitempty"`
}

type SystemMobileTunnelNetwork struct {
	Id        *float64 `json:"id,omitempty"`
	Interface *string  `json:"interface,omitempty"`
	Prefix    *string  `json:"prefix,omitempty"`
}
