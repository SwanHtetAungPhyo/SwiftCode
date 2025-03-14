package repo

import (
	"github.com/SwanHtetAungPhyo/swifcode/internal/model"
	"github.com/SwanHtetAungPhyo/swifcode/internal/pkg/logging"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func ModelMigration(db *gorm.DB) error {
	logging.Logger.Info("starting migration", zap.String("migration", "migration"))
	err := db.AutoMigrate(&model.SwiftCode{})
	if err != nil {
		logging.Logger.Error("error migrating models", zap.Error(err))
		return err
	}
	logging.Logger.Info("migration complete", zap.String("migration", "migration"))
	return nil
}
