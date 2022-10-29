package app

import (
	"context"
	"fmt"
	"net"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func (a *Application) runGRPCServer(ctx context.Context, grpcServer *grpc.Server) error {
	go func() {
		<-ctx.Done()
		a.logger.Info("grpc: initiated graceful server stop")
		grpcServer.GracefulStop()
		a.logger.Info("grpc: server gracefully stoped")
	}()

	select {
	case <-ctx.Done():
		a.logger.Error("grpc: not started, context already done")
		return ctx.Err()
	default:
		a.logger.Info(fmt.Sprintf("grpc: starting server on port %d", a.config.GRPC.Port))
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.config.GRPC.Port))
		if err != nil {
			return errors.WithMessage(err, "grpc")
		}
		return errors.WithMessage(grpcServer.Serve(lis), "grpc")
	}
}
