package models

const SystemNetworkVisibilityPath = "system/network-visibility/"

type SystemNetworkVisibility struct {
	DestinationHostnameVisibility *string `json:"destination-hostname-visibility,omitempty"`
	DestinationLocation           *string `json:"destination-location,omitempty"`
	DestinationVisibility         *string `json:"destination-visibility,omitempty"`
	HostnameLimit                 *int64  `json:"hostname-limit,omitempty"`
	HostnameTtl                   *int64  `json:"hostname-ttl,omitempty"`
	SourceLocation                *string `json:"source-location,omitempty"`
}
