package db

import (
	"time"

	"github.com/eflem00/go-example-app/common"
	"github.com/eflem00/go-example-app/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "gorm.io/gorm/logger"
)

func RetryConnect(settings *common.Settings, config *gorm.Config, logger *common.Logger) *gorm.DB {
	db, err := gorm.Open(postgres.Open(settings.PgConnectionString), config)

	if err != nil {
		logger.Err(err, "Failed to connect to db, sleeping and retying in 5 seconds...")
		time.Sleep(time.Second * 5)
		return RetryConnect(settings, config, logger)
	}

	return db
}

func NewDb(settings *common.Settings, lgr *common.Logger) *gorm.DB {
	config := gorm.Config{
		// Logger: logger.Default.LogMode(logger.Silent),
		FullSaveAssociations: true,
	}

	db := RetryConnect(settings, &config, lgr)

	if settings.IsDev() {
		lgr.Info("automigrating DB...")

		db.AutoMigrate(&entities.Profile{})
		db.AutoMigrate(&entities.Element{})
	}

	return db
}
