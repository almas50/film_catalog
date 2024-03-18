package main

import (
	"FilmCatalog/internal/handler"
	"FilmCatalog/internal/service"
	"FilmCatalog/internal/storage/postgres"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"net/http"
	"time"

	_ "FilmCatalog/docs"
)

type config struct {
	db_name     string `env:"POSTGRES_DB"`
	db_user     string `env:"POSTGRES_USER"`
	db_password string `env:"POSTGRES_PASSWORD"`
}

// @title Swagger FilmCatalog API
// @version 1.0
// @description Catalog of films and actors

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @securityDefinitions.basic BasicAuth
// @in header

// @host localhost:8080
// @BasePath /

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	var validate *validator.Validate
	validate = validator.New()

	mux := http.NewServeMux()
	s, _ := postgres.NewStorage()
	service, _ := service.NewService(s)
	Handler := handler.Handler{Service: service, Validator: validate}
	Handler.Register(mux)

	mux.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	serv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}
	log.Fatal(serv.ListenAndServe())
}
