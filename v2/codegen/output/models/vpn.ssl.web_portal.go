package models

const VpnsslwebPortalPath = "vpn.ssl.web/portal/"

type VpnsslwebPortal struct {
	AllowUserAccess                  *string                                           `json:"allow-user-access,omitempty"`
	AutoConnect                      *string                                           `json:"auto-connect,omitempty"`
	BookmarkGroup                    []VpnsslwebPortalBookmarkGroup                    `json:"bookmark-group,omitempty"`
	Clipboard                        *string                                           `json:"clipboard,omitempty"`
	CustomLang                       *string                                           `json:"custom-lang,omitempty"`
	CustomizeForticlientDownloadUrl  *string                                           `json:"customize-forticlient-download-url,omitempty"`
	DisplayBookmark                  *string                                           `json:"display-bookmark,omitempty"`
	DisplayConnectionTools           *string                                           `json:"display-connection-tools,omitempty"`
	DisplayHistory                   *string                                           `json:"display-history,omitempty"`
	DisplayStatus                    *string                                           `json:"display-status,omitempty"`
	DnsServer1                       *string                                           `json:"dns-server1,omitempty"`
	DnsServer2                       *string                                           `json:"dns-server2,omitempty"`
	DnsSuffix                        *string                                           `json:"dns-suffix,omitempty"`
	ExclusiveRouting                 *string                                           `json:"exclusive-routing,omitempty"`
	ForticlientDownload              *string                                           `json:"forticlient-download,omitempty"`
	ForticlientDownloadMethod        *string                                           `json:"forticlient-download-method,omitempty"`
	Heading                          *string                                           `json:"heading,omitempty"`
	HideSsoCredential                *string                                           `json:"hide-sso-credential,omitempty"`
	HostCheck                        *string                                           `json:"host-check,omitempty"`
	HostCheckInterval                *int64                                            `json:"host-check-interval,omitempty"`
	HostCheckPolicy                  []VpnsslwebPortalHostCheckPolicy                  `json:"host-check-policy,omitempty"`
	IpMode                           *string                                           `json:"ip-mode,omitempty"`
	IpPools                          []VpnsslwebPortalIpPools                          `json:"ip-pools,omitempty"`
	Ipv6DnsServer1                   *string                                           `json:"ipv6-dns-server1,omitempty"`
	Ipv6DnsServer2                   *string                                           `json:"ipv6-dns-server2,omitempty"`
	Ipv6ExclusiveRouting             *string                                           `json:"ipv6-exclusive-routing,omitempty"`
	Ipv6Pools                        []VpnsslwebPortalIpv6Pools                        `json:"ipv6-pools,omitempty"`
	Ipv6ServiceRestriction           *string                                           `json:"ipv6-service-restriction,omitempty"`
	Ipv6SplitTunneling               *string                                           `json:"ipv6-split-tunneling,omitempty"`
	Ipv6SplitTunnelingRoutingAddress []VpnsslwebPortalIpv6SplitTunnelingRoutingAddress `json:"ipv6-split-tunneling-routing-address,omitempty"`
	Ipv6SplitTunnelingRoutingNegate  *string                                           `json:"ipv6-split-tunneling-routing-negate,omitempty"`
	Ipv6TunnelMode                   *string                                           `json:"ipv6-tunnel-mode,omitempty"`
	Ipv6WinsServer1                  *string                                           `json:"ipv6-wins-server1,omitempty"`
	Ipv6WinsServer2                  *string                                           `json:"ipv6-wins-server2,omitempty"`
	KeepAlive                        *string                                           `json:"keep-alive,omitempty"`
	LimitUserLogins                  *string                                           `json:"limit-user-logins,omitempty"`
	MacAddrAction                    *string                                           `json:"mac-addr-action,omitempty"`
	MacAddrCheck                     *string                                           `json:"mac-addr-check,omitempty"`
	MacAddrCheckRule                 []VpnsslwebPortalMacAddrCheckRule                 `json:"mac-addr-check-rule,omitempty"`
	MacosForticlientDownloadUrl      *string                                           `json:"macos-forticlient-download-url,omitempty"`
	Name                             *string                                           `json:"name,omitempty"`
	OsCheck                          *string                                           `json:"os-check,omitempty"`
	OsCheckList                      []VpnsslwebPortalOsCheckList                      `json:"os-check-list,omitempty"`
	PreferIpv6Dns                    *string                                           `json:"prefer-ipv6-dns,omitempty"`
	RedirUrl                         *string                                           `json:"redir-url,omitempty"`
	RewriteIpUriUi                   *string                                           `json:"rewrite-ip-uri-ui,omitempty"`
	SavePassword                     *string                                           `json:"save-password,omitempty"`
	ServiceRestriction               *string                                           `json:"service-restriction,omitempty"`
	SkipCheckForBrowser              *string                                           `json:"skip-check-for-browser,omitempty"`
	SkipCheckForUnsupportedOs        *string                                           `json:"skip-check-for-unsupported-os,omitempty"`
	SmbMaxVersion                    *string                                           `json:"smb-max-version,omitempty"`
	SmbMinVersion                    *string                                           `json:"smb-min-version,omitempty"`
	SmbNtlmv1Auth                    *string                                           `json:"smb-ntlmv1-auth,omitempty"`
	Smbv1                            *string                                           `json:"smbv1,omitempty"`
	SplitDns                         []VpnsslwebPortalSplitDns                         `json:"split-dns,omitempty"`
	SplitTunneling                   *string                                           `json:"split-tunneling,omitempty"`
	SplitTunnelingRoutingAddress     []VpnsslwebPortalSplitTunnelingRoutingAddress     `json:"split-tunneling-routing-address,omitempty"`
	SplitTunnelingRoutingNegate      *string                                           `json:"split-tunneling-routing-negate,omitempty"`
	Theme                            *string                                           `json:"theme,omitempty"`
	TunnelMode                       *string                                           `json:"tunnel-mode,omitempty"`
	UseSdwan                         *string                                           `json:"use-sdwan,omitempty"`
	UserBookmark                     *string                                           `json:"user-bookmark,omitempty"`
	UserGroupBookmark                *string                                           `json:"user-group-bookmark,omitempty"`
	WebMode                          *string                                           `json:"web-mode,omitempty"`
	WindowsForticlientDownloadUrl    *string                                           `json:"windows-forticlient-download-url,omitempty"`
	WinsServer1                      *string                                           `json:"wins-server1,omitempty"`
	WinsServer2                      *string                                           `json:"wins-server2,omitempty"`
}

