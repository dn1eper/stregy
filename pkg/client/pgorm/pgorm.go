package pgorm

import (
	"context"
	"fmt"
	goLog "log"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	dbLog "gorm.io/gorm/logger"
)

func NewLogger(level logrus.Level) dbLog.Interface {
	var dbLevel dbLog.LogLevel
	switch level {
	case log.DebugLevel, log.TraceLevel:
		dbLevel = dbLog.Info
	case log.InfoLevel, log.WarnLevel:
		dbLevel = dbLog.Warn
	case log.ErrorLevel, log.FatalLevel, log.PanicLevel:
		dbLevel = dbLog.Error
	}

	return dbLog.New(
		goLog.New(os.Stdout, "\r\n", goLog.LstdFlags), // io writer
		dbLog.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  dbLevel,     // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
}

func NewClient(ctx context.Context, username, password, host, port, database string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v", host, port, username, password, database)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: NewLogger(log.GetLevel()),
	})
}
