package resources

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"                   // mysql driver initialization
	_ "github.com/golang-migrate/migrate/v4/source/file" // local file system driver for migrations initialization
	"go.uber.org/zap"
	"gorm.io/gorm"

	"imperial-fleet-inventory/services/spaceship/domain/model"
)

func ProvisionDatabase(host, name, username, password string, logger *zap.Logger) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/?charset=utf8mb4&parseTime=True&loc=UTC", username, password, host)
	var db *sql.DB
	var err error

	// Checks if we can connect to the db
	trySQLConn := func() bool {
		db, err = sql.Open("mysql", connectionString)
		return err == nil
	}

	// Give it a couple of tries, sometimes the DB isn't really ready yet
	retries := 10
	delay := 1 * time.Second
	for i := 0; i < retries; i++ {
		if trySQLConn() {
			break
		}

		time.Sleep(delay)
	}

	// Trying one last time, and logging if we failed
	if db == nil {
		db, err = sql.Open("mysql", connectionString)
		if err != nil {
			logger.Error("Failed connect to database!", zap.Error(err))
			return
		}
	}

	defer db.Close()

	_, err = db.Query(fmt.Sprintf("USE %s", name))
	if err == nil {
		logger.Info(fmt.Sprintf("Database '%s' already exist!", name))
		return
	}

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", name))
	if err != nil {
		logger.Error(fmt.Sprintf("Failed create '%s' database!", name), zap.Error(err))
		return
	}

	logger.Info(fmt.Sprintf("Successfully created '%s' database..", name))

}

func InitDB(host string, port int, name string, user string, password string) (*sql.DB, error) {
	DSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", user, password, host, port, name)
	db, err := sql.Open("mysql", DSN)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func MigrateDB(db *gorm.DB, logger *zap.Logger) error {

	err := db.AutoMigrate(&model.Spaceship{}, &model.SpaceshipArmament{})
	if err != nil {
		logger.Warn("Auto-migration finished with err", zap.Error(err))
	}

	return nil
}
