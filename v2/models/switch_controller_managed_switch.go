package models

const SwitchControllerManagedSwitchPath = "switch-controller/managed-switch/"

type SwitchControllerManagedSwitch struct {
	The8021XSettings          *SwitchControllerManagedSwitch8021XSettings     `json:"802-1X-settings,omitempty"`
	AccessProfile             *string                                         `json:"access-profile,omitempty"`
	CustomCommand             *[]SwitchControllerManagedSwitchCustomCommand   `json:"custom-command,omitempty"`
	DelayedRestartTrigger     *int64                                          `json:"delayed-restart-trigger,omitempty"`
	Description               *string                                         `json:"description,omitempty"`
	DhcpServerAccessList      *string                                         `json:"dhcp-server-access-list,omitempty"`
	DirectlyConnected         *int64                                          `json:"directly-connected,omitempty"`
	DynamicCapability         *string                                         `json:"dynamic-capability,omitempty"`
	DynamicallyDiscovered     *int64                                          `json:"dynamically-discovered,omitempty"`
	FirmwareProvision         *string                                         `json:"firmware-provision,omitempty"`
	FirmwareProvisionLatest   *string                                         `json:"firmware-provision-latest,omitempty"`
	FirmwareProvisionVersion  *string                                         `json:"firmware-provision-version,omitempty"`
	FlowIdentity              *string                                         `json:"flow-identity,omitempty"`
	FswWan1Admin              *string                                         `json:"fsw-wan1-admin,omitempty"`
	FswWan1Peer               *string                                         `json:"fsw-wan1-peer,omitempty"`
	IgmpSnooping              *SwitchControllerManagedSwitchIgmpSnooping      `json:"igmp-snooping,omitempty"`
	IpSourceGuard             *[]SwitchControllerManagedSwitchIpSourceGuard   `json:"ip-source-guard,omitempty"`
	L3Discovered              *int64                                          `json:"l3-discovered,omitempty"`
	MaxAllowedTrunkMembers    *int64                                          `json:"max-allowed-trunk-members,omitempty"`
	MclagIgmpSnoopingAware    *string                                         `json:"mclag-igmp-snooping-aware,omitempty"`
	Mirror                    *[]SwitchControllerManagedSwitchMirror          `json:"mirror,omitempty"`
	Name                      *string                                         `json:"name,omitempty"`
	OverrideSnmpCommunity     *string                                         `json:"override-snmp-community,omitempty"`
	OverrideSnmpSysinfo       *string                                         `json:"override-snmp-sysinfo,omitempty"`
	OverrideSnmpTrapThreshold *string                                         `json:"override-snmp-trap-threshold,omitempty"`
	OverrideSnmpUser          *string                                         `json:"override-snmp-user,omitempty"`
	OwnerVdom                 *string                                         `json:"owner-vdom,omitempty"`
	PoeDetectionType          *int64                                          `json:"poe-detection-type,omitempty"`
	PoeLldpDetection          *string                                         `json:"poe-lldp-detection,omitempty"`
	PoePreStandardDetection   *string                                         `json:"poe-pre-standard-detection,omitempty"`
	Ports                     *[]SwitchControllerManagedSwitchPorts           `json:"ports,omitempty"`
	PreProvisioned            *int64                                          `json:"pre-provisioned,omitempty"`
	QosDropPolicy             *string                                         `json:"qos-drop-policy,omitempty"`
	QosRedProbability         *int64                                          `json:"qos-red-probability,omitempty"`
	RemoteLog                 *[]SwitchControllerManagedSwitchRemoteLog       `json:"remote-log,omitempty"`
	SnmpCommunity             *[]SwitchControllerManagedSwitchSnmpCommunity   `json:"snmp-community,omitempty"`
	SnmpSysinfo               *SwitchControllerManagedSwitchSnmpSysinfo       `json:"snmp-sysinfo,omitempty"`
	SnmpTrapThreshold         *SwitchControllerManagedSwitchSnmpTrapThreshold `json:"snmp-trap-threshold,omitempty"`
	SnmpUser                  *[]SwitchControllerManagedSwitchSnmpUser        `json:"snmp-user,omitempty"`
	StagedImageVersion        *string                                         `json:"staged-image-version,omitempty"`
	StaticMac                 *[]SwitchControllerManagedSwitchStaticMac       `json:"static-mac,omitempty"`
	StormControl              *SwitchControllerManagedSwitchStormControl      `json:"storm-control,omitempty"`
	StpInstance               *[]SwitchControllerManagedSwitchStpInstance     `json:"stp-instance,omitempty"`
	StpSettings               *SwitchControllerManagedSwitchStpSettings       `json:"stp-settings,omitempty"`
	SwitchDeviceTag           *string                                         `json:"switch-device-tag,omitempty"`
	SwitchDhcp_opt43_key      *string                                         `json:"switch-dhcp_opt43_key,omitempty"`
	SwitchId                  *string                                         `json:"switch-id,omitempty"`
	SwitchLog                 *SwitchControllerManagedSwitchSwitchLog         `json:"switch-log,omitempty"`
	SwitchProfile             *string                                         `json:"switch-profile,omitempty"`
	TdrSupported              *string                                         `json:"tdr-supported,omitempty"`
	Type                      *string                                         `json:"type,omitempty"`
	Version                   *int64                                          `json:"version,omitempty"`
}

