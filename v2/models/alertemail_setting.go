package models

const AlertemailSettingPath = "alertemail/setting/"

type AlertemailSetting struct {
	FDSLicenseExpiringDays            *float64 `json:"FDS-license-expiring-days,omitempty"`
	FDSLicenseExpiringWarning         *string  `json:"FDS-license-expiring-warning,omitempty"`
	FDSUpdateLogs                     *string  `json:"FDS-update-logs,omitempty"`
	FIPSCCErrors                      *string  `json:"FIPS-CC-errors,omitempty"`
	FSSODisconnectLogs                *string  `json:"FSSO-disconnect-logs,omitempty"`
	HALogs                            *string  `json:"HA-logs,omitempty"`
	IPSLogs                           *string  `json:"IPS-logs,omitempty"`
	IPsecErrorsLogs                   *string  `json:"IPsec-errors-logs,omitempty"`
	PPPErrorsLogs                     *string  `json:"PPP-errors-logs,omitempty"`
	AdminLoginLogs                    *string  `json:"admin-login-logs,omitempty"`
	AlertInterval                     *float64 `json:"alert-interval,omitempty"`
	AmcInterfaceBypassMode            *string  `json:"amc-interface-bypass-mode,omitempty"`
	AntivirusLogs                     *string  `json:"antivirus-logs,omitempty"`
	ConfigurationChangesLogs          *string  `json:"configuration-changes-logs,omitempty"`
	CriticalInterval                  *float64 `json:"critical-interval,omitempty"`
	DebugInterval                     *float64 `json:"debug-interval,omitempty"`
	EmailInterval                     *float64 `json:"email-interval,omitempty"`
	EmergencyInterval                 *float64 `json:"emergency-interval,omitempty"`
	ErrorInterval                     *float64 `json:"error-interval,omitempty"`
	FilterMode                        *string  `json:"filter-mode,omitempty"`
	FirewallAuthenticationFailureLogs *string  `json:"firewall-authentication-failure-logs,omitempty"`
	FortiguardLogQuotaWarning         *string  `json:"fortiguard-log-quota-warning,omitempty"`
	InformationInterval               *float64 `json:"information-interval,omitempty"`
	LocalDiskUsage                    *float64 `json:"local-disk-usage,omitempty"`
	LogDiskUsageWarning               *string  `json:"log-disk-usage-warning,omitempty"`
	Mailto1                           *string  `json:"mailto1,omitempty"`
	Mailto2                           *string  `json:"mailto2,omitempty"`
	Mailto3                           *string  `json:"mailto3,omitempty"`
	NotificationInterval              *float64 `json:"notification-interval,omitempty"`
	Severity                          *string  `json:"severity,omitempty"`
	SshLogs                           *string  `json:"ssh-logs,omitempty"`
	SslvpnAuthenticationErrorsLogs    *string  `json:"sslvpn-authentication-errors-logs,omitempty"`
	Username                          *string  `json:"username,omitempty"`
	ViolationTrafficLogs              *string  `json:"violation-traffic-logs,omitempty"`
	WarningInterval                   *float64 `json:"warning-interval,omitempty"`
	WebfilterLogs                     *string  `json:"webfilter-logs,omitempty"`
}
