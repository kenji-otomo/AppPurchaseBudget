package repository

import (
	"time"

	"gorm.io/gorm"
)

type App struct {
	ID             *int64
	Name           string
	LastPurchaseAt time.Time
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
}

func (g *App) Create(tx *gorm.DB) error {
	g.LastPurchaseAt = time.Now()
	return tx.Create(&g).Error
}

func GetApps() ([]*App, error) {

	apps := []*App{}
	err := db.Find(&apps).Error
	if err != nil {
		return nil, err
	}

	return apps, nil
}

func FetchAppByName(name string) (*App, error) {
	app := &App{}
	err := db.Where("name = ?", name).First(app).Error

	return app, err
}
