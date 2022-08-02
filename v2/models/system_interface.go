package models

const SystemInterfacePath = "system/interface/"

type SystemInterface struct {
	AcName                                *string                                  `json:"ac-name,omitempty"`
	Aggregate                             *string                                  `json:"aggregate,omitempty"`
	Algorithm                             *string                                  `json:"algorithm,omitempty"`
	Alias                                 *string                                  `json:"alias,omitempty"`
	Allowaccess                           *string                                  `json:"allowaccess,omitempty"`
	ApDiscover                            *string                                  `json:"ap-discover,omitempty"`
	Arpforward                            *string                                  `json:"arpforward,omitempty"`
	AuthCert                              *string                                  `json:"auth-cert,omitempty"`
	AuthPortalAddr                        *string                                  `json:"auth-portal-addr,omitempty"`
	AuthType                              *string                                  `json:"auth-type,omitempty"`
	AutoAuthExtensionDevice               *string                                  `json:"auto-auth-extension-device,omitempty"`
	BandwidthMeasureTime                  *int64                                   `json:"bandwidth-measure-time,omitempty"`
	Bfd                                   *string                                  `json:"bfd,omitempty"`
	BfdDesiredMinTx                       *int64                                   `json:"bfd-desired-min-tx,omitempty"`
	BfdDetectMult                         *int64                                   `json:"bfd-detect-mult,omitempty"`
	BfdRequiredMinRx                      *int64                                   `json:"bfd-required-min-rx,omitempty"`
	BroadcastForward                      *string                                  `json:"broadcast-forward,omitempty"`
	CliConnStatus                         *int64                                   `json:"cli-conn-status,omitempty"`
	ClientOptions                         *[]SystemInterfaceClientOptions          `json:"client-options,omitempty"`
	Color                                 *int64                                   `json:"color,omitempty"`
	DedicatedTo                           *string                                  `json:"dedicated-to,omitempty"`
	Defaultgw                             *string                                  `json:"defaultgw,omitempty"`
	Description                           *string                                  `json:"description,omitempty"`
	DetectedPeerMtu                       *int64                                   `json:"detected-peer-mtu,omitempty"`
	Detectprotocol                        *string                                  `json:"detectprotocol,omitempty"`
	Detectserver                          *string                                  `json:"detectserver,omitempty"`
	DeviceIdentification                  *string                                  `json:"device-identification,omitempty"`
	DeviceUserIdentification              *string                                  `json:"device-user-identification,omitempty"`
	Devindex                              *int64                                   `json:"devindex,omitempty"`
	DhcpClasslessRouteAddition            *string                                  `json:"dhcp-classless-route-addition,omitempty"`
	DhcpClientIdentifier                  *string                                  `json:"dhcp-client-identifier,omitempty"`
	DhcpRelayAgentOption                  *string                                  `json:"dhcp-relay-agent-option,omitempty"`
	DhcpRelayInterface                    *string                                  `json:"dhcp-relay-interface,omitempty"`
	DhcpRelayInterfaceSelectMethod        *string                                  `json:"dhcp-relay-interface-select-method,omitempty"`
	DhcpRelayIp                           *string                                  `json:"dhcp-relay-ip,omitempty"`
	DhcpRelayLinkSelection                *string                                  `json:"dhcp-relay-link-selection,omitempty"`
	DhcpRelayRequestAllServer             *string                                  `json:"dhcp-relay-request-all-server,omitempty"`
	DhcpRelayService                      *string                                  `json:"dhcp-relay-service,omitempty"`
	DhcpRelayType                         *string                                  `json:"dhcp-relay-type,omitempty"`
	DhcpRenewTime                         *int64                                   `json:"dhcp-renew-time,omitempty"`
	DhcpSnoopingServerList                *[]SystemInterfaceDhcpSnoopingServerList `json:"dhcp-snooping-server-list,omitempty"`
	DiscRetryTimeout                      *int64                                   `json:"disc-retry-timeout,omitempty"`
	DisconnectThreshold                   *int64                                   `json:"disconnect-threshold,omitempty"`
	Distance                              *int64                                   `json:"distance,omitempty"`
	DnsServerOverride                     *string                                  `json:"dns-server-override,omitempty"`
	DnsServerProtocol                     *string                                  `json:"dns-server-protocol,omitempty"`
	DropFragment                          *string                                  `json:"drop-fragment,omitempty"`
	DropOverlappedFragment                *string                                  `json:"drop-overlapped-fragment,omitempty"`
	EgressShapingProfile                  *string                                  `json:"egress-shaping-profile,omitempty"`
	EstimatedDownstreamBandwidth          *int64                                   `json:"estimated-downstream-bandwidth,omitempty"`
	EstimatedUpstreamBandwidth            *int64                                   `json:"estimated-upstream-bandwidth,omitempty"`
	ExplicitFtpProxy                      *string                                  `json:"explicit-ftp-proxy,omitempty"`
	ExplicitWebProxy                      *string                                  `json:"explicit-web-proxy,omitempty"`
	External                              *string                                  `json:"external,omitempty"`
	FailActionOnExtender                  *string                                  `json:"fail-action-on-extender,omitempty"`
	FailAlertInterfaces                   *[]SystemInterfaceFailAlertInterfaces    `json:"fail-alert-interfaces,omitempty"`
	FailAlertMethod                       *string                                  `json:"fail-alert-method,omitempty"`
	FailDetect                            *string                                  `json:"fail-detect,omitempty"`
	FailDetectOption                      *string                                  `json:"fail-detect-option,omitempty"`
	Fortilink                             *string                                  `json:"fortilink,omitempty"`
	FortilinkBackupLink                   *int64                                   `json:"fortilink-backup-link,omitempty"`
	FortilinkNeighborDetect               *string                                  `json:"fortilink-neighbor-detect,omitempty"`
	FortilinkSplitInterface               *string                                  `json:"fortilink-split-interface,omitempty"`
	FortilinkStacking                     *string                                  `json:"fortilink-stacking,omitempty"`
	ForwardDomain                         *int64                                   `json:"forward-domain,omitempty"`
	ForwardErrorCorrection                *string                                  `json:"forward-error-correction,omitempty"`
	Gwdetect                              *string                                  `json:"gwdetect,omitempty"`
	HaPriority                            *int64                                   `json:"ha-priority,omitempty"`
	IcmpAcceptRedirect                    *string                                  `json:"icmp-accept-redirect,omitempty"`
	IcmpSendRedirect                      *string                                  `json:"icmp-send-redirect,omitempty"`
	IdentAccept                           *string                                  `json:"ident-accept,omitempty"`
	IdleTimeout                           *int64                                   `json:"idle-timeout,omitempty"`
	Inbandwidth                           *int64                                   `json:"inbandwidth,omitempty"`
	IngressShapingProfile                 *string                                  `json:"ingress-shaping-profile,omitempty"`
	IngressSpilloverThreshold             *int64                                   `json:"ingress-spillover-threshold,omitempty"`
	Interface                             *string                                  `json:"interface,omitempty"`
	Internal                              *int64                                   `json:"internal,omitempty"`
	Ip                                    *string                                  `json:"ip,omitempty"`
	IpManagedByFortiipam                  *string                                  `json:"ip-managed-by-fortiipam,omitempty"`
	Ipmac                                 *string                                  `json:"ipmac,omitempty"`
	IpsSnifferMode                        *string                                  `json:"ips-sniffer-mode,omitempty"`
	Ipunnumbered                          *string                                  `json:"ipunnumbered,omitempty"`
	Ipv6                                  *SystemInterfaceIpv6                     `json:"ipv6,omitempty"`
	L2forward                             *string                                  `json:"l2forward,omitempty"`
	L2tpClient                            *string                                  `json:"l2tp-client,omitempty"`
	L2tpClientSettings                    *SystemInterfaceL2tpClientSettings       `json:"l2tp-client-settings,omitempty"`
	LacpHaSlave                           *string                                  `json:"lacp-ha-slave,omitempty"`
	LacpMode                              *string                                  `json:"lacp-mode,omitempty"`
	LacpSpeed                             *string                                  `json:"lacp-speed,omitempty"`
	LcpEchoInterval                       *int64                                   `json:"lcp-echo-interval,omitempty"`
	LcpMaxEchoFails                       *int64                                   `json:"lcp-max-echo-fails,omitempty"`
	LinkUpDelay                           *int64                                   `json:"link-up-delay,omitempty"`
	LldpNetworkPolicy                     *string                                  `json:"lldp-network-policy,omitempty"`
	LldpReception                         *string                                  `json:"lldp-reception,omitempty"`
	LldpTransmission                      *string                                  `json:"lldp-transmission,omitempty"`
	Macaddr                               *string                                  `json:"macaddr,omitempty"`
	ManagedSubnetworkSize                 *string                                  `json:"managed-subnetwork-size,omitempty"`
	ManagementIp                          *string                                  `json:"management-ip,omitempty"`
	MeasuredDownstreamBandwidth           *int64                                   `json:"measured-downstream-bandwidth,omitempty"`
	MeasuredUpstreamBandwidth             *int64                                   `json:"measured-upstream-bandwidth,omitempty"`
	Mediatype                             *string                                  `json:"mediatype,omitempty"`
	Member                                *[]SystemInterfaceMember                 `json:"member,omitempty"`
	MinLinks                              *int64                                   `json:"min-links,omitempty"`
	MinLinksDown                          *string                                  `json:"min-links-down,omitempty"`
	Mode                                  *string                                  `json:"mode,omitempty"`
	MonitorBandwidth                      *string                                  `json:"monitor-bandwidth,omitempty"`
	Mtu                                   *int64                                   `json:"mtu,omitempty"`
	MtuOverride                           *string                                  `json:"mtu-override,omitempty"`
	Name                                  *string                                  `json:"name,omitempty"`
	Ndiscforward                          *string                                  `json:"ndiscforward,omitempty"`
	NetbiosForward                        *string                                  `json:"netbios-forward,omitempty"`
	NetflowSampler                        *string                                  `json:"netflow-sampler,omitempty"`
	NpQosProfile                          *int64                                   `json:"np-qos-profile,omitempty"`
	Outbandwidth                          *int64                                   `json:"outbandwidth,omitempty"`
	PadtRetryTimeout                      *int64                                   `json:"padt-retry-timeout,omitempty"`
	Password                              *string                                  `json:"password,omitempty"`
	PingServStatus                        *int64                                   `json:"ping-serv-status,omitempty"`
	PollingInterval                       *int64                                   `json:"polling-interval,omitempty"`
	PppoeUnnumberedNegotiate              *string                                  `json:"pppoe-unnumbered-negotiate,omitempty"`
	PptpAuthType                          *string                                  `json:"pptp-auth-type,omitempty"`
	PptpClient                            *string                                  `json:"pptp-client,omitempty"`
	PptpPassword                          *string                                  `json:"pptp-password,omitempty"`
	PptpServerIp                          *string                                  `json:"pptp-server-ip,omitempty"`
	PptpTimeout                           *int64                                   `json:"pptp-timeout,omitempty"`
	PptpUser                              *string                                  `json:"pptp-user,omitempty"`
	PreserveSessionRoute                  *string                                  `json:"preserve-session-route,omitempty"`
	Priority                              *int64                                   `json:"priority,omitempty"`
	PriorityOverride                      *string                                  `json:"priority-override,omitempty"`
	ProxyCaptivePortal                    *string                                  `json:"proxy-captive-portal,omitempty"`
	ReachableTime                         *int64                                   `json:"reachable-time,omitempty"`
	RedundantInterface                    *string                                  `json:"redundant-interface,omitempty"`
	RemoteIp                              *string                                  `json:"remote-ip,omitempty"`
	ReplacemsgOverrideGroup               *string                                  `json:"replacemsg-override-group,omitempty"`
	Role                                  *string                                  `json:"role,omitempty"`
	SampleDirection                       *string                                  `json:"sample-direction,omitempty"`
	SampleRate                            *int64                                   `json:"sample-rate,omitempty"`
	SecondaryIp                           *string                                  `json:"secondary-IP,omitempty"`
	Secondaryip                           *[]SystemInterfaceSecondaryip            `json:"secondaryip,omitempty"`
	Security8021XDynamicVlanId            *int64                                   `json:"security-8021x-dynamic-vlan-id,omitempty"`
	Security8021XMaster                   *string                                  `json:"security-8021x-master,omitempty"`
	Security8021XMode                     *string                                  `json:"security-8021x-mode,omitempty"`
	SecurityExemptList                    *string                                  `json:"security-exempt-list,omitempty"`
	SecurityExternalLogout                *string                                  `json:"security-external-logout,omitempty"`
	SecurityExternalWeb                   *string                                  `json:"security-external-web,omitempty"`
	SecurityGroups                        *[]SystemInterfaceSecurityGroups         `json:"security-groups,omitempty"`
	SecurityMacAuthBypass                 *string                                  `json:"security-mac-auth-bypass,omitempty"`
	SecurityMode                          *string                                  `json:"security-mode,omitempty"`
	SecurityRedirectUrl                   *string                                  `json:"security-redirect-url,omitempty"`
	ServiceName                           *string                                  `json:"service-name,omitempty"`
	SflowSampler                          *string                                  `json:"sflow-sampler,omitempty"`
	SnmpIndex                             *int64                                   `json:"snmp-index,omitempty"`
	Speed                                 *string                                  `json:"speed,omitempty"`
	SpilloverThreshold                    *int64                                   `json:"spillover-threshold,omitempty"`
	SrcCheck                              *string                                  `json:"src-check,omitempty"`
	Status                                *string                                  `json:"status,omitempty"`
	Stp                                   *string                                  `json:"stp,omitempty"`
	StpHaSecondary                        *string                                  `json:"stp-ha-secondary,omitempty"`
	StpHaSlave                            *string                                  `json:"stp-ha-slave,omitempty"`
	Stpforward                            *string                                  `json:"stpforward,omitempty"`
	StpforwardMode                        *string                                  `json:"stpforward-mode,omitempty"`
	Subst                                 *string                                  `json:"subst,omitempty"`
	SubstituteDstMac                      *string                                  `json:"substitute-dst-mac,omitempty"`
	SwcFirstCreate                        *int64                                   `json:"swc-first-create,omitempty"`
	SwcVlan                               *int64                                   `json:"swc-vlan,omitempty"`
	Switch                                *string                                  `json:"switch,omitempty"`
	SwitchControllerAccessVlan            *string                                  `json:"switch-controller-access-vlan,omitempty"`
	SwitchControllerArpInspection         *string                                  `json:"switch-controller-arp-inspection,omitempty"`
	SwitchControllerDhcpSnooping          *string                                  `json:"switch-controller-dhcp-snooping,omitempty"`
	SwitchControllerDhcpSnoopingOption82  *string                                  `json:"switch-controller-dhcp-snooping-option82,omitempty"`
	SwitchControllerDhcpSnoopingVerifyMac *string                                  `json:"switch-controller-dhcp-snooping-verify-mac,omitempty"`
	SwitchControllerDynamic               *string                                  `json:"switch-controller-dynamic,omitempty"`
	SwitchControllerFeature               *string                                  `json:"switch-controller-feature,omitempty"`
	SwitchControllerIgmpSnooping          *string                                  `json:"switch-controller-igmp-snooping,omitempty"`
	SwitchControllerIgmpSnoopingFastLeave *string                                  `json:"switch-controller-igmp-snooping-fast-leave,omitempty"`
	SwitchControllerIgmpSnoopingProxy     *string                                  `json:"switch-controller-igmp-snooping-proxy,omitempty"`
	SwitchControllerIotScanning           *string                                  `json:"switch-controller-iot-scanning,omitempty"`
	SwitchControllerLearningLimit         *int64                                   `json:"switch-controller-learning-limit,omitempty"`
	SwitchControllerMgmtVlan              *int64                                   `json:"switch-controller-mgmt-vlan,omitempty"`
	SwitchControllerNac                   *string                                  `json:"switch-controller-nac,omitempty"`
	SwitchControllerRspanMode             *string                                  `json:"switch-controller-rspan-mode,omitempty"`
	SwitchControllerSourceIp              *string                                  `json:"switch-controller-source-ip,omitempty"`
	SwitchControllerTrafficPolicy         *string                                  `json:"switch-controller-traffic-policy,omitempty"`
	SystemId                              *string                                  `json:"system-id,omitempty"`
	SystemIdType                          *string                                  `json:"system-id-type,omitempty"`
	Tagging                               *[]SystemInterfaceTagging                `json:"tagging,omitempty"`
	TcpMss                                *int64                                   `json:"tcp-mss,omitempty"`
	Trunk                                 *string                                  `json:"trunk,omitempty"`
	TrustIp1                              *string                                  `json:"trust-ip-1,omitempty"`
	TrustIp2                              *string                                  `json:"trust-ip-2,omitempty"`
	TrustIp3                              *string                                  `json:"trust-ip-3,omitempty"`
	TrustIp61                             *string                                  `json:"trust-ip6-1,omitempty"`
	TrustIp62                             *string                                  `json:"trust-ip6-2,omitempty"`
	TrustIp63                             *string                                  `json:"trust-ip6-3,omitempty"`
	Type                                  *string                                  `json:"type,omitempty"`
	Username                              *string                                  `json:"username,omitempty"`
	Vdom                                  *string                                  `json:"vdom,omitempty"`
	Vindex                                *int64                                   `json:"vindex,omitempty"`
	VlanProtocol                          *string                                  `json:"vlan-protocol,omitempty"`
	Vlanforward                           *string                                  `json:"vlanforward,omitempty"`
	Vlanid                                *int64                                   `json:"vlanid,omitempty"`
	Vrf                                   *int64                                   `json:"vrf,omitempty"`
	Vrrp                                  *[]SystemInterfaceVrrp                   `json:"vrrp,omitempty"`
	VrrpVirtualMac                        *string                                  `json:"vrrp-virtual-mac,omitempty"`
	Wccp                                  *string                                  `json:"wccp,omitempty"`
	Weight                                *int64                                   `json:"weight,omitempty"`
	WinsIp                                *string                                  `json:"wins-ip,omitempty"`
}

