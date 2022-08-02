package models

const WirelessControllerHotspot20IconPath = "wireless-controller.hotspot20/icon/"

type WirelessControllerHotspot20Icon struct {
	IconList *[]WirelessControllerHotspot20IconIconList `json:"icon-list,omitempty"`
	Name     *string                                    `json:"name,omitempty"`
}

const WirelessControllerHotspot20IconIconListPath = "wireless-controller.hotspot20/icon/icon-list/"

type WirelessControllerHotspot20IconIconList struct {
	File   *string `json:"file,omitempty"`
	Height *int64  `json:"height,omitempty"`
	Lang   *string `json:"lang,omitempty"`
	Name   *string `json:"name,omitempty"`
	Type   *string `json:"type,omitempty"`
	Width  *int64  `json:"width,omitempty"`
}
