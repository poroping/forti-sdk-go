package models

const WirelessControllerRegionPath = "wireless-controller/region/"

type WirelessControllerRegion struct {
	Comments  *string  `json:"comments,omitempty"`
	Grayscale *string  `json:"grayscale,omitempty"`
	Name      *string  `json:"name,omitempty"`
	Opacity   *float64 `json:"opacity,omitempty"`
}
