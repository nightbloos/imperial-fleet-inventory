package orm

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm/logger"
)

// GormLogger is a logger for Gorm
type GormLogger struct {
	logger       *zap.Logger
	debugEnabled bool
}

// NewGormLogger creates new GormLogger
func NewGormLogger(logger *zap.Logger) logger.Interface {
	return &GormLogger{
		logger:       logger,
		debugEnabled: os.Getenv("LOG_DB_DEBUG_ENABLED") == "1",
	}
}

// LogMode log mode
func (l *GormLogger) LogMode(_ logger.LogLevel) logger.Interface {
	return l
}

// Info log info
func (l *GormLogger) Info(_ context.Context, s string, i ...interface{}) {
	l.logger.Info(fmt.Sprintf(s, i...))
}

// Warn log warn
func (l *GormLogger) Warn(_ context.Context, s string, i ...interface{}) {
	l.logger.Warn(fmt.Sprintf(s, i...))
}

// Error log error
func (l *GormLogger) Error(_ context.Context, s string, i ...interface{}) {
	l.logger.Error(fmt.Sprintf(s, i...))
}

// Trace log trace
func (l *GormLogger) Trace(_ context.Context, begin time.Time, fc func() (string, int64), err error) {
	if !l.debugEnabled {
		return
	}
	queryString, rowsAffected := fc()
	l.logger.Debug(fmt.Sprintf("sql:%s, rows affetted: %d, begin: %s, error: %s", queryString, rowsAffected, begin.String(), err))
}
