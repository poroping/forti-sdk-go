package monitor

import (
	"encoding/base64"
	"encoding/json"

	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/forti-sdk-go/v2/request"
)

type MonitorSystemVpnCertificateLocalImport struct {
	Name        string `json:"certname"`
	Type        string `json:"type"`
	Certificate string `json:"file_content"`
	Password    string `json:"password,omitempty"`
	PrivateKey  string `json:"key_file_content,omitempty"`
	Scope       string `json:"scope"`
}

const MonitorSystemVpnCertificateLocalImportPath = "vpn-certificate/local/import/"

// Uploads certificate to Fortigate
func (c *Client) MonitorSystemVpnCertificateLocalImport(payload *MonitorSystemVpnCertificateLocalImport, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	payload.Certificate = base64.StdEncoding.EncodeToString([]byte(payload.Certificate))
	if payload.PrivateKey != "" {
		payload.PrivateKey = base64.StdEncoding.EncodeToString([]byte(payload.PrivateKey))
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req := &models.CmdbRequest{}
	req.HTTPMethod = "POST"
	req.Payload = body
	req.Path = models.MonitorBasePath + MonitorSystemVpnCertificateLocalImportPath
	req.Params = *params

	res, err := request.CreateUpdate(c.config, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
