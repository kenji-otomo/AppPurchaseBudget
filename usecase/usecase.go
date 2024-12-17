package usecase

import (
	"github.com/kenji-otomo/AppPurchaseBudget/domain/history"
	"github.com/kenji-otomo/AppPurchaseBudget/repository"
	"gorm.io/gorm"
)

func GetApps() ([]*repository.App, error) {

	results, err := repository.GetApps()

	return results, err
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

func CreateApp(name string) error {

	app := &repository.App{
		Name: name,
	}

	tx := repository.BeginTransaction()

	if err := app.Create(tx); err != nil {
		tx.Rollback()
		return err
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
