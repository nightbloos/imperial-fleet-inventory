package resources

import (
	"database/sql"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"imperial-fleet-inventory/common/orm"
)

// InitGORM initializes gorm connection
func InitGORM(db *sql.DB, logger *zap.Logger) (*gorm.DB, error) {
	return gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{Logger: orm.NewGormLogger(logger)})
}
