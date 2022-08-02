package cmdb

import (
	"github.com/poroping/forti-sdk-go/v2/config"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func New(config *config.Config) Endpoints {
	return &Client{config: config}
}

type Client struct {
	config *config.Config
}

type Endpoints interface {
	CreateSystemCustomLanguage(payload *models.SystemCustomLanguage, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemCustomLanguage(mkey string, params *models.CmdbRequestParams) (*models.SystemCustomLanguage, error)
	UpdateSystemCustomLanguage(mkey string, payload *models.SystemCustomLanguage, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemCustomLanguage(mkey string, params *models.CmdbRequestParams) error
	ListSystemCustomLanguage(params *models.CmdbRequestParams) (*[]models.SystemCustomLanguage, error)
	CreateSystemFssoPolling(payload *models.SystemFssoPolling, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemFssoPolling(mkey string, params *models.CmdbRequestParams) (*models.SystemFssoPolling, error)
	UpdateSystemFssoPolling(mkey string, payload *models.SystemFssoPolling, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemFssoPolling(mkey string, params *models.CmdbRequestParams) error
	CreateCertificateRemote(payload *models.CertificateRemote, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadCertificateRemote(mkey string, params *models.CmdbRequestParams) (*models.CertificateRemote, error)
	UpdateCertificateRemote(mkey string, payload *models.CertificateRemote, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteCertificateRemote(mkey string, params *models.CmdbRequestParams) error
	ListCertificateRemote(params *models.CmdbRequestParams) (*[]models.CertificateRemote, error)
	CreateMgmtDataStatus(payload *models.MgmtDataStatus, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadMgmtDataStatus(mkey string, params *models.CmdbRequestParams) (*models.MgmtDataStatus, error)
	UpdateMgmtDataStatus(mkey string, payload *models.MgmtDataStatus, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteMgmtDataStatus(mkey string, params *models.CmdbRequestParams) error
	CreateSystemPasswordPolicy(payload *models.SystemPasswordPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemPasswordPolicy(mkey string, params *models.CmdbRequestParams) (*models.SystemPasswordPolicy, error)
	UpdateSystemPasswordPolicy(mkey string, payload *models.SystemPasswordPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemPasswordPolicy(mkey string, params *models.CmdbRequestParams) error
	CreateUserTacacs(payload *models.UserTacacs, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadUserTacacs(mkey string, params *models.CmdbRequestParams) (*models.UserTacacs, error)
	UpdateUserTacacs(mkey string, payload *models.UserTacacs, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteUserTacacs(mkey string, params *models.CmdbRequestParams) error
	ListUserTacacs(params *models.CmdbRequestParams) (*[]models.UserTacacs, error)
	CreateSystemGeoipCountry(payload *models.SystemGeoipCountry, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemGeoipCountry(mkey string, params *models.CmdbRequestParams) (*models.SystemGeoipCountry, error)
	UpdateSystemGeoipCountry(mkey string, payload *models.SystemGeoipCountry, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemGeoipCountry(mkey string, params *models.CmdbRequestParams) error
	ListSystemGeoipCountry(params *models.CmdbRequestParams) (*[]models.SystemGeoipCountry, error)
	CreateWirelessControllerHotspot20H2qpWanMetric(payload *models.WirelessControllerHotspot20H2qpWanMetric, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerHotspot20H2qpWanMetric(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerHotspot20H2qpWanMetric, error)
	UpdateWirelessControllerHotspot20H2qpWanMetric(mkey string, payload *models.WirelessControllerHotspot20H2qpWanMetric, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerHotspot20H2qpWanMetric(mkey string, params *models.CmdbRequestParams) error
	ListWirelessControllerHotspot20H2qpWanMetric(params *models.CmdbRequestParams) (*[]models.WirelessControllerHotspot20H2qpWanMetric, error)
	CreateFirewallShaperPerIpShaper(payload *models.FirewallShaperPerIpShaper, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallShaperPerIpShaper(mkey string, params *models.CmdbRequestParams) (*models.FirewallShaperPerIpShaper, error)
	UpdateFirewallShaperPerIpShaper(mkey string, payload *models.FirewallShaperPerIpShaper, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallShaperPerIpShaper(mkey string, params *models.CmdbRequestParams) error
	ListFirewallShaperPerIpShaper(params *models.CmdbRequestParams) (*[]models.FirewallShaperPerIpShaper, error)
	CreateWebfilterContentHeader(payload *models.WebfilterContentHeader, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWebfilterContentHeader(mkey string, params *models.CmdbRequestParams) (*models.WebfilterContentHeader, error)
	UpdateWebfilterContentHeader(mkey string, payload *models.WebfilterContentHeader, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWebfilterContentHeader(mkey string, params *models.CmdbRequestParams) error
	ListWebfilterContentHeader(params *models.CmdbRequestParams) (*[]models.WebfilterContentHeader, error)
	CreateWebProxyWisp(payload *models.WebProxyWisp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWebProxyWisp(mkey string, params *models.CmdbRequestParams) (*models.WebProxyWisp, error)
	UpdateWebProxyWisp(mkey string, payload *models.WebProxyWisp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWebProxyWisp(mkey string, params *models.CmdbRequestParams) error
	ListWebProxyWisp(params *models.CmdbRequestParams) (*[]models.WebProxyWisp, error)
	CreateLogDiskFilter(payload *models.LogDiskFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogDiskFilter(mkey string, params *models.CmdbRequestParams) (*models.LogDiskFilter, error)
	UpdateLogDiskFilter(mkey string, payload *models.LogDiskFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogDiskFilter(mkey string, params *models.CmdbRequestParams) error
	CreateCredentialStoreDomainController(payload *models.CredentialStoreDomainController, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadCredentialStoreDomainController(mkey string, params *models.CmdbRequestParams) (*models.CredentialStoreDomainController, error)
	UpdateCredentialStoreDomainController(mkey string, payload *models.CredentialStoreDomainController, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteCredentialStoreDomainController(mkey string, params *models.CmdbRequestParams) error
	ListCredentialStoreDomainController(params *models.CmdbRequestParams) (*[]models.CredentialStoreDomainController, error)
	CreateSystemPhysicalSwitch(payload *models.SystemPhysicalSwitch, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemPhysicalSwitch(mkey string, params *models.CmdbRequestParams) (*models.SystemPhysicalSwitch, error)
	UpdateSystemPhysicalSwitch(mkey string, payload *models.SystemPhysicalSwitch, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemPhysicalSwitch(mkey string, params *models.CmdbRequestParams) error
	ListSystemPhysicalSwitch(params *models.CmdbRequestParams) (*[]models.SystemPhysicalSwitch, error)
	CreateSwitchControllerAutoConfigPolicy(payload *models.SwitchControllerAutoConfigPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerAutoConfigPolicy(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerAutoConfigPolicy, error)
	UpdateSwitchControllerAutoConfigPolicy(mkey string, payload *models.SwitchControllerAutoConfigPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerAutoConfigPolicy(mkey string, params *models.CmdbRequestParams) error
	ListSwitchControllerAutoConfigPolicy(params *models.CmdbRequestParams) (*[]models.SwitchControllerAutoConfigPolicy, error)
	CreateFirewallIpv6EhFilter(payload *models.FirewallIpv6EhFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallIpv6EhFilter(mkey string, params *models.CmdbRequestParams) (*models.FirewallIpv6EhFilter, error)
	UpdateFirewallIpv6EhFilter(mkey string, payload *models.FirewallIpv6EhFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallIpv6EhFilter(mkey string, params *models.CmdbRequestParams) error
	CreateVpnIpsecPhase2(payload *models.VpnIpsecPhase2, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnIpsecPhase2(mkey string, params *models.CmdbRequestParams) (*models.VpnIpsecPhase2, error)
	UpdateVpnIpsecPhase2(mkey string, payload *models.VpnIpsecPhase2, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnIpsecPhase2(mkey string, params *models.CmdbRequestParams) error
	ListVpnIpsecPhase2(params *models.CmdbRequestParams) (*[]models.VpnIpsecPhase2, error)
	CreateWebfilterFtgdLocalRating(payload *models.WebfilterFtgdLocalRating, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWebfilterFtgdLocalRating(mkey string, params *models.CmdbRequestParams) (*models.WebfilterFtgdLocalRating, error)
	UpdateWebfilterFtgdLocalRating(mkey string, payload *models.WebfilterFtgdLocalRating, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWebfilterFtgdLocalRating(mkey string, params *models.CmdbRequestParams) error
	ListWebfilterFtgdLocalRating(params *models.CmdbRequestParams) (*[]models.WebfilterFtgdLocalRating, error)
	CreateRouterCommunityList(payload *models.RouterCommunityList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterCommunityList(mkey string, params *models.CmdbRequestParams) (*models.RouterCommunityList, error)
	UpdateRouterCommunityList(mkey string, payload *models.RouterCommunityList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterCommunityList(mkey string, params *models.CmdbRequestParams) error
	ListRouterCommunityList(params *models.CmdbRequestParams) (*[]models.RouterCommunityList, error)
	CreateRouterInfo(payload *models.RouterInfo, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterInfo(mkey string, params *models.CmdbRequestParams) (*models.RouterInfo, error)
	UpdateRouterInfo(mkey string, payload *models.RouterInfo, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterInfo(mkey string, params *models.CmdbRequestParams) error
	CreateSystemVirtualWirePair(payload *models.SystemVirtualWirePair, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemVirtualWirePair(mkey string, params *models.CmdbRequestParams) (*models.SystemVirtualWirePair, error)
	UpdateSystemVirtualWirePair(mkey string, payload *models.SystemVirtualWirePair, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemVirtualWirePair(mkey string, params *models.CmdbRequestParams) error
	ListSystemVirtualWirePair(params *models.CmdbRequestParams) (*[]models.SystemVirtualWirePair, error)
	CreateSwitchControllerPtpSettings(payload *models.SwitchControllerPtpSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerPtpSettings(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerPtpSettings, error)
	UpdateSwitchControllerPtpSettings(mkey string, payload *models.SwitchControllerPtpSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerPtpSettings(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallVip(payload *models.FirewallVip, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallVip(mkey string, params *models.CmdbRequestParams) (*models.FirewallVip, error)
	UpdateFirewallVip(mkey string, payload *models.FirewallVip, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallVip(mkey string, params *models.CmdbRequestParams) error
	ListFirewallVip(params *models.CmdbRequestParams) (*[]models.FirewallVip, error)
	CreateApplicationName(payload *models.ApplicationName, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadApplicationName(mkey string, params *models.CmdbRequestParams) (*models.ApplicationName, error)
	UpdateApplicationName(mkey string, payload *models.ApplicationName, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteApplicationName(mkey string, params *models.CmdbRequestParams) error
	ListApplicationName(params *models.CmdbRequestParams) (*[]models.ApplicationName, error)
	CreateSystemInfoAdminSsh(payload *models.SystemInfoAdminSsh, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemInfoAdminSsh(mkey string, params *models.CmdbRequestParams) (*models.SystemInfoAdminSsh, error)
	UpdateSystemInfoAdminSsh(mkey string, payload *models.SystemInfoAdminSsh, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemInfoAdminSsh(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallScheduleRecurring(payload *models.FirewallScheduleRecurring, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallScheduleRecurring(mkey string, params *models.CmdbRequestParams) (*models.FirewallScheduleRecurring, error)
	UpdateFirewallScheduleRecurring(mkey string, payload *models.FirewallScheduleRecurring, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallScheduleRecurring(mkey string, params *models.CmdbRequestParams) error
	ListFirewallScheduleRecurring(params *models.CmdbRequestParams) (*[]models.FirewallScheduleRecurring, error)
	CreateFirewallSslSshProfile(payload *models.FirewallSslSshProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallSslSshProfile(mkey string, params *models.CmdbRequestParams) (*models.FirewallSslSshProfile, error)
	UpdateFirewallSslSshProfile(mkey string, payload *models.FirewallSslSshProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallSslSshProfile(mkey string, params *models.CmdbRequestParams) error
	ListFirewallSslSshProfile(params *models.CmdbRequestParams) (*[]models.FirewallSslSshProfile, error)
	CreateLogTacacsaccounting2Setting(payload *models.LogTacacsaccounting2Setting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogTacacsaccounting2Setting(mkey string, params *models.CmdbRequestParams) (*models.LogTacacsaccounting2Setting, error)
	UpdateLogTacacsaccounting2Setting(mkey string, payload *models.LogTacacsaccounting2Setting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogTacacsaccounting2Setting(mkey string, params *models.CmdbRequestParams) error
	CreateSystemInfoAdminStatus(payload *models.SystemInfoAdminStatus, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemInfoAdminStatus(mkey string, params *models.CmdbRequestParams) (*models.SystemInfoAdminStatus, error)
	UpdateSystemInfoAdminStatus(mkey string, payload *models.SystemInfoAdminStatus, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemInfoAdminStatus(mkey string, params *models.CmdbRequestParams) error
	CreateSystemIpipTunnel(payload *models.SystemIpipTunnel, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemIpipTunnel(mkey string, params *models.CmdbRequestParams) (*models.SystemIpipTunnel, error)
	UpdateSystemIpipTunnel(mkey string, payload *models.SystemIpipTunnel, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemIpipTunnel(mkey string, params *models.CmdbRequestParams) error
	ListSystemIpipTunnel(params *models.CmdbRequestParams) (*[]models.SystemIpipTunnel, error)
	CreateVpnIpsecTunnelDetails(payload *models.VpnIpsecTunnelDetails, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnIpsecTunnelDetails(mkey string, params *models.CmdbRequestParams) (*models.VpnIpsecTunnelDetails, error)
	UpdateVpnIpsecTunnelDetails(mkey string, payload *models.VpnIpsecTunnelDetails, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnIpsecTunnelDetails(mkey string, params *models.CmdbRequestParams) error
	CreateSystemMgmtCsum(payload *models.SystemMgmtCsum, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemMgmtCsum(mkey string, params *models.CmdbRequestParams) (*models.SystemMgmtCsum, error)
	UpdateSystemMgmtCsum(mkey string, payload *models.SystemMgmtCsum, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemMgmtCsum(mkey string, params *models.CmdbRequestParams) error
	CreateSystemSnmpUser(payload *models.SystemSnmpUser, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemSnmpUser(mkey string, params *models.CmdbRequestParams) (*models.SystemSnmpUser, error)
	UpdateSystemSnmpUser(mkey string, payload *models.SystemSnmpUser, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemSnmpUser(mkey string, params *models.CmdbRequestParams) error
	ListSystemSnmpUser(params *models.CmdbRequestParams) (*[]models.SystemSnmpUser, error)
	CreateSystemProxyArp(payload *models.SystemProxyArp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemProxyArp(mkey string, params *models.CmdbRequestParams) (*models.SystemProxyArp, error)
	UpdateSystemProxyArp(mkey string, payload *models.SystemProxyArp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemProxyArp(mkey string, params *models.CmdbRequestParams) error
	ListSystemProxyArp(params *models.CmdbRequestParams) (*[]models.SystemProxyArp, error)
	CreateWirelessControllerVap(payload *models.WirelessControllerVap, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerVap(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerVap, error)
	UpdateWirelessControllerVap(mkey string, payload *models.WirelessControllerVap, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerVap(mkey string, params *models.CmdbRequestParams) error
	ListWirelessControllerVap(params *models.CmdbRequestParams) (*[]models.WirelessControllerVap, error)
	CreateEmailfilterFortishield(payload *models.EmailfilterFortishield, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadEmailfilterFortishield(mkey string, params *models.CmdbRequestParams) (*models.EmailfilterFortishield, error)
	UpdateEmailfilterFortishield(mkey string, payload *models.EmailfilterFortishield, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteEmailfilterFortishield(mkey string, params *models.CmdbRequestParams) error
	CreateLogSetting(payload *models.LogSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogSetting(mkey string, params *models.CmdbRequestParams) (*models.LogSetting, error)
	UpdateLogSetting(mkey string, payload *models.LogSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogSetting(mkey string, params *models.CmdbRequestParams) error
	CreateSystemReplacemsgAutomation(payload *models.SystemReplacemsgAutomation, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemReplacemsgAutomation(mkey string, params *models.CmdbRequestParams) (*models.SystemReplacemsgAutomation, error)
	UpdateSystemReplacemsgAutomation(mkey string, payload *models.SystemReplacemsgAutomation, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemReplacemsgAutomation(mkey string, params *models.CmdbRequestParams) error
	ListSystemReplacemsgAutomation(params *models.CmdbRequestParams) (*[]models.SystemReplacemsgAutomation, error)
	CreateFirewallInternetServiceReputation(payload *models.FirewallInternetServiceReputation, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallInternetServiceReputation(mkey string, params *models.CmdbRequestParams) (*models.FirewallInternetServiceReputation, error)
	UpdateFirewallInternetServiceReputation(mkey string, payload *models.FirewallInternetServiceReputation, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallInternetServiceReputation(mkey string, params *models.CmdbRequestParams) error
	ListFirewallInternetServiceReputation(params *models.CmdbRequestParams) (*[]models.FirewallInternetServiceReputation, error)
	CreateVpnSslWebHostCheckSoftware(payload *models.VpnSslWebHostCheckSoftware, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnSslWebHostCheckSoftware(mkey string, params *models.CmdbRequestParams) (*models.VpnSslWebHostCheckSoftware, error)
	UpdateVpnSslWebHostCheckSoftware(mkey string, payload *models.VpnSslWebHostCheckSoftware, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnSslWebHostCheckSoftware(mkey string, params *models.CmdbRequestParams) error
	ListVpnSslWebHostCheckSoftware(params *models.CmdbRequestParams) (*[]models.VpnSslWebHostCheckSoftware, error)
	CreateSystemIpam(payload *models.SystemIpam, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemIpam(mkey string, params *models.CmdbRequestParams) (*models.SystemIpam, error)
	UpdateSystemIpam(mkey string, payload *models.SystemIpam, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemIpam(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerGlobal(payload *models.WirelessControllerGlobal, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerGlobal(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerGlobal, error)
	UpdateWirelessControllerGlobal(mkey string, payload *models.WirelessControllerGlobal, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerGlobal(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerApcfgProfile(payload *models.WirelessControllerApcfgProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerApcfgProfile(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerApcfgProfile, error)
	UpdateWirelessControllerApcfgProfile(mkey string, payload *models.WirelessControllerApcfgProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerApcfgProfile(mkey string, params *models.CmdbRequestParams) error
	ListWirelessControllerApcfgProfile(params *models.CmdbRequestParams) (*[]models.WirelessControllerApcfgProfile, error)
	CreateDlpSensitivity(payload *models.DlpSensitivity, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadDlpSensitivity(mkey string, params *models.CmdbRequestParams) (*models.DlpSensitivity, error)
	UpdateDlpSensitivity(mkey string, payload *models.DlpSensitivity, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteDlpSensitivity(mkey string, params *models.CmdbRequestParams) error
	ListDlpSensitivity(params *models.CmdbRequestParams) (*[]models.DlpSensitivity, error)
	CreateRouterOspf(payload *models.RouterOspf, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterOspf(mkey string, params *models.CmdbRequestParams) (*models.RouterOspf, error)
	UpdateRouterOspf(mkey string, payload *models.RouterOspf, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterOspf(mkey string, params *models.CmdbRequestParams) error
	CreateSwitchControllerNacDevice(payload *models.SwitchControllerNacDevice, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerNacDevice(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerNacDevice, error)
	UpdateSwitchControllerNacDevice(mkey string, payload *models.SwitchControllerNacDevice, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerNacDevice(mkey string, params *models.CmdbRequestParams) error
	ListSwitchControllerNacDevice(params *models.CmdbRequestParams) (*[]models.SwitchControllerNacDevice, error)
	CreateSystemFortianalyzerConnectivity(payload *models.SystemFortianalyzerConnectivity, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemFortianalyzerConnectivity(mkey string, params *models.CmdbRequestParams) (*models.SystemFortianalyzerConnectivity, error)
	UpdateSystemFortianalyzerConnectivity(mkey string, payload *models.SystemFortianalyzerConnectivity, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemFortianalyzerConnectivity(mkey string, params *models.CmdbRequestParams) error
	CreateVpnCertificateLocal(payload *models.VpnCertificateLocal, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnCertificateLocal(mkey string, params *models.CmdbRequestParams) (*models.VpnCertificateLocal, error)
	UpdateVpnCertificateLocal(mkey string, payload *models.VpnCertificateLocal, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnCertificateLocal(mkey string, params *models.CmdbRequestParams) error
	ListVpnCertificateLocal(params *models.CmdbRequestParams) (*[]models.VpnCertificateLocal, error)
	CreateEmailfilterIptrust(payload *models.EmailfilterIptrust, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadEmailfilterIptrust(mkey string, params *models.CmdbRequestParams) (*models.EmailfilterIptrust, error)
	UpdateEmailfilterIptrust(mkey string, payload *models.EmailfilterIptrust, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteEmailfilterIptrust(mkey string, params *models.CmdbRequestParams) error
	ListEmailfilterIptrust(params *models.CmdbRequestParams) (*[]models.EmailfilterIptrust, error)
	CreateLogFortianalyzer2Setting(payload *models.LogFortianalyzer2Setting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogFortianalyzer2Setting(mkey string, params *models.CmdbRequestParams) (*models.LogFortianalyzer2Setting, error)
	UpdateLogFortianalyzer2Setting(mkey string, payload *models.LogFortianalyzer2Setting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogFortianalyzer2Setting(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerScan(payload *models.WirelessControllerScan, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerScan(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerScan, error)
	UpdateWirelessControllerScan(mkey string, payload *models.WirelessControllerScan, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerScan(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallInternetServiceName(payload *models.FirewallInternetServiceName, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallInternetServiceName(mkey string, params *models.CmdbRequestParams) (*models.FirewallInternetServiceName, error)
	UpdateFirewallInternetServiceName(mkey string, payload *models.FirewallInternetServiceName, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallInternetServiceName(mkey string, params *models.CmdbRequestParams) error
	ListFirewallInternetServiceName(params *models.CmdbRequestParams) (*[]models.FirewallInternetServiceName, error)
	CreateLogSyslogd3Filter(payload *models.LogSyslogd3Filter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogSyslogd3Filter(mkey string, params *models.CmdbRequestParams) (*models.LogSyslogd3Filter, error)
	UpdateLogSyslogd3Filter(mkey string, payload *models.LogSyslogd3Filter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogSyslogd3Filter(mkey string, params *models.CmdbRequestParams) error
	CreateLogFortiguardFilter(payload *models.LogFortiguardFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogFortiguardFilter(mkey string, params *models.CmdbRequestParams) (*models.LogFortiguardFilter, error)
	UpdateLogFortiguardFilter(mkey string, payload *models.LogFortiguardFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogFortiguardFilter(mkey string, params *models.CmdbRequestParams) error
	CreateSwitchControllerInitialConfigVlans(payload *models.SwitchControllerInitialConfigVlans, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerInitialConfigVlans(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerInitialConfigVlans, error)
	UpdateSwitchControllerInitialConfigVlans(mkey string, payload *models.SwitchControllerInitialConfigVlans, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerInitialConfigVlans(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallAddress6(payload *models.FirewallAddress6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallAddress6(mkey string, params *models.CmdbRequestParams) (*models.FirewallAddress6, error)
	UpdateFirewallAddress6(mkey string, payload *models.FirewallAddress6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallAddress6(mkey string, params *models.CmdbRequestParams) error
	ListFirewallAddress6(params *models.CmdbRequestParams) (*[]models.FirewallAddress6, error)
	CreateVpnSslMonitor(payload *models.VpnSslMonitor, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnSslMonitor(mkey string, params *models.CmdbRequestParams) (*models.VpnSslMonitor, error)
	UpdateVpnSslMonitor(mkey string, payload *models.VpnSslMonitor, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnSslMonitor(mkey string, params *models.CmdbRequestParams) error
	CreateSwitchControllerSnmpCommunity(payload *models.SwitchControllerSnmpCommunity, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerSnmpCommunity(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerSnmpCommunity, error)
	UpdateSwitchControllerSnmpCommunity(mkey string, payload *models.SwitchControllerSnmpCommunity, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerSnmpCommunity(mkey string, params *models.CmdbRequestParams) error
	ListSwitchControllerSnmpCommunity(params *models.CmdbRequestParams) (*[]models.SwitchControllerSnmpCommunity, error)
	CreateVpnIpsecPhase1(payload *models.VpnIpsecPhase1, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnIpsecPhase1(mkey string, params *models.CmdbRequestParams) (*models.VpnIpsecPhase1, error)
	UpdateVpnIpsecPhase1(mkey string, payload *models.VpnIpsecPhase1, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnIpsecPhase1(mkey string, params *models.CmdbRequestParams) error
	ListVpnIpsecPhase1(params *models.CmdbRequestParams) (*[]models.VpnIpsecPhase1, error)
	CreateVpnOcvpn(payload *models.VpnOcvpn, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnOcvpn(mkey string, params *models.CmdbRequestParams) (*models.VpnOcvpn, error)
	UpdateVpnOcvpn(mkey string, payload *models.VpnOcvpn, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnOcvpn(mkey string, params *models.CmdbRequestParams) error
	CreateCertificateCa(payload *models.CertificateCa, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadCertificateCa(mkey string, params *models.CmdbRequestParams) (*models.CertificateCa, error)
	UpdateCertificateCa(mkey string, payload *models.CertificateCa, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteCertificateCa(mkey string, params *models.CmdbRequestParams) error
	ListCertificateCa(params *models.CmdbRequestParams) (*[]models.CertificateCa, error)
	CreateSshFilterProfile(payload *models.SshFilterProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSshFilterProfile(mkey string, params *models.CmdbRequestParams) (*models.SshFilterProfile, error)
	UpdateSshFilterProfile(mkey string, payload *models.SshFilterProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSshFilterProfile(mkey string, params *models.CmdbRequestParams) error
	ListSshFilterProfile(params *models.CmdbRequestParams) (*[]models.SshFilterProfile, error)
	CreateSystemNetworkVisibility(payload *models.SystemNetworkVisibility, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemNetworkVisibility(mkey string, params *models.CmdbRequestParams) (*models.SystemNetworkVisibility, error)
	UpdateSystemNetworkVisibility(mkey string, payload *models.SystemNetworkVisibility, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemNetworkVisibility(mkey string, params *models.CmdbRequestParams) error
	CreateIpsViewMap(payload *models.IpsViewMap, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadIpsViewMap(mkey string, params *models.CmdbRequestParams) (*models.IpsViewMap, error)
	UpdateIpsViewMap(mkey string, payload *models.IpsViewMap, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteIpsViewMap(mkey string, params *models.CmdbRequestParams) error
	ListIpsViewMap(params *models.CmdbRequestParams) (*[]models.IpsViewMap, error)
	CreateAntivirusQuarantine(payload *models.AntivirusQuarantine, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadAntivirusQuarantine(mkey string, params *models.CmdbRequestParams) (*models.AntivirusQuarantine, error)
	UpdateAntivirusQuarantine(mkey string, payload *models.AntivirusQuarantine, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteAntivirusQuarantine(mkey string, params *models.CmdbRequestParams) error
	CreateSystemDhcpServer(payload *models.SystemDhcpServer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemDhcpServer(mkey string, params *models.CmdbRequestParams) (*models.SystemDhcpServer, error)
	UpdateSystemDhcpServer(mkey string, payload *models.SystemDhcpServer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemDhcpServer(mkey string, params *models.CmdbRequestParams) error
	ListSystemDhcpServer(params *models.CmdbRequestParams) (*[]models.SystemDhcpServer, error)
	CreateWirelessControllerInterController(payload *models.WirelessControllerInterController, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerInterController(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerInterController, error)
	UpdateWirelessControllerInterController(mkey string, payload *models.WirelessControllerInterController, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerInterController(mkey string, params *models.CmdbRequestParams) error
	CreateSwitchControllerSflow(payload *models.SwitchControllerSflow, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerSflow(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerSflow, error)
	UpdateSwitchControllerSflow(mkey string, payload *models.SwitchControllerSflow, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerSflow(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallVip6(payload *models.FirewallVip6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallVip6(mkey string, params *models.CmdbRequestParams) (*models.FirewallVip6, error)
	UpdateFirewallVip6(mkey string, payload *models.FirewallVip6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallVip6(mkey string, params *models.CmdbRequestParams) error
	ListFirewallVip6(params *models.CmdbRequestParams) (*[]models.FirewallVip6, error)
	CreateFirewallProfileGroup(payload *models.FirewallProfileGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallProfileGroup(mkey string, params *models.CmdbRequestParams) (*models.FirewallProfileGroup, error)
	UpdateFirewallProfileGroup(mkey string, payload *models.FirewallProfileGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallProfileGroup(mkey string, params *models.CmdbRequestParams) error
	ListFirewallProfileGroup(params *models.CmdbRequestParams) (*[]models.FirewallProfileGroup, error)
	CreateDlpFilepattern(payload *models.DlpFilepattern, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadDlpFilepattern(mkey string, params *models.CmdbRequestParams) (*models.DlpFilepattern, error)
	UpdateDlpFilepattern(mkey string, payload *models.DlpFilepattern, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteDlpFilepattern(mkey string, params *models.CmdbRequestParams) error
	ListDlpFilepattern(params *models.CmdbRequestParams) (*[]models.DlpFilepattern, error)
	CreateAntivirusHeuristic(payload *models.AntivirusHeuristic, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadAntivirusHeuristic(mkey string, params *models.CmdbRequestParams) (*models.AntivirusHeuristic, error)
	UpdateAntivirusHeuristic(mkey string, payload *models.AntivirusHeuristic, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteAntivirusHeuristic(mkey string, params *models.CmdbRequestParams) error
	CreateSystemAlias(payload *models.SystemAlias, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemAlias(mkey string, params *models.CmdbRequestParams) (*models.SystemAlias, error)
	UpdateSystemAlias(mkey string, payload *models.SystemAlias, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemAlias(mkey string, params *models.CmdbRequestParams) error
	ListSystemAlias(params *models.CmdbRequestParams) (*[]models.SystemAlias, error)
	CreateFirewallSshLocalCa(payload *models.FirewallSshLocalCa, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallSshLocalCa(mkey string, params *models.CmdbRequestParams) (*models.FirewallSshLocalCa, error)
	UpdateFirewallSshLocalCa(mkey string, payload *models.FirewallSshLocalCa, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallSshLocalCa(mkey string, params *models.CmdbRequestParams) error
	ListFirewallSshLocalCa(params *models.CmdbRequestParams) (*[]models.FirewallSshLocalCa, error)
	CreateFirewallInterfacePolicy6(payload *models.FirewallInterfacePolicy6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallInterfacePolicy6(mkey string, params *models.CmdbRequestParams) (*models.FirewallInterfacePolicy6, error)
	UpdateFirewallInterfacePolicy6(mkey string, payload *models.FirewallInterfacePolicy6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallInterfacePolicy6(mkey string, params *models.CmdbRequestParams) error
	ListFirewallInterfacePolicy6(params *models.CmdbRequestParams) (*[]models.FirewallInterfacePolicy6, error)
	CreateVpnSslWebUserGroupBookmark(payload *models.VpnSslWebUserGroupBookmark, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnSslWebUserGroupBookmark(mkey string, params *models.CmdbRequestParams) (*models.VpnSslWebUserGroupBookmark, error)
	UpdateVpnSslWebUserGroupBookmark(mkey string, payload *models.VpnSslWebUserGroupBookmark, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnSslWebUserGroupBookmark(mkey string, params *models.CmdbRequestParams) error
	ListVpnSslWebUserGroupBookmark(params *models.CmdbRequestParams) (*[]models.VpnSslWebUserGroupBookmark, error)
	CreateLogSyslogd3Setting(payload *models.LogSyslogd3Setting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogSyslogd3Setting(mkey string, params *models.CmdbRequestParams) (*models.LogSyslogd3Setting, error)
	UpdateLogSyslogd3Setting(mkey string, payload *models.LogSyslogd3Setting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogSyslogd3Setting(mkey string, params *models.CmdbRequestParams) error
	CreateSwitchControllerLocation(payload *models.SwitchControllerLocation, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerLocation(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerLocation, error)
	UpdateSwitchControllerLocation(mkey string, payload *models.SwitchControllerLocation, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerLocation(mkey string, params *models.CmdbRequestParams) error
	ListSwitchControllerLocation(params *models.CmdbRequestParams) (*[]models.SwitchControllerLocation, error)
	CreateFirewallSslSetting(payload *models.FirewallSslSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallSslSetting(mkey string, params *models.CmdbRequestParams) (*models.FirewallSslSetting, error)
	UpdateFirewallSslSetting(mkey string, payload *models.FirewallSslSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallSslSetting(mkey string, params *models.CmdbRequestParams) error
	CreateVpnIpsecPhase2Interface(payload *models.VpnIpsecPhase2Interface, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnIpsecPhase2Interface(mkey string, params *models.CmdbRequestParams) (*models.VpnIpsecPhase2Interface, error)
	UpdateVpnIpsecPhase2Interface(mkey string, payload *models.VpnIpsecPhase2Interface, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnIpsecPhase2Interface(mkey string, params *models.CmdbRequestParams) error
	ListVpnIpsecPhase2Interface(params *models.CmdbRequestParams) (*[]models.VpnIpsecPhase2Interface, error)
	CreateVpnStatusL2tp(payload *models.VpnStatusL2tp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnStatusL2tp(mkey string, params *models.CmdbRequestParams) (*models.VpnStatusL2tp, error)
	UpdateVpnStatusL2tp(mkey string, payload *models.VpnStatusL2tp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnStatusL2tp(mkey string, params *models.CmdbRequestParams) error
	CreateCertificateLocal(payload *models.CertificateLocal, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadCertificateLocal(mkey string, params *models.CmdbRequestParams) (*models.CertificateLocal, error)
	UpdateCertificateLocal(mkey string, payload *models.CertificateLocal, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteCertificateLocal(mkey string, params *models.CmdbRequestParams) error
	ListCertificateLocal(params *models.CmdbRequestParams) (*[]models.CertificateLocal, error)
	CreateApplicationGroup(payload *models.ApplicationGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadApplicationGroup(mkey string, params *models.CmdbRequestParams) (*models.ApplicationGroup, error)
	UpdateApplicationGroup(mkey string, payload *models.ApplicationGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteApplicationGroup(mkey string, params *models.CmdbRequestParams) error
	ListApplicationGroup(params *models.CmdbRequestParams) (*[]models.ApplicationGroup, error)
	CreateMonitoringNpuHpe(payload *models.MonitoringNpuHpe, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadMonitoringNpuHpe(mkey string, params *models.CmdbRequestParams) (*models.MonitoringNpuHpe, error)
	UpdateMonitoringNpuHpe(mkey string, payload *models.MonitoringNpuHpe, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteMonitoringNpuHpe(mkey string, params *models.CmdbRequestParams) error
	CreateSwitchControllerSwitchGroup(payload *models.SwitchControllerSwitchGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerSwitchGroup(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerSwitchGroup, error)
	UpdateSwitchControllerSwitchGroup(mkey string, payload *models.SwitchControllerSwitchGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerSwitchGroup(mkey string, params *models.CmdbRequestParams) error
	ListSwitchControllerSwitchGroup(params *models.CmdbRequestParams) (*[]models.SwitchControllerSwitchGroup, error)
	CreateWirelessControllerSnmp(payload *models.WirelessControllerSnmp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerSnmp(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerSnmp, error)
	UpdateWirelessControllerSnmp(mkey string, payload *models.WirelessControllerSnmp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerSnmp(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerApStatus(payload *models.WirelessControllerApStatus, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerApStatus(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerApStatus, error)
	UpdateWirelessControllerApStatus(mkey string, payload *models.WirelessControllerApStatus, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerApStatus(mkey string, params *models.CmdbRequestParams) error
	ListWirelessControllerApStatus(params *models.CmdbRequestParams) (*[]models.WirelessControllerApStatus, error)
	CreateSwitchControllerSecurityPolicyLocalAccess(payload *models.SwitchControllerSecurityPolicyLocalAccess, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerSecurityPolicyLocalAccess(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerSecurityPolicyLocalAccess, error)
	UpdateSwitchControllerSecurityPolicyLocalAccess(mkey string, payload *models.SwitchControllerSecurityPolicyLocalAccess, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerSecurityPolicyLocalAccess(mkey string, params *models.CmdbRequestParams) error
	ListSwitchControllerSecurityPolicyLocalAccess(params *models.CmdbRequestParams) (*[]models.SwitchControllerSecurityPolicyLocalAccess, error)
	CreateFtpProxyExplicit(payload *models.FtpProxyExplicit, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFtpProxyExplicit(mkey string, params *models.CmdbRequestParams) (*models.FtpProxyExplicit, error)
	UpdateFtpProxyExplicit(mkey string, payload *models.FtpProxyExplicit, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFtpProxyExplicit(mkey string, params *models.CmdbRequestParams) error
	CreateLogFortianalyzerCloudFilter(payload *models.LogFortianalyzerCloudFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogFortianalyzerCloudFilter(mkey string, params *models.CmdbRequestParams) (*models.LogFortianalyzerCloudFilter, error)
	UpdateLogFortianalyzerCloudFilter(mkey string, payload *models.LogFortianalyzerCloudFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogFortianalyzerCloudFilter(mkey string, params *models.CmdbRequestParams) error
	CreateUserPeer(payload *models.UserPeer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadUserPeer(mkey string, params *models.CmdbRequestParams) (*models.UserPeer, error)
	UpdateUserPeer(mkey string, payload *models.UserPeer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteUserPeer(mkey string, params *models.CmdbRequestParams) error
	ListUserPeer(params *models.CmdbRequestParams) (*[]models.UserPeer, error)
	CreateExtenderControllerExtender(payload *models.ExtenderControllerExtender, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadExtenderControllerExtender(mkey string, params *models.CmdbRequestParams) (*models.ExtenderControllerExtender, error)
	UpdateExtenderControllerExtender(mkey string, payload *models.ExtenderControllerExtender, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteExtenderControllerExtender(mkey string, params *models.CmdbRequestParams) error
	ListExtenderControllerExtender(params *models.CmdbRequestParams) (*[]models.ExtenderControllerExtender, error)
	CreateSystemFortiguard(payload *models.SystemFortiguard, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemFortiguard(mkey string, params *models.CmdbRequestParams) (*models.SystemFortiguard, error)
	UpdateSystemFortiguard(mkey string, payload *models.SystemFortiguard, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemFortiguard(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallProxyAddress(payload *models.FirewallProxyAddress, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallProxyAddress(mkey string, params *models.CmdbRequestParams) (*models.FirewallProxyAddress, error)
	UpdateFirewallProxyAddress(mkey string, payload *models.FirewallProxyAddress, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallProxyAddress(mkey string, params *models.CmdbRequestParams) error
	ListFirewallProxyAddress(params *models.CmdbRequestParams) (*[]models.FirewallProxyAddress, error)
	CreateLogSyslogd4Setting(payload *models.LogSyslogd4Setting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogSyslogd4Setting(mkey string, params *models.CmdbRequestParams) (*models.LogSyslogd4Setting, error)
	UpdateLogSyslogd4Setting(mkey string, payload *models.LogSyslogd4Setting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogSyslogd4Setting(mkey string, params *models.CmdbRequestParams) error
	CreateSwitchControllerGlobal(payload *models.SwitchControllerGlobal, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerGlobal(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerGlobal, error)
	UpdateSwitchControllerGlobal(mkey string, payload *models.SwitchControllerGlobal, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerGlobal(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerWtpStatus(payload *models.WirelessControllerWtpStatus, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerWtpStatus(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerWtpStatus, error)
	UpdateWirelessControllerWtpStatus(mkey string, payload *models.WirelessControllerWtpStatus, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerWtpStatus(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallScheduleGroup(payload *models.FirewallScheduleGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallScheduleGroup(mkey string, params *models.CmdbRequestParams) (*models.FirewallScheduleGroup, error)
	UpdateFirewallScheduleGroup(mkey string, payload *models.FirewallScheduleGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallScheduleGroup(mkey string, params *models.CmdbRequestParams) error
	ListFirewallScheduleGroup(params *models.CmdbRequestParams) (*[]models.FirewallScheduleGroup, error)
	CreateLogNullDeviceSetting(payload *models.LogNullDeviceSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogNullDeviceSetting(mkey string, params *models.CmdbRequestParams) (*models.LogNullDeviceSetting, error)
	UpdateLogNullDeviceSetting(mkey string, payload *models.LogNullDeviceSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogNullDeviceSetting(mkey string, params *models.CmdbRequestParams) error
	CreateSystemIpsUrlfilterDns(payload *models.SystemIpsUrlfilterDns, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemIpsUrlfilterDns(mkey string, params *models.CmdbRequestParams) (*models.SystemIpsUrlfilterDns, error)
	UpdateSystemIpsUrlfilterDns(mkey string, payload *models.SystemIpsUrlfilterDns, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemIpsUrlfilterDns(mkey string, params *models.CmdbRequestParams) error
	ListSystemIpsUrlfilterDns(params *models.CmdbRequestParams) (*[]models.SystemIpsUrlfilterDns, error)
	CreateSystemSessionInfoExpectation(payload *models.SystemSessionInfoExpectation, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemSessionInfoExpectation(mkey string, params *models.CmdbRequestParams) (*models.SystemSessionInfoExpectation, error)
	UpdateSystemSessionInfoExpectation(mkey string, payload *models.SystemSessionInfoExpectation, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemSessionInfoExpectation(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerTimers(payload *models.WirelessControllerTimers, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerTimers(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerTimers, error)
	UpdateWirelessControllerTimers(mkey string, payload *models.WirelessControllerTimers, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerTimers(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerUtmProfile(payload *models.WirelessControllerUtmProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerUtmProfile(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerUtmProfile, error)
	UpdateWirelessControllerUtmProfile(mkey string, payload *models.WirelessControllerUtmProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerUtmProfile(mkey string, params *models.CmdbRequestParams) error
	ListWirelessControllerUtmProfile(params *models.CmdbRequestParams) (*[]models.WirelessControllerUtmProfile, error)
	CreateFirewallInternetServiceList(payload *models.FirewallInternetServiceList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallInternetServiceList(mkey string, params *models.CmdbRequestParams) (*models.FirewallInternetServiceList, error)
	UpdateFirewallInternetServiceList(mkey string, payload *models.FirewallInternetServiceList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallInternetServiceList(mkey string, params *models.CmdbRequestParams) error
	ListFirewallInternetServiceList(params *models.CmdbRequestParams) (*[]models.FirewallInternetServiceList, error)
	CreateFirewallDnstranslation(payload *models.FirewallDnstranslation, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallDnstranslation(mkey string, params *models.CmdbRequestParams) (*models.FirewallDnstranslation, error)
	UpdateFirewallDnstranslation(mkey string, payload *models.FirewallDnstranslation, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallDnstranslation(mkey string, params *models.CmdbRequestParams) error
	ListFirewallDnstranslation(params *models.CmdbRequestParams) (*[]models.FirewallDnstranslation, error)
	CreateWebfilterProfile(payload *models.WebfilterProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWebfilterProfile(mkey string, params *models.CmdbRequestParams) (*models.WebfilterProfile, error)
	UpdateWebfilterProfile(mkey string, payload *models.WebfilterProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWebfilterProfile(mkey string, params *models.CmdbRequestParams) error
	ListWebfilterProfile(params *models.CmdbRequestParams) (*[]models.WebfilterProfile, error)
	CreateWebfilterCategories(payload *models.WebfilterCategories, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWebfilterCategories(mkey string, params *models.CmdbRequestParams) (*models.WebfilterCategories, error)
	UpdateWebfilterCategories(mkey string, payload *models.WebfilterCategories, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWebfilterCategories(mkey string, params *models.CmdbRequestParams) error
	CreateSystemMobileTunnel(payload *models.SystemMobileTunnel, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemMobileTunnel(mkey string, params *models.CmdbRequestParams) (*models.SystemMobileTunnel, error)
	UpdateSystemMobileTunnel(mkey string, payload *models.SystemMobileTunnel, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemMobileTunnel(mkey string, params *models.CmdbRequestParams) error
	ListSystemMobileTunnel(params *models.CmdbRequestParams) (*[]models.SystemMobileTunnel, error)
	CreateEmailfilterBlockAllowList(payload *models.EmailfilterBlockAllowList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadEmailfilterBlockAllowList(mkey string, params *models.CmdbRequestParams) (*models.EmailfilterBlockAllowList, error)
	UpdateEmailfilterBlockAllowList(mkey string, payload *models.EmailfilterBlockAllowList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteEmailfilterBlockAllowList(mkey string, params *models.CmdbRequestParams) error
	ListEmailfilterBlockAllowList(params *models.CmdbRequestParams) (*[]models.EmailfilterBlockAllowList, error)
	CreateUserPasswordPolicy(payload *models.UserPasswordPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadUserPasswordPolicy(mkey string, params *models.CmdbRequestParams) (*models.UserPasswordPolicy, error)
	UpdateUserPasswordPolicy(mkey string, payload *models.UserPasswordPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteUserPasswordPolicy(mkey string, params *models.CmdbRequestParams) error
	ListUserPasswordPolicy(params *models.CmdbRequestParams) (*[]models.UserPasswordPolicy, error)
	CreateRouterPolicy(payload *models.RouterPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterPolicy(mkey string, params *models.CmdbRequestParams) (*models.RouterPolicy, error)
	UpdateRouterPolicy(mkey string, payload *models.RouterPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterPolicy(mkey string, params *models.CmdbRequestParams) error
	ListRouterPolicy(params *models.CmdbRequestParams) (*[]models.RouterPolicy, error)
	CreateSctpFilterProfile(payload *models.SctpFilterProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSctpFilterProfile(mkey string, params *models.CmdbRequestParams) (*models.SctpFilterProfile, error)
	UpdateSctpFilterProfile(mkey string, payload *models.SctpFilterProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSctpFilterProfile(mkey string, params *models.CmdbRequestParams) error
	ListSctpFilterProfile(params *models.CmdbRequestParams) (*[]models.SctpFilterProfile, error)
	CreateLogSyslogd3OverrideSetting(payload *models.LogSyslogd3OverrideSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogSyslogd3OverrideSetting(mkey string, params *models.CmdbRequestParams) (*models.LogSyslogd3OverrideSetting, error)
	UpdateLogSyslogd3OverrideSetting(mkey string, payload *models.LogSyslogd3OverrideSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogSyslogd3OverrideSetting(mkey string, params *models.CmdbRequestParams) error
	CreateSystemFortiguardLogService(payload *models.SystemFortiguardLogService, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemFortiguardLogService(mkey string, params *models.CmdbRequestParams) (*models.SystemFortiguardLogService, error)
	UpdateSystemFortiguardLogService(mkey string, payload *models.SystemFortiguardLogService, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemFortiguardLogService(mkey string, params *models.CmdbRequestParams) error
	CreateSystemGreTunnel(payload *models.SystemGreTunnel, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemGreTunnel(mkey string, params *models.CmdbRequestParams) (*models.SystemGreTunnel, error)
	UpdateSystemGreTunnel(mkey string, payload *models.SystemGreTunnel, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemGreTunnel(mkey string, params *models.CmdbRequestParams) error
	ListSystemGreTunnel(params *models.CmdbRequestParams) (*[]models.SystemGreTunnel, error)
	CreateSystemSpeedTestSchedule(payload *models.SystemSpeedTestSchedule, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemSpeedTestSchedule(mkey string, params *models.CmdbRequestParams) (*models.SystemSpeedTestSchedule, error)
	UpdateSystemSpeedTestSchedule(mkey string, payload *models.SystemSpeedTestSchedule, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemSpeedTestSchedule(mkey string, params *models.CmdbRequestParams) error
	ListSystemSpeedTestSchedule(params *models.CmdbRequestParams) (*[]models.SystemSpeedTestSchedule, error)
	CreateSystemCsf(payload *models.SystemCsf, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemCsf(mkey string, params *models.CmdbRequestParams) (*models.SystemCsf, error)
	UpdateSystemCsf(mkey string, payload *models.SystemCsf, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemCsf(mkey string, params *models.CmdbRequestParams) error
	CreateSwitchControllerSnmpTrapThreshold(payload *models.SwitchControllerSnmpTrapThreshold, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerSnmpTrapThreshold(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerSnmpTrapThreshold, error)
	UpdateSwitchControllerSnmpTrapThreshold(mkey string, payload *models.SwitchControllerSnmpTrapThreshold, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerSnmpTrapThreshold(mkey string, params *models.CmdbRequestParams) error
	CreateLogFortianalyzer2OverrideSetting(payload *models.LogFortianalyzer2OverrideSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogFortianalyzer2OverrideSetting(mkey string, params *models.CmdbRequestParams) (*models.LogFortianalyzer2OverrideSetting, error)
	UpdateLogFortianalyzer2OverrideSetting(mkey string, payload *models.LogFortianalyzer2OverrideSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogFortianalyzer2OverrideSetting(mkey string, params *models.CmdbRequestParams) error
	CreateSystemReplacemsgFtp(payload *models.SystemReplacemsgFtp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemReplacemsgFtp(mkey string, params *models.CmdbRequestParams) (*models.SystemReplacemsgFtp, error)
	UpdateSystemReplacemsgFtp(mkey string, payload *models.SystemReplacemsgFtp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemReplacemsgFtp(mkey string, params *models.CmdbRequestParams) error
	ListSystemReplacemsgFtp(params *models.CmdbRequestParams) (*[]models.SystemReplacemsgFtp, error)
	CreateFirewallSshLocalKey(payload *models.FirewallSshLocalKey, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallSshLocalKey(mkey string, params *models.CmdbRequestParams) (*models.FirewallSshLocalKey, error)
	UpdateFirewallSshLocalKey(mkey string, payload *models.FirewallSshLocalKey, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallSshLocalKey(mkey string, params *models.CmdbRequestParams) error
	ListFirewallSshLocalKey(params *models.CmdbRequestParams) (*[]models.FirewallSshLocalKey, error)
	CreateVpnIpsecForticlient(payload *models.VpnIpsecForticlient, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnIpsecForticlient(mkey string, params *models.CmdbRequestParams) (*models.VpnIpsecForticlient, error)
	UpdateVpnIpsecForticlient(mkey string, payload *models.VpnIpsecForticlient, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnIpsecForticlient(mkey string, params *models.CmdbRequestParams) error
	ListVpnIpsecForticlient(params *models.CmdbRequestParams) (*[]models.VpnIpsecForticlient, error)
	CreateWebfilterStatus(payload *models.WebfilterStatus, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWebfilterStatus(mkey string, params *models.CmdbRequestParams) (*models.WebfilterStatus, error)
	UpdateWebfilterStatus(mkey string, payload *models.WebfilterStatus, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWebfilterStatus(mkey string, params *models.CmdbRequestParams) error
	CreateSystem3GModemCustom(payload *models.System3GModemCustom, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystem3GModemCustom(mkey string, params *models.CmdbRequestParams) (*models.System3GModemCustom, error)
	UpdateSystem3GModemCustom(mkey string, payload *models.System3GModemCustom, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystem3GModemCustom(mkey string, params *models.CmdbRequestParams) error
	ListSystem3GModemCustom(params *models.CmdbRequestParams) (*[]models.System3GModemCustom, error)
	CreateSwitchControllerTrafficSniffer(payload *models.SwitchControllerTrafficSniffer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerTrafficSniffer(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerTrafficSniffer, error)
	UpdateSwitchControllerTrafficSniffer(mkey string, payload *models.SwitchControllerTrafficSniffer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerTrafficSniffer(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallIpropeList(payload *models.FirewallIpropeList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallIpropeList(mkey string, params *models.CmdbRequestParams) (*models.FirewallIpropeList, error)
	UpdateFirewallIpropeList(mkey string, payload *models.FirewallIpropeList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallIpropeList(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerAddress(payload *models.WirelessControllerAddress, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerAddress(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerAddress, error)
	UpdateWirelessControllerAddress(mkey string, payload *models.WirelessControllerAddress, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerAddress(mkey string, params *models.CmdbRequestParams) error
	ListWirelessControllerAddress(params *models.CmdbRequestParams) (*[]models.WirelessControllerAddress, error)
	CreateSystemPppoeInterface(payload *models.SystemPppoeInterface, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemPppoeInterface(mkey string, params *models.CmdbRequestParams) (*models.SystemPppoeInterface, error)
	UpdateSystemPppoeInterface(mkey string, payload *models.SystemPppoeInterface, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemPppoeInterface(mkey string, params *models.CmdbRequestParams) error
	ListSystemPppoeInterface(params *models.CmdbRequestParams) (*[]models.SystemPppoeInterface, error)
	CreateWirelessControllerHotspot20Icon(payload *models.WirelessControllerHotspot20Icon, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerHotspot20Icon(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerHotspot20Icon, error)
	UpdateWirelessControllerHotspot20Icon(mkey string, payload *models.WirelessControllerHotspot20Icon, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerHotspot20Icon(mkey string, params *models.CmdbRequestParams) error
	ListWirelessControllerHotspot20Icon(params *models.CmdbRequestParams) (*[]models.WirelessControllerHotspot20Icon, error)
	CreateFirewallDecryptedTrafficMirror(payload *models.FirewallDecryptedTrafficMirror, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallDecryptedTrafficMirror(mkey string, params *models.CmdbRequestParams) (*models.FirewallDecryptedTrafficMirror, error)
	UpdateFirewallDecryptedTrafficMirror(mkey string, payload *models.FirewallDecryptedTrafficMirror, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallDecryptedTrafficMirror(mkey string, params *models.CmdbRequestParams) error
	ListFirewallDecryptedTrafficMirror(params *models.CmdbRequestParams) (*[]models.FirewallDecryptedTrafficMirror, error)
	CreateFirewallInterfacePolicy(payload *models.FirewallInterfacePolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallInterfacePolicy(mkey string, params *models.CmdbRequestParams) (*models.FirewallInterfacePolicy, error)
	UpdateFirewallInterfacePolicy(mkey string, payload *models.FirewallInterfacePolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallInterfacePolicy(mkey string, params *models.CmdbRequestParams) error
	ListFirewallInterfacePolicy(params *models.CmdbRequestParams) (*[]models.FirewallInterfacePolicy, error)
	CreateFirewallCentralSnatMap(payload *models.FirewallCentralSnatMap, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallCentralSnatMap(mkey string, params *models.CmdbRequestParams) (*models.FirewallCentralSnatMap, error)
	UpdateFirewallCentralSnatMap(mkey string, payload *models.FirewallCentralSnatMap, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallCentralSnatMap(mkey string, params *models.CmdbRequestParams) error
	ListFirewallCentralSnatMap(params *models.CmdbRequestParams) (*[]models.FirewallCentralSnatMap, error)
	CreateUserExchange(payload *models.UserExchange, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadUserExchange(mkey string, params *models.CmdbRequestParams) (*models.UserExchange, error)
	UpdateUserExchange(mkey string, payload *models.UserExchange, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteUserExchange(mkey string, params *models.CmdbRequestParams) error
	ListUserExchange(params *models.CmdbRequestParams) (*[]models.UserExchange, error)
	CreateSystemNpu(payload *models.SystemNpu, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemNpu(mkey string, params *models.CmdbRequestParams) (*models.SystemNpu, error)
	UpdateSystemNpu(mkey string, payload *models.SystemNpu, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemNpu(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerHotspot20AnqpIpAddressType(payload *models.WirelessControllerHotspot20AnqpIpAddressType, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerHotspot20AnqpIpAddressType(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerHotspot20AnqpIpAddressType, error)
	UpdateWirelessControllerHotspot20AnqpIpAddressType(mkey string, payload *models.WirelessControllerHotspot20AnqpIpAddressType, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerHotspot20AnqpIpAddressType(mkey string, params *models.CmdbRequestParams) error
	ListWirelessControllerHotspot20AnqpIpAddressType(params *models.CmdbRequestParams) (*[]models.WirelessControllerHotspot20AnqpIpAddressType, error)
	CreateFirewallSshSetting(payload *models.FirewallSshSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallSshSetting(mkey string, params *models.CmdbRequestParams) (*models.FirewallSshSetting, error)
	UpdateFirewallSshSetting(mkey string, payload *models.FirewallSshSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallSshSetting(mkey string, params *models.CmdbRequestParams) error
	CreateWafMainClass(payload *models.WafMainClass, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWafMainClass(mkey string, params *models.CmdbRequestParams) (*models.WafMainClass, error)
	UpdateWafMainClass(mkey string, payload *models.WafMainClass, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWafMainClass(mkey string, params *models.CmdbRequestParams) error
	ListWafMainClass(params *models.CmdbRequestParams) (*[]models.WafMainClass, error)
	CreateFirewallIppool_grp(payload *models.FirewallIppool_grp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallIppool_grp(mkey string, params *models.CmdbRequestParams) (*models.FirewallIppool_grp, error)
	UpdateFirewallIppool_grp(mkey string, payload *models.FirewallIppool_grp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallIppool_grp(mkey string, params *models.CmdbRequestParams) error
	ListFirewallIppool_grp(params *models.CmdbRequestParams) (*[]models.FirewallIppool_grp, error)
	CreateSystemSsoForticloudAdmin(payload *models.SystemSsoForticloudAdmin, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemSsoForticloudAdmin(mkey string, params *models.CmdbRequestParams) (*models.SystemSsoForticloudAdmin, error)
	UpdateSystemSsoForticloudAdmin(mkey string, payload *models.SystemSsoForticloudAdmin, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemSsoForticloudAdmin(mkey string, params *models.CmdbRequestParams) error
	ListSystemSsoForticloudAdmin(params *models.CmdbRequestParams) (*[]models.SystemSsoForticloudAdmin, error)
	CreateSystemHa(payload *models.SystemHa, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemHa(mkey string, params *models.CmdbRequestParams) (*models.SystemHa, error)
	UpdateSystemHa(mkey string, payload *models.SystemHa, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemHa(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerHotspot20AnqpVenueUrl(payload *models.WirelessControllerHotspot20AnqpVenueUrl, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerHotspot20AnqpVenueUrl(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerHotspot20AnqpVenueUrl, error)
	UpdateWirelessControllerHotspot20AnqpVenueUrl(mkey string, payload *models.WirelessControllerHotspot20AnqpVenueUrl, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerHotspot20AnqpVenueUrl(mkey string, params *models.CmdbRequestParams) error
	ListWirelessControllerHotspot20AnqpVenueUrl(params *models.CmdbRequestParams) (*[]models.WirelessControllerHotspot20AnqpVenueUrl, error)
	CreateWirelessControllerLog(payload *models.WirelessControllerLog, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerLog(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerLog, error)
	UpdateWirelessControllerLog(mkey string, payload *models.WirelessControllerLog, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerLog(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerVapGroup(payload *models.WirelessControllerVapGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerVapGroup(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerVapGroup, error)
	UpdateWirelessControllerVapGroup(mkey string, payload *models.WirelessControllerVapGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerVapGroup(mkey string, params *models.CmdbRequestParams) error
	ListWirelessControllerVapGroup(params *models.CmdbRequestParams) (*[]models.WirelessControllerVapGroup, error)
	CreateSwitchControllerStpInstance(payload *models.SwitchControllerStpInstance, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerStpInstance(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerStpInstance, error)
	UpdateSwitchControllerStpInstance(mkey string, payload *models.SwitchControllerStpInstance, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerStpInstance(mkey string, params *models.CmdbRequestParams) error
	ListSwitchControllerStpInstance(params *models.CmdbRequestParams) (*[]models.SwitchControllerStpInstance, error)
	CreateFirewallInternetServiceExtension(payload *models.FirewallInternetServiceExtension, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallInternetServiceExtension(mkey string, params *models.CmdbRequestParams) (*models.FirewallInternetServiceExtension, error)
	UpdateFirewallInternetServiceExtension(mkey string, payload *models.FirewallInternetServiceExtension, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallInternetServiceExtension(mkey string, params *models.CmdbRequestParams) error
	ListFirewallInternetServiceExtension(params *models.CmdbRequestParams) (*[]models.FirewallInternetServiceExtension, error)
	CreateVpnSslWebUserBookmark(payload *models.VpnSslWebUserBookmark, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnSslWebUserBookmark(mkey string, params *models.CmdbRequestParams) (*models.VpnSslWebUserBookmark, error)
	UpdateVpnSslWebUserBookmark(mkey string, payload *models.VpnSslWebUserBookmark, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnSslWebUserBookmark(mkey string, params *models.CmdbRequestParams) error
	ListVpnSslWebUserBookmark(params *models.CmdbRequestParams) (*[]models.VpnSslWebUserBookmark, error)
	CreateSystemVdom(payload *models.SystemVdom, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemVdom(mkey string, params *models.CmdbRequestParams) (*models.SystemVdom, error)
	UpdateSystemVdom(mkey string, payload *models.SystemVdom, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemVdom(mkey string, params *models.CmdbRequestParams) error
	ListSystemVdom(params *models.CmdbRequestParams) (*[]models.SystemVdom, error)
	CreateUserKrbKeytab(payload *models.UserKrbKeytab, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadUserKrbKeytab(mkey string, params *models.CmdbRequestParams) (*models.UserKrbKeytab, error)
	UpdateUserKrbKeytab(mkey string, payload *models.UserKrbKeytab, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteUserKrbKeytab(mkey string, params *models.CmdbRequestParams) error
	ListUserKrbKeytab(params *models.CmdbRequestParams) (*[]models.UserKrbKeytab, error)
	CreateRouterBgp(payload *models.RouterBgp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterBgp(mkey string, params *models.CmdbRequestParams) (*models.RouterBgp, error)
	UpdateRouterBgp(mkey string, payload *models.RouterBgp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterBgp(mkey string, params *models.CmdbRequestParams) error
	CreateIpsDecoder(payload *models.IpsDecoder, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadIpsDecoder(mkey string, params *models.CmdbRequestParams) (*models.IpsDecoder, error)
	UpdateIpsDecoder(mkey string, payload *models.IpsDecoder, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteIpsDecoder(mkey string, params *models.CmdbRequestParams) error
	ListIpsDecoder(params *models.CmdbRequestParams) (*[]models.IpsDecoder, error)
	CreateWebfilterUrlfilter(payload *models.WebfilterUrlfilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWebfilterUrlfilter(mkey string, params *models.CmdbRequestParams) (*models.WebfilterUrlfilter, error)
	UpdateWebfilterUrlfilter(mkey string, payload *models.WebfilterUrlfilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWebfilterUrlfilter(mkey string, params *models.CmdbRequestParams) error
	ListWebfilterUrlfilter(params *models.CmdbRequestParams) (*[]models.WebfilterUrlfilter, error)
	CreateLogMemorySetting(payload *models.LogMemorySetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogMemorySetting(mkey string, params *models.CmdbRequestParams) (*models.LogMemorySetting, error)
	UpdateLogMemorySetting(mkey string, payload *models.LogMemorySetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogMemorySetting(mkey string, params *models.CmdbRequestParams) error
	CreateSwitchControllerAutoConfigCustom(payload *models.SwitchControllerAutoConfigCustom, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerAutoConfigCustom(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerAutoConfigCustom, error)
	UpdateSwitchControllerAutoConfigCustom(mkey string, payload *models.SwitchControllerAutoConfigCustom, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerAutoConfigCustom(mkey string, params *models.CmdbRequestParams) error
	ListSwitchControllerAutoConfigCustom(params *models.CmdbRequestParams) (*[]models.SwitchControllerAutoConfigCustom, error)
	CreateFirewallProxyPolicy(payload *models.FirewallProxyPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallProxyPolicy(mkey string, params *models.CmdbRequestParams) (*models.FirewallProxyPolicy, error)
	UpdateFirewallProxyPolicy(mkey string, payload *models.FirewallProxyPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallProxyPolicy(mkey string, params *models.CmdbRequestParams) error
	ListFirewallProxyPolicy(params *models.CmdbRequestParams) (*[]models.FirewallProxyPolicy, error)
	CreateVpnIpsecPhase1Interface(payload *models.VpnIpsecPhase1Interface, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnIpsecPhase1Interface(mkey string, params *models.CmdbRequestParams) (*models.VpnIpsecPhase1Interface, error)
	UpdateVpnIpsecPhase1Interface(mkey string, payload *models.VpnIpsecPhase1Interface, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnIpsecPhase1Interface(mkey string, params *models.CmdbRequestParams) error
	ListVpnIpsecPhase1Interface(params *models.CmdbRequestParams) (*[]models.VpnIpsecPhase1Interface, error)
	CreateLogFortianalyzer2Filter(payload *models.LogFortianalyzer2Filter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogFortianalyzer2Filter(mkey string, params *models.CmdbRequestParams) (*models.LogFortianalyzer2Filter, error)
	UpdateLogFortianalyzer2Filter(mkey string, payload *models.LogFortianalyzer2Filter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogFortianalyzer2Filter(mkey string, params *models.CmdbRequestParams) error
	CreateUserQuarantine(payload *models.UserQuarantine, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadUserQuarantine(mkey string, params *models.CmdbRequestParams) (*models.UserQuarantine, error)
	UpdateUserQuarantine(mkey string, payload *models.UserQuarantine, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteUserQuarantine(mkey string, params *models.CmdbRequestParams) error
	CreateRouterBfd(payload *models.RouterBfd, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterBfd(mkey string, params *models.CmdbRequestParams) (*models.RouterBfd, error)
	UpdateRouterBfd(mkey string, payload *models.RouterBfd, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterBfd(mkey string, params *models.CmdbRequestParams) error
	CreateHardwareNic(payload *models.HardwareNic, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadHardwareNic(mkey string, params *models.CmdbRequestParams) (*models.HardwareNic, error)
	UpdateHardwareNic(mkey string, payload *models.HardwareNic, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteHardwareNic(mkey string, params *models.CmdbRequestParams) error
	CreateSwitchControllerTrafficPolicy(payload *models.SwitchControllerTrafficPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerTrafficPolicy(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerTrafficPolicy, error)
	UpdateSwitchControllerTrafficPolicy(mkey string, payload *models.SwitchControllerTrafficPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerTrafficPolicy(mkey string, params *models.CmdbRequestParams) error
	ListSwitchControllerTrafficPolicy(params *models.CmdbRequestParams) (*[]models.SwitchControllerTrafficPolicy, error)
	CreateFirewallVip64(payload *models.FirewallVip64, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallVip64(mkey string, params *models.CmdbRequestParams) (*models.FirewallVip64, error)
	UpdateFirewallVip64(mkey string, payload *models.FirewallVip64, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallVip64(mkey string, params *models.CmdbRequestParams) error
	ListFirewallVip64(params *models.CmdbRequestParams) (*[]models.FirewallVip64, error)
	CreateSwitchControllerPortPolicy(payload *models.SwitchControllerPortPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerPortPolicy(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerPortPolicy, error)
	UpdateSwitchControllerPortPolicy(mkey string, payload *models.SwitchControllerPortPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerPortPolicy(mkey string, params *models.CmdbRequestParams) error
	ListSwitchControllerPortPolicy(params *models.CmdbRequestParams) (*[]models.SwitchControllerPortPolicy, error)
	CreateWirelessControllerSetting(payload *models.WirelessControllerSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerSetting(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerSetting, error)
	UpdateWirelessControllerSetting(mkey string, payload *models.WirelessControllerSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerSetting(mkey string, params *models.CmdbRequestParams) error
	CreateCertificateCrl(payload *models.CertificateCrl, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadCertificateCrl(mkey string, params *models.CmdbRequestParams) (*models.CertificateCrl, error)
	UpdateCertificateCrl(mkey string, payload *models.CertificateCrl, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteCertificateCrl(mkey string, params *models.CmdbRequestParams) error
	ListCertificateCrl(params *models.CmdbRequestParams) (*[]models.CertificateCrl, error)
	CreateFirewallVipgrp46(payload *models.FirewallVipgrp46, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallVipgrp46(mkey string, params *models.CmdbRequestParams) (*models.FirewallVipgrp46, error)
	UpdateFirewallVipgrp46(mkey string, payload *models.FirewallVipgrp46, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallVipgrp46(mkey string, params *models.CmdbRequestParams) error
	ListFirewallVipgrp46(params *models.CmdbRequestParams) (*[]models.FirewallVipgrp46, error)
	CreateSystemFipsCc(payload *models.SystemFipsCc, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemFipsCc(mkey string, params *models.CmdbRequestParams) (*models.SystemFipsCc, error)
	UpdateSystemFipsCc(mkey string, payload *models.SystemFipsCc, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemFipsCc(mkey string, params *models.CmdbRequestParams) error
	CreateExtenderModemStatus(payload *models.ExtenderModemStatus, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadExtenderModemStatus(mkey string, params *models.CmdbRequestParams) (*models.ExtenderModemStatus, error)
	UpdateExtenderModemStatus(mkey string, payload *models.ExtenderModemStatus, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteExtenderModemStatus(mkey string, params *models.CmdbRequestParams) error
	CreateSystemStartupErrorLog(payload *models.SystemStartupErrorLog, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemStartupErrorLog(mkey string, params *models.CmdbRequestParams) (*models.SystemStartupErrorLog, error)
	UpdateSystemStartupErrorLog(mkey string, payload *models.SystemStartupErrorLog, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemStartupErrorLog(mkey string, params *models.CmdbRequestParams) error
	CreateSystemIpsecAggregate(payload *models.SystemIpsecAggregate, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemIpsecAggregate(mkey string, params *models.CmdbRequestParams) (*models.SystemIpsecAggregate, error)
	UpdateSystemIpsecAggregate(mkey string, payload *models.SystemIpsecAggregate, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemIpsecAggregate(mkey string, params *models.CmdbRequestParams) error
	ListSystemIpsecAggregate(params *models.CmdbRequestParams) (*[]models.SystemIpsecAggregate, error)
	CreateFirewallInternetServiceSld(payload *models.FirewallInternetServiceSld, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallInternetServiceSld(mkey string, params *models.CmdbRequestParams) (*models.FirewallInternetServiceSld, error)
	UpdateFirewallInternetServiceSld(mkey string, payload *models.FirewallInternetServiceSld, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallInternetServiceSld(mkey string, params *models.CmdbRequestParams) error
	ListFirewallInternetServiceSld(params *models.CmdbRequestParams) (*[]models.FirewallInternetServiceSld, error)
	CreateRouterRip(payload *models.RouterRip, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterRip(mkey string, params *models.CmdbRequestParams) (*models.RouterRip, error)
	UpdateRouterRip(mkey string, payload *models.RouterRip, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterRip(mkey string, params *models.CmdbRequestParams) error
	CreateSystemTosBasedPriority(payload *models.SystemTosBasedPriority, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemTosBasedPriority(mkey string, params *models.CmdbRequestParams) (*models.SystemTosBasedPriority, error)
	UpdateSystemTosBasedPriority(mkey string, payload *models.SystemTosBasedPriority, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemTosBasedPriority(mkey string, params *models.CmdbRequestParams) error
	ListSystemTosBasedPriority(params *models.CmdbRequestParams) (*[]models.SystemTosBasedPriority, error)
	CreateVpnIpsecManualkey(payload *models.VpnIpsecManualkey, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnIpsecManualkey(mkey string, params *models.CmdbRequestParams) (*models.VpnIpsecManualkey, error)
	UpdateVpnIpsecManualkey(mkey string, payload *models.VpnIpsecManualkey, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnIpsecManualkey(mkey string, params *models.CmdbRequestParams) error
	ListVpnIpsecManualkey(params *models.CmdbRequestParams) (*[]models.VpnIpsecManualkey, error)
	CreateEmailfilterBword(payload *models.EmailfilterBword, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadEmailfilterBword(mkey string, params *models.CmdbRequestParams) (*models.EmailfilterBword, error)
	UpdateEmailfilterBword(mkey string, payload *models.EmailfilterBword, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteEmailfilterBword(mkey string, params *models.CmdbRequestParams) error
	ListEmailfilterBword(params *models.CmdbRequestParams) (*[]models.EmailfilterBword, error)
	CreateSystemSessionTtl(payload *models.SystemSessionTtl, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemSessionTtl(mkey string, params *models.CmdbRequestParams) (*models.SystemSessionTtl, error)
	UpdateSystemSessionTtl(mkey string, payload *models.SystemSessionTtl, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemSessionTtl(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerHotspot20H2qpTermsAndConditions(payload *models.WirelessControllerHotspot20H2qpTermsAndConditions, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerHotspot20H2qpTermsAndConditions(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerHotspot20H2qpTermsAndConditions, error)
	UpdateWirelessControllerHotspot20H2qpTermsAndConditions(mkey string, payload *models.WirelessControllerHotspot20H2qpTermsAndConditions, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerHotspot20H2qpTermsAndConditions(mkey string, params *models.CmdbRequestParams) error
	ListWirelessControllerHotspot20H2qpTermsAndConditions(params *models.CmdbRequestParams) (*[]models.WirelessControllerHotspot20H2qpTermsAndConditions, error)
	CreateFirewallAddress6Template(payload *models.FirewallAddress6Template, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallAddress6Template(mkey string, params *models.CmdbRequestParams) (*models.FirewallAddress6Template, error)
	UpdateFirewallAddress6Template(mkey string, payload *models.FirewallAddress6Template, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallAddress6Template(mkey string, params *models.CmdbRequestParams) error
	ListFirewallAddress6Template(params *models.CmdbRequestParams) (*[]models.FirewallAddress6Template, error)
	CreateFirewallInternetService(payload *models.FirewallInternetService, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallInternetService(mkey string, params *models.CmdbRequestParams) (*models.FirewallInternetService, error)
	UpdateFirewallInternetService(mkey string, payload *models.FirewallInternetService, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallInternetService(mkey string, params *models.CmdbRequestParams) error
	ListFirewallInternetService(params *models.CmdbRequestParams) (*[]models.FirewallInternetService, error)
	CreateFirewallAccessProxy6(payload *models.FirewallAccessProxy6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallAccessProxy6(mkey string, params *models.CmdbRequestParams) (*models.FirewallAccessProxy6, error)
	UpdateFirewallAccessProxy6(mkey string, payload *models.FirewallAccessProxy6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallAccessProxy6(mkey string, params *models.CmdbRequestParams) error
	ListFirewallAccessProxy6(params *models.CmdbRequestParams) (*[]models.FirewallAccessProxy6, error)
	CreateVpnIpsecConcentrator(payload *models.VpnIpsecConcentrator, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnIpsecConcentrator(mkey string, params *models.CmdbRequestParams) (*models.VpnIpsecConcentrator, error)
	UpdateVpnIpsecConcentrator(mkey string, payload *models.VpnIpsecConcentrator, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnIpsecConcentrator(mkey string, params *models.CmdbRequestParams) error
	ListVpnIpsecConcentrator(params *models.CmdbRequestParams) (*[]models.VpnIpsecConcentrator, error)
	CreateSystemReplacemsgTrafficQuota(payload *models.SystemReplacemsgTrafficQuota, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemReplacemsgTrafficQuota(mkey string, params *models.CmdbRequestParams) (*models.SystemReplacemsgTrafficQuota, error)
	UpdateSystemReplacemsgTrafficQuota(mkey string, payload *models.SystemReplacemsgTrafficQuota, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemReplacemsgTrafficQuota(mkey string, params *models.CmdbRequestParams) error
	ListSystemReplacemsgTrafficQuota(params *models.CmdbRequestParams) (*[]models.SystemReplacemsgTrafficQuota, error)
	CreateSystemSdwan(payload *models.SystemSdwan, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemSdwan(mkey string, params *models.CmdbRequestParams) (*models.SystemSdwan, error)
	UpdateSystemSdwan(mkey string, payload *models.SystemSdwan, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemSdwan(mkey string, params *models.CmdbRequestParams) error
	CreateSystemFederatedUpgrade(payload *models.SystemFederatedUpgrade, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemFederatedUpgrade(mkey string, params *models.CmdbRequestParams) (*models.SystemFederatedUpgrade, error)
	UpdateSystemFederatedUpgrade(mkey string, payload *models.SystemFederatedUpgrade, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemFederatedUpgrade(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallMulticastAddress(payload *models.FirewallMulticastAddress, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallMulticastAddress(mkey string, params *models.CmdbRequestParams) (*models.FirewallMulticastAddress, error)
	UpdateFirewallMulticastAddress(mkey string, payload *models.FirewallMulticastAddress, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallMulticastAddress(mkey string, params *models.CmdbRequestParams) error
	ListFirewallMulticastAddress(params *models.CmdbRequestParams) (*[]models.FirewallMulticastAddress, error)
	CreateFirewallWildcardFqdnCustom(payload *models.FirewallWildcardFqdnCustom, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallWildcardFqdnCustom(mkey string, params *models.CmdbRequestParams) (*models.FirewallWildcardFqdnCustom, error)
	UpdateFirewallWildcardFqdnCustom(mkey string, payload *models.FirewallWildcardFqdnCustom, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallWildcardFqdnCustom(mkey string, params *models.CmdbRequestParams) error
	ListFirewallWildcardFqdnCustom(params *models.CmdbRequestParams) (*[]models.FirewallWildcardFqdnCustom, error)
	CreateWebfilterContent(payload *models.WebfilterContent, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWebfilterContent(mkey string, params *models.CmdbRequestParams) (*models.WebfilterContent, error)
	UpdateWebfilterContent(mkey string, payload *models.WebfilterContent, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWebfilterContent(mkey string, params *models.CmdbRequestParams) error
	ListWebfilterContent(params *models.CmdbRequestParams) (*[]models.WebfilterContent, error)
	CreateLogWebtrendsSetting(payload *models.LogWebtrendsSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogWebtrendsSetting(mkey string, params *models.CmdbRequestParams) (*models.LogWebtrendsSetting, error)
	UpdateLogWebtrendsSetting(mkey string, payload *models.LogWebtrendsSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogWebtrendsSetting(mkey string, params *models.CmdbRequestParams) error
	CreateWafProfile(payload *models.WafProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWafProfile(mkey string, params *models.CmdbRequestParams) (*models.WafProfile, error)
	UpdateWafProfile(mkey string, payload *models.WafProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWafProfile(mkey string, params *models.CmdbRequestParams) error
	ListWafProfile(params *models.CmdbRequestParams) (*[]models.WafProfile, error)
	CreateSystemPasswordPolicyGuestAdmin(payload *models.SystemPasswordPolicyGuestAdmin, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemPasswordPolicyGuestAdmin(mkey string, params *models.CmdbRequestParams) (*models.SystemPasswordPolicyGuestAdmin, error)
	UpdateSystemPasswordPolicyGuestAdmin(mkey string, payload *models.SystemPasswordPolicyGuestAdmin, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemPasswordPolicyGuestAdmin(mkey string, params *models.CmdbRequestParams) error
	CreateEndpointControlFctems(payload *models.EndpointControlFctems, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadEndpointControlFctems(mkey string, params *models.CmdbRequestParams) (*models.EndpointControlFctems, error)
	UpdateEndpointControlFctems(mkey string, payload *models.EndpointControlFctems, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteEndpointControlFctems(mkey string, params *models.CmdbRequestParams) error
	ListEndpointControlFctems(params *models.CmdbRequestParams) (*[]models.EndpointControlFctems, error)
	CreateIpsRule(payload *models.IpsRule, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadIpsRule(mkey string, params *models.CmdbRequestParams) (*models.IpsRule, error)
	UpdateIpsRule(mkey string, payload *models.IpsRule, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteIpsRule(mkey string, params *models.CmdbRequestParams) error
	ListIpsRule(params *models.CmdbRequestParams) (*[]models.IpsRule, error)
	CreateWebProxyUrlMatch(payload *models.WebProxyUrlMatch, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWebProxyUrlMatch(mkey string, params *models.CmdbRequestParams) (*models.WebProxyUrlMatch, error)
	UpdateWebProxyUrlMatch(mkey string, payload *models.WebProxyUrlMatch, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWebProxyUrlMatch(mkey string, params *models.CmdbRequestParams) error
	ListWebProxyUrlMatch(params *models.CmdbRequestParams) (*[]models.WebProxyUrlMatch, error)
	CreateUserLocal(payload *models.UserLocal, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadUserLocal(mkey string, params *models.CmdbRequestParams) (*models.UserLocal, error)
	UpdateUserLocal(mkey string, payload *models.UserLocal, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteUserLocal(mkey string, params *models.CmdbRequestParams) error
	ListUserLocal(params *models.CmdbRequestParams) (*[]models.UserLocal, error)
	CreateRouterStatic6(payload *models.RouterStatic6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterStatic6(mkey string, params *models.CmdbRequestParams) (*models.RouterStatic6, error)
	UpdateRouterStatic6(mkey string, payload *models.RouterStatic6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterStatic6(mkey string, params *models.CmdbRequestParams) error
	ListRouterStatic6(params *models.CmdbRequestParams) (*[]models.RouterStatic6, error)
	CreateSystemSettings(payload *models.SystemSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemSettings(mkey string, params *models.CmdbRequestParams) (*models.SystemSettings, error)
	UpdateSystemSettings(mkey string, payload *models.SystemSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemSettings(mkey string, params *models.CmdbRequestParams) error
	CreateExtenderLteCarrierByMccMnc(payload *models.ExtenderLteCarrierByMccMnc, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadExtenderLteCarrierByMccMnc(mkey string, params *models.CmdbRequestParams) (*models.ExtenderLteCarrierByMccMnc, error)
	UpdateExtenderLteCarrierByMccMnc(mkey string, payload *models.ExtenderLteCarrierByMccMnc, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteExtenderLteCarrierByMccMnc(mkey string, params *models.CmdbRequestParams) error
	CreateFileFilterProfile(payload *models.FileFilterProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFileFilterProfile(mkey string, params *models.CmdbRequestParams) (*models.FileFilterProfile, error)
	UpdateFileFilterProfile(mkey string, payload *models.FileFilterProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFileFilterProfile(mkey string, params *models.CmdbRequestParams) error
	ListFileFilterProfile(params *models.CmdbRequestParams) (*[]models.FileFilterProfile, error)
	CreateSystemNetflow(payload *models.SystemNetflow, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemNetflow(mkey string, params *models.CmdbRequestParams) (*models.SystemNetflow, error)
	UpdateSystemNetflow(mkey string, payload *models.SystemNetflow, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemNetflow(mkey string, params *models.CmdbRequestParams) error
	CreateSystemReplacemsgMail(payload *models.SystemReplacemsgMail, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemReplacemsgMail(mkey string, params *models.CmdbRequestParams) (*models.SystemReplacemsgMail, error)
	UpdateSystemReplacemsgMail(mkey string, payload *models.SystemReplacemsgMail, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemReplacemsgMail(mkey string, params *models.CmdbRequestParams) error
	ListSystemReplacemsgMail(params *models.CmdbRequestParams) (*[]models.SystemReplacemsgMail, error)
	CreateSystemZone(payload *models.SystemZone, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemZone(mkey string, params *models.CmdbRequestParams) (*models.SystemZone, error)
	UpdateSystemZone(mkey string, payload *models.SystemZone, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemZone(mkey string, params *models.CmdbRequestParams) error
	ListSystemZone(params *models.CmdbRequestParams) (*[]models.SystemZone, error)
	CreateSystemAcme(payload *models.SystemAcme, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemAcme(mkey string, params *models.CmdbRequestParams) (*models.SystemAcme, error)
	UpdateSystemAcme(mkey string, payload *models.SystemAcme, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemAcme(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerRfAnalysis(payload *models.WirelessControllerRfAnalysis, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerRfAnalysis(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerRfAnalysis, error)
	UpdateWirelessControllerRfAnalysis(mkey string, payload *models.WirelessControllerRfAnalysis, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerRfAnalysis(mkey string, params *models.CmdbRequestParams) error
	CreateRouterIsis(payload *models.RouterIsis, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterIsis(mkey string, params *models.CmdbRequestParams) (*models.RouterIsis, error)
	UpdateRouterIsis(mkey string, payload *models.RouterIsis, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterIsis(mkey string, params *models.CmdbRequestParams) error
	CreateRouterMulticast6(payload *models.RouterMulticast6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterMulticast6(mkey string, params *models.CmdbRequestParams) (*models.RouterMulticast6, error)
	UpdateRouterMulticast6(mkey string, payload *models.RouterMulticast6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterMulticast6(mkey string, params *models.CmdbRequestParams) error
	CreateSystemGlobal(payload *models.SystemGlobal, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemGlobal(mkey string, params *models.CmdbRequestParams) (*models.SystemGlobal, error)
	UpdateSystemGlobal(mkey string, payload *models.SystemGlobal, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemGlobal(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallTrafficClass(payload *models.FirewallTrafficClass, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallTrafficClass(mkey string, params *models.CmdbRequestParams) (*models.FirewallTrafficClass, error)
	UpdateFirewallTrafficClass(mkey string, payload *models.FirewallTrafficClass, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallTrafficClass(mkey string, params *models.CmdbRequestParams) error
	ListFirewallTrafficClass(params *models.CmdbRequestParams) (*[]models.FirewallTrafficClass, error)
	CreateVpnCertificateSetting(payload *models.VpnCertificateSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnCertificateSetting(mkey string, params *models.CmdbRequestParams) (*models.VpnCertificateSetting, error)
	UpdateVpnCertificateSetting(mkey string, payload *models.VpnCertificateSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnCertificateSetting(mkey string, params *models.CmdbRequestParams) error
	CreateVpnStatusSslHwAccelerationStatus(payload *models.VpnStatusSslHwAccelerationStatus, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnStatusSslHwAccelerationStatus(mkey string, params *models.CmdbRequestParams) (*models.VpnStatusSslHwAccelerationStatus, error)
	UpdateVpnStatusSslHwAccelerationStatus(mkey string, payload *models.VpnStatusSslHwAccelerationStatus, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnStatusSslHwAccelerationStatus(mkey string, params *models.CmdbRequestParams) error
	CreateUserFssoPolling(payload *models.UserFssoPolling, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadUserFssoPolling(mkey string, params *models.CmdbRequestParams) (*models.UserFssoPolling, error)
	UpdateUserFssoPolling(mkey string, payload *models.UserFssoPolling, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteUserFssoPolling(mkey string, params *models.CmdbRequestParams) error
	ListUserFssoPolling(params *models.CmdbRequestParams) (*[]models.UserFssoPolling, error)
	CreateSystemWccp(payload *models.SystemWccp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemWccp(mkey string, params *models.CmdbRequestParams) (*models.SystemWccp, error)
	UpdateSystemWccp(mkey string, payload *models.SystemWccp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemWccp(mkey string, params *models.CmdbRequestParams) error
	ListSystemWccp(params *models.CmdbRequestParams) (*[]models.SystemWccp, error)
	CreateSystemDns64(payload *models.SystemDns64, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemDns64(mkey string, params *models.CmdbRequestParams) (*models.SystemDns64, error)
	UpdateSystemDns64(mkey string, payload *models.SystemDns64, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemDns64(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerMpskProfile(payload *models.WirelessControllerMpskProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerMpskProfile(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerMpskProfile, error)
	UpdateWirelessControllerMpskProfile(mkey string, payload *models.WirelessControllerMpskProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerMpskProfile(mkey string, params *models.CmdbRequestParams) error
	ListWirelessControllerMpskProfile(params *models.CmdbRequestParams) (*[]models.WirelessControllerMpskProfile, error)
	CreateSwitchControllerStormControlPolicy(payload *models.SwitchControllerStormControlPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerStormControlPolicy(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerStormControlPolicy, error)
	UpdateSwitchControllerStormControlPolicy(mkey string, payload *models.SwitchControllerStormControlPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerStormControlPolicy(mkey string, params *models.CmdbRequestParams) error
	ListSwitchControllerStormControlPolicy(params *models.CmdbRequestParams) (*[]models.SwitchControllerStormControlPolicy, error)
	CreateFirewallInternetServiceOwner(payload *models.FirewallInternetServiceOwner, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallInternetServiceOwner(mkey string, params *models.CmdbRequestParams) (*models.FirewallInternetServiceOwner, error)
	UpdateFirewallInternetServiceOwner(mkey string, payload *models.FirewallInternetServiceOwner, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallInternetServiceOwner(mkey string, params *models.CmdbRequestParams) error
	ListFirewallInternetServiceOwner(params *models.CmdbRequestParams) (*[]models.FirewallInternetServiceOwner, error)
	CreateRouterAccessList6(payload *models.RouterAccessList6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterAccessList6(mkey string, params *models.CmdbRequestParams) (*models.RouterAccessList6, error)
	UpdateRouterAccessList6(mkey string, payload *models.RouterAccessList6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterAccessList6(mkey string, params *models.CmdbRequestParams) error
	ListRouterAccessList6(params *models.CmdbRequestParams) (*[]models.RouterAccessList6, error)
	CreateSystemModem(payload *models.SystemModem, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemModem(mkey string, params *models.CmdbRequestParams) (*models.SystemModem, error)
	UpdateSystemModem(mkey string, payload *models.SystemModem, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemModem(mkey string, params *models.CmdbRequestParams) error
	CreateSystemSessionInfoTtl(payload *models.SystemSessionInfoTtl, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemSessionInfoTtl(mkey string, params *models.CmdbRequestParams) (*models.SystemSessionInfoTtl, error)
	UpdateSystemSessionInfoTtl(mkey string, payload *models.SystemSessionInfoTtl, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemSessionInfoTtl(mkey string, params *models.CmdbRequestParams) error
	CreateSystemAutomationTrigger(payload *models.SystemAutomationTrigger, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemAutomationTrigger(mkey string, params *models.CmdbRequestParams) (*models.SystemAutomationTrigger, error)
	UpdateSystemAutomationTrigger(mkey string, payload *models.SystemAutomationTrigger, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemAutomationTrigger(mkey string, params *models.CmdbRequestParams) error
	ListSystemAutomationTrigger(params *models.CmdbRequestParams) (*[]models.SystemAutomationTrigger, error)
	CreateAlertemailSetting(payload *models.AlertemailSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadAlertemailSetting(mkey string, params *models.CmdbRequestParams) (*models.AlertemailSetting, error)
	UpdateAlertemailSetting(mkey string, payload *models.AlertemailSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteAlertemailSetting(mkey string, params *models.CmdbRequestParams) error
	CreateExtenderSessionInfo(payload *models.ExtenderSessionInfo, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadExtenderSessionInfo(mkey string, params *models.CmdbRequestParams) (*models.ExtenderSessionInfo, error)
	UpdateExtenderSessionInfo(mkey string, payload *models.ExtenderSessionInfo, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteExtenderSessionInfo(mkey string, params *models.CmdbRequestParams) error
	CreateSystemArp(payload *models.SystemArp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemArp(mkey string, params *models.CmdbRequestParams) (*models.SystemArp, error)
	UpdateSystemArp(mkey string, payload *models.SystemArp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemArp(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallAuthPortal(payload *models.FirewallAuthPortal, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallAuthPortal(mkey string, params *models.CmdbRequestParams) (*models.FirewallAuthPortal, error)
	UpdateFirewallAuthPortal(mkey string, payload *models.FirewallAuthPortal, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallAuthPortal(mkey string, params *models.CmdbRequestParams) error
	CreateSystemSwitchInterface(payload *models.SystemSwitchInterface, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemSwitchInterface(mkey string, params *models.CmdbRequestParams) (*models.SystemSwitchInterface, error)
	UpdateSystemSwitchInterface(mkey string, payload *models.SystemSwitchInterface, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemSwitchInterface(mkey string, params *models.CmdbRequestParams) error
	ListSystemSwitchInterface(params *models.CmdbRequestParams) (*[]models.SystemSwitchInterface, error)
	CreateLogWebtrendsFilter(payload *models.LogWebtrendsFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogWebtrendsFilter(mkey string, params *models.CmdbRequestParams) (*models.LogWebtrendsFilter, error)
	UpdateLogWebtrendsFilter(mkey string, payload *models.LogWebtrendsFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogWebtrendsFilter(mkey string, params *models.CmdbRequestParams) error
	CreateLogFortianalyzerCloudOverrideSetting(payload *models.LogFortianalyzerCloudOverrideSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogFortianalyzerCloudOverrideSetting(mkey string, params *models.CmdbRequestParams) (*models.LogFortianalyzerCloudOverrideSetting, error)
	UpdateLogFortianalyzerCloudOverrideSetting(mkey string, payload *models.LogFortianalyzerCloudOverrideSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogFortianalyzerCloudOverrideSetting(mkey string, params *models.CmdbRequestParams) error
	CreateWafSignature(payload *models.WafSignature, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWafSignature(mkey string, params *models.CmdbRequestParams) (*models.WafSignature, error)
	UpdateWafSignature(mkey string, payload *models.WafSignature, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWafSignature(mkey string, params *models.CmdbRequestParams) error
	ListWafSignature(params *models.CmdbRequestParams) (*[]models.WafSignature, error)
	CreateAuthenticationSetting(payload *models.AuthenticationSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadAuthenticationSetting(mkey string, params *models.CmdbRequestParams) (*models.AuthenticationSetting, error)
	UpdateAuthenticationSetting(mkey string, payload *models.AuthenticationSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteAuthenticationSetting(mkey string, params *models.CmdbRequestParams) error
	CreateSystemDedicatedMgmt(payload *models.SystemDedicatedMgmt, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemDedicatedMgmt(mkey string, params *models.CmdbRequestParams) (*models.SystemDedicatedMgmt, error)
	UpdateSystemDedicatedMgmt(mkey string, payload *models.SystemDedicatedMgmt, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemDedicatedMgmt(mkey string, params *models.CmdbRequestParams) error
	CreateSwitchControllerVirtualPortPool(payload *models.SwitchControllerVirtualPortPool, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerVirtualPortPool(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerVirtualPortPool, error)
	UpdateSwitchControllerVirtualPortPool(mkey string, payload *models.SwitchControllerVirtualPortPool, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerVirtualPortPool(mkey string, params *models.CmdbRequestParams) error
	ListSwitchControllerVirtualPortPool(params *models.CmdbRequestParams) (*[]models.SwitchControllerVirtualPortPool, error)
	CreateSwitchControllerNetworkMonitorSettings(payload *models.SwitchControllerNetworkMonitorSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerNetworkMonitorSettings(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerNetworkMonitorSettings, error)
	UpdateSwitchControllerNetworkMonitorSettings(mkey string, payload *models.SwitchControllerNetworkMonitorSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerNetworkMonitorSettings(mkey string, params *models.CmdbRequestParams) error
	CreateVpnStatusSslList(payload *models.VpnStatusSslList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnStatusSslList(mkey string, params *models.CmdbRequestParams) (*models.VpnStatusSslList, error)
	UpdateVpnStatusSslList(mkey string, payload *models.VpnStatusSslList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnStatusSslList(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallPolicy64(payload *models.FirewallPolicy64, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallPolicy64(mkey string, params *models.CmdbRequestParams) (*models.FirewallPolicy64, error)
	UpdateFirewallPolicy64(mkey string, payload *models.FirewallPolicy64, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallPolicy64(mkey string, params *models.CmdbRequestParams) error
	ListFirewallPolicy64(params *models.CmdbRequestParams) (*[]models.FirewallPolicy64, error)
	CreateSystemVdomRadiusServer(payload *models.SystemVdomRadiusServer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemVdomRadiusServer(mkey string, params *models.CmdbRequestParams) (*models.SystemVdomRadiusServer, error)
	UpdateSystemVdomRadiusServer(mkey string, payload *models.SystemVdomRadiusServer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemVdomRadiusServer(mkey string, params *models.CmdbRequestParams) error
	ListSystemVdomRadiusServer(params *models.CmdbRequestParams) (*[]models.SystemVdomRadiusServer, error)
	CreateFirewallAccessProxyVirtualHost(payload *models.FirewallAccessProxyVirtualHost, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallAccessProxyVirtualHost(mkey string, params *models.CmdbRequestParams) (*models.FirewallAccessProxyVirtualHost, error)
	UpdateFirewallAccessProxyVirtualHost(mkey string, payload *models.FirewallAccessProxyVirtualHost, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallAccessProxyVirtualHost(mkey string, params *models.CmdbRequestParams) error
	ListFirewallAccessProxyVirtualHost(params *models.CmdbRequestParams) (*[]models.FirewallAccessProxyVirtualHost, error)
	CreateLogThreatWeight(payload *models.LogThreatWeight, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogThreatWeight(mkey string, params *models.CmdbRequestParams) (*models.LogThreatWeight, error)
	UpdateLogThreatWeight(mkey string, payload *models.LogThreatWeight, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogThreatWeight(mkey string, params *models.CmdbRequestParams) error
	CreateSystemSpeedTestServer(payload *models.SystemSpeedTestServer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemSpeedTestServer(mkey string, params *models.CmdbRequestParams) (*models.SystemSpeedTestServer, error)
	UpdateSystemSpeedTestServer(mkey string, payload *models.SystemSpeedTestServer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemSpeedTestServer(mkey string, params *models.CmdbRequestParams) error
	ListSystemSpeedTestServer(params *models.CmdbRequestParams) (*[]models.SystemSpeedTestServer, error)
	CreateSystemDscpBasedPriority(payload *models.SystemDscpBasedPriority, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemDscpBasedPriority(mkey string, params *models.CmdbRequestParams) (*models.SystemDscpBasedPriority, error)
	UpdateSystemDscpBasedPriority(mkey string, payload *models.SystemDscpBasedPriority, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemDscpBasedPriority(mkey string, params *models.CmdbRequestParams) error
	ListSystemDscpBasedPriority(params *models.CmdbRequestParams) (*[]models.SystemDscpBasedPriority, error)
	CreateSwitchControllerStormControl(payload *models.SwitchControllerStormControl, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerStormControl(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerStormControl, error)
	UpdateSwitchControllerStormControl(mkey string, payload *models.SwitchControllerStormControl, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerStormControl(mkey string, params *models.CmdbRequestParams) error
	CreateEmailfilterOptions(payload *models.EmailfilterOptions, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadEmailfilterOptions(mkey string, params *models.CmdbRequestParams) (*models.EmailfilterOptions, error)
	UpdateEmailfilterOptions(mkey string, payload *models.EmailfilterOptions, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteEmailfilterOptions(mkey string, params *models.CmdbRequestParams) error
	CreateUserPeergrp(payload *models.UserPeergrp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadUserPeergrp(mkey string, params *models.CmdbRequestParams) (*models.UserPeergrp, error)
	UpdateUserPeergrp(mkey string, payload *models.UserPeergrp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteUserPeergrp(mkey string, params *models.CmdbRequestParams) error
	ListUserPeergrp(params *models.CmdbRequestParams) (*[]models.UserPeergrp, error)
	CreateSystemSflow(payload *models.SystemSflow, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemSflow(mkey string, params *models.CmdbRequestParams) (*models.SystemSflow, error)
	UpdateSystemSflow(mkey string, payload *models.SystemSflow, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemSflow(mkey string, params *models.CmdbRequestParams) error
	CreateVpnSslClient(payload *models.VpnSslClient, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnSslClient(mkey string, params *models.CmdbRequestParams) (*models.VpnSslClient, error)
	UpdateVpnSslClient(mkey string, payload *models.VpnSslClient, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnSslClient(mkey string, params *models.CmdbRequestParams) error
	ListVpnSslClient(params *models.CmdbRequestParams) (*[]models.VpnSslClient, error)
	CreateSystemReplacemsgSslvpn(payload *models.SystemReplacemsgSslvpn, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemReplacemsgSslvpn(mkey string, params *models.CmdbRequestParams) (*models.SystemReplacemsgSslvpn, error)
	UpdateSystemReplacemsgSslvpn(mkey string, payload *models.SystemReplacemsgSslvpn, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemReplacemsgSslvpn(mkey string, params *models.CmdbRequestParams) error
	ListSystemReplacemsgSslvpn(params *models.CmdbRequestParams) (*[]models.SystemReplacemsgSslvpn, error)
	CreateExtenderSysInfo(payload *models.ExtenderSysInfo, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadExtenderSysInfo(mkey string, params *models.CmdbRequestParams) (*models.ExtenderSysInfo, error)
	UpdateExtenderSysInfo(mkey string, payload *models.ExtenderSysInfo, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteExtenderSysInfo(mkey string, params *models.CmdbRequestParams) error
	CreateVideofilterYoutubeKey(payload *models.VideofilterYoutubeKey, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVideofilterYoutubeKey(mkey string, params *models.CmdbRequestParams) (*models.VideofilterYoutubeKey, error)
	UpdateVideofilterYoutubeKey(mkey string, payload *models.VideofilterYoutubeKey, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVideofilterYoutubeKey(mkey string, params *models.CmdbRequestParams) error
	ListVideofilterYoutubeKey(params *models.CmdbRequestParams) (*[]models.VideofilterYoutubeKey, error)
	CreateFirewallInternetServiceAppend(payload *models.FirewallInternetServiceAppend, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallInternetServiceAppend(mkey string, params *models.CmdbRequestParams) (*models.FirewallInternetServiceAppend, error)
	UpdateFirewallInternetServiceAppend(mkey string, payload *models.FirewallInternetServiceAppend, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallInternetServiceAppend(mkey string, params *models.CmdbRequestParams) error
	CreateSystemSsoAdmin(payload *models.SystemSsoAdmin, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemSsoAdmin(mkey string, params *models.CmdbRequestParams) (*models.SystemSsoAdmin, error)
	UpdateSystemSsoAdmin(mkey string, payload *models.SystemSsoAdmin, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemSsoAdmin(mkey string, params *models.CmdbRequestParams) error
	ListSystemSsoAdmin(params *models.CmdbRequestParams) (*[]models.SystemSsoAdmin, error)
	CreateSystemHaMonitor(payload *models.SystemHaMonitor, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemHaMonitor(mkey string, params *models.CmdbRequestParams) (*models.SystemHaMonitor, error)
	UpdateSystemHaMonitor(mkey string, payload *models.SystemHaMonitor, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemHaMonitor(mkey string, params *models.CmdbRequestParams) error
	CreateVpnIpsecFec(payload *models.VpnIpsecFec, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnIpsecFec(mkey string, params *models.CmdbRequestParams) (*models.VpnIpsecFec, error)
	UpdateVpnIpsecFec(mkey string, payload *models.VpnIpsecFec, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnIpsecFec(mkey string, params *models.CmdbRequestParams) error
	ListVpnIpsecFec(params *models.CmdbRequestParams) (*[]models.VpnIpsecFec, error)
	CreateLogFortianalyzerOverrideSetting(payload *models.LogFortianalyzerOverrideSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogFortianalyzerOverrideSetting(mkey string, params *models.CmdbRequestParams) (*models.LogFortianalyzerOverrideSetting, error)
	UpdateLogFortianalyzerOverrideSetting(mkey string, payload *models.LogFortianalyzerOverrideSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogFortianalyzerOverrideSetting(mkey string, params *models.CmdbRequestParams) error
	CreateRouterSetting(payload *models.RouterSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterSetting(mkey string, params *models.CmdbRequestParams) (*models.RouterSetting, error)
	UpdateRouterSetting(mkey string, payload *models.RouterSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterSetting(mkey string, params *models.CmdbRequestParams) error
	CreateSystemLteModem(payload *models.SystemLteModem, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemLteModem(mkey string, params *models.CmdbRequestParams) (*models.SystemLteModem, error)
	UpdateSystemLteModem(mkey string, payload *models.SystemLteModem, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemLteModem(mkey string, params *models.CmdbRequestParams) error
	CreateSystemNdProxy(payload *models.SystemNdProxy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemNdProxy(mkey string, params *models.CmdbRequestParams) (*models.SystemNdProxy, error)
	UpdateSystemNdProxy(mkey string, payload *models.SystemNdProxy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemNdProxy(mkey string, params *models.CmdbRequestParams) error
	CreateWebProxyProfile(payload *models.WebProxyProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWebProxyProfile(mkey string, params *models.CmdbRequestParams) (*models.WebProxyProfile, error)
	UpdateWebProxyProfile(mkey string, payload *models.WebProxyProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWebProxyProfile(mkey string, params *models.CmdbRequestParams) error
	ListWebProxyProfile(params *models.CmdbRequestParams) (*[]models.WebProxyProfile, error)
	CreateLogSyslogd2OverrideFilter(payload *models.LogSyslogd2OverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogSyslogd2OverrideFilter(mkey string, params *models.CmdbRequestParams) (*models.LogSyslogd2OverrideFilter, error)
	UpdateLogSyslogd2OverrideFilter(mkey string, payload *models.LogSyslogd2OverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogSyslogd2OverrideFilter(mkey string, params *models.CmdbRequestParams) error
	CreateSystemSessionInfoList(payload *models.SystemSessionInfoList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemSessionInfoList(mkey string, params *models.CmdbRequestParams) (*models.SystemSessionInfoList, error)
	UpdateSystemSessionInfoList(mkey string, payload *models.SystemSessionInfoList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemSessionInfoList(mkey string, params *models.CmdbRequestParams) error
	CreateSystemLldpNetworkPolicy(payload *models.SystemLldpNetworkPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemLldpNetworkPolicy(mkey string, params *models.CmdbRequestParams) (*models.SystemLldpNetworkPolicy, error)
	UpdateSystemLldpNetworkPolicy(mkey string, payload *models.SystemLldpNetworkPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemLldpNetworkPolicy(mkey string, params *models.CmdbRequestParams) error
	ListSystemLldpNetworkPolicy(params *models.CmdbRequestParams) (*[]models.SystemLldpNetworkPolicy, error)
	CreateWirelessControllerAddrgrp(payload *models.WirelessControllerAddrgrp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerAddrgrp(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerAddrgrp, error)
	UpdateWirelessControllerAddrgrp(mkey string, payload *models.WirelessControllerAddrgrp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerAddrgrp(mkey string, params *models.CmdbRequestParams) error
	ListWirelessControllerAddrgrp(params *models.CmdbRequestParams) (*[]models.WirelessControllerAddrgrp, error)
	CreateSwitchControllerFortilinkSettings(payload *models.SwitchControllerFortilinkSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerFortilinkSettings(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerFortilinkSettings, error)
	UpdateSwitchControllerFortilinkSettings(mkey string, payload *models.SwitchControllerFortilinkSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerFortilinkSettings(mkey string, params *models.CmdbRequestParams) error
	ListSwitchControllerFortilinkSettings(params *models.CmdbRequestParams) (*[]models.SwitchControllerFortilinkSettings, error)
	CreateFirewallCity(payload *models.FirewallCity, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallCity(mkey string, params *models.CmdbRequestParams) (*models.FirewallCity, error)
	UpdateFirewallCity(mkey string, payload *models.FirewallCity, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallCity(mkey string, params *models.CmdbRequestParams) error
	ListFirewallCity(params *models.CmdbRequestParams) (*[]models.FirewallCity, error)
	CreateFirewallInternetServiceCustom(payload *models.FirewallInternetServiceCustom, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallInternetServiceCustom(mkey string, params *models.CmdbRequestParams) (*models.FirewallInternetServiceCustom, error)
	UpdateFirewallInternetServiceCustom(mkey string, payload *models.FirewallInternetServiceCustom, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallInternetServiceCustom(mkey string, params *models.CmdbRequestParams) error
	ListFirewallInternetServiceCustom(params *models.CmdbRequestParams) (*[]models.FirewallInternetServiceCustom, error)
	CreateFirewallIppool6(payload *models.FirewallIppool6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallIppool6(mkey string, params *models.CmdbRequestParams) (*models.FirewallIppool6, error)
	UpdateFirewallIppool6(mkey string, payload *models.FirewallIppool6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallIppool6(mkey string, params *models.CmdbRequestParams) error
	ListFirewallIppool6(params *models.CmdbRequestParams) (*[]models.FirewallIppool6, error)
	CreateRouterPolicy6(payload *models.RouterPolicy6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterPolicy6(mkey string, params *models.CmdbRequestParams) (*models.RouterPolicy6, error)
	UpdateRouterPolicy6(mkey string, payload *models.RouterPolicy6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterPolicy6(mkey string, params *models.CmdbRequestParams) error
	ListRouterPolicy6(params *models.CmdbRequestParams) (*[]models.RouterPolicy6, error)
	CreateSystemAutoupdateTunneling(payload *models.SystemAutoupdateTunneling, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemAutoupdateTunneling(mkey string, params *models.CmdbRequestParams) (*models.SystemAutoupdateTunneling, error)
	UpdateSystemAutoupdateTunneling(mkey string, payload *models.SystemAutoupdateTunneling, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemAutoupdateTunneling(mkey string, params *models.CmdbRequestParams) error
	CreateSystemReplacemsgWebproxy(payload *models.SystemReplacemsgWebproxy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemReplacemsgWebproxy(mkey string, params *models.CmdbRequestParams) (*models.SystemReplacemsgWebproxy, error)
	UpdateSystemReplacemsgWebproxy(mkey string, payload *models.SystemReplacemsgWebproxy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemReplacemsgWebproxy(mkey string, params *models.CmdbRequestParams) error
	ListSystemReplacemsgWebproxy(params *models.CmdbRequestParams) (*[]models.SystemReplacemsgWebproxy, error)
	CreateSwitchControllerMacPolicy(payload *models.SwitchControllerMacPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerMacPolicy(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerMacPolicy, error)
	UpdateSwitchControllerMacPolicy(mkey string, payload *models.SwitchControllerMacPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerMacPolicy(mkey string, params *models.CmdbRequestParams) error
	ListSwitchControllerMacPolicy(params *models.CmdbRequestParams) (*[]models.SwitchControllerMacPolicy, error)
	CreateUserRadius(payload *models.UserRadius, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadUserRadius(mkey string, params *models.CmdbRequestParams) (*models.UserRadius, error)
	UpdateUserRadius(mkey string, payload *models.UserRadius, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteUserRadius(mkey string, params *models.CmdbRequestParams) error
	ListUserRadius(params *models.CmdbRequestParams) (*[]models.UserRadius, error)
	CreateRouterBfd6(payload *models.RouterBfd6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterBfd6(mkey string, params *models.CmdbRequestParams) (*models.RouterBfd6, error)
	UpdateRouterBfd6(mkey string, payload *models.RouterBfd6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterBfd6(mkey string, params *models.CmdbRequestParams) error
	CreateSystemSmsServer(payload *models.SystemSmsServer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemSmsServer(mkey string, params *models.CmdbRequestParams) (*models.SystemSmsServer, error)
	UpdateSystemSmsServer(mkey string, payload *models.SystemSmsServer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemSmsServer(mkey string, params *models.CmdbRequestParams) error
	ListSystemSmsServer(params *models.CmdbRequestParams) (*[]models.SystemSmsServer, error)
	CreateSystemReplacemsgNacQuar(payload *models.SystemReplacemsgNacQuar, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemReplacemsgNacQuar(mkey string, params *models.CmdbRequestParams) (*models.SystemReplacemsgNacQuar, error)
	UpdateSystemReplacemsgNacQuar(mkey string, payload *models.SystemReplacemsgNacQuar, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemReplacemsgNacQuar(mkey string, params *models.CmdbRequestParams) error
	ListSystemReplacemsgNacQuar(params *models.CmdbRequestParams) (*[]models.SystemReplacemsgNacQuar, error)
	CreateSwitchControllerSecurityPolicy8021X(payload *models.SwitchControllerSecurityPolicy8021X, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerSecurityPolicy8021X(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerSecurityPolicy8021X, error)
	UpdateSwitchControllerSecurityPolicy8021X(mkey string, payload *models.SwitchControllerSecurityPolicy8021X, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerSecurityPolicy8021X(mkey string, params *models.CmdbRequestParams) error
	ListSwitchControllerSecurityPolicy8021X(params *models.CmdbRequestParams) (*[]models.SwitchControllerSecurityPolicy8021X, error)
	CreateSwitchControllerSnmpSysinfo(payload *models.SwitchControllerSnmpSysinfo, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerSnmpSysinfo(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerSnmpSysinfo, error)
	UpdateSwitchControllerSnmpSysinfo(mkey string, payload *models.SwitchControllerSnmpSysinfo, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerSnmpSysinfo(mkey string, params *models.CmdbRequestParams) error
	CreateWebProxyExplicit(payload *models.WebProxyExplicit, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWebProxyExplicit(mkey string, params *models.CmdbRequestParams) (*models.WebProxyExplicit, error)
	UpdateWebProxyExplicit(mkey string, payload *models.WebProxyExplicit, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWebProxyExplicit(mkey string, params *models.CmdbRequestParams) error
	CreateSystemAccprofile(payload *models.SystemAccprofile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemAccprofile(mkey string, params *models.CmdbRequestParams) (*models.SystemAccprofile, error)
	UpdateSystemAccprofile(mkey string, payload *models.SystemAccprofile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemAccprofile(mkey string, params *models.CmdbRequestParams) error
	ListSystemAccprofile(params *models.CmdbRequestParams) (*[]models.SystemAccprofile, error)
	CreateFirewallCountry(payload *models.FirewallCountry, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallCountry(mkey string, params *models.CmdbRequestParams) (*models.FirewallCountry, error)
	UpdateFirewallCountry(mkey string, payload *models.FirewallCountry, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallCountry(mkey string, params *models.CmdbRequestParams) error
	ListFirewallCountry(params *models.CmdbRequestParams) (*[]models.FirewallCountry, error)
	CreateWebfilterFtgdLocalCat(payload *models.WebfilterFtgdLocalCat, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWebfilterFtgdLocalCat(mkey string, params *models.CmdbRequestParams) (*models.WebfilterFtgdLocalCat, error)
	UpdateWebfilterFtgdLocalCat(mkey string, payload *models.WebfilterFtgdLocalCat, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWebfilterFtgdLocalCat(mkey string, params *models.CmdbRequestParams) error
	ListWebfilterFtgdLocalCat(params *models.CmdbRequestParams) (*[]models.WebfilterFtgdLocalCat, error)
	CreateLogFortiguardOverrideSetting(payload *models.LogFortiguardOverrideSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogFortiguardOverrideSetting(mkey string, params *models.CmdbRequestParams) (*models.LogFortiguardOverrideSetting, error)
	UpdateLogFortiguardOverrideSetting(mkey string, payload *models.LogFortiguardOverrideSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogFortiguardOverrideSetting(mkey string, params *models.CmdbRequestParams) error
	CreateLogTacacsaccounting3Filter(payload *models.LogTacacsaccounting3Filter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogTacacsaccounting3Filter(mkey string, params *models.CmdbRequestParams) (*models.LogTacacsaccounting3Filter, error)
	UpdateLogTacacsaccounting3Filter(mkey string, payload *models.LogTacacsaccounting3Filter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogTacacsaccounting3Filter(mkey string, params *models.CmdbRequestParams) error
	CreateSwitchControllerDynamicPortPolicy(payload *models.SwitchControllerDynamicPortPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerDynamicPortPolicy(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerDynamicPortPolicy, error)
	UpdateSwitchControllerDynamicPortPolicy(mkey string, payload *models.SwitchControllerDynamicPortPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerDynamicPortPolicy(mkey string, params *models.CmdbRequestParams) error
	ListSwitchControllerDynamicPortPolicy(params *models.CmdbRequestParams) (*[]models.SwitchControllerDynamicPortPolicy, error)
	CreateWebfilterIpsUrlfilterCacheSetting(payload *models.WebfilterIpsUrlfilterCacheSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWebfilterIpsUrlfilterCacheSetting(mkey string, params *models.CmdbRequestParams) (*models.WebfilterIpsUrlfilterCacheSetting, error)
	UpdateWebfilterIpsUrlfilterCacheSetting(mkey string, payload *models.WebfilterIpsUrlfilterCacheSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWebfilterIpsUrlfilterCacheSetting(mkey string, params *models.CmdbRequestParams) error
	CreateApplicationCustom(payload *models.ApplicationCustom, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadApplicationCustom(mkey string, params *models.CmdbRequestParams) (*models.ApplicationCustom, error)
	UpdateApplicationCustom(mkey string, payload *models.ApplicationCustom, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteApplicationCustom(mkey string, params *models.CmdbRequestParams) error
	ListApplicationCustom(params *models.CmdbRequestParams) (*[]models.ApplicationCustom, error)
	CreateUserFortitoken(payload *models.UserFortitoken, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadUserFortitoken(mkey string, params *models.CmdbRequestParams) (*models.UserFortitoken, error)
	UpdateUserFortitoken(mkey string, payload *models.UserFortitoken, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteUserFortitoken(mkey string, params *models.CmdbRequestParams) error
	ListUserFortitoken(params *models.CmdbRequestParams) (*[]models.UserFortitoken, error)
	CreateFirewallInternetServiceBotnet(payload *models.FirewallInternetServiceBotnet, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallInternetServiceBotnet(mkey string, params *models.CmdbRequestParams) (*models.FirewallInternetServiceBotnet, error)
	UpdateFirewallInternetServiceBotnet(mkey string, payload *models.FirewallInternetServiceBotnet, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallInternetServiceBotnet(mkey string, params *models.CmdbRequestParams) error
	ListFirewallInternetServiceBotnet(params *models.CmdbRequestParams) (*[]models.FirewallInternetServiceBotnet, error)
	CreateSystemReplacemsgImage(payload *models.SystemReplacemsgImage, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemReplacemsgImage(mkey string, params *models.CmdbRequestParams) (*models.SystemReplacemsgImage, error)
	UpdateSystemReplacemsgImage(mkey string, payload *models.SystemReplacemsgImage, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemReplacemsgImage(mkey string, params *models.CmdbRequestParams) error
	ListSystemReplacemsgImage(params *models.CmdbRequestParams) (*[]models.SystemReplacemsgImage, error)
	CreateSystemVdomProperty(payload *models.SystemVdomProperty, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemVdomProperty(mkey string, params *models.CmdbRequestParams) (*models.SystemVdomProperty, error)
	UpdateSystemVdomProperty(mkey string, payload *models.SystemVdomProperty, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemVdomProperty(mkey string, params *models.CmdbRequestParams) error
	ListSystemVdomProperty(params *models.CmdbRequestParams) (*[]models.SystemVdomProperty, error)
	CreateFirewallAddress(payload *models.FirewallAddress, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallAddress(mkey string, params *models.CmdbRequestParams) (*models.FirewallAddress, error)
	UpdateFirewallAddress(mkey string, payload *models.FirewallAddress, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallAddress(mkey string, params *models.CmdbRequestParams) error
	ListFirewallAddress(params *models.CmdbRequestParams) (*[]models.FirewallAddress, error)
	CreateLogFortianalyzerFilter(payload *models.LogFortianalyzerFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogFortianalyzerFilter(mkey string, params *models.CmdbRequestParams) (*models.LogFortianalyzerFilter, error)
	UpdateLogFortianalyzerFilter(mkey string, payload *models.LogFortianalyzerFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogFortianalyzerFilter(mkey string, params *models.CmdbRequestParams) error
	CreateSystemVdomLink(payload *models.SystemVdomLink, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemVdomLink(mkey string, params *models.CmdbRequestParams) (*models.SystemVdomLink, error)
	UpdateSystemVdomLink(mkey string, payload *models.SystemVdomLink, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemVdomLink(mkey string, params *models.CmdbRequestParams) error
	ListSystemVdomLink(params *models.CmdbRequestParams) (*[]models.SystemVdomLink, error)
	CreateVpnSslWebRealm(payload *models.VpnSslWebRealm, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnSslWebRealm(mkey string, params *models.CmdbRequestParams) (*models.VpnSslWebRealm, error)
	UpdateVpnSslWebRealm(mkey string, payload *models.VpnSslWebRealm, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnSslWebRealm(mkey string, params *models.CmdbRequestParams) error
	ListVpnSslWebRealm(params *models.CmdbRequestParams) (*[]models.VpnSslWebRealm, error)
	CreateRouterMulticastFlow(payload *models.RouterMulticastFlow, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterMulticastFlow(mkey string, params *models.CmdbRequestParams) (*models.RouterMulticastFlow, error)
	UpdateRouterMulticastFlow(mkey string, payload *models.RouterMulticastFlow, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterMulticastFlow(mkey string, params *models.CmdbRequestParams) error
	ListRouterMulticastFlow(params *models.CmdbRequestParams) (*[]models.RouterMulticastFlow, error)
	CreateSystemSessionInfoFullStat(payload *models.SystemSessionInfoFullStat, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemSessionInfoFullStat(mkey string, params *models.CmdbRequestParams) (*models.SystemSessionInfoFullStat, error)
	UpdateSystemSessionInfoFullStat(mkey string, payload *models.SystemSessionInfoFullStat, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemSessionInfoFullStat(mkey string, params *models.CmdbRequestParams) error
	CreateSystemExternalResource(payload *models.SystemExternalResource, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemExternalResource(mkey string, params *models.CmdbRequestParams) (*models.SystemExternalResource, error)
	UpdateSystemExternalResource(mkey string, payload *models.SystemExternalResource, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemExternalResource(mkey string, params *models.CmdbRequestParams) error
	ListSystemExternalResource(params *models.CmdbRequestParams) (*[]models.SystemExternalResource, error)
	CreateFirewallShapingProfile(payload *models.FirewallShapingProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallShapingProfile(mkey string, params *models.CmdbRequestParams) (*models.FirewallShapingProfile, error)
	UpdateFirewallShapingProfile(mkey string, payload *models.FirewallShapingProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallShapingProfile(mkey string, params *models.CmdbRequestParams) error
	ListFirewallShapingProfile(params *models.CmdbRequestParams) (*[]models.FirewallShapingProfile, error)
	CreateSystemDns(payload *models.SystemDns, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemDns(mkey string, params *models.CmdbRequestParams) (*models.SystemDns, error)
	UpdateSystemDns(mkey string, payload *models.SystemDns, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemDns(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallIpmacbindingTable(payload *models.FirewallIpmacbindingTable, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallIpmacbindingTable(mkey string, params *models.CmdbRequestParams) (*models.FirewallIpmacbindingTable, error)
	UpdateFirewallIpmacbindingTable(mkey string, payload *models.FirewallIpmacbindingTable, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallIpmacbindingTable(mkey string, params *models.CmdbRequestParams) error
	ListFirewallIpmacbindingTable(params *models.CmdbRequestParams) (*[]models.FirewallIpmacbindingTable, error)
	CreateVpnIpsecTunnelName(payload *models.VpnIpsecTunnelName, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnIpsecTunnelName(mkey string, params *models.CmdbRequestParams) (*models.VpnIpsecTunnelName, error)
	UpdateVpnIpsecTunnelName(mkey string, payload *models.VpnIpsecTunnelName, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnIpsecTunnelName(mkey string, params *models.CmdbRequestParams) error
	CreateEmailfilterDnsbl(payload *models.EmailfilterDnsbl, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadEmailfilterDnsbl(mkey string, params *models.CmdbRequestParams) (*models.EmailfilterDnsbl, error)
	UpdateEmailfilterDnsbl(mkey string, payload *models.EmailfilterDnsbl, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteEmailfilterDnsbl(mkey string, params *models.CmdbRequestParams) error
	ListEmailfilterDnsbl(params *models.CmdbRequestParams) (*[]models.EmailfilterDnsbl, error)
	CreateRouterStatic(payload *models.RouterStatic, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterStatic(mkey string, params *models.CmdbRequestParams) (*models.RouterStatic, error)
	UpdateRouterStatic(mkey string, payload *models.RouterStatic, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterStatic(mkey string, params *models.CmdbRequestParams) error
	ListRouterStatic(params *models.CmdbRequestParams) (*[]models.RouterStatic, error)
	CreateRouterOspf6(payload *models.RouterOspf6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterOspf6(mkey string, params *models.CmdbRequestParams) (*models.RouterOspf6, error)
	UpdateRouterOspf6(mkey string, payload *models.RouterOspf6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterOspf6(mkey string, params *models.CmdbRequestParams) error
	CreateSystemHaNonsyncCsum(payload *models.SystemHaNonsyncCsum, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemHaNonsyncCsum(mkey string, params *models.CmdbRequestParams) (*models.SystemHaNonsyncCsum, error)
	UpdateSystemHaNonsyncCsum(mkey string, payload *models.SystemHaNonsyncCsum, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemHaNonsyncCsum(mkey string, params *models.CmdbRequestParams) error
	CreateSystemIps(payload *models.SystemIps, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemIps(mkey string, params *models.CmdbRequestParams) (*models.SystemIps, error)
	UpdateSystemIps(mkey string, payload *models.SystemIps, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemIps(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerHotspot20Anqp3GppCellular(payload *models.WirelessControllerHotspot20Anqp3GppCellular, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerHotspot20Anqp3GppCellular(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerHotspot20Anqp3GppCellular, error)
	UpdateWirelessControllerHotspot20Anqp3GppCellular(mkey string, payload *models.WirelessControllerHotspot20Anqp3GppCellular, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerHotspot20Anqp3GppCellular(mkey string, params *models.CmdbRequestParams) error
	ListWirelessControllerHotspot20Anqp3GppCellular(params *models.CmdbRequestParams) (*[]models.WirelessControllerHotspot20Anqp3GppCellular, error)
	CreateWirelessControllerHotspot20H2qpOsuProvider(payload *models.WirelessControllerHotspot20H2qpOsuProvider, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerHotspot20H2qpOsuProvider(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerHotspot20H2qpOsuProvider, error)
	UpdateWirelessControllerHotspot20H2qpOsuProvider(mkey string, payload *models.WirelessControllerHotspot20H2qpOsuProvider, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerHotspot20H2qpOsuProvider(mkey string, params *models.CmdbRequestParams) error
	ListWirelessControllerHotspot20H2qpOsuProvider(params *models.CmdbRequestParams) (*[]models.WirelessControllerHotspot20H2qpOsuProvider, error)
	CreateSwitchControllerAutoConfigDefault(payload *models.SwitchControllerAutoConfigDefault, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerAutoConfigDefault(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerAutoConfigDefault, error)
	UpdateSwitchControllerAutoConfigDefault(mkey string, payload *models.SwitchControllerAutoConfigDefault, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerAutoConfigDefault(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallProxyAddrgrp(payload *models.FirewallProxyAddrgrp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallProxyAddrgrp(mkey string, params *models.CmdbRequestParams) (*models.FirewallProxyAddrgrp, error)
	UpdateFirewallProxyAddrgrp(mkey string, payload *models.FirewallProxyAddrgrp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallProxyAddrgrp(mkey string, params *models.CmdbRequestParams) error
	ListFirewallProxyAddrgrp(params *models.CmdbRequestParams) (*[]models.FirewallProxyAddrgrp, error)
	CreateFirewallDosPolicy(payload *models.FirewallDosPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallDosPolicy(mkey string, params *models.CmdbRequestParams) (*models.FirewallDosPolicy, error)
	UpdateFirewallDosPolicy(mkey string, payload *models.FirewallDosPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallDosPolicy(mkey string, params *models.CmdbRequestParams) error
	ListFirewallDosPolicy(params *models.CmdbRequestParams) (*[]models.FirewallDosPolicy, error)
	CreateVpnCertificateCa(payload *models.VpnCertificateCa, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnCertificateCa(mkey string, params *models.CmdbRequestParams) (*models.VpnCertificateCa, error)
	UpdateVpnCertificateCa(mkey string, payload *models.VpnCertificateCa, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnCertificateCa(mkey string, params *models.CmdbRequestParams) error
	ListVpnCertificateCa(params *models.CmdbRequestParams) (*[]models.VpnCertificateCa, error)
	CreateSystemResourceLimits(payload *models.SystemResourceLimits, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemResourceLimits(mkey string, params *models.CmdbRequestParams) (*models.SystemResourceLimits, error)
	UpdateSystemResourceLimits(mkey string, payload *models.SystemResourceLimits, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemResourceLimits(mkey string, params *models.CmdbRequestParams) error
	CreateSystemSnmpCommunity(payload *models.SystemSnmpCommunity, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemSnmpCommunity(mkey string, params *models.CmdbRequestParams) (*models.SystemSnmpCommunity, error)
	UpdateSystemSnmpCommunity(mkey string, payload *models.SystemSnmpCommunity, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemSnmpCommunity(mkey string, params *models.CmdbRequestParams) error
	ListSystemSnmpCommunity(params *models.CmdbRequestParams) (*[]models.SystemSnmpCommunity, error)
	CreateSystemManagementTunnel(payload *models.SystemManagementTunnel, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemManagementTunnel(mkey string, params *models.CmdbRequestParams) (*models.SystemManagementTunnel, error)
	UpdateSystemManagementTunnel(mkey string, payload *models.SystemManagementTunnel, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemManagementTunnel(mkey string, params *models.CmdbRequestParams) error
	CreateSwitchControllerSwitchProfile(payload *models.SwitchControllerSwitchProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerSwitchProfile(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerSwitchProfile, error)
	UpdateSwitchControllerSwitchProfile(mkey string, payload *models.SwitchControllerSwitchProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerSwitchProfile(mkey string, params *models.CmdbRequestParams) error
	ListSwitchControllerSwitchProfile(params *models.CmdbRequestParams) (*[]models.SwitchControllerSwitchProfile, error)
	CreateSwitchControllerManagedSwitch(payload *models.SwitchControllerManagedSwitch, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerManagedSwitch(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerManagedSwitch, error)
	UpdateSwitchControllerManagedSwitch(mkey string, payload *models.SwitchControllerManagedSwitch, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerManagedSwitch(mkey string, params *models.CmdbRequestParams) error
	ListSwitchControllerManagedSwitch(params *models.CmdbRequestParams) (*[]models.SwitchControllerManagedSwitch, error)
	CreateFirewallAcl6(payload *models.FirewallAcl6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallAcl6(mkey string, params *models.CmdbRequestParams) (*models.FirewallAcl6, error)
	UpdateFirewallAcl6(mkey string, payload *models.FirewallAcl6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallAcl6(mkey string, params *models.CmdbRequestParams) error
	ListFirewallAcl6(params *models.CmdbRequestParams) (*[]models.FirewallAcl6, error)
	CreateSystemReplacemsgGroup(payload *models.SystemReplacemsgGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemReplacemsgGroup(mkey string, params *models.CmdbRequestParams) (*models.SystemReplacemsgGroup, error)
	UpdateSystemReplacemsgGroup(mkey string, payload *models.SystemReplacemsgGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemReplacemsgGroup(mkey string, params *models.CmdbRequestParams) error
	ListSystemReplacemsgGroup(params *models.CmdbRequestParams) (*[]models.SystemReplacemsgGroup, error)
	CreateSystemConsole(payload *models.SystemConsole, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemConsole(mkey string, params *models.CmdbRequestParams) (*models.SystemConsole, error)
	UpdateSystemConsole(mkey string, payload *models.SystemConsole, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemConsole(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallVendorMac(payload *models.FirewallVendorMac, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallVendorMac(mkey string, params *models.CmdbRequestParams) (*models.FirewallVendorMac, error)
	UpdateFirewallVendorMac(mkey string, payload *models.FirewallVendorMac, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallVendorMac(mkey string, params *models.CmdbRequestParams) error
	ListFirewallVendorMac(params *models.CmdbRequestParams) (*[]models.FirewallVendorMac, error)
	CreateUserCertificate(payload *models.UserCertificate, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadUserCertificate(mkey string, params *models.CmdbRequestParams) (*models.UserCertificate, error)
	UpdateUserCertificate(mkey string, payload *models.UserCertificate, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteUserCertificate(mkey string, params *models.CmdbRequestParams) error
	ListUserCertificate(params *models.CmdbRequestParams) (*[]models.UserCertificate, error)
	CreateSystemSmcNtp(payload *models.SystemSmcNtp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemSmcNtp(mkey string, params *models.CmdbRequestParams) (*models.SystemSmcNtp, error)
	UpdateSystemSmcNtp(mkey string, payload *models.SystemSmcNtp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemSmcNtp(mkey string, params *models.CmdbRequestParams) error
	CreateSystemArpTable(payload *models.SystemArpTable, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemArpTable(mkey string, params *models.CmdbRequestParams) (*models.SystemArpTable, error)
	UpdateSystemArpTable(mkey string, payload *models.SystemArpTable, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemArpTable(mkey string, params *models.CmdbRequestParams) error
	ListSystemArpTable(params *models.CmdbRequestParams) (*[]models.SystemArpTable, error)
	CreateSystemAutomationDestination(payload *models.SystemAutomationDestination, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemAutomationDestination(mkey string, params *models.CmdbRequestParams) (*models.SystemAutomationDestination, error)
	UpdateSystemAutomationDestination(mkey string, payload *models.SystemAutomationDestination, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemAutomationDestination(mkey string, params *models.CmdbRequestParams) error
	ListSystemAutomationDestination(params *models.CmdbRequestParams) (*[]models.SystemAutomationDestination, error)
	CreateWirelessControllerArrpProfile(payload *models.WirelessControllerArrpProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerArrpProfile(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerArrpProfile, error)
	UpdateWirelessControllerArrpProfile(mkey string, payload *models.WirelessControllerArrpProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerArrpProfile(mkey string, params *models.CmdbRequestParams) error
	ListWirelessControllerArrpProfile(params *models.CmdbRequestParams) (*[]models.WirelessControllerArrpProfile, error)
	CreateSwitchControllerSystem(payload *models.SwitchControllerSystem, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerSystem(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerSystem, error)
	UpdateSwitchControllerSystem(mkey string, payload *models.SwitchControllerSystem, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerSystem(mkey string, params *models.CmdbRequestParams) error
	CreateSwitchControllerSnmpUser(payload *models.SwitchControllerSnmpUser, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerSnmpUser(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerSnmpUser, error)
	UpdateSwitchControllerSnmpUser(mkey string, payload *models.SwitchControllerSnmpUser, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerSnmpUser(mkey string, params *models.CmdbRequestParams) error
	ListSwitchControllerSnmpUser(params *models.CmdbRequestParams) (*[]models.SwitchControllerSnmpUser, error)
	CreateDlpSettings(payload *models.DlpSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadDlpSettings(mkey string, params *models.CmdbRequestParams) (*models.DlpSettings, error)
	UpdateDlpSettings(mkey string, payload *models.DlpSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteDlpSettings(mkey string, params *models.CmdbRequestParams) error
	CreateSystemReplacemsgIcap(payload *models.SystemReplacemsgIcap, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemReplacemsgIcap(mkey string, params *models.CmdbRequestParams) (*models.SystemReplacemsgIcap, error)
	UpdateSystemReplacemsgIcap(mkey string, payload *models.SystemReplacemsgIcap, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemReplacemsgIcap(mkey string, params *models.CmdbRequestParams) error
	ListSystemReplacemsgIcap(params *models.CmdbRequestParams) (*[]models.SystemReplacemsgIcap, error)
	CreateSwitchControllerLldpSettings(payload *models.SwitchControllerLldpSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerLldpSettings(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerLldpSettings, error)
	UpdateSwitchControllerLldpSettings(mkey string, payload *models.SwitchControllerLldpSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerLldpSettings(mkey string, params *models.CmdbRequestParams) error
	CreateSwitchControllerQosIpDscpMap(payload *models.SwitchControllerQosIpDscpMap, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerQosIpDscpMap(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerQosIpDscpMap, error)
	UpdateSwitchControllerQosIpDscpMap(mkey string, payload *models.SwitchControllerQosIpDscpMap, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerQosIpDscpMap(mkey string, params *models.CmdbRequestParams) error
	ListSwitchControllerQosIpDscpMap(params *models.CmdbRequestParams) (*[]models.SwitchControllerQosIpDscpMap, error)
	CreateSwitchControllerVlanPolicy(payload *models.SwitchControllerVlanPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerVlanPolicy(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerVlanPolicy, error)
	UpdateSwitchControllerVlanPolicy(mkey string, payload *models.SwitchControllerVlanPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerVlanPolicy(mkey string, params *models.CmdbRequestParams) error
	ListSwitchControllerVlanPolicy(params *models.CmdbRequestParams) (*[]models.SwitchControllerVlanPolicy, error)
	CreateFirewallLocalInPolicy6(payload *models.FirewallLocalInPolicy6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallLocalInPolicy6(mkey string, params *models.CmdbRequestParams) (*models.FirewallLocalInPolicy6, error)
	UpdateFirewallLocalInPolicy6(mkey string, payload *models.FirewallLocalInPolicy6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallLocalInPolicy6(mkey string, params *models.CmdbRequestParams) error
	ListFirewallLocalInPolicy6(params *models.CmdbRequestParams) (*[]models.FirewallLocalInPolicy6, error)
	CreateUserAdgrp(payload *models.UserAdgrp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadUserAdgrp(mkey string, params *models.CmdbRequestParams) (*models.UserAdgrp, error)
	UpdateUserAdgrp(mkey string, payload *models.UserAdgrp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteUserAdgrp(mkey string, params *models.CmdbRequestParams) error
	ListUserAdgrp(params *models.CmdbRequestParams) (*[]models.UserAdgrp, error)
	CreateRouterInfo6(payload *models.RouterInfo6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterInfo6(mkey string, params *models.CmdbRequestParams) (*models.RouterInfo6, error)
	UpdateRouterInfo6(mkey string, payload *models.RouterInfo6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterInfo6(mkey string, params *models.CmdbRequestParams) error
	CreateSystemIpsUrlfilterDns6(payload *models.SystemIpsUrlfilterDns6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemIpsUrlfilterDns6(mkey string, params *models.CmdbRequestParams) (*models.SystemIpsUrlfilterDns6, error)
	UpdateSystemIpsUrlfilterDns6(mkey string, payload *models.SystemIpsUrlfilterDns6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemIpsUrlfilterDns6(mkey string, params *models.CmdbRequestParams) error
	ListSystemIpsUrlfilterDns6(params *models.CmdbRequestParams) (*[]models.SystemIpsUrlfilterDns6, error)
	CreateLogGuiDisplay(payload *models.LogGuiDisplay, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogGuiDisplay(mkey string, params *models.CmdbRequestParams) (*models.LogGuiDisplay, error)
	UpdateLogGuiDisplay(mkey string, payload *models.LogGuiDisplay, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogGuiDisplay(mkey string, params *models.CmdbRequestParams) error
	CreateSystemObjectTagging(payload *models.SystemObjectTagging, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemObjectTagging(mkey string, params *models.CmdbRequestParams) (*models.SystemObjectTagging, error)
	UpdateSystemObjectTagging(mkey string, payload *models.SystemObjectTagging, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemObjectTagging(mkey string, params *models.CmdbRequestParams) error
	ListSystemObjectTagging(params *models.CmdbRequestParams) (*[]models.SystemObjectTagging, error)
	CreateFirewallProute(payload *models.FirewallProute, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallProute(mkey string, params *models.CmdbRequestParams) (*models.FirewallProute, error)
	UpdateFirewallProute(mkey string, payload *models.FirewallProute, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallProute(mkey string, params *models.CmdbRequestParams) error
	CreateVpnIpsecStatsCrypto(payload *models.VpnIpsecStatsCrypto, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnIpsecStatsCrypto(mkey string, params *models.CmdbRequestParams) (*models.VpnIpsecStatsCrypto, error)
	UpdateVpnIpsecStatsCrypto(mkey string, payload *models.VpnIpsecStatsCrypto, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnIpsecStatsCrypto(mkey string, params *models.CmdbRequestParams) error
	CreateLogSyslogdFilter(payload *models.LogSyslogdFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogSyslogdFilter(mkey string, params *models.CmdbRequestParams) (*models.LogSyslogdFilter, error)
	UpdateLogSyslogdFilter(mkey string, payload *models.LogSyslogdFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogSyslogdFilter(mkey string, params *models.CmdbRequestParams) error
	CreateLogMemoryGlobalSetting(payload *models.LogMemoryGlobalSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogMemoryGlobalSetting(mkey string, params *models.CmdbRequestParams) (*models.LogMemoryGlobalSetting, error)
	UpdateLogMemoryGlobalSetting(mkey string, payload *models.LogMemoryGlobalSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogMemoryGlobalSetting(mkey string, params *models.CmdbRequestParams) error
	CreateLogFortianalyzer3Filter(payload *models.LogFortianalyzer3Filter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogFortianalyzer3Filter(mkey string, params *models.CmdbRequestParams) (*models.LogFortianalyzer3Filter, error)
	UpdateLogFortianalyzer3Filter(mkey string, payload *models.LogFortianalyzer3Filter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogFortianalyzer3Filter(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallVipgrp(payload *models.FirewallVipgrp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallVipgrp(mkey string, params *models.CmdbRequestParams) (*models.FirewallVipgrp, error)
	UpdateFirewallVipgrp(mkey string, payload *models.FirewallVipgrp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallVipgrp(mkey string, params *models.CmdbRequestParams) error
	ListFirewallVipgrp(params *models.CmdbRequestParams) (*[]models.FirewallVipgrp, error)
	CreateFirewallVipgrp6(payload *models.FirewallVipgrp6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallVipgrp6(mkey string, params *models.CmdbRequestParams) (*models.FirewallVipgrp6, error)
	UpdateFirewallVipgrp6(mkey string, payload *models.FirewallVipgrp6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallVipgrp6(mkey string, params *models.CmdbRequestParams) error
	ListFirewallVipgrp6(params *models.CmdbRequestParams) (*[]models.FirewallVipgrp6, error)
	CreateLogTacacsaccounting3Setting(payload *models.LogTacacsaccounting3Setting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogTacacsaccounting3Setting(mkey string, params *models.CmdbRequestParams) (*models.LogTacacsaccounting3Setting, error)
	UpdateLogTacacsaccounting3Setting(mkey string, payload *models.LogTacacsaccounting3Setting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogTacacsaccounting3Setting(mkey string, params *models.CmdbRequestParams) error
	CreateUserFsso(payload *models.UserFsso, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadUserFsso(mkey string, params *models.CmdbRequestParams) (*models.UserFsso, error)
	UpdateUserFsso(mkey string, payload *models.UserFsso, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteUserFsso(mkey string, params *models.CmdbRequestParams) error
	ListUserFsso(params *models.CmdbRequestParams) (*[]models.UserFsso, error)
	CreateSystemGeneve(payload *models.SystemGeneve, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemGeneve(mkey string, params *models.CmdbRequestParams) (*models.SystemGeneve, error)
	UpdateSystemGeneve(mkey string, payload *models.SystemGeneve, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemGeneve(mkey string, params *models.CmdbRequestParams) error
	ListSystemGeneve(params *models.CmdbRequestParams) (*[]models.SystemGeneve, error)
	CreateSwitchControllerStpSettings(payload *models.SwitchControllerStpSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerStpSettings(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerStpSettings, error)
	UpdateSwitchControllerStpSettings(mkey string, payload *models.SwitchControllerStpSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerStpSettings(mkey string, params *models.CmdbRequestParams) error
	CreateVpnSslSettings(payload *models.VpnSslSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnSslSettings(mkey string, params *models.CmdbRequestParams) (*models.VpnSslSettings, error)
	UpdateVpnSslSettings(mkey string, payload *models.VpnSslSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnSslSettings(mkey string, params *models.CmdbRequestParams) error
	CreateWebfilterFtgdStatistics(payload *models.WebfilterFtgdStatistics, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWebfilterFtgdStatistics(mkey string, params *models.CmdbRequestParams) (*models.WebfilterFtgdStatistics, error)
	UpdateWebfilterFtgdStatistics(mkey string, payload *models.WebfilterFtgdStatistics, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWebfilterFtgdStatistics(mkey string, params *models.CmdbRequestParams) error
	CreateSystemIke(payload *models.SystemIke, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemIke(mkey string, params *models.CmdbRequestParams) (*models.SystemIke, error)
	UpdateSystemIke(mkey string, payload *models.SystemIke, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemIke(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerHotspot20AnqpRoamingConsortium(payload *models.WirelessControllerHotspot20AnqpRoamingConsortium, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerHotspot20AnqpRoamingConsortium(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerHotspot20AnqpRoamingConsortium, error)
	UpdateWirelessControllerHotspot20AnqpRoamingConsortium(mkey string, payload *models.WirelessControllerHotspot20AnqpRoamingConsortium, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerHotspot20AnqpRoamingConsortium(mkey string, params *models.CmdbRequestParams) error
	ListWirelessControllerHotspot20AnqpRoamingConsortium(params *models.CmdbRequestParams) (*[]models.WirelessControllerHotspot20AnqpRoamingConsortium, error)
	CreateFirewallIpropeAppctrlStatus(payload *models.FirewallIpropeAppctrlStatus, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallIpropeAppctrlStatus(mkey string, params *models.CmdbRequestParams) (*models.FirewallIpropeAppctrlStatus, error)
	UpdateFirewallIpropeAppctrlStatus(mkey string, payload *models.FirewallIpropeAppctrlStatus, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallIpropeAppctrlStatus(mkey string, params *models.CmdbRequestParams) error
	CreateIpsSensor(payload *models.IpsSensor, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadIpsSensor(mkey string, params *models.CmdbRequestParams) (*models.IpsSensor, error)
	UpdateIpsSensor(mkey string, payload *models.IpsSensor, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteIpsSensor(mkey string, params *models.CmdbRequestParams) error
	ListIpsSensor(params *models.CmdbRequestParams) (*[]models.IpsSensor, error)
	CreateLogFortianalyzer3OverrideSetting(payload *models.LogFortianalyzer3OverrideSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogFortianalyzer3OverrideSetting(mkey string, params *models.CmdbRequestParams) (*models.LogFortianalyzer3OverrideSetting, error)
	UpdateLogFortianalyzer3OverrideSetting(mkey string, payload *models.LogFortianalyzer3OverrideSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogFortianalyzer3OverrideSetting(mkey string, params *models.CmdbRequestParams) error
	CreateSystemVxlan(payload *models.SystemVxlan, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemVxlan(mkey string, params *models.CmdbRequestParams) (*models.SystemVxlan, error)
	UpdateSystemVxlan(mkey string, payload *models.SystemVxlan, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemVxlan(mkey string, params *models.CmdbRequestParams) error
	ListSystemVxlan(params *models.CmdbRequestParams) (*[]models.SystemVxlan, error)
	CreateSwitchController8021XSettings(payload *models.SwitchController8021XSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchController8021XSettings(mkey string, params *models.CmdbRequestParams) (*models.SwitchController8021XSettings, error)
	UpdateSwitchController8021XSettings(mkey string, payload *models.SwitchController8021XSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchController8021XSettings(mkey string, params *models.CmdbRequestParams) error
	CreateEmailfilterProfile(payload *models.EmailfilterProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadEmailfilterProfile(mkey string, params *models.CmdbRequestParams) (*models.EmailfilterProfile, error)
	UpdateEmailfilterProfile(mkey string, payload *models.EmailfilterProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteEmailfilterProfile(mkey string, params *models.CmdbRequestParams) error
	ListEmailfilterProfile(params *models.CmdbRequestParams) (*[]models.EmailfilterProfile, error)
	CreateLogSyslogd4OverrideSetting(payload *models.LogSyslogd4OverrideSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogSyslogd4OverrideSetting(mkey string, params *models.CmdbRequestParams) (*models.LogSyslogd4OverrideSetting, error)
	UpdateLogSyslogd4OverrideSetting(mkey string, payload *models.LogSyslogd4OverrideSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogSyslogd4OverrideSetting(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallPolicy46(payload *models.FirewallPolicy46, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallPolicy46(mkey string, params *models.CmdbRequestParams) (*models.FirewallPolicy46, error)
	UpdateFirewallPolicy46(mkey string, payload *models.FirewallPolicy46, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallPolicy46(mkey string, params *models.CmdbRequestParams) error
	ListFirewallPolicy46(params *models.CmdbRequestParams) (*[]models.FirewallPolicy46, error)
	CreateSystemDnsServer(payload *models.SystemDnsServer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemDnsServer(mkey string, params *models.CmdbRequestParams) (*models.SystemDnsServer, error)
	UpdateSystemDnsServer(mkey string, payload *models.SystemDnsServer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemDnsServer(mkey string, params *models.CmdbRequestParams) error
	ListSystemDnsServer(params *models.CmdbRequestParams) (*[]models.SystemDnsServer, error)
	CreateSystemAutoInstall(payload *models.SystemAutoInstall, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemAutoInstall(mkey string, params *models.CmdbRequestParams) (*models.SystemAutoInstall, error)
	UpdateSystemAutoInstall(mkey string, payload *models.SystemAutoInstall, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemAutoInstall(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerSyslogProfile(payload *models.WirelessControllerSyslogProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerSyslogProfile(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerSyslogProfile, error)
	UpdateWirelessControllerSyslogProfile(mkey string, payload *models.WirelessControllerSyslogProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerSyslogProfile(mkey string, params *models.CmdbRequestParams) error
	ListWirelessControllerSyslogProfile(params *models.CmdbRequestParams) (*[]models.WirelessControllerSyslogProfile, error)
	CreateVpnCertificateOcspServer(payload *models.VpnCertificateOcspServer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnCertificateOcspServer(mkey string, params *models.CmdbRequestParams) (*models.VpnCertificateOcspServer, error)
	UpdateVpnCertificateOcspServer(mkey string, payload *models.VpnCertificateOcspServer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnCertificateOcspServer(mkey string, params *models.CmdbRequestParams) error
	ListVpnCertificateOcspServer(params *models.CmdbRequestParams) (*[]models.VpnCertificateOcspServer, error)
	CreateVpnPptp(payload *models.VpnPptp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnPptp(mkey string, params *models.CmdbRequestParams) (*models.VpnPptp, error)
	UpdateVpnPptp(mkey string, payload *models.VpnPptp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnPptp(mkey string, params *models.CmdbRequestParams) error
	CreateLogSyslogdOverrideSetting(payload *models.LogSyslogdOverrideSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogSyslogdOverrideSetting(mkey string, params *models.CmdbRequestParams) (*models.LogSyslogdOverrideSetting, error)
	UpdateLogSyslogdOverrideSetting(mkey string, payload *models.LogSyslogdOverrideSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogSyslogdOverrideSetting(mkey string, params *models.CmdbRequestParams) error
	CreateSystemDdns(payload *models.SystemDdns, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemDdns(mkey string, params *models.CmdbRequestParams) (*models.SystemDdns, error)
	UpdateSystemDdns(mkey string, payload *models.SystemDdns, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemDdns(mkey string, params *models.CmdbRequestParams) error
	ListSystemDdns(params *models.CmdbRequestParams) (*[]models.SystemDdns, error)
	CreateIpsSession(payload *models.IpsSession, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadIpsSession(mkey string, params *models.CmdbRequestParams) (*models.IpsSession, error)
	UpdateIpsSession(mkey string, payload *models.IpsSession, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteIpsSession(mkey string, params *models.CmdbRequestParams) error
	CreateApplicationList(payload *models.ApplicationList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadApplicationList(mkey string, params *models.CmdbRequestParams) (*models.ApplicationList, error)
	UpdateApplicationList(mkey string, payload *models.ApplicationList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteApplicationList(mkey string, params *models.CmdbRequestParams) error
	ListApplicationList(params *models.CmdbRequestParams) (*[]models.ApplicationList, error)
	CreateFirewallInternetServiceCustomGroup(payload *models.FirewallInternetServiceCustomGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallInternetServiceCustomGroup(mkey string, params *models.CmdbRequestParams) (*models.FirewallInternetServiceCustomGroup, error)
	UpdateFirewallInternetServiceCustomGroup(mkey string, payload *models.FirewallInternetServiceCustomGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallInternetServiceCustomGroup(mkey string, params *models.CmdbRequestParams) error
	ListFirewallInternetServiceCustomGroup(params *models.CmdbRequestParams) (*[]models.FirewallInternetServiceCustomGroup, error)
	CreateWirelessControllerBleProfile(payload *models.WirelessControllerBleProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerBleProfile(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerBleProfile, error)
	UpdateWirelessControllerBleProfile(mkey string, payload *models.WirelessControllerBleProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerBleProfile(mkey string, params *models.CmdbRequestParams) error
	ListWirelessControllerBleProfile(params *models.CmdbRequestParams) (*[]models.WirelessControllerBleProfile, error)
	CreateFirewallAddrgrp(payload *models.FirewallAddrgrp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallAddrgrp(mkey string, params *models.CmdbRequestParams) (*models.FirewallAddrgrp, error)
	UpdateFirewallAddrgrp(mkey string, payload *models.FirewallAddrgrp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallAddrgrp(mkey string, params *models.CmdbRequestParams) error
	ListFirewallAddrgrp(params *models.CmdbRequestParams) (*[]models.FirewallAddrgrp, error)
	CreateRouterAccessList(payload *models.RouterAccessList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterAccessList(mkey string, params *models.CmdbRequestParams) (*models.RouterAccessList, error)
	UpdateRouterAccessList(mkey string, payload *models.RouterAccessList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterAccessList(mkey string, params *models.CmdbRequestParams) error
	ListRouterAccessList(params *models.CmdbRequestParams) (*[]models.RouterAccessList, error)
	CreateWirelessControllerHotspot20H2qpOperatorName(payload *models.WirelessControllerHotspot20H2qpOperatorName, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerHotspot20H2qpOperatorName(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerHotspot20H2qpOperatorName, error)
	UpdateWirelessControllerHotspot20H2qpOperatorName(mkey string, payload *models.WirelessControllerHotspot20H2qpOperatorName, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerHotspot20H2qpOperatorName(mkey string, params *models.CmdbRequestParams) error
	ListWirelessControllerHotspot20H2qpOperatorName(params *models.CmdbRequestParams) (*[]models.WirelessControllerHotspot20H2qpOperatorName, error)
	CreateFirewallServiceCategory(payload *models.FirewallServiceCategory, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallServiceCategory(mkey string, params *models.CmdbRequestParams) (*models.FirewallServiceCategory, error)
	UpdateFirewallServiceCategory(mkey string, payload *models.FirewallServiceCategory, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallServiceCategory(mkey string, params *models.CmdbRequestParams) error
	ListFirewallServiceCategory(params *models.CmdbRequestParams) (*[]models.FirewallServiceCategory, error)
	CreateFirewallInternetServiceIpblVendor(payload *models.FirewallInternetServiceIpblVendor, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallInternetServiceIpblVendor(mkey string, params *models.CmdbRequestParams) (*models.FirewallInternetServiceIpblVendor, error)
	UpdateFirewallInternetServiceIpblVendor(mkey string, payload *models.FirewallInternetServiceIpblVendor, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallInternetServiceIpblVendor(mkey string, params *models.CmdbRequestParams) error
	ListFirewallInternetServiceIpblVendor(params *models.CmdbRequestParams) (*[]models.FirewallInternetServiceIpblVendor, error)
	CreateVpnCertificateRemote(payload *models.VpnCertificateRemote, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnCertificateRemote(mkey string, params *models.CmdbRequestParams) (*models.VpnCertificateRemote, error)
	UpdateVpnCertificateRemote(mkey string, payload *models.VpnCertificateRemote, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnCertificateRemote(mkey string, params *models.CmdbRequestParams) error
	ListVpnCertificateRemote(params *models.CmdbRequestParams) (*[]models.VpnCertificateRemote, error)
	CreateApplicationRuleSettings(payload *models.ApplicationRuleSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadApplicationRuleSettings(mkey string, params *models.CmdbRequestParams) (*models.ApplicationRuleSettings, error)
	UpdateApplicationRuleSettings(mkey string, payload *models.ApplicationRuleSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteApplicationRuleSettings(mkey string, params *models.CmdbRequestParams) error
	ListApplicationRuleSettings(params *models.CmdbRequestParams) (*[]models.ApplicationRuleSettings, error)
	CreateVideofilterProfile(payload *models.VideofilterProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVideofilterProfile(mkey string, params *models.CmdbRequestParams) (*models.VideofilterProfile, error)
	UpdateVideofilterProfile(mkey string, payload *models.VideofilterProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVideofilterProfile(mkey string, params *models.CmdbRequestParams) error
	ListVideofilterProfile(params *models.CmdbRequestParams) (*[]models.VideofilterProfile, error)
	CreateRouterKeyChain(payload *models.RouterKeyChain, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterKeyChain(mkey string, params *models.CmdbRequestParams) (*models.RouterKeyChain, error)
	UpdateRouterKeyChain(mkey string, payload *models.RouterKeyChain, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterKeyChain(mkey string, params *models.CmdbRequestParams) error
	ListRouterKeyChain(params *models.CmdbRequestParams) (*[]models.RouterKeyChain, error)
	CreateIpsecTunnel(payload *models.IpsecTunnel, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadIpsecTunnel(mkey string, params *models.CmdbRequestParams) (*models.IpsecTunnel, error)
	UpdateIpsecTunnel(mkey string, payload *models.IpsecTunnel, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteIpsecTunnel(mkey string, params *models.CmdbRequestParams) error
	CreateSystemReplacemsgAlertmail(payload *models.SystemReplacemsgAlertmail, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemReplacemsgAlertmail(mkey string, params *models.CmdbRequestParams) (*models.SystemReplacemsgAlertmail, error)
	UpdateSystemReplacemsgAlertmail(mkey string, payload *models.SystemReplacemsgAlertmail, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemReplacemsgAlertmail(mkey string, params *models.CmdbRequestParams) error
	ListSystemReplacemsgAlertmail(params *models.CmdbRequestParams) (*[]models.SystemReplacemsgAlertmail, error)
	CreateWirelessControllerWidsProfile(payload *models.WirelessControllerWidsProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerWidsProfile(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerWidsProfile, error)
	UpdateWirelessControllerWidsProfile(mkey string, payload *models.WirelessControllerWidsProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerWidsProfile(mkey string, params *models.CmdbRequestParams) error
	ListWirelessControllerWidsProfile(params *models.CmdbRequestParams) (*[]models.WirelessControllerWidsProfile, error)
	CreateFirewallIppool(payload *models.FirewallIppool, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallIppool(mkey string, params *models.CmdbRequestParams) (*models.FirewallIppool, error)
	UpdateFirewallIppool(mkey string, payload *models.FirewallIppool, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallIppool(mkey string, params *models.CmdbRequestParams) error
	ListFirewallIppool(params *models.CmdbRequestParams) (*[]models.FirewallIppool, error)
	CreateFirewallSshHostKey(payload *models.FirewallSshHostKey, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallSshHostKey(mkey string, params *models.CmdbRequestParams) (*models.FirewallSshHostKey, error)
	UpdateFirewallSshHostKey(mkey string, payload *models.FirewallSshHostKey, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallSshHostKey(mkey string, params *models.CmdbRequestParams) error
	ListFirewallSshHostKey(params *models.CmdbRequestParams) (*[]models.FirewallSshHostKey, error)
	CreateIpsGlobal(payload *models.IpsGlobal, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadIpsGlobal(mkey string, params *models.CmdbRequestParams) (*models.IpsGlobal, error)
	UpdateIpsGlobal(mkey string, payload *models.IpsGlobal, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteIpsGlobal(mkey string, params *models.CmdbRequestParams) error
	CreateEmailfilterMheader(payload *models.EmailfilterMheader, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadEmailfilterMheader(mkey string, params *models.CmdbRequestParams) (*models.EmailfilterMheader, error)
	UpdateEmailfilterMheader(mkey string, payload *models.EmailfilterMheader, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteEmailfilterMheader(mkey string, params *models.CmdbRequestParams) error
	ListEmailfilterMheader(params *models.CmdbRequestParams) (*[]models.EmailfilterMheader, error)
	CreateLogTacacsaccountingSetting(payload *models.LogTacacsaccountingSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogTacacsaccountingSetting(mkey string, params *models.CmdbRequestParams) (*models.LogTacacsaccountingSetting, error)
	UpdateLogTacacsaccountingSetting(mkey string, payload *models.LogTacacsaccountingSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogTacacsaccountingSetting(mkey string, params *models.CmdbRequestParams) error
	CreateSystemSitTunnel(payload *models.SystemSitTunnel, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemSitTunnel(mkey string, params *models.CmdbRequestParams) (*models.SystemSitTunnel, error)
	UpdateSystemSitTunnel(mkey string, payload *models.SystemSitTunnel, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemSitTunnel(mkey string, params *models.CmdbRequestParams) error
	ListSystemSitTunnel(params *models.CmdbRequestParams) (*[]models.SystemSitTunnel, error)
	CreateFirewallProute6(payload *models.FirewallProute6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallProute6(mkey string, params *models.CmdbRequestParams) (*models.FirewallProute6, error)
	UpdateFirewallProute6(mkey string, payload *models.FirewallProute6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallProute6(mkey string, params *models.CmdbRequestParams) error
	CreateLogEventfilter(payload *models.LogEventfilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogEventfilter(mkey string, params *models.CmdbRequestParams) (*models.LogEventfilter, error)
	UpdateLogEventfilter(mkey string, payload *models.LogEventfilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogEventfilter(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerWlchanlistlic(payload *models.WirelessControllerWlchanlistlic, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerWlchanlistlic(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerWlchanlistlic, error)
	UpdateWirelessControllerWlchanlistlic(mkey string, payload *models.WirelessControllerWlchanlistlic, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerWlchanlistlic(mkey string, params *models.CmdbRequestParams) error
	CreateSystemSessionInfoStatistics(payload *models.SystemSessionInfoStatistics, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemSessionInfoStatistics(mkey string, params *models.CmdbRequestParams) (*models.SystemSessionInfoStatistics, error)
	UpdateSystemSessionInfoStatistics(mkey string, payload *models.SystemSessionInfoStatistics, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemSessionInfoStatistics(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerWtp(payload *models.WirelessControllerWtp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerWtp(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerWtp, error)
	UpdateWirelessControllerWtp(mkey string, payload *models.WirelessControllerWtp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerWtp(mkey string, params *models.CmdbRequestParams) error
	ListWirelessControllerWtp(params *models.CmdbRequestParams) (*[]models.WirelessControllerWtp, error)
	CreateFirewallIpropeAppctrlList(payload *models.FirewallIpropeAppctrlList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallIpropeAppctrlList(mkey string, params *models.CmdbRequestParams) (*models.FirewallIpropeAppctrlList, error)
	UpdateFirewallIpropeAppctrlList(mkey string, payload *models.FirewallIpropeAppctrlList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallIpropeAppctrlList(mkey string, params *models.CmdbRequestParams) error
	CreateVpnIpsecStatsTunnel(payload *models.VpnIpsecStatsTunnel, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnIpsecStatsTunnel(mkey string, params *models.CmdbRequestParams) (*models.VpnIpsecStatsTunnel, error)
	UpdateVpnIpsecStatsTunnel(mkey string, payload *models.VpnIpsecStatsTunnel, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnIpsecStatsTunnel(mkey string, params *models.CmdbRequestParams) error
	CreateLogNullDeviceFilter(payload *models.LogNullDeviceFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogNullDeviceFilter(mkey string, params *models.CmdbRequestParams) (*models.LogNullDeviceFilter, error)
	UpdateLogNullDeviceFilter(mkey string, payload *models.LogNullDeviceFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogNullDeviceFilter(mkey string, params *models.CmdbRequestParams) error
	CreateSystemCentralMgmt(payload *models.SystemCentralMgmt, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemCentralMgmt(mkey string, params *models.CmdbRequestParams) (*models.SystemCentralMgmt, error)
	UpdateSystemCentralMgmt(mkey string, payload *models.SystemCentralMgmt, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemCentralMgmt(mkey string, params *models.CmdbRequestParams) error
	CreateSystemAutoUpdateVersions(payload *models.SystemAutoUpdateVersions, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemAutoUpdateVersions(mkey string, params *models.CmdbRequestParams) (*models.SystemAutoUpdateVersions, error)
	UpdateSystemAutoUpdateVersions(mkey string, payload *models.SystemAutoUpdateVersions, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemAutoUpdateVersions(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerHotspot20H2qpOsuProviderNai(payload *models.WirelessControllerHotspot20H2qpOsuProviderNai, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerHotspot20H2qpOsuProviderNai(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerHotspot20H2qpOsuProviderNai, error)
	UpdateWirelessControllerHotspot20H2qpOsuProviderNai(mkey string, payload *models.WirelessControllerHotspot20H2qpOsuProviderNai, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerHotspot20H2qpOsuProviderNai(mkey string, params *models.CmdbRequestParams) error
	ListWirelessControllerHotspot20H2qpOsuProviderNai(params *models.CmdbRequestParams) (*[]models.WirelessControllerHotspot20H2qpOsuProviderNai, error)
	CreateSwitchControllerSwitchInterfaceTag(payload *models.SwitchControllerSwitchInterfaceTag, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerSwitchInterfaceTag(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerSwitchInterfaceTag, error)
	UpdateSwitchControllerSwitchInterfaceTag(mkey string, payload *models.SwitchControllerSwitchInterfaceTag, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerSwitchInterfaceTag(mkey string, params *models.CmdbRequestParams) error
	ListSwitchControllerSwitchInterfaceTag(params *models.CmdbRequestParams) (*[]models.SwitchControllerSwitchInterfaceTag, error)
	CreateSwitchControllerRemoteLog(payload *models.SwitchControllerRemoteLog, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerRemoteLog(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerRemoteLog, error)
	UpdateSwitchControllerRemoteLog(mkey string, payload *models.SwitchControllerRemoteLog, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerRemoteLog(mkey string, params *models.CmdbRequestParams) error
	ListSwitchControllerRemoteLog(params *models.CmdbRequestParams) (*[]models.SwitchControllerRemoteLog, error)
	CreateFirewallMulticastPolicy(payload *models.FirewallMulticastPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallMulticastPolicy(mkey string, params *models.CmdbRequestParams) (*models.FirewallMulticastPolicy, error)
	UpdateFirewallMulticastPolicy(mkey string, payload *models.FirewallMulticastPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallMulticastPolicy(mkey string, params *models.CmdbRequestParams) error
	ListFirewallMulticastPolicy(params *models.CmdbRequestParams) (*[]models.FirewallMulticastPolicy, error)
	CreateFirewallSniffer(payload *models.FirewallSniffer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallSniffer(mkey string, params *models.CmdbRequestParams) (*models.FirewallSniffer, error)
	UpdateFirewallSniffer(mkey string, payload *models.FirewallSniffer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallSniffer(mkey string, params *models.CmdbRequestParams) error
	ListFirewallSniffer(params *models.CmdbRequestParams) (*[]models.FirewallSniffer, error)
	CreateSystemSession6(payload *models.SystemSession6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemSession6(mkey string, params *models.CmdbRequestParams) (*models.SystemSession6, error)
	UpdateSystemSession6(mkey string, payload *models.SystemSession6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemSession6(mkey string, params *models.CmdbRequestParams) error
	CreateLogFortianalyzerCloudSetting(payload *models.LogFortianalyzerCloudSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogFortianalyzerCloudSetting(mkey string, params *models.CmdbRequestParams) (*models.LogFortianalyzerCloudSetting, error)
	UpdateLogFortianalyzerCloudSetting(mkey string, payload *models.LogFortianalyzerCloudSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogFortianalyzerCloudSetting(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerBonjourProfile(payload *models.WirelessControllerBonjourProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerBonjourProfile(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerBonjourProfile, error)
	UpdateWirelessControllerBonjourProfile(mkey string, payload *models.WirelessControllerBonjourProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerBonjourProfile(mkey string, params *models.CmdbRequestParams) error
	ListWirelessControllerBonjourProfile(params *models.CmdbRequestParams) (*[]models.WirelessControllerBonjourProfile, error)
	CreateWirelessControllerNacProfile(payload *models.WirelessControllerNacProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerNacProfile(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerNacProfile, error)
	UpdateWirelessControllerNacProfile(mkey string, payload *models.WirelessControllerNacProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerNacProfile(mkey string, params *models.CmdbRequestParams) error
	ListWirelessControllerNacProfile(params *models.CmdbRequestParams) (*[]models.WirelessControllerNacProfile, error)
	CreateSwitchControllerQosDot1pMap(payload *models.SwitchControllerQosDot1pMap, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerQosDot1pMap(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerQosDot1pMap, error)
	UpdateSwitchControllerQosDot1pMap(mkey string, payload *models.SwitchControllerQosDot1pMap, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerQosDot1pMap(mkey string, params *models.CmdbRequestParams) error
	ListSwitchControllerQosDot1pMap(params *models.CmdbRequestParams) (*[]models.SwitchControllerQosDot1pMap, error)
	CreateAntivirusProfile(payload *models.AntivirusProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadAntivirusProfile(mkey string, params *models.CmdbRequestParams) (*models.AntivirusProfile, error)
	UpdateAntivirusProfile(mkey string, payload *models.AntivirusProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteAntivirusProfile(mkey string, params *models.CmdbRequestParams) error
	ListAntivirusProfile(params *models.CmdbRequestParams) (*[]models.AntivirusProfile, error)
	CreateSystemIpConflictStatus(payload *models.SystemIpConflictStatus, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemIpConflictStatus(mkey string, params *models.CmdbRequestParams) (*models.SystemIpConflictStatus, error)
	UpdateSystemIpConflictStatus(mkey string, payload *models.SystemIpConflictStatus, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemIpConflictStatus(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallMulticastAddress6(payload *models.FirewallMulticastAddress6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallMulticastAddress6(mkey string, params *models.CmdbRequestParams) (*models.FirewallMulticastAddress6, error)
	UpdateFirewallMulticastAddress6(mkey string, payload *models.FirewallMulticastAddress6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallMulticastAddress6(mkey string, params *models.CmdbRequestParams) error
	ListFirewallMulticastAddress6(params *models.CmdbRequestParams) (*[]models.FirewallMulticastAddress6, error)
	CreateFirewallAccessProxySshClientCert(payload *models.FirewallAccessProxySshClientCert, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallAccessProxySshClientCert(mkey string, params *models.CmdbRequestParams) (*models.FirewallAccessProxySshClientCert, error)
	UpdateFirewallAccessProxySshClientCert(mkey string, payload *models.FirewallAccessProxySshClientCert, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallAccessProxySshClientCert(mkey string, params *models.CmdbRequestParams) error
	ListFirewallAccessProxySshClientCert(params *models.CmdbRequestParams) (*[]models.FirewallAccessProxySshClientCert, error)
	CreateWebfilterOverride(payload *models.WebfilterOverride, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWebfilterOverride(mkey string, params *models.CmdbRequestParams) (*models.WebfilterOverride, error)
	UpdateWebfilterOverride(mkey string, payload *models.WebfilterOverride, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWebfilterOverride(mkey string, params *models.CmdbRequestParams) error
	ListWebfilterOverride(params *models.CmdbRequestParams) (*[]models.WebfilterOverride, error)
	CreateSwitchControllerSwitchLog(payload *models.SwitchControllerSwitchLog, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerSwitchLog(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerSwitchLog, error)
	UpdateSwitchControllerSwitchLog(mkey string, payload *models.SwitchControllerSwitchLog, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerSwitchLog(mkey string, params *models.CmdbRequestParams) error
	CreateVpnSslWebPortal(payload *models.VpnSslWebPortal, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnSslWebPortal(mkey string, params *models.CmdbRequestParams) (*models.VpnSslWebPortal, error)
	UpdateVpnSslWebPortal(mkey string, payload *models.VpnSslWebPortal, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnSslWebPortal(mkey string, params *models.CmdbRequestParams) error
	ListVpnSslWebPortal(params *models.CmdbRequestParams) (*[]models.VpnSslWebPortal, error)
	CreateWebProxyGlobal(payload *models.WebProxyGlobal, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWebProxyGlobal(mkey string, params *models.CmdbRequestParams) (*models.WebProxyGlobal, error)
	UpdateWebProxyGlobal(mkey string, payload *models.WebProxyGlobal, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWebProxyGlobal(mkey string, params *models.CmdbRequestParams) error
	CreateUserNacPolicy(payload *models.UserNacPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadUserNacPolicy(mkey string, params *models.CmdbRequestParams) (*models.UserNacPolicy, error)
	UpdateUserNacPolicy(mkey string, payload *models.UserNacPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteUserNacPolicy(mkey string, params *models.CmdbRequestParams) error
	ListUserNacPolicy(params *models.CmdbRequestParams) (*[]models.UserNacPolicy, error)
	CreateSystemDhcp6Server(payload *models.SystemDhcp6Server, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemDhcp6Server(mkey string, params *models.CmdbRequestParams) (*models.SystemDhcp6Server, error)
	UpdateSystemDhcp6Server(mkey string, payload *models.SystemDhcp6Server, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemDhcp6Server(mkey string, params *models.CmdbRequestParams) error
	ListSystemDhcp6Server(params *models.CmdbRequestParams) (*[]models.SystemDhcp6Server, error)
	CreateWirelessControllerHotspot20AnqpNaiRealm(payload *models.WirelessControllerHotspot20AnqpNaiRealm, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerHotspot20AnqpNaiRealm(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerHotspot20AnqpNaiRealm, error)
	UpdateWirelessControllerHotspot20AnqpNaiRealm(mkey string, payload *models.WirelessControllerHotspot20AnqpNaiRealm, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerHotspot20AnqpNaiRealm(mkey string, params *models.CmdbRequestParams) error
	ListWirelessControllerHotspot20AnqpNaiRealm(params *models.CmdbRequestParams) (*[]models.WirelessControllerHotspot20AnqpNaiRealm, error)
	CreateDlpSensor(payload *models.DlpSensor, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadDlpSensor(mkey string, params *models.CmdbRequestParams) (*models.DlpSensor, error)
	UpdateDlpSensor(mkey string, payload *models.DlpSensor, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteDlpSensor(mkey string, params *models.CmdbRequestParams) error
	ListDlpSensor(params *models.CmdbRequestParams) (*[]models.DlpSensor, error)
	CreateLogSyslogd3OverrideFilter(payload *models.LogSyslogd3OverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogSyslogd3OverrideFilter(mkey string, params *models.CmdbRequestParams) (*models.LogSyslogd3OverrideFilter, error)
	UpdateLogSyslogd3OverrideFilter(mkey string, payload *models.LogSyslogd3OverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogSyslogd3OverrideFilter(mkey string, params *models.CmdbRequestParams) error
	CreateAuthenticationScheme(payload *models.AuthenticationScheme, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadAuthenticationScheme(mkey string, params *models.CmdbRequestParams) (*models.AuthenticationScheme, error)
	UpdateAuthenticationScheme(mkey string, payload *models.AuthenticationScheme, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteAuthenticationScheme(mkey string, params *models.CmdbRequestParams) error
	ListAuthenticationScheme(params *models.CmdbRequestParams) (*[]models.AuthenticationScheme, error)
	CreateSystemPerformanceFirewallPacketDistribution(payload *models.SystemPerformanceFirewallPacketDistribution, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemPerformanceFirewallPacketDistribution(mkey string, params *models.CmdbRequestParams) (*models.SystemPerformanceFirewallPacketDistribution, error)
	UpdateSystemPerformanceFirewallPacketDistribution(mkey string, payload *models.SystemPerformanceFirewallPacketDistribution, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemPerformanceFirewallPacketDistribution(mkey string, params *models.CmdbRequestParams) error
	CreateLogTacacsaccountingFilter(payload *models.LogTacacsaccountingFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogTacacsaccountingFilter(mkey string, params *models.CmdbRequestParams) (*models.LogTacacsaccountingFilter, error)
	UpdateLogTacacsaccountingFilter(mkey string, payload *models.LogTacacsaccountingFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogTacacsaccountingFilter(mkey string, params *models.CmdbRequestParams) error
	CreateSystemApiUser(payload *models.SystemApiUser, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemApiUser(mkey string, params *models.CmdbRequestParams) (*models.SystemApiUser, error)
	UpdateSystemApiUser(mkey string, payload *models.SystemApiUser, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemApiUser(mkey string, params *models.CmdbRequestParams) error
	ListSystemApiUser(params *models.CmdbRequestParams) (*[]models.SystemApiUser, error)
	CreateSystemPerformanceFirewallStatistics(payload *models.SystemPerformanceFirewallStatistics, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemPerformanceFirewallStatistics(mkey string, params *models.CmdbRequestParams) (*models.SystemPerformanceFirewallStatistics, error)
	UpdateSystemPerformanceFirewallStatistics(mkey string, payload *models.SystemPerformanceFirewallStatistics, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemPerformanceFirewallStatistics(mkey string, params *models.CmdbRequestParams) error
	CreateSwitchControllerQuarantine(payload *models.SwitchControllerQuarantine, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerQuarantine(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerQuarantine, error)
	UpdateSwitchControllerQuarantine(mkey string, payload *models.SwitchControllerQuarantine, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerQuarantine(mkey string, params *models.CmdbRequestParams) error
	CreateUserGroup(payload *models.UserGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadUserGroup(mkey string, params *models.CmdbRequestParams) (*models.UserGroup, error)
	UpdateUserGroup(mkey string, payload *models.UserGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteUserGroup(mkey string, params *models.CmdbRequestParams) error
	ListUserGroup(params *models.CmdbRequestParams) (*[]models.UserGroup, error)
	CreateAntivirusSettings(payload *models.AntivirusSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadAntivirusSettings(mkey string, params *models.CmdbRequestParams) (*models.AntivirusSettings, error)
	UpdateAntivirusSettings(mkey string, payload *models.AntivirusSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteAntivirusSettings(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallVip46(payload *models.FirewallVip46, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallVip46(mkey string, params *models.CmdbRequestParams) (*models.FirewallVip46, error)
	UpdateFirewallVip46(mkey string, payload *models.FirewallVip46, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallVip46(mkey string, params *models.CmdbRequestParams) error
	ListFirewallVip46(params *models.CmdbRequestParams) (*[]models.FirewallVip46, error)
	CreateSystemVdomNetflow(payload *models.SystemVdomNetflow, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemVdomNetflow(mkey string, params *models.CmdbRequestParams) (*models.SystemVdomNetflow, error)
	UpdateSystemVdomNetflow(mkey string, payload *models.SystemVdomNetflow, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemVdomNetflow(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallShapingPolicy(payload *models.FirewallShapingPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallShapingPolicy(mkey string, params *models.CmdbRequestParams) (*models.FirewallShapingPolicy, error)
	UpdateFirewallShapingPolicy(mkey string, payload *models.FirewallShapingPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallShapingPolicy(mkey string, params *models.CmdbRequestParams) error
	ListFirewallShapingPolicy(params *models.CmdbRequestParams) (*[]models.FirewallShapingPolicy, error)
	CreateSystemReplacemsgHttp(payload *models.SystemReplacemsgHttp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemReplacemsgHttp(mkey string, params *models.CmdbRequestParams) (*models.SystemReplacemsgHttp, error)
	UpdateSystemReplacemsgHttp(mkey string, payload *models.SystemReplacemsgHttp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemReplacemsgHttp(mkey string, params *models.CmdbRequestParams) error
	ListSystemReplacemsgHttp(params *models.CmdbRequestParams) (*[]models.SystemReplacemsgHttp, error)
	CreateSystemIpv6Tunnel(payload *models.SystemIpv6Tunnel, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemIpv6Tunnel(mkey string, params *models.CmdbRequestParams) (*models.SystemIpv6Tunnel, error)
	UpdateSystemIpv6Tunnel(mkey string, payload *models.SystemIpv6Tunnel, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemIpv6Tunnel(mkey string, params *models.CmdbRequestParams) error
	ListSystemIpv6Tunnel(params *models.CmdbRequestParams) (*[]models.SystemIpv6Tunnel, error)
	CreateSystemClusterSync(payload *models.SystemClusterSync, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemClusterSync(mkey string, params *models.CmdbRequestParams) (*models.SystemClusterSync, error)
	UpdateSystemClusterSync(mkey string, payload *models.SystemClusterSync, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemClusterSync(mkey string, params *models.CmdbRequestParams) error
	ListSystemClusterSync(params *models.CmdbRequestParams) (*[]models.SystemClusterSync, error)
	CreateSystemAutoUpdateStatus(payload *models.SystemAutoUpdateStatus, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemAutoUpdateStatus(mkey string, params *models.CmdbRequestParams) (*models.SystemAutoUpdateStatus, error)
	UpdateSystemAutoUpdateStatus(mkey string, payload *models.SystemAutoUpdateStatus, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemAutoUpdateStatus(mkey string, params *models.CmdbRequestParams) error
	CreateWebProxyForwardServer(payload *models.WebProxyForwardServer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWebProxyForwardServer(mkey string, params *models.CmdbRequestParams) (*models.WebProxyForwardServer, error)
	UpdateWebProxyForwardServer(mkey string, payload *models.WebProxyForwardServer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWebProxyForwardServer(mkey string, params *models.CmdbRequestParams) error
	ListWebProxyForwardServer(params *models.CmdbRequestParams) (*[]models.WebProxyForwardServer, error)
	CreateLogFortianalyzerSetting(payload *models.LogFortianalyzerSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogFortianalyzerSetting(mkey string, params *models.CmdbRequestParams) (*models.LogFortianalyzerSetting, error)
	UpdateLogFortianalyzerSetting(mkey string, payload *models.LogFortianalyzerSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogFortianalyzerSetting(mkey string, params *models.CmdbRequestParams) error
	CreateIcapProfile(payload *models.IcapProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadIcapProfile(mkey string, params *models.CmdbRequestParams) (*models.IcapProfile, error)
	UpdateIcapProfile(mkey string, payload *models.IcapProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteIcapProfile(mkey string, params *models.CmdbRequestParams) error
	ListIcapProfile(params *models.CmdbRequestParams) (*[]models.IcapProfile, error)
	CreateSystemVdomDns(payload *models.SystemVdomDns, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemVdomDns(mkey string, params *models.CmdbRequestParams) (*models.SystemVdomDns, error)
	UpdateSystemVdomDns(mkey string, payload *models.SystemVdomDns, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemVdomDns(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerHotspot20AnqpNetworkAuthType(payload *models.WirelessControllerHotspot20AnqpNetworkAuthType, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerHotspot20AnqpNetworkAuthType(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerHotspot20AnqpNetworkAuthType, error)
	UpdateWirelessControllerHotspot20AnqpNetworkAuthType(mkey string, payload *models.WirelessControllerHotspot20AnqpNetworkAuthType, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerHotspot20AnqpNetworkAuthType(mkey string, params *models.CmdbRequestParams) error
	ListWirelessControllerHotspot20AnqpNetworkAuthType(params *models.CmdbRequestParams) (*[]models.WirelessControllerHotspot20AnqpNetworkAuthType, error)
	CreateSystemReplacemsgSpam(payload *models.SystemReplacemsgSpam, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemReplacemsgSpam(mkey string, params *models.CmdbRequestParams) (*models.SystemReplacemsgSpam, error)
	UpdateSystemReplacemsgSpam(mkey string, payload *models.SystemReplacemsgSpam, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemReplacemsgSpam(mkey string, params *models.CmdbRequestParams) error
	ListSystemReplacemsgSpam(params *models.CmdbRequestParams) (*[]models.SystemReplacemsgSpam, error)
	CreateSystemCmdb(payload *models.SystemCmdb, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemCmdb(mkey string, params *models.CmdbRequestParams) (*models.SystemCmdb, error)
	UpdateSystemCmdb(mkey string, payload *models.SystemCmdb, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemCmdb(mkey string, params *models.CmdbRequestParams) error
	CreateSystemChecksumStatus(payload *models.SystemChecksumStatus, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemChecksumStatus(mkey string, params *models.CmdbRequestParams) (*models.SystemChecksumStatus, error)
	UpdateSystemChecksumStatus(mkey string, payload *models.SystemChecksumStatus, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemChecksumStatus(mkey string, params *models.CmdbRequestParams) error
	CreateSystemFortisandbox(payload *models.SystemFortisandbox, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemFortisandbox(mkey string, params *models.CmdbRequestParams) (*models.SystemFortisandbox, error)
	UpdateSystemFortisandbox(mkey string, payload *models.SystemFortisandbox, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemFortisandbox(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallServiceGroup(payload *models.FirewallServiceGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallServiceGroup(mkey string, params *models.CmdbRequestParams) (*models.FirewallServiceGroup, error)
	UpdateFirewallServiceGroup(mkey string, payload *models.FirewallServiceGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallServiceGroup(mkey string, params *models.CmdbRequestParams) error
	ListFirewallServiceGroup(params *models.CmdbRequestParams) (*[]models.FirewallServiceGroup, error)
	CreateSystemReplacemsgFortiguardWf(payload *models.SystemReplacemsgFortiguardWf, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemReplacemsgFortiguardWf(mkey string, params *models.CmdbRequestParams) (*models.SystemReplacemsgFortiguardWf, error)
	UpdateSystemReplacemsgFortiguardWf(mkey string, payload *models.SystemReplacemsgFortiguardWf, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemReplacemsgFortiguardWf(mkey string, params *models.CmdbRequestParams) error
	ListSystemReplacemsgFortiguardWf(params *models.CmdbRequestParams) (*[]models.SystemReplacemsgFortiguardWf, error)
	CreateFirewallWildcardFqdnGroup(payload *models.FirewallWildcardFqdnGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallWildcardFqdnGroup(mkey string, params *models.CmdbRequestParams) (*models.FirewallWildcardFqdnGroup, error)
	UpdateFirewallWildcardFqdnGroup(mkey string, payload *models.FirewallWildcardFqdnGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallWildcardFqdnGroup(mkey string, params *models.CmdbRequestParams) error
	ListFirewallWildcardFqdnGroup(params *models.CmdbRequestParams) (*[]models.FirewallWildcardFqdnGroup, error)
	CreateSystemNat64(payload *models.SystemNat64, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemNat64(mkey string, params *models.CmdbRequestParams) (*models.SystemNat64, error)
	UpdateSystemNat64(mkey string, payload *models.SystemNat64, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemNat64(mkey string, params *models.CmdbRequestParams) error
	CreateSystemFortimanager(payload *models.SystemFortimanager, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemFortimanager(mkey string, params *models.CmdbRequestParams) (*models.SystemFortimanager, error)
	UpdateSystemFortimanager(mkey string, payload *models.SystemFortimanager, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemFortimanager(mkey string, params *models.CmdbRequestParams) error
	CreateSystemAutoupdateSchedule(payload *models.SystemAutoupdateSchedule, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemAutoupdateSchedule(mkey string, params *models.CmdbRequestParams) (*models.SystemAutoupdateSchedule, error)
	UpdateSystemAutoupdateSchedule(mkey string, payload *models.SystemAutoupdateSchedule, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemAutoupdateSchedule(mkey string, params *models.CmdbRequestParams) error
	CreateSystemPtp(payload *models.SystemPtp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemPtp(mkey string, params *models.CmdbRequestParams) (*models.SystemPtp, error)
	UpdateSystemPtp(mkey string, payload *models.SystemPtp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemPtp(mkey string, params *models.CmdbRequestParams) error
	CreateWebfilterSearchEngine(payload *models.WebfilterSearchEngine, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWebfilterSearchEngine(mkey string, params *models.CmdbRequestParams) (*models.WebfilterSearchEngine, error)
	UpdateWebfilterSearchEngine(mkey string, payload *models.WebfilterSearchEngine, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWebfilterSearchEngine(mkey string, params *models.CmdbRequestParams) error
	ListWebfilterSearchEngine(params *models.CmdbRequestParams) (*[]models.WebfilterSearchEngine, error)
	CreateLogFortianalyzerOverrideFilter(payload *models.LogFortianalyzerOverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogFortianalyzerOverrideFilter(mkey string, params *models.CmdbRequestParams) (*models.LogFortianalyzerOverrideFilter, error)
	UpdateLogFortianalyzerOverrideFilter(mkey string, payload *models.LogFortianalyzerOverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogFortianalyzerOverrideFilter(mkey string, params *models.CmdbRequestParams) error
	CreateUserSetting(payload *models.UserSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadUserSetting(mkey string, params *models.CmdbRequestParams) (*models.UserSetting, error)
	UpdateUserSetting(mkey string, payload *models.UserSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteUserSetting(mkey string, params *models.CmdbRequestParams) error
	CreateSystemAutoScript(payload *models.SystemAutoScript, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemAutoScript(mkey string, params *models.CmdbRequestParams) (*models.SystemAutoScript, error)
	UpdateSystemAutoScript(mkey string, payload *models.SystemAutoScript, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemAutoScript(mkey string, params *models.CmdbRequestParams) error
	ListSystemAutoScript(params *models.CmdbRequestParams) (*[]models.SystemAutoScript, error)
	CreateWirelessControllerRegion(payload *models.WirelessControllerRegion, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerRegion(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerRegion, error)
	UpdateWirelessControllerRegion(mkey string, payload *models.WirelessControllerRegion, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerRegion(mkey string, params *models.CmdbRequestParams) error
	ListWirelessControllerRegion(params *models.CmdbRequestParams) (*[]models.WirelessControllerRegion, error)
	CreateWirelessControllerWtpGroup(payload *models.WirelessControllerWtpGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerWtpGroup(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerWtpGroup, error)
	UpdateWirelessControllerWtpGroup(mkey string, payload *models.WirelessControllerWtpGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerWtpGroup(mkey string, params *models.CmdbRequestParams) error
	ListWirelessControllerWtpGroup(params *models.CmdbRequestParams) (*[]models.WirelessControllerWtpGroup, error)
	CreateSwitchControllerQosQosPolicy(payload *models.SwitchControllerQosQosPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerQosQosPolicy(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerQosQosPolicy, error)
	UpdateSwitchControllerQosQosPolicy(mkey string, payload *models.SwitchControllerQosQosPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerQosQosPolicy(mkey string, params *models.CmdbRequestParams) error
	ListSwitchControllerQosQosPolicy(params *models.CmdbRequestParams) (*[]models.SwitchControllerQosQosPolicy, error)
	CreateSystemFm(payload *models.SystemFm, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemFm(mkey string, params *models.CmdbRequestParams) (*models.SystemFm, error)
	UpdateSystemFm(mkey string, payload *models.SystemFm, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemFm(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerHotspot20H2qpConnCapability(payload *models.WirelessControllerHotspot20H2qpConnCapability, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerHotspot20H2qpConnCapability(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerHotspot20H2qpConnCapability, error)
	UpdateWirelessControllerHotspot20H2qpConnCapability(mkey string, payload *models.WirelessControllerHotspot20H2qpConnCapability, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerHotspot20H2qpConnCapability(mkey string, params *models.CmdbRequestParams) error
	ListWirelessControllerHotspot20H2qpConnCapability(params *models.CmdbRequestParams) (*[]models.WirelessControllerHotspot20H2qpConnCapability, error)
	CreateSystemAutomationAction(payload *models.SystemAutomationAction, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemAutomationAction(mkey string, params *models.CmdbRequestParams) (*models.SystemAutomationAction, error)
	UpdateSystemAutomationAction(mkey string, payload *models.SystemAutomationAction, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemAutomationAction(mkey string, params *models.CmdbRequestParams) error
	ListSystemAutomationAction(params *models.CmdbRequestParams) (*[]models.SystemAutomationAction, error)
	CreateFirewallSecurityPolicy(payload *models.FirewallSecurityPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallSecurityPolicy(mkey string, params *models.CmdbRequestParams) (*models.FirewallSecurityPolicy, error)
	UpdateFirewallSecurityPolicy(mkey string, payload *models.FirewallSecurityPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallSecurityPolicy(mkey string, params *models.CmdbRequestParams) error
	ListFirewallSecurityPolicy(params *models.CmdbRequestParams) (*[]models.FirewallSecurityPolicy, error)
	CreateLogFortiguardSetting(payload *models.LogFortiguardSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogFortiguardSetting(mkey string, params *models.CmdbRequestParams) (*models.LogFortiguardSetting, error)
	UpdateLogFortiguardSetting(mkey string, payload *models.LogFortiguardSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogFortiguardSetting(mkey string, params *models.CmdbRequestParams) error
	CreateRouterRouteMap(payload *models.RouterRouteMap, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterRouteMap(mkey string, params *models.CmdbRequestParams) (*models.RouterRouteMap, error)
	UpdateRouterRouteMap(mkey string, payload *models.RouterRouteMap, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterRouteMap(mkey string, params *models.CmdbRequestParams) error
	ListRouterRouteMap(params *models.CmdbRequestParams) (*[]models.RouterRouteMap, error)
	CreateExtenderExtenderInfo(payload *models.ExtenderExtenderInfo, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadExtenderExtenderInfo(mkey string, params *models.CmdbRequestParams) (*models.ExtenderExtenderInfo, error)
	UpdateExtenderExtenderInfo(mkey string, payload *models.ExtenderExtenderInfo, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteExtenderExtenderInfo(mkey string, params *models.CmdbRequestParams) error
	CreateSystemSourceIpStatus(payload *models.SystemSourceIpStatus, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemSourceIpStatus(mkey string, params *models.CmdbRequestParams) (*models.SystemSourceIpStatus, error)
	UpdateSystemSourceIpStatus(mkey string, payload *models.SystemSourceIpStatus, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemSourceIpStatus(mkey string, params *models.CmdbRequestParams) error
	CreateLogFortianalyzer3Setting(payload *models.LogFortianalyzer3Setting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogFortianalyzer3Setting(mkey string, params *models.CmdbRequestParams) (*models.LogFortianalyzer3Setting, error)
	UpdateLogFortianalyzer3Setting(mkey string, payload *models.LogFortianalyzer3Setting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogFortianalyzer3Setting(mkey string, params *models.CmdbRequestParams) error
	CreateVoipProfile(payload *models.VoipProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVoipProfile(mkey string, params *models.CmdbRequestParams) (*models.VoipProfile, error)
	UpdateVoipProfile(mkey string, payload *models.VoipProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVoipProfile(mkey string, params *models.CmdbRequestParams) error
	ListVoipProfile(params *models.CmdbRequestParams) (*[]models.VoipProfile, error)
	CreateRouterAspathList(payload *models.RouterAspathList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterAspathList(mkey string, params *models.CmdbRequestParams) (*models.RouterAspathList, error)
	UpdateRouterAspathList(mkey string, payload *models.RouterAspathList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterAspathList(mkey string, params *models.CmdbRequestParams) error
	ListRouterAspathList(params *models.CmdbRequestParams) (*[]models.RouterAspathList, error)
	CreateRouterMulticast(payload *models.RouterMulticast, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterMulticast(mkey string, params *models.CmdbRequestParams) (*models.RouterMulticast, error)
	UpdateRouterMulticast(mkey string, payload *models.RouterMulticast, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterMulticast(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallProfileProtocolOptions(payload *models.FirewallProfileProtocolOptions, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallProfileProtocolOptions(mkey string, params *models.CmdbRequestParams) (*models.FirewallProfileProtocolOptions, error)
	UpdateFirewallProfileProtocolOptions(mkey string, payload *models.FirewallProfileProtocolOptions, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallProfileProtocolOptions(mkey string, params *models.CmdbRequestParams) error
	ListFirewallProfileProtocolOptions(params *models.CmdbRequestParams) (*[]models.FirewallProfileProtocolOptions, error)
	CreateFirewallInternetServiceIpblReason(payload *models.FirewallInternetServiceIpblReason, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallInternetServiceIpblReason(mkey string, params *models.CmdbRequestParams) (*models.FirewallInternetServiceIpblReason, error)
	UpdateFirewallInternetServiceIpblReason(mkey string, payload *models.FirewallInternetServiceIpblReason, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallInternetServiceIpblReason(mkey string, params *models.CmdbRequestParams) error
	ListFirewallInternetServiceIpblReason(params *models.CmdbRequestParams) (*[]models.FirewallInternetServiceIpblReason, error)
	CreateFirewallIdentityBasedRoute(payload *models.FirewallIdentityBasedRoute, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallIdentityBasedRoute(mkey string, params *models.CmdbRequestParams) (*models.FirewallIdentityBasedRoute, error)
	UpdateFirewallIdentityBasedRoute(mkey string, payload *models.FirewallIdentityBasedRoute, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallIdentityBasedRoute(mkey string, params *models.CmdbRequestParams) error
	ListFirewallIdentityBasedRoute(params *models.CmdbRequestParams) (*[]models.FirewallIdentityBasedRoute, error)
	CreateWebfilterFortiguard(payload *models.WebfilterFortiguard, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWebfilterFortiguard(mkey string, params *models.CmdbRequestParams) (*models.WebfilterFortiguard, error)
	UpdateWebfilterFortiguard(mkey string, payload *models.WebfilterFortiguard, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWebfilterFortiguard(mkey string, params *models.CmdbRequestParams) error
	CreateLogSyslogdOverrideFilter(payload *models.LogSyslogdOverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogSyslogdOverrideFilter(mkey string, params *models.CmdbRequestParams) (*models.LogSyslogdOverrideFilter, error)
	UpdateLogSyslogdOverrideFilter(mkey string, payload *models.LogSyslogdOverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogSyslogdOverrideFilter(mkey string, params *models.CmdbRequestParams) error
	CreateUserSecurityExemptList(payload *models.UserSecurityExemptList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadUserSecurityExemptList(mkey string, params *models.CmdbRequestParams) (*models.UserSecurityExemptList, error)
	UpdateUserSecurityExemptList(mkey string, payload *models.UserSecurityExemptList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteUserSecurityExemptList(mkey string, params *models.CmdbRequestParams) error
	ListUserSecurityExemptList(params *models.CmdbRequestParams) (*[]models.UserSecurityExemptList, error)
	CreateHardwareMemory(payload *models.HardwareMemory, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadHardwareMemory(mkey string, params *models.CmdbRequestParams) (*models.HardwareMemory, error)
	UpdateHardwareMemory(mkey string, payload *models.HardwareMemory, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteHardwareMemory(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerWtpProfile(payload *models.WirelessControllerWtpProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerWtpProfile(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerWtpProfile, error)
	UpdateWirelessControllerWtpProfile(mkey string, payload *models.WirelessControllerWtpProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerWtpProfile(mkey string, params *models.CmdbRequestParams) error
	ListWirelessControllerWtpProfile(params *models.CmdbRequestParams) (*[]models.WirelessControllerWtpProfile, error)
	CreateSystemAutomationStitch(payload *models.SystemAutomationStitch, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemAutomationStitch(mkey string, params *models.CmdbRequestParams) (*models.SystemAutomationStitch, error)
	UpdateSystemAutomationStitch(mkey string, payload *models.SystemAutomationStitch, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemAutomationStitch(mkey string, params *models.CmdbRequestParams) error
	ListSystemAutomationStitch(params *models.CmdbRequestParams) (*[]models.SystemAutomationStitch, error)
	CreateSystemInterface(payload *models.SystemInterface, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemInterface(mkey string, params *models.CmdbRequestParams) (*models.SystemInterface, error)
	UpdateSystemInterface(mkey string, payload *models.SystemInterface, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemInterface(mkey string, params *models.CmdbRequestParams) error
	ListSystemInterface(params *models.CmdbRequestParams) (*[]models.SystemInterface, error)
	CreateSystemCentralManagement(payload *models.SystemCentralManagement, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemCentralManagement(mkey string, params *models.CmdbRequestParams) (*models.SystemCentralManagement, error)
	UpdateSystemCentralManagement(mkey string, payload *models.SystemCentralManagement, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemCentralManagement(mkey string, params *models.CmdbRequestParams) error
	CreateSystemProbeResponse(payload *models.SystemProbeResponse, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemProbeResponse(mkey string, params *models.CmdbRequestParams) (*models.SystemProbeResponse, error)
	UpdateSystemProbeResponse(mkey string, payload *models.SystemProbeResponse, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemProbeResponse(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallIpmacbindingSetting(payload *models.FirewallIpmacbindingSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallIpmacbindingSetting(mkey string, params *models.CmdbRequestParams) (*models.FirewallIpmacbindingSetting, error)
	UpdateFirewallIpmacbindingSetting(mkey string, payload *models.FirewallIpmacbindingSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallIpmacbindingSetting(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallTtlPolicy(payload *models.FirewallTtlPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallTtlPolicy(mkey string, params *models.CmdbRequestParams) (*models.FirewallTtlPolicy, error)
	UpdateFirewallTtlPolicy(mkey string, payload *models.FirewallTtlPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallTtlPolicy(mkey string, params *models.CmdbRequestParams) error
	ListFirewallTtlPolicy(params *models.CmdbRequestParams) (*[]models.FirewallTtlPolicy, error)
	CreateLogSyslogd2Setting(payload *models.LogSyslogd2Setting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogSyslogd2Setting(mkey string, params *models.CmdbRequestParams) (*models.LogSyslogd2Setting, error)
	UpdateLogSyslogd2Setting(mkey string, payload *models.LogSyslogd2Setting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogSyslogd2Setting(mkey string, params *models.CmdbRequestParams) error
	CreateHardwareCpu(payload *models.HardwareCpu, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadHardwareCpu(mkey string, params *models.CmdbRequestParams) (*models.HardwareCpu, error)
	UpdateHardwareCpu(mkey string, payload *models.HardwareCpu, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteHardwareCpu(mkey string, params *models.CmdbRequestParams) error
	CreateSystemStp(payload *models.SystemStp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemStp(mkey string, params *models.CmdbRequestParams) (*models.SystemStp, error)
	UpdateSystemStp(mkey string, payload *models.SystemStp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemStp(mkey string, params *models.CmdbRequestParams) error
	CreateSwitchControllerNacSettings(payload *models.SwitchControllerNacSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerNacSettings(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerNacSettings, error)
	UpdateSwitchControllerNacSettings(mkey string, payload *models.SwitchControllerNacSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerNacSettings(mkey string, params *models.CmdbRequestParams) error
	ListSwitchControllerNacSettings(params *models.CmdbRequestParams) (*[]models.SwitchControllerNacSettings, error)
	CreateFirewallServiceCustom(payload *models.FirewallServiceCustom, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallServiceCustom(mkey string, params *models.CmdbRequestParams) (*models.FirewallServiceCustom, error)
	UpdateFirewallServiceCustom(mkey string, payload *models.FirewallServiceCustom, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallServiceCustom(mkey string, params *models.CmdbRequestParams) error
	ListFirewallServiceCustom(params *models.CmdbRequestParams) (*[]models.FirewallServiceCustom, error)
	CreateFirewallLdbMonitor(payload *models.FirewallLdbMonitor, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallLdbMonitor(mkey string, params *models.CmdbRequestParams) (*models.FirewallLdbMonitor, error)
	UpdateFirewallLdbMonitor(mkey string, payload *models.FirewallLdbMonitor, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallLdbMonitor(mkey string, params *models.CmdbRequestParams) error
	ListFirewallLdbMonitor(params *models.CmdbRequestParams) (*[]models.FirewallLdbMonitor, error)
	CreateWebProxyDebugUrl(payload *models.WebProxyDebugUrl, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWebProxyDebugUrl(mkey string, params *models.CmdbRequestParams) (*models.WebProxyDebugUrl, error)
	UpdateWebProxyDebugUrl(mkey string, payload *models.WebProxyDebugUrl, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWebProxyDebugUrl(mkey string, params *models.CmdbRequestParams) error
	ListWebProxyDebugUrl(params *models.CmdbRequestParams) (*[]models.WebProxyDebugUrl, error)
	CreateLogFortianalyzer2OverrideFilter(payload *models.LogFortianalyzer2OverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogFortianalyzer2OverrideFilter(mkey string, params *models.CmdbRequestParams) (*models.LogFortianalyzer2OverrideFilter, error)
	UpdateLogFortianalyzer2OverrideFilter(mkey string, payload *models.LogFortianalyzer2OverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogFortianalyzer2OverrideFilter(mkey string, params *models.CmdbRequestParams) error
	CreateRouterPrefixList6(payload *models.RouterPrefixList6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterPrefixList6(mkey string, params *models.CmdbRequestParams) (*models.RouterPrefixList6, error)
	UpdateRouterPrefixList6(mkey string, payload *models.RouterPrefixList6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterPrefixList6(mkey string, params *models.CmdbRequestParams) error
	ListRouterPrefixList6(params *models.CmdbRequestParams) (*[]models.RouterPrefixList6, error)
	CreateSwitchControllerPtpPolicy(payload *models.SwitchControllerPtpPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerPtpPolicy(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerPtpPolicy, error)
	UpdateSwitchControllerPtpPolicy(mkey string, payload *models.SwitchControllerPtpPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerPtpPolicy(mkey string, params *models.CmdbRequestParams) error
	ListSwitchControllerPtpPolicy(params *models.CmdbRequestParams) (*[]models.SwitchControllerPtpPolicy, error)
	CreateFirewallIpTranslation(payload *models.FirewallIpTranslation, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallIpTranslation(mkey string, params *models.CmdbRequestParams) (*models.FirewallIpTranslation, error)
	UpdateFirewallIpTranslation(mkey string, payload *models.FirewallIpTranslation, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallIpTranslation(mkey string, params *models.CmdbRequestParams) error
	ListFirewallIpTranslation(params *models.CmdbRequestParams) (*[]models.FirewallIpTranslation, error)
	CreateFirewallAcl(payload *models.FirewallAcl, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallAcl(mkey string, params *models.CmdbRequestParams) (*models.FirewallAcl, error)
	UpdateFirewallAcl(mkey string, payload *models.FirewallAcl, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallAcl(mkey string, params *models.CmdbRequestParams) error
	ListFirewallAcl(params *models.CmdbRequestParams) (*[]models.FirewallAcl, error)
	CreateWirelessControllerHotspot20QosMap(payload *models.WirelessControllerHotspot20QosMap, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerHotspot20QosMap(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerHotspot20QosMap, error)
	UpdateWirelessControllerHotspot20QosMap(mkey string, payload *models.WirelessControllerHotspot20QosMap, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerHotspot20QosMap(mkey string, params *models.CmdbRequestParams) error
	ListWirelessControllerHotspot20QosMap(params *models.CmdbRequestParams) (*[]models.WirelessControllerHotspot20QosMap, error)
	CreateSwitchControllerQosQueuePolicy(payload *models.SwitchControllerQosQueuePolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerQosQueuePolicy(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerQosQueuePolicy, error)
	UpdateSwitchControllerQosQueuePolicy(mkey string, payload *models.SwitchControllerQosQueuePolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerQosQueuePolicy(mkey string, params *models.CmdbRequestParams) error
	ListSwitchControllerQosQueuePolicy(params *models.CmdbRequestParams) (*[]models.SwitchControllerQosQueuePolicy, error)
	CreateIpsSettings(payload *models.IpsSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadIpsSettings(mkey string, params *models.CmdbRequestParams) (*models.IpsSettings, error)
	UpdateIpsSettings(mkey string, payload *models.IpsSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteIpsSettings(mkey string, params *models.CmdbRequestParams) error
	CreateLogSyslogd2Filter(payload *models.LogSyslogd2Filter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogSyslogd2Filter(mkey string, params *models.CmdbRequestParams) (*models.LogSyslogd2Filter, error)
	UpdateLogSyslogd2Filter(mkey string, payload *models.LogSyslogd2Filter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogSyslogd2Filter(mkey string, params *models.CmdbRequestParams) error
	CreateSystemIpv6NeighborCache(payload *models.SystemIpv6NeighborCache, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemIpv6NeighborCache(mkey string, params *models.CmdbRequestParams) (*models.SystemIpv6NeighborCache, error)
	UpdateSystemIpv6NeighborCache(mkey string, payload *models.SystemIpv6NeighborCache, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemIpv6NeighborCache(mkey string, params *models.CmdbRequestParams) error
	ListSystemIpv6NeighborCache(params *models.CmdbRequestParams) (*[]models.SystemIpv6NeighborCache, error)
	CreateFirewallInternetServiceGroup(payload *models.FirewallInternetServiceGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallInternetServiceGroup(mkey string, params *models.CmdbRequestParams) (*models.FirewallInternetServiceGroup, error)
	UpdateFirewallInternetServiceGroup(mkey string, payload *models.FirewallInternetServiceGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallInternetServiceGroup(mkey string, params *models.CmdbRequestParams) error
	ListFirewallInternetServiceGroup(params *models.CmdbRequestParams) (*[]models.FirewallInternetServiceGroup, error)
	CreateLogSyslogd4Filter(payload *models.LogSyslogd4Filter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogSyslogd4Filter(mkey string, params *models.CmdbRequestParams) (*models.LogSyslogd4Filter, error)
	UpdateLogSyslogd4Filter(mkey string, payload *models.LogSyslogd4Filter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogSyslogd4Filter(mkey string, params *models.CmdbRequestParams) error
	CreateSystemLinkMonitor(payload *models.SystemLinkMonitor, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemLinkMonitor(mkey string, params *models.CmdbRequestParams) (*models.SystemLinkMonitor, error)
	UpdateSystemLinkMonitor(mkey string, payload *models.SystemLinkMonitor, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemLinkMonitor(mkey string, params *models.CmdbRequestParams) error
	ListSystemLinkMonitor(params *models.CmdbRequestParams) (*[]models.SystemLinkMonitor, error)
	CreateSwitchControllerLldpProfile(payload *models.SwitchControllerLldpProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerLldpProfile(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerLldpProfile, error)
	UpdateSwitchControllerLldpProfile(mkey string, payload *models.SwitchControllerLldpProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerLldpProfile(mkey string, params *models.CmdbRequestParams) error
	ListSwitchControllerLldpProfile(params *models.CmdbRequestParams) (*[]models.SwitchControllerLldpProfile, error)
	CreateFirewallAddrgrp6(payload *models.FirewallAddrgrp6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallAddrgrp6(mkey string, params *models.CmdbRequestParams) (*models.FirewallAddrgrp6, error)
	UpdateFirewallAddrgrp6(mkey string, payload *models.FirewallAddrgrp6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallAddrgrp6(mkey string, params *models.CmdbRequestParams) error
	ListFirewallAddrgrp6(params *models.CmdbRequestParams) (*[]models.FirewallAddrgrp6, error)
	CreateFirewallPolicy(payload *models.FirewallPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallPolicy(mkey string, params *models.CmdbRequestParams) (*models.FirewallPolicy, error)
	UpdateFirewallPolicy(mkey string, payload *models.FirewallPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallPolicy(mkey string, params *models.CmdbRequestParams) error
	ListFirewallPolicy(params *models.CmdbRequestParams) (*[]models.FirewallPolicy, error)
	CreateSystemReplacemsgUtm(payload *models.SystemReplacemsgUtm, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemReplacemsgUtm(mkey string, params *models.CmdbRequestParams) (*models.SystemReplacemsgUtm, error)
	UpdateSystemReplacemsgUtm(mkey string, payload *models.SystemReplacemsgUtm, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemReplacemsgUtm(mkey string, params *models.CmdbRequestParams) error
	ListSystemReplacemsgUtm(params *models.CmdbRequestParams) (*[]models.SystemReplacemsgUtm, error)
	CreateSystemMacAddressTable(payload *models.SystemMacAddressTable, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemMacAddressTable(mkey string, params *models.CmdbRequestParams) (*models.SystemMacAddressTable, error)
	UpdateSystemMacAddressTable(mkey string, payload *models.SystemMacAddressTable, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemMacAddressTable(mkey string, params *models.CmdbRequestParams) error
	ListSystemMacAddressTable(params *models.CmdbRequestParams) (*[]models.SystemMacAddressTable, error)
	CreateSystemVneTunnel(payload *models.SystemVneTunnel, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemVneTunnel(mkey string, params *models.CmdbRequestParams) (*models.SystemVneTunnel, error)
	UpdateSystemVneTunnel(mkey string, payload *models.SystemVneTunnel, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemVneTunnel(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerQosProfile(payload *models.WirelessControllerQosProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerQosProfile(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerQosProfile, error)
	UpdateWirelessControllerQosProfile(mkey string, payload *models.WirelessControllerQosProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerQosProfile(mkey string, params *models.CmdbRequestParams) error
	ListWirelessControllerQosProfile(params *models.CmdbRequestParams) (*[]models.WirelessControllerQosProfile, error)
	CreateSwitchControllerCustomCommand(payload *models.SwitchControllerCustomCommand, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerCustomCommand(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerCustomCommand, error)
	UpdateSwitchControllerCustomCommand(mkey string, payload *models.SwitchControllerCustomCommand, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerCustomCommand(mkey string, params *models.CmdbRequestParams) error
	ListSwitchControllerCustomCommand(params *models.CmdbRequestParams) (*[]models.SwitchControllerCustomCommand, error)
	CreateIcapServer(payload *models.IcapServer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadIcapServer(mkey string, params *models.CmdbRequestParams) (*models.IcapServer, error)
	UpdateIcapServer(mkey string, payload *models.IcapServer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteIcapServer(mkey string, params *models.CmdbRequestParams) error
	ListIcapServer(params *models.CmdbRequestParams) (*[]models.IcapServer, error)
	CreateWafSubClass(payload *models.WafSubClass, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWafSubClass(mkey string, params *models.CmdbRequestParams) (*models.WafSubClass, error)
	UpdateWafSubClass(mkey string, payload *models.WafSubClass, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWafSubClass(mkey string, params *models.CmdbRequestParams) error
	ListWafSubClass(params *models.CmdbRequestParams) (*[]models.WafSubClass, error)
	CreateFirewallVipgrp64(payload *models.FirewallVipgrp64, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallVipgrp64(mkey string, params *models.CmdbRequestParams) (*models.FirewallVipgrp64, error)
	UpdateFirewallVipgrp64(mkey string, payload *models.FirewallVipgrp64, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallVipgrp64(mkey string, params *models.CmdbRequestParams) error
	ListFirewallVipgrp64(params *models.CmdbRequestParams) (*[]models.FirewallVipgrp64, error)
	CreateSystemStandaloneCluster(payload *models.SystemStandaloneCluster, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemStandaloneCluster(mkey string, params *models.CmdbRequestParams) (*models.SystemStandaloneCluster, error)
	UpdateSystemStandaloneCluster(mkey string, payload *models.SystemStandaloneCluster, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemStandaloneCluster(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallMulticastPolicy6(payload *models.FirewallMulticastPolicy6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallMulticastPolicy6(mkey string, params *models.CmdbRequestParams) (*models.FirewallMulticastPolicy6, error)
	UpdateFirewallMulticastPolicy6(mkey string, payload *models.FirewallMulticastPolicy6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallMulticastPolicy6(mkey string, params *models.CmdbRequestParams) error
	ListFirewallMulticastPolicy6(params *models.CmdbRequestParams) (*[]models.FirewallMulticastPolicy6, error)
	CreateLogDiskSetting(payload *models.LogDiskSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogDiskSetting(mkey string, params *models.CmdbRequestParams) (*models.LogDiskSetting, error)
	UpdateLogDiskSetting(mkey string, payload *models.LogDiskSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogDiskSetting(mkey string, params *models.CmdbRequestParams) error
	CreateRouterPrefixList(payload *models.RouterPrefixList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterPrefixList(mkey string, params *models.CmdbRequestParams) (*models.RouterPrefixList, error)
	UpdateRouterPrefixList(mkey string, payload *models.RouterPrefixList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterPrefixList(mkey string, params *models.CmdbRequestParams) error
	ListRouterPrefixList(params *models.CmdbRequestParams) (*[]models.RouterPrefixList, error)
	CreateSystemGeoipOverride(payload *models.SystemGeoipOverride, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemGeoipOverride(mkey string, params *models.CmdbRequestParams) (*models.SystemGeoipOverride, error)
	UpdateSystemGeoipOverride(mkey string, payload *models.SystemGeoipOverride, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemGeoipOverride(mkey string, params *models.CmdbRequestParams) error
	ListSystemGeoipOverride(params *models.CmdbRequestParams) (*[]models.SystemGeoipOverride, error)
	CreateWebProxyForwardServerGroup(payload *models.WebProxyForwardServerGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWebProxyForwardServerGroup(mkey string, params *models.CmdbRequestParams) (*models.WebProxyForwardServerGroup, error)
	UpdateWebProxyForwardServerGroup(mkey string, payload *models.WebProxyForwardServerGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWebProxyForwardServerGroup(mkey string, params *models.CmdbRequestParams) error
	ListWebProxyForwardServerGroup(params *models.CmdbRequestParams) (*[]models.WebProxyForwardServerGroup, error)
	CreateLogCustomField(payload *models.LogCustomField, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogCustomField(mkey string, params *models.CmdbRequestParams) (*models.LogCustomField, error)
	UpdateLogCustomField(mkey string, payload *models.LogCustomField, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogCustomField(mkey string, params *models.CmdbRequestParams) error
	ListLogCustomField(params *models.CmdbRequestParams) (*[]models.LogCustomField, error)
	CreateSystemVdomSflow(payload *models.SystemVdomSflow, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemVdomSflow(mkey string, params *models.CmdbRequestParams) (*models.SystemVdomSflow, error)
	UpdateSystemVdomSflow(mkey string, payload *models.SystemVdomSflow, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemVdomSflow(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallDosPolicy6(payload *models.FirewallDosPolicy6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallDosPolicy6(mkey string, params *models.CmdbRequestParams) (*models.FirewallDosPolicy6, error)
	UpdateFirewallDosPolicy6(mkey string, payload *models.FirewallDosPolicy6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallDosPolicy6(mkey string, params *models.CmdbRequestParams) error
	ListFirewallDosPolicy6(params *models.CmdbRequestParams) (*[]models.FirewallDosPolicy6, error)
	CreateWebfilterIpsUrlfilterSetting(payload *models.WebfilterIpsUrlfilterSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWebfilterIpsUrlfilterSetting(mkey string, params *models.CmdbRequestParams) (*models.WebfilterIpsUrlfilterSetting, error)
	UpdateWebfilterIpsUrlfilterSetting(mkey string, payload *models.WebfilterIpsUrlfilterSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWebfilterIpsUrlfilterSetting(mkey string, params *models.CmdbRequestParams) error
	CreateIpsRuleSettings(payload *models.IpsRuleSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadIpsRuleSettings(mkey string, params *models.CmdbRequestParams) (*models.IpsRuleSettings, error)
	UpdateIpsRuleSettings(mkey string, payload *models.IpsRuleSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteIpsRuleSettings(mkey string, params *models.CmdbRequestParams) error
	ListIpsRuleSettings(params *models.CmdbRequestParams) (*[]models.IpsRuleSettings, error)
	CreateSystemFortiguardService(payload *models.SystemFortiguardService, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemFortiguardService(mkey string, params *models.CmdbRequestParams) (*models.SystemFortiguardService, error)
	UpdateSystemFortiguardService(mkey string, payload *models.SystemFortiguardService, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemFortiguardService(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerHotspot20HsProfile(payload *models.WirelessControllerHotspot20HsProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerHotspot20HsProfile(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerHotspot20HsProfile, error)
	UpdateWirelessControllerHotspot20HsProfile(mkey string, payload *models.WirelessControllerHotspot20HsProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerHotspot20HsProfile(mkey string, params *models.CmdbRequestParams) error
	ListWirelessControllerHotspot20HsProfile(params *models.CmdbRequestParams) (*[]models.WirelessControllerHotspot20HsProfile, error)
	CreateWirelessControllerSsidPolicy(payload *models.WirelessControllerSsidPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerSsidPolicy(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerSsidPolicy, error)
	UpdateWirelessControllerSsidPolicy(mkey string, payload *models.WirelessControllerSsidPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerSsidPolicy(mkey string, params *models.CmdbRequestParams) error
	ListWirelessControllerSsidPolicy(params *models.CmdbRequestParams) (*[]models.WirelessControllerSsidPolicy, error)
	CreateSwitchControllerInitialConfigTemplate(payload *models.SwitchControllerInitialConfigTemplate, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerInitialConfigTemplate(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerInitialConfigTemplate, error)
	UpdateSwitchControllerInitialConfigTemplate(mkey string, payload *models.SwitchControllerInitialConfigTemplate, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerInitialConfigTemplate(mkey string, params *models.CmdbRequestParams) error
	ListSwitchControllerInitialConfigTemplate(params *models.CmdbRequestParams) (*[]models.SwitchControllerInitialConfigTemplate, error)
	CreateRouterAuthPath(payload *models.RouterAuthPath, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterAuthPath(mkey string, params *models.CmdbRequestParams) (*models.RouterAuthPath, error)
	UpdateRouterAuthPath(mkey string, payload *models.RouterAuthPath, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterAuthPath(mkey string, params *models.CmdbRequestParams) error
	ListRouterAuthPath(params *models.CmdbRequestParams) (*[]models.RouterAuthPath, error)
	CreateSystemReplacemsgAuth(payload *models.SystemReplacemsgAuth, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemReplacemsgAuth(mkey string, params *models.CmdbRequestParams) (*models.SystemReplacemsgAuth, error)
	UpdateSystemReplacemsgAuth(mkey string, payload *models.SystemReplacemsgAuth, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemReplacemsgAuth(mkey string, params *models.CmdbRequestParams) error
	ListSystemReplacemsgAuth(params *models.CmdbRequestParams) (*[]models.SystemReplacemsgAuth, error)
	CreateWirelessControllerWagProfile(payload *models.WirelessControllerWagProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerWagProfile(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerWagProfile, error)
	UpdateWirelessControllerWagProfile(mkey string, payload *models.WirelessControllerWagProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerWagProfile(mkey string, params *models.CmdbRequestParams) error
	ListWirelessControllerWagProfile(params *models.CmdbRequestParams) (*[]models.WirelessControllerWagProfile, error)
	CreateUserSaml(payload *models.UserSaml, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadUserSaml(mkey string, params *models.CmdbRequestParams) (*models.UserSaml, error)
	UpdateUserSaml(mkey string, payload *models.UserSaml, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteUserSaml(mkey string, params *models.CmdbRequestParams) error
	ListUserSaml(params *models.CmdbRequestParams) (*[]models.UserSaml, error)
	CreateSystemAdmin(payload *models.SystemAdmin, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemAdmin(mkey string, params *models.CmdbRequestParams) (*models.SystemAdmin, error)
	UpdateSystemAdmin(mkey string, payload *models.SystemAdmin, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemAdmin(mkey string, params *models.CmdbRequestParams) error
	ListSystemAdmin(params *models.CmdbRequestParams) (*[]models.SystemAdmin, error)
	CreateLogMemoryFilter(payload *models.LogMemoryFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogMemoryFilter(mkey string, params *models.CmdbRequestParams) (*models.LogMemoryFilter, error)
	UpdateLogMemoryFilter(mkey string, payload *models.LogMemoryFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogMemoryFilter(mkey string, params *models.CmdbRequestParams) error
	CreateDnsfilterProfile(payload *models.DnsfilterProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadDnsfilterProfile(mkey string, params *models.CmdbRequestParams) (*models.DnsfilterProfile, error)
	UpdateDnsfilterProfile(mkey string, payload *models.DnsfilterProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteDnsfilterProfile(mkey string, params *models.CmdbRequestParams) error
	ListDnsfilterProfile(params *models.CmdbRequestParams) (*[]models.DnsfilterProfile, error)
	CreateExtenderLteCarrierList(payload *models.ExtenderLteCarrierList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadExtenderLteCarrierList(mkey string, params *models.CmdbRequestParams) (*models.ExtenderLteCarrierList, error)
	UpdateExtenderLteCarrierList(mkey string, payload *models.ExtenderLteCarrierList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteExtenderLteCarrierList(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerAccessControlList(payload *models.WirelessControllerAccessControlList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerAccessControlList(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerAccessControlList, error)
	UpdateWirelessControllerAccessControlList(mkey string, payload *models.WirelessControllerAccessControlList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerAccessControlList(mkey string, params *models.CmdbRequestParams) error
	ListWirelessControllerAccessControlList(params *models.CmdbRequestParams) (*[]models.WirelessControllerAccessControlList, error)
	CreateSystemDnsDatabase(payload *models.SystemDnsDatabase, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemDnsDatabase(mkey string, params *models.CmdbRequestParams) (*models.SystemDnsDatabase, error)
	UpdateSystemDnsDatabase(mkey string, payload *models.SystemDnsDatabase, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemDnsDatabase(mkey string, params *models.CmdbRequestParams) error
	ListSystemDnsDatabase(params *models.CmdbRequestParams) (*[]models.SystemDnsDatabase, error)
	CreateSystemAlarm(payload *models.SystemAlarm, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemAlarm(mkey string, params *models.CmdbRequestParams) (*models.SystemAlarm, error)
	UpdateSystemAlarm(mkey string, payload *models.SystemAlarm, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemAlarm(mkey string, params *models.CmdbRequestParams) error
	CreateSystemSession(payload *models.SystemSession, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemSession(mkey string, params *models.CmdbRequestParams) (*models.SystemSession, error)
	UpdateSystemSession(mkey string, payload *models.SystemSession, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemSession(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallShaperTrafficShaper(payload *models.FirewallShaperTrafficShaper, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallShaperTrafficShaper(mkey string, params *models.CmdbRequestParams) (*models.FirewallShaperTrafficShaper, error)
	UpdateFirewallShaperTrafficShaper(mkey string, payload *models.FirewallShaperTrafficShaper, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallShaperTrafficShaper(mkey string, params *models.CmdbRequestParams) error
	ListFirewallShaperTrafficShaper(params *models.CmdbRequestParams) (*[]models.FirewallShaperTrafficShaper, error)
	CreateVpnCertificateCrl(payload *models.VpnCertificateCrl, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnCertificateCrl(mkey string, params *models.CmdbRequestParams) (*models.VpnCertificateCrl, error)
	UpdateVpnCertificateCrl(mkey string, payload *models.VpnCertificateCrl, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnCertificateCrl(mkey string, params *models.CmdbRequestParams) error
	ListVpnCertificateCrl(params *models.CmdbRequestParams) (*[]models.VpnCertificateCrl, error)
	CreateLogSyslogd4OverrideFilter(payload *models.LogSyslogd4OverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogSyslogd4OverrideFilter(mkey string, params *models.CmdbRequestParams) (*models.LogSyslogd4OverrideFilter, error)
	UpdateLogSyslogd4OverrideFilter(mkey string, payload *models.LogSyslogd4OverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogSyslogd4OverrideFilter(mkey string, params *models.CmdbRequestParams) error
	CreateLogFortianalyzerCloudOverrideFilter(payload *models.LogFortianalyzerCloudOverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogFortianalyzerCloudOverrideFilter(mkey string, params *models.CmdbRequestParams) (*models.LogFortianalyzerCloudOverrideFilter, error)
	UpdateLogFortianalyzerCloudOverrideFilter(mkey string, payload *models.LogFortianalyzerCloudOverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogFortianalyzerCloudOverrideFilter(mkey string, params *models.CmdbRequestParams) error
	CreateHardwareStatus(payload *models.HardwareStatus, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadHardwareStatus(mkey string, params *models.CmdbRequestParams) (*models.HardwareStatus, error)
	UpdateHardwareStatus(mkey string, payload *models.HardwareStatus, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteHardwareStatus(mkey string, params *models.CmdbRequestParams) error
	CreateSystemNtp(payload *models.SystemNtp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemNtp(mkey string, params *models.CmdbRequestParams) (*models.SystemNtp, error)
	UpdateSystemNtp(mkey string, payload *models.SystemNtp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemNtp(mkey string, params *models.CmdbRequestParams) error
	CreateSystemSaml(payload *models.SystemSaml, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemSaml(mkey string, params *models.CmdbRequestParams) (*models.SystemSaml, error)
	UpdateSystemSaml(mkey string, payload *models.SystemSaml, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemSaml(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallShaperTraffic(payload *models.FirewallShaperTraffic, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallShaperTraffic(mkey string, params *models.CmdbRequestParams) (*models.FirewallShaperTraffic, error)
	UpdateFirewallShaperTraffic(mkey string, payload *models.FirewallShaperTraffic, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallShaperTraffic(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallLocalInPolicy(payload *models.FirewallLocalInPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallLocalInPolicy(mkey string, params *models.CmdbRequestParams) (*models.FirewallLocalInPolicy, error)
	UpdateFirewallLocalInPolicy(mkey string, payload *models.FirewallLocalInPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallLocalInPolicy(mkey string, params *models.CmdbRequestParams) error
	ListFirewallLocalInPolicy(params *models.CmdbRequestParams) (*[]models.FirewallLocalInPolicy, error)
	CreateWebfilterIpsUrlfilterSetting6(payload *models.WebfilterIpsUrlfilterSetting6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWebfilterIpsUrlfilterSetting6(mkey string, params *models.CmdbRequestParams) (*models.WebfilterIpsUrlfilterSetting6, error)
	UpdateWebfilterIpsUrlfilterSetting6(mkey string, payload *models.WebfilterIpsUrlfilterSetting6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWebfilterIpsUrlfilterSetting6(mkey string, params *models.CmdbRequestParams) error
	CreateLogTacacsaccounting2Filter(payload *models.LogTacacsaccounting2Filter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogTacacsaccounting2Filter(mkey string, params *models.CmdbRequestParams) (*models.LogTacacsaccounting2Filter, error)
	UpdateLogTacacsaccounting2Filter(mkey string, payload *models.LogTacacsaccounting2Filter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogTacacsaccounting2Filter(mkey string, params *models.CmdbRequestParams) error
	CreateSystemNp6xlite(payload *models.SystemNp6xlite, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemNp6xlite(mkey string, params *models.CmdbRequestParams) (*models.SystemNp6xlite, error)
	UpdateSystemNp6xlite(mkey string, payload *models.SystemNp6xlite, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemNp6xlite(mkey string, params *models.CmdbRequestParams) error
	ListSystemNp6xlite(params *models.CmdbRequestParams) (*[]models.SystemNp6xlite, error)
	CreateSystemEmailServer(payload *models.SystemEmailServer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemEmailServer(mkey string, params *models.CmdbRequestParams) (*models.SystemEmailServer, error)
	UpdateSystemEmailServer(mkey string, payload *models.SystemEmailServer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemEmailServer(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallAccessProxy(payload *models.FirewallAccessProxy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallAccessProxy(mkey string, params *models.CmdbRequestParams) (*models.FirewallAccessProxy, error)
	UpdateFirewallAccessProxy(mkey string, payload *models.FirewallAccessProxy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallAccessProxy(mkey string, params *models.CmdbRequestParams) error
	ListFirewallAccessProxy(params *models.CmdbRequestParams) (*[]models.FirewallAccessProxy, error)
	CreateVpnL2tp(payload *models.VpnL2tp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnL2tp(mkey string, params *models.CmdbRequestParams) (*models.VpnL2tp, error)
	UpdateVpnL2tp(mkey string, payload *models.VpnL2tp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnL2tp(mkey string, params *models.CmdbRequestParams) error
	CreateWebfilterOverrideUsr(payload *models.WebfilterOverrideUsr, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWebfilterOverrideUsr(mkey string, params *models.CmdbRequestParams) (*models.WebfilterOverrideUsr, error)
	UpdateWebfilterOverrideUsr(mkey string, payload *models.WebfilterOverrideUsr, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWebfilterOverrideUsr(mkey string, params *models.CmdbRequestParams) error
	CreateIpsCustom(payload *models.IpsCustom, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadIpsCustom(mkey string, params *models.CmdbRequestParams) (*models.IpsCustom, error)
	UpdateIpsCustom(mkey string, payload *models.IpsCustom, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteIpsCustom(mkey string, params *models.CmdbRequestParams) error
	ListIpsCustom(params *models.CmdbRequestParams) (*[]models.IpsCustom, error)
	CreateDnsfilterDomainFilter(payload *models.DnsfilterDomainFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadDnsfilterDomainFilter(mkey string, params *models.CmdbRequestParams) (*models.DnsfilterDomainFilter, error)
	UpdateDnsfilterDomainFilter(mkey string, payload *models.DnsfilterDomainFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteDnsfilterDomainFilter(mkey string, params *models.CmdbRequestParams) error
	ListDnsfilterDomainFilter(params *models.CmdbRequestParams) (*[]models.DnsfilterDomainFilter, error)
	CreateAuthenticationRule(payload *models.AuthenticationRule, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadAuthenticationRule(mkey string, params *models.CmdbRequestParams) (*models.AuthenticationRule, error)
	UpdateAuthenticationRule(mkey string, payload *models.AuthenticationRule, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteAuthenticationRule(mkey string, params *models.CmdbRequestParams) error
	ListAuthenticationRule(params *models.CmdbRequestParams) (*[]models.AuthenticationRule, error)
	CreateSystemStatus(payload *models.SystemStatus, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemStatus(mkey string, params *models.CmdbRequestParams) (*models.SystemStatus, error)
	UpdateSystemStatus(mkey string, payload *models.SystemStatus, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemStatus(mkey string, params *models.CmdbRequestParams) error
	CreateSystemNpuVlink(payload *models.SystemNpuVlink, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemNpuVlink(mkey string, params *models.CmdbRequestParams) (*models.SystemNpuVlink, error)
	UpdateSystemNpuVlink(mkey string, payload *models.SystemNpuVlink, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemNpuVlink(mkey string, params *models.CmdbRequestParams) error
	ListSystemNpuVlink(params *models.CmdbRequestParams) (*[]models.SystemNpuVlink, error)
	CreateRouterRipng(payload *models.RouterRipng, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterRipng(mkey string, params *models.CmdbRequestParams) (*models.RouterRipng, error)
	UpdateRouterRipng(mkey string, payload *models.RouterRipng, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterRipng(mkey string, params *models.CmdbRequestParams) error
	CreateSystemFortiai(payload *models.SystemFortiai, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemFortiai(mkey string, params *models.CmdbRequestParams) (*models.SystemFortiai, error)
	UpdateSystemFortiai(mkey string, payload *models.SystemFortiai, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemFortiai(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerHotspot20AnqpVenueName(payload *models.WirelessControllerHotspot20AnqpVenueName, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerHotspot20AnqpVenueName(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerHotspot20AnqpVenueName, error)
	UpdateWirelessControllerHotspot20AnqpVenueName(mkey string, payload *models.WirelessControllerHotspot20AnqpVenueName, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerHotspot20AnqpVenueName(mkey string, params *models.CmdbRequestParams) error
	ListWirelessControllerHotspot20AnqpVenueName(params *models.CmdbRequestParams) (*[]models.WirelessControllerHotspot20AnqpVenueName, error)
	CreateLogFortiguardOverrideFilter(payload *models.LogFortiguardOverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogFortiguardOverrideFilter(mkey string, params *models.CmdbRequestParams) (*models.LogFortiguardOverrideFilter, error)
	UpdateLogFortiguardOverrideFilter(mkey string, payload *models.LogFortiguardOverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogFortiguardOverrideFilter(mkey string, params *models.CmdbRequestParams) error
	CreateLogFortianalyzer3OverrideFilter(payload *models.LogFortianalyzer3OverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogFortianalyzer3OverrideFilter(mkey string, params *models.CmdbRequestParams) (*models.LogFortianalyzer3OverrideFilter, error)
	UpdateLogFortianalyzer3OverrideFilter(mkey string, payload *models.LogFortianalyzer3OverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogFortianalyzer3OverrideFilter(mkey string, params *models.CmdbRequestParams) error
	CreateSystemAutoupdatePushUpdate(payload *models.SystemAutoupdatePushUpdate, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemAutoupdatePushUpdate(mkey string, params *models.CmdbRequestParams) (*models.SystemAutoupdatePushUpdate, error)
	UpdateSystemAutoupdatePushUpdate(mkey string, payload *models.SystemAutoupdatePushUpdate, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemAutoupdatePushUpdate(mkey string, params *models.CmdbRequestParams) error
	CreateSystemSdnConnector(payload *models.SystemSdnConnector, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemSdnConnector(mkey string, params *models.CmdbRequestParams) (*models.SystemSdnConnector, error)
	UpdateSystemSdnConnector(mkey string, payload *models.SystemSdnConnector, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemSdnConnector(mkey string, params *models.CmdbRequestParams) error
	ListSystemSdnConnector(params *models.CmdbRequestParams) (*[]models.SystemSdnConnector, error)
	CreateSystemVdomException(payload *models.SystemVdomException, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemVdomException(mkey string, params *models.CmdbRequestParams) (*models.SystemVdomException, error)
	UpdateSystemVdomException(mkey string, payload *models.SystemVdomException, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemVdomException(mkey string, params *models.CmdbRequestParams) error
	ListSystemVdomException(params *models.CmdbRequestParams) (*[]models.SystemVdomException, error)
	CreateFirewallScheduleOnetime(payload *models.FirewallScheduleOnetime, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallScheduleOnetime(mkey string, params *models.CmdbRequestParams) (*models.FirewallScheduleOnetime, error)
	UpdateFirewallScheduleOnetime(mkey string, payload *models.FirewallScheduleOnetime, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallScheduleOnetime(mkey string, params *models.CmdbRequestParams) error
	ListFirewallScheduleOnetime(params *models.CmdbRequestParams) (*[]models.FirewallScheduleOnetime, error)
	CreateVpnIpsecManualkeyInterface(payload *models.VpnIpsecManualkeyInterface, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnIpsecManualkeyInterface(mkey string, params *models.CmdbRequestParams) (*models.VpnIpsecManualkeyInterface, error)
	UpdateVpnIpsecManualkeyInterface(mkey string, payload *models.VpnIpsecManualkeyInterface, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnIpsecManualkeyInterface(mkey string, params *models.CmdbRequestParams) error
	ListVpnIpsecManualkeyInterface(params *models.CmdbRequestParams) (*[]models.VpnIpsecManualkeyInterface, error)
	CreateVpnIpsecTunnelSummary(payload *models.VpnIpsecTunnelSummary, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnIpsecTunnelSummary(mkey string, params *models.CmdbRequestParams) (*models.VpnIpsecTunnelSummary, error)
	UpdateVpnIpsecTunnelSummary(mkey string, payload *models.VpnIpsecTunnelSummary, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnIpsecTunnelSummary(mkey string, params *models.CmdbRequestParams) error
	CreateSystemSnmpSysinfo(payload *models.SystemSnmpSysinfo, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemSnmpSysinfo(mkey string, params *models.CmdbRequestParams) (*models.SystemSnmpSysinfo, error)
	UpdateSystemSnmpSysinfo(mkey string, payload *models.SystemSnmpSysinfo, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemSnmpSysinfo(mkey string, params *models.CmdbRequestParams) error
	CreateVpnStatusPptp(payload *models.VpnStatusPptp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnStatusPptp(mkey string, params *models.CmdbRequestParams) (*models.VpnStatusPptp, error)
	UpdateVpnStatusPptp(mkey string, payload *models.VpnStatusPptp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnStatusPptp(mkey string, params *models.CmdbRequestParams) error
	CreateDlpFpDocSource(payload *models.DlpFpDocSource, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadDlpFpDocSource(mkey string, params *models.CmdbRequestParams) (*models.DlpFpDocSource, error)
	UpdateDlpFpDocSource(mkey string, payload *models.DlpFpDocSource, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteDlpFpDocSource(mkey string, params *models.CmdbRequestParams) error
	ListDlpFpDocSource(params *models.CmdbRequestParams) (*[]models.DlpFpDocSource, error)
	CreateEmailfilterBwl(payload *models.EmailfilterBwl, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadEmailfilterBwl(mkey string, params *models.CmdbRequestParams) (*models.EmailfilterBwl, error)
	UpdateEmailfilterBwl(mkey string, payload *models.EmailfilterBwl, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteEmailfilterBwl(mkey string, params *models.CmdbRequestParams) error
	ListEmailfilterBwl(params *models.CmdbRequestParams) (*[]models.EmailfilterBwl, error)
	CreateFirewallRegion(payload *models.FirewallRegion, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallRegion(mkey string, params *models.CmdbRequestParams) (*models.FirewallRegion, error)
	UpdateFirewallRegion(mkey string, payload *models.FirewallRegion, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallRegion(mkey string, params *models.CmdbRequestParams) error
	ListFirewallRegion(params *models.CmdbRequestParams) (*[]models.FirewallRegion, error)
	CreateSystemReplacemsgAdmin(payload *models.SystemReplacemsgAdmin, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemReplacemsgAdmin(mkey string, params *models.CmdbRequestParams) (*models.SystemReplacemsgAdmin, error)
	UpdateSystemReplacemsgAdmin(mkey string, payload *models.SystemReplacemsgAdmin, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemReplacemsgAdmin(mkey string, params *models.CmdbRequestParams) error
	ListSystemReplacemsgAdmin(params *models.CmdbRequestParams) (*[]models.SystemReplacemsgAdmin, error)
	CreateLogSyslogdSetting(payload *models.LogSyslogdSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogSyslogdSetting(mkey string, params *models.CmdbRequestParams) (*models.LogSyslogdSetting, error)
	UpdateLogSyslogdSetting(mkey string, payload *models.LogSyslogdSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogSyslogdSetting(mkey string, params *models.CmdbRequestParams) error
	CreateUserLdap(payload *models.UserLdap, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadUserLdap(mkey string, params *models.CmdbRequestParams) (*models.UserLdap, error)
	UpdateUserLdap(mkey string, payload *models.UserLdap, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteUserLdap(mkey string, params *models.CmdbRequestParams) error
	ListUserLdap(params *models.CmdbRequestParams) (*[]models.UserLdap, error)
	CreateUserDomainController(payload *models.UserDomainController, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadUserDomainController(mkey string, params *models.CmdbRequestParams) (*models.UserDomainController, error)
	UpdateUserDomainController(mkey string, payload *models.UserDomainController, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteUserDomainController(mkey string, params *models.CmdbRequestParams) error
	ListUserDomainController(params *models.CmdbRequestParams) (*[]models.UserDomainController, error)
	CreateSystemStorage(payload *models.SystemStorage, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemStorage(mkey string, params *models.CmdbRequestParams) (*models.SystemStorage, error)
	UpdateSystemStorage(mkey string, payload *models.SystemStorage, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemStorage(mkey string, params *models.CmdbRequestParams) error
	ListSystemStorage(params *models.CmdbRequestParams) (*[]models.SystemStorage, error)
	CreateWirelessControllerClientInfo(payload *models.WirelessControllerClientInfo, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerClientInfo(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerClientInfo, error)
	UpdateWirelessControllerClientInfo(mkey string, payload *models.WirelessControllerClientInfo, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerClientInfo(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallVendorMacSummary(payload *models.FirewallVendorMacSummary, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallVendorMacSummary(mkey string, params *models.CmdbRequestParams) (*models.FirewallVendorMacSummary, error)
	UpdateFirewallVendorMacSummary(mkey string, payload *models.FirewallVendorMacSummary, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallVendorMacSummary(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallShaperPerIp(payload *models.FirewallShaperPerIp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallShaperPerIp(mkey string, params *models.CmdbRequestParams) (*models.FirewallShaperPerIp, error)
	UpdateFirewallShaperPerIp(mkey string, payload *models.FirewallShaperPerIp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallShaperPerIp(mkey string, params *models.CmdbRequestParams) error
	CreateLogSyslogd2OverrideSetting(payload *models.LogSyslogd2OverrideSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogSyslogd2OverrideSetting(mkey string, params *models.CmdbRequestParams) (*models.LogSyslogd2OverrideSetting, error)
	UpdateLogSyslogd2OverrideSetting(mkey string, payload *models.LogSyslogd2OverrideSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogSyslogd2OverrideSetting(mkey string, params *models.CmdbRequestParams) error
	CreateUserPop3(payload *models.UserPop3, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadUserPop3(mkey string, params *models.CmdbRequestParams) (*models.UserPop3, error)
	UpdateUserPop3(mkey string, payload *models.UserPop3, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteUserPop3(mkey string, params *models.CmdbRequestParams) error
	ListUserPop3(params *models.CmdbRequestParams) (*[]models.UserPop3, error)
	CreateExtenderControllerDataplan(payload *models.ExtenderControllerDataplan, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadExtenderControllerDataplan(mkey string, params *models.CmdbRequestParams) (*models.ExtenderControllerDataplan, error)
	UpdateExtenderControllerDataplan(mkey string, payload *models.ExtenderControllerDataplan, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteExtenderControllerDataplan(mkey string, params *models.CmdbRequestParams) error
	ListExtenderControllerDataplan(params *models.CmdbRequestParams) (*[]models.ExtenderControllerDataplan, error)
	CreateSystemVirtualSwitch(payload *models.SystemVirtualSwitch, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemVirtualSwitch(mkey string, params *models.CmdbRequestParams) (*models.SystemVirtualSwitch, error)
	UpdateSystemVirtualSwitch(mkey string, payload *models.SystemVirtualSwitch, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemVirtualSwitch(mkey string, params *models.CmdbRequestParams) error
	ListSystemVirtualSwitch(params *models.CmdbRequestParams) (*[]models.SystemVirtualSwitch, error)
	CreateSystemFtmPush(payload *models.SystemFtmPush, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemFtmPush(mkey string, params *models.CmdbRequestParams) (*models.SystemFtmPush, error)
	UpdateSystemFtmPush(mkey string, payload *models.SystemFtmPush, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemFtmPush(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerHotspot20H2qpAdviceOfCharge(payload *models.WirelessControllerHotspot20H2qpAdviceOfCharge, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerHotspot20H2qpAdviceOfCharge(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerHotspot20H2qpAdviceOfCharge, error)
	UpdateWirelessControllerHotspot20H2qpAdviceOfCharge(mkey string, payload *models.WirelessControllerHotspot20H2qpAdviceOfCharge, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerHotspot20H2qpAdviceOfCharge(mkey string, params *models.CmdbRequestParams) error
	ListWirelessControllerHotspot20H2qpAdviceOfCharge(params *models.CmdbRequestParams) (*[]models.WirelessControllerHotspot20H2qpAdviceOfCharge, error)
	CreateSwitchControllerIgmpSnooping(payload *models.SwitchControllerIgmpSnooping, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerIgmpSnooping(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerIgmpSnooping, error)
	UpdateSwitchControllerIgmpSnooping(mkey string, payload *models.SwitchControllerIgmpSnooping, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerIgmpSnooping(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallInternetServiceAddition(payload *models.FirewallInternetServiceAddition, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallInternetServiceAddition(mkey string, params *models.CmdbRequestParams) (*models.FirewallInternetServiceAddition, error)
	UpdateFirewallInternetServiceAddition(mkey string, payload *models.FirewallInternetServiceAddition, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallInternetServiceAddition(mkey string, params *models.CmdbRequestParams) error
	ListFirewallInternetServiceAddition(params *models.CmdbRequestParams) (*[]models.FirewallInternetServiceAddition, error)
	CreateLogNpuServer(payload *models.LogNpuServer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogNpuServer(mkey string, params *models.CmdbRequestParams) (*models.LogNpuServer, error)
	UpdateLogNpuServer(mkey string, payload *models.LogNpuServer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogNpuServer(mkey string, params *models.CmdbRequestParams) error
	CreateSystemSessionHelper(payload *models.SystemSessionHelper, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemSessionHelper(mkey string, params *models.CmdbRequestParams) (*models.SystemSessionHelper, error)
	UpdateSystemSessionHelper(mkey string, payload *models.SystemSessionHelper, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemSessionHelper(mkey string, params *models.CmdbRequestParams) error
	ListSystemSessionHelper(params *models.CmdbRequestParams) (*[]models.SystemSessionHelper, error)
	CreateSwitchControllerFlowTracking(payload *models.SwitchControllerFlowTracking, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerFlowTracking(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerFlowTracking, error)
	UpdateSwitchControllerFlowTracking(mkey string, payload *models.SwitchControllerFlowTracking, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerFlowTracking(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallInternetServiceDefinition(payload *models.FirewallInternetServiceDefinition, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallInternetServiceDefinition(mkey string, params *models.CmdbRequestParams) (*models.FirewallInternetServiceDefinition, error)
	UpdateFirewallInternetServiceDefinition(mkey string, payload *models.FirewallInternetServiceDefinition, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallInternetServiceDefinition(mkey string, params *models.CmdbRequestParams) error
	ListFirewallInternetServiceDefinition(params *models.CmdbRequestParams) (*[]models.FirewallInternetServiceDefinition, error)
	CreateVideofilterYoutubeChannelFilter(payload *models.VideofilterYoutubeChannelFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVideofilterYoutubeChannelFilter(mkey string, params *models.CmdbRequestParams) (*models.VideofilterYoutubeChannelFilter, error)
	UpdateVideofilterYoutubeChannelFilter(mkey string, payload *models.VideofilterYoutubeChannelFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVideofilterYoutubeChannelFilter(mkey string, params *models.CmdbRequestParams) error
	ListVideofilterYoutubeChannelFilter(params *models.CmdbRequestParams) (*[]models.VideofilterYoutubeChannelFilter, error)
	CreateExtenderControllerExtenderProfile(payload *models.ExtenderControllerExtenderProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadExtenderControllerExtenderProfile(mkey string, params *models.CmdbRequestParams) (*models.ExtenderControllerExtenderProfile, error)
	UpdateExtenderControllerExtenderProfile(mkey string, payload *models.ExtenderControllerExtenderProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteExtenderControllerExtenderProfile(mkey string, params *models.CmdbRequestParams) error
	ListExtenderControllerExtenderProfile(params *models.CmdbRequestParams) (*[]models.ExtenderControllerExtenderProfile, error)
	CreateFirewallSslServer(payload *models.FirewallSslServer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallSslServer(mkey string, params *models.CmdbRequestParams) (*models.FirewallSslServer, error)
	UpdateFirewallSslServer(mkey string, payload *models.FirewallSslServer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallSslServer(mkey string, params *models.CmdbRequestParams) error
	ListFirewallSslServer(params *models.CmdbRequestParams) (*[]models.FirewallSslServer, error)
}
