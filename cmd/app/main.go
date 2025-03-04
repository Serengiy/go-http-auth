package main

import (
	"auth_app/internal/config"
	"auth_app/internal/http/handler"
	"auth_app/internal/http/validators"
	"auth_app/internal/repository"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	cfg := config.MustLoad()

	repository.InitDB(cfg)
	validators.InitiateValidator()
	r := chi.NewRouter()

	handler.RegisterAppHandlers(r)

	log.Printf("Server running on %s", cfg.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), r)
	if err != nil {
		panic(err)
	}

}