const SystemInterfaceClientOptionsPath = "system/interface/client-options/"

type SystemInterfaceClientOptions struct {
	Code  *int64  `json:"code,omitempty"`
	Id    *int64  `json:"id,omitempty"`
	Ip    *string `json:"ip,omitempty"`
	Type  *string `json:"type,omitempty"`
	Value *string `json:"value,omitempty"`
}

const SystemInterfaceDhcpSnoopingServerListPath = "system/interface/dhcp-snooping-server-list/"

type SystemInterfaceDhcpSnoopingServerList struct {
	Name     *string `json:"name,omitempty"`
	ServerIp *string `json:"server-ip,omitempty"`
}

const SystemInterfaceFailAlertInterfacesPath = "system/interface/fail-alert-interfaces/"

type SystemInterfaceFailAlertInterfaces struct {
	Name *string `json:"name,omitempty"`
}

const SystemInterfaceIpv6Path = "system/interface/ipv6/"

type SystemInterfaceIpv6 struct {
	Autoconf                *string                                      `json:"autoconf,omitempty"`
	CliConn6Status          *int64                                       `json:"cli-conn6-status,omitempty"`
	Dhcp6ClientOptions      *string                                      `json:"dhcp6-client-options,omitempty"`
	Dhcp6IapdList           *[]SystemInterfaceIpv6Dhcp6IapdList          `json:"dhcp6-iapd-list,omitempty"`
	Dhcp6InformationRequest *string                                      `json:"dhcp6-information-request,omitempty"`
	Dhcp6PrefixDelegation   *string                                      `json:"dhcp6-prefix-delegation,omitempty"`
	Dhcp6PrefixHint         *string                                      `json:"dhcp6-prefix-hint,omitempty"`
	Dhcp6PrefixHintPlt      *int64                                       `json:"dhcp6-prefix-hint-plt,omitempty"`
	Dhcp6PrefixHintVlt      *int64                                       `json:"dhcp6-prefix-hint-vlt,omitempty"`
	Dhcp6RelayIp            *string                                      `json:"dhcp6-relay-ip,omitempty"`
	Dhcp6RelayService       *string                                      `json:"dhcp6-relay-service,omitempty"`
	Dhcp6RelayType          *string                                      `json:"dhcp6-relay-type,omitempty"`
	Icmp6SendRedirect       *string                                      `json:"icmp6-send-redirect,omitempty"`
	InterfaceIdentifier     *string                                      `json:"interface-identifier,omitempty"`
	Ip6Address              *string                                      `json:"ip6-address,omitempty"`
	Ip6Allowaccess          *string                                      `json:"ip6-allowaccess,omitempty"`
	Ip6DefaultLife          *int64                                       `json:"ip6-default-life,omitempty"`
	Ip6DelegatedPrefixIaid  *int64                                       `json:"ip6-delegated-prefix-iaid,omitempty"`
	Ip6DelegatedPrefixList  *[]SystemInterfaceIpv6Ip6DelegatedPrefixList `json:"ip6-delegated-prefix-list,omitempty"`
	Ip6DnsServerOverride    *string                                      `json:"ip6-dns-server-override,omitempty"`
	Ip6ExtraAddr            *[]SystemInterfaceIpv6Ip6ExtraAddr           `json:"ip6-extra-addr,omitempty"`
	Ip6HopLimit             *int64                                       `json:"ip6-hop-limit,omitempty"`
	Ip6LinkMtu              *int64                                       `json:"ip6-link-mtu,omitempty"`
	Ip6ManageFlag           *string                                      `json:"ip6-manage-flag,omitempty"`
	Ip6MaxInterval          *int64                                       `json:"ip6-max-interval,omitempty"`
	Ip6MinInterval          *int64                                       `json:"ip6-min-interval,omitempty"`
	Ip6Mode                 *string                                      `json:"ip6-mode,omitempty"`
	Ip6OtherFlag            *string                                      `json:"ip6-other-flag,omitempty"`
	Ip6PrefixList           *[]SystemInterfaceIpv6Ip6PrefixList          `json:"ip6-prefix-list,omitempty"`
	Ip6PrefixMode           *string                                      `json:"ip6-prefix-mode,omitempty"`
	Ip6ReachableTime        *int64                                       `json:"ip6-reachable-time,omitempty"`
	Ip6RetransTime          *int64                                       `json:"ip6-retrans-time,omitempty"`
	Ip6SendAdv              *string                                      `json:"ip6-send-adv,omitempty"`
	Ip6Subnet               *string                                      `json:"ip6-subnet,omitempty"`
	Ip6UpstreamInterface    *string                                      `json:"ip6-upstream-interface,omitempty"`
	NdCert                  *string                                      `json:"nd-cert,omitempty"`
	NdCgaModifier           *string                                      `json:"nd-cga-modifier,omitempty"`
	NdMode                  *string                                      `json:"nd-mode,omitempty"`
	NdSecurityLevel         *int64                                       `json:"nd-security-level,omitempty"`
	NdTimestampDelta        *int64                                       `json:"nd-timestamp-delta,omitempty"`
	NdTimestampFuzz         *int64                                       `json:"nd-timestamp-fuzz,omitempty"`
	RaSendMtu               *string                                      `json:"ra-send-mtu,omitempty"`
	UniqueAutoconfAddr      *string                                      `json:"unique-autoconf-addr,omitempty"`
	Vrip6_link_local        *string                                      `json:"vrip6_link_local,omitempty"`
	VrrpVirtualMac6         *string                                      `json:"vrrp-virtual-mac6,omitempty"`
	Vrrp6                   *[]SystemInterfaceIpv6Vrrp6                  `json:"vrrp6,omitempty"`
}

