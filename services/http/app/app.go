package app

import (
	"context"
	"math/rand"
	"time"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"imperial-fleet-inventory/common/http"
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

	spaceshipAPI := api.NewSpaceshipServer(a.logger)
	spaceshipAPI.Register(ginRouter)

	errGrp, ctx := errgroup.WithContext(ctx)

	errGrp.Go(func() error {
		return http.NewServer(a.config.HTTP.Port, a.logger).Run(ctx, ginRouter)
	})

	return errGrp.Wait()
}
