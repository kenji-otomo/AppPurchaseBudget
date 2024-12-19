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

type AppWithSum struct {
	App
	Amount int64
}

func (a *App) Create(tx *gorm.DB) error {
	a.LastPurchaseAt = time.Now()
	return tx.Create(&a).Error
}

func (a *App) UpdateName(tx *gorm.DB) error {
	now := time.Now()
	return tx.Model(&a).Updates(App{Name: a.Name, UpdatedAt: &now}).Error
}

func GetApps() ([]*App, error) {
	apps := []*App{}
	err := db.Find(&apps).Error
	return apps, err
}

func FetchAppByName(name string) (*App, error) {
	app := &App{}
	err := db.Where("name = ?", name).First(app).Error

	return app, err
}

func FetchPurchaseData(start, end time.Time) ([]*AppWithSum, error) {
	appSums := []*AppWithSum{}

	err := db.Select("a.*, SUM(p.amount) amount").
		Table("apps a").
		Joins("INNER JOIN purchase_histories p ON a.id = p.app_id ").
		Where("p.created_at BETWEEN ? AND ?", start, end).
		Group("a.id").Limit(10).Order("amount DESC").Scan(&appSums).Error

	return appSums, err
}

func GetAppsOrderByAmount(start, end time.Time) ([]*AppWithSum, error) {
	appSums := []*AppWithSum{}

	err := db.Raw(`
	SELECT 
		a.* , 
  		CASE 
    		WHEN SUM(p.amount) IS NULL THEN 0
    		ELSE SUM(p.amount)
  		END as amount 
	FROM apps a
	LEFT OUTER JOIN(
  		SELECT * FROM
  		purchase_histories 
  		WHERE created_at BETWEEN ? AND ?
	) p ON a.id = p.app_id
	GROUP BY a.id 
	ORDER BY amount DESC;
	`, start, end).Scan(&appSums).Error

	return appSums, err
}
