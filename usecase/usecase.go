package usecase

import (
	"time"

	"github.com/kenji-otomo/AppPurchaseBudget/domain/app"
	"github.com/kenji-otomo/AppPurchaseBudget/domain/history"
	"github.com/kenji-otomo/AppPurchaseBudget/repository"
	"gorm.io/gorm"
)

func GetApps() ([]*repository.App, error) {

	now := time.Now()
	year, month, _ := now.Date()

	// 月初
	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	// 月末
	lastDay := firstDay.AddDate(0, 1, 0).Add(-1 * time.Second)

	results, err := repository.GetAppsOrderByAmount(firstDay, lastDay)

	r := []*repository.App{}
	for _, v := range results {
		r = append(r, &v.App)
	}

	return r, err
}

func CheckDuplicateApp(name string) (*repository.App, error) {

	app, err := repository.FetchAppByName(name)
	if err != nil {

		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return app, err
}

func CreateApp(name string) (*repository.App, error) {

	app := &repository.App{
		Name: name,
	}

	tx := repository.BeginTransaction()

	if err := app.Create(tx); err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return app, nil
}

func UpdateAppName(reqs []*app.UpdateAppRequest) error {

	tx := repository.BeginTransaction()

	for _, req := range reqs {
		rApp := repository.App{
			ID:   req.ID,
			Name: req.Name,
		}

		if err := rApp.UpdateName(tx); err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

func GetHistories() ([]*repository.PurchaseHistory, error) {
	results, err := repository.ListPurchaseHistories()

	return results, err
}

func CreateHitory(h history.HistoryRequest) error {

	// エンティティに変換
	his := history.NewHistory(h.AppID, h.Amount, h.PurchaseDate)

	// DB用に変換
	hisRep := repository.NewPurchaseHistory(his)

	tx := repository.BeginTransaction()

	// 登録
	if err := hisRep.Create(tx); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func FetchHistoryData() ([]*repository.AppWithSum, error) {

	now := time.Now()
	year, month, _ := now.Date()

	// 月初
	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	// 月末
	lastDay := firstDay.AddDate(0, 1, 0).Add(-1 * time.Second)

	data, err := repository.FetchPurchaseData(firstDay, lastDay)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func FetchBudgetByType(t int64) (*repository.Budget, error) {
	bt := repository.BudgetType(t)
	budget, err := repository.FetchBudgetByType(bt)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	year, month, _ := now.Date()

	// 月初
	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	// 月末
	lastDay := firstDay.AddDate(0, 1, 0).Add(-1 * time.Second)

	totalAmount, err := repository.FetchTotalPurchaseAmount(firstDay, lastDay)
	if err != nil {
		return nil, err
	}

	if totalAmount != nil {
		budget.Amount -= *totalAmount
	}

	return budget, err
}
