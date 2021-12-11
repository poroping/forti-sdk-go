package models

const SystemAccprofilePath = "system/accprofile/"

type SystemAccprofile struct {
	Admintimeout         *int64                             `json:"admintimeout,omitempty"`
	AdmintimeoutOverride *string                            `json:"admintimeout-override,omitempty"`
	Authgrp              *string                            `json:"authgrp,omitempty"`
	Comments             *string                            `json:"comments,omitempty"`
	Ftviewgrp            *string                            `json:"ftviewgrp,omitempty"`
	Fwgrp                *string                            `json:"fwgrp,omitempty"`
	FwgrpPermission      []SystemAccprofileFwgrpPermission  `json:"fwgrp-permission,omitempty"`
	Loggrp               *string                            `json:"loggrp,omitempty"`
	LoggrpPermission     []SystemAccprofileLoggrpPermission `json:"loggrp-permission,omitempty"`
	Name                 *string                            `json:"name,omitempty"`
	Netgrp               *string                            `json:"netgrp,omitempty"`
	NetgrpPermission     []SystemAccprofileNetgrpPermission `json:"netgrp-permission,omitempty"`
	Scope                *string                            `json:"scope,omitempty"`
	Secfabgrp            *string                            `json:"secfabgrp,omitempty"`
	Sysgrp               *string                            `json:"sysgrp,omitempty"`
	SysgrpPermission     []SystemAccprofileSysgrpPermission `json:"sysgrp-permission,omitempty"`
	SystemDiagnostics    *string                            `json:"system-diagnostics,omitempty"`
	Utmgrp               *string                            `json:"utmgrp,omitempty"`
	UtmgrpPermission     []SystemAccprofileUtmgrpPermission `json:"utmgrp-permission,omitempty"`
	Vpngrp               *string                            `json:"vpngrp,omitempty"`
	Wifi                 *string                            `json:"wifi,omitempty"`
}

type SystemAccprofileFwgrpPermission struct {
	Address  *string `json:"address,omitempty"`
	Others   *string `json:"others,omitempty"`
	Policy   *string `json:"policy,omitempty"`
	Schedule *string `json:"schedule,omitempty"`
	Service  *string `json:"service,omitempty"`
}

type SystemAccprofileLoggrpPermission struct {
	Config       *string `json:"config,omitempty"`
	DataAccess   *string `json:"data-access,omitempty"`
	ReportAccess *string `json:"report-access,omitempty"`
	ThreatWeight *string `json:"threat-weight,omitempty"`
}

type SystemAccprofileNetgrpPermission struct {
	Cfg           *string `json:"cfg,omitempty"`
	PacketCapture *string `json:"packet-capture,omitempty"`
	RouteCfg      *string `json:"route-cfg,omitempty"`
}

type SystemAccprofileSysgrpPermission struct {
	Admin *string `json:"admin,omitempty"`
	Cfg   *string `json:"cfg,omitempty"`
	Mnt   *string `json:"mnt,omitempty"`
	Upd   *string `json:"upd,omitempty"`
}

type SystemAccprofileUtmgrpPermission struct {
	Antivirus          *string `json:"antivirus,omitempty"`
	ApplicationControl *string `json:"application-control,omitempty"`
	DataLossPrevention *string `json:"data-loss-prevention,omitempty"`
	Dnsfilter          *string `json:"dnsfilter,omitempty"`
	Emailfilter        *string `json:"emailfilter,omitempty"`
	EndpointControl    *string `json:"endpoint-control,omitempty"`
	FileFilter         *string `json:"file-filter,omitempty"`
	Icap               *string `json:"icap,omitempty"`
	Ips                *string `json:"ips,omitempty"`
	Voip               *string `json:"voip,omitempty"`
	Waf                *string `json:"waf,omitempty"`
	Webfilter          *string `json:"webfilter,omitempty"`
}
