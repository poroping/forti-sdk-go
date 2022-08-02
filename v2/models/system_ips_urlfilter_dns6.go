package models

const SystemIpsUrlfilterDns6Path = "system/ips-urlfilter-dns6/"

type SystemIpsUrlfilterDns6 struct {
	Address6 *string `json:"address6,omitempty"`
	Status   *string `json:"status,omitempty"`
}
