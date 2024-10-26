package server

import (
	"context"
	"crypto/tls"
	"log/slog"
	"net/http"
	"time"

	"github.com/oogway93/golangArchitecture/internal/adapter/config"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(config *config.Container, handler http.Handler) error {
	cer, err := tls.LoadX509KeyPair(config.HTTP.TLSCertPath, config.HTTP.TLSKeyPath)
	if err != nil {
		slog.Warn("Cannot Load tls certification", "error", err.Error())
		return err
	}
	s.httpServer = &http.Server{
		Addr: ":" + config.HTTP.Port,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cer},
		},
		Handler:        handler,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return s.httpServer.ListenAndServeTLS("", "")
}

func (s *Server) ShutDown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
