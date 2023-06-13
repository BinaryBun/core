package server

import (
	"Core/internal/product/delivery/http"
	"Core/internal/product/usecase"
	"fmt"
	log "github.com/sirupsen/logrus"
	"regexp"
	"time"
)

func (s *Server) MapHandlers() (err error) {
	HTTPService := http.NewHttpService(s.cfg, s.restyCl)
	uc := usecase.NewProductUC(s.cfg, HTTPService)

	uc.Print()
	return
}

func (s *Server) checkUrl() error {
	const urlPattern = `^(https?://)?[^\s/$.?#]+\.[^\s/]{2,}(/[^\s]*)?$`

	if s.startUrl == "" {
		s.startUrl = "https://google.com"
		return nil
	}

	match, _ := regexp.MatchString(urlPattern, s.startUrl)
	if !match {
		return fmt.Errorf("incorrect URL: %s", s.startUrl)
	}

	return nil
}

func (s *Server) initLogger() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339, // Set your desired timestamp format here.
	})
}
