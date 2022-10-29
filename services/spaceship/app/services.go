package app

import (
	"context"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	spaceshipProto "imperial-fleet-inventory/api/langs/go/spaceship/grpc"

	grpcRequestMetadata "imperial-fleet-inventory/common/request_metadata/grpc"
	"imperial-fleet-inventory/services/spaceship/api"
	"imperial-fleet-inventory/services/spaceship/repository"
	"imperial-fleet-inventory/services/spaceship/spaceship"
)

func (a *Application) initServices(
	grpCtx context.Context,
	errGrp *errgroup.Group,
) {
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(grpcRequestMetadata.NewGRPCServerInterceptor()))
	spaceshipRepo := repository.NewSpaceshipRepo(a.db)
	spaceshipService := spaceship.NewService(spaceshipRepo, a.logger)
	spaceshipProto.RegisterSpaceshipServiceServer(grpcServer, api.NewSpaceshipServer(spaceshipService))

	errGrp.Go(func() error {
		return a.runGRPCServer(grpCtx, grpcServer)
	})
}
