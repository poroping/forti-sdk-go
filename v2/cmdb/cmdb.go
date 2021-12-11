package cmdb

import (
	"github.com/poroping/forti-sdk-go/v2/config"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func New(config config.Config) Endpoints {
	return &Client{config: config}
}

type Client struct {
	config config.Config
}

type Endpoints interface {
	CreateRouterStatic(payload *models.RouterStatic, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	ReadRouterStatic(mkey string, params *models.CmdbRequestParams) (*models.RouterStatic, error)
	UpdateRouterStatic(mkey string, payload *models.RouterStatic, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	DeleteRouterStatic(mkey string, params *models.CmdbRequestParams) error
}
