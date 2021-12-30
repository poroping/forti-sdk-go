package models

const IpsSettingsPath = "ips/settings/"

type IpsSettings struct {
	IpsPacketQuota      *int64 `json:"ips-packet-quota,omitempty"`
	PacketLogHistory    *int64 `json:"packet-log-history,omitempty"`
	PacketLogMemory     *int64 `json:"packet-log-memory,omitempty"`
	PacketLogPostAttack *int64 `json:"packet-log-post-attack,omitempty"`
}
