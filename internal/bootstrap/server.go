package bootsrap

import (
	"context"
	"net/http"

	"github.com/zmonsoonz/go-test-api/internal/platform/config"
)

type Server struct {
    httpServer *http.Server
}

func (s *Server) Run(cfg config.HttpConfig, handler http.Handler) error{
	s.httpServer = &http.Server{
		Addr:         ":" + cfg.Port,
		Handler: handler,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
    return s.httpServer.Shutdown(ctx)
}