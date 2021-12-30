package models

const VpnipsecManualkeyInterfacePath = "vpn/ipsec/manualkey-interface/"

type VpnipsecManualkeyInterface struct {
	AddrType   *string `json:"addr-type,omitempty"`
	AuthAlg    *string `json:"auth-alg,omitempty"`
	AuthKey    *string `json:"auth-key,omitempty"`
	EncAlg     *string `json:"enc-alg,omitempty"`
	EncKey     *string `json:"enc-key,omitempty"`
	Interface  *string `json:"interface,omitempty"`
	IpVersion  *string `json:"ip-version,omitempty"`
	LocalGw    *string `json:"local-gw,omitempty"`
	LocalGw6   *string `json:"local-gw6,omitempty"`
	LocalSpi   *string `json:"local-spi,omitempty"`
	Name       *string `json:"name,omitempty"`
	NpuOffload *string `json:"npu-offload,omitempty"`
	RemoteGw   *string `json:"remote-gw,omitempty"`
	RemoteGw6  *string `json:"remote-gw6,omitempty"`
	RemoteSpi  *string `json:"remote-spi,omitempty"`
}