const SystemInterfaceIpv6Dhcp6IapdListPath = "system/interface/ipv6/dhcp6-iapd-list/"

type SystemInterfaceIpv6Dhcp6IapdList struct {
	Iaid          *int64  `json:"iaid,omitempty"`
	PrefixHint    *string `json:"prefix-hint,omitempty"`
	PrefixHintPlt *int64  `json:"prefix-hint-plt,omitempty"`
	PrefixHintVlt *int64  `json:"prefix-hint-vlt,omitempty"`
}

const SystemInterfaceIpv6Ip6DelegatedPrefixListPath = "system/interface/ipv6/ip6-delegated-prefix-list/"

type SystemInterfaceIpv6Ip6DelegatedPrefixList struct {
	AutonomousFlag      *string `json:"autonomous-flag,omitempty"`
	DelegatedPrefixIaid *int64  `json:"delegated-prefix-iaid,omitempty"`
	OnlinkFlag          *string `json:"onlink-flag,omitempty"`
	PrefixId            *int64  `json:"prefix-id,omitempty"`
	Rdnss               *string `json:"rdnss,omitempty"`
	RdnssService        *string `json:"rdnss-service,omitempty"`
	Subnet              *string `json:"subnet,omitempty"`
	UpstreamInterface   *string `json:"upstream-interface,omitempty"`
}

