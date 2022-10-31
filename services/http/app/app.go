package app

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	spaceshipProto "imperial-fleet-inventory/api/langs/go/spaceship/grpc"
	"imperial-fleet-inventory/common/http"
	grpcRequestMetadata "imperial-fleet-inventory/common/request_metadata/grpc"
	"imperial-fleet-inventory/services/http/api"
)

type Application struct {
	config Config
	logger *zap.Logger
}

func NewApplication() *Application {
	return &Application{}
}

func (a *Application) Run(ctx context.Context) error {
	rand.Seed(time.Now().Unix())

	if err := a.initConfig(); err != nil {
		return err
	}
	if err := a.initLogger(); err != nil {
		return err
	}

	ginRouter := a.initGinRouter()

	spaceshipService, err := a.initSpaceshipClient(ctx)
	if err != nil {
		return err
	}
	spaceshipAPI := api.NewSpaceshipServer(spaceshipService, a.logger)
	spaceshipAPI.Register(ginRouter)

	errGrp, ctx := errgroup.WithContext(ctx)

	errGrp.Go(func() error {
		return http.NewServer(a.config.HTTP.Port, a.logger).Run(ctx, ginRouter)
	})

	return errGrp.Wait()
}

func (a *Application) initSpaceshipClient(ctx context.Context) (spaceshipProto.SpaceshipServiceClient, error) {
	conn, err := grpc.DialContext(
		ctx,
		fmt.Sprintf("%s:%d", a.config.SpaceshipService.Host, a.config.SpaceshipService.Port),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpcRequestMetadata.NewClientInterceptor()))
	if err != nil {
		return nil, err
	}
	return spaceshipProto.NewSpaceshipServiceClient(conn), err
}
