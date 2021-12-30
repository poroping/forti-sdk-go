package models

const WirelessControllerSettingPath = "wireless-controller/setting/"

type WirelessControllerSetting struct {
	AccountId                        *string                                            `json:"account-id,omitempty"`
	Country                          *string                                            `json:"country,omitempty"`
	DarrpOptimize                    *int64                                             `json:"darrp-optimize,omitempty"`
	DarrpOptimizeSchedules           *[]WirelessControllerSettingDarrpOptimizeSchedules `json:"darrp-optimize-schedules,omitempty"`
	DeviceHoldoff                    *int64                                             `json:"device-holdoff,omitempty"`
	DeviceIdle                       *int64                                             `json:"device-idle,omitempty"`
	DeviceWeight                     *int64                                             `json:"device-weight,omitempty"`
	DuplicateSsid                    *string                                            `json:"duplicate-ssid,omitempty"`
	FakeSsidAction                   *string                                            `json:"fake-ssid-action,omitempty"`
	FapcCompatibility                *string                                            `json:"fapc-compatibility,omitempty"`
	FirmwareProvisionOnAuthorization *string                                            `json:"firmware-provision-on-authorization,omitempty"`
	OffendingSsid                    *[]WirelessControllerSettingOffendingSsid          `json:"offending-ssid,omitempty"`
	PhishingSsidDetect               *string                                            `json:"phishing-ssid-detect,omitempty"`
	WfaCompatibility                 *string                                            `json:"wfa-compatibility,omitempty"`
}

type WirelessControllerSettingDarrpOptimizeSchedules struct {
	Name *string `json:"name,omitempty"`
}

type WirelessControllerSettingOffendingSsid struct {
	Action      *string `json:"action,omitempty"`
	Id          *int64  `json:"id,omitempty"`
	SsidPattern *string `json:"ssid-pattern,omitempty"`
}
