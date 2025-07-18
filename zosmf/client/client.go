package client

import (
	"zosmf/config"
	"zosmf/ds"
	"zosmf/jobs"
	"zosmf/transport"
)

// Client centraliza o acesso aos serviços do SDK z/OSMF.
type Client struct {
	Config    *config.Config
	Transport *transport.Transport
	DS        *ds.Service
	Job       *jobs.Service
}

// New cria uma nova instância de Client com os serviços inicializados.
func New(cfg *config.Config) *Client {
	tr := transport.New(cfg)
	return &Client{
		Config:    cfg,
		Transport: tr,
		DS:        ds.New(tr),
		Job:       jobs.New(tr),
	}
}
