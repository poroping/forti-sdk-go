package models

const WirelessControllerHotspot20AnqpNetworkAuthTypePath = "wireless-controller.hotspot20/anqp-network-auth-type/"

type WirelessControllerHotspot20AnqpNetworkAuthType struct {
	AuthType *string `json:"auth-type,omitempty"`
	Name     *string `json:"name,omitempty"`
	Url      *string `json:"url,omitempty"`
}
