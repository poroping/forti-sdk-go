package models

const VpnipsecFecPath = "vpn.ipsec/fec/"

type VpnipsecFec struct {
	Mappings *[]VpnipsecFecMappings `json:"mappings,omitempty"`
	Name     *string                `json:"name,omitempty"`
}

type VpnipsecFecMappings struct {
	BandwidthBiThreshold   *float64 `json:"bandwidth-bi-threshold,omitempty"`
	BandwidthDownThreshold *float64 `json:"bandwidth-down-threshold,omitempty"`
	BandwidthUpThreshold   *float64 `json:"bandwidth-up-threshold,omitempty"`
	Base                   *float64 `json:"base,omitempty"`
	LatencyThreshold       *float64 `json:"latency-threshold,omitempty"`
	PacketLossThreshold    *float64 `json:"packet-loss-threshold,omitempty"`
	Redundant              *float64 `json:"redundant,omitempty"`
	Seqno                  *float64 `json:"seqno,omitempty"`
}
