package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/kenji-otomo/AppPurchaseBudget/controller"
	"github.com/kenji-otomo/AppPurchaseBudget/infra"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("環境変数の読み取りに失敗しました", err)
	}

	err := infra.DBOpen()
	if err != nil {
		log.Fatal("DB接続に失敗しました", err)
	}

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
	}))

	// ルーティング
	controller.Route(r)

	fmt.Println("接続OK")
	http.ListenAndServe(":8000", r)
}
