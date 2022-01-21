package models

const WirelessControllerHotspot20H2qpOsuProviderNaiPath = "wireless-controller.hotspot20/h2qp-osu-provider-nai/"

type WirelessControllerHotspot20H2qpOsuProviderNai struct {
	NaiList *[]WirelessControllerHotspot20H2qpOsuProviderNaiNaiList `json:"nai-list,omitempty"`
	Name    *string                                                 `json:"name,omitempty"`
}

type WirelessControllerHotspot20H2qpOsuProviderNaiNaiList struct {
	Name   *string `json:"name,omitempty"`
	OsuNai *string `json:"osu-nai,omitempty"`
}
