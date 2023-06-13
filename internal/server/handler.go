package server

import (
	log "github.com/sirupsen/logrus"
	"time"
)

func (s *Server) initLogger() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339, // Set your desired timestamp format here.
	})
}
