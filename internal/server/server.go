package server

import (
	"Core/config"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	startUrl string
	cfg      *config.Config
	restyCl  *resty.Client
}

func NewServer(
	url string,
	config *config.Config,
	app *resty.Client,
) *Server {
	return &Server{
		startUrl: url,
		cfg:      config,
		restyCl:  app,
	}
}

func (s *Server) Run() {
	s.initLogger()

	err := s.checkUrl()
	if err != nil {
		log.Error(err)
		return
	}

	err = s.MapHandlers()
	if err != nil {
		log.Error(err)
		return
	}
}
