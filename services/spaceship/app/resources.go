package app

import (
	"strconv"

	"go.uber.org/zap"

	"imperial-fleet-inventory/services/spaceship/resources"
)

func (a *Application) initDB() error {
	port, err := strconv.Atoi(a.config.DB.Port)
	if err != nil {
		return err
	}

	resources.ProvisionDatabase(a.config.DB.Host, a.config.DB.Database, a.config.DB.Username, a.config.DB.Password, a.logger)

	db, err := resources.InitDB(a.config.DB.Host, port, a.config.DB.Database, a.config.DB.Username, a.config.DB.Password)
	if err != nil {
		return err
	}

	gormDB, err := resources.InitGORM(db, a.logger)
	if err != nil {
		return err
	}

	if err := resources.MigrateDB(gormDB, a.logger); err != nil {
		return err
	}

	closeFn := func() {
		if err := db.Close(); err != nil {
			a.logger.Error("db close error", zap.Error(err))
		}
	}

	a.db = gormDB
	a.dbClosFn = closeFn
	return nil
}
