package usecase

import (
	"net/http"

	"github.com/kenji-otomo/AppPurchaseBudget/repository"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello World!"))
}

func GetApps() ([]*repository.AppDao, error) {

	results, err := repository.GetApps()

	return results, err
}

func CreateApp(name string) error {

	app := &repository.AppDao{
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
