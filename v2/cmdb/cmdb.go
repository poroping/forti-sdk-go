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
	CreateAlertemailSetting(payload *models.AlertemailSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadAlertemailSetting(mkey string, params *models.CmdbRequestParams) (*models.AlertemailSetting, error)
	UpdateAlertemailSetting(mkey string, payload *models.AlertemailSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteAlertemailSetting(mkey string, params *models.CmdbRequestParams) error
	CreateAntivirusHeuristic(payload *models.AntivirusHeuristic, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadAntivirusHeuristic(mkey string, params *models.CmdbRequestParams) (*models.AntivirusHeuristic, error)
	UpdateAntivirusHeuristic(mkey string, payload *models.AntivirusHeuristic, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteAntivirusHeuristic(mkey string, params *models.CmdbRequestParams) error
	CreateAntivirusProfile(payload *models.AntivirusProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadAntivirusProfile(mkey string, params *models.CmdbRequestParams) (*models.AntivirusProfile, error)
	UpdateAntivirusProfile(mkey string, payload *models.AntivirusProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteAntivirusProfile(mkey string, params *models.CmdbRequestParams) error
	CreateAntivirusQuarantine(payload *models.AntivirusQuarantine, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadAntivirusQuarantine(mkey string, params *models.CmdbRequestParams) (*models.AntivirusQuarantine, error)
	UpdateAntivirusQuarantine(mkey string, payload *models.AntivirusQuarantine, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteAntivirusQuarantine(mkey string, params *models.CmdbRequestParams) error
	CreateAntivirusSettings(payload *models.AntivirusSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadAntivirusSettings(mkey string, params *models.CmdbRequestParams) (*models.AntivirusSettings, error)
	UpdateAntivirusSettings(mkey string, payload *models.AntivirusSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteAntivirusSettings(mkey string, params *models.CmdbRequestParams) error
	CreateApplicationCustom(payload *models.ApplicationCustom, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadApplicationCustom(mkey string, params *models.CmdbRequestParams) (*models.ApplicationCustom, error)
	UpdateApplicationCustom(mkey string, payload *models.ApplicationCustom, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteApplicationCustom(mkey string, params *models.CmdbRequestParams) error
	CreateApplicationGroup(payload *models.ApplicationGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadApplicationGroup(mkey string, params *models.CmdbRequestParams) (*models.ApplicationGroup, error)
	UpdateApplicationGroup(mkey string, payload *models.ApplicationGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteApplicationGroup(mkey string, params *models.CmdbRequestParams) error
	CreateApplicationList(payload *models.ApplicationList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadApplicationList(mkey string, params *models.CmdbRequestParams) (*models.ApplicationList, error)
	UpdateApplicationList(mkey string, payload *models.ApplicationList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteApplicationList(mkey string, params *models.CmdbRequestParams) error
	CreateApplicationName(payload *models.ApplicationName, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadApplicationName(mkey string, params *models.CmdbRequestParams) (*models.ApplicationName, error)
	UpdateApplicationName(mkey string, payload *models.ApplicationName, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteApplicationName(mkey string, params *models.CmdbRequestParams) error
	CreateApplicationRuleSettings(payload *models.ApplicationRuleSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadApplicationRuleSettings(mkey string, params *models.CmdbRequestParams) (*models.ApplicationRuleSettings, error)
	UpdateApplicationRuleSettings(mkey string, payload *models.ApplicationRuleSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteApplicationRuleSettings(mkey string, params *models.CmdbRequestParams) error
	CreateAuthenticationRule(payload *models.AuthenticationRule, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadAuthenticationRule(mkey string, params *models.CmdbRequestParams) (*models.AuthenticationRule, error)
	UpdateAuthenticationRule(mkey string, payload *models.AuthenticationRule, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteAuthenticationRule(mkey string, params *models.CmdbRequestParams) error
	CreateAuthenticationScheme(payload *models.AuthenticationScheme, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadAuthenticationScheme(mkey string, params *models.CmdbRequestParams) (*models.AuthenticationScheme, error)
	UpdateAuthenticationScheme(mkey string, payload *models.AuthenticationScheme, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteAuthenticationScheme(mkey string, params *models.CmdbRequestParams) error
	CreateAuthenticationSetting(payload *models.AuthenticationSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadAuthenticationSetting(mkey string, params *models.CmdbRequestParams) (*models.AuthenticationSetting, error)
	UpdateAuthenticationSetting(mkey string, payload *models.AuthenticationSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteAuthenticationSetting(mkey string, params *models.CmdbRequestParams) error
	CreateCertificateCa(payload *models.CertificateCa, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadCertificateCa(mkey string, params *models.CmdbRequestParams) (*models.CertificateCa, error)
	UpdateCertificateCa(mkey string, payload *models.CertificateCa, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteCertificateCa(mkey string, params *models.CmdbRequestParams) error
	CreateCertificateCrl(payload *models.CertificateCrl, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadCertificateCrl(mkey string, params *models.CmdbRequestParams) (*models.CertificateCrl, error)
	UpdateCertificateCrl(mkey string, payload *models.CertificateCrl, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteCertificateCrl(mkey string, params *models.CmdbRequestParams) error
	CreateCertificateLocal(payload *models.CertificateLocal, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadCertificateLocal(mkey string, params *models.CmdbRequestParams) (*models.CertificateLocal, error)
	UpdateCertificateLocal(mkey string, payload *models.CertificateLocal, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteCertificateLocal(mkey string, params *models.CmdbRequestParams) error
	CreateCertificateRemote(payload *models.CertificateRemote, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadCertificateRemote(mkey string, params *models.CmdbRequestParams) (*models.CertificateRemote, error)
	UpdateCertificateRemote(mkey string, payload *models.CertificateRemote, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteCertificateRemote(mkey string, params *models.CmdbRequestParams) error
	CreateCifsDomainController(payload *models.CifsDomainController, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadCifsDomainController(mkey string, params *models.CmdbRequestParams) (*models.CifsDomainController, error)
	UpdateCifsDomainController(mkey string, payload *models.CifsDomainController, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteCifsDomainController(mkey string, params *models.CmdbRequestParams) error
	CreateCifsProfile(payload *models.CifsProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadCifsProfile(mkey string, params *models.CmdbRequestParams) (*models.CifsProfile, error)
	UpdateCifsProfile(mkey string, payload *models.CifsProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteCifsProfile(mkey string, params *models.CmdbRequestParams) error
	CreateCredentialStoreDomainController(payload *models.CredentialStoreDomainController, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadCredentialStoreDomainController(mkey string, params *models.CmdbRequestParams) (*models.CredentialStoreDomainController, error)
	UpdateCredentialStoreDomainController(mkey string, payload *models.CredentialStoreDomainController, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteCredentialStoreDomainController(mkey string, params *models.CmdbRequestParams) error
	CreateDlpFilepattern(payload *models.DlpFilepattern, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadDlpFilepattern(mkey string, params *models.CmdbRequestParams) (*models.DlpFilepattern, error)
	UpdateDlpFilepattern(mkey string, payload *models.DlpFilepattern, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteDlpFilepattern(mkey string, params *models.CmdbRequestParams) error
	CreateDlpFpDocSource(payload *models.DlpFpDocSource, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadDlpFpDocSource(mkey string, params *models.CmdbRequestParams) (*models.DlpFpDocSource, error)
	UpdateDlpFpDocSource(mkey string, payload *models.DlpFpDocSource, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteDlpFpDocSource(mkey string, params *models.CmdbRequestParams) error
	CreateDlpSensitivity(payload *models.DlpSensitivity, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadDlpSensitivity(mkey string, params *models.CmdbRequestParams) (*models.DlpSensitivity, error)
	UpdateDlpSensitivity(mkey string, payload *models.DlpSensitivity, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteDlpSensitivity(mkey string, params *models.CmdbRequestParams) error
	CreateDlpSensor(payload *models.DlpSensor, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadDlpSensor(mkey string, params *models.CmdbRequestParams) (*models.DlpSensor, error)
	UpdateDlpSensor(mkey string, payload *models.DlpSensor, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteDlpSensor(mkey string, params *models.CmdbRequestParams) error
	CreateDlpSettings(payload *models.DlpSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadDlpSettings(mkey string, params *models.CmdbRequestParams) (*models.DlpSettings, error)
	UpdateDlpSettings(mkey string, payload *models.DlpSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteDlpSettings(mkey string, params *models.CmdbRequestParams) error
	CreateDnsfilterDomainFilter(payload *models.DnsfilterDomainFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadDnsfilterDomainFilter(mkey string, params *models.CmdbRequestParams) (*models.DnsfilterDomainFilter, error)
	UpdateDnsfilterDomainFilter(mkey string, payload *models.DnsfilterDomainFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteDnsfilterDomainFilter(mkey string, params *models.CmdbRequestParams) error
	CreateDnsfilterProfile(payload *models.DnsfilterProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadDnsfilterProfile(mkey string, params *models.CmdbRequestParams) (*models.DnsfilterProfile, error)
	UpdateDnsfilterProfile(mkey string, payload *models.DnsfilterProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteDnsfilterProfile(mkey string, params *models.CmdbRequestParams) error
	CreateEmailfilterBlockAllowList(payload *models.EmailfilterBlockAllowList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadEmailfilterBlockAllowList(mkey string, params *models.CmdbRequestParams) (*models.EmailfilterBlockAllowList, error)
	UpdateEmailfilterBlockAllowList(mkey string, payload *models.EmailfilterBlockAllowList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteEmailfilterBlockAllowList(mkey string, params *models.CmdbRequestParams) error
	CreateEmailfilterBwl(payload *models.EmailfilterBwl, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadEmailfilterBwl(mkey string, params *models.CmdbRequestParams) (*models.EmailfilterBwl, error)
	UpdateEmailfilterBwl(mkey string, payload *models.EmailfilterBwl, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteEmailfilterBwl(mkey string, params *models.CmdbRequestParams) error
	CreateEmailfilterBword(payload *models.EmailfilterBword, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadEmailfilterBword(mkey string, params *models.CmdbRequestParams) (*models.EmailfilterBword, error)
	UpdateEmailfilterBword(mkey string, payload *models.EmailfilterBword, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteEmailfilterBword(mkey string, params *models.CmdbRequestParams) error
	CreateEmailfilterDnsbl(payload *models.EmailfilterDnsbl, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadEmailfilterDnsbl(mkey string, params *models.CmdbRequestParams) (*models.EmailfilterDnsbl, error)
	UpdateEmailfilterDnsbl(mkey string, payload *models.EmailfilterDnsbl, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteEmailfilterDnsbl(mkey string, params *models.CmdbRequestParams) error
	CreateEmailfilterFortishield(payload *models.EmailfilterFortishield, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadEmailfilterFortishield(mkey string, params *models.CmdbRequestParams) (*models.EmailfilterFortishield, error)
	UpdateEmailfilterFortishield(mkey string, payload *models.EmailfilterFortishield, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteEmailfilterFortishield(mkey string, params *models.CmdbRequestParams) error
	CreateEmailfilterIptrust(payload *models.EmailfilterIptrust, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadEmailfilterIptrust(mkey string, params *models.CmdbRequestParams) (*models.EmailfilterIptrust, error)
	UpdateEmailfilterIptrust(mkey string, payload *models.EmailfilterIptrust, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteEmailfilterIptrust(mkey string, params *models.CmdbRequestParams) error
	CreateEmailfilterMheader(payload *models.EmailfilterMheader, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadEmailfilterMheader(mkey string, params *models.CmdbRequestParams) (*models.EmailfilterMheader, error)
	UpdateEmailfilterMheader(mkey string, payload *models.EmailfilterMheader, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteEmailfilterMheader(mkey string, params *models.CmdbRequestParams) error
	CreateEmailfilterOptions(payload *models.EmailfilterOptions, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadEmailfilterOptions(mkey string, params *models.CmdbRequestParams) (*models.EmailfilterOptions, error)
	UpdateEmailfilterOptions(mkey string, payload *models.EmailfilterOptions, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteEmailfilterOptions(mkey string, params *models.CmdbRequestParams) error
	CreateEmailfilterProfile(payload *models.EmailfilterProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadEmailfilterProfile(mkey string, params *models.CmdbRequestParams) (*models.EmailfilterProfile, error)
	UpdateEmailfilterProfile(mkey string, payload *models.EmailfilterProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteEmailfilterProfile(mkey string, params *models.CmdbRequestParams) error
	CreateEndpointControlFctems(payload *models.EndpointControlFctems, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadEndpointControlFctems(mkey string, params *models.CmdbRequestParams) (*models.EndpointControlFctems, error)
	UpdateEndpointControlFctems(mkey string, payload *models.EndpointControlFctems, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteEndpointControlFctems(mkey string, params *models.CmdbRequestParams) error
	CreateEndpointControlSettings(payload *models.EndpointControlSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadEndpointControlSettings(mkey string, params *models.CmdbRequestParams) (*models.EndpointControlSettings, error)
	UpdateEndpointControlSettings(mkey string, payload *models.EndpointControlSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteEndpointControlSettings(mkey string, params *models.CmdbRequestParams) error
	CreateExtenderControllerDataplan(payload *models.ExtenderControllerDataplan, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadExtenderControllerDataplan(mkey string, params *models.CmdbRequestParams) (*models.ExtenderControllerDataplan, error)
	UpdateExtenderControllerDataplan(mkey string, payload *models.ExtenderControllerDataplan, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteExtenderControllerDataplan(mkey string, params *models.CmdbRequestParams) error
	CreateExtenderControllerExtender(payload *models.ExtenderControllerExtender, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadExtenderControllerExtender(mkey string, params *models.CmdbRequestParams) (*models.ExtenderControllerExtender, error)
	UpdateExtenderControllerExtender(mkey string, payload *models.ExtenderControllerExtender, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteExtenderControllerExtender(mkey string, params *models.CmdbRequestParams) error
	CreateExtenderControllerExtenderProfile(payload *models.ExtenderControllerExtenderProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadExtenderControllerExtenderProfile(mkey string, params *models.CmdbRequestParams) (*models.ExtenderControllerExtenderProfile, error)
	UpdateExtenderControllerExtenderProfile(mkey string, payload *models.ExtenderControllerExtenderProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteExtenderControllerExtenderProfile(mkey string, params *models.CmdbRequestParams) error
	CreateFileFilterProfile(payload *models.FileFilterProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFileFilterProfile(mkey string, params *models.CmdbRequestParams) (*models.FileFilterProfile, error)
	UpdateFileFilterProfile(mkey string, payload *models.FileFilterProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFileFilterProfile(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallconsolidatedPolicy(payload *models.FirewallconsolidatedPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallconsolidatedPolicy(mkey string, params *models.CmdbRequestParams) (*models.FirewallconsolidatedPolicy, error)
	UpdateFirewallconsolidatedPolicy(mkey string, payload *models.FirewallconsolidatedPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallconsolidatedPolicy(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallipmacbindingSetting(payload *models.FirewallipmacbindingSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallipmacbindingSetting(mkey string, params *models.CmdbRequestParams) (*models.FirewallipmacbindingSetting, error)
	UpdateFirewallipmacbindingSetting(mkey string, payload *models.FirewallipmacbindingSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallipmacbindingSetting(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallipmacbindingTable(payload *models.FirewallipmacbindingTable, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallipmacbindingTable(mkey string, params *models.CmdbRequestParams) (*models.FirewallipmacbindingTable, error)
	UpdateFirewallipmacbindingTable(mkey string, payload *models.FirewallipmacbindingTable, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallipmacbindingTable(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallscheduleGroup(payload *models.FirewallscheduleGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallscheduleGroup(mkey string, params *models.CmdbRequestParams) (*models.FirewallscheduleGroup, error)
	UpdateFirewallscheduleGroup(mkey string, payload *models.FirewallscheduleGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallscheduleGroup(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallscheduleOnetime(payload *models.FirewallscheduleOnetime, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallscheduleOnetime(mkey string, params *models.CmdbRequestParams) (*models.FirewallscheduleOnetime, error)
	UpdateFirewallscheduleOnetime(mkey string, payload *models.FirewallscheduleOnetime, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallscheduleOnetime(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallscheduleRecurring(payload *models.FirewallscheduleRecurring, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallscheduleRecurring(mkey string, params *models.CmdbRequestParams) (*models.FirewallscheduleRecurring, error)
	UpdateFirewallscheduleRecurring(mkey string, payload *models.FirewallscheduleRecurring, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallscheduleRecurring(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallserviceCategory(payload *models.FirewallserviceCategory, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallserviceCategory(mkey string, params *models.CmdbRequestParams) (*models.FirewallserviceCategory, error)
	UpdateFirewallserviceCategory(mkey string, payload *models.FirewallserviceCategory, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallserviceCategory(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallserviceCustom(payload *models.FirewallserviceCustom, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallserviceCustom(mkey string, params *models.CmdbRequestParams) (*models.FirewallserviceCustom, error)
	UpdateFirewallserviceCustom(mkey string, payload *models.FirewallserviceCustom, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallserviceCustom(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallserviceGroup(payload *models.FirewallserviceGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallserviceGroup(mkey string, params *models.CmdbRequestParams) (*models.FirewallserviceGroup, error)
	UpdateFirewallserviceGroup(mkey string, payload *models.FirewallserviceGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallserviceGroup(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallshaperPerIpShaper(payload *models.FirewallshaperPerIpShaper, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallshaperPerIpShaper(mkey string, params *models.CmdbRequestParams) (*models.FirewallshaperPerIpShaper, error)
	UpdateFirewallshaperPerIpShaper(mkey string, payload *models.FirewallshaperPerIpShaper, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallshaperPerIpShaper(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallshaperTrafficShaper(payload *models.FirewallshaperTrafficShaper, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallshaperTrafficShaper(mkey string, params *models.CmdbRequestParams) (*models.FirewallshaperTrafficShaper, error)
	UpdateFirewallshaperTrafficShaper(mkey string, payload *models.FirewallshaperTrafficShaper, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallshaperTrafficShaper(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallsshHostKey(payload *models.FirewallsshHostKey, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallsshHostKey(mkey string, params *models.CmdbRequestParams) (*models.FirewallsshHostKey, error)
	UpdateFirewallsshHostKey(mkey string, payload *models.FirewallsshHostKey, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallsshHostKey(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallsshLocalCa(payload *models.FirewallsshLocalCa, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallsshLocalCa(mkey string, params *models.CmdbRequestParams) (*models.FirewallsshLocalCa, error)
	UpdateFirewallsshLocalCa(mkey string, payload *models.FirewallsshLocalCa, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallsshLocalCa(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallsshLocalKey(payload *models.FirewallsshLocalKey, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallsshLocalKey(mkey string, params *models.CmdbRequestParams) (*models.FirewallsshLocalKey, error)
	UpdateFirewallsshLocalKey(mkey string, payload *models.FirewallsshLocalKey, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallsshLocalKey(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallsshSetting(payload *models.FirewallsshSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallsshSetting(mkey string, params *models.CmdbRequestParams) (*models.FirewallsshSetting, error)
	UpdateFirewallsshSetting(mkey string, payload *models.FirewallsshSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallsshSetting(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallsslSetting(payload *models.FirewallsslSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallsslSetting(mkey string, params *models.CmdbRequestParams) (*models.FirewallsslSetting, error)
	UpdateFirewallsslSetting(mkey string, payload *models.FirewallsslSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallsslSetting(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallwildcardFqdnCustom(payload *models.FirewallwildcardFqdnCustom, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallwildcardFqdnCustom(mkey string, params *models.CmdbRequestParams) (*models.FirewallwildcardFqdnCustom, error)
	UpdateFirewallwildcardFqdnCustom(mkey string, payload *models.FirewallwildcardFqdnCustom, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallwildcardFqdnCustom(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallwildcardFqdnGroup(payload *models.FirewallwildcardFqdnGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallwildcardFqdnGroup(mkey string, params *models.CmdbRequestParams) (*models.FirewallwildcardFqdnGroup, error)
	UpdateFirewallwildcardFqdnGroup(mkey string, payload *models.FirewallwildcardFqdnGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallwildcardFqdnGroup(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallDoSPolicy(payload *models.FirewallDoSPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallDoSPolicy(mkey string, params *models.CmdbRequestParams) (*models.FirewallDoSPolicy, error)
	UpdateFirewallDoSPolicy(mkey string, payload *models.FirewallDoSPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallDoSPolicy(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallDoSPolicy6(payload *models.FirewallDoSPolicy6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallDoSPolicy6(mkey string, params *models.CmdbRequestParams) (*models.FirewallDoSPolicy6, error)
	UpdateFirewallDoSPolicy6(mkey string, payload *models.FirewallDoSPolicy6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallDoSPolicy6(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallAccessProxy(payload *models.FirewallAccessProxy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallAccessProxy(mkey string, params *models.CmdbRequestParams) (*models.FirewallAccessProxy, error)
	UpdateFirewallAccessProxy(mkey string, payload *models.FirewallAccessProxy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallAccessProxy(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallAccessProxy6(payload *models.FirewallAccessProxy6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallAccessProxy6(mkey string, params *models.CmdbRequestParams) (*models.FirewallAccessProxy6, error)
	UpdateFirewallAccessProxy6(mkey string, payload *models.FirewallAccessProxy6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallAccessProxy6(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallAccessProxySshClientCert(payload *models.FirewallAccessProxySshClientCert, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallAccessProxySshClientCert(mkey string, params *models.CmdbRequestParams) (*models.FirewallAccessProxySshClientCert, error)
	UpdateFirewallAccessProxySshClientCert(mkey string, payload *models.FirewallAccessProxySshClientCert, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallAccessProxySshClientCert(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallAccessProxyVirtualHost(payload *models.FirewallAccessProxyVirtualHost, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallAccessProxyVirtualHost(mkey string, params *models.CmdbRequestParams) (*models.FirewallAccessProxyVirtualHost, error)
	UpdateFirewallAccessProxyVirtualHost(mkey string, payload *models.FirewallAccessProxyVirtualHost, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallAccessProxyVirtualHost(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallAddress(payload *models.FirewallAddress, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallAddress(mkey string, params *models.CmdbRequestParams) (*models.FirewallAddress, error)
	UpdateFirewallAddress(mkey string, payload *models.FirewallAddress, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallAddress(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallAddress6(payload *models.FirewallAddress6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallAddress6(mkey string, params *models.CmdbRequestParams) (*models.FirewallAddress6, error)
	UpdateFirewallAddress6(mkey string, payload *models.FirewallAddress6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallAddress6(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallAddress6Template(payload *models.FirewallAddress6Template, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallAddress6Template(mkey string, params *models.CmdbRequestParams) (*models.FirewallAddress6Template, error)
	UpdateFirewallAddress6Template(mkey string, payload *models.FirewallAddress6Template, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallAddress6Template(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallAddrgrp(payload *models.FirewallAddrgrp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallAddrgrp(mkey string, params *models.CmdbRequestParams) (*models.FirewallAddrgrp, error)
	UpdateFirewallAddrgrp(mkey string, payload *models.FirewallAddrgrp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallAddrgrp(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallAddrgrp6(payload *models.FirewallAddrgrp6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallAddrgrp6(mkey string, params *models.CmdbRequestParams) (*models.FirewallAddrgrp6, error)
	UpdateFirewallAddrgrp6(mkey string, payload *models.FirewallAddrgrp6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallAddrgrp6(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallAuthPortal(payload *models.FirewallAuthPortal, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallAuthPortal(mkey string, params *models.CmdbRequestParams) (*models.FirewallAuthPortal, error)
	UpdateFirewallAuthPortal(mkey string, payload *models.FirewallAuthPortal, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallAuthPortal(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallCentralSnatMap(payload *models.FirewallCentralSnatMap, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallCentralSnatMap(mkey string, params *models.CmdbRequestParams) (*models.FirewallCentralSnatMap, error)
	UpdateFirewallCentralSnatMap(mkey string, payload *models.FirewallCentralSnatMap, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallCentralSnatMap(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallCity(payload *models.FirewallCity, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallCity(mkey string, params *models.CmdbRequestParams) (*models.FirewallCity, error)
	UpdateFirewallCity(mkey string, payload *models.FirewallCity, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallCity(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallCountry(payload *models.FirewallCountry, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallCountry(mkey string, params *models.CmdbRequestParams) (*models.FirewallCountry, error)
	UpdateFirewallCountry(mkey string, payload *models.FirewallCountry, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallCountry(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallDecryptedTrafficMirror(payload *models.FirewallDecryptedTrafficMirror, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallDecryptedTrafficMirror(mkey string, params *models.CmdbRequestParams) (*models.FirewallDecryptedTrafficMirror, error)
	UpdateFirewallDecryptedTrafficMirror(mkey string, payload *models.FirewallDecryptedTrafficMirror, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallDecryptedTrafficMirror(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallDnstranslation(payload *models.FirewallDnstranslation, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallDnstranslation(mkey string, params *models.CmdbRequestParams) (*models.FirewallDnstranslation, error)
	UpdateFirewallDnstranslation(mkey string, payload *models.FirewallDnstranslation, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallDnstranslation(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallIdentityBasedRoute(payload *models.FirewallIdentityBasedRoute, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallIdentityBasedRoute(mkey string, params *models.CmdbRequestParams) (*models.FirewallIdentityBasedRoute, error)
	UpdateFirewallIdentityBasedRoute(mkey string, payload *models.FirewallIdentityBasedRoute, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallIdentityBasedRoute(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallInterfacePolicy(payload *models.FirewallInterfacePolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallInterfacePolicy(mkey string, params *models.CmdbRequestParams) (*models.FirewallInterfacePolicy, error)
	UpdateFirewallInterfacePolicy(mkey string, payload *models.FirewallInterfacePolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallInterfacePolicy(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallInterfacePolicy6(payload *models.FirewallInterfacePolicy6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallInterfacePolicy6(mkey string, params *models.CmdbRequestParams) (*models.FirewallInterfacePolicy6, error)
	UpdateFirewallInterfacePolicy6(mkey string, payload *models.FirewallInterfacePolicy6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallInterfacePolicy6(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallInternetService(payload *models.FirewallInternetService, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallInternetService(mkey string, params *models.CmdbRequestParams) (*models.FirewallInternetService, error)
	UpdateFirewallInternetService(mkey string, payload *models.FirewallInternetService, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallInternetService(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallInternetServiceAddition(payload *models.FirewallInternetServiceAddition, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallInternetServiceAddition(mkey string, params *models.CmdbRequestParams) (*models.FirewallInternetServiceAddition, error)
	UpdateFirewallInternetServiceAddition(mkey string, payload *models.FirewallInternetServiceAddition, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallInternetServiceAddition(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallInternetServiceAppend(payload *models.FirewallInternetServiceAppend, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallInternetServiceAppend(mkey string, params *models.CmdbRequestParams) (*models.FirewallInternetServiceAppend, error)
	UpdateFirewallInternetServiceAppend(mkey string, payload *models.FirewallInternetServiceAppend, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallInternetServiceAppend(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallInternetServiceBotnet(payload *models.FirewallInternetServiceBotnet, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallInternetServiceBotnet(mkey string, params *models.CmdbRequestParams) (*models.FirewallInternetServiceBotnet, error)
	UpdateFirewallInternetServiceBotnet(mkey string, payload *models.FirewallInternetServiceBotnet, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallInternetServiceBotnet(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallInternetServiceCustom(payload *models.FirewallInternetServiceCustom, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallInternetServiceCustom(mkey string, params *models.CmdbRequestParams) (*models.FirewallInternetServiceCustom, error)
	UpdateFirewallInternetServiceCustom(mkey string, payload *models.FirewallInternetServiceCustom, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallInternetServiceCustom(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallInternetServiceCustomGroup(payload *models.FirewallInternetServiceCustomGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallInternetServiceCustomGroup(mkey string, params *models.CmdbRequestParams) (*models.FirewallInternetServiceCustomGroup, error)
	UpdateFirewallInternetServiceCustomGroup(mkey string, payload *models.FirewallInternetServiceCustomGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallInternetServiceCustomGroup(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallInternetServiceDefinition(payload *models.FirewallInternetServiceDefinition, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallInternetServiceDefinition(mkey string, params *models.CmdbRequestParams) (*models.FirewallInternetServiceDefinition, error)
	UpdateFirewallInternetServiceDefinition(mkey string, payload *models.FirewallInternetServiceDefinition, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallInternetServiceDefinition(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallInternetServiceExtension(payload *models.FirewallInternetServiceExtension, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallInternetServiceExtension(mkey string, params *models.CmdbRequestParams) (*models.FirewallInternetServiceExtension, error)
	UpdateFirewallInternetServiceExtension(mkey string, payload *models.FirewallInternetServiceExtension, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallInternetServiceExtension(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallInternetServiceGroup(payload *models.FirewallInternetServiceGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallInternetServiceGroup(mkey string, params *models.CmdbRequestParams) (*models.FirewallInternetServiceGroup, error)
	UpdateFirewallInternetServiceGroup(mkey string, payload *models.FirewallInternetServiceGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallInternetServiceGroup(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallInternetServiceIpblReason(payload *models.FirewallInternetServiceIpblReason, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallInternetServiceIpblReason(mkey string, params *models.CmdbRequestParams) (*models.FirewallInternetServiceIpblReason, error)
	UpdateFirewallInternetServiceIpblReason(mkey string, payload *models.FirewallInternetServiceIpblReason, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallInternetServiceIpblReason(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallInternetServiceIpblVendor(payload *models.FirewallInternetServiceIpblVendor, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallInternetServiceIpblVendor(mkey string, params *models.CmdbRequestParams) (*models.FirewallInternetServiceIpblVendor, error)
	UpdateFirewallInternetServiceIpblVendor(mkey string, payload *models.FirewallInternetServiceIpblVendor, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallInternetServiceIpblVendor(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallInternetServiceList(payload *models.FirewallInternetServiceList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallInternetServiceList(mkey string, params *models.CmdbRequestParams) (*models.FirewallInternetServiceList, error)
	UpdateFirewallInternetServiceList(mkey string, payload *models.FirewallInternetServiceList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallInternetServiceList(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallInternetServiceName(payload *models.FirewallInternetServiceName, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallInternetServiceName(mkey string, params *models.CmdbRequestParams) (*models.FirewallInternetServiceName, error)
	UpdateFirewallInternetServiceName(mkey string, payload *models.FirewallInternetServiceName, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallInternetServiceName(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallInternetServiceOwner(payload *models.FirewallInternetServiceOwner, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallInternetServiceOwner(mkey string, params *models.CmdbRequestParams) (*models.FirewallInternetServiceOwner, error)
	UpdateFirewallInternetServiceOwner(mkey string, payload *models.FirewallInternetServiceOwner, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallInternetServiceOwner(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallInternetServiceReputation(payload *models.FirewallInternetServiceReputation, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallInternetServiceReputation(mkey string, params *models.CmdbRequestParams) (*models.FirewallInternetServiceReputation, error)
	UpdateFirewallInternetServiceReputation(mkey string, payload *models.FirewallInternetServiceReputation, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallInternetServiceReputation(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallInternetServiceSld(payload *models.FirewallInternetServiceSld, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallInternetServiceSld(mkey string, params *models.CmdbRequestParams) (*models.FirewallInternetServiceSld, error)
	UpdateFirewallInternetServiceSld(mkey string, payload *models.FirewallInternetServiceSld, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallInternetServiceSld(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallIpTranslation(payload *models.FirewallIpTranslation, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallIpTranslation(mkey string, params *models.CmdbRequestParams) (*models.FirewallIpTranslation, error)
	UpdateFirewallIpTranslation(mkey string, payload *models.FirewallIpTranslation, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallIpTranslation(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallIppool(payload *models.FirewallIppool, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallIppool(mkey string, params *models.CmdbRequestParams) (*models.FirewallIppool, error)
	UpdateFirewallIppool(mkey string, payload *models.FirewallIppool, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallIppool(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallIppool6(payload *models.FirewallIppool6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallIppool6(mkey string, params *models.CmdbRequestParams) (*models.FirewallIppool6, error)
	UpdateFirewallIppool6(mkey string, payload *models.FirewallIppool6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallIppool6(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallLdbMonitor(payload *models.FirewallLdbMonitor, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallLdbMonitor(mkey string, params *models.CmdbRequestParams) (*models.FirewallLdbMonitor, error)
	UpdateFirewallLdbMonitor(mkey string, payload *models.FirewallLdbMonitor, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallLdbMonitor(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallLocalInPolicy(payload *models.FirewallLocalInPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallLocalInPolicy(mkey string, params *models.CmdbRequestParams) (*models.FirewallLocalInPolicy, error)
	UpdateFirewallLocalInPolicy(mkey string, payload *models.FirewallLocalInPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallLocalInPolicy(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallLocalInPolicy6(payload *models.FirewallLocalInPolicy6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallLocalInPolicy6(mkey string, params *models.CmdbRequestParams) (*models.FirewallLocalInPolicy6, error)
	UpdateFirewallLocalInPolicy6(mkey string, payload *models.FirewallLocalInPolicy6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallLocalInPolicy6(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallMulticastAddress(payload *models.FirewallMulticastAddress, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallMulticastAddress(mkey string, params *models.CmdbRequestParams) (*models.FirewallMulticastAddress, error)
	UpdateFirewallMulticastAddress(mkey string, payload *models.FirewallMulticastAddress, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallMulticastAddress(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallMulticastAddress6(payload *models.FirewallMulticastAddress6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallMulticastAddress6(mkey string, params *models.CmdbRequestParams) (*models.FirewallMulticastAddress6, error)
	UpdateFirewallMulticastAddress6(mkey string, payload *models.FirewallMulticastAddress6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallMulticastAddress6(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallMulticastPolicy(payload *models.FirewallMulticastPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallMulticastPolicy(mkey string, params *models.CmdbRequestParams) (*models.FirewallMulticastPolicy, error)
	UpdateFirewallMulticastPolicy(mkey string, payload *models.FirewallMulticastPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallMulticastPolicy(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallMulticastPolicy6(payload *models.FirewallMulticastPolicy6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallMulticastPolicy6(mkey string, params *models.CmdbRequestParams) (*models.FirewallMulticastPolicy6, error)
	UpdateFirewallMulticastPolicy6(mkey string, payload *models.FirewallMulticastPolicy6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallMulticastPolicy6(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallPolicy(payload *models.FirewallPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallPolicy(mkey string, params *models.CmdbRequestParams) (*models.FirewallPolicy, error)
	UpdateFirewallPolicy(mkey string, payload *models.FirewallPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallPolicy(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallPolicy46(payload *models.FirewallPolicy46, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallPolicy46(mkey string, params *models.CmdbRequestParams) (*models.FirewallPolicy46, error)
	UpdateFirewallPolicy46(mkey string, payload *models.FirewallPolicy46, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallPolicy46(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallPolicy6(payload *models.FirewallPolicy6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallPolicy6(mkey string, params *models.CmdbRequestParams) (*models.FirewallPolicy6, error)
	UpdateFirewallPolicy6(mkey string, payload *models.FirewallPolicy6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallPolicy6(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallPolicy64(payload *models.FirewallPolicy64, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallPolicy64(mkey string, params *models.CmdbRequestParams) (*models.FirewallPolicy64, error)
	UpdateFirewallPolicy64(mkey string, payload *models.FirewallPolicy64, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallPolicy64(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallProfileGroup(payload *models.FirewallProfileGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallProfileGroup(mkey string, params *models.CmdbRequestParams) (*models.FirewallProfileGroup, error)
	UpdateFirewallProfileGroup(mkey string, payload *models.FirewallProfileGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallProfileGroup(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallProfileProtocolOptions(payload *models.FirewallProfileProtocolOptions, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallProfileProtocolOptions(mkey string, params *models.CmdbRequestParams) (*models.FirewallProfileProtocolOptions, error)
	UpdateFirewallProfileProtocolOptions(mkey string, payload *models.FirewallProfileProtocolOptions, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallProfileProtocolOptions(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallProxyAddress(payload *models.FirewallProxyAddress, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallProxyAddress(mkey string, params *models.CmdbRequestParams) (*models.FirewallProxyAddress, error)
	UpdateFirewallProxyAddress(mkey string, payload *models.FirewallProxyAddress, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallProxyAddress(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallProxyAddrgrp(payload *models.FirewallProxyAddrgrp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallProxyAddrgrp(mkey string, params *models.CmdbRequestParams) (*models.FirewallProxyAddrgrp, error)
	UpdateFirewallProxyAddrgrp(mkey string, payload *models.FirewallProxyAddrgrp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallProxyAddrgrp(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallProxyPolicy(payload *models.FirewallProxyPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallProxyPolicy(mkey string, params *models.CmdbRequestParams) (*models.FirewallProxyPolicy, error)
	UpdateFirewallProxyPolicy(mkey string, payload *models.FirewallProxyPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallProxyPolicy(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallRegion(payload *models.FirewallRegion, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallRegion(mkey string, params *models.CmdbRequestParams) (*models.FirewallRegion, error)
	UpdateFirewallRegion(mkey string, payload *models.FirewallRegion, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallRegion(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallSecurityPolicy(payload *models.FirewallSecurityPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallSecurityPolicy(mkey string, params *models.CmdbRequestParams) (*models.FirewallSecurityPolicy, error)
	UpdateFirewallSecurityPolicy(mkey string, payload *models.FirewallSecurityPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallSecurityPolicy(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallShapingPolicy(payload *models.FirewallShapingPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallShapingPolicy(mkey string, params *models.CmdbRequestParams) (*models.FirewallShapingPolicy, error)
	UpdateFirewallShapingPolicy(mkey string, payload *models.FirewallShapingPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallShapingPolicy(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallShapingProfile(payload *models.FirewallShapingProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallShapingProfile(mkey string, params *models.CmdbRequestParams) (*models.FirewallShapingProfile, error)
	UpdateFirewallShapingProfile(mkey string, payload *models.FirewallShapingProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallShapingProfile(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallSniffer(payload *models.FirewallSniffer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallSniffer(mkey string, params *models.CmdbRequestParams) (*models.FirewallSniffer, error)
	UpdateFirewallSniffer(mkey string, payload *models.FirewallSniffer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallSniffer(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallSslServer(payload *models.FirewallSslServer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallSslServer(mkey string, params *models.CmdbRequestParams) (*models.FirewallSslServer, error)
	UpdateFirewallSslServer(mkey string, payload *models.FirewallSslServer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallSslServer(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallSslSshProfile(payload *models.FirewallSslSshProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallSslSshProfile(mkey string, params *models.CmdbRequestParams) (*models.FirewallSslSshProfile, error)
	UpdateFirewallSslSshProfile(mkey string, payload *models.FirewallSslSshProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallSslSshProfile(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallTrafficClass(payload *models.FirewallTrafficClass, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallTrafficClass(mkey string, params *models.CmdbRequestParams) (*models.FirewallTrafficClass, error)
	UpdateFirewallTrafficClass(mkey string, payload *models.FirewallTrafficClass, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallTrafficClass(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallTtlPolicy(payload *models.FirewallTtlPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallTtlPolicy(mkey string, params *models.CmdbRequestParams) (*models.FirewallTtlPolicy, error)
	UpdateFirewallTtlPolicy(mkey string, payload *models.FirewallTtlPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallTtlPolicy(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallVendorMac(payload *models.FirewallVendorMac, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallVendorMac(mkey string, params *models.CmdbRequestParams) (*models.FirewallVendorMac, error)
	UpdateFirewallVendorMac(mkey string, payload *models.FirewallVendorMac, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallVendorMac(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallVendorMacSummary(payload *models.FirewallVendorMacSummary, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallVendorMacSummary(mkey string, params *models.CmdbRequestParams) (*models.FirewallVendorMacSummary, error)
	UpdateFirewallVendorMacSummary(mkey string, payload *models.FirewallVendorMacSummary, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallVendorMacSummary(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallVip(payload *models.FirewallVip, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallVip(mkey string, params *models.CmdbRequestParams) (*models.FirewallVip, error)
	UpdateFirewallVip(mkey string, payload *models.FirewallVip, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallVip(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallVip46(payload *models.FirewallVip46, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallVip46(mkey string, params *models.CmdbRequestParams) (*models.FirewallVip46, error)
	UpdateFirewallVip46(mkey string, payload *models.FirewallVip46, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallVip46(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallVip6(payload *models.FirewallVip6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallVip6(mkey string, params *models.CmdbRequestParams) (*models.FirewallVip6, error)
	UpdateFirewallVip6(mkey string, payload *models.FirewallVip6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallVip6(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallVip64(payload *models.FirewallVip64, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallVip64(mkey string, params *models.CmdbRequestParams) (*models.FirewallVip64, error)
	UpdateFirewallVip64(mkey string, payload *models.FirewallVip64, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallVip64(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallVipgrp(payload *models.FirewallVipgrp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallVipgrp(mkey string, params *models.CmdbRequestParams) (*models.FirewallVipgrp, error)
	UpdateFirewallVipgrp(mkey string, payload *models.FirewallVipgrp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallVipgrp(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallVipgrp46(payload *models.FirewallVipgrp46, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallVipgrp46(mkey string, params *models.CmdbRequestParams) (*models.FirewallVipgrp46, error)
	UpdateFirewallVipgrp46(mkey string, payload *models.FirewallVipgrp46, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallVipgrp46(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallVipgrp6(payload *models.FirewallVipgrp6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallVipgrp6(mkey string, params *models.CmdbRequestParams) (*models.FirewallVipgrp6, error)
	UpdateFirewallVipgrp6(mkey string, payload *models.FirewallVipgrp6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallVipgrp6(mkey string, params *models.CmdbRequestParams) error
	CreateFirewallVipgrp64(payload *models.FirewallVipgrp64, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFirewallVipgrp64(mkey string, params *models.CmdbRequestParams) (*models.FirewallVipgrp64, error)
	UpdateFirewallVipgrp64(mkey string, payload *models.FirewallVipgrp64, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFirewallVipgrp64(mkey string, params *models.CmdbRequestParams) error
	CreateFtpProxyExplicit(payload *models.FtpProxyExplicit, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadFtpProxyExplicit(mkey string, params *models.CmdbRequestParams) (*models.FtpProxyExplicit, error)
	UpdateFtpProxyExplicit(mkey string, payload *models.FtpProxyExplicit, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteFtpProxyExplicit(mkey string, params *models.CmdbRequestParams) error
	CreateIcapProfile(payload *models.IcapProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadIcapProfile(mkey string, params *models.CmdbRequestParams) (*models.IcapProfile, error)
	UpdateIcapProfile(mkey string, payload *models.IcapProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteIcapProfile(mkey string, params *models.CmdbRequestParams) error
	CreateIcapServer(payload *models.IcapServer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadIcapServer(mkey string, params *models.CmdbRequestParams) (*models.IcapServer, error)
	UpdateIcapServer(mkey string, payload *models.IcapServer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteIcapServer(mkey string, params *models.CmdbRequestParams) error
	CreateIpsCustom(payload *models.IpsCustom, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadIpsCustom(mkey string, params *models.CmdbRequestParams) (*models.IpsCustom, error)
	UpdateIpsCustom(mkey string, payload *models.IpsCustom, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteIpsCustom(mkey string, params *models.CmdbRequestParams) error
	CreateIpsDecoder(payload *models.IpsDecoder, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadIpsDecoder(mkey string, params *models.CmdbRequestParams) (*models.IpsDecoder, error)
	UpdateIpsDecoder(mkey string, payload *models.IpsDecoder, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteIpsDecoder(mkey string, params *models.CmdbRequestParams) error
	CreateIpsGlobal(payload *models.IpsGlobal, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadIpsGlobal(mkey string, params *models.CmdbRequestParams) (*models.IpsGlobal, error)
	UpdateIpsGlobal(mkey string, payload *models.IpsGlobal, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteIpsGlobal(mkey string, params *models.CmdbRequestParams) error
	CreateIpsRule(payload *models.IpsRule, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadIpsRule(mkey string, params *models.CmdbRequestParams) (*models.IpsRule, error)
	UpdateIpsRule(mkey string, payload *models.IpsRule, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteIpsRule(mkey string, params *models.CmdbRequestParams) error
	CreateIpsRuleSettings(payload *models.IpsRuleSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadIpsRuleSettings(mkey string, params *models.CmdbRequestParams) (*models.IpsRuleSettings, error)
	UpdateIpsRuleSettings(mkey string, payload *models.IpsRuleSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteIpsRuleSettings(mkey string, params *models.CmdbRequestParams) error
	CreateIpsSensor(payload *models.IpsSensor, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadIpsSensor(mkey string, params *models.CmdbRequestParams) (*models.IpsSensor, error)
	UpdateIpsSensor(mkey string, payload *models.IpsSensor, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteIpsSensor(mkey string, params *models.CmdbRequestParams) error
	CreateIpsSettings(payload *models.IpsSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadIpsSettings(mkey string, params *models.CmdbRequestParams) (*models.IpsSettings, error)
	UpdateIpsSettings(mkey string, payload *models.IpsSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteIpsSettings(mkey string, params *models.CmdbRequestParams) error
	CreateIpsViewMap(payload *models.IpsViewMap, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadIpsViewMap(mkey string, params *models.CmdbRequestParams) (*models.IpsViewMap, error)
	UpdateIpsViewMap(mkey string, payload *models.IpsViewMap, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteIpsViewMap(mkey string, params *models.CmdbRequestParams) error
	CreateLogdiskFilter(payload *models.LogdiskFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogdiskFilter(mkey string, params *models.CmdbRequestParams) (*models.LogdiskFilter, error)
	UpdateLogdiskFilter(mkey string, payload *models.LogdiskFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogdiskFilter(mkey string, params *models.CmdbRequestParams) error
	CreateLogdiskSetting(payload *models.LogdiskSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogdiskSetting(mkey string, params *models.CmdbRequestParams) (*models.LogdiskSetting, error)
	UpdateLogdiskSetting(mkey string, payload *models.LogdiskSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogdiskSetting(mkey string, params *models.CmdbRequestParams) error
	CreateLogfortianalyzer2Filter(payload *models.Logfortianalyzer2Filter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogfortianalyzer2Filter(mkey string, params *models.CmdbRequestParams) (*models.Logfortianalyzer2Filter, error)
	UpdateLogfortianalyzer2Filter(mkey string, payload *models.Logfortianalyzer2Filter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogfortianalyzer2Filter(mkey string, params *models.CmdbRequestParams) error
	CreateLogfortianalyzer2OverrideFilter(payload *models.Logfortianalyzer2OverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogfortianalyzer2OverrideFilter(mkey string, params *models.CmdbRequestParams) (*models.Logfortianalyzer2OverrideFilter, error)
	UpdateLogfortianalyzer2OverrideFilter(mkey string, payload *models.Logfortianalyzer2OverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogfortianalyzer2OverrideFilter(mkey string, params *models.CmdbRequestParams) error
	CreateLogfortianalyzer2OverrideSetting(payload *models.Logfortianalyzer2OverrideSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogfortianalyzer2OverrideSetting(mkey string, params *models.CmdbRequestParams) (*models.Logfortianalyzer2OverrideSetting, error)
	UpdateLogfortianalyzer2OverrideSetting(mkey string, payload *models.Logfortianalyzer2OverrideSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogfortianalyzer2OverrideSetting(mkey string, params *models.CmdbRequestParams) error
	CreateLogfortianalyzer2Setting(payload *models.Logfortianalyzer2Setting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogfortianalyzer2Setting(mkey string, params *models.CmdbRequestParams) (*models.Logfortianalyzer2Setting, error)
	UpdateLogfortianalyzer2Setting(mkey string, payload *models.Logfortianalyzer2Setting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogfortianalyzer2Setting(mkey string, params *models.CmdbRequestParams) error
	CreateLogfortianalyzer3Filter(payload *models.Logfortianalyzer3Filter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogfortianalyzer3Filter(mkey string, params *models.CmdbRequestParams) (*models.Logfortianalyzer3Filter, error)
	UpdateLogfortianalyzer3Filter(mkey string, payload *models.Logfortianalyzer3Filter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogfortianalyzer3Filter(mkey string, params *models.CmdbRequestParams) error
	CreateLogfortianalyzer3OverrideFilter(payload *models.Logfortianalyzer3OverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogfortianalyzer3OverrideFilter(mkey string, params *models.CmdbRequestParams) (*models.Logfortianalyzer3OverrideFilter, error)
	UpdateLogfortianalyzer3OverrideFilter(mkey string, payload *models.Logfortianalyzer3OverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogfortianalyzer3OverrideFilter(mkey string, params *models.CmdbRequestParams) error
	CreateLogfortianalyzer3OverrideSetting(payload *models.Logfortianalyzer3OverrideSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogfortianalyzer3OverrideSetting(mkey string, params *models.CmdbRequestParams) (*models.Logfortianalyzer3OverrideSetting, error)
	UpdateLogfortianalyzer3OverrideSetting(mkey string, payload *models.Logfortianalyzer3OverrideSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogfortianalyzer3OverrideSetting(mkey string, params *models.CmdbRequestParams) error
	CreateLogfortianalyzer3Setting(payload *models.Logfortianalyzer3Setting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogfortianalyzer3Setting(mkey string, params *models.CmdbRequestParams) (*models.Logfortianalyzer3Setting, error)
	UpdateLogfortianalyzer3Setting(mkey string, payload *models.Logfortianalyzer3Setting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogfortianalyzer3Setting(mkey string, params *models.CmdbRequestParams) error
	CreateLogfortianalyzerCloudFilter(payload *models.LogfortianalyzerCloudFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogfortianalyzerCloudFilter(mkey string, params *models.CmdbRequestParams) (*models.LogfortianalyzerCloudFilter, error)
	UpdateLogfortianalyzerCloudFilter(mkey string, payload *models.LogfortianalyzerCloudFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogfortianalyzerCloudFilter(mkey string, params *models.CmdbRequestParams) error
	CreateLogfortianalyzerCloudOverrideFilter(payload *models.LogfortianalyzerCloudOverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogfortianalyzerCloudOverrideFilter(mkey string, params *models.CmdbRequestParams) (*models.LogfortianalyzerCloudOverrideFilter, error)
	UpdateLogfortianalyzerCloudOverrideFilter(mkey string, payload *models.LogfortianalyzerCloudOverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogfortianalyzerCloudOverrideFilter(mkey string, params *models.CmdbRequestParams) error
	CreateLogfortianalyzerCloudOverrideSetting(payload *models.LogfortianalyzerCloudOverrideSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogfortianalyzerCloudOverrideSetting(mkey string, params *models.CmdbRequestParams) (*models.LogfortianalyzerCloudOverrideSetting, error)
	UpdateLogfortianalyzerCloudOverrideSetting(mkey string, payload *models.LogfortianalyzerCloudOverrideSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogfortianalyzerCloudOverrideSetting(mkey string, params *models.CmdbRequestParams) error
	CreateLogfortianalyzerCloudSetting(payload *models.LogfortianalyzerCloudSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogfortianalyzerCloudSetting(mkey string, params *models.CmdbRequestParams) (*models.LogfortianalyzerCloudSetting, error)
	UpdateLogfortianalyzerCloudSetting(mkey string, payload *models.LogfortianalyzerCloudSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogfortianalyzerCloudSetting(mkey string, params *models.CmdbRequestParams) error
	CreateLogfortianalyzerFilter(payload *models.LogfortianalyzerFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogfortianalyzerFilter(mkey string, params *models.CmdbRequestParams) (*models.LogfortianalyzerFilter, error)
	UpdateLogfortianalyzerFilter(mkey string, payload *models.LogfortianalyzerFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogfortianalyzerFilter(mkey string, params *models.CmdbRequestParams) error
	CreateLogfortianalyzerOverrideFilter(payload *models.LogfortianalyzerOverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogfortianalyzerOverrideFilter(mkey string, params *models.CmdbRequestParams) (*models.LogfortianalyzerOverrideFilter, error)
	UpdateLogfortianalyzerOverrideFilter(mkey string, payload *models.LogfortianalyzerOverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogfortianalyzerOverrideFilter(mkey string, params *models.CmdbRequestParams) error
	CreateLogfortianalyzerOverrideSetting(payload *models.LogfortianalyzerOverrideSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogfortianalyzerOverrideSetting(mkey string, params *models.CmdbRequestParams) (*models.LogfortianalyzerOverrideSetting, error)
	UpdateLogfortianalyzerOverrideSetting(mkey string, payload *models.LogfortianalyzerOverrideSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogfortianalyzerOverrideSetting(mkey string, params *models.CmdbRequestParams) error
	CreateLogfortianalyzerSetting(payload *models.LogfortianalyzerSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogfortianalyzerSetting(mkey string, params *models.CmdbRequestParams) (*models.LogfortianalyzerSetting, error)
	UpdateLogfortianalyzerSetting(mkey string, payload *models.LogfortianalyzerSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogfortianalyzerSetting(mkey string, params *models.CmdbRequestParams) error
	CreateLogfortiguardFilter(payload *models.LogfortiguardFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogfortiguardFilter(mkey string, params *models.CmdbRequestParams) (*models.LogfortiguardFilter, error)
	UpdateLogfortiguardFilter(mkey string, payload *models.LogfortiguardFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogfortiguardFilter(mkey string, params *models.CmdbRequestParams) error
	CreateLogfortiguardOverrideFilter(payload *models.LogfortiguardOverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogfortiguardOverrideFilter(mkey string, params *models.CmdbRequestParams) (*models.LogfortiguardOverrideFilter, error)
	UpdateLogfortiguardOverrideFilter(mkey string, payload *models.LogfortiguardOverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogfortiguardOverrideFilter(mkey string, params *models.CmdbRequestParams) error
	CreateLogfortiguardOverrideSetting(payload *models.LogfortiguardOverrideSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogfortiguardOverrideSetting(mkey string, params *models.CmdbRequestParams) (*models.LogfortiguardOverrideSetting, error)
	UpdateLogfortiguardOverrideSetting(mkey string, payload *models.LogfortiguardOverrideSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogfortiguardOverrideSetting(mkey string, params *models.CmdbRequestParams) error
	CreateLogfortiguardSetting(payload *models.LogfortiguardSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogfortiguardSetting(mkey string, params *models.CmdbRequestParams) (*models.LogfortiguardSetting, error)
	UpdateLogfortiguardSetting(mkey string, payload *models.LogfortiguardSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogfortiguardSetting(mkey string, params *models.CmdbRequestParams) error
	CreateLogmemoryFilter(payload *models.LogmemoryFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogmemoryFilter(mkey string, params *models.CmdbRequestParams) (*models.LogmemoryFilter, error)
	UpdateLogmemoryFilter(mkey string, payload *models.LogmemoryFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogmemoryFilter(mkey string, params *models.CmdbRequestParams) error
	CreateLogmemoryGlobalSetting(payload *models.LogmemoryGlobalSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogmemoryGlobalSetting(mkey string, params *models.CmdbRequestParams) (*models.LogmemoryGlobalSetting, error)
	UpdateLogmemoryGlobalSetting(mkey string, payload *models.LogmemoryGlobalSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogmemoryGlobalSetting(mkey string, params *models.CmdbRequestParams) error
	CreateLogmemorySetting(payload *models.LogmemorySetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogmemorySetting(mkey string, params *models.CmdbRequestParams) (*models.LogmemorySetting, error)
	UpdateLogmemorySetting(mkey string, payload *models.LogmemorySetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogmemorySetting(mkey string, params *models.CmdbRequestParams) error
	CreateLognullDeviceFilter(payload *models.LognullDeviceFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLognullDeviceFilter(mkey string, params *models.CmdbRequestParams) (*models.LognullDeviceFilter, error)
	UpdateLognullDeviceFilter(mkey string, payload *models.LognullDeviceFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLognullDeviceFilter(mkey string, params *models.CmdbRequestParams) error
	CreateLognullDeviceSetting(payload *models.LognullDeviceSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLognullDeviceSetting(mkey string, params *models.CmdbRequestParams) (*models.LognullDeviceSetting, error)
	UpdateLognullDeviceSetting(mkey string, payload *models.LognullDeviceSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLognullDeviceSetting(mkey string, params *models.CmdbRequestParams) error
	CreateLogsyslogd2Filter(payload *models.Logsyslogd2Filter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogsyslogd2Filter(mkey string, params *models.CmdbRequestParams) (*models.Logsyslogd2Filter, error)
	UpdateLogsyslogd2Filter(mkey string, payload *models.Logsyslogd2Filter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogsyslogd2Filter(mkey string, params *models.CmdbRequestParams) error
	CreateLogsyslogd2OverrideFilter(payload *models.Logsyslogd2OverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogsyslogd2OverrideFilter(mkey string, params *models.CmdbRequestParams) (*models.Logsyslogd2OverrideFilter, error)
	UpdateLogsyslogd2OverrideFilter(mkey string, payload *models.Logsyslogd2OverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogsyslogd2OverrideFilter(mkey string, params *models.CmdbRequestParams) error
	CreateLogsyslogd2OverrideSetting(payload *models.Logsyslogd2OverrideSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogsyslogd2OverrideSetting(mkey string, params *models.CmdbRequestParams) (*models.Logsyslogd2OverrideSetting, error)
	UpdateLogsyslogd2OverrideSetting(mkey string, payload *models.Logsyslogd2OverrideSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogsyslogd2OverrideSetting(mkey string, params *models.CmdbRequestParams) error
	CreateLogsyslogd2Setting(payload *models.Logsyslogd2Setting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogsyslogd2Setting(mkey string, params *models.CmdbRequestParams) (*models.Logsyslogd2Setting, error)
	UpdateLogsyslogd2Setting(mkey string, payload *models.Logsyslogd2Setting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogsyslogd2Setting(mkey string, params *models.CmdbRequestParams) error
	CreateLogsyslogd3Filter(payload *models.Logsyslogd3Filter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogsyslogd3Filter(mkey string, params *models.CmdbRequestParams) (*models.Logsyslogd3Filter, error)
	UpdateLogsyslogd3Filter(mkey string, payload *models.Logsyslogd3Filter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogsyslogd3Filter(mkey string, params *models.CmdbRequestParams) error
	CreateLogsyslogd3OverrideFilter(payload *models.Logsyslogd3OverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogsyslogd3OverrideFilter(mkey string, params *models.CmdbRequestParams) (*models.Logsyslogd3OverrideFilter, error)
	UpdateLogsyslogd3OverrideFilter(mkey string, payload *models.Logsyslogd3OverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogsyslogd3OverrideFilter(mkey string, params *models.CmdbRequestParams) error
	CreateLogsyslogd3OverrideSetting(payload *models.Logsyslogd3OverrideSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogsyslogd3OverrideSetting(mkey string, params *models.CmdbRequestParams) (*models.Logsyslogd3OverrideSetting, error)
	UpdateLogsyslogd3OverrideSetting(mkey string, payload *models.Logsyslogd3OverrideSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogsyslogd3OverrideSetting(mkey string, params *models.CmdbRequestParams) error
	CreateLogsyslogd3Setting(payload *models.Logsyslogd3Setting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogsyslogd3Setting(mkey string, params *models.CmdbRequestParams) (*models.Logsyslogd3Setting, error)
	UpdateLogsyslogd3Setting(mkey string, payload *models.Logsyslogd3Setting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogsyslogd3Setting(mkey string, params *models.CmdbRequestParams) error
	CreateLogsyslogd4Filter(payload *models.Logsyslogd4Filter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogsyslogd4Filter(mkey string, params *models.CmdbRequestParams) (*models.Logsyslogd4Filter, error)
	UpdateLogsyslogd4Filter(mkey string, payload *models.Logsyslogd4Filter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogsyslogd4Filter(mkey string, params *models.CmdbRequestParams) error
	CreateLogsyslogd4OverrideFilter(payload *models.Logsyslogd4OverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogsyslogd4OverrideFilter(mkey string, params *models.CmdbRequestParams) (*models.Logsyslogd4OverrideFilter, error)
	UpdateLogsyslogd4OverrideFilter(mkey string, payload *models.Logsyslogd4OverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogsyslogd4OverrideFilter(mkey string, params *models.CmdbRequestParams) error
	CreateLogsyslogd4OverrideSetting(payload *models.Logsyslogd4OverrideSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogsyslogd4OverrideSetting(mkey string, params *models.CmdbRequestParams) (*models.Logsyslogd4OverrideSetting, error)
	UpdateLogsyslogd4OverrideSetting(mkey string, payload *models.Logsyslogd4OverrideSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogsyslogd4OverrideSetting(mkey string, params *models.CmdbRequestParams) error
	CreateLogsyslogd4Setting(payload *models.Logsyslogd4Setting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogsyslogd4Setting(mkey string, params *models.CmdbRequestParams) (*models.Logsyslogd4Setting, error)
	UpdateLogsyslogd4Setting(mkey string, payload *models.Logsyslogd4Setting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogsyslogd4Setting(mkey string, params *models.CmdbRequestParams) error
	CreateLogsyslogdFilter(payload *models.LogsyslogdFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogsyslogdFilter(mkey string, params *models.CmdbRequestParams) (*models.LogsyslogdFilter, error)
	UpdateLogsyslogdFilter(mkey string, payload *models.LogsyslogdFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogsyslogdFilter(mkey string, params *models.CmdbRequestParams) error
	CreateLogsyslogdOverrideFilter(payload *models.LogsyslogdOverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogsyslogdOverrideFilter(mkey string, params *models.CmdbRequestParams) (*models.LogsyslogdOverrideFilter, error)
	UpdateLogsyslogdOverrideFilter(mkey string, payload *models.LogsyslogdOverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogsyslogdOverrideFilter(mkey string, params *models.CmdbRequestParams) error
	CreateLogsyslogdOverrideSetting(payload *models.LogsyslogdOverrideSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogsyslogdOverrideSetting(mkey string, params *models.CmdbRequestParams) (*models.LogsyslogdOverrideSetting, error)
	UpdateLogsyslogdOverrideSetting(mkey string, payload *models.LogsyslogdOverrideSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogsyslogdOverrideSetting(mkey string, params *models.CmdbRequestParams) error
	CreateLogsyslogdSetting(payload *models.LogsyslogdSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogsyslogdSetting(mkey string, params *models.CmdbRequestParams) (*models.LogsyslogdSetting, error)
	UpdateLogsyslogdSetting(mkey string, payload *models.LogsyslogdSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogsyslogdSetting(mkey string, params *models.CmdbRequestParams) error
	CreateLogtacacsaccounting2Filter(payload *models.Logtacacsaccounting2Filter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogtacacsaccounting2Filter(mkey string, params *models.CmdbRequestParams) (*models.Logtacacsaccounting2Filter, error)
	UpdateLogtacacsaccounting2Filter(mkey string, payload *models.Logtacacsaccounting2Filter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogtacacsaccounting2Filter(mkey string, params *models.CmdbRequestParams) error
	CreateLogtacacsaccounting2Setting(payload *models.Logtacacsaccounting2Setting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogtacacsaccounting2Setting(mkey string, params *models.CmdbRequestParams) (*models.Logtacacsaccounting2Setting, error)
	UpdateLogtacacsaccounting2Setting(mkey string, payload *models.Logtacacsaccounting2Setting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogtacacsaccounting2Setting(mkey string, params *models.CmdbRequestParams) error
	CreateLogtacacsaccounting3Filter(payload *models.Logtacacsaccounting3Filter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogtacacsaccounting3Filter(mkey string, params *models.CmdbRequestParams) (*models.Logtacacsaccounting3Filter, error)
	UpdateLogtacacsaccounting3Filter(mkey string, payload *models.Logtacacsaccounting3Filter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogtacacsaccounting3Filter(mkey string, params *models.CmdbRequestParams) error
	CreateLogtacacsaccounting3Setting(payload *models.Logtacacsaccounting3Setting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogtacacsaccounting3Setting(mkey string, params *models.CmdbRequestParams) (*models.Logtacacsaccounting3Setting, error)
	UpdateLogtacacsaccounting3Setting(mkey string, payload *models.Logtacacsaccounting3Setting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogtacacsaccounting3Setting(mkey string, params *models.CmdbRequestParams) error
	CreateLogtacacsaccountingFilter(payload *models.LogtacacsaccountingFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogtacacsaccountingFilter(mkey string, params *models.CmdbRequestParams) (*models.LogtacacsaccountingFilter, error)
	UpdateLogtacacsaccountingFilter(mkey string, payload *models.LogtacacsaccountingFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogtacacsaccountingFilter(mkey string, params *models.CmdbRequestParams) error
	CreateLogtacacsaccountingSetting(payload *models.LogtacacsaccountingSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogtacacsaccountingSetting(mkey string, params *models.CmdbRequestParams) (*models.LogtacacsaccountingSetting, error)
	UpdateLogtacacsaccountingSetting(mkey string, payload *models.LogtacacsaccountingSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogtacacsaccountingSetting(mkey string, params *models.CmdbRequestParams) error
	CreateLogwebtrendsFilter(payload *models.LogwebtrendsFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogwebtrendsFilter(mkey string, params *models.CmdbRequestParams) (*models.LogwebtrendsFilter, error)
	UpdateLogwebtrendsFilter(mkey string, payload *models.LogwebtrendsFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogwebtrendsFilter(mkey string, params *models.CmdbRequestParams) error
	CreateLogwebtrendsSetting(payload *models.LogwebtrendsSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogwebtrendsSetting(mkey string, params *models.CmdbRequestParams) (*models.LogwebtrendsSetting, error)
	UpdateLogwebtrendsSetting(mkey string, payload *models.LogwebtrendsSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogwebtrendsSetting(mkey string, params *models.CmdbRequestParams) error
	CreateLogCustomField(payload *models.LogCustomField, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogCustomField(mkey string, params *models.CmdbRequestParams) (*models.LogCustomField, error)
	UpdateLogCustomField(mkey string, payload *models.LogCustomField, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogCustomField(mkey string, params *models.CmdbRequestParams) error
	CreateLogEventfilter(payload *models.LogEventfilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogEventfilter(mkey string, params *models.CmdbRequestParams) (*models.LogEventfilter, error)
	UpdateLogEventfilter(mkey string, payload *models.LogEventfilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogEventfilter(mkey string, params *models.CmdbRequestParams) error
	CreateLogGuiDisplay(payload *models.LogGuiDisplay, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogGuiDisplay(mkey string, params *models.CmdbRequestParams) (*models.LogGuiDisplay, error)
	UpdateLogGuiDisplay(mkey string, payload *models.LogGuiDisplay, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogGuiDisplay(mkey string, params *models.CmdbRequestParams) error
	CreateLogSetting(payload *models.LogSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogSetting(mkey string, params *models.CmdbRequestParams) (*models.LogSetting, error)
	UpdateLogSetting(mkey string, payload *models.LogSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogSetting(mkey string, params *models.CmdbRequestParams) error
	CreateLogThreatWeight(payload *models.LogThreatWeight, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadLogThreatWeight(mkey string, params *models.CmdbRequestParams) (*models.LogThreatWeight, error)
	UpdateLogThreatWeight(mkey string, payload *models.LogThreatWeight, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteLogThreatWeight(mkey string, params *models.CmdbRequestParams) error
	CreateReportChart(payload *models.ReportChart, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadReportChart(mkey string, params *models.CmdbRequestParams) (*models.ReportChart, error)
	UpdateReportChart(mkey string, payload *models.ReportChart, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteReportChart(mkey string, params *models.CmdbRequestParams) error
	CreateReportDataset(payload *models.ReportDataset, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadReportDataset(mkey string, params *models.CmdbRequestParams) (*models.ReportDataset, error)
	UpdateReportDataset(mkey string, payload *models.ReportDataset, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteReportDataset(mkey string, params *models.CmdbRequestParams) error
	CreateReportLayout(payload *models.ReportLayout, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadReportLayout(mkey string, params *models.CmdbRequestParams) (*models.ReportLayout, error)
	UpdateReportLayout(mkey string, payload *models.ReportLayout, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteReportLayout(mkey string, params *models.CmdbRequestParams) error
	CreateReportSetting(payload *models.ReportSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadReportSetting(mkey string, params *models.CmdbRequestParams) (*models.ReportSetting, error)
	UpdateReportSetting(mkey string, payload *models.ReportSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteReportSetting(mkey string, params *models.CmdbRequestParams) error
	CreateReportStyle(payload *models.ReportStyle, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadReportStyle(mkey string, params *models.CmdbRequestParams) (*models.ReportStyle, error)
	UpdateReportStyle(mkey string, payload *models.ReportStyle, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteReportStyle(mkey string, params *models.CmdbRequestParams) error
	CreateReportTheme(payload *models.ReportTheme, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadReportTheme(mkey string, params *models.CmdbRequestParams) (*models.ReportTheme, error)
	UpdateReportTheme(mkey string, payload *models.ReportTheme, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteReportTheme(mkey string, params *models.CmdbRequestParams) error
	CreateRouteraccessList6Rule(payload *models.RouteraccessList6Rule, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouteraccessList6Rule(mkey string, params *models.CmdbRequestParams) (*models.RouteraccessList6Rule, error)
	UpdateRouteraccessList6Rule(mkey string, payload *models.RouteraccessList6Rule, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouteraccessList6Rule(mkey string, params *models.CmdbRequestParams) error
	CreateRouteraccessListRule(payload *models.RouteraccessListRule, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouteraccessListRule(mkey string, params *models.CmdbRequestParams) (*models.RouteraccessListRule, error)
	UpdateRouteraccessListRule(mkey string, payload *models.RouteraccessListRule, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouteraccessListRule(mkey string, params *models.CmdbRequestParams) error
	CreateRouterbgpNeighbor(payload *models.RouterbgpNeighbor, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterbgpNeighbor(mkey string, params *models.CmdbRequestParams) (*models.RouterbgpNeighbor, error)
	UpdateRouterbgpNeighbor(mkey string, payload *models.RouterbgpNeighbor, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterbgpNeighbor(mkey string, params *models.CmdbRequestParams) error
	CreateRouterbgpNeighborGroup(payload *models.RouterbgpNeighborGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterbgpNeighborGroup(mkey string, params *models.CmdbRequestParams) (*models.RouterbgpNeighborGroup, error)
	UpdateRouterbgpNeighborGroup(mkey string, payload *models.RouterbgpNeighborGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterbgpNeighborGroup(mkey string, params *models.CmdbRequestParams) error
	CreateRouterbgpNeighborRange(payload *models.RouterbgpNeighborRange, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterbgpNeighborRange(mkey string, params *models.CmdbRequestParams) (*models.RouterbgpNeighborRange, error)
	UpdateRouterbgpNeighborRange(mkey string, payload *models.RouterbgpNeighborRange, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterbgpNeighborRange(mkey string, params *models.CmdbRequestParams) error
	CreateRouterbgpNetwork(payload *models.RouterbgpNetwork, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterbgpNetwork(mkey string, params *models.CmdbRequestParams) (*models.RouterbgpNetwork, error)
	UpdateRouterbgpNetwork(mkey string, payload *models.RouterbgpNetwork, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterbgpNetwork(mkey string, params *models.CmdbRequestParams) error
	CreateRouterbgpNetwork6(payload *models.RouterbgpNetwork6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterbgpNetwork6(mkey string, params *models.CmdbRequestParams) (*models.RouterbgpNetwork6, error)
	UpdateRouterbgpNetwork6(mkey string, payload *models.RouterbgpNetwork6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterbgpNetwork6(mkey string, params *models.CmdbRequestParams) error
	CreateRouterbgpRedistribute(payload *models.RouterbgpRedistribute, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterbgpRedistribute(mkey string, params *models.CmdbRequestParams) (*models.RouterbgpRedistribute, error)
	UpdateRouterbgpRedistribute(mkey string, payload *models.RouterbgpRedistribute, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterbgpRedistribute(mkey string, params *models.CmdbRequestParams) error
	CreateRouterospf6Area(payload *models.Routerospf6Area, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterospf6Area(mkey string, params *models.CmdbRequestParams) (*models.Routerospf6Area, error)
	UpdateRouterospf6Area(mkey string, payload *models.Routerospf6Area, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterospf6Area(mkey string, params *models.CmdbRequestParams) error
	CreateRouterospf6Ospf6Interface(payload *models.Routerospf6Ospf6Interface, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterospf6Ospf6Interface(mkey string, params *models.CmdbRequestParams) (*models.Routerospf6Ospf6Interface, error)
	UpdateRouterospf6Ospf6Interface(mkey string, payload *models.Routerospf6Ospf6Interface, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterospf6Ospf6Interface(mkey string, params *models.CmdbRequestParams) error
	CreateRouterospf6Redistribute(payload *models.Routerospf6Redistribute, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterospf6Redistribute(mkey string, params *models.CmdbRequestParams) (*models.Routerospf6Redistribute, error)
	UpdateRouterospf6Redistribute(mkey string, payload *models.Routerospf6Redistribute, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterospf6Redistribute(mkey string, params *models.CmdbRequestParams) error
	CreateRouterospfArea(payload *models.RouterospfArea, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterospfArea(mkey string, params *models.CmdbRequestParams) (*models.RouterospfArea, error)
	UpdateRouterospfArea(mkey string, payload *models.RouterospfArea, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterospfArea(mkey string, params *models.CmdbRequestParams) error
	CreateRouterospfNeighbor(payload *models.RouterospfNeighbor, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterospfNeighbor(mkey string, params *models.CmdbRequestParams) (*models.RouterospfNeighbor, error)
	UpdateRouterospfNeighbor(mkey string, payload *models.RouterospfNeighbor, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterospfNeighbor(mkey string, params *models.CmdbRequestParams) error
	CreateRouterospfNetwork(payload *models.RouterospfNetwork, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterospfNetwork(mkey string, params *models.CmdbRequestParams) (*models.RouterospfNetwork, error)
	UpdateRouterospfNetwork(mkey string, payload *models.RouterospfNetwork, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterospfNetwork(mkey string, params *models.CmdbRequestParams) error
	CreateRouterospfOspfInterface(payload *models.RouterospfOspfInterface, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterospfOspfInterface(mkey string, params *models.CmdbRequestParams) (*models.RouterospfOspfInterface, error)
	UpdateRouterospfOspfInterface(mkey string, payload *models.RouterospfOspfInterface, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterospfOspfInterface(mkey string, params *models.CmdbRequestParams) error
	CreateRouterospfRedistribute(payload *models.RouterospfRedistribute, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterospfRedistribute(mkey string, params *models.CmdbRequestParams) (*models.RouterospfRedistribute, error)
	UpdateRouterospfRedistribute(mkey string, payload *models.RouterospfRedistribute, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterospfRedistribute(mkey string, params *models.CmdbRequestParams) error
	CreateRouterprefixList6Rule(payload *models.RouterprefixList6Rule, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterprefixList6Rule(mkey string, params *models.CmdbRequestParams) (*models.RouterprefixList6Rule, error)
	UpdateRouterprefixList6Rule(mkey string, payload *models.RouterprefixList6Rule, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterprefixList6Rule(mkey string, params *models.CmdbRequestParams) error
	CreateRouterprefixListRule(payload *models.RouterprefixListRule, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterprefixListRule(mkey string, params *models.CmdbRequestParams) (*models.RouterprefixListRule, error)
	UpdateRouterprefixListRule(mkey string, payload *models.RouterprefixListRule, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterprefixListRule(mkey string, params *models.CmdbRequestParams) error
	CreateRouterrouteMapRule(payload *models.RouterrouteMapRule, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterrouteMapRule(mkey string, params *models.CmdbRequestParams) (*models.RouterrouteMapRule, error)
	UpdateRouterrouteMapRule(mkey string, payload *models.RouterrouteMapRule, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterrouteMapRule(mkey string, params *models.CmdbRequestParams) error
	CreateRouterAccessList(payload *models.RouterAccessList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterAccessList(mkey string, params *models.CmdbRequestParams) (*models.RouterAccessList, error)
	UpdateRouterAccessList(mkey string, payload *models.RouterAccessList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterAccessList(mkey string, params *models.CmdbRequestParams) error
	CreateRouterAccessList6(payload *models.RouterAccessList6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterAccessList6(mkey string, params *models.CmdbRequestParams) (*models.RouterAccessList6, error)
	UpdateRouterAccessList6(mkey string, payload *models.RouterAccessList6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterAccessList6(mkey string, params *models.CmdbRequestParams) error
	CreateRouterAspathList(payload *models.RouterAspathList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterAspathList(mkey string, params *models.CmdbRequestParams) (*models.RouterAspathList, error)
	UpdateRouterAspathList(mkey string, payload *models.RouterAspathList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterAspathList(mkey string, params *models.CmdbRequestParams) error
	CreateRouterAuthPath(payload *models.RouterAuthPath, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterAuthPath(mkey string, params *models.CmdbRequestParams) (*models.RouterAuthPath, error)
	UpdateRouterAuthPath(mkey string, payload *models.RouterAuthPath, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterAuthPath(mkey string, params *models.CmdbRequestParams) error
	CreateRouterBfd(payload *models.RouterBfd, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterBfd(mkey string, params *models.CmdbRequestParams) (*models.RouterBfd, error)
	UpdateRouterBfd(mkey string, payload *models.RouterBfd, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterBfd(mkey string, params *models.CmdbRequestParams) error
	CreateRouterBfd6(payload *models.RouterBfd6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterBfd6(mkey string, params *models.CmdbRequestParams) (*models.RouterBfd6, error)
	UpdateRouterBfd6(mkey string, payload *models.RouterBfd6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterBfd6(mkey string, params *models.CmdbRequestParams) error
	CreateRouterBgp(payload *models.RouterBgp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterBgp(mkey string, params *models.CmdbRequestParams) (*models.RouterBgp, error)
	UpdateRouterBgp(mkey string, payload *models.RouterBgp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterBgp(mkey string, params *models.CmdbRequestParams) error
	CreateRouterCommunityList(payload *models.RouterCommunityList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterCommunityList(mkey string, params *models.CmdbRequestParams) (*models.RouterCommunityList, error)
	UpdateRouterCommunityList(mkey string, payload *models.RouterCommunityList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterCommunityList(mkey string, params *models.CmdbRequestParams) error
	CreateRouterIsis(payload *models.RouterIsis, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterIsis(mkey string, params *models.CmdbRequestParams) (*models.RouterIsis, error)
	UpdateRouterIsis(mkey string, payload *models.RouterIsis, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterIsis(mkey string, params *models.CmdbRequestParams) error
	CreateRouterKeyChain(payload *models.RouterKeyChain, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterKeyChain(mkey string, params *models.CmdbRequestParams) (*models.RouterKeyChain, error)
	UpdateRouterKeyChain(mkey string, payload *models.RouterKeyChain, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterKeyChain(mkey string, params *models.CmdbRequestParams) error
	CreateRouterMulticast(payload *models.RouterMulticast, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterMulticast(mkey string, params *models.CmdbRequestParams) (*models.RouterMulticast, error)
	UpdateRouterMulticast(mkey string, payload *models.RouterMulticast, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterMulticast(mkey string, params *models.CmdbRequestParams) error
	CreateRouterMulticast6(payload *models.RouterMulticast6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterMulticast6(mkey string, params *models.CmdbRequestParams) (*models.RouterMulticast6, error)
	UpdateRouterMulticast6(mkey string, payload *models.RouterMulticast6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterMulticast6(mkey string, params *models.CmdbRequestParams) error
	CreateRouterMulticastFlow(payload *models.RouterMulticastFlow, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterMulticastFlow(mkey string, params *models.CmdbRequestParams) (*models.RouterMulticastFlow, error)
	UpdateRouterMulticastFlow(mkey string, payload *models.RouterMulticastFlow, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterMulticastFlow(mkey string, params *models.CmdbRequestParams) error
	CreateRouterOspf(payload *models.RouterOspf, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterOspf(mkey string, params *models.CmdbRequestParams) (*models.RouterOspf, error)
	UpdateRouterOspf(mkey string, payload *models.RouterOspf, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterOspf(mkey string, params *models.CmdbRequestParams) error
	CreateRouterOspf6(payload *models.RouterOspf6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterOspf6(mkey string, params *models.CmdbRequestParams) (*models.RouterOspf6, error)
	UpdateRouterOspf6(mkey string, payload *models.RouterOspf6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterOspf6(mkey string, params *models.CmdbRequestParams) error
	CreateRouterPolicy(payload *models.RouterPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterPolicy(mkey string, params *models.CmdbRequestParams) (*models.RouterPolicy, error)
	UpdateRouterPolicy(mkey string, payload *models.RouterPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterPolicy(mkey string, params *models.CmdbRequestParams) error
	CreateRouterPolicy6(payload *models.RouterPolicy6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterPolicy6(mkey string, params *models.CmdbRequestParams) (*models.RouterPolicy6, error)
	UpdateRouterPolicy6(mkey string, payload *models.RouterPolicy6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterPolicy6(mkey string, params *models.CmdbRequestParams) error
	CreateRouterPrefixList(payload *models.RouterPrefixList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterPrefixList(mkey string, params *models.CmdbRequestParams) (*models.RouterPrefixList, error)
	UpdateRouterPrefixList(mkey string, payload *models.RouterPrefixList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterPrefixList(mkey string, params *models.CmdbRequestParams) error
	CreateRouterPrefixList6(payload *models.RouterPrefixList6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterPrefixList6(mkey string, params *models.CmdbRequestParams) (*models.RouterPrefixList6, error)
	UpdateRouterPrefixList6(mkey string, payload *models.RouterPrefixList6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterPrefixList6(mkey string, params *models.CmdbRequestParams) error
	CreateRouterRip(payload *models.RouterRip, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterRip(mkey string, params *models.CmdbRequestParams) (*models.RouterRip, error)
	UpdateRouterRip(mkey string, payload *models.RouterRip, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterRip(mkey string, params *models.CmdbRequestParams) error
	CreateRouterRipng(payload *models.RouterRipng, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterRipng(mkey string, params *models.CmdbRequestParams) (*models.RouterRipng, error)
	UpdateRouterRipng(mkey string, payload *models.RouterRipng, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterRipng(mkey string, params *models.CmdbRequestParams) error
	CreateRouterRouteMap(payload *models.RouterRouteMap, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterRouteMap(mkey string, params *models.CmdbRequestParams) (*models.RouterRouteMap, error)
	UpdateRouterRouteMap(mkey string, payload *models.RouterRouteMap, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterRouteMap(mkey string, params *models.CmdbRequestParams) error
	CreateRouterSetting(payload *models.RouterSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterSetting(mkey string, params *models.CmdbRequestParams) (*models.RouterSetting, error)
	UpdateRouterSetting(mkey string, payload *models.RouterSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterSetting(mkey string, params *models.CmdbRequestParams) error
	CreateRouterStatic(payload *models.RouterStatic, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterStatic(mkey string, params *models.CmdbRequestParams) (*models.RouterStatic, error)
	UpdateRouterStatic(mkey string, payload *models.RouterStatic, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterStatic(mkey string, params *models.CmdbRequestParams) error
	CreateRouterStatic6(payload *models.RouterStatic6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterStatic6(mkey string, params *models.CmdbRequestParams) (*models.RouterStatic6, error)
	UpdateRouterStatic6(mkey string, payload *models.RouterStatic6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterStatic6(mkey string, params *models.CmdbRequestParams) error
	CreateSctpFilterProfile(payload *models.SctpFilterProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSctpFilterProfile(mkey string, params *models.CmdbRequestParams) (*models.SctpFilterProfile, error)
	UpdateSctpFilterProfile(mkey string, payload *models.SctpFilterProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSctpFilterProfile(mkey string, params *models.CmdbRequestParams) error
	CreateSshFilterProfile(payload *models.SshFilterProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSshFilterProfile(mkey string, params *models.CmdbRequestParams) (*models.SshFilterProfile, error)
	UpdateSshFilterProfile(mkey string, payload *models.SshFilterProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSshFilterProfile(mkey string, params *models.CmdbRequestParams) error
	CreateSwitchControllerautoConfigDefault(payload *models.SwitchControllerautoConfigDefault, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerautoConfigDefault(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerautoConfigDefault, error)
	UpdateSwitchControllerautoConfigDefault(mkey string, payload *models.SwitchControllerautoConfigDefault, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerautoConfigDefault(mkey string, params *models.CmdbRequestParams) error
	CreateSwitchControllerinitialConfigTemplate(payload *models.SwitchControllerinitialConfigTemplate, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerinitialConfigTemplate(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerinitialConfigTemplate, error)
	UpdateSwitchControllerinitialConfigTemplate(mkey string, payload *models.SwitchControllerinitialConfigTemplate, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerinitialConfigTemplate(mkey string, params *models.CmdbRequestParams) error
	CreateSwitchControllerinitialConfigVlans(payload *models.SwitchControllerinitialConfigVlans, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerinitialConfigVlans(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerinitialConfigVlans, error)
	UpdateSwitchControllerinitialConfigVlans(mkey string, payload *models.SwitchControllerinitialConfigVlans, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerinitialConfigVlans(mkey string, params *models.CmdbRequestParams) error
	CreateSwitchControllersecurityPolicy8021X(payload *models.SwitchControllersecurityPolicy8021X, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllersecurityPolicy8021X(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllersecurityPolicy8021X, error)
	UpdateSwitchControllersecurityPolicy8021X(mkey string, payload *models.SwitchControllersecurityPolicy8021X, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllersecurityPolicy8021X(mkey string, params *models.CmdbRequestParams) error
	CreateSwitchControllerDynamicPortPolicy(payload *models.SwitchControllerDynamicPortPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerDynamicPortPolicy(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerDynamicPortPolicy, error)
	UpdateSwitchControllerDynamicPortPolicy(mkey string, payload *models.SwitchControllerDynamicPortPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerDynamicPortPolicy(mkey string, params *models.CmdbRequestParams) error
	CreateSwitchControllerFortilinkSettings(payload *models.SwitchControllerFortilinkSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerFortilinkSettings(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerFortilinkSettings, error)
	UpdateSwitchControllerFortilinkSettings(mkey string, payload *models.SwitchControllerFortilinkSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerFortilinkSettings(mkey string, params *models.CmdbRequestParams) error
	CreateSwitchControllerGlobal(payload *models.SwitchControllerGlobal, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerGlobal(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerGlobal, error)
	UpdateSwitchControllerGlobal(mkey string, payload *models.SwitchControllerGlobal, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerGlobal(mkey string, params *models.CmdbRequestParams) error
	CreateSwitchControllerLldpProfile(payload *models.SwitchControllerLldpProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerLldpProfile(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerLldpProfile, error)
	UpdateSwitchControllerLldpProfile(mkey string, payload *models.SwitchControllerLldpProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerLldpProfile(mkey string, params *models.CmdbRequestParams) error
	CreateSwitchControllerLldpSettings(payload *models.SwitchControllerLldpSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerLldpSettings(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerLldpSettings, error)
	UpdateSwitchControllerLldpSettings(mkey string, payload *models.SwitchControllerLldpSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerLldpSettings(mkey string, params *models.CmdbRequestParams) error
	CreateSwitchControllerLocation(payload *models.SwitchControllerLocation, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerLocation(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerLocation, error)
	UpdateSwitchControllerLocation(mkey string, payload *models.SwitchControllerLocation, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerLocation(mkey string, params *models.CmdbRequestParams) error
	CreateSwitchControllerMacPolicy(payload *models.SwitchControllerMacPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerMacPolicy(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerMacPolicy, error)
	UpdateSwitchControllerMacPolicy(mkey string, payload *models.SwitchControllerMacPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerMacPolicy(mkey string, params *models.CmdbRequestParams) error
	CreateSwitchControllerManagedSwitch(payload *models.SwitchControllerManagedSwitch, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerManagedSwitch(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerManagedSwitch, error)
	UpdateSwitchControllerManagedSwitch(mkey string, payload *models.SwitchControllerManagedSwitch, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerManagedSwitch(mkey string, params *models.CmdbRequestParams) error
	CreateSwitchControllerNacDevice(payload *models.SwitchControllerNacDevice, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerNacDevice(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerNacDevice, error)
	UpdateSwitchControllerNacDevice(mkey string, payload *models.SwitchControllerNacDevice, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerNacDevice(mkey string, params *models.CmdbRequestParams) error
	CreateSwitchControllerNacSettings(payload *models.SwitchControllerNacSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerNacSettings(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerNacSettings, error)
	UpdateSwitchControllerNacSettings(mkey string, payload *models.SwitchControllerNacSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerNacSettings(mkey string, params *models.CmdbRequestParams) error
	CreateSwitchControllerPortPolicy(payload *models.SwitchControllerPortPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerPortPolicy(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerPortPolicy, error)
	UpdateSwitchControllerPortPolicy(mkey string, payload *models.SwitchControllerPortPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerPortPolicy(mkey string, params *models.CmdbRequestParams) error
	CreateSwitchControllerSnmpCommunity(payload *models.SwitchControllerSnmpCommunity, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerSnmpCommunity(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerSnmpCommunity, error)
	UpdateSwitchControllerSnmpCommunity(mkey string, payload *models.SwitchControllerSnmpCommunity, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerSnmpCommunity(mkey string, params *models.CmdbRequestParams) error
	CreateSwitchControllerStpInstance(payload *models.SwitchControllerStpInstance, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerStpInstance(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerStpInstance, error)
	UpdateSwitchControllerStpInstance(mkey string, payload *models.SwitchControllerStpInstance, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerStpInstance(mkey string, params *models.CmdbRequestParams) error
	CreateSwitchControllerStpSettings(payload *models.SwitchControllerStpSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerStpSettings(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerStpSettings, error)
	UpdateSwitchControllerStpSettings(mkey string, payload *models.SwitchControllerStpSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerStpSettings(mkey string, params *models.CmdbRequestParams) error
	CreateSwitchControllerSwitchGroup(payload *models.SwitchControllerSwitchGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerSwitchGroup(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerSwitchGroup, error)
	UpdateSwitchControllerSwitchGroup(mkey string, payload *models.SwitchControllerSwitchGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerSwitchGroup(mkey string, params *models.CmdbRequestParams) error
	CreateSwitchControllerSystem(payload *models.SwitchControllerSystem, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerSystem(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerSystem, error)
	UpdateSwitchControllerSystem(mkey string, payload *models.SwitchControllerSystem, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerSystem(mkey string, params *models.CmdbRequestParams) error
	CreateSwitchControllerVlanPolicy(payload *models.SwitchControllerVlanPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSwitchControllerVlanPolicy(mkey string, params *models.CmdbRequestParams) (*models.SwitchControllerVlanPolicy, error)
	UpdateSwitchControllerVlanPolicy(mkey string, payload *models.SwitchControllerVlanPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSwitchControllerVlanPolicy(mkey string, params *models.CmdbRequestParams) error
	CreateSystem3gModemCustom(payload *models.System3gModemCustom, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystem3gModemCustom(mkey string, params *models.CmdbRequestParams) (*models.System3gModemCustom, error)
	UpdateSystem3gModemCustom(mkey string, payload *models.System3gModemCustom, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystem3gModemCustom(mkey string, params *models.CmdbRequestParams) error
	CreateSystemautoupdatePushUpdate(payload *models.SystemautoupdatePushUpdate, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemautoupdatePushUpdate(mkey string, params *models.CmdbRequestParams) (*models.SystemautoupdatePushUpdate, error)
	UpdateSystemautoupdatePushUpdate(mkey string, payload *models.SystemautoupdatePushUpdate, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemautoupdatePushUpdate(mkey string, params *models.CmdbRequestParams) error
	CreateSystemautoupdateSchedule(payload *models.SystemautoupdateSchedule, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemautoupdateSchedule(mkey string, params *models.CmdbRequestParams) (*models.SystemautoupdateSchedule, error)
	UpdateSystemautoupdateSchedule(mkey string, payload *models.SystemautoupdateSchedule, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemautoupdateSchedule(mkey string, params *models.CmdbRequestParams) error
	CreateSystemautoupdateTunneling(payload *models.SystemautoupdateTunneling, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemautoupdateTunneling(mkey string, params *models.CmdbRequestParams) (*models.SystemautoupdateTunneling, error)
	UpdateSystemautoupdateTunneling(mkey string, payload *models.SystemautoupdateTunneling, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemautoupdateTunneling(mkey string, params *models.CmdbRequestParams) error
	CreateSystemdhcp6Server(payload *models.Systemdhcp6Server, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemdhcp6Server(mkey string, params *models.CmdbRequestParams) (*models.Systemdhcp6Server, error)
	UpdateSystemdhcp6Server(mkey string, payload *models.Systemdhcp6Server, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemdhcp6Server(mkey string, params *models.CmdbRequestParams) error
	CreateSystemdhcpServer(payload *models.SystemdhcpServer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemdhcpServer(mkey string, params *models.CmdbRequestParams) (*models.SystemdhcpServer, error)
	UpdateSystemdhcpServer(mkey string, payload *models.SystemdhcpServer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemdhcpServer(mkey string, params *models.CmdbRequestParams) error
	CreateSystemlldpNetworkPolicy(payload *models.SystemlldpNetworkPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemlldpNetworkPolicy(mkey string, params *models.CmdbRequestParams) (*models.SystemlldpNetworkPolicy, error)
	UpdateSystemlldpNetworkPolicy(mkey string, payload *models.SystemlldpNetworkPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemlldpNetworkPolicy(mkey string, params *models.CmdbRequestParams) error
	CreateSystemreplacemsgAdmin(payload *models.SystemreplacemsgAdmin, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemreplacemsgAdmin(mkey string, params *models.CmdbRequestParams) (*models.SystemreplacemsgAdmin, error)
	UpdateSystemreplacemsgAdmin(mkey string, payload *models.SystemreplacemsgAdmin, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemreplacemsgAdmin(mkey string, params *models.CmdbRequestParams) error
	CreateSystemreplacemsgAlertmail(payload *models.SystemreplacemsgAlertmail, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemreplacemsgAlertmail(mkey string, params *models.CmdbRequestParams) (*models.SystemreplacemsgAlertmail, error)
	UpdateSystemreplacemsgAlertmail(mkey string, payload *models.SystemreplacemsgAlertmail, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemreplacemsgAlertmail(mkey string, params *models.CmdbRequestParams) error
	CreateSystemreplacemsgAuth(payload *models.SystemreplacemsgAuth, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemreplacemsgAuth(mkey string, params *models.CmdbRequestParams) (*models.SystemreplacemsgAuth, error)
	UpdateSystemreplacemsgAuth(mkey string, payload *models.SystemreplacemsgAuth, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemreplacemsgAuth(mkey string, params *models.CmdbRequestParams) error
	CreateSystemreplacemsgAutomation(payload *models.SystemreplacemsgAutomation, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemreplacemsgAutomation(mkey string, params *models.CmdbRequestParams) (*models.SystemreplacemsgAutomation, error)
	UpdateSystemreplacemsgAutomation(mkey string, payload *models.SystemreplacemsgAutomation, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemreplacemsgAutomation(mkey string, params *models.CmdbRequestParams) error
	CreateSystemreplacemsgDeviceDetectionPortal(payload *models.SystemreplacemsgDeviceDetectionPortal, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemreplacemsgDeviceDetectionPortal(mkey string, params *models.CmdbRequestParams) (*models.SystemreplacemsgDeviceDetectionPortal, error)
	UpdateSystemreplacemsgDeviceDetectionPortal(mkey string, payload *models.SystemreplacemsgDeviceDetectionPortal, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemreplacemsgDeviceDetectionPortal(mkey string, params *models.CmdbRequestParams) error
	CreateSystemreplacemsgFortiguardWf(payload *models.SystemreplacemsgFortiguardWf, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemreplacemsgFortiguardWf(mkey string, params *models.CmdbRequestParams) (*models.SystemreplacemsgFortiguardWf, error)
	UpdateSystemreplacemsgFortiguardWf(mkey string, payload *models.SystemreplacemsgFortiguardWf, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemreplacemsgFortiguardWf(mkey string, params *models.CmdbRequestParams) error
	CreateSystemreplacemsgFtp(payload *models.SystemreplacemsgFtp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemreplacemsgFtp(mkey string, params *models.CmdbRequestParams) (*models.SystemreplacemsgFtp, error)
	UpdateSystemreplacemsgFtp(mkey string, payload *models.SystemreplacemsgFtp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemreplacemsgFtp(mkey string, params *models.CmdbRequestParams) error
	CreateSystemreplacemsgHttp(payload *models.SystemreplacemsgHttp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemreplacemsgHttp(mkey string, params *models.CmdbRequestParams) (*models.SystemreplacemsgHttp, error)
	UpdateSystemreplacemsgHttp(mkey string, payload *models.SystemreplacemsgHttp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemreplacemsgHttp(mkey string, params *models.CmdbRequestParams) error
	CreateSystemreplacemsgIcap(payload *models.SystemreplacemsgIcap, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemreplacemsgIcap(mkey string, params *models.CmdbRequestParams) (*models.SystemreplacemsgIcap, error)
	UpdateSystemreplacemsgIcap(mkey string, payload *models.SystemreplacemsgIcap, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemreplacemsgIcap(mkey string, params *models.CmdbRequestParams) error
	CreateSystemreplacemsgMail(payload *models.SystemreplacemsgMail, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemreplacemsgMail(mkey string, params *models.CmdbRequestParams) (*models.SystemreplacemsgMail, error)
	UpdateSystemreplacemsgMail(mkey string, payload *models.SystemreplacemsgMail, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemreplacemsgMail(mkey string, params *models.CmdbRequestParams) error
	CreateSystemreplacemsgNacQuar(payload *models.SystemreplacemsgNacQuar, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemreplacemsgNacQuar(mkey string, params *models.CmdbRequestParams) (*models.SystemreplacemsgNacQuar, error)
	UpdateSystemreplacemsgNacQuar(mkey string, payload *models.SystemreplacemsgNacQuar, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemreplacemsgNacQuar(mkey string, params *models.CmdbRequestParams) error
	CreateSystemreplacemsgNntp(payload *models.SystemreplacemsgNntp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemreplacemsgNntp(mkey string, params *models.CmdbRequestParams) (*models.SystemreplacemsgNntp, error)
	UpdateSystemreplacemsgNntp(mkey string, payload *models.SystemreplacemsgNntp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemreplacemsgNntp(mkey string, params *models.CmdbRequestParams) error
	CreateSystemreplacemsgSpam(payload *models.SystemreplacemsgSpam, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemreplacemsgSpam(mkey string, params *models.CmdbRequestParams) (*models.SystemreplacemsgSpam, error)
	UpdateSystemreplacemsgSpam(mkey string, payload *models.SystemreplacemsgSpam, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemreplacemsgSpam(mkey string, params *models.CmdbRequestParams) error
	CreateSystemreplacemsgSslvpn(payload *models.SystemreplacemsgSslvpn, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemreplacemsgSslvpn(mkey string, params *models.CmdbRequestParams) (*models.SystemreplacemsgSslvpn, error)
	UpdateSystemreplacemsgSslvpn(mkey string, payload *models.SystemreplacemsgSslvpn, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemreplacemsgSslvpn(mkey string, params *models.CmdbRequestParams) error
	CreateSystemreplacemsgTrafficQuota(payload *models.SystemreplacemsgTrafficQuota, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemreplacemsgTrafficQuota(mkey string, params *models.CmdbRequestParams) (*models.SystemreplacemsgTrafficQuota, error)
	UpdateSystemreplacemsgTrafficQuota(mkey string, payload *models.SystemreplacemsgTrafficQuota, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemreplacemsgTrafficQuota(mkey string, params *models.CmdbRequestParams) error
	CreateSystemreplacemsgUtm(payload *models.SystemreplacemsgUtm, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemreplacemsgUtm(mkey string, params *models.CmdbRequestParams) (*models.SystemreplacemsgUtm, error)
	UpdateSystemreplacemsgUtm(mkey string, payload *models.SystemreplacemsgUtm, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemreplacemsgUtm(mkey string, params *models.CmdbRequestParams) error
	CreateSystemreplacemsgWebproxy(payload *models.SystemreplacemsgWebproxy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemreplacemsgWebproxy(mkey string, params *models.CmdbRequestParams) (*models.SystemreplacemsgWebproxy, error)
	UpdateSystemreplacemsgWebproxy(mkey string, payload *models.SystemreplacemsgWebproxy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemreplacemsgWebproxy(mkey string, params *models.CmdbRequestParams) error
	CreateSystemsdwanDuplication(payload *models.SystemsdwanDuplication, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemsdwanDuplication(mkey string, params *models.CmdbRequestParams) (*models.SystemsdwanDuplication, error)
	UpdateSystemsdwanDuplication(mkey string, payload *models.SystemsdwanDuplication, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemsdwanDuplication(mkey string, params *models.CmdbRequestParams) error
	CreateSystemsdwanHealthCheck(payload *models.SystemsdwanHealthCheck, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemsdwanHealthCheck(mkey string, params *models.CmdbRequestParams) (*models.SystemsdwanHealthCheck, error)
	UpdateSystemsdwanHealthCheck(mkey string, payload *models.SystemsdwanHealthCheck, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemsdwanHealthCheck(mkey string, params *models.CmdbRequestParams) error
	CreateSystemsdwanMembers(payload *models.SystemsdwanMembers, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemsdwanMembers(mkey string, params *models.CmdbRequestParams) (*models.SystemsdwanMembers, error)
	UpdateSystemsdwanMembers(mkey string, payload *models.SystemsdwanMembers, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemsdwanMembers(mkey string, params *models.CmdbRequestParams) error
	CreateSystemsdwanNeighbor(payload *models.SystemsdwanNeighbor, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemsdwanNeighbor(mkey string, params *models.CmdbRequestParams) (*models.SystemsdwanNeighbor, error)
	UpdateSystemsdwanNeighbor(mkey string, payload *models.SystemsdwanNeighbor, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemsdwanNeighbor(mkey string, params *models.CmdbRequestParams) error
	CreateSystemsdwanService(payload *models.SystemsdwanService, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemsdwanService(mkey string, params *models.CmdbRequestParams) (*models.SystemsdwanService, error)
	UpdateSystemsdwanService(mkey string, payload *models.SystemsdwanService, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemsdwanService(mkey string, params *models.CmdbRequestParams) error
	CreateSystemsdwanZone(payload *models.SystemsdwanZone, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemsdwanZone(mkey string, params *models.CmdbRequestParams) (*models.SystemsdwanZone, error)
	UpdateSystemsdwanZone(mkey string, payload *models.SystemsdwanZone, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemsdwanZone(mkey string, params *models.CmdbRequestParams) error
	CreateSystemsnmpCommunity(payload *models.SystemsnmpCommunity, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemsnmpCommunity(mkey string, params *models.CmdbRequestParams) (*models.SystemsnmpCommunity, error)
	UpdateSystemsnmpCommunity(mkey string, payload *models.SystemsnmpCommunity, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemsnmpCommunity(mkey string, params *models.CmdbRequestParams) error
	CreateSystemsnmpSysinfo(payload *models.SystemsnmpSysinfo, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemsnmpSysinfo(mkey string, params *models.CmdbRequestParams) (*models.SystemsnmpSysinfo, error)
	UpdateSystemsnmpSysinfo(mkey string, payload *models.SystemsnmpSysinfo, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemsnmpSysinfo(mkey string, params *models.CmdbRequestParams) error
	CreateSystemsnmpUser(payload *models.SystemsnmpUser, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemsnmpUser(mkey string, params *models.CmdbRequestParams) (*models.SystemsnmpUser, error)
	UpdateSystemsnmpUser(mkey string, payload *models.SystemsnmpUser, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemsnmpUser(mkey string, params *models.CmdbRequestParams) error
	CreateSystemAccprofile(payload *models.SystemAccprofile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemAccprofile(mkey string, params *models.CmdbRequestParams) (*models.SystemAccprofile, error)
	UpdateSystemAccprofile(mkey string, payload *models.SystemAccprofile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemAccprofile(mkey string, params *models.CmdbRequestParams) error
	CreateSystemAcme(payload *models.SystemAcme, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemAcme(mkey string, params *models.CmdbRequestParams) (*models.SystemAcme, error)
	UpdateSystemAcme(mkey string, payload *models.SystemAcme, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemAcme(mkey string, params *models.CmdbRequestParams) error
	CreateSystemAdmin(payload *models.SystemAdmin, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemAdmin(mkey string, params *models.CmdbRequestParams) (*models.SystemAdmin, error)
	UpdateSystemAdmin(mkey string, payload *models.SystemAdmin, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemAdmin(mkey string, params *models.CmdbRequestParams) error
	CreateSystemAlarm(payload *models.SystemAlarm, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemAlarm(mkey string, params *models.CmdbRequestParams) (*models.SystemAlarm, error)
	UpdateSystemAlarm(mkey string, payload *models.SystemAlarm, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemAlarm(mkey string, params *models.CmdbRequestParams) error
	CreateSystemApiUser(payload *models.SystemApiUser, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemApiUser(mkey string, params *models.CmdbRequestParams) (*models.SystemApiUser, error)
	UpdateSystemApiUser(mkey string, payload *models.SystemApiUser, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemApiUser(mkey string, params *models.CmdbRequestParams) error
	CreateSystemArpTable(payload *models.SystemArpTable, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemArpTable(mkey string, params *models.CmdbRequestParams) (*models.SystemArpTable, error)
	UpdateSystemArpTable(mkey string, payload *models.SystemArpTable, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemArpTable(mkey string, params *models.CmdbRequestParams) error
	CreateSystemAutoInstall(payload *models.SystemAutoInstall, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemAutoInstall(mkey string, params *models.CmdbRequestParams) (*models.SystemAutoInstall, error)
	UpdateSystemAutoInstall(mkey string, payload *models.SystemAutoInstall, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemAutoInstall(mkey string, params *models.CmdbRequestParams) error
	CreateSystemAutoScript(payload *models.SystemAutoScript, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemAutoScript(mkey string, params *models.CmdbRequestParams) (*models.SystemAutoScript, error)
	UpdateSystemAutoScript(mkey string, payload *models.SystemAutoScript, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemAutoScript(mkey string, params *models.CmdbRequestParams) error
	CreateSystemAutomationAction(payload *models.SystemAutomationAction, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemAutomationAction(mkey string, params *models.CmdbRequestParams) (*models.SystemAutomationAction, error)
	UpdateSystemAutomationAction(mkey string, payload *models.SystemAutomationAction, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemAutomationAction(mkey string, params *models.CmdbRequestParams) error
	CreateSystemAutomationDestination(payload *models.SystemAutomationDestination, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemAutomationDestination(mkey string, params *models.CmdbRequestParams) (*models.SystemAutomationDestination, error)
	UpdateSystemAutomationDestination(mkey string, payload *models.SystemAutomationDestination, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemAutomationDestination(mkey string, params *models.CmdbRequestParams) error
	CreateSystemAutomationStitch(payload *models.SystemAutomationStitch, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemAutomationStitch(mkey string, params *models.CmdbRequestParams) (*models.SystemAutomationStitch, error)
	UpdateSystemAutomationStitch(mkey string, payload *models.SystemAutomationStitch, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemAutomationStitch(mkey string, params *models.CmdbRequestParams) error
	CreateSystemAutomationTrigger(payload *models.SystemAutomationTrigger, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemAutomationTrigger(mkey string, params *models.CmdbRequestParams) (*models.SystemAutomationTrigger, error)
	UpdateSystemAutomationTrigger(mkey string, payload *models.SystemAutomationTrigger, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemAutomationTrigger(mkey string, params *models.CmdbRequestParams) error
	CreateSystemCentralManagement(payload *models.SystemCentralManagement, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemCentralManagement(mkey string, params *models.CmdbRequestParams) (*models.SystemCentralManagement, error)
	UpdateSystemCentralManagement(mkey string, payload *models.SystemCentralManagement, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemCentralManagement(mkey string, params *models.CmdbRequestParams) error
	CreateSystemClusterSync(payload *models.SystemClusterSync, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemClusterSync(mkey string, params *models.CmdbRequestParams) (*models.SystemClusterSync, error)
	UpdateSystemClusterSync(mkey string, payload *models.SystemClusterSync, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemClusterSync(mkey string, params *models.CmdbRequestParams) error
	CreateSystemConsole(payload *models.SystemConsole, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemConsole(mkey string, params *models.CmdbRequestParams) (*models.SystemConsole, error)
	UpdateSystemConsole(mkey string, payload *models.SystemConsole, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemConsole(mkey string, params *models.CmdbRequestParams) error
	CreateSystemCsf(payload *models.SystemCsf, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemCsf(mkey string, params *models.CmdbRequestParams) (*models.SystemCsf, error)
	UpdateSystemCsf(mkey string, payload *models.SystemCsf, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemCsf(mkey string, params *models.CmdbRequestParams) error
	CreateSystemCustomLanguage(payload *models.SystemCustomLanguage, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemCustomLanguage(mkey string, params *models.CmdbRequestParams) (*models.SystemCustomLanguage, error)
	UpdateSystemCustomLanguage(mkey string, payload *models.SystemCustomLanguage, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemCustomLanguage(mkey string, params *models.CmdbRequestParams) error
	CreateSystemDdns(payload *models.SystemDdns, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemDdns(mkey string, params *models.CmdbRequestParams) (*models.SystemDdns, error)
	UpdateSystemDdns(mkey string, payload *models.SystemDdns, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemDdns(mkey string, params *models.CmdbRequestParams) error
	CreateSystemDedicatedMgmt(payload *models.SystemDedicatedMgmt, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemDedicatedMgmt(mkey string, params *models.CmdbRequestParams) (*models.SystemDedicatedMgmt, error)
	UpdateSystemDedicatedMgmt(mkey string, payload *models.SystemDedicatedMgmt, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemDedicatedMgmt(mkey string, params *models.CmdbRequestParams) error
	CreateSystemDns(payload *models.SystemDns, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemDns(mkey string, params *models.CmdbRequestParams) (*models.SystemDns, error)
	UpdateSystemDns(mkey string, payload *models.SystemDns, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemDns(mkey string, params *models.CmdbRequestParams) error
	CreateSystemDns64(payload *models.SystemDns64, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemDns64(mkey string, params *models.CmdbRequestParams) (*models.SystemDns64, error)
	UpdateSystemDns64(mkey string, payload *models.SystemDns64, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemDns64(mkey string, params *models.CmdbRequestParams) error
	CreateSystemDnsDatabase(payload *models.SystemDnsDatabase, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemDnsDatabase(mkey string, params *models.CmdbRequestParams) (*models.SystemDnsDatabase, error)
	UpdateSystemDnsDatabase(mkey string, payload *models.SystemDnsDatabase, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemDnsDatabase(mkey string, params *models.CmdbRequestParams) error
	CreateSystemDnsServer(payload *models.SystemDnsServer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemDnsServer(mkey string, params *models.CmdbRequestParams) (*models.SystemDnsServer, error)
	UpdateSystemDnsServer(mkey string, payload *models.SystemDnsServer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemDnsServer(mkey string, params *models.CmdbRequestParams) error
	CreateSystemDscpBasedPriority(payload *models.SystemDscpBasedPriority, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemDscpBasedPriority(mkey string, params *models.CmdbRequestParams) (*models.SystemDscpBasedPriority, error)
	UpdateSystemDscpBasedPriority(mkey string, payload *models.SystemDscpBasedPriority, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemDscpBasedPriority(mkey string, params *models.CmdbRequestParams) error
	CreateSystemEmailServer(payload *models.SystemEmailServer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemEmailServer(mkey string, params *models.CmdbRequestParams) (*models.SystemEmailServer, error)
	UpdateSystemEmailServer(mkey string, payload *models.SystemEmailServer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemEmailServer(mkey string, params *models.CmdbRequestParams) error
	CreateSystemExternalResource(payload *models.SystemExternalResource, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemExternalResource(mkey string, params *models.CmdbRequestParams) (*models.SystemExternalResource, error)
	UpdateSystemExternalResource(mkey string, payload *models.SystemExternalResource, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemExternalResource(mkey string, params *models.CmdbRequestParams) error
	CreateSystemFederatedUpgrade(payload *models.SystemFederatedUpgrade, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemFederatedUpgrade(mkey string, params *models.CmdbRequestParams) (*models.SystemFederatedUpgrade, error)
	UpdateSystemFederatedUpgrade(mkey string, payload *models.SystemFederatedUpgrade, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemFederatedUpgrade(mkey string, params *models.CmdbRequestParams) error
	CreateSystemFipsCc(payload *models.SystemFipsCc, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemFipsCc(mkey string, params *models.CmdbRequestParams) (*models.SystemFipsCc, error)
	UpdateSystemFipsCc(mkey string, payload *models.SystemFipsCc, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemFipsCc(mkey string, params *models.CmdbRequestParams) error
	CreateSystemFortiai(payload *models.SystemFortiai, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemFortiai(mkey string, params *models.CmdbRequestParams) (*models.SystemFortiai, error)
	UpdateSystemFortiai(mkey string, payload *models.SystemFortiai, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemFortiai(mkey string, params *models.CmdbRequestParams) error
	CreateSystemFortiguard(payload *models.SystemFortiguard, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemFortiguard(mkey string, params *models.CmdbRequestParams) (*models.SystemFortiguard, error)
	UpdateSystemFortiguard(mkey string, payload *models.SystemFortiguard, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemFortiguard(mkey string, params *models.CmdbRequestParams) error
	CreateSystemFortimanager(payload *models.SystemFortimanager, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemFortimanager(mkey string, params *models.CmdbRequestParams) (*models.SystemFortimanager, error)
	UpdateSystemFortimanager(mkey string, payload *models.SystemFortimanager, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemFortimanager(mkey string, params *models.CmdbRequestParams) error
	CreateSystemFortisandbox(payload *models.SystemFortisandbox, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemFortisandbox(mkey string, params *models.CmdbRequestParams) (*models.SystemFortisandbox, error)
	UpdateSystemFortisandbox(mkey string, payload *models.SystemFortisandbox, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemFortisandbox(mkey string, params *models.CmdbRequestParams) error
	CreateSystemFssoPolling(payload *models.SystemFssoPolling, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemFssoPolling(mkey string, params *models.CmdbRequestParams) (*models.SystemFssoPolling, error)
	UpdateSystemFssoPolling(mkey string, payload *models.SystemFssoPolling, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemFssoPolling(mkey string, params *models.CmdbRequestParams) error
	CreateSystemFtmPush(payload *models.SystemFtmPush, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemFtmPush(mkey string, params *models.CmdbRequestParams) (*models.SystemFtmPush, error)
	UpdateSystemFtmPush(mkey string, payload *models.SystemFtmPush, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemFtmPush(mkey string, params *models.CmdbRequestParams) error
	CreateSystemGeneve(payload *models.SystemGeneve, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemGeneve(mkey string, params *models.CmdbRequestParams) (*models.SystemGeneve, error)
	UpdateSystemGeneve(mkey string, payload *models.SystemGeneve, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemGeneve(mkey string, params *models.CmdbRequestParams) error
	CreateSystemGeoipCountry(payload *models.SystemGeoipCountry, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemGeoipCountry(mkey string, params *models.CmdbRequestParams) (*models.SystemGeoipCountry, error)
	UpdateSystemGeoipCountry(mkey string, payload *models.SystemGeoipCountry, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemGeoipCountry(mkey string, params *models.CmdbRequestParams) error
	CreateSystemGeoipOverride(payload *models.SystemGeoipOverride, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemGeoipOverride(mkey string, params *models.CmdbRequestParams) (*models.SystemGeoipOverride, error)
	UpdateSystemGeoipOverride(mkey string, payload *models.SystemGeoipOverride, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemGeoipOverride(mkey string, params *models.CmdbRequestParams) error
	CreateSystemGlobal(payload *models.SystemGlobal, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemGlobal(mkey string, params *models.CmdbRequestParams) (*models.SystemGlobal, error)
	UpdateSystemGlobal(mkey string, payload *models.SystemGlobal, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemGlobal(mkey string, params *models.CmdbRequestParams) error
	CreateSystemGreTunnel(payload *models.SystemGreTunnel, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemGreTunnel(mkey string, params *models.CmdbRequestParams) (*models.SystemGreTunnel, error)
	UpdateSystemGreTunnel(mkey string, payload *models.SystemGreTunnel, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemGreTunnel(mkey string, params *models.CmdbRequestParams) error
	CreateSystemHa(payload *models.SystemHa, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemHa(mkey string, params *models.CmdbRequestParams) (*models.SystemHa, error)
	UpdateSystemHa(mkey string, payload *models.SystemHa, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemHa(mkey string, params *models.CmdbRequestParams) error
	CreateSystemHaMonitor(payload *models.SystemHaMonitor, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemHaMonitor(mkey string, params *models.CmdbRequestParams) (*models.SystemHaMonitor, error)
	UpdateSystemHaMonitor(mkey string, payload *models.SystemHaMonitor, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemHaMonitor(mkey string, params *models.CmdbRequestParams) error
	CreateSystemIke(payload *models.SystemIke, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemIke(mkey string, params *models.CmdbRequestParams) (*models.SystemIke, error)
	UpdateSystemIke(mkey string, payload *models.SystemIke, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemIke(mkey string, params *models.CmdbRequestParams) error
	CreateSystemInterface(payload *models.SystemInterface, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemInterface(mkey string, params *models.CmdbRequestParams) (*models.SystemInterface, error)
	UpdateSystemInterface(mkey string, payload *models.SystemInterface, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemInterface(mkey string, params *models.CmdbRequestParams) error
	CreateSystemIpam(payload *models.SystemIpam, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemIpam(mkey string, params *models.CmdbRequestParams) (*models.SystemIpam, error)
	UpdateSystemIpam(mkey string, payload *models.SystemIpam, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemIpam(mkey string, params *models.CmdbRequestParams) error
	CreateSystemIpipTunnel(payload *models.SystemIpipTunnel, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemIpipTunnel(mkey string, params *models.CmdbRequestParams) (*models.SystemIpipTunnel, error)
	UpdateSystemIpipTunnel(mkey string, payload *models.SystemIpipTunnel, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemIpipTunnel(mkey string, params *models.CmdbRequestParams) error
	CreateSystemIps(payload *models.SystemIps, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemIps(mkey string, params *models.CmdbRequestParams) (*models.SystemIps, error)
	UpdateSystemIps(mkey string, payload *models.SystemIps, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemIps(mkey string, params *models.CmdbRequestParams) error
	CreateSystemIpsUrlfilterDns(payload *models.SystemIpsUrlfilterDns, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemIpsUrlfilterDns(mkey string, params *models.CmdbRequestParams) (*models.SystemIpsUrlfilterDns, error)
	UpdateSystemIpsUrlfilterDns(mkey string, payload *models.SystemIpsUrlfilterDns, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemIpsUrlfilterDns(mkey string, params *models.CmdbRequestParams) error
	CreateSystemIpsUrlfilterDns6(payload *models.SystemIpsUrlfilterDns6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemIpsUrlfilterDns6(mkey string, params *models.CmdbRequestParams) (*models.SystemIpsUrlfilterDns6, error)
	UpdateSystemIpsUrlfilterDns6(mkey string, payload *models.SystemIpsUrlfilterDns6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemIpsUrlfilterDns6(mkey string, params *models.CmdbRequestParams) error
	CreateSystemIpsecAggregate(payload *models.SystemIpsecAggregate, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemIpsecAggregate(mkey string, params *models.CmdbRequestParams) (*models.SystemIpsecAggregate, error)
	UpdateSystemIpsecAggregate(mkey string, payload *models.SystemIpsecAggregate, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemIpsecAggregate(mkey string, params *models.CmdbRequestParams) error
	CreateSystemIpv6NeighborCache(payload *models.SystemIpv6NeighborCache, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemIpv6NeighborCache(mkey string, params *models.CmdbRequestParams) (*models.SystemIpv6NeighborCache, error)
	UpdateSystemIpv6NeighborCache(mkey string, payload *models.SystemIpv6NeighborCache, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemIpv6NeighborCache(mkey string, params *models.CmdbRequestParams) error
	CreateSystemIpv6Tunnel(payload *models.SystemIpv6Tunnel, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemIpv6Tunnel(mkey string, params *models.CmdbRequestParams) (*models.SystemIpv6Tunnel, error)
	UpdateSystemIpv6Tunnel(mkey string, payload *models.SystemIpv6Tunnel, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemIpv6Tunnel(mkey string, params *models.CmdbRequestParams) error
	CreateSystemLinkMonitor(payload *models.SystemLinkMonitor, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemLinkMonitor(mkey string, params *models.CmdbRequestParams) (*models.SystemLinkMonitor, error)
	UpdateSystemLinkMonitor(mkey string, payload *models.SystemLinkMonitor, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemLinkMonitor(mkey string, params *models.CmdbRequestParams) error
	CreateSystemLteModem(payload *models.SystemLteModem, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemLteModem(mkey string, params *models.CmdbRequestParams) (*models.SystemLteModem, error)
	UpdateSystemLteModem(mkey string, payload *models.SystemLteModem, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemLteModem(mkey string, params *models.CmdbRequestParams) error
	CreateSystemMacAddressTable(payload *models.SystemMacAddressTable, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemMacAddressTable(mkey string, params *models.CmdbRequestParams) (*models.SystemMacAddressTable, error)
	UpdateSystemMacAddressTable(mkey string, payload *models.SystemMacAddressTable, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemMacAddressTable(mkey string, params *models.CmdbRequestParams) error
	CreateSystemMobileTunnel(payload *models.SystemMobileTunnel, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemMobileTunnel(mkey string, params *models.CmdbRequestParams) (*models.SystemMobileTunnel, error)
	UpdateSystemMobileTunnel(mkey string, payload *models.SystemMobileTunnel, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemMobileTunnel(mkey string, params *models.CmdbRequestParams) error
	CreateSystemModem(payload *models.SystemModem, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemModem(mkey string, params *models.CmdbRequestParams) (*models.SystemModem, error)
	UpdateSystemModem(mkey string, payload *models.SystemModem, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemModem(mkey string, params *models.CmdbRequestParams) error
	CreateSystemNat64(payload *models.SystemNat64, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemNat64(mkey string, params *models.CmdbRequestParams) (*models.SystemNat64, error)
	UpdateSystemNat64(mkey string, payload *models.SystemNat64, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemNat64(mkey string, params *models.CmdbRequestParams) error
	CreateSystemNdProxy(payload *models.SystemNdProxy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemNdProxy(mkey string, params *models.CmdbRequestParams) (*models.SystemNdProxy, error)
	UpdateSystemNdProxy(mkey string, payload *models.SystemNdProxy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemNdProxy(mkey string, params *models.CmdbRequestParams) error
	CreateSystemNetflow(payload *models.SystemNetflow, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemNetflow(mkey string, params *models.CmdbRequestParams) (*models.SystemNetflow, error)
	UpdateSystemNetflow(mkey string, payload *models.SystemNetflow, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemNetflow(mkey string, params *models.CmdbRequestParams) error
	CreateSystemNetworkVisibility(payload *models.SystemNetworkVisibility, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemNetworkVisibility(mkey string, params *models.CmdbRequestParams) (*models.SystemNetworkVisibility, error)
	UpdateSystemNetworkVisibility(mkey string, payload *models.SystemNetworkVisibility, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemNetworkVisibility(mkey string, params *models.CmdbRequestParams) error
	CreateSystemNpu(payload *models.SystemNpu, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemNpu(mkey string, params *models.CmdbRequestParams) (*models.SystemNpu, error)
	UpdateSystemNpu(mkey string, payload *models.SystemNpu, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemNpu(mkey string, params *models.CmdbRequestParams) error
	CreateSystemNtp(payload *models.SystemNtp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemNtp(mkey string, params *models.CmdbRequestParams) (*models.SystemNtp, error)
	UpdateSystemNtp(mkey string, payload *models.SystemNtp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemNtp(mkey string, params *models.CmdbRequestParams) error
	CreateSystemObjectTagging(payload *models.SystemObjectTagging, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemObjectTagging(mkey string, params *models.CmdbRequestParams) (*models.SystemObjectTagging, error)
	UpdateSystemObjectTagging(mkey string, payload *models.SystemObjectTagging, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemObjectTagging(mkey string, params *models.CmdbRequestParams) error
	CreateSystemPasswordPolicy(payload *models.SystemPasswordPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemPasswordPolicy(mkey string, params *models.CmdbRequestParams) (*models.SystemPasswordPolicy, error)
	UpdateSystemPasswordPolicy(mkey string, payload *models.SystemPasswordPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemPasswordPolicy(mkey string, params *models.CmdbRequestParams) error
	CreateSystemPasswordPolicyGuestAdmin(payload *models.SystemPasswordPolicyGuestAdmin, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemPasswordPolicyGuestAdmin(mkey string, params *models.CmdbRequestParams) (*models.SystemPasswordPolicyGuestAdmin, error)
	UpdateSystemPasswordPolicyGuestAdmin(mkey string, payload *models.SystemPasswordPolicyGuestAdmin, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemPasswordPolicyGuestAdmin(mkey string, params *models.CmdbRequestParams) error
	CreateSystemPhysicalSwitch(payload *models.SystemPhysicalSwitch, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemPhysicalSwitch(mkey string, params *models.CmdbRequestParams) (*models.SystemPhysicalSwitch, error)
	UpdateSystemPhysicalSwitch(mkey string, payload *models.SystemPhysicalSwitch, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemPhysicalSwitch(mkey string, params *models.CmdbRequestParams) error
	CreateSystemPppoeInterface(payload *models.SystemPppoeInterface, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemPppoeInterface(mkey string, params *models.CmdbRequestParams) (*models.SystemPppoeInterface, error)
	UpdateSystemPppoeInterface(mkey string, payload *models.SystemPppoeInterface, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemPppoeInterface(mkey string, params *models.CmdbRequestParams) error
	CreateSystemProbeResponse(payload *models.SystemProbeResponse, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemProbeResponse(mkey string, params *models.CmdbRequestParams) (*models.SystemProbeResponse, error)
	UpdateSystemProbeResponse(mkey string, payload *models.SystemProbeResponse, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemProbeResponse(mkey string, params *models.CmdbRequestParams) error
	CreateSystemProxyArp(payload *models.SystemProxyArp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemProxyArp(mkey string, params *models.CmdbRequestParams) (*models.SystemProxyArp, error)
	UpdateSystemProxyArp(mkey string, payload *models.SystemProxyArp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemProxyArp(mkey string, params *models.CmdbRequestParams) error
	CreateSystemPtp(payload *models.SystemPtp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemPtp(mkey string, params *models.CmdbRequestParams) (*models.SystemPtp, error)
	UpdateSystemPtp(mkey string, payload *models.SystemPtp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemPtp(mkey string, params *models.CmdbRequestParams) error
	CreateSystemReplacemsgGroup(payload *models.SystemReplacemsgGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemReplacemsgGroup(mkey string, params *models.CmdbRequestParams) (*models.SystemReplacemsgGroup, error)
	UpdateSystemReplacemsgGroup(mkey string, payload *models.SystemReplacemsgGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemReplacemsgGroup(mkey string, params *models.CmdbRequestParams) error
	CreateSystemReplacemsgImage(payload *models.SystemReplacemsgImage, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemReplacemsgImage(mkey string, params *models.CmdbRequestParams) (*models.SystemReplacemsgImage, error)
	UpdateSystemReplacemsgImage(mkey string, payload *models.SystemReplacemsgImage, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemReplacemsgImage(mkey string, params *models.CmdbRequestParams) error
	CreateSystemResourceLimits(payload *models.SystemResourceLimits, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemResourceLimits(mkey string, params *models.CmdbRequestParams) (*models.SystemResourceLimits, error)
	UpdateSystemResourceLimits(mkey string, payload *models.SystemResourceLimits, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemResourceLimits(mkey string, params *models.CmdbRequestParams) error
	CreateSystemSaml(payload *models.SystemSaml, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemSaml(mkey string, params *models.CmdbRequestParams) (*models.SystemSaml, error)
	UpdateSystemSaml(mkey string, payload *models.SystemSaml, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemSaml(mkey string, params *models.CmdbRequestParams) error
	CreateSystemSdnConnector(payload *models.SystemSdnConnector, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemSdnConnector(mkey string, params *models.CmdbRequestParams) (*models.SystemSdnConnector, error)
	UpdateSystemSdnConnector(mkey string, payload *models.SystemSdnConnector, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemSdnConnector(mkey string, params *models.CmdbRequestParams) error
	CreateSystemSdwan(payload *models.SystemSdwan, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemSdwan(mkey string, params *models.CmdbRequestParams) (*models.SystemSdwan, error)
	UpdateSystemSdwan(mkey string, payload *models.SystemSdwan, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemSdwan(mkey string, params *models.CmdbRequestParams) error
	CreateSystemSessionHelper(payload *models.SystemSessionHelper, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemSessionHelper(mkey string, params *models.CmdbRequestParams) (*models.SystemSessionHelper, error)
	UpdateSystemSessionHelper(mkey string, payload *models.SystemSessionHelper, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemSessionHelper(mkey string, params *models.CmdbRequestParams) error
	CreateSystemSessionTtl(payload *models.SystemSessionTtl, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemSessionTtl(mkey string, params *models.CmdbRequestParams) (*models.SystemSessionTtl, error)
	UpdateSystemSessionTtl(mkey string, payload *models.SystemSessionTtl, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemSessionTtl(mkey string, params *models.CmdbRequestParams) error
	CreateSystemSettings(payload *models.SystemSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemSettings(mkey string, params *models.CmdbRequestParams) (*models.SystemSettings, error)
	UpdateSystemSettings(mkey string, payload *models.SystemSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemSettings(mkey string, params *models.CmdbRequestParams) error
	CreateSystemSflow(payload *models.SystemSflow, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemSflow(mkey string, params *models.CmdbRequestParams) (*models.SystemSflow, error)
	UpdateSystemSflow(mkey string, payload *models.SystemSflow, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemSflow(mkey string, params *models.CmdbRequestParams) error
	CreateSystemSitTunnel(payload *models.SystemSitTunnel, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemSitTunnel(mkey string, params *models.CmdbRequestParams) (*models.SystemSitTunnel, error)
	UpdateSystemSitTunnel(mkey string, payload *models.SystemSitTunnel, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemSitTunnel(mkey string, params *models.CmdbRequestParams) error
	CreateSystemSmsServer(payload *models.SystemSmsServer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemSmsServer(mkey string, params *models.CmdbRequestParams) (*models.SystemSmsServer, error)
	UpdateSystemSmsServer(mkey string, payload *models.SystemSmsServer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemSmsServer(mkey string, params *models.CmdbRequestParams) error
	CreateSystemSpeedTestSchedule(payload *models.SystemSpeedTestSchedule, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemSpeedTestSchedule(mkey string, params *models.CmdbRequestParams) (*models.SystemSpeedTestSchedule, error)
	UpdateSystemSpeedTestSchedule(mkey string, payload *models.SystemSpeedTestSchedule, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemSpeedTestSchedule(mkey string, params *models.CmdbRequestParams) error
	CreateSystemSpeedTestServer(payload *models.SystemSpeedTestServer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemSpeedTestServer(mkey string, params *models.CmdbRequestParams) (*models.SystemSpeedTestServer, error)
	UpdateSystemSpeedTestServer(mkey string, payload *models.SystemSpeedTestServer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemSpeedTestServer(mkey string, params *models.CmdbRequestParams) error
	CreateSystemSsoAdmin(payload *models.SystemSsoAdmin, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemSsoAdmin(mkey string, params *models.CmdbRequestParams) (*models.SystemSsoAdmin, error)
	UpdateSystemSsoAdmin(mkey string, payload *models.SystemSsoAdmin, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemSsoAdmin(mkey string, params *models.CmdbRequestParams) error
	CreateSystemSsoForticloudAdmin(payload *models.SystemSsoForticloudAdmin, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemSsoForticloudAdmin(mkey string, params *models.CmdbRequestParams) (*models.SystemSsoForticloudAdmin, error)
	UpdateSystemSsoForticloudAdmin(mkey string, payload *models.SystemSsoForticloudAdmin, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemSsoForticloudAdmin(mkey string, params *models.CmdbRequestParams) error
	CreateSystemStandaloneCluster(payload *models.SystemStandaloneCluster, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemStandaloneCluster(mkey string, params *models.CmdbRequestParams) (*models.SystemStandaloneCluster, error)
	UpdateSystemStandaloneCluster(mkey string, payload *models.SystemStandaloneCluster, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemStandaloneCluster(mkey string, params *models.CmdbRequestParams) error
	CreateSystemStorage(payload *models.SystemStorage, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemStorage(mkey string, params *models.CmdbRequestParams) (*models.SystemStorage, error)
	UpdateSystemStorage(mkey string, payload *models.SystemStorage, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemStorage(mkey string, params *models.CmdbRequestParams) error
	CreateSystemStp(payload *models.SystemStp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemStp(mkey string, params *models.CmdbRequestParams) (*models.SystemStp, error)
	UpdateSystemStp(mkey string, payload *models.SystemStp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemStp(mkey string, params *models.CmdbRequestParams) error
	CreateSystemSwitchInterface(payload *models.SystemSwitchInterface, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemSwitchInterface(mkey string, params *models.CmdbRequestParams) (*models.SystemSwitchInterface, error)
	UpdateSystemSwitchInterface(mkey string, payload *models.SystemSwitchInterface, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemSwitchInterface(mkey string, params *models.CmdbRequestParams) error
	CreateSystemTosBasedPriority(payload *models.SystemTosBasedPriority, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemTosBasedPriority(mkey string, params *models.CmdbRequestParams) (*models.SystemTosBasedPriority, error)
	UpdateSystemTosBasedPriority(mkey string, payload *models.SystemTosBasedPriority, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemTosBasedPriority(mkey string, params *models.CmdbRequestParams) error
	CreateSystemVdom(payload *models.SystemVdom, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemVdom(mkey string, params *models.CmdbRequestParams) (*models.SystemVdom, error)
	UpdateSystemVdom(mkey string, payload *models.SystemVdom, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemVdom(mkey string, params *models.CmdbRequestParams) error
	CreateSystemVdomDns(payload *models.SystemVdomDns, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemVdomDns(mkey string, params *models.CmdbRequestParams) (*models.SystemVdomDns, error)
	UpdateSystemVdomDns(mkey string, payload *models.SystemVdomDns, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemVdomDns(mkey string, params *models.CmdbRequestParams) error
	CreateSystemVdomException(payload *models.SystemVdomException, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemVdomException(mkey string, params *models.CmdbRequestParams) (*models.SystemVdomException, error)
	UpdateSystemVdomException(mkey string, payload *models.SystemVdomException, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemVdomException(mkey string, params *models.CmdbRequestParams) error
	CreateSystemVdomLink(payload *models.SystemVdomLink, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemVdomLink(mkey string, params *models.CmdbRequestParams) (*models.SystemVdomLink, error)
	UpdateSystemVdomLink(mkey string, payload *models.SystemVdomLink, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemVdomLink(mkey string, params *models.CmdbRequestParams) error
	CreateSystemVdomNetflow(payload *models.SystemVdomNetflow, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemVdomNetflow(mkey string, params *models.CmdbRequestParams) (*models.SystemVdomNetflow, error)
	UpdateSystemVdomNetflow(mkey string, payload *models.SystemVdomNetflow, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemVdomNetflow(mkey string, params *models.CmdbRequestParams) error
	CreateSystemVdomProperty(payload *models.SystemVdomProperty, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemVdomProperty(mkey string, params *models.CmdbRequestParams) (*models.SystemVdomProperty, error)
	UpdateSystemVdomProperty(mkey string, payload *models.SystemVdomProperty, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemVdomProperty(mkey string, params *models.CmdbRequestParams) error
	CreateSystemVdomRadiusServer(payload *models.SystemVdomRadiusServer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemVdomRadiusServer(mkey string, params *models.CmdbRequestParams) (*models.SystemVdomRadiusServer, error)
	UpdateSystemVdomRadiusServer(mkey string, payload *models.SystemVdomRadiusServer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemVdomRadiusServer(mkey string, params *models.CmdbRequestParams) error
	CreateSystemVdomSflow(payload *models.SystemVdomSflow, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemVdomSflow(mkey string, params *models.CmdbRequestParams) (*models.SystemVdomSflow, error)
	UpdateSystemVdomSflow(mkey string, payload *models.SystemVdomSflow, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemVdomSflow(mkey string, params *models.CmdbRequestParams) error
	CreateSystemVirtualSwitch(payload *models.SystemVirtualSwitch, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemVirtualSwitch(mkey string, params *models.CmdbRequestParams) (*models.SystemVirtualSwitch, error)
	UpdateSystemVirtualSwitch(mkey string, payload *models.SystemVirtualSwitch, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemVirtualSwitch(mkey string, params *models.CmdbRequestParams) error
	CreateSystemVirtualWanLink(payload *models.SystemVirtualWanLink, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemVirtualWanLink(mkey string, params *models.CmdbRequestParams) (*models.SystemVirtualWanLink, error)
	UpdateSystemVirtualWanLink(mkey string, payload *models.SystemVirtualWanLink, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemVirtualWanLink(mkey string, params *models.CmdbRequestParams) error
	CreateSystemVirtualWirePair(payload *models.SystemVirtualWirePair, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemVirtualWirePair(mkey string, params *models.CmdbRequestParams) (*models.SystemVirtualWirePair, error)
	UpdateSystemVirtualWirePair(mkey string, payload *models.SystemVirtualWirePair, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemVirtualWirePair(mkey string, params *models.CmdbRequestParams) error
	CreateSystemVneTunnel(payload *models.SystemVneTunnel, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemVneTunnel(mkey string, params *models.CmdbRequestParams) (*models.SystemVneTunnel, error)
	UpdateSystemVneTunnel(mkey string, payload *models.SystemVneTunnel, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemVneTunnel(mkey string, params *models.CmdbRequestParams) error
	CreateSystemVxlan(payload *models.SystemVxlan, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemVxlan(mkey string, params *models.CmdbRequestParams) (*models.SystemVxlan, error)
	UpdateSystemVxlan(mkey string, payload *models.SystemVxlan, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemVxlan(mkey string, params *models.CmdbRequestParams) error
	CreateSystemWccp(payload *models.SystemWccp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemWccp(mkey string, params *models.CmdbRequestParams) (*models.SystemWccp, error)
	UpdateSystemWccp(mkey string, payload *models.SystemWccp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemWccp(mkey string, params *models.CmdbRequestParams) error
	CreateSystemZone(payload *models.SystemZone, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadSystemZone(mkey string, params *models.CmdbRequestParams) (*models.SystemZone, error)
	UpdateSystemZone(mkey string, payload *models.SystemZone, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteSystemZone(mkey string, params *models.CmdbRequestParams) error
	CreateUserAdgrp(payload *models.UserAdgrp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadUserAdgrp(mkey string, params *models.CmdbRequestParams) (*models.UserAdgrp, error)
	UpdateUserAdgrp(mkey string, payload *models.UserAdgrp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteUserAdgrp(mkey string, params *models.CmdbRequestParams) error
	CreateUserCertificate(payload *models.UserCertificate, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadUserCertificate(mkey string, params *models.CmdbRequestParams) (*models.UserCertificate, error)
	UpdateUserCertificate(mkey string, payload *models.UserCertificate, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteUserCertificate(mkey string, params *models.CmdbRequestParams) error
	CreateUserDomainController(payload *models.UserDomainController, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadUserDomainController(mkey string, params *models.CmdbRequestParams) (*models.UserDomainController, error)
	UpdateUserDomainController(mkey string, payload *models.UserDomainController, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteUserDomainController(mkey string, params *models.CmdbRequestParams) error
	CreateUserExchange(payload *models.UserExchange, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadUserExchange(mkey string, params *models.CmdbRequestParams) (*models.UserExchange, error)
	UpdateUserExchange(mkey string, payload *models.UserExchange, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteUserExchange(mkey string, params *models.CmdbRequestParams) error
	CreateUserFortitoken(payload *models.UserFortitoken, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadUserFortitoken(mkey string, params *models.CmdbRequestParams) (*models.UserFortitoken, error)
	UpdateUserFortitoken(mkey string, payload *models.UserFortitoken, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteUserFortitoken(mkey string, params *models.CmdbRequestParams) error
	CreateUserFsso(payload *models.UserFsso, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadUserFsso(mkey string, params *models.CmdbRequestParams) (*models.UserFsso, error)
	UpdateUserFsso(mkey string, payload *models.UserFsso, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteUserFsso(mkey string, params *models.CmdbRequestParams) error
	CreateUserFssoPolling(payload *models.UserFssoPolling, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadUserFssoPolling(mkey string, params *models.CmdbRequestParams) (*models.UserFssoPolling, error)
	UpdateUserFssoPolling(mkey string, payload *models.UserFssoPolling, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteUserFssoPolling(mkey string, params *models.CmdbRequestParams) error
	CreateUserGroup(payload *models.UserGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadUserGroup(mkey string, params *models.CmdbRequestParams) (*models.UserGroup, error)
	UpdateUserGroup(mkey string, payload *models.UserGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteUserGroup(mkey string, params *models.CmdbRequestParams) error
	CreateUserKrbKeytab(payload *models.UserKrbKeytab, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadUserKrbKeytab(mkey string, params *models.CmdbRequestParams) (*models.UserKrbKeytab, error)
	UpdateUserKrbKeytab(mkey string, payload *models.UserKrbKeytab, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteUserKrbKeytab(mkey string, params *models.CmdbRequestParams) error
	CreateUserLdap(payload *models.UserLdap, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadUserLdap(mkey string, params *models.CmdbRequestParams) (*models.UserLdap, error)
	UpdateUserLdap(mkey string, payload *models.UserLdap, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteUserLdap(mkey string, params *models.CmdbRequestParams) error
	CreateUserLocal(payload *models.UserLocal, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadUserLocal(mkey string, params *models.CmdbRequestParams) (*models.UserLocal, error)
	UpdateUserLocal(mkey string, payload *models.UserLocal, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteUserLocal(mkey string, params *models.CmdbRequestParams) error
	CreateUserNacPolicy(payload *models.UserNacPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadUserNacPolicy(mkey string, params *models.CmdbRequestParams) (*models.UserNacPolicy, error)
	UpdateUserNacPolicy(mkey string, payload *models.UserNacPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteUserNacPolicy(mkey string, params *models.CmdbRequestParams) error
	CreateUserPasswordPolicy(payload *models.UserPasswordPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadUserPasswordPolicy(mkey string, params *models.CmdbRequestParams) (*models.UserPasswordPolicy, error)
	UpdateUserPasswordPolicy(mkey string, payload *models.UserPasswordPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteUserPasswordPolicy(mkey string, params *models.CmdbRequestParams) error
	CreateUserPeer(payload *models.UserPeer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadUserPeer(mkey string, params *models.CmdbRequestParams) (*models.UserPeer, error)
	UpdateUserPeer(mkey string, payload *models.UserPeer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteUserPeer(mkey string, params *models.CmdbRequestParams) error
	CreateUserPeergrp(payload *models.UserPeergrp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadUserPeergrp(mkey string, params *models.CmdbRequestParams) (*models.UserPeergrp, error)
	UpdateUserPeergrp(mkey string, payload *models.UserPeergrp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteUserPeergrp(mkey string, params *models.CmdbRequestParams) error
	CreateUserPop3(payload *models.UserPop3, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadUserPop3(mkey string, params *models.CmdbRequestParams) (*models.UserPop3, error)
	UpdateUserPop3(mkey string, payload *models.UserPop3, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteUserPop3(mkey string, params *models.CmdbRequestParams) error
	CreateUserQuarantine(payload *models.UserQuarantine, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadUserQuarantine(mkey string, params *models.CmdbRequestParams) (*models.UserQuarantine, error)
	UpdateUserQuarantine(mkey string, payload *models.UserQuarantine, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteUserQuarantine(mkey string, params *models.CmdbRequestParams) error
	CreateUserRadius(payload *models.UserRadius, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadUserRadius(mkey string, params *models.CmdbRequestParams) (*models.UserRadius, error)
	UpdateUserRadius(mkey string, payload *models.UserRadius, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteUserRadius(mkey string, params *models.CmdbRequestParams) error
	CreateUserSaml(payload *models.UserSaml, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadUserSaml(mkey string, params *models.CmdbRequestParams) (*models.UserSaml, error)
	UpdateUserSaml(mkey string, payload *models.UserSaml, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteUserSaml(mkey string, params *models.CmdbRequestParams) error
	CreateUserSecurityExemptList(payload *models.UserSecurityExemptList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadUserSecurityExemptList(mkey string, params *models.CmdbRequestParams) (*models.UserSecurityExemptList, error)
	UpdateUserSecurityExemptList(mkey string, payload *models.UserSecurityExemptList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteUserSecurityExemptList(mkey string, params *models.CmdbRequestParams) error
	CreateUserSetting(payload *models.UserSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadUserSetting(mkey string, params *models.CmdbRequestParams) (*models.UserSetting, error)
	UpdateUserSetting(mkey string, payload *models.UserSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteUserSetting(mkey string, params *models.CmdbRequestParams) error
	CreateUserTacacs(payload *models.UserTacacs, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadUserTacacs(mkey string, params *models.CmdbRequestParams) (*models.UserTacacs, error)
	UpdateUserTacacs(mkey string, payload *models.UserTacacs, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteUserTacacs(mkey string, params *models.CmdbRequestParams) error
	CreateVideofilterProfile(payload *models.VideofilterProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVideofilterProfile(mkey string, params *models.CmdbRequestParams) (*models.VideofilterProfile, error)
	UpdateVideofilterProfile(mkey string, payload *models.VideofilterProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVideofilterProfile(mkey string, params *models.CmdbRequestParams) error
	CreateVideofilterYoutubeChannelFilter(payload *models.VideofilterYoutubeChannelFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVideofilterYoutubeChannelFilter(mkey string, params *models.CmdbRequestParams) (*models.VideofilterYoutubeChannelFilter, error)
	UpdateVideofilterYoutubeChannelFilter(mkey string, payload *models.VideofilterYoutubeChannelFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVideofilterYoutubeChannelFilter(mkey string, params *models.CmdbRequestParams) error
	CreateVideofilterYoutubeKey(payload *models.VideofilterYoutubeKey, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVideofilterYoutubeKey(mkey string, params *models.CmdbRequestParams) (*models.VideofilterYoutubeKey, error)
	UpdateVideofilterYoutubeKey(mkey string, payload *models.VideofilterYoutubeKey, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVideofilterYoutubeKey(mkey string, params *models.CmdbRequestParams) error
	CreateVoipProfile(payload *models.VoipProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVoipProfile(mkey string, params *models.CmdbRequestParams) (*models.VoipProfile, error)
	UpdateVoipProfile(mkey string, payload *models.VoipProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVoipProfile(mkey string, params *models.CmdbRequestParams) error
	CreateVpncertificateCa(payload *models.VpncertificateCa, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpncertificateCa(mkey string, params *models.CmdbRequestParams) (*models.VpncertificateCa, error)
	UpdateVpncertificateCa(mkey string, payload *models.VpncertificateCa, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpncertificateCa(mkey string, params *models.CmdbRequestParams) error
	CreateVpncertificateCrl(payload *models.VpncertificateCrl, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpncertificateCrl(mkey string, params *models.CmdbRequestParams) (*models.VpncertificateCrl, error)
	UpdateVpncertificateCrl(mkey string, payload *models.VpncertificateCrl, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpncertificateCrl(mkey string, params *models.CmdbRequestParams) error
	CreateVpncertificateLocal(payload *models.VpncertificateLocal, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpncertificateLocal(mkey string, params *models.CmdbRequestParams) (*models.VpncertificateLocal, error)
	UpdateVpncertificateLocal(mkey string, payload *models.VpncertificateLocal, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpncertificateLocal(mkey string, params *models.CmdbRequestParams) error
	CreateVpncertificateOcspServer(payload *models.VpncertificateOcspServer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpncertificateOcspServer(mkey string, params *models.CmdbRequestParams) (*models.VpncertificateOcspServer, error)
	UpdateVpncertificateOcspServer(mkey string, payload *models.VpncertificateOcspServer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpncertificateOcspServer(mkey string, params *models.CmdbRequestParams) error
	CreateVpncertificateRemote(payload *models.VpncertificateRemote, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpncertificateRemote(mkey string, params *models.CmdbRequestParams) (*models.VpncertificateRemote, error)
	UpdateVpncertificateRemote(mkey string, payload *models.VpncertificateRemote, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpncertificateRemote(mkey string, params *models.CmdbRequestParams) error
	CreateVpncertificateSetting(payload *models.VpncertificateSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpncertificateSetting(mkey string, params *models.CmdbRequestParams) (*models.VpncertificateSetting, error)
	UpdateVpncertificateSetting(mkey string, payload *models.VpncertificateSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpncertificateSetting(mkey string, params *models.CmdbRequestParams) error
	CreateVpnipsecConcentrator(payload *models.VpnipsecConcentrator, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnipsecConcentrator(mkey string, params *models.CmdbRequestParams) (*models.VpnipsecConcentrator, error)
	UpdateVpnipsecConcentrator(mkey string, payload *models.VpnipsecConcentrator, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnipsecConcentrator(mkey string, params *models.CmdbRequestParams) error
	CreateVpnipsecFec(payload *models.VpnipsecFec, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnipsecFec(mkey string, params *models.CmdbRequestParams) (*models.VpnipsecFec, error)
	UpdateVpnipsecFec(mkey string, payload *models.VpnipsecFec, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnipsecFec(mkey string, params *models.CmdbRequestParams) error
	CreateVpnipsecForticlient(payload *models.VpnipsecForticlient, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnipsecForticlient(mkey string, params *models.CmdbRequestParams) (*models.VpnipsecForticlient, error)
	UpdateVpnipsecForticlient(mkey string, payload *models.VpnipsecForticlient, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnipsecForticlient(mkey string, params *models.CmdbRequestParams) error
	CreateVpnipsecManualkey(payload *models.VpnipsecManualkey, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnipsecManualkey(mkey string, params *models.CmdbRequestParams) (*models.VpnipsecManualkey, error)
	UpdateVpnipsecManualkey(mkey string, payload *models.VpnipsecManualkey, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnipsecManualkey(mkey string, params *models.CmdbRequestParams) error
	CreateVpnipsecManualkeyInterface(payload *models.VpnipsecManualkeyInterface, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnipsecManualkeyInterface(mkey string, params *models.CmdbRequestParams) (*models.VpnipsecManualkeyInterface, error)
	UpdateVpnipsecManualkeyInterface(mkey string, payload *models.VpnipsecManualkeyInterface, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnipsecManualkeyInterface(mkey string, params *models.CmdbRequestParams) error
	CreateVpnipsecPhase1(payload *models.VpnipsecPhase1, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnipsecPhase1(mkey string, params *models.CmdbRequestParams) (*models.VpnipsecPhase1, error)
	UpdateVpnipsecPhase1(mkey string, payload *models.VpnipsecPhase1, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnipsecPhase1(mkey string, params *models.CmdbRequestParams) error
	CreateVpnipsecPhase1Interface(payload *models.VpnipsecPhase1Interface, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnipsecPhase1Interface(mkey string, params *models.CmdbRequestParams) (*models.VpnipsecPhase1Interface, error)
	UpdateVpnipsecPhase1Interface(mkey string, payload *models.VpnipsecPhase1Interface, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnipsecPhase1Interface(mkey string, params *models.CmdbRequestParams) error
	CreateVpnipsecPhase2(payload *models.VpnipsecPhase2, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnipsecPhase2(mkey string, params *models.CmdbRequestParams) (*models.VpnipsecPhase2, error)
	UpdateVpnipsecPhase2(mkey string, payload *models.VpnipsecPhase2, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnipsecPhase2(mkey string, params *models.CmdbRequestParams) error
	CreateVpnipsecPhase2Interface(payload *models.VpnipsecPhase2Interface, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnipsecPhase2Interface(mkey string, params *models.CmdbRequestParams) (*models.VpnipsecPhase2Interface, error)
	UpdateVpnipsecPhase2Interface(mkey string, payload *models.VpnipsecPhase2Interface, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnipsecPhase2Interface(mkey string, params *models.CmdbRequestParams) error
	CreateVpnsslsettingsAuthenticationRule(payload *models.VpnsslsettingsAuthenticationRule, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnsslsettingsAuthenticationRule(mkey string, params *models.CmdbRequestParams) (*models.VpnsslsettingsAuthenticationRule, error)
	UpdateVpnsslsettingsAuthenticationRule(mkey string, payload *models.VpnsslsettingsAuthenticationRule, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnsslsettingsAuthenticationRule(mkey string, params *models.CmdbRequestParams) error
	CreateVpnsslwebHostCheckSoftware(payload *models.VpnsslwebHostCheckSoftware, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnsslwebHostCheckSoftware(mkey string, params *models.CmdbRequestParams) (*models.VpnsslwebHostCheckSoftware, error)
	UpdateVpnsslwebHostCheckSoftware(mkey string, payload *models.VpnsslwebHostCheckSoftware, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnsslwebHostCheckSoftware(mkey string, params *models.CmdbRequestParams) error
	CreateVpnsslwebPortal(payload *models.VpnsslwebPortal, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnsslwebPortal(mkey string, params *models.CmdbRequestParams) (*models.VpnsslwebPortal, error)
	UpdateVpnsslwebPortal(mkey string, payload *models.VpnsslwebPortal, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnsslwebPortal(mkey string, params *models.CmdbRequestParams) error
	CreateVpnsslwebRealm(payload *models.VpnsslwebRealm, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnsslwebRealm(mkey string, params *models.CmdbRequestParams) (*models.VpnsslwebRealm, error)
	UpdateVpnsslwebRealm(mkey string, payload *models.VpnsslwebRealm, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnsslwebRealm(mkey string, params *models.CmdbRequestParams) error
	CreateVpnsslwebUserBookmark(payload *models.VpnsslwebUserBookmark, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnsslwebUserBookmark(mkey string, params *models.CmdbRequestParams) (*models.VpnsslwebUserBookmark, error)
	UpdateVpnsslwebUserBookmark(mkey string, payload *models.VpnsslwebUserBookmark, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnsslwebUserBookmark(mkey string, params *models.CmdbRequestParams) error
	CreateVpnsslwebUserGroupBookmark(payload *models.VpnsslwebUserGroupBookmark, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnsslwebUserGroupBookmark(mkey string, params *models.CmdbRequestParams) (*models.VpnsslwebUserGroupBookmark, error)
	UpdateVpnsslwebUserGroupBookmark(mkey string, payload *models.VpnsslwebUserGroupBookmark, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnsslwebUserGroupBookmark(mkey string, params *models.CmdbRequestParams) error
	CreateVpnsslClient(payload *models.VpnsslClient, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnsslClient(mkey string, params *models.CmdbRequestParams) (*models.VpnsslClient, error)
	UpdateVpnsslClient(mkey string, payload *models.VpnsslClient, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnsslClient(mkey string, params *models.CmdbRequestParams) error
	CreateVpnsslSettings(payload *models.VpnsslSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnsslSettings(mkey string, params *models.CmdbRequestParams) (*models.VpnsslSettings, error)
	UpdateVpnsslSettings(mkey string, payload *models.VpnsslSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnsslSettings(mkey string, params *models.CmdbRequestParams) error
	CreateVpnL2tp(payload *models.VpnL2tp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnL2tp(mkey string, params *models.CmdbRequestParams) (*models.VpnL2tp, error)
	UpdateVpnL2tp(mkey string, payload *models.VpnL2tp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnL2tp(mkey string, params *models.CmdbRequestParams) error
	CreateVpnOcvpn(payload *models.VpnOcvpn, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnOcvpn(mkey string, params *models.CmdbRequestParams) (*models.VpnOcvpn, error)
	UpdateVpnOcvpn(mkey string, payload *models.VpnOcvpn, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnOcvpn(mkey string, params *models.CmdbRequestParams) error
	CreateVpnPptp(payload *models.VpnPptp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadVpnPptp(mkey string, params *models.CmdbRequestParams) (*models.VpnPptp, error)
	UpdateVpnPptp(mkey string, payload *models.VpnPptp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteVpnPptp(mkey string, params *models.CmdbRequestParams) error
	CreateWafMainClass(payload *models.WafMainClass, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWafMainClass(mkey string, params *models.CmdbRequestParams) (*models.WafMainClass, error)
	UpdateWafMainClass(mkey string, payload *models.WafMainClass, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWafMainClass(mkey string, params *models.CmdbRequestParams) error
	CreateWafProfile(payload *models.WafProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWafProfile(mkey string, params *models.CmdbRequestParams) (*models.WafProfile, error)
	UpdateWafProfile(mkey string, payload *models.WafProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWafProfile(mkey string, params *models.CmdbRequestParams) error
	CreateWafSignature(payload *models.WafSignature, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWafSignature(mkey string, params *models.CmdbRequestParams) (*models.WafSignature, error)
	UpdateWafSignature(mkey string, payload *models.WafSignature, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWafSignature(mkey string, params *models.CmdbRequestParams) error
	CreateWanoptAuthGroup(payload *models.WanoptAuthGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWanoptAuthGroup(mkey string, params *models.CmdbRequestParams) (*models.WanoptAuthGroup, error)
	UpdateWanoptAuthGroup(mkey string, payload *models.WanoptAuthGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWanoptAuthGroup(mkey string, params *models.CmdbRequestParams) error
	CreateWanoptCacheService(payload *models.WanoptCacheService, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWanoptCacheService(mkey string, params *models.CmdbRequestParams) (*models.WanoptCacheService, error)
	UpdateWanoptCacheService(mkey string, payload *models.WanoptCacheService, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWanoptCacheService(mkey string, params *models.CmdbRequestParams) error
	CreateWanoptContentDeliveryNetworkRule(payload *models.WanoptContentDeliveryNetworkRule, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWanoptContentDeliveryNetworkRule(mkey string, params *models.CmdbRequestParams) (*models.WanoptContentDeliveryNetworkRule, error)
	UpdateWanoptContentDeliveryNetworkRule(mkey string, payload *models.WanoptContentDeliveryNetworkRule, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWanoptContentDeliveryNetworkRule(mkey string, params *models.CmdbRequestParams) error
	CreateWanoptPeer(payload *models.WanoptPeer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWanoptPeer(mkey string, params *models.CmdbRequestParams) (*models.WanoptPeer, error)
	UpdateWanoptPeer(mkey string, payload *models.WanoptPeer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWanoptPeer(mkey string, params *models.CmdbRequestParams) error
	CreateWanoptProfile(payload *models.WanoptProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWanoptProfile(mkey string, params *models.CmdbRequestParams) (*models.WanoptProfile, error)
	UpdateWanoptProfile(mkey string, payload *models.WanoptProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWanoptProfile(mkey string, params *models.CmdbRequestParams) error
	CreateWanoptRemoteStorage(payload *models.WanoptRemoteStorage, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWanoptRemoteStorage(mkey string, params *models.CmdbRequestParams) (*models.WanoptRemoteStorage, error)
	UpdateWanoptRemoteStorage(mkey string, payload *models.WanoptRemoteStorage, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWanoptRemoteStorage(mkey string, params *models.CmdbRequestParams) error
	CreateWanoptSettings(payload *models.WanoptSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWanoptSettings(mkey string, params *models.CmdbRequestParams) (*models.WanoptSettings, error)
	UpdateWanoptSettings(mkey string, payload *models.WanoptSettings, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWanoptSettings(mkey string, params *models.CmdbRequestParams) error
	CreateWanoptWebcache(payload *models.WanoptWebcache, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWanoptWebcache(mkey string, params *models.CmdbRequestParams) (*models.WanoptWebcache, error)
	UpdateWanoptWebcache(mkey string, payload *models.WanoptWebcache, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWanoptWebcache(mkey string, params *models.CmdbRequestParams) error
	CreateWebProxyDebugUrl(payload *models.WebProxyDebugUrl, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWebProxyDebugUrl(mkey string, params *models.CmdbRequestParams) (*models.WebProxyDebugUrl, error)
	UpdateWebProxyDebugUrl(mkey string, payload *models.WebProxyDebugUrl, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWebProxyDebugUrl(mkey string, params *models.CmdbRequestParams) error
	CreateWebProxyExplicit(payload *models.WebProxyExplicit, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWebProxyExplicit(mkey string, params *models.CmdbRequestParams) (*models.WebProxyExplicit, error)
	UpdateWebProxyExplicit(mkey string, payload *models.WebProxyExplicit, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWebProxyExplicit(mkey string, params *models.CmdbRequestParams) error
	CreateWebProxyForwardServer(payload *models.WebProxyForwardServer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWebProxyForwardServer(mkey string, params *models.CmdbRequestParams) (*models.WebProxyForwardServer, error)
	UpdateWebProxyForwardServer(mkey string, payload *models.WebProxyForwardServer, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWebProxyForwardServer(mkey string, params *models.CmdbRequestParams) error
	CreateWebProxyForwardServerGroup(payload *models.WebProxyForwardServerGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWebProxyForwardServerGroup(mkey string, params *models.CmdbRequestParams) (*models.WebProxyForwardServerGroup, error)
	UpdateWebProxyForwardServerGroup(mkey string, payload *models.WebProxyForwardServerGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWebProxyForwardServerGroup(mkey string, params *models.CmdbRequestParams) error
	CreateWebProxyGlobal(payload *models.WebProxyGlobal, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWebProxyGlobal(mkey string, params *models.CmdbRequestParams) (*models.WebProxyGlobal, error)
	UpdateWebProxyGlobal(mkey string, payload *models.WebProxyGlobal, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWebProxyGlobal(mkey string, params *models.CmdbRequestParams) error
	CreateWebProxyProfile(payload *models.WebProxyProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWebProxyProfile(mkey string, params *models.CmdbRequestParams) (*models.WebProxyProfile, error)
	UpdateWebProxyProfile(mkey string, payload *models.WebProxyProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWebProxyProfile(mkey string, params *models.CmdbRequestParams) error
	CreateWebProxyUrlMatch(payload *models.WebProxyUrlMatch, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWebProxyUrlMatch(mkey string, params *models.CmdbRequestParams) (*models.WebProxyUrlMatch, error)
	UpdateWebProxyUrlMatch(mkey string, payload *models.WebProxyUrlMatch, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWebProxyUrlMatch(mkey string, params *models.CmdbRequestParams) error
	CreateWebProxyWisp(payload *models.WebProxyWisp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWebProxyWisp(mkey string, params *models.CmdbRequestParams) (*models.WebProxyWisp, error)
	UpdateWebProxyWisp(mkey string, payload *models.WebProxyWisp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWebProxyWisp(mkey string, params *models.CmdbRequestParams) error
	CreateWebfilterContent(payload *models.WebfilterContent, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWebfilterContent(mkey string, params *models.CmdbRequestParams) (*models.WebfilterContent, error)
	UpdateWebfilterContent(mkey string, payload *models.WebfilterContent, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWebfilterContent(mkey string, params *models.CmdbRequestParams) error
	CreateWebfilterContentHeader(payload *models.WebfilterContentHeader, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWebfilterContentHeader(mkey string, params *models.CmdbRequestParams) (*models.WebfilterContentHeader, error)
	UpdateWebfilterContentHeader(mkey string, payload *models.WebfilterContentHeader, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWebfilterContentHeader(mkey string, params *models.CmdbRequestParams) error
	CreateWebfilterFortiguard(payload *models.WebfilterFortiguard, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWebfilterFortiguard(mkey string, params *models.CmdbRequestParams) (*models.WebfilterFortiguard, error)
	UpdateWebfilterFortiguard(mkey string, payload *models.WebfilterFortiguard, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWebfilterFortiguard(mkey string, params *models.CmdbRequestParams) error
	CreateWebfilterFtgdLocalCat(payload *models.WebfilterFtgdLocalCat, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWebfilterFtgdLocalCat(mkey string, params *models.CmdbRequestParams) (*models.WebfilterFtgdLocalCat, error)
	UpdateWebfilterFtgdLocalCat(mkey string, payload *models.WebfilterFtgdLocalCat, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWebfilterFtgdLocalCat(mkey string, params *models.CmdbRequestParams) error
	CreateWebfilterFtgdLocalRating(payload *models.WebfilterFtgdLocalRating, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWebfilterFtgdLocalRating(mkey string, params *models.CmdbRequestParams) (*models.WebfilterFtgdLocalRating, error)
	UpdateWebfilterFtgdLocalRating(mkey string, payload *models.WebfilterFtgdLocalRating, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWebfilterFtgdLocalRating(mkey string, params *models.CmdbRequestParams) error
	CreateWebfilterIpsUrlfilterCacheSetting(payload *models.WebfilterIpsUrlfilterCacheSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWebfilterIpsUrlfilterCacheSetting(mkey string, params *models.CmdbRequestParams) (*models.WebfilterIpsUrlfilterCacheSetting, error)
	UpdateWebfilterIpsUrlfilterCacheSetting(mkey string, payload *models.WebfilterIpsUrlfilterCacheSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWebfilterIpsUrlfilterCacheSetting(mkey string, params *models.CmdbRequestParams) error
	CreateWebfilterIpsUrlfilterSetting(payload *models.WebfilterIpsUrlfilterSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWebfilterIpsUrlfilterSetting(mkey string, params *models.CmdbRequestParams) (*models.WebfilterIpsUrlfilterSetting, error)
	UpdateWebfilterIpsUrlfilterSetting(mkey string, payload *models.WebfilterIpsUrlfilterSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWebfilterIpsUrlfilterSetting(mkey string, params *models.CmdbRequestParams) error
	CreateWebfilterIpsUrlfilterSetting6(payload *models.WebfilterIpsUrlfilterSetting6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWebfilterIpsUrlfilterSetting6(mkey string, params *models.CmdbRequestParams) (*models.WebfilterIpsUrlfilterSetting6, error)
	UpdateWebfilterIpsUrlfilterSetting6(mkey string, payload *models.WebfilterIpsUrlfilterSetting6, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWebfilterIpsUrlfilterSetting6(mkey string, params *models.CmdbRequestParams) error
	CreateWebfilterOverride(payload *models.WebfilterOverride, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWebfilterOverride(mkey string, params *models.CmdbRequestParams) (*models.WebfilterOverride, error)
	UpdateWebfilterOverride(mkey string, payload *models.WebfilterOverride, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWebfilterOverride(mkey string, params *models.CmdbRequestParams) error
	CreateWebfilterProfile(payload *models.WebfilterProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWebfilterProfile(mkey string, params *models.CmdbRequestParams) (*models.WebfilterProfile, error)
	UpdateWebfilterProfile(mkey string, payload *models.WebfilterProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWebfilterProfile(mkey string, params *models.CmdbRequestParams) error
	CreateWebfilterSearchEngine(payload *models.WebfilterSearchEngine, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWebfilterSearchEngine(mkey string, params *models.CmdbRequestParams) (*models.WebfilterSearchEngine, error)
	UpdateWebfilterSearchEngine(mkey string, payload *models.WebfilterSearchEngine, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWebfilterSearchEngine(mkey string, params *models.CmdbRequestParams) error
	CreateWebfilterUrlfilter(payload *models.WebfilterUrlfilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWebfilterUrlfilter(mkey string, params *models.CmdbRequestParams) (*models.WebfilterUrlfilter, error)
	UpdateWebfilterUrlfilter(mkey string, payload *models.WebfilterUrlfilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWebfilterUrlfilter(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerhotspot20Anqp3gppCellular(payload *models.WirelessControllerhotspot20Anqp3gppCellular, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerhotspot20Anqp3gppCellular(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerhotspot20Anqp3gppCellular, error)
	UpdateWirelessControllerhotspot20Anqp3gppCellular(mkey string, payload *models.WirelessControllerhotspot20Anqp3gppCellular, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerhotspot20Anqp3gppCellular(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerhotspot20AnqpIpAddressType(payload *models.WirelessControllerhotspot20AnqpIpAddressType, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerhotspot20AnqpIpAddressType(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerhotspot20AnqpIpAddressType, error)
	UpdateWirelessControllerhotspot20AnqpIpAddressType(mkey string, payload *models.WirelessControllerhotspot20AnqpIpAddressType, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerhotspot20AnqpIpAddressType(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerhotspot20AnqpNaiRealm(payload *models.WirelessControllerhotspot20AnqpNaiRealm, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerhotspot20AnqpNaiRealm(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerhotspot20AnqpNaiRealm, error)
	UpdateWirelessControllerhotspot20AnqpNaiRealm(mkey string, payload *models.WirelessControllerhotspot20AnqpNaiRealm, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerhotspot20AnqpNaiRealm(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerhotspot20AnqpNetworkAuthType(payload *models.WirelessControllerhotspot20AnqpNetworkAuthType, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerhotspot20AnqpNetworkAuthType(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerhotspot20AnqpNetworkAuthType, error)
	UpdateWirelessControllerhotspot20AnqpNetworkAuthType(mkey string, payload *models.WirelessControllerhotspot20AnqpNetworkAuthType, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerhotspot20AnqpNetworkAuthType(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerhotspot20AnqpRoamingConsortium(payload *models.WirelessControllerhotspot20AnqpRoamingConsortium, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerhotspot20AnqpRoamingConsortium(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerhotspot20AnqpRoamingConsortium, error)
	UpdateWirelessControllerhotspot20AnqpRoamingConsortium(mkey string, payload *models.WirelessControllerhotspot20AnqpRoamingConsortium, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerhotspot20AnqpRoamingConsortium(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerhotspot20AnqpVenueName(payload *models.WirelessControllerhotspot20AnqpVenueName, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerhotspot20AnqpVenueName(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerhotspot20AnqpVenueName, error)
	UpdateWirelessControllerhotspot20AnqpVenueName(mkey string, payload *models.WirelessControllerhotspot20AnqpVenueName, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerhotspot20AnqpVenueName(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerhotspot20AnqpVenueUrl(payload *models.WirelessControllerhotspot20AnqpVenueUrl, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerhotspot20AnqpVenueUrl(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerhotspot20AnqpVenueUrl, error)
	UpdateWirelessControllerhotspot20AnqpVenueUrl(mkey string, payload *models.WirelessControllerhotspot20AnqpVenueUrl, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerhotspot20AnqpVenueUrl(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerhotspot20H2qpAdviceOfCharge(payload *models.WirelessControllerhotspot20H2qpAdviceOfCharge, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerhotspot20H2qpAdviceOfCharge(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerhotspot20H2qpAdviceOfCharge, error)
	UpdateWirelessControllerhotspot20H2qpAdviceOfCharge(mkey string, payload *models.WirelessControllerhotspot20H2qpAdviceOfCharge, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerhotspot20H2qpAdviceOfCharge(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerhotspot20H2qpConnCapability(payload *models.WirelessControllerhotspot20H2qpConnCapability, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerhotspot20H2qpConnCapability(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerhotspot20H2qpConnCapability, error)
	UpdateWirelessControllerhotspot20H2qpConnCapability(mkey string, payload *models.WirelessControllerhotspot20H2qpConnCapability, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerhotspot20H2qpConnCapability(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerhotspot20H2qpOperatorName(payload *models.WirelessControllerhotspot20H2qpOperatorName, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerhotspot20H2qpOperatorName(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerhotspot20H2qpOperatorName, error)
	UpdateWirelessControllerhotspot20H2qpOperatorName(mkey string, payload *models.WirelessControllerhotspot20H2qpOperatorName, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerhotspot20H2qpOperatorName(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerhotspot20H2qpOsuProvider(payload *models.WirelessControllerhotspot20H2qpOsuProvider, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerhotspot20H2qpOsuProvider(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerhotspot20H2qpOsuProvider, error)
	UpdateWirelessControllerhotspot20H2qpOsuProvider(mkey string, payload *models.WirelessControllerhotspot20H2qpOsuProvider, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerhotspot20H2qpOsuProvider(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerhotspot20H2qpOsuProviderNai(payload *models.WirelessControllerhotspot20H2qpOsuProviderNai, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerhotspot20H2qpOsuProviderNai(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerhotspot20H2qpOsuProviderNai, error)
	UpdateWirelessControllerhotspot20H2qpOsuProviderNai(mkey string, payload *models.WirelessControllerhotspot20H2qpOsuProviderNai, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerhotspot20H2qpOsuProviderNai(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerhotspot20H2qpTermsAndConditions(payload *models.WirelessControllerhotspot20H2qpTermsAndConditions, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerhotspot20H2qpTermsAndConditions(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerhotspot20H2qpTermsAndConditions, error)
	UpdateWirelessControllerhotspot20H2qpTermsAndConditions(mkey string, payload *models.WirelessControllerhotspot20H2qpTermsAndConditions, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerhotspot20H2qpTermsAndConditions(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerhotspot20H2qpWanMetric(payload *models.WirelessControllerhotspot20H2qpWanMetric, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerhotspot20H2qpWanMetric(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerhotspot20H2qpWanMetric, error)
	UpdateWirelessControllerhotspot20H2qpWanMetric(mkey string, payload *models.WirelessControllerhotspot20H2qpWanMetric, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerhotspot20H2qpWanMetric(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerhotspot20HsProfile(payload *models.WirelessControllerhotspot20HsProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerhotspot20HsProfile(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerhotspot20HsProfile, error)
	UpdateWirelessControllerhotspot20HsProfile(mkey string, payload *models.WirelessControllerhotspot20HsProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerhotspot20HsProfile(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerhotspot20Icon(payload *models.WirelessControllerhotspot20Icon, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerhotspot20Icon(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerhotspot20Icon, error)
	UpdateWirelessControllerhotspot20Icon(mkey string, payload *models.WirelessControllerhotspot20Icon, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerhotspot20Icon(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerhotspot20QosMap(payload *models.WirelessControllerhotspot20QosMap, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerhotspot20QosMap(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerhotspot20QosMap, error)
	UpdateWirelessControllerhotspot20QosMap(mkey string, payload *models.WirelessControllerhotspot20QosMap, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerhotspot20QosMap(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerAccessControlList(payload *models.WirelessControllerAccessControlList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerAccessControlList(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerAccessControlList, error)
	UpdateWirelessControllerAccessControlList(mkey string, payload *models.WirelessControllerAccessControlList, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerAccessControlList(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerAddress(payload *models.WirelessControllerAddress, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerAddress(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerAddress, error)
	UpdateWirelessControllerAddress(mkey string, payload *models.WirelessControllerAddress, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerAddress(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerAddrgrp(payload *models.WirelessControllerAddrgrp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerAddrgrp(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerAddrgrp, error)
	UpdateWirelessControllerAddrgrp(mkey string, payload *models.WirelessControllerAddrgrp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerAddrgrp(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerApStatus(payload *models.WirelessControllerApStatus, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerApStatus(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerApStatus, error)
	UpdateWirelessControllerApStatus(mkey string, payload *models.WirelessControllerApStatus, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerApStatus(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerApcfgProfile(payload *models.WirelessControllerApcfgProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerApcfgProfile(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerApcfgProfile, error)
	UpdateWirelessControllerApcfgProfile(mkey string, payload *models.WirelessControllerApcfgProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerApcfgProfile(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerArrpProfile(payload *models.WirelessControllerArrpProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerArrpProfile(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerArrpProfile, error)
	UpdateWirelessControllerArrpProfile(mkey string, payload *models.WirelessControllerArrpProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerArrpProfile(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerBleProfile(payload *models.WirelessControllerBleProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerBleProfile(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerBleProfile, error)
	UpdateWirelessControllerBleProfile(mkey string, payload *models.WirelessControllerBleProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerBleProfile(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerBonjourProfile(payload *models.WirelessControllerBonjourProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerBonjourProfile(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerBonjourProfile, error)
	UpdateWirelessControllerBonjourProfile(mkey string, payload *models.WirelessControllerBonjourProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerBonjourProfile(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerGlobal(payload *models.WirelessControllerGlobal, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerGlobal(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerGlobal, error)
	UpdateWirelessControllerGlobal(mkey string, payload *models.WirelessControllerGlobal, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerGlobal(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerInterController(payload *models.WirelessControllerInterController, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerInterController(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerInterController, error)
	UpdateWirelessControllerInterController(mkey string, payload *models.WirelessControllerInterController, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerInterController(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerLog(payload *models.WirelessControllerLog, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerLog(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerLog, error)
	UpdateWirelessControllerLog(mkey string, payload *models.WirelessControllerLog, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerLog(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerMpskProfile(payload *models.WirelessControllerMpskProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerMpskProfile(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerMpskProfile, error)
	UpdateWirelessControllerMpskProfile(mkey string, payload *models.WirelessControllerMpskProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerMpskProfile(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerNacProfile(payload *models.WirelessControllerNacProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerNacProfile(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerNacProfile, error)
	UpdateWirelessControllerNacProfile(mkey string, payload *models.WirelessControllerNacProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerNacProfile(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerQosProfile(payload *models.WirelessControllerQosProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerQosProfile(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerQosProfile, error)
	UpdateWirelessControllerQosProfile(mkey string, payload *models.WirelessControllerQosProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerQosProfile(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerRegion(payload *models.WirelessControllerRegion, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerRegion(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerRegion, error)
	UpdateWirelessControllerRegion(mkey string, payload *models.WirelessControllerRegion, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerRegion(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerSetting(payload *models.WirelessControllerSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerSetting(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerSetting, error)
	UpdateWirelessControllerSetting(mkey string, payload *models.WirelessControllerSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerSetting(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerSnmp(payload *models.WirelessControllerSnmp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerSnmp(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerSnmp, error)
	UpdateWirelessControllerSnmp(mkey string, payload *models.WirelessControllerSnmp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerSnmp(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerSsidPolicy(payload *models.WirelessControllerSsidPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerSsidPolicy(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerSsidPolicy, error)
	UpdateWirelessControllerSsidPolicy(mkey string, payload *models.WirelessControllerSsidPolicy, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerSsidPolicy(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerSyslogProfile(payload *models.WirelessControllerSyslogProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerSyslogProfile(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerSyslogProfile, error)
	UpdateWirelessControllerSyslogProfile(mkey string, payload *models.WirelessControllerSyslogProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerSyslogProfile(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerTimers(payload *models.WirelessControllerTimers, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerTimers(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerTimers, error)
	UpdateWirelessControllerTimers(mkey string, payload *models.WirelessControllerTimers, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerTimers(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerUtmProfile(payload *models.WirelessControllerUtmProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerUtmProfile(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerUtmProfile, error)
	UpdateWirelessControllerUtmProfile(mkey string, payload *models.WirelessControllerUtmProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerUtmProfile(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerVap(payload *models.WirelessControllerVap, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerVap(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerVap, error)
	UpdateWirelessControllerVap(mkey string, payload *models.WirelessControllerVap, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerVap(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerVapGroup(payload *models.WirelessControllerVapGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerVapGroup(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerVapGroup, error)
	UpdateWirelessControllerVapGroup(mkey string, payload *models.WirelessControllerVapGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerVapGroup(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerWagProfile(payload *models.WirelessControllerWagProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerWagProfile(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerWagProfile, error)
	UpdateWirelessControllerWagProfile(mkey string, payload *models.WirelessControllerWagProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerWagProfile(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerWidsProfile(payload *models.WirelessControllerWidsProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerWidsProfile(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerWidsProfile, error)
	UpdateWirelessControllerWidsProfile(mkey string, payload *models.WirelessControllerWidsProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerWidsProfile(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerWtp(payload *models.WirelessControllerWtp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerWtp(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerWtp, error)
	UpdateWirelessControllerWtp(mkey string, payload *models.WirelessControllerWtp, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerWtp(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerWtpGroup(payload *models.WirelessControllerWtpGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerWtpGroup(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerWtpGroup, error)
	UpdateWirelessControllerWtpGroup(mkey string, payload *models.WirelessControllerWtpGroup, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerWtpGroup(mkey string, params *models.CmdbRequestParams) error
	CreateWirelessControllerWtpProfile(payload *models.WirelessControllerWtpProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadWirelessControllerWtpProfile(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerWtpProfile, error)
	UpdateWirelessControllerWtpProfile(mkey string, payload *models.WirelessControllerWtpProfile, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteWirelessControllerWtpProfile(mkey string, params *models.CmdbRequestParams) error
}