const SystemInterfaceIpv6Ip6ExtraAddrPath = "system/interface/ipv6/ip6-extra-addr/"

type SystemInterfaceIpv6Ip6ExtraAddr struct {
	Prefix *string `json:"prefix,omitempty"`
}

const SystemInterfaceIpv6Ip6PrefixListPath = "system/interface/ipv6/ip6-prefix-list/"

type SystemInterfaceIpv6Ip6PrefixList struct {
	AutonomousFlag    *string                                  `json:"autonomous-flag,omitempty"`
	Dnssl             *[]SystemInterfaceIpv6Ip6PrefixListDnssl `json:"dnssl,omitempty"`
	OnlinkFlag        *string                                  `json:"onlink-flag,omitempty"`
	PreferredLifeTime *int64                                   `json:"preferred-life-time,omitempty"`
	Prefix            *string                                  `json:"prefix,omitempty"`
	Rdnss             *string                                  `json:"rdnss,omitempty"`
	ValidLifeTime     *int64                                   `json:"valid-life-time,omitempty"`
}

const SystemInterfaceIpv6Ip6PrefixListDnsslPath = "system/interface/ipv6/ip6-prefix-list/dnssl/"

type SystemInterfaceIpv6Ip6PrefixListDnssl struct {
	Domain *string `json:"domain,omitempty"`
}

