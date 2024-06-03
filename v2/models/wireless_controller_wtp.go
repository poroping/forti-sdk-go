package models

const WirelessControllerWtpPath = "wireless-controller/wtp/"

type WirelessControllerWtp struct {
	Admin                          *string                                   `json:"admin,omitempty"`
	Allowaccess                    *string                                   `json:"allowaccess,omitempty"`
	ApcfgProfile                   *string                                   `json:"apcfg-profile,omitempty"`
	BleMajorId                     *int64                                    `json:"ble-major-id,omitempty"`
	BleMinorId                     *int64                                    `json:"ble-minor-id,omitempty"`
	BonjourProfile                 *string                                   `json:"bonjour-profile,omitempty"`
	CoordinateLatitude             *string                                   `json:"coordinate-latitude,omitempty"`
	CoordinateLongitude            *string                                   `json:"coordinate-longitude,omitempty"`
	FirmwareProvision              *string                                   `json:"firmware-provision,omitempty"`
	FirmwareProvisionLatest        *string                                   `json:"firmware-provision-latest,omitempty"`
	ImageDownload                  *string                                   `json:"image-download,omitempty"`
	Index                          *int64                                    `json:"index,omitempty"`
	IpFragmentPreventing           *string                                   `json:"ip-fragment-preventing,omitempty"`
	Lan                            *WirelessControllerWtpLan                 `json:"lan,omitempty"`
	LedState                       *string                                   `json:"led-state,omitempty"`
	Location                       *string                                   `json:"location,omitempty"`
	LoginPasswd                    *string                                   `json:"login-passwd,omitempty"`
	LoginPasswdChange              *string                                   `json:"login-passwd-change,omitempty"`
	MeshBridgeEnable               *string                                   `json:"mesh-bridge-enable,omitempty"`
	Name                           *string                                   `json:"name,omitempty"`
	OverrideAllowaccess            *string                                   `json:"override-allowaccess,omitempty"`
	OverrideIpFragment             *string                                   `json:"override-ip-fragment,omitempty"`
	OverrideLan                    *string                                   `json:"override-lan,omitempty"`
	OverrideLedState               *string                                   `json:"override-led-state,omitempty"`
	OverrideLoginPasswdChange      *string                                   `json:"override-login-passwd-change,omitempty"`
	OverrideSplitTunnel            *string                                   `json:"override-split-tunnel,omitempty"`
	OverrideWanPortMode            *string                                   `json:"override-wan-port-mode,omitempty"`
	Radio1                         *WirelessControllerWtpRadio1              `json:"radio-1,omitempty"`
	Radio2                         *WirelessControllerWtpRadio2              `json:"radio-2,omitempty"`
	Radio3                         *WirelessControllerWtpRadio3              `json:"radio-3,omitempty"`
	Radio4                         *WirelessControllerWtpRadio4              `json:"radio-4,omitempty"`
	Region                         *string                                   `json:"region,omitempty"`
	RegionX                        *string                                   `json:"region-x,omitempty"`
	RegionY                        *string                                   `json:"region-y,omitempty"`
	SplitTunnelingAcl              *[]WirelessControllerWtpSplitTunnelingAcl `json:"split-tunneling-acl,omitempty"`
	SplitTunnelingAclLocalApSubnet *string                                   `json:"split-tunneling-acl-local-ap-subnet,omitempty"`
	SplitTunnelingAclPath          *string                                   `json:"split-tunneling-acl-path,omitempty"`
	TunMtuDownlink                 *int64                                    `json:"tun-mtu-downlink,omitempty"`
	TunMtuUplink                   *int64                                    `json:"tun-mtu-uplink,omitempty"`
	Uuid                           *string                                   `json:"uuid,omitempty"`
	WanPortMode                    *string                                   `json:"wan-port-mode,omitempty"`
	WtpId                          *string                                   `json:"wtp-id,omitempty"`
	WtpMode                        *string                                   `json:"wtp-mode,omitempty"`
	WtpProfile                     *string                                   `json:"wtp-profile,omitempty"`
}

