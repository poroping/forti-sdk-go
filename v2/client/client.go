package client

import (
	"crypto/tls"
	"log"
	"net/http"
	"time"

	"github.com/poroping/forti-sdk-go/v2/auth"
	"github.com/poroping/forti-sdk-go/v2/cmdb"
	"github.com/poroping/forti-sdk-go/v2/config"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/forti-sdk-go/v2/monitor"
	"github.com/poroping/forti-sdk-go/v2/request"
	"github.com/poroping/forti-sdk-go/v2/utils"
)

type FortiSDKClient struct {
	Config  config.Config
	Cmdb    cmdb.Endpoints
	Monitor monitor.Endpoints
	Utils   utils.Endpoints
}

func NewClient(auth *auth.Auth) (*FortiSDKClient, error) {
	c, err := NewClientBase(auth)
	if err != nil {
		return nil, err
	}

	c.Cmdb = cmdb.New(&c.Config)
	c.Monitor = monitor.New(&c.Config)
	c.Utils = utils.New(&c.Config)

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
	dc.Timeout = 30 * time.Second

	c := &FortiSDKClient{}
	c.Config.Auth = auth
	// c.Config.HTTPCon = retryClient.StandardClient()
	c.Config.HTTPCon = dc
	c.Config.FwTarget = auth.Hostname
	c.Config.Fv = ""
	c.Config.UserAgent = "FortiOS Go-SDK"
	c.Config.Retries = 5

	if !auth.AutoVersion {
		log.Print("[INFO] Skipping FortiOS version check.")
	} else {
		log.Print("[INFO] Attempting to determine FortiOS version.")
		req := &models.CmdbRequest{}
		req.HTTPMethod = "GET"
		req.NoVdom = true
		req.Path = models.MonitorBasePath + "system/status/"
		res, err := request.Read(&c.Config, req)
		if err != nil || res == nil {
			log.Print("[WARN] Error attempting to determine FortiOS version, this check can be skipped.")
			return c, err
		}
		if res.Version != nil {
			log.Printf("[INFO] FortiOS version detected as %s", *res.Version)
			c.Config.Fv = *res.Version
		}
	}

	return c, nil
}
