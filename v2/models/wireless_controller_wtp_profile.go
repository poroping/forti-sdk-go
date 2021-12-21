package models

const WirelessControllerWtpProfilePath = "wireless-controller/wtp-profile/"

type WirelessControllerWtpProfile struct {
	Allowaccess                    *string                                          `json:"allowaccess,omitempty"`
	ApCountry                      *string                                          `json:"ap-country,omitempty"`
	ApHandoff                      *string                                          `json:"ap-handoff,omitempty"`
	ApcfgProfile                   *string                                          `json:"apcfg-profile,omitempty"`
	BleProfile                     *string                                          `json:"ble-profile,omitempty"`
	Comment                        *string                                          `json:"comment,omitempty"`
	ConsoleLogin                   *string                                          `json:"console-login,omitempty"`
	ControlMessageOffload          *string                                          `json:"control-message-offload,omitempty"`
	DenyMacList                    *[]WirelessControllerWtpProfileDenyMacList       `json:"deny-mac-list,omitempty"`
	DtlsInKernel                   *string                                          `json:"dtls-in-kernel,omitempty"`
	DtlsPolicy                     *string                                          `json:"dtls-policy,omitempty"`
	EnergyEfficientEthernet        *string                                          `json:"energy-efficient-ethernet,omitempty"`
	EslSesDongle                   *[]WirelessControllerWtpProfileEslSesDongle      `json:"esl-ses-dongle,omitempty"`
	ExtInfoEnable                  *string                                          `json:"ext-info-enable,omitempty"`
	FrequencyHandoff               *string                                          `json:"frequency-handoff,omitempty"`
	HandoffRoaming                 *string                                          `json:"handoff-roaming,omitempty"`
	HandoffRssi                    *float64                                         `json:"handoff-rssi,omitempty"`
	HandoffStaThresh               *float64                                         `json:"handoff-sta-thresh,omitempty"`
	IndoorOutdoorDeployment        *string                                          `json:"indoor-outdoor-deployment,omitempty"`
	IpFragmentPreventing           *string                                          `json:"ip-fragment-preventing,omitempty"`
	Lan                            *[]WirelessControllerWtpProfileLan               `json:"lan,omitempty"`
	Lbs                            *[]WirelessControllerWtpProfileLbs               `json:"lbs,omitempty"`
	LedSchedules                   *[]WirelessControllerWtpProfileLedSchedules      `json:"led-schedules,omitempty"`
	LedState                       *string                                          `json:"led-state,omitempty"`
	Lldp                           *string                                          `json:"lldp,omitempty"`
	LoginPasswd                    *string                                          `json:"login-passwd,omitempty"`
	LoginPasswdChange              *string                                          `json:"login-passwd-change,omitempty"`
	MaxClients                     *float64                                         `json:"max-clients,omitempty"`
	Name                           *string                                          `json:"name,omitempty"`
	Platform                       *[]WirelessControllerWtpProfilePlatform          `json:"platform,omitempty"`
	PoeMode                        *string                                          `json:"poe-mode,omitempty"`
	Radio1                         *[]WirelessControllerWtpProfileRadio1            `json:"radio-1,omitempty"`
	Radio2                         *[]WirelessControllerWtpProfileRadio2            `json:"radio-2,omitempty"`
	Radio3                         *[]WirelessControllerWtpProfileRadio3            `json:"radio-3,omitempty"`
	Radio4                         *[]WirelessControllerWtpProfileRadio4            `json:"radio-4,omitempty"`
	SplitTunnelingAcl              *[]WirelessControllerWtpProfileSplitTunnelingAcl `json:"split-tunneling-acl,omitempty"`
	SplitTunnelingAclLocalApSubnet *string                                          `json:"split-tunneling-acl-local-ap-subnet,omitempty"`
	SplitTunnelingAclPath          *string                                          `json:"split-tunneling-acl-path,omitempty"`
	SyslogProfile                  *string                                          `json:"syslog-profile,omitempty"`
	TunMtuDownlink                 *float64                                         `json:"tun-mtu-downlink,omitempty"`
	TunMtuUplink                   *float64                                         `json:"tun-mtu-uplink,omitempty"`
	WanPortAuth                    *string                                          `json:"wan-port-auth,omitempty"`
	WanPortAuthMethods             *string                                          `json:"wan-port-auth-methods,omitempty"`
	WanPortAuthPassword            *string                                          `json:"wan-port-auth-password,omitempty"`
	WanPortAuthUsrname             *string                                          `json:"wan-port-auth-usrname,omitempty"`
	WanPortMode                    *string                                          `json:"wan-port-mode,omitempty"`
}

