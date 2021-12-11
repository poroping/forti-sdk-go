package monitor

import (
	"github.com/poroping/forti-sdk-go/v2/config"
)

func New(config config.Config) Endpoints {
	return &Client{config: config}
}

type Client struct {
	config config.Config
}

type Endpoints interface {
}
