package models

const SystemGlobalPath = "system/global/"

type SystemGlobal struct {
	AdminConcurrent                      *string `json:"admin-concurrent,omitempty"`
	AdminConsoleTimeout                  *int64  `json:"admin-console-timeout,omitempty"`
	AdminForticloudSsoLogin              *string `json:"admin-forticloud-sso-login,omitempty"`
	AdminHstsMaxAge                      *int64  `json:"admin-hsts-max-age,omitempty"`
	AdminHttpsPkiRequired                *string `json:"admin-https-pki-required,omitempty"`
	AdminHttpsRedirect                   *string `json:"admin-https-redirect,omitempty"`
	AdminHttpsSslBannedCiphers           *string `json:"admin-https-ssl-banned-ciphers,omitempty"`
	AdminHttpsSslCiphersuites            *string `json:"admin-https-ssl-ciphersuites,omitempty"`
	AdminHttpsSslVersions                *string `json:"admin-https-ssl-versions,omitempty"`
	AdminLockoutDuration                 *int64  `json:"admin-lockout-duration,omitempty"`
	AdminLockoutThreshold                *int64  `json:"admin-lockout-threshold,omitempty"`
	AdminLoginMax                        *int64  `json:"admin-login-max,omitempty"`
	AdminMaintainer                      *string `json:"admin-maintainer,omitempty"`
	AdminPort                            *int64  `json:"admin-port,omitempty"`
	AdminResetButton                     *string `json:"admin-reset-button,omitempty"`
	AdminRestrictLocal                   *string `json:"admin-restrict-local,omitempty"`
	AdminScp                             *string `json:"admin-scp,omitempty"`
	AdminServerCert                      *string `json:"admin-server-cert,omitempty"`
	AdminSport                           *int64  `json:"admin-sport,omitempty"`
	AdminSshGraceTime                    *int64  `json:"admin-ssh-grace-time,omitempty"`
	AdminSshPassword                     *string `json:"admin-ssh-password,omitempty"`
	AdminSshPort                         *int64  `json:"admin-ssh-port,omitempty"`
	AdminSshV1                           *string `json:"admin-ssh-v1,omitempty"`
	AdminTelnet                          *string `json:"admin-telnet,omitempty"`
	AdminTelnetPort                      *int64  `json:"admin-telnet-port,omitempty"`
	Admintimeout                         *int64  `json:"admintimeout,omitempty"`
	Alias                                *string `json:"alias,omitempty"`
	AllowTrafficRedirect                 *string `json:"allow-traffic-redirect,omitempty"`
	AntiReplay                           *string `json:"anti-replay,omitempty"`
	ArpMaxEntry                          *int64  `json:"arp-max-entry,omitempty"`
	AuthCert                             *string `json:"auth-cert,omitempty"`
	AuthHttpPort                         *int64  `json:"auth-http-port,omitempty"`
	AuthHttpsPort                        *int64  `json:"auth-https-port,omitempty"`
	AuthKeepalive                        *string `json:"auth-keepalive,omitempty"`
	AuthSessionLimit                     *string `json:"auth-session-limit,omitempty"`
	AutoAuthExtensionDevice              *string `json:"auto-auth-extension-device,omitempty"`
	AutorunLogFsck                       *string `json:"autorun-log-fsck,omitempty"`
	AvFailopen                           *string `json:"av-failopen,omitempty"`
	AvFailopenSession                    *string `json:"av-failopen-session,omitempty"`
	BatchCmdb                            *string `json:"batch-cmdb,omitempty"`
	BlockSessionTimer                    *int64  `json:"block-session-timer,omitempty"`
	BrFdbMaxEntry                        *int64  `json:"br-fdb-max-entry,omitempty"`
	CertChainMax                         *int64  `json:"cert-chain-max,omitempty"`
	CfgRevertTimeout                     *int64  `json:"cfg-revert-timeout,omitempty"`
	CfgSave                              *string `json:"cfg-save,omitempty"`
	CheckProtocolHeader                  *string `json:"check-protocol-header,omitempty"`
	CheckResetRange                      *string `json:"check-reset-range,omitempty"`
	CliAuditLog                          *string `json:"cli-audit-log,omitempty"`
	CloudCommunication                   *string `json:"cloud-communication,omitempty"`
	CltCertReq                           *string `json:"clt-cert-req,omitempty"`
	CmdbsvrAffinity                      *string `json:"cmdbsvr-affinity,omitempty"`
	CpuUseThreshold                      *int64  `json:"cpu-use-threshold,omitempty"`
	CsrCaAttribute                       *string `json:"csr-ca-attribute,omitempty"`
	DailyRestart                         *string `json:"daily-restart,omitempty"`
	DefaultServiceSourcePort             *string `json:"default-service-source-port,omitempty"`
	DeviceIdleTimeout                    *int64  `json:"device-idle-timeout,omitempty"`
	DhParams                             *string `json:"dh-params,omitempty"`
	DnsproxyWorkerCount                  *int64  `json:"dnsproxy-worker-count,omitempty"`
	Dst                                  *string `json:"dst,omitempty"`
	EditVdomPrompt                       *string `json:"edit-vdom-prompt,omitempty"`
	ExtenderControllerReservedNetwork    *string `json:"extender-controller-reserved-network,omitempty"`
	Failtime                             *int64  `json:"failtime,omitempty"`
	FazDiskBufferSize                    *int64  `json:"faz-disk-buffer-size,omitempty"`
	FdsStatistics                        *string `json:"fds-statistics,omitempty"`
	FdsStatisticsPeriod                  *int64  `json:"fds-statistics-period,omitempty"`
	FgdAlertSubscription                 *string `json:"fgd-alert-subscription,omitempty"`
	Fortiextender                        *string `json:"fortiextender,omitempty"`
	FortiextenderDataPort                *int64  `json:"fortiextender-data-port,omitempty"`
	FortiextenderDiscoveryLockdown       *string `json:"fortiextender-discovery-lockdown,omitempty"`
	FortiextenderVlanMode                *string `json:"fortiextender-vlan-mode,omitempty"`
	FortiservicePort                     *int64  `json:"fortiservice-port,omitempty"`
	FortitokenCloud                      *string `json:"fortitoken-cloud,omitempty"`
	GuiAllowDefaultHostname              *string `json:"gui-allow-default-hostname,omitempty"`
	GuiCertificates                      *string `json:"gui-certificates,omitempty"`
	GuiCustomLanguage                    *string `json:"gui-custom-language,omitempty"`
	GuiDateFormat                        *string `json:"gui-date-format,omitempty"`
	GuiDateTimeSource                    *string `json:"gui-date-time-source,omitempty"`
	GuiDeviceLatitude                    *string `json:"gui-device-latitude,omitempty"`
	GuiDeviceLongitude                   *string `json:"gui-device-longitude,omitempty"`
	GuiDisplayHostname                   *string `json:"gui-display-hostname,omitempty"`
	GuiFirmwareUpgradeWarning            *string `json:"gui-firmware-upgrade-warning,omitempty"`
	GuiForticareRegistrationSetupWarning *string `json:"gui-forticare-registration-setup-warning,omitempty"`
	GuiFortigateCloudSandbox             *string `json:"gui-fortigate-cloud-sandbox,omitempty"`
	GuiIpv6                              *string `json:"gui-ipv6,omitempty"`
	GuiLocalOut                          *string `json:"gui-local-out,omitempty"`
	GuiReplacementMessageGroups          *string `json:"gui-replacement-message-groups,omitempty"`
	GuiRestApiCache                      *string `json:"gui-rest-api-cache,omitempty"`
	GuiTheme                             *string `json:"gui-theme,omitempty"`
	GuiWirelessOpensecurity              *string `json:"gui-wireless-opensecurity,omitempty"`
	HaAffinity                           *string `json:"ha-affinity,omitempty"`
	HonorDf                              *string `json:"honor-df,omitempty"`
	Hostname                             *string `json:"hostname,omitempty"`
	IgmpStateLimit                       *int64  `json:"igmp-state-limit,omitempty"`
	Interval                             *int64  `json:"interval,omitempty"`
	IpSrcPortRange                       *string `json:"ip-src-port-range,omitempty"`
	IpsecAsicOffload                     *string `json:"ipsec-asic-offload,omitempty"`
	IpsecHaSeqjumpRate                   *int64  `json:"ipsec-ha-seqjump-rate,omitempty"`
	IpsecHmacOffload                     *string `json:"ipsec-hmac-offload,omitempty"`
	IpsecSoftDecAsync                    *string `json:"ipsec-soft-dec-async,omitempty"`
	Ipv6AcceptDad                        *int64  `json:"ipv6-accept-dad,omitempty"`
	Ipv6AllowAnycastProbe                *string `json:"ipv6-allow-anycast-probe,omitempty"`
	Ipv6AllowTrafficRedirect             *string `json:"ipv6-allow-traffic-redirect,omitempty"`
	IrqTimeAccounting                    *string `json:"irq-time-accounting,omitempty"`
	Language                             *string `json:"language,omitempty"`
	Ldapconntimeout                      *int64  `json:"ldapconntimeout,omitempty"`
	LldpReception                        *string `json:"lldp-reception,omitempty"`
	LldpTransmission                     *string `json:"lldp-transmission,omitempty"`
	LogSslConnection                     *string `json:"log-ssl-connection,omitempty"`
	LogUuidAddress                       *string `json:"log-uuid-address,omitempty"`
	LoginTimestamp                       *string `json:"login-timestamp,omitempty"`
	LongVdomName                         *string `json:"long-vdom-name,omitempty"`
	ManagementIp                         *string `json:"management-ip,omitempty"`
	ManagementPort                       *int64  `json:"management-port,omitempty"`
	ManagementPortUseAdminSport          *string `json:"management-port-use-admin-sport,omitempty"`
	ManagementVdom                       *string `json:"management-vdom,omitempty"`
	MaxRouteCacheSize                    *int64  `json:"max-route-cache-size,omitempty"`
	MemoryUseThresholdExtreme            *int64  `json:"memory-use-threshold-extreme,omitempty"`
	MemoryUseThresholdGreen              *int64  `json:"memory-use-threshold-green,omitempty"`
	MemoryUseThresholdRed                *int64  `json:"memory-use-threshold-red,omitempty"`
	MiglogdChildren                      *int64  `json:"miglogd-children,omitempty"`
	MultiFactorAuthentication            *string `json:"multi-factor-authentication,omitempty"`
	NdpMaxEntry                          *int64  `json:"ndp-max-entry,omitempty"`
	PmtuDiscovery                        *string `json:"pmtu-discovery,omitempty"`
	PolicyAuthConcurrent                 *int64  `json:"policy-auth-concurrent,omitempty"`
	PostLoginBanner                      *string `json:"post-login-banner,omitempty"`
	PreLoginBanner                       *string `json:"pre-login-banner,omitempty"`
	PrivateDataEncryption                *string `json:"private-data-encryption,omitempty"`
	ProxyAuthLifetime                    *string `json:"proxy-auth-lifetime,omitempty"`
	ProxyAuthLifetimeTimeout             *int64  `json:"proxy-auth-lifetime-timeout,omitempty"`
	ProxyAuthTimeout                     *int64  `json:"proxy-auth-timeout,omitempty"`
	ProxyHardwareAcceleration            *string `json:"proxy-hardware-acceleration,omitempty"`
	ProxyReAuthenticationMode            *string `json:"proxy-re-authentication-mode,omitempty"`
	ProxyResourceMode                    *string `json:"proxy-resource-mode,omitempty"`
	ProxyWorkerCount                     *int64  `json:"proxy-worker-count,omitempty"`
	RadiusPort                           *int64  `json:"radius-port,omitempty"`
	RebootUponConfigRestore              *string `json:"reboot-upon-config-restore,omitempty"`
	Refresh                              *int64  `json:"refresh,omitempty"`
	Remoteauthtimeout                    *int64  `json:"remoteauthtimeout,omitempty"`
	ResetSessionlessTcp                  *string `json:"reset-sessionless-tcp,omitempty"`
	RestartTime                          *string `json:"restart-time,omitempty"`
	RevisionBackupOnLogout               *string `json:"revision-backup-on-logout,omitempty"`
	RevisionImageAutoBackup              *string `json:"revision-image-auto-backup,omitempty"`
	ScanunitCount                        *int64  `json:"scanunit-count,omitempty"`
	SecurityRatingResultSubmission       *string `json:"security-rating-result-submission,omitempty"`
	SecurityRatingRunOnSchedule          *string `json:"security-rating-run-on-schedule,omitempty"`
	SendPmtuIcmp                         *string `json:"send-pmtu-icmp,omitempty"`
	SnatRouteChange                      *string `json:"snat-route-change,omitempty"`
	SpecialFile23Support                 *string `json:"special-file-23-support,omitempty"`
	SpeedtestServer                      *string `json:"speedtest-server,omitempty"`
	SshEncAlgo                           *string `json:"ssh-enc-algo,omitempty"`
	SshKexAlgo                           *string `json:"ssh-kex-algo,omitempty"`
	SshMacAlgo                           *string `json:"ssh-mac-algo,omitempty"`
	SslMinProtoVersion                   *string `json:"ssl-min-proto-version,omitempty"`
	SslStaticKeyCiphers                  *string `json:"ssl-static-key-ciphers,omitempty"`
	SslvpnCipherHardwareAcceleration     *string `json:"sslvpn-cipher-hardware-acceleration,omitempty"`
	SslvpnEmsSnCheck                     *string `json:"sslvpn-ems-sn-check,omitempty"`
	SslvpnKxpHardwareAcceleration        *string `json:"sslvpn-kxp-hardware-acceleration,omitempty"`
	SslvpnMaxWorkerCount                 *int64  `json:"sslvpn-max-worker-count,omitempty"`
	SslvpnPluginVersionCheck             *string `json:"sslvpn-plugin-version-check,omitempty"`
	StrictDirtySessionCheck              *string `json:"strict-dirty-session-check,omitempty"`
	StrongCrypto                         *string `json:"strong-crypto,omitempty"`
	SwitchController                     *string `json:"switch-controller,omitempty"`
	SwitchControllerReservedNetwork      *string `json:"switch-controller-reserved-network,omitempty"`
	SysPerfLogInterval                   *int64  `json:"sys-perf-log-interval,omitempty"`
	TcpHalfcloseTimer                    *int64  `json:"tcp-halfclose-timer,omitempty"`
	TcpHalfopenTimer                     *int64  `json:"tcp-halfopen-timer,omitempty"`
	TcpOption                            *string `json:"tcp-option,omitempty"`
	TcpRstTimer                          *int64  `json:"tcp-rst-timer,omitempty"`
	TcpTimewaitTimer                     *int64  `json:"tcp-timewait-timer,omitempty"`
	Tftp                                 *string `json:"tftp,omitempty"`
	Timezone                             *string `json:"timezone,omitempty"`
	TrafficPriority                      *string `json:"traffic-priority,omitempty"`
	TrafficPriorityLevel                 *string `json:"traffic-priority-level,omitempty"`
	TwoFactorEmailExpiry                 *int64  `json:"two-factor-email-expiry,omitempty"`
	TwoFactorFacExpiry                   *int64  `json:"two-factor-fac-expiry,omitempty"`
	TwoFactorFtkExpiry                   *int64  `json:"two-factor-ftk-expiry,omitempty"`
	TwoFactorFtmExpiry                   *int64  `json:"two-factor-ftm-expiry,omitempty"`
	TwoFactorSmsExpiry                   *int64  `json:"two-factor-sms-expiry,omitempty"`
	UdpIdleTimer                         *int64  `json:"udp-idle-timer,omitempty"`
	UrlFilterCount                       *int64  `json:"url-filter-count,omitempty"`
	UserDeviceStoreMaxDevices            *int64  `json:"user-device-store-max-devices,omitempty"`
	UserDeviceStoreMaxUnifiedMem         *int64  `json:"user-device-store-max-unified-mem,omitempty"`
	UserDeviceStoreMaxUsers              *int64  `json:"user-device-store-max-users,omitempty"`
	UserServerCert                       *string `json:"user-server-cert,omitempty"`
	VdomMode                             *string `json:"vdom-mode,omitempty"`
	VipArpRange                          *string `json:"vip-arp-range,omitempty"`
	VirtualSwitchVlan                    *string `json:"virtual-switch-vlan,omitempty"`
	WadCsvcCsCount                       *int64  `json:"wad-csvc-cs-count,omitempty"`
	WadCsvcDbCount                       *int64  `json:"wad-csvc-db-count,omitempty"`
	WadMemoryChangeGranularity           *int64  `json:"wad-memory-change-granularity,omitempty"`
	WadSourceAffinity                    *string `json:"wad-source-affinity,omitempty"`
	WadWorkerCount                       *int64  `json:"wad-worker-count,omitempty"`
	WifiCaCertificate                    *string `json:"wifi-ca-certificate,omitempty"`
	WifiCertificate                      *string `json:"wifi-certificate,omitempty"`
	Wimax4gUsb                           *string `json:"wimax-4g-usb,omitempty"`
	WirelessController                   *string `json:"wireless-controller,omitempty"`
	WirelessControllerPort               *int64  `json:"wireless-controller-port,omitempty"`
}
