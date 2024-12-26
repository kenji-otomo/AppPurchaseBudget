package usecase

import (
	"time"

	"github.com/kenji-otomo/AppPurchaseBudget/domain/app"
	"github.com/kenji-otomo/AppPurchaseBudget/domain/budget"
	"github.com/kenji-otomo/AppPurchaseBudget/domain/history"
	"github.com/kenji-otomo/AppPurchaseBudget/repository"
	"gorm.io/gorm"
)

// 課金対象一覧取得（課金登録画面にて使用）
func GetApps() ([]*app.App, error) {

	now := time.Now()
	firstDay, lastDay := generateDataRangeForShowApp(now)

	results, err := repository.GetAppsOrderByAmount(firstDay, lastDay)

	r := []*app.App{}
	for _, v := range results {
		r = append(r, app.NewApp(v.ID, v.Name, &v.LastPurchaseAt))
	}

	return r, err
}

// 課金名が重複していないかチェック
func CheckDuplicateApp(name string) (*app.App, error) {

	rApp, err := repository.FetchAppByName(name)
	if err != nil {

		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return app.NewApp(rApp.ID, rApp.Name, &rApp.LastPurchaseAt), err
}

// 課金対象登録
func CreateApp(name string) (*app.App, error) {

	rApp := &repository.App{
		Name: name,
	}

	tx := repository.BeginTransaction()

	if err := rApp.Create(tx); err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return app.NewApp(rApp.ID, rApp.Name, &rApp.LastPurchaseAt), nil
}

// 課金対象の名称変更
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

// 課金明細の一覧取得
func GetHistories() ([]*history.History, error) {

	now := time.Now()
	firstDay, lastDay := generateDataRangeForShowApp(now)

	results, err := repository.ListPurchaseHistories(firstDay, lastDay)

	histories := []*history.History{}
	for _, r := range results {
		histories = append(histories, history.NewHistory(r.ID, r.AppID, r.Amount, r.PurchaseDate, &r.Name))
	}

	return histories, err
}

// 課金明細登録
func CreateHitory(h history.HistoryRequest) error {

	// エンティティに変換
	his := history.NewHistory(nil, h.AppID, h.Amount, h.PurchaseDate, nil)

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

// 1ヶ月間の課金明細を取得
func FetchHistoryData() ([]*app.AppWithSum, error) {

	now := time.Now()
	firstDay, lastDay := generateDataRange(now)

	data, err := repository.FetchPurchaseData(firstDay, lastDay)
	if err != nil {
		return nil, err
	}

	results := []*app.AppWithSum{}
	for _, d := range data {
		results = append(results, app.NewAppWithSum(app.NewApp(d.App.ID, d.App.Name, &d.App.LastPurchaseAt), d.Amount))
	}

	return results, nil
}

// タイプを指定して課金設定額情報の取得
func FetchBudgetByType(t int64) (*budget.Budget, error) {
	bt := budget.BudgetType(t)
	rBudget, err := repository.FetchBudgetByType(bt)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	firstDay, lastDay := generateDataRange(now)

	totalAmount, err := repository.FetchTotalPurchaseAmount(firstDay, lastDay)
	if err != nil {
		return nil, err
	}

	if totalAmount != nil {
		rBudget.Amount -= *totalAmount
	}

	return budget.NewBudget(rBudget.ID, rBudget.Type, rBudget.Amount), nil
}
