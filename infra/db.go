package infra

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/kenji-otomo/AppPurchaseBudget/config"
	"github.com/kenji-otomo/AppPurchaseBudget/repository"
	_ "github.com/lib/pq"
)

func DBOpen() error {

	cfg, err := config.Load()
	if err != nil {
		return err
	}

	db, err := gorm.Open(postgres.Open(cfg.Dns), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	// 接続が有効であるか確認する
	pingErr := sqlDB.Ping()
	if pingErr != nil {
		return pingErr
	}

	repository.NewDB(db)

	return nil
}
