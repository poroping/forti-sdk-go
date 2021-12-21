package models

const WirelessControllerVapPath = "wireless-controller/vap/"

type WirelessControllerVap struct {
	AccessControlList                   *string                                         `json:"access-control-list,omitempty"`
	AcctInterimInterval                 *float64                                        `json:"acct-interim-interval,omitempty"`
	AdditionalAkms                      *string                                         `json:"additional-akms,omitempty"`
	AddressGroup                        *string                                         `json:"address-group,omitempty"`
	AntivirusProfile                    *string                                         `json:"antivirus-profile,omitempty"`
	ApplicationList                     *string                                         `json:"application-list,omitempty"`
	AtfWeight                           *float64                                        `json:"atf-weight,omitempty"`
	Auth                                *string                                         `json:"auth,omitempty"`
	BeaconAdvertising                   *string                                         `json:"beacon-advertising,omitempty"`
	BroadcastSsid                       *string                                         `json:"broadcast-ssid,omitempty"`
	BroadcastSuppression                *string                                         `json:"broadcast-suppression,omitempty"`
	BssColorPartial                     *string                                         `json:"bss-color-partial,omitempty"`
	BstmDisassociationImminent          *string                                         `json:"bstm-disassociation-imminent,omitempty"`
	BstmLoadBalancingDisassocTimer      *float64                                        `json:"bstm-load-balancing-disassoc-timer,omitempty"`
	BstmRssiDisassocTimer               *float64                                        `json:"bstm-rssi-disassoc-timer,omitempty"`
	CaptivePortalAcName                 *string                                         `json:"captive-portal-ac-name,omitempty"`
	CaptivePortalAuthTimeout            *float64                                        `json:"captive-portal-auth-timeout,omitempty"`
	CaptivePortalMacauthRadiusSecret    *string                                         `json:"captive-portal-macauth-radius-secret,omitempty"`
	CaptivePortalMacauthRadiusServer    *string                                         `json:"captive-portal-macauth-radius-server,omitempty"`
	CaptivePortalRadiusSecret           *string                                         `json:"captive-portal-radius-secret,omitempty"`
	CaptivePortalRadiusServer           *string                                         `json:"captive-portal-radius-server,omitempty"`
	CaptivePortalSessionTimeoutInterval *float64                                        `json:"captive-portal-session-timeout-interval,omitempty"`
	DhcpAddressEnforcement              *string                                         `json:"dhcp-address-enforcement,omitempty"`
	DhcpLeaseTime                       *float64                                        `json:"dhcp-lease-time,omitempty"`
	DhcpOption43Insertion               *string                                         `json:"dhcp-option43-insertion,omitempty"`
	DhcpOption82CircuitIdInsertion      *string                                         `json:"dhcp-option82-circuit-id-insertion,omitempty"`
	DhcpOption82Insertion               *string                                         `json:"dhcp-option82-insertion,omitempty"`
	DhcpOption82RemoteIdInsertion       *string                                         `json:"dhcp-option82-remote-id-insertion,omitempty"`
	DynamicVlan                         *string                                         `json:"dynamic-vlan,omitempty"`
	EapReauth                           *string                                         `json:"eap-reauth,omitempty"`
	EapReauthIntv                       *float64                                        `json:"eap-reauth-intv,omitempty"`
	EapolKeyRetries                     *string                                         `json:"eapol-key-retries,omitempty"`
	Encrypt                             *string                                         `json:"encrypt,omitempty"`
	ExternalFastRoaming                 *string                                         `json:"external-fast-roaming,omitempty"`
	ExternalLogout                      *string                                         `json:"external-logout,omitempty"`
	ExternalWeb                         *string                                         `json:"external-web,omitempty"`
	ExternalWebFormat                   *string                                         `json:"external-web-format,omitempty"`
	FastBssTransition                   *string                                         `json:"fast-bss-transition,omitempty"`
	FastRoaming                         *string                                         `json:"fast-roaming,omitempty"`
	FtMobilityDomain                    *float64                                        `json:"ft-mobility-domain,omitempty"`
	FtOverDs                            *string                                         `json:"ft-over-ds,omitempty"`
	FtR0KeyLifetime                     *float64                                        `json:"ft-r0-key-lifetime,omitempty"`
	GasComebackDelay                    *float64                                        `json:"gas-comeback-delay,omitempty"`
	GasFragmentationLimit               *float64                                        `json:"gas-fragmentation-limit,omitempty"`
	GtkRekey                            *string                                         `json:"gtk-rekey,omitempty"`
	GtkRekeyIntv                        *float64                                        `json:"gtk-rekey-intv,omitempty"`
	HighEfficiency                      *string                                         `json:"high-efficiency,omitempty"`
	Hotspot20Profile                    *string                                         `json:"hotspot20-profile,omitempty"`
	IgmpSnooping                        *string                                         `json:"igmp-snooping,omitempty"`
	IntraVapPrivacy                     *string                                         `json:"intra-vap-privacy,omitempty"`
	Ip                                  *string                                         `json:"ip,omitempty"`
	IpsSensor                           *string                                         `json:"ips-sensor,omitempty"`
	Ipv6Rules                           *string                                         `json:"ipv6-rules,omitempty"`
	Key                                 *string                                         `json:"key,omitempty"`
	Keyindex                            *float64                                        `json:"keyindex,omitempty"`
	Ldpc                                *string                                         `json:"ldpc,omitempty"`
	LocalAuthentication                 *string                                         `json:"local-authentication,omitempty"`
	LocalBridging                       *string                                         `json:"local-bridging,omitempty"`
	LocalLan                            *string                                         `json:"local-lan,omitempty"`
	LocalStandalone                     *string                                         `json:"local-standalone,omitempty"`
	LocalStandaloneDns                  *string                                         `json:"local-standalone-dns,omitempty"`
	LocalStandaloneDnsIp                *string                                         `json:"local-standalone-dns-ip,omitempty"`
	LocalStandaloneNat                  *string                                         `json:"local-standalone-nat,omitempty"`
	MacAuthBypass                       *string                                         `json:"mac-auth-bypass,omitempty"`
	MacCalledStationDelimiter           *string                                         `json:"mac-called-station-delimiter,omitempty"`
	MacCallingStationDelimiter          *string                                         `json:"mac-calling-station-delimiter,omitempty"`
	MacCase                             *string                                         `json:"mac-case,omitempty"`
	MacFilter                           *string                                         `json:"mac-filter,omitempty"`
	MacFilterList                       *[]WirelessControllerVapMacFilterList           `json:"mac-filter-list,omitempty"`
	MacFilterPolicyOther                *string                                         `json:"mac-filter-policy-other,omitempty"`
	MacPasswordDelimiter                *string                                         `json:"mac-password-delimiter,omitempty"`
	MacUsernameDelimiter                *string                                         `json:"mac-username-delimiter,omitempty"`
	MaxClients                          *float64                                        `json:"max-clients,omitempty"`
	MaxClientsAp                        *float64                                        `json:"max-clients-ap,omitempty"`
	Mbo                                 *string                                         `json:"mbo,omitempty"`
	MboCellDataConnPref                 *string                                         `json:"mbo-cell-data-conn-pref,omitempty"`
	MeDisableThresh                     *float64                                        `json:"me-disable-thresh,omitempty"`
	MeshBackhaul                        *string                                         `json:"mesh-backhaul,omitempty"`
	Mpsk                                *string                                         `json:"mpsk,omitempty"`
	MpskConcurrentClients               *float64                                        `json:"mpsk-concurrent-clients,omitempty"`
	MpskKey                             *[]WirelessControllerVapMpskKey                 `json:"mpsk-key,omitempty"`
	MpskProfile                         *string                                         `json:"mpsk-profile,omitempty"`
	MuMimo                              *string                                         `json:"mu-mimo,omitempty"`
	MulticastEnhance                    *string                                         `json:"multicast-enhance,omitempty"`
	MulticastRate                       *string                                         `json:"multicast-rate,omitempty"`
	Nac                                 *string                                         `json:"nac,omitempty"`
	NacProfile                          *string                                         `json:"nac-profile,omitempty"`
	Name                                *string                                         `json:"name,omitempty"`
	NeighborReportDualBand              *string                                         `json:"neighbor-report-dual-band,omitempty"`
	Okc                                 *string                                         `json:"okc,omitempty"`
	Osen                                *string                                         `json:"osen,omitempty"`
	OweGroups                           *string                                         `json:"owe-groups,omitempty"`
	OweTransition                       *string                                         `json:"owe-transition,omitempty"`
	OweTransitionSsid                   *string                                         `json:"owe-transition-ssid,omitempty"`
	Passphrase                          *string                                         `json:"passphrase,omitempty"`
	Pmf                                 *string                                         `json:"pmf,omitempty"`
	PmfAssocComebackTimeout             *float64                                        `json:"pmf-assoc-comeback-timeout,omitempty"`
	PmfSaQueryRetryTimeout              *float64                                        `json:"pmf-sa-query-retry-timeout,omitempty"`
	PortMacauth                         *string                                         `json:"port-macauth,omitempty"`
	PortMacauthReauthTimeout            *float64                                        `json:"port-macauth-reauth-timeout,omitempty"`
	PortMacauthTimeout                  *float64                                        `json:"port-macauth-timeout,omitempty"`
	PortalMessageOverrideGroup          *string                                         `json:"portal-message-override-group,omitempty"`
	PortalMessageOverrides              *[]WirelessControllerVapPortalMessageOverrides  `json:"portal-message-overrides,omitempty"`
	PortalType                          *string                                         `json:"portal-type,omitempty"`
	PrimaryWagProfile                   *string                                         `json:"primary-wag-profile,omitempty"`
	ProbeRespSuppression                *string                                         `json:"probe-resp-suppression,omitempty"`
	ProbeRespThreshold                  *string                                         `json:"probe-resp-threshold,omitempty"`
	PtkRekey                            *string                                         `json:"ptk-rekey,omitempty"`
	PtkRekeyIntv                        *float64                                        `json:"ptk-rekey-intv,omitempty"`
	QosProfile                          *string                                         `json:"qos-profile,omitempty"`
	Quarantine                          *string                                         `json:"quarantine,omitempty"`
	Radio2gThreshold                    *string                                         `json:"radio-2g-threshold,omitempty"`
	Radio5gThreshold                    *string                                         `json:"radio-5g-threshold,omitempty"`
	RadioSensitivity                    *string                                         `json:"radio-sensitivity,omitempty"`
	RadiusMacAuth                       *string                                         `json:"radius-mac-auth,omitempty"`
	RadiusMacAuthServer                 *string                                         `json:"radius-mac-auth-server,omitempty"`
	RadiusMacAuthUsergroups             *[]WirelessControllerVapRadiusMacAuthUsergroups `json:"radius-mac-auth-usergroups,omitempty"`
	RadiusMacMpskAuth                   *string                                         `json:"radius-mac-mpsk-auth,omitempty"`
	RadiusMacMpskTimeout                *float64                                        `json:"radius-mac-mpsk-timeout,omitempty"`
	RadiusServer                        *string                                         `json:"radius-server,omitempty"`
	Rates11a                            *string                                         `json:"rates-11a,omitempty"`
	Rates11acSs12                       *string                                         `json:"rates-11ac-ss12,omitempty"`
	Rates11acSs34                       *string                                         `json:"rates-11ac-ss34,omitempty"`
	Rates11axSs12                       *string                                         `json:"rates-11ax-ss12,omitempty"`
	Rates11axSs34                       *string                                         `json:"rates-11ax-ss34,omitempty"`
	Rates11bg                           *string                                         `json:"rates-11bg,omitempty"`
	Rates11nSs12                        *string                                         `json:"rates-11n-ss12,omitempty"`
	Rates11nSs34                        *string                                         `json:"rates-11n-ss34,omitempty"`
	SaeGroups                           *string                                         `json:"sae-groups,omitempty"`
	SaePassword                         *string                                         `json:"sae-password,omitempty"`
	ScanBotnetConnections               *string                                         `json:"scan-botnet-connections,omitempty"`
	Schedule                            *[]WirelessControllerVapSchedule                `json:"schedule,omitempty"`
	SecondaryWagProfile                 *string                                         `json:"secondary-wag-profile,omitempty"`
	Security                            *string                                         `json:"security,omitempty"`
	SecurityExemptList                  *string                                         `json:"security-exempt-list,omitempty"`
	SecurityRedirectUrl                 *string                                         `json:"security-redirect-url,omitempty"`
	SelectedUsergroups                  *[]WirelessControllerVapSelectedUsergroups      `json:"selected-usergroups,omitempty"`
	SplitTunneling                      *string                                         `json:"split-tunneling,omitempty"`
	Ssid                                *string                                         `json:"ssid,omitempty"`
	StickyClientRemove                  *string                                         `json:"sticky-client-remove,omitempty"`
	StickyClientThreshold2g             *string                                         `json:"sticky-client-threshold-2g,omitempty"`
	StickyClientThreshold5g             *string                                         `json:"sticky-client-threshold-5g,omitempty"`
	TargetWakeTime                      *string                                         `json:"target-wake-time,omitempty"`
	TkipCounterMeasure                  *string                                         `json:"tkip-counter-measure,omitempty"`
	TunnelEchoInterval                  *float64                                        `json:"tunnel-echo-interval,omitempty"`
	TunnelFallbackInterval              *float64                                        `json:"tunnel-fallback-interval,omitempty"`
	Usergroup                           *[]WirelessControllerVapUsergroup               `json:"usergroup,omitempty"`
	UtmLog                              *string                                         `json:"utm-log,omitempty"`
	UtmProfile                          *string                                         `json:"utm-profile,omitempty"`
	UtmStatus                           *string                                         `json:"utm-status,omitempty"`
	VlanAuto                            *string                                         `json:"vlan-auto,omitempty"`
	VlanPool                            *[]WirelessControllerVapVlanPool                `json:"vlan-pool,omitempty"`
	VlanPooling                         *string                                         `json:"vlan-pooling,omitempty"`
	Vlanid                              *float64                                        `json:"vlanid,omitempty"`
	VoiceEnterprise                     *string                                         `json:"voice-enterprise,omitempty"`
	WebfilterProfile                    *string                                         `json:"webfilter-profile,omitempty"`
}

