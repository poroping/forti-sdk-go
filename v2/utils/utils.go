package utils

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
	GenericApi(payload *UtilsGenericApi, params *models.CmdbRequestParams) (*string, error)
	SortCentralSnatMap(params *models.CmdbRequestParams, sortBy, sortDirection string) (err error)
	CentralSnatMapListIsSorted(params *models.CmdbRequestParams, sortBy, sortDirection string) (sorted bool, err error)
	SortFirewallPolicy(params *models.CmdbRequestParams, sortBy, sortDirection string) (err error)
	FirewallPolicyListIsSorted(params *models.CmdbRequestParams, sortBy, sortDirection string) (sorted bool, err error)
	SortFirewallDosPolicy(params *models.CmdbRequestParams, sortBy, sortDirection string) (err error)
	FirewallDosPolicyListIsSorted(params *models.CmdbRequestParams, sortBy, sortDirection string) (sorted bool, err error)
	SortFirewallDosPolicy6(params *models.CmdbRequestParams, sortBy, sortDirection string) (err error)
	FirewallDosPolicy6ListIsSorted(params *models.CmdbRequestParams, sortBy, sortDirection string) (sorted bool, err error)
	SortFirewallInterfacePolicy(params *models.CmdbRequestParams, sortBy, sortDirection string) (err error)
	FirewallInterfacePolicyListIsSorted(params *models.CmdbRequestParams, sortBy, sortDirection string) (sorted bool, err error)
	SortFirewallInterfacePolicy6(params *models.CmdbRequestParams, sortBy, sortDirection string) (err error)
	FirewallInterfacePolicy6ListIsSorted(params *models.CmdbRequestParams, sortBy, sortDirection string) (sorted bool, err error)
	SortFirewallLocalInPolicy(params *models.CmdbRequestParams, sortBy, sortDirection string) (err error)
	FirewallLocalInPolicyListIsSorted(params *models.CmdbRequestParams, sortBy, sortDirection string) (sorted bool, err error)
	SortFirewallLocalInPolicy6(params *models.CmdbRequestParams, sortBy, sortDirection string) (err error)
	FirewallLocalInPolicy6ListIsSorted(params *models.CmdbRequestParams, sortBy, sortDirection string) (sorted bool, err error)
	SortFirewallMulticastPolicy(params *models.CmdbRequestParams, sortBy, sortDirection string) (err error)
	FirewallMulticastPolicyListIsSorted(params *models.CmdbRequestParams, sortBy, sortDirection string) (sorted bool, err error)
	SortFirewallMulticastPolicy6(params *models.CmdbRequestParams, sortBy, sortDirection string) (err error)
	FirewallMulticastPolicy6ListIsSorted(params *models.CmdbRequestParams, sortBy, sortDirection string) (sorted bool, err error)
	SortFirewallProxyPolicy(params *models.CmdbRequestParams, sortBy, sortDirection string) (err error)
	FirewallProxyPolicyListIsSorted(params *models.CmdbRequestParams, sortBy, sortDirection string) (sorted bool, err error)
	SortFirewallSecurityPolicy(params *models.CmdbRequestParams, sortBy, sortDirection string) (err error)
	FirewallSecurityPolicyListIsSorted(params *models.CmdbRequestParams, sortBy, sortDirection string) (sorted bool, err error)
	SortFirewallShapingPolicy(params *models.CmdbRequestParams, sortBy, sortDirection string) (err error)
	FirewallShapingPolicyListIsSorted(params *models.CmdbRequestParams, sortBy, sortDirection string) (sorted bool, err error)
}