type SwitchControllerManagedSwitch8021XSettings struct {
	LinkDownAuth     *string `json:"link-down-auth,omitempty"`
	LocalOverride    *string `json:"local-override,omitempty"`
	MaxReauthAttempt *int64  `json:"max-reauth-attempt,omitempty"`
	ReauthPeriod     *int64  `json:"reauth-period,omitempty"`
	TxPeriod         *int64  `json:"tx-period,omitempty"`
}

type SwitchControllerManagedSwitchCustomCommand struct {
	CommandEntry *string `json:"command-entry,omitempty"`
	CommandName  *string `json:"command-name,omitempty"`
}

type SwitchControllerManagedSwitchIgmpSnooping struct {
	AgingTime             *int64                                            `json:"aging-time,omitempty"`
	FloodUnknownMulticast *string                                           `json:"flood-unknown-multicast,omitempty"`
	LocalOverride         *string                                           `json:"local-override,omitempty"`
	Vlans                 *[]SwitchControllerManagedSwitchIgmpSnoopingVlans `json:"vlans,omitempty"`
}

type SwitchControllerManagedSwitchIgmpSnoopingVlans struct {
	Proxy       *string `json:"proxy,omitempty"`
	Querier     *string `json:"querier,omitempty"`
	QuerierAddr *string `json:"querier-addr,omitempty"`
	Version     *int64  `json:"version,omitempty"`
	VlanName    *string `json:"vlan-name,omitempty"`
}

type SwitchControllerManagedSwitchIpSourceGuard struct {
	BindingEntry *[]SwitchControllerManagedSwitchIpSourceGuardBindingEntry `json:"binding-entry,omitempty"`
	Description  *string                                                   `json:"description,omitempty"`
	Port         *string                                                   `json:"port,omitempty"`
}

type SwitchControllerManagedSwitchIpSourceGuardBindingEntry struct {
	EntryName *string `json:"entry-name,omitempty"`
	Ip        *string `json:"ip,omitempty"`
	Mac       *string `json:"mac,omitempty"`
}

type SwitchControllerManagedSwitchMirror struct {
	Dst             *string                                          `json:"dst,omitempty"`
	Name            *string                                          `json:"name,omitempty"`
	SrcEgress       *[]SwitchControllerManagedSwitchMirrorSrcEgress  `json:"src-egress,omitempty"`
	SrcIngress      *[]SwitchControllerManagedSwitchMirrorSrcIngress `json:"src-ingress,omitempty"`
	Status          *string                                          `json:"status,omitempty"`
	SwitchingPacket *string                                          `json:"switching-packet,omitempty"`
}

type SwitchControllerManagedSwitchMirrorSrcEgress struct {
	Name *string `json:"name,omitempty"`
}

type SwitchControllerManagedSwitchMirrorSrcIngress struct {
	Name *string `json:"name,omitempty"`
}

