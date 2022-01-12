package monitor

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
	MonitorRouterBgpNeighbors() (*[]MonitorRouterBgpNeighborsResult, error)
	MonitorRouterBgpPaths() (*[]MonitorRouterBgpPathsResult, error)
	MonitorRouterOspfNeighbors() (*[]MonitorRouterOspfNeighborsResult, error)
	MonitorRouterIPv4() (*[]MonitorRouterIPv4Result, error)
	MonitorSystemCertificateDownload(params *models.CmdbRequestParams) (cert *string, err error)
	MonitorSystemVpnCertificateLocalImport(payload *MonitorSystemVpnCertificateLocalImport, params *models.CmdbRequestParams) (*models.CmdbResponse, error)
	MonitorSystemVpnCertificateRemoteImport(payload *MonitorSystemVpnCertificateRemoteImport, params *models.CmdbRequestParams) (*string, error)
}
