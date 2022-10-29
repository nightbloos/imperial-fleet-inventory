package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

type Server struct {
	port   int
	logger *zap.Logger
}

func NewServer(port int, logger *zap.Logger) *Server {
	return &Server{port: port, logger: logger}
}

func (s *Server) Run(ctx context.Context, handler http.Handler) error {
	srv := &http.Server{Addr: fmt.Sprintf(":%d", s.port), Handler: handler}

	shutdownCompletedCh := make(chan struct{})
	go func() {
		<-ctx.Done()
		s.logger.Info("http: initiated graceful server stop")
		if err := srv.Shutdown(context.Background()); err != nil {
			s.logger.Error("http: server forced to shutdown", zap.Error(err))
		}
		close(shutdownCompletedCh)
	}()

	s.logger.Info(fmt.Sprintf("http: starting server on port %d", s.port))
	err := srv.ListenAndServe()
	if err != nil && errors.Is(err, http.ErrServerClosed) {
		<-shutdownCompletedCh
		s.logger.Info("http: server stopped")
		return nil
	}
	return err
}
