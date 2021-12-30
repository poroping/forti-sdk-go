package client

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/poroping/forti-sdk-go/v2/auth"
	"github.com/poroping/forti-sdk-go/v2/cmdb"
	"github.com/poroping/forti-sdk-go/v2/config"
	"github.com/poroping/forti-sdk-go/v2/monitor"
)

type FortiSDKClient struct {
	Config  config.Config
	Cmdb    cmdb.Endpoints
	Monitor monitor.Endpoints
}

func NewClient(auth *auth.Auth) (*FortiSDKClient, error) {
	c, _ := NewClientBase(auth)

	c.Cmdb = cmdb.New(&c.Config)
	c.Monitor = cmdb.New(&c.Config)

	return c, nil
}

func NewClientBase(auth *auth.Auth) (*FortiSDKClient, error) {
	insecure := auth.Insecure

	// retryClient := retryablehttp.NewClient()
	// retryClient.RetryMax = 6
	// retryClient.HTTPClient.Transport = tr
	// retryClient.HTTPClient.Timeout = 10 * time.Second

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: insecure},
	}

	dc := &http.Client{}
	dc.Transport = tr
	dc.Timeout = 10 * time.Second

	c := &FortiSDKClient{}
	c.Config.Auth = auth
	// c.Config.HTTPCon = retryClient.StandardClient()
	c.Config.HTTPCon = dc
	c.Config.FwTarget = auth.Hostname
	c.Config.Fv = ""
	c.Config.UserAgent = "FortiOS Go-SDK"
	c.Config.Retries = 5

	return c, nil
}
