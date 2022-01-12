package monitor

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/forti-sdk-go/v2/request"
)

type MonitorSystemVpnCertificateRemoteImport struct {
	Certificate string `json:"file_content"`
	Scope       string `json:"scope"`
}

const MonitorSystemVpnCertificateRemoteImportPath = "vpn-certificate/remote/import/"

// Uploads certificate to Fortigate
func (c *Client) MonitorSystemVpnCertificateRemoteImport(payload *MonitorSystemVpnCertificateRemoteImport, params *models.CmdbRequestParams) (*string, error) {
	payload.Certificate = base64.StdEncoding.EncodeToString([]byte(payload.Certificate))

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req := &models.CmdbRequest{}
	req.HTTPMethod = "POST"
	req.Payload = body
	req.Path = models.MonitorBasePath + MonitorSystemVpnCertificateRemoteImportPath
	req.Params = *params

	var before []string
	before, err = c.getRemoteCertName(params)
	if err != nil {
		err = fmt.Errorf("cannot get list of existing cert names: %s", err)
		return nil, err
	}

	_, err = request.CreateUpdate(c.config, req)
	if err != nil {
		return nil, err
	}

	var after []string
	after, err = c.getRemoteCertName(params)
	if err != nil {
		err = fmt.Errorf("cannot get updated list of cert names: %s", err)
		return nil, err
	}

	diff := difference(after, before)

	if len(diff) != 1 {
		err = fmt.Errorf("error finding new cert")
		return nil, err
	}

	return &diff[0], nil
}

func (c *Client) getRemoteCertName(params *models.CmdbRequestParams) (name []string, err error) {
	req := &models.CmdbRequest{}
	req.HTTPMethod = "GET"
	req.Payload = nil
	req.Path = models.CmdbBasePath + "vpn.certificate/remote/"
	req.Params = *params
	filter := []string{"name=@REMOTE_Cert_"}
	format := []string{"name"}
	req.Params.Filter = &filter
	req.Params.Format = &format

	res, err := request.Read(c.config, req)
	if err != nil {
		return nil, err
	}

	names := make([]string, 0)

	if tmp, ok := res.Results.([]interface{}); ok {
		jsontmp, err := json.Marshal(tmp)
		if err != nil {
			return nil, err
		}
		v := []map[string]interface{}{}
		err = json.Unmarshal(jsontmp, &v)
		if err != nil {
			return nil, err
		}
		for _, result := range v {
			name := result["name"].(string)
			names = append(names, name)
		}
		return names, nil
	}
	return names, fmt.Errorf("getRemoteCertName fell through")
}

func difference(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}