const SystemInterfaceIpv6Vrrp6Path = "system/interface/ipv6/vrrp6/"

type SystemInterfaceIpv6Vrrp6 struct {
	AcceptMode  *string `json:"accept-mode,omitempty"`
	AdvInterval *int64  `json:"adv-interval,omitempty"`
	Preempt     *string `json:"preempt,omitempty"`
	Priority    *int64  `json:"priority,omitempty"`
	StartTime   *int64  `json:"start-time,omitempty"`
	Status      *string `json:"status,omitempty"`
	Vrdst6      *string `json:"vrdst6,omitempty"`
	Vrgrp       *int64  `json:"vrgrp,omitempty"`
	Vrid        *int64  `json:"vrid,omitempty"`
	Vrip6       *string `json:"vrip6,omitempty"`
}

const SystemInterfaceL2tpClientSettingsPath = "system/interface/l2tp-client-settings/"

type SystemInterfaceL2tpClientSettings struct {
	AuthType  *string `json:"auth-type,omitempty"`
	Defaultgw *string `json:"defaultgw,omitempty"`
	Distance  *int64  `json:"distance,omitempty"`
	Ip        *string `json:"ip,omitempty"`
	Mtu       *int64  `json:"mtu,omitempty"`
	Password  *string `json:"password,omitempty"`
	PeerHost  *string `json:"peer-host,omitempty"`
	PeerMask  *string `json:"peer-mask,omitempty"`
	PeerPort  *int64  `json:"peer-port,omitempty"`
	Priority  *int64  `json:"priority,omitempty"`
	User      *string `json:"user,omitempty"`
}

