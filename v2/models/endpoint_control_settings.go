package models

const EndpointControlSettingsPath = "endpoint-control/settings/"

type EndpointControlSettings struct {
	ForticlientDisconnectUnsupportedClient *string `json:"forticlient-disconnect-unsupported-client,omitempty"`
	ForticlientKeepaliveInterval           *int64  `json:"forticlient-keepalive-interval,omitempty"`
	ForticlientSysUpdateInterval           *int64  `json:"forticlient-sys-update-interval,omitempty"`
	ForticlientUserAvatar                  *string `json:"forticlient-user-avatar,omitempty"`
}
