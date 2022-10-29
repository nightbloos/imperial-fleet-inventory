package app

import (
	"context"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

// Application is the main application struct
type Application struct {
	config   Config
	db       *gorm.DB
	dbClosFn func()
	logger   *zap.Logger
}

// NewApplication creates a new application
func NewApplication() *Application {
	return &Application{}
}

// Run runs the application
func (a *Application) Run(ctx context.Context) error {

	if err := a.initConfig(); err != nil {
		return err
	}
	if err := a.initLogger(); err != nil {
		return err
	}
	if err := a.initDB(); err != nil {
		return err
	}
	defer a.dbClosFn()

	errGrp, ctx := errgroup.WithContext(ctx)

	a.initServices(ctx, errGrp)

	return errGrp.Wait()
}
