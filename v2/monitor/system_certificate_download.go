package monitor

import (
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/forti-sdk-go/v2/request"
)

type MonitorSystemCertificateDownload struct {
	Type string `json:"type"`
}

const MonitorSystemCertificateDownloadPath = "system/certificate/download/"

func (c *Client) MonitorSystemCertificateDownload(params *models.CmdbRequestParams) (cert *string, err error) {
	req := &models.CmdbRequest{}
	req.HTTPMethod = "GET"
	req.Payload = nil
	req.Path = models.MonitorBasePath + MonitorSystemCertificateDownloadPath
	req.Params = *params

	res, err := request.ReadString(c.config, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
