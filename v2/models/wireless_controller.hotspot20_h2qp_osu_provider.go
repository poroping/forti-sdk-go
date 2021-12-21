package models

const WirelessControllerhotspot20H2qpOsuProviderPath = "wireless-controller.hotspot20/h2qp-osu-provider/"

type WirelessControllerhotspot20H2qpOsuProvider struct {
	FriendlyName       *[]WirelessControllerhotspot20H2qpOsuProviderFriendlyName       `json:"friendly-name,omitempty"`
	Icon               *string                                                         `json:"icon,omitempty"`
	Name               *string                                                         `json:"name,omitempty"`
	OsuMethod          *string                                                         `json:"osu-method,omitempty"`
	OsuNai             *string                                                         `json:"osu-nai,omitempty"`
	ServerUri          *string                                                         `json:"server-uri,omitempty"`
	ServiceDescription *[]WirelessControllerhotspot20H2qpOsuProviderServiceDescription `json:"service-description,omitempty"`
}

type WirelessControllerhotspot20H2qpOsuProviderFriendlyName struct {
	FriendlyName *string `json:"friendly-name,omitempty"`
	Index        *int64  `json:"index,omitempty"`
	Lang         *string `json:"lang,omitempty"`
}

type WirelessControllerhotspot20H2qpOsuProviderServiceDescription struct {
	Lang               *string `json:"lang,omitempty"`
	ServiceDescription *string `json:"service-description,omitempty"`
	ServiceId          *int64  `json:"service-id,omitempty"`
}