type SwitchControllerManagedSwitchPorts struct {
	AccessMode               *string                                            `json:"access-mode,omitempty"`
	AggregatorMode           *string                                            `json:"aggregator-mode,omitempty"`
	AllowedVlans             *[]SwitchControllerManagedSwitchPortsAllowedVlans  `json:"allowed-vlans,omitempty"`
	AllowedVlansAll          *string                                            `json:"allowed-vlans-all,omitempty"`
	ArpInspectionTrust       *string                                            `json:"arp-inspection-trust,omitempty"`
	Bundle                   *string                                            `json:"bundle,omitempty"`
	Description              *string                                            `json:"description,omitempty"`
	DhcpSnoopOption82Trust   *string                                            `json:"dhcp-snoop-option82-trust,omitempty"`
	DhcpSnooping             *string                                            `json:"dhcp-snooping,omitempty"`
	DiscardMode              *string                                            `json:"discard-mode,omitempty"`
	EdgePort                 *string                                            `json:"edge-port,omitempty"`
	ExportTags               *[]SwitchControllerManagedSwitchPortsExportTags    `json:"export-tags,omitempty"`
	ExportTo                 *string                                            `json:"export-to,omitempty"`
	ExportToPool             *string                                            `json:"export-to-pool,omitempty"`
	FecCapable               *int64                                             `json:"fec-capable,omitempty"`
	FecState                 *string                                            `json:"fec-state,omitempty"`
	FgtPeerDeviceName        *string                                            `json:"fgt-peer-device-name,omitempty"`
	FgtPeerPortName          *string                                            `json:"fgt-peer-port-name,omitempty"`
	FiberPort                *int64                                             `json:"fiber-port,omitempty"`
	Flags                    *int64                                             `json:"flags,omitempty"`
	FlowControl              *string                                            `json:"flow-control,omitempty"`
	FortilinkPort            *int64                                             `json:"fortilink-port,omitempty"`
	IgmpSnooping             *string                                            `json:"igmp-snooping,omitempty"`
	IgmpsFloodReports        *string                                            `json:"igmps-flood-reports,omitempty"`
	IgmpsFloodTraffic        *string                                            `json:"igmps-flood-traffic,omitempty"`
	InterfaceTags            *[]SwitchControllerManagedSwitchPortsInterfaceTags `json:"interface-tags,omitempty"`
	IpSourceGuard            *string                                            `json:"ip-source-guard,omitempty"`
	IslLocalTrunkName        *string                                            `json:"isl-local-trunk-name,omitempty"`
	IslPeerDeviceName        *string                                            `json:"isl-peer-device-name,omitempty"`
	IslPeerPortName          *string                                            `json:"isl-peer-port-name,omitempty"`
	LacpSpeed                *string                                            `json:"lacp-speed,omitempty"`
	LearningLimit            *int64                                             `json:"learning-limit,omitempty"`
	LldpProfile              *string                                            `json:"lldp-profile,omitempty"`
	LldpStatus               *string                                            `json:"lldp-status,omitempty"`
	LoopGuard                *string                                            `json:"loop-guard,omitempty"`
	LoopGuardTimeout         *int64                                             `json:"loop-guard-timeout,omitempty"`
	MacAddr                  *string                                            `json:"mac-addr,omitempty"`
	MatchedDppIntfTags       *string                                            `json:"matched-dpp-intf-tags,omitempty"`
	MatchedDppPolicy         *string                                            `json:"matched-dpp-policy,omitempty"`
	MaxBundle                *int64                                             `json:"max-bundle,omitempty"`
	Mclag                    *string                                            `json:"mclag,omitempty"`
	MclagIclPort             *int64                                             `json:"mclag-icl-port,omitempty"`
	MediaType                *string                                            `json:"media-type,omitempty"`
	MemberWithdrawalBehavior *string                                            `json:"member-withdrawal-behavior,omitempty"`
	Members                  *[]SwitchControllerManagedSwitchPortsMembers       `json:"members,omitempty"`
	MinBundle                *int64                                             `json:"min-bundle,omitempty"`
	Mode                     *string                                            `json:"mode,omitempty"`
	P2pPort                  *int64                                             `json:"p2p-port,omitempty"`
	PacketSampleRate         *int64                                             `json:"packet-sample-rate,omitempty"`
	PacketSampler            *string                                            `json:"packet-sampler,omitempty"`
	PauseMeter               *int64                                             `json:"pause-meter,omitempty"`
	PauseMeterResume         *string                                            `json:"pause-meter-resume,omitempty"`
	PoeCapable               *int64                                             `json:"poe-capable,omitempty"`
	PoeMaxPower              *string                                            `json:"poe-max-power,omitempty"`
	PoePreStandardDetection  *string                                            `json:"poe-pre-standard-detection,omitempty"`
	PoeStandard              *string                                            `json:"poe-standard,omitempty"`
	PoeStatus                *string                                            `json:"poe-status,omitempty"`
	PortName                 *string                                            `json:"port-name,omitempty"`
	PortNumber               *int64                                             `json:"port-number,omitempty"`
	PortOwner                *string                                            `json:"port-owner,omitempty"`
	PortPolicy               *string                                            `json:"port-policy,omitempty"`
	PortPrefixType           *int64                                             `json:"port-prefix-type,omitempty"`
	PortSecurityPolicy       *string                                            `json:"port-security-policy,omitempty"`
	PortSelectionCriteria    *string                                            `json:"port-selection-criteria,omitempty"`
	PtpPolicy                *string                                            `json:"ptp-policy,omitempty"`
	QosPolicy                *string                                            `json:"qos-policy,omitempty"`
	RpvstPort                *string                                            `json:"rpvst-port,omitempty"`
	SampleDirection          *string                                            `json:"sample-direction,omitempty"`
	SflowCounterInterval     *int64                                             `json:"sflow-counter-interval,omitempty"`
	Speed                    *string                                            `json:"speed,omitempty"`
	StackingPort             *int64                                             `json:"stacking-port,omitempty"`
	Status                   *string                                            `json:"status,omitempty"`
	StickyMac                *string                                            `json:"sticky-mac,omitempty"`
	StormControlPolicy       *string                                            `json:"storm-control-policy,omitempty"`
	StpBpduGuard             *string                                            `json:"stp-bpdu-guard,omitempty"`
	StpBpduGuardTimeout      *int64                                             `json:"stp-bpdu-guard-timeout,omitempty"`
	StpRootGuard             *string                                            `json:"stp-root-guard,omitempty"`
	StpState                 *string                                            `json:"stp-state,omitempty"`
	SwitchId                 *string                                            `json:"switch-id,omitempty"`
	Type                     *string                                            `json:"type,omitempty"`
	UntaggedVlans            *[]SwitchControllerManagedSwitchPortsUntaggedVlans `json:"untagged-vlans,omitempty"`
	Vlan                     *string                                            `json:"vlan,omitempty"`
}

