package models

const WirelessControllerWagProfilePath = "wireless-controller/wag-profile/"

type WirelessControllerWagProfile struct {
	Comment             *string `json:"comment,omitempty"`
	DhcpIpAddr          *string `json:"dhcp-ip-addr,omitempty"`
	Name                *string `json:"name,omitempty"`
	PingInterval        *int64  `json:"ping-interval,omitempty"`
	PingNumber          *int64  `json:"ping-number,omitempty"`
	ReturnPacketTimeout *int64  `json:"return-packet-timeout,omitempty"`
	TunnelType          *string `json:"tunnel-type,omitempty"`
	WagIp               *string `json:"wag-ip,omitempty"`
	WagPort             *int64  `json:"wag-port,omitempty"`
}