type WirelessControllerWtpProfileDenyMacList struct {
	Id  *float64 `json:"id,omitempty"`
	Mac *string  `json:"mac,omitempty"`
}

type WirelessControllerWtpProfileEslSesDongle struct {
	ApcAddrType         *string  `json:"apc-addr-type,omitempty"`
	ApcFqdn             *string  `json:"apc-fqdn,omitempty"`
	ApcIp               *string  `json:"apc-ip,omitempty"`
	ApcPort             *float64 `json:"apc-port,omitempty"`
	CoexLevel           *string  `json:"coex-level,omitempty"`
	ComplianceLevel     *string  `json:"compliance-level,omitempty"`
	EslChannel          *string  `json:"esl-channel,omitempty"`
	OutputPower         *string  `json:"output-power,omitempty"`
	ScdEnable           *string  `json:"scd-enable,omitempty"`
	TlsCertVerification *string  `json:"tls-cert-verification,omitempty"`
	TlsFqdnVerification *string  `json:"tls-fqdn-verification,omitempty"`
}

type WirelessControllerWtpProfileLan struct {
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

type WirelessControllerWtpProfileLbs struct {
	Aeroscout                   *string  `json:"aeroscout,omitempty"`
	AeroscoutApMac              *string  `json:"aeroscout-ap-mac,omitempty"`
	AeroscoutMmuReport          *string  `json:"aeroscout-mmu-report,omitempty"`
	AeroscoutMu                 *string  `json:"aeroscout-mu,omitempty"`
	AeroscoutMuFactor           *float64 `json:"aeroscout-mu-factor,omitempty"`
	AeroscoutMuTimeout          *float64 `json:"aeroscout-mu-timeout,omitempty"`
	AeroscoutServerIp           *string  `json:"aeroscout-server-ip,omitempty"`
	AeroscoutServerPort         *float64 `json:"aeroscout-server-port,omitempty"`
	EkahauBlinkMode             *string  `json:"ekahau-blink-mode,omitempty"`
	EkahauTag                   *string  `json:"ekahau-tag,omitempty"`
	ErcServerIp                 *string  `json:"erc-server-ip,omitempty"`
	ErcServerPort               *float64 `json:"erc-server-port,omitempty"`
	Fortipresence               *string  `json:"fortipresence,omitempty"`
	FortipresenceBle            *string  `json:"fortipresence-ble,omitempty"`
	FortipresenceFrequency      *float64 `json:"fortipresence-frequency,omitempty"`
	FortipresencePort           *float64 `json:"fortipresence-port,omitempty"`
	FortipresenceProject        *string  `json:"fortipresence-project,omitempty"`
	FortipresenceRogue          *string  `json:"fortipresence-rogue,omitempty"`
	FortipresenceSecret         *string  `json:"fortipresence-secret,omitempty"`
	FortipresenceServer         *string  `json:"fortipresence-server,omitempty"`
	FortipresenceServerAddrType *string  `json:"fortipresence-server-addr-type,omitempty"`
	FortipresenceServerFqdn     *string  `json:"fortipresence-server-fqdn,omitempty"`
	FortipresenceUnassoc        *string  `json:"fortipresence-unassoc,omitempty"`
	StationLocate               *string  `json:"station-locate,omitempty"`
}

type WirelessControllerWtpProfileLedSchedules struct {
	Name *string `json:"name,omitempty"`
}

type WirelessControllerWtpProfilePlatform struct {
	Ddscan *string `json:"ddscan,omitempty"`
	Mode   *string `json:"mode,omitempty"`
	Type   *string `json:"type,omitempty"`
}

type WirelessControllerWtpProfileRadio1 struct {
	AirtimeFairness           *string                                      `json:"airtime-fairness,omitempty"`
	Amsdu                     *string                                      `json:"amsdu,omitempty"`
	ApHandoff                 *string                                      `json:"ap-handoff,omitempty"`
	ApSnifferAddr             *string                                      `json:"ap-sniffer-addr,omitempty"`
	ApSnifferBufsize          *float64                                     `json:"ap-sniffer-bufsize,omitempty"`
	ApSnifferChan             *float64                                     `json:"ap-sniffer-chan,omitempty"`
	ApSnifferCtl              *string                                      `json:"ap-sniffer-ctl,omitempty"`
	ApSnifferData             *string                                      `json:"ap-sniffer-data,omitempty"`
	ApSnifferMgmtBeacon       *string                                      `json:"ap-sniffer-mgmt-beacon,omitempty"`
	ApSnifferMgmtOther        *string                                      `json:"ap-sniffer-mgmt-other,omitempty"`
	ApSnifferMgmtProbe        *string                                      `json:"ap-sniffer-mgmt-probe,omitempty"`
	AutoPowerHigh             *float64                                     `json:"auto-power-high,omitempty"`
	AutoPowerLevel            *string                                      `json:"auto-power-level,omitempty"`
	AutoPowerLow              *float64                                     `json:"auto-power-low,omitempty"`
	AutoPowerTarget           *string                                      `json:"auto-power-target,omitempty"`
	Band                      *string                                      `json:"band,omitempty"`
	Band5gType                *string                                      `json:"band-5g-type,omitempty"`
	BandwidthAdmissionControl *string                                      `json:"bandwidth-admission-control,omitempty"`
	BandwidthCapacity         *float64                                     `json:"bandwidth-capacity,omitempty"`
	BeaconInterval            *float64                                     `json:"beacon-interval,omitempty"`
	BssColor                  *float64                                     `json:"bss-color,omitempty"`
	BssColorMode              *string                                      `json:"bss-color-mode,omitempty"`
	CallAdmissionControl      *string                                      `json:"call-admission-control,omitempty"`
	CallCapacity              *float64                                     `json:"call-capacity,omitempty"`
	Channel                   *[]WirelessControllerWtpProfileRadio1Channel `json:"channel,omitempty"`
	ChannelBonding            *string                                      `json:"channel-bonding,omitempty"`
	ChannelUtilization        *string                                      `json:"channel-utilization,omitempty"`
	Coexistence               *string                                      `json:"coexistence,omitempty"`
	Darrp                     *string                                      `json:"darrp,omitempty"`
	Drma                      *string                                      `json:"drma,omitempty"`
	DrmaSensitivity           *string                                      `json:"drma-sensitivity,omitempty"`
	Dtim                      *float64                                     `json:"dtim,omitempty"`
	FragThreshold             *float64                                     `json:"frag-threshold,omitempty"`
	FrequencyHandoff          *string                                      `json:"frequency-handoff,omitempty"`
	IperfProtocol             *string                                      `json:"iperf-protocol,omitempty"`
	IperfServerPort           *float64                                     `json:"iperf-server-port,omitempty"`
	MaxClients                *float64                                     `json:"max-clients,omitempty"`
	MaxDistance               *float64                                     `json:"max-distance,omitempty"`
	Mode                      *string                                      `json:"mode,omitempty"`
	PowerLevel                *float64                                     `json:"power-level,omitempty"`
	PowerMode                 *string                                      `json:"power-mode,omitempty"`
	PowerValue                *float64                                     `json:"power-value,omitempty"`
	PowersaveOptimize         *string                                      `json:"powersave-optimize,omitempty"`
	ProtectionMode            *string                                      `json:"protection-mode,omitempty"`
	RtsThreshold              *float64                                     `json:"rts-threshold,omitempty"`
	SamBssid                  *string                                      `json:"sam-bssid,omitempty"`
	SamCaptivePortal          *string                                      `json:"sam-captive-portal,omitempty"`
	SamCwpFailureString       *string                                      `json:"sam-cwp-failure-string,omitempty"`
	SamCwpMatchString         *string                                      `json:"sam-cwp-match-string,omitempty"`
	SamCwpPassword            *string                                      `json:"sam-cwp-password,omitempty"`
	SamCwpSuccessString       *string                                      `json:"sam-cwp-success-string,omitempty"`
	SamCwpTestUrl             *string                                      `json:"sam-cwp-test-url,omitempty"`
	SamCwpUsername            *string                                      `json:"sam-cwp-username,omitempty"`
	SamPassword               *string                                      `json:"sam-password,omitempty"`
	SamReportIntv             *float64                                     `json:"sam-report-intv,omitempty"`
	SamSecurityType           *string                                      `json:"sam-security-type,omitempty"`
	SamServer                 *string                                      `json:"sam-server,omitempty"`
	SamServerFqdn             *string                                      `json:"sam-server-fqdn,omitempty"`
	SamServerIp               *string                                      `json:"sam-server-ip,omitempty"`
	SamServerType             *string                                      `json:"sam-server-type,omitempty"`
	SamSsid                   *string                                      `json:"sam-ssid,omitempty"`
	SamTest                   *string                                      `json:"sam-test,omitempty"`
	SamUsername               *string                                      `json:"sam-username,omitempty"`
	ShortGuardInterval        *string                                      `json:"short-guard-interval,omitempty"`
	SpectrumAnalysis          *string                                      `json:"spectrum-analysis,omitempty"`
	TransmitOptimize          *string                                      `json:"transmit-optimize,omitempty"`
	VapAll                    *string                                      `json:"vap-all,omitempty"`
	Vaps                      *[]WirelessControllerWtpProfileRadio1Vaps    `json:"vaps,omitempty"`
	WidsProfile               *string                                      `json:"wids-profile,omitempty"`
	ZeroWaitDfs               *string                                      `json:"zero-wait-dfs,omitempty"`
}

type WirelessControllerWtpProfileRadio1Channel struct {
	Chan *string `json:"chan,omitempty"`
}

type WirelessControllerWtpProfileRadio1Vaps struct {
	Name *string `json:"name,omitempty"`
}

type WirelessControllerWtpProfileRadio2 struct {
	AirtimeFairness           *string                                      `json:"airtime-fairness,omitempty"`
	Amsdu                     *string                                      `json:"amsdu,omitempty"`
	ApHandoff                 *string                                      `json:"ap-handoff,omitempty"`
	ApSnifferAddr             *string                                      `json:"ap-sniffer-addr,omitempty"`
	ApSnifferBufsize          *float64                                     `json:"ap-sniffer-bufsize,omitempty"`
	ApSnifferChan             *float64                                     `json:"ap-sniffer-chan,omitempty"`
	ApSnifferCtl              *string                                      `json:"ap-sniffer-ctl,omitempty"`
	ApSnifferData             *string                                      `json:"ap-sniffer-data,omitempty"`
	ApSnifferMgmtBeacon       *string                                      `json:"ap-sniffer-mgmt-beacon,omitempty"`
	ApSnifferMgmtOther        *string                                      `json:"ap-sniffer-mgmt-other,omitempty"`
	ApSnifferMgmtProbe        *string                                      `json:"ap-sniffer-mgmt-probe,omitempty"`
	AutoPowerHigh             *float64                                     `json:"auto-power-high,omitempty"`
	AutoPowerLevel            *string                                      `json:"auto-power-level,omitempty"`
	AutoPowerLow              *float64                                     `json:"auto-power-low,omitempty"`
	AutoPowerTarget           *string                                      `json:"auto-power-target,omitempty"`
	Band                      *string                                      `json:"band,omitempty"`
	Band5gType                *string                                      `json:"band-5g-type,omitempty"`
	BandwidthAdmissionControl *string                                      `json:"bandwidth-admission-control,omitempty"`
	BandwidthCapacity         *float64                                     `json:"bandwidth-capacity,omitempty"`
	BeaconInterval            *float64                                     `json:"beacon-interval,omitempty"`
	BssColor                  *float64                                     `json:"bss-color,omitempty"`
	BssColorMode              *string                                      `json:"bss-color-mode,omitempty"`
	CallAdmissionControl      *string                                      `json:"call-admission-control,omitempty"`
	CallCapacity              *float64                                     `json:"call-capacity,omitempty"`
	Channel                   *[]WirelessControllerWtpProfileRadio2Channel `json:"channel,omitempty"`
	ChannelBonding            *string                                      `json:"channel-bonding,omitempty"`
	ChannelUtilization        *string                                      `json:"channel-utilization,omitempty"`
	Coexistence               *string                                      `json:"coexistence,omitempty"`
	Darrp                     *string                                      `json:"darrp,omitempty"`
	Drma                      *string                                      `json:"drma,omitempty"`
	DrmaSensitivity           *string                                      `json:"drma-sensitivity,omitempty"`
	Dtim                      *float64                                     `json:"dtim,omitempty"`
	FragThreshold             *float64                                     `json:"frag-threshold,omitempty"`
	FrequencyHandoff          *string                                      `json:"frequency-handoff,omitempty"`
	IperfProtocol             *string                                      `json:"iperf-protocol,omitempty"`
	IperfServerPort           *float64                                     `json:"iperf-server-port,omitempty"`
	MaxClients                *float64                                     `json:"max-clients,omitempty"`
	MaxDistance               *float64                                     `json:"max-distance,omitempty"`
	Mode                      *string                                      `json:"mode,omitempty"`
	PowerLevel                *float64                                     `json:"power-level,omitempty"`
	PowerMode                 *string                                      `json:"power-mode,omitempty"`
	PowerValue                *float64                                     `json:"power-value,omitempty"`
	PowersaveOptimize         *string                                      `json:"powersave-optimize,omitempty"`
	ProtectionMode            *string                                      `json:"protection-mode,omitempty"`
	RtsThreshold              *float64                                     `json:"rts-threshold,omitempty"`
	SamBssid                  *string                                      `json:"sam-bssid,omitempty"`
	SamCaptivePortal          *string                                      `json:"sam-captive-portal,omitempty"`
	SamCwpFailureString       *string                                      `json:"sam-cwp-failure-string,omitempty"`
	SamCwpMatchString         *string                                      `json:"sam-cwp-match-string,omitempty"`
	SamCwpPassword            *string                                      `json:"sam-cwp-password,omitempty"`
	SamCwpSuccessString       *string                                      `json:"sam-cwp-success-string,omitempty"`
	SamCwpTestUrl             *string                                      `json:"sam-cwp-test-url,omitempty"`
	SamCwpUsername            *string                                      `json:"sam-cwp-username,omitempty"`
	SamPassword               *string                                      `json:"sam-password,omitempty"`
	SamReportIntv             *float64                                     `json:"sam-report-intv,omitempty"`
	SamSecurityType           *string                                      `json:"sam-security-type,omitempty"`
	SamServer                 *string                                      `json:"sam-server,omitempty"`
	SamServerFqdn             *string                                      `json:"sam-server-fqdn,omitempty"`
	SamServerIp               *string                                      `json:"sam-server-ip,omitempty"`
	SamServerType             *string                                      `json:"sam-server-type,omitempty"`
	SamSsid                   *string                                      `json:"sam-ssid,omitempty"`
	SamTest                   *string                                      `json:"sam-test,omitempty"`
	SamUsername               *string                                      `json:"sam-username,omitempty"`
	ShortGuardInterval        *string                                      `json:"short-guard-interval,omitempty"`
	SpectrumAnalysis          *string                                      `json:"spectrum-analysis,omitempty"`
	TransmitOptimize          *string                                      `json:"transmit-optimize,omitempty"`
	VapAll                    *string                                      `json:"vap-all,omitempty"`
	Vaps                      *[]WirelessControllerWtpProfileRadio2Vaps    `json:"vaps,omitempty"`
	WidsProfile               *string                                      `json:"wids-profile,omitempty"`
	ZeroWaitDfs               *string                                      `json:"zero-wait-dfs,omitempty"`
}

type WirelessControllerWtpProfileRadio2Channel struct {
	Chan *string `json:"chan,omitempty"`
}

type WirelessControllerWtpProfileRadio2Vaps struct {
	Name *string `json:"name,omitempty"`
}

type WirelessControllerWtpProfileRadio3 struct {
	AirtimeFairness           *string                                      `json:"airtime-fairness,omitempty"`
	Amsdu                     *string                                      `json:"amsdu,omitempty"`
	ApHandoff                 *string                                      `json:"ap-handoff,omitempty"`
	ApSnifferAddr             *string                                      `json:"ap-sniffer-addr,omitempty"`
	ApSnifferBufsize          *float64                                     `json:"ap-sniffer-bufsize,omitempty"`
	ApSnifferChan             *float64                                     `json:"ap-sniffer-chan,omitempty"`
	ApSnifferCtl              *string                                      `json:"ap-sniffer-ctl,omitempty"`
	ApSnifferData             *string                                      `json:"ap-sniffer-data,omitempty"`
	ApSnifferMgmtBeacon       *string                                      `json:"ap-sniffer-mgmt-beacon,omitempty"`
	ApSnifferMgmtOther        *string                                      `json:"ap-sniffer-mgmt-other,omitempty"`
	ApSnifferMgmtProbe        *string                                      `json:"ap-sniffer-mgmt-probe,omitempty"`
	AutoPowerHigh             *float64                                     `json:"auto-power-high,omitempty"`
	AutoPowerLevel            *string                                      `json:"auto-power-level,omitempty"`
	AutoPowerLow              *float64                                     `json:"auto-power-low,omitempty"`
	AutoPowerTarget           *string                                      `json:"auto-power-target,omitempty"`
	Band                      *string                                      `json:"band,omitempty"`
	Band5gType                *string                                      `json:"band-5g-type,omitempty"`
	BandwidthAdmissionControl *string                                      `json:"bandwidth-admission-control,omitempty"`
	BandwidthCapacity         *float64                                     `json:"bandwidth-capacity,omitempty"`
	BeaconInterval            *float64                                     `json:"beacon-interval,omitempty"`
	BssColor                  *float64                                     `json:"bss-color,omitempty"`
	BssColorMode              *string                                      `json:"bss-color-mode,omitempty"`
	CallAdmissionControl      *string                                      `json:"call-admission-control,omitempty"`
	CallCapacity              *float64                                     `json:"call-capacity,omitempty"`
	Channel                   *[]WirelessControllerWtpProfileRadio3Channel `json:"channel,omitempty"`
	ChannelBonding            *string                                      `json:"channel-bonding,omitempty"`
	ChannelUtilization        *string                                      `json:"channel-utilization,omitempty"`
	Coexistence               *string                                      `json:"coexistence,omitempty"`
	Darrp                     *string                                      `json:"darrp,omitempty"`
	Drma                      *string                                      `json:"drma,omitempty"`
	DrmaSensitivity           *string                                      `json:"drma-sensitivity,omitempty"`
	Dtim                      *float64                                     `json:"dtim,omitempty"`
	FragThreshold             *float64                                     `json:"frag-threshold,omitempty"`
	FrequencyHandoff          *string                                      `json:"frequency-handoff,omitempty"`
	IperfProtocol             *string                                      `json:"iperf-protocol,omitempty"`
	IperfServerPort           *float64                                     `json:"iperf-server-port,omitempty"`
	MaxClients                *float64                                     `json:"max-clients,omitempty"`
	MaxDistance               *float64                                     `json:"max-distance,omitempty"`
	Mode                      *string                                      `json:"mode,omitempty"`
	PowerLevel                *float64                                     `json:"power-level,omitempty"`
	PowerMode                 *string                                      `json:"power-mode,omitempty"`
	PowerValue                *float64                                     `json:"power-value,omitempty"`
	PowersaveOptimize         *string                                      `json:"powersave-optimize,omitempty"`
	ProtectionMode            *string                                      `json:"protection-mode,omitempty"`
	RtsThreshold              *float64                                     `json:"rts-threshold,omitempty"`
	SamBssid                  *string                                      `json:"sam-bssid,omitempty"`
	SamCaptivePortal          *string                                      `json:"sam-captive-portal,omitempty"`
	SamCwpFailureString       *string                                      `json:"sam-cwp-failure-string,omitempty"`
	SamCwpMatchString         *string                                      `json:"sam-cwp-match-string,omitempty"`
	SamCwpPassword            *string                                      `json:"sam-cwp-password,omitempty"`
	SamCwpSuccessString       *string                                      `json:"sam-cwp-success-string,omitempty"`
	SamCwpTestUrl             *string                                      `json:"sam-cwp-test-url,omitempty"`
	SamCwpUsername            *string                                      `json:"sam-cwp-username,omitempty"`
	SamPassword               *string                                      `json:"sam-password,omitempty"`
	SamReportIntv             *float64                                     `json:"sam-report-intv,omitempty"`
	SamSecurityType           *string                                      `json:"sam-security-type,omitempty"`
	SamServer                 *string                                      `json:"sam-server,omitempty"`
	SamServerFqdn             *string                                      `json:"sam-server-fqdn,omitempty"`
	SamServerIp               *string                                      `json:"sam-server-ip,omitempty"`
	SamServerType             *string                                      `json:"sam-server-type,omitempty"`
	SamSsid                   *string                                      `json:"sam-ssid,omitempty"`
	SamTest                   *string                                      `json:"sam-test,omitempty"`
	SamUsername               *string                                      `json:"sam-username,omitempty"`
	ShortGuardInterval        *string                                      `json:"short-guard-interval,omitempty"`
	SpectrumAnalysis          *string                                      `json:"spectrum-analysis,omitempty"`
	TransmitOptimize          *string                                      `json:"transmit-optimize,omitempty"`
	VapAll                    *string                                      `json:"vap-all,omitempty"`
	Vaps                      *[]WirelessControllerWtpProfileRadio3Vaps    `json:"vaps,omitempty"`
	WidsProfile               *string                                      `json:"wids-profile,omitempty"`
	ZeroWaitDfs               *string                                      `json:"zero-wait-dfs,omitempty"`
}

type WirelessControllerWtpProfileRadio3Channel struct {
	Chan *string `json:"chan,omitempty"`
}

type WirelessControllerWtpProfileRadio3Vaps struct {
	Name *string `json:"name,omitempty"`
}

type WirelessControllerWtpProfileRadio4 struct {
	AirtimeFairness           *string                                      `json:"airtime-fairness,omitempty"`
	Amsdu                     *string                                      `json:"amsdu,omitempty"`
	ApHandoff                 *string                                      `json:"ap-handoff,omitempty"`
	ApSnifferAddr             *string                                      `json:"ap-sniffer-addr,omitempty"`
	ApSnifferBufsize          *float64                                     `json:"ap-sniffer-bufsize,omitempty"`
	ApSnifferChan             *float64                                     `json:"ap-sniffer-chan,omitempty"`
	ApSnifferCtl              *string                                      `json:"ap-sniffer-ctl,omitempty"`
	ApSnifferData             *string                                      `json:"ap-sniffer-data,omitempty"`
	ApSnifferMgmtBeacon       *string                                      `json:"ap-sniffer-mgmt-beacon,omitempty"`
	ApSnifferMgmtOther        *string                                      `json:"ap-sniffer-mgmt-other,omitempty"`
	ApSnifferMgmtProbe        *string                                      `json:"ap-sniffer-mgmt-probe,omitempty"`
	AutoPowerHigh             *float64                                     `json:"auto-power-high,omitempty"`
	AutoPowerLevel            *string                                      `json:"auto-power-level,omitempty"`
	AutoPowerLow              *float64                                     `json:"auto-power-low,omitempty"`
	AutoPowerTarget           *string                                      `json:"auto-power-target,omitempty"`
	Band                      *string                                      `json:"band,omitempty"`
	Band5gType                *string                                      `json:"band-5g-type,omitempty"`
	BandwidthAdmissionControl *string                                      `json:"bandwidth-admission-control,omitempty"`
	BandwidthCapacity         *float64                                     `json:"bandwidth-capacity,omitempty"`
	BeaconInterval            *float64                                     `json:"beacon-interval,omitempty"`
	BssColor                  *float64                                     `json:"bss-color,omitempty"`
	BssColorMode              *string                                      `json:"bss-color-mode,omitempty"`
	CallAdmissionControl      *string                                      `json:"call-admission-control,omitempty"`
	CallCapacity              *float64                                     `json:"call-capacity,omitempty"`
	Channel                   *[]WirelessControllerWtpProfileRadio4Channel `json:"channel,omitempty"`
	ChannelBonding            *string                                      `json:"channel-bonding,omitempty"`
	ChannelUtilization        *string                                      `json:"channel-utilization,omitempty"`
	Coexistence               *string                                      `json:"coexistence,omitempty"`
	Darrp                     *string                                      `json:"darrp,omitempty"`
	Drma                      *string                                      `json:"drma,omitempty"`
	DrmaSensitivity           *string                                      `json:"drma-sensitivity,omitempty"`
	Dtim                      *float64                                     `json:"dtim,omitempty"`
	FragThreshold             *float64                                     `json:"frag-threshold,omitempty"`
	FrequencyHandoff          *string                                      `json:"frequency-handoff,omitempty"`
	IperfProtocol             *string                                      `json:"iperf-protocol,omitempty"`
	IperfServerPort           *float64                                     `json:"iperf-server-port,omitempty"`
	MaxClients                *float64                                     `json:"max-clients,omitempty"`
	MaxDistance               *float64                                     `json:"max-distance,omitempty"`
	Mode                      *string                                      `json:"mode,omitempty"`
	PowerLevel                *float64                                     `json:"power-level,omitempty"`
	PowerMode                 *string                                      `json:"power-mode,omitempty"`
	PowerValue                *float64                                     `json:"power-value,omitempty"`
	PowersaveOptimize         *string                                      `json:"powersave-optimize,omitempty"`
	ProtectionMode            *string                                      `json:"protection-mode,omitempty"`
	RtsThreshold              *float64                                     `json:"rts-threshold,omitempty"`
	SamBssid                  *string                                      `json:"sam-bssid,omitempty"`
	SamCaptivePortal          *string                                      `json:"sam-captive-portal,omitempty"`
	SamCwpFailureString       *string                                      `json:"sam-cwp-failure-string,omitempty"`
	SamCwpMatchString         *string                                      `json:"sam-cwp-match-string,omitempty"`
	SamCwpPassword            *string                                      `json:"sam-cwp-password,omitempty"`
	SamCwpSuccessString       *string                                      `json:"sam-cwp-success-string,omitempty"`
	SamCwpTestUrl             *string                                      `json:"sam-cwp-test-url,omitempty"`
	SamCwpUsername            *string                                      `json:"sam-cwp-username,omitempty"`
	SamPassword               *string                                      `json:"sam-password,omitempty"`
	SamReportIntv             *float64                                     `json:"sam-report-intv,omitempty"`
	SamSecurityType           *string                                      `json:"sam-security-type,omitempty"`
	SamServer                 *string                                      `json:"sam-server,omitempty"`
	SamServerFqdn             *string                                      `json:"sam-server-fqdn,omitempty"`
	SamServerIp               *string                                      `json:"sam-server-ip,omitempty"`
	SamServerType             *string                                      `json:"sam-server-type,omitempty"`
	SamSsid                   *string                                      `json:"sam-ssid,omitempty"`
	SamTest                   *string                                      `json:"sam-test,omitempty"`
	SamUsername               *string                                      `json:"sam-username,omitempty"`
	ShortGuardInterval        *string                                      `json:"short-guard-interval,omitempty"`
	SpectrumAnalysis          *string                                      `json:"spectrum-analysis,omitempty"`
	TransmitOptimize          *string                                      `json:"transmit-optimize,omitempty"`
	VapAll                    *string                                      `json:"vap-all,omitempty"`
	Vaps                      *[]WirelessControllerWtpProfileRadio4Vaps    `json:"vaps,omitempty"`
	WidsProfile               *string                                      `json:"wids-profile,omitempty"`
	ZeroWaitDfs               *string                                      `json:"zero-wait-dfs,omitempty"`
}

type WirelessControllerWtpProfileRadio4Channel struct {
	Chan *string `json:"chan,omitempty"`
}

type WirelessControllerWtpProfileRadio4Vaps struct {
	Name *string `json:"name,omitempty"`
}

type WirelessControllerWtpProfileSplitTunnelingAcl struct {
	DestIp *string  `json:"dest-ip,omitempty"`
	Id     *float64 `json:"id,omitempty"`
}