type WirelessControllerVapMacFilterList struct {
	Id              *float64 `json:"id,omitempty"`
	Mac             *string  `json:"mac,omitempty"`
	MacFilterPolicy *string  `json:"mac-filter-policy,omitempty"`
}

type WirelessControllerVapMpskKey struct {
	Comment           *string                                      `json:"comment,omitempty"`
	ConcurrentClients *string                                      `json:"concurrent-clients,omitempty"`
	KeyName           *string                                      `json:"key-name,omitempty"`
	MpskSchedules     *[]WirelessControllerVapMpskKeyMpskSchedules `json:"mpsk-schedules,omitempty"`
	Passphrase        *string                                      `json:"passphrase,omitempty"`
}

type WirelessControllerVapMpskKeyMpskSchedules struct {
	Name *string `json:"name,omitempty"`
}

type WirelessControllerVapPortalMessageOverrides struct {
	AuthDisclaimerPage  *string `json:"auth-disclaimer-page,omitempty"`
	AuthLoginFailedPage *string `json:"auth-login-failed-page,omitempty"`
	AuthLoginPage       *string `json:"auth-login-page,omitempty"`
	AuthRejectPage      *string `json:"auth-reject-page,omitempty"`
}

type WirelessControllerVapRadiusMacAuthUsergroups struct {
	Name *string `json:"name,omitempty"`
}

type WirelessControllerVapSchedule struct {
	Name *string `json:"name,omitempty"`
}

type WirelessControllerVapSelectedUsergroups struct {
	Name *string `json:"name,omitempty"`
}

type WirelessControllerVapUsergroup struct {
	Name *string `json:"name,omitempty"`
}

type WirelessControllerVapVlanPool struct {
	Id       *float64 `json:"id,omitempty"`
	WtpGroup *string  `json:"wtp-group,omitempty"`
}
