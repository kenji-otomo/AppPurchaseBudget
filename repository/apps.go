package repository

import (
	"time"

	"gorm.io/gorm"
)

type AppDao struct {
	ID             *int64
	Name           string
	LastPurchaseAt *time.Time
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
}

func GetApps() ([]*AppDao, error) {

	apps := []*AppDao{}
	err := db.Find(&apps).Error
	if err != nil {
		return nil, err
	}

	return apps, nil
}

func (g *AppDao) Create(tx *gorm.DB) error {
	return tx.Create(&g).Error
}
