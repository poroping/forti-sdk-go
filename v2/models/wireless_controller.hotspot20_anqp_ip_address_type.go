package models

const WirelessControllerHotspot20AnqpIpAddressTypePath = "wireless-controller.hotspot20/anqp-ip-address-type/"

type WirelessControllerHotspot20AnqpIpAddressType struct {
	Ipv4AddressType *string `json:"ipv4-address-type,omitempty"`
	Ipv6AddressType *string `json:"ipv6-address-type,omitempty"`
	Name            *string `json:"name,omitempty"`
}