type WirelessControllerWtpLan struct {
	PortEslMode *string `json:"port-esl-mode,omitempty"`
	PortEslSsid *string `json:"port-esl-ssid,omitempty"`
	PortMode    *string `json:"port-mode,omitempty"`
	PortSsid    *string `json:"port-ssid,omitempty"`
	Port1Mode   *string `json:"port1-mode,omitempty"`
	Port1Ssid   *string `json:"port1-ssid,omitempty"`
	Port2Mode   *string `json:"port2-mode,omitempty"`
	Port2Ssid   *string `json:"port2-ssid,omitempty"`
	Port3Mode   *string `json:"port3-mode,omitempty"`
	Port3Ssid   *string `json:"port3-ssid,omitempty"`
	Port4Mode   *string `json:"port4-mode,omitempty"`
	Port4Ssid   *string `json:"port4-ssid,omitempty"`
	Port5Mode   *string `json:"port5-mode,omitempty"`
	Port5Ssid   *string `json:"port5-ssid,omitempty"`
	Port6Mode   *string `json:"port6-mode,omitempty"`
	Port6Ssid   *string `json:"port6-ssid,omitempty"`
	Port7Mode   *string `json:"port7-mode,omitempty"`
	Port7Ssid   *string `json:"port7-ssid,omitempty"`
	Port8Mode   *string `json:"port8-mode,omitempty"`
	Port8Ssid   *string `json:"port8-ssid,omitempty"`
}

type WirelessControllerWtpRadio1 struct {
	AutoPowerHigh    *int64                                `json:"auto-power-high,omitempty"`
	AutoPowerLevel   *string                               `json:"auto-power-level,omitempty"`
	AutoPowerLow     *int64                                `json:"auto-power-low,omitempty"`
	AutoPowerTarget  *string                               `json:"auto-power-target,omitempty"`
	Band             *string                               `json:"band,omitempty"`
	Channel          *[]WirelessControllerWtpRadio1Channel `json:"channel,omitempty"`
	DrmaManualMode   *string                               `json:"drma-manual-mode,omitempty"`
	OverrideAnalysis *string                               `json:"override-analysis,omitempty"`
	OverrideBand     *string                               `json:"override-band,omitempty"`
	OverrideChannel  *string                               `json:"override-channel,omitempty"`
	OverrideTxpower  *string                               `json:"override-txpower,omitempty"`
	OverrideVaps     *string                               `json:"override-vaps,omitempty"`
	PowerLevel       *int64                                `json:"power-level,omitempty"`
	PowerMode        *string                               `json:"power-mode,omitempty"`
	PowerValue       *int64                                `json:"power-value,omitempty"`
	SpectrumAnalysis *string                               `json:"spectrum-analysis,omitempty"`
	VapAll           *string                               `json:"vap-all,omitempty"`
	Vaps             *[]WirelessControllerWtpRadio1Vaps    `json:"vaps,omitempty"`
}

type WirelessControllerWtpRadio1Channel struct {
	Chan *string `json:"chan,omitempty"`
}

type WirelessControllerWtpRadio1Vaps struct {
	Name *string `json:"name,omitempty"`
}

type WirelessControllerWtpRadio2 struct {
	AutoPowerHigh    *int64                                `json:"auto-power-high,omitempty"`
	AutoPowerLevel   *string                               `json:"auto-power-level,omitempty"`
	AutoPowerLow     *int64                                `json:"auto-power-low,omitempty"`
	AutoPowerTarget  *string                               `json:"auto-power-target,omitempty"`
	Band             *string                               `json:"band,omitempty"`
	Channel          *[]WirelessControllerWtpRadio2Channel `json:"channel,omitempty"`
	DrmaManualMode   *string                               `json:"drma-manual-mode,omitempty"`
	OverrideAnalysis *string                               `json:"override-analysis,omitempty"`
	OverrideBand     *string                               `json:"override-band,omitempty"`
	OverrideChannel  *string                               `json:"override-channel,omitempty"`
	OverrideTxpower  *string                               `json:"override-txpower,omitempty"`
	OverrideVaps     *string                               `json:"override-vaps,omitempty"`
	PowerLevel       *int64                                `json:"power-level,omitempty"`
	PowerMode        *string                               `json:"power-mode,omitempty"`
	PowerValue       *int64                                `json:"power-value,omitempty"`
	SpectrumAnalysis *string                               `json:"spectrum-analysis,omitempty"`
	VapAll           *string                               `json:"vap-all,omitempty"`
	Vaps             *[]WirelessControllerWtpRadio2Vaps    `json:"vaps,omitempty"`
}