type SwitchControllerManagedSwitchPortsAllowedVlans struct {
	VlanName *string `json:"vlan-name,omitempty"`
}

type SwitchControllerManagedSwitchPortsExportTags struct {
	TagName *string `json:"tag-name,omitempty"`
}

type SwitchControllerManagedSwitchPortsInterfaceTags struct {
	TagName *string `json:"tag-name,omitempty"`
}

type SwitchControllerManagedSwitchPortsMembers struct {
	MemberName *string `json:"member-name,omitempty"`
}

type SwitchControllerManagedSwitchPortsUntaggedVlans struct {
	VlanName *string `json:"vlan-name,omitempty"`
}

type SwitchControllerManagedSwitchRemoteLog struct {
	Csv      *string `json:"csv,omitempty"`
	Facility *string `json:"facility,omitempty"`
	Name     *string `json:"name,omitempty"`
	Port     *int64  `json:"port,omitempty"`
	Server   *string `json:"server,omitempty"`
	Severity *string `json:"severity,omitempty"`
	Status   *string `json:"status,omitempty"`
}

type SwitchControllerManagedSwitchSnmpCommunity struct {
	Events         *string                                            `json:"events,omitempty"`
	Hosts          *[]SwitchControllerManagedSwitchSnmpCommunityHosts `json:"hosts,omitempty"`
	Id             *int64                                             `json:"id,omitempty"`
	Name           *string                                            `json:"name,omitempty"`
	QueryV1Port    *int64                                             `json:"query-v1-port,omitempty"`
	QueryV1Status  *string                                            `json:"query-v1-status,omitempty"`
	QueryV2cPort   *int64                                             `json:"query-v2c-port,omitempty"`
	QueryV2cStatus *string                                            `json:"query-v2c-status,omitempty"`
	Status         *string                                            `json:"status,omitempty"`
	TrapV1Lport    *int64                                             `json:"trap-v1-lport,omitempty"`
	TrapV1Rport    *int64                                             `json:"trap-v1-rport,omitempty"`
	TrapV1Status   *string                                            `json:"trap-v1-status,omitempty"`
	TrapV2cLport   *int64                                             `json:"trap-v2c-lport,omitempty"`
	TrapV2cRport   *int64                                             `json:"trap-v2c-rport,omitempty"`
	TrapV2cStatus  *string                                            `json:"trap-v2c-status,omitempty"`
}