const SystemInterfaceMemberPath = "system/interface/member/"

type SystemInterfaceMember struct {
	InterfaceName *string `json:"interface-name,omitempty"`
}

const SystemInterfaceSecondaryipPath = "system/interface/secondaryip/"

type SystemInterfaceSecondaryip struct {
	Allowaccess    *string `json:"allowaccess,omitempty"`
	Detectprotocol *string `json:"detectprotocol,omitempty"`
	Detectserver   *string `json:"detectserver,omitempty"`
	Gwdetect       *string `json:"gwdetect,omitempty"`
	HaPriority     *int64  `json:"ha-priority,omitempty"`
	Id             *int64  `json:"id,omitempty"`
	Ip             *string `json:"ip,omitempty"`
	PingServStatus *int64  `json:"ping-serv-status,omitempty"`
}

const SystemInterfaceSecurityGroupsPath = "system/interface/security-groups/"

type SystemInterfaceSecurityGroups struct {
	Name *string `json:"name,omitempty"`
}

const SystemInterfaceTaggingPath = "system/interface/tagging/"

type SystemInterfaceTagging struct {
	Category *string                       `json:"category,omitempty"`
	Name     *string                       `json:"name,omitempty"`
	Tags     *[]SystemInterfaceTaggingTags `json:"tags,omitempty"`
}

