package models

const SystemGreTunnelPath = "system/gre-tunnel/"

type SystemGreTunnel struct {
	AutoAsicOffload            *string `json:"auto-asic-offload,omitempty"`
	ChecksumReception          *string `json:"checksum-reception,omitempty"`
	ChecksumTransmission       *string `json:"checksum-transmission,omitempty"`
	Diffservcode               *string `json:"diffservcode,omitempty"`
	DscpCopying                *string `json:"dscp-copying,omitempty"`
	Interface                  *string `json:"interface,omitempty"`
	IpVersion                  *string `json:"ip-version,omitempty"`
	KeepaliveFailtimes         *int64  `json:"keepalive-failtimes,omitempty"`
	KeepaliveInterval          *int64  `json:"keepalive-interval,omitempty"`
	KeyInbound                 *int64  `json:"key-inbound,omitempty"`
	KeyOutbound                *int64  `json:"key-outbound,omitempty"`
	LocalGw                    *string `json:"local-gw,omitempty"`
	LocalGw6                   *string `json:"local-gw6,omitempty"`
	Name                       *string `json:"name,omitempty"`
	RemoteGw                   *string `json:"remote-gw,omitempty"`
	RemoteGw6                  *string `json:"remote-gw6,omitempty"`
	SequenceNumberReception    *string `json:"sequence-number-reception,omitempty"`
	SequenceNumberTransmission *string `json:"sequence-number-transmission,omitempty"`
	UseSdwan                   *string `json:"use-sdwan,omitempty"`
}
