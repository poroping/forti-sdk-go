package models

const EndpointControlSettingsPath = "endpoint-control/settings/"

type EndpointControlSettings struct {
	ForticlientDisconnectUnsupportedClient *string  `json:"forticlient-disconnect-unsupported-client,omitempty"`
	ForticlientKeepaliveInterval           *float64 `json:"forticlient-keepalive-interval,omitempty"`
	ForticlientSysUpdateInterval           *float64 `json:"forticlient-sys-update-interval,omitempty"`
	ForticlientUserAvatar                  *string  `json:"forticlient-user-avatar,omitempty"`
}
