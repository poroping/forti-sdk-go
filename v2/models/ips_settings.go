package models

const IpsSettingsPath = "ips/settings/"

type IpsSettings struct {
	IpsPacketQuota      *float64 `json:"ips-packet-quota,omitempty"`
	PacketLogHistory    *float64 `json:"packet-log-history,omitempty"`
	PacketLogMemory     *float64 `json:"packet-log-memory,omitempty"`
	PacketLogPostAttack *float64 `json:"packet-log-post-attack,omitempty"`
}
