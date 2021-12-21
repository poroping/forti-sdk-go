package models

const WirelessControllerhotspot20IconPath = "wireless-controller.hotspot20/icon/"

type WirelessControllerhotspot20Icon struct {
	IconList *[]WirelessControllerhotspot20IconIconList `json:"icon-list,omitempty"`
	Name     *string                                    `json:"name,omitempty"`
}

type WirelessControllerhotspot20IconIconList struct {
	File   *string `json:"file,omitempty"`
	Height *int64  `json:"height,omitempty"`
	Lang   *string `json:"lang,omitempty"`
	Name   *string `json:"name,omitempty"`
	Type   *string `json:"type,omitempty"`
	Width  *int64  `json:"width,omitempty"`
}