const SystemInterfaceTaggingTagsPath = "system/interface/tagging/tags/"

type SystemInterfaceTaggingTags struct {
	Name *string `json:"name,omitempty"`
}

const SystemInterfaceVrrpPath = "system/interface/vrrp/"

type SystemInterfaceVrrp struct {
	AcceptMode         *string                        `json:"accept-mode,omitempty"`
	AdvInterval        *int64                         `json:"adv-interval,omitempty"`
	IgnoreDefaultRoute *string                        `json:"ignore-default-route,omitempty"`
	Preempt            *string                        `json:"preempt,omitempty"`
	Priority           *int64                         `json:"priority,omitempty"`
	ProxyArp           *[]SystemInterfaceVrrpProxyArp `json:"proxy-arp,omitempty"`
	StartTime          *int64                         `json:"start-time,omitempty"`
	Status             *string                        `json:"status,omitempty"`
	Version            *string                        `json:"version,omitempty"`
	Vrdst              *string                        `json:"vrdst,omitempty"`
	VrdstPriority      *int64                         `json:"vrdst-priority,omitempty"`
	Vrgrp              *int64                         `json:"vrgrp,omitempty"`
	Vrid               *int64                         `json:"vrid,omitempty"`
	Vrip               *string                        `json:"vrip,omitempty"`
}

const SystemInterfaceVrrpProxyArpPath = "system/interface/vrrp/proxy-arp/"

type SystemInterfaceVrrpProxyArp struct {
	Id *int64  `json:"id,omitempty"`
	Ip *string `json:"ip,omitempty"`
}