type WirelessControllerWtpRadio2Channel struct {
	Chan *string `json:"chan,omitempty"`
}

type WirelessControllerWtpRadio2Vaps struct {
	Name *string `json:"name,omitempty"`
}

type WirelessControllerWtpRadio3 struct {
	AutoPowerHigh    *int64                                `json:"auto-power-high,omitempty"`
	AutoPowerLevel   *string                               `json:"auto-power-level,omitempty"`
	AutoPowerLow     *int64                                `json:"auto-power-low,omitempty"`
	AutoPowerTarget  *string                               `json:"auto-power-target,omitempty"`
	Band             *string                               `json:"band,omitempty"`
	Channel          *[]WirelessControllerWtpRadio3Channel `json:"channel,omitempty"`
	DrmaManualMode   *string                               `json:"drma-manual-mode,omitempty"`
	OverrideAnalysis *string                               `json:"override-analysis,omitempty"`
	OverrideBand     *string                               `json:"override-band,omitempty"`
	OverrideChannel  *string                               `json:"override-channel,omitempty"`
	OverrideTxpower  *string                               `json:"override-txpower,omitempty"`
	OverrideVaps     *string                               `json:"override-vaps,omitempty"`
	PowerLevel       *int64                                `json:"power-level,omitempty"`
	PowerMode        *string                               `json:"power-mode,omitempty"`
	PowerValue       *int64                                `json:"power-value,omitempty"`
	SpectrumAnalysis *string                               `json:"spectrum-analysis,omitempty"`
	VapAll           *string                               `json:"vap-all,omitempty"`
	Vaps             *[]WirelessControllerWtpRadio3Vaps    `json:"vaps,omitempty"`
}

type WirelessControllerWtpRadio3Channel struct {
	Chan *string `json:"chan,omitempty"`
}

type WirelessControllerWtpRadio3Vaps struct {
	Name *string `json:"name,omitempty"`
}

type WirelessControllerWtpRadio4 struct {
	AutoPowerHigh    *int64                                `json:"auto-power-high,omitempty"`
	AutoPowerLevel   *string                               `json:"auto-power-level,omitempty"`
	AutoPowerLow     *int64                                `json:"auto-power-low,omitempty"`
	AutoPowerTarget  *string                               `json:"auto-power-target,omitempty"`
	Band             *string                               `json:"band,omitempty"`
	Channel          *[]WirelessControllerWtpRadio4Channel `json:"channel,omitempty"`
	DrmaManualMode   *string                               `json:"drma-manual-mode,omitempty"`
	OverrideAnalysis *string                               `json:"override-analysis,omitempty"`
	OverrideBand     *string                               `json:"override-band,omitempty"`
	OverrideChannel  *string                               `json:"override-channel,omitempty"`
	OverrideTxpower  *string                               `json:"override-txpower,omitempty"`
	OverrideVaps     *string                               `json:"override-vaps,omitempty"`
	PowerLevel       *int64                                `json:"power-level,omitempty"`
	PowerMode        *string                               `json:"power-mode,omitempty"`
	PowerValue       *int64                                `json:"power-value,omitempty"`
	SpectrumAnalysis *string                               `json:"spectrum-analysis,omitempty"`
	VapAll           *string                               `json:"vap-all,omitempty"`
	Vaps             *[]WirelessControllerWtpRadio4Vaps    `json:"vaps,omitempty"`
}

type WirelessControllerWtpRadio4Channel struct {
	Chan *string `json:"chan,omitempty"`
}

type WirelessControllerWtpRadio4Vaps struct {
	Name *string `json:"name,omitempty"`
}

type WirelessControllerWtpSplitTunnelingAcl struct {
	DestIp *string `json:"dest-ip,omitempty"`
	Id     *int64  `json:"id,omitempty"`
}
