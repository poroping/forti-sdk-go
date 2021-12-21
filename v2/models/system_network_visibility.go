package models

const SystemNetworkVisibilityPath = "system/network-visibility/"

type SystemNetworkVisibility struct {
	DestinationHostnameVisibility *string  `json:"destination-hostname-visibility,omitempty"`
	DestinationLocation           *string  `json:"destination-location,omitempty"`
	DestinationVisibility         *string  `json:"destination-visibility,omitempty"`
	HostnameLimit                 *float64 `json:"hostname-limit,omitempty"`
	HostnameTtl                   *float64 `json:"hostname-ttl,omitempty"`
	SourceLocation                *string  `json:"source-location,omitempty"`
}
