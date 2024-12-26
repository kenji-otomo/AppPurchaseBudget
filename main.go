package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/kenji-otomo/AppPurchaseBudget/config"
	"github.com/kenji-otomo/AppPurchaseBudget/controller"
	"github.com/kenji-otomo/AppPurchaseBudget/infra"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("環境変数の読み取りに失敗しました", err)
	}
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("環境変数の生成に失敗しました", err)
	}

	if err := infra.DBOpen(); err != nil {
		log.Fatal("DB接続に失敗しました", err)
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{cfg.VueURL},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
	}))

	// ルーティング
	controller.Route(r)

	fmt.Println("接続OK")
	http.ListenAndServe(":8080", r)
}