type SwitchControllerManagedSwitchSnmpCommunityHosts struct {
	Id *int64  `json:"id,omitempty"`
	Ip *string `json:"ip,omitempty"`
}

type SwitchControllerManagedSwitchSnmpSysinfo struct {
	ContactInfo *string `json:"contact-info,omitempty"`
	Description *string `json:"description,omitempty"`
	EngineId    *string `json:"engine-id,omitempty"`
	Location    *string `json:"location,omitempty"`
	Status      *string `json:"status,omitempty"`
}

type SwitchControllerManagedSwitchSnmpTrapThreshold struct {
	TrapHighCpuThreshold   *int64 `json:"trap-high-cpu-threshold,omitempty"`
	TrapLogFullThreshold   *int64 `json:"trap-log-full-threshold,omitempty"`
	TrapLowMemoryThreshold *int64 `json:"trap-low-memory-threshold,omitempty"`
}

type SwitchControllerManagedSwitchSnmpUser struct {
	AuthProto     *string `json:"auth-proto,omitempty"`
	AuthPwd       *string `json:"auth-pwd,omitempty"`
	Name          *string `json:"name,omitempty"`
	PrivProto     *string `json:"priv-proto,omitempty"`
	PrivPwd       *string `json:"priv-pwd,omitempty"`
	Queries       *string `json:"queries,omitempty"`
	QueryPort     *int64  `json:"query-port,omitempty"`
	SecurityLevel *string `json:"security-level,omitempty"`
}

type SwitchControllerManagedSwitchStaticMac struct {
	Description *string `json:"description,omitempty"`
	Id          *int64  `json:"id,omitempty"`
	Interface   *string `json:"interface,omitempty"`
	Mac         *string `json:"mac,omitempty"`
	Type        *string `json:"type,omitempty"`
	Vlan        *string `json:"vlan,omitempty"`
}

type SwitchControllerManagedSwitchStormControl struct {
	Broadcast        *string `json:"broadcast,omitempty"`
	LocalOverride    *string `json:"local-override,omitempty"`
	Rate             *int64  `json:"rate,omitempty"`
	UnknownMulticast *string `json:"unknown-multicast,omitempty"`
	UnknownUnicast   *string `json:"unknown-unicast,omitempty"`
}

type SwitchControllerManagedSwitchStpInstance struct {
	Id       *string `json:"id,omitempty"`
	Priority *string `json:"priority,omitempty"`
}

type SwitchControllerManagedSwitchStpSettings struct {
	ForwardTime   *int64  `json:"forward-time,omitempty"`
	HelloTime     *int64  `json:"hello-time,omitempty"`
	LocalOverride *string `json:"local-override,omitempty"`
	MaxAge        *int64  `json:"max-age,omitempty"`
	MaxHops       *int64  `json:"max-hops,omitempty"`
	Name          *string `json:"name,omitempty"`
	PendingTimer  *int64  `json:"pending-timer,omitempty"`
	Revision      *int64  `json:"revision,omitempty"`
}

type SwitchControllerManagedSwitchSwitchLog struct {
	LocalOverride *string `json:"local-override,omitempty"`
	Severity      *string `json:"severity,omitempty"`
	Status        *string `json:"status,omitempty"`
}
