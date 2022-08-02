package models

const WirelessControllerHotspot20H2qpOsuProviderPath = "wireless-controller.hotspot20/h2qp-osu-provider/"

type WirelessControllerHotspot20H2qpOsuProvider struct {
	FriendlyName       *[]WirelessControllerHotspot20H2qpOsuProviderFriendlyName       `json:"friendly-name,omitempty"`
	Icon               *string                                                         `json:"icon,omitempty"`
	Name               *string                                                         `json:"name,omitempty"`
	OsuMethod          *string                                                         `json:"osu-method,omitempty"`
	OsuNai             *string                                                         `json:"osu-nai,omitempty"`
	ServerUri          *string                                                         `json:"server-uri,omitempty"`
	ServiceDescription *[]WirelessControllerHotspot20H2qpOsuProviderServiceDescription `json:"service-description,omitempty"`
}

const WirelessControllerHotspot20H2qpOsuProviderFriendlyNamePath = "wireless-controller.hotspot20/h2qp-osu-provider/friendly-name/"

type WirelessControllerHotspot20H2qpOsuProviderFriendlyName struct {
	FriendlyName *string `json:"friendly-name,omitempty"`
	Index        *int64  `json:"index,omitempty"`
	Lang         *string `json:"lang,omitempty"`
}

const WirelessControllerHotspot20H2qpOsuProviderServiceDescriptionPath = "wireless-controller.hotspot20/h2qp-osu-provider/service-description/"

type WirelessControllerHotspot20H2qpOsuProviderServiceDescription struct {
	Lang               *string `json:"lang,omitempty"`
	ServiceDescription *string `json:"service-description,omitempty"`
	ServiceId          *int64  `json:"service-id,omitempty"`
}
