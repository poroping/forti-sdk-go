package models

const VpnIpsecManualkeyPath = "vpn.ipsec/manualkey/"

type VpnIpsecManualkey struct {
	Authentication *string `json:"authentication,omitempty"`
	Authkey        *string `json:"authkey,omitempty"`
	Enckey         *string `json:"enckey,omitempty"`
	Encryption     *string `json:"encryption,omitempty"`
	Interface      *string `json:"interface,omitempty"`
	LocalGw        *string `json:"local-gw,omitempty"`
	Localspi       *string `json:"localspi,omitempty"`
	Name           *string `json:"name,omitempty"`
	NpuOffload     *string `json:"npu-offload,omitempty"`
	RemoteGw       *string `json:"remote-gw,omitempty"`
	Remotespi      *string `json:"remotespi,omitempty"`
}