type VpnsslwebPortalBookmarkGroup struct {
	Bookmarks []VpnsslwebPortalBookmarkGroupBookmarks `json:"bookmarks,omitempty"`
	Name      *string                                 `json:"name,omitempty"`
}

type VpnsslwebPortalBookmarkGroupBookmarks struct {
	AdditionalParams      *string                                         `json:"additional-params,omitempty"`
	Apptype               *string                                         `json:"apptype,omitempty"`
	ColorDepth            *string                                         `json:"color-depth,omitempty"`
	Description           *string                                         `json:"description,omitempty"`
	Domain                *string                                         `json:"domain,omitempty"`
	Folder                *string                                         `json:"folder,omitempty"`
	FormData              []VpnsslwebPortalBookmarkGroupBookmarksFormData `json:"form-data,omitempty"`
	Host                  *string                                         `json:"host,omitempty"`
	KeyboardLayout        *string                                         `json:"keyboard-layout,omitempty"`
	LoadBalancingInfo     *string                                         `json:"load-balancing-info,omitempty"`
	LogonPassword         *string                                         `json:"logon-password,omitempty"`
	LogonUser             *string                                         `json:"logon-user,omitempty"`
	Name                  *string                                         `json:"name,omitempty"`
	Port                  *int64                                          `json:"port,omitempty"`
	PreconnectionBlob     *string                                         `json:"preconnection-blob,omitempty"`
	PreconnectionId       *int64                                          `json:"preconnection-id,omitempty"`
	RestrictedAdmin       *string                                         `json:"restricted-admin,omitempty"`
	Security              *string                                         `json:"security,omitempty"`
	SendPreconnectionId   *string                                         `json:"send-preconnection-id,omitempty"`
	Sso                   *string                                         `json:"sso,omitempty"`
	SsoCredential         *string                                         `json:"sso-credential,omitempty"`
	SsoCredentialSentOnce *string                                         `json:"sso-credential-sent-once,omitempty"`
	SsoPassword           *string                                         `json:"sso-password,omitempty"`
	SsoUsername           *string                                         `json:"sso-username,omitempty"`
	Url                   *string                                         `json:"url,omitempty"`
}

type VpnsslwebPortalBookmarkGroupBookmarksFormData struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

type VpnsslwebPortalHostCheckPolicy struct {
	Name *string `json:"name,omitempty"`
}

type VpnsslwebPortalIpPools struct {
	Name *string `json:"name,omitempty"`
}

type VpnsslwebPortalIpv6Pools struct {
	Name *string `json:"name,omitempty"`
}

type VpnsslwebPortalIpv6SplitTunnelingRoutingAddress struct {
	Name *string `json:"name,omitempty"`
}

type VpnsslwebPortalMacAddrCheckRule struct {
	MacAddrList []VpnsslwebPortalMacAddrCheckRuleMacAddrList `json:"mac-addr-list,omitempty"`
	MacAddrMask *int64                                       `json:"mac-addr-mask,omitempty"`
	Name        *string                                      `json:"name,omitempty"`
}

type VpnsslwebPortalMacAddrCheckRuleMacAddrList struct {
	Addr *string `json:"addr,omitempty"`
}

type VpnsslwebPortalOsCheckList struct {
	Action           *string `json:"action,omitempty"`
	LatestPatchLevel *string `json:"latest-patch-level,omitempty"`
	Name             *string `json:"name,omitempty"`
	Tolerance        *int64  `json:"tolerance,omitempty"`
}

type VpnsslwebPortalSplitDns struct {
	DnsServer1     *string `json:"dns-server1,omitempty"`
	DnsServer2     *string `json:"dns-server2,omitempty"`
	Domains        *string `json:"domains,omitempty"`
	Id             *int64  `json:"id,omitempty"`
	Ipv6DnsServer1 *string `json:"ipv6-dns-server1,omitempty"`
	Ipv6DnsServer2 *string `json:"ipv6-dns-server2,omitempty"`
}

type VpnsslwebPortalSplitTunnelingRoutingAddress struct {
	Name *string `json:"name,omitempty"`
}
