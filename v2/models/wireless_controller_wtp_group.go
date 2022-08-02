package models

const WirelessControllerWtpGroupPath = "wireless-controller/wtp-group/"

type WirelessControllerWtpGroup struct {
	Name         *string                           `json:"name,omitempty"`
	PlatformType *string                           `json:"platform-type,omitempty"`
	Wtps         *[]WirelessControllerWtpGroupWtps `json:"wtps,omitempty"`
}

const WirelessControllerWtpGroupWtpsPath = "wireless-controller/wtp-group/wtps/"

type WirelessControllerWtpGroupWtps struct {
	WtpId *string `json:"wtp-id,omitempty"`
}
