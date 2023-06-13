package server

import (
	"Core/internal/product/delivery/http"
	"fmt"
	log "github.com/sirupsen/logrus"
	"regexp"
	"time"
)

func (s *Server) MapHandlers() (err error) {
	HTTPService := http.NewHttpService(s.cfg, s.restyCl)
	_, err = HTTPService.GetRequestBody(s.startUrl)
	if err != nil {
		return
	}
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
