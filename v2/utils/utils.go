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
	GenericApi(payload *UtilsGenericApi, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	SortFirewallPolicy(params *models.CmdbRequestParams, sortBy, sortDirection string) (err error)
	FirewallPolicyListIsSorted(params *models.CmdbRequestParams, sortBy, sortDirection string) (sorted bool, err error)
}
