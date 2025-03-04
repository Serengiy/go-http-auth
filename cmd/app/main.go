package main

import (
	"auth_app/internal/config"
	"auth_app/internal/http/handler"
	"auth_app/internal/http/validators"
	"auth_app/internal/repository"
	"auth_app/internal/service"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	cfg := config.MustLoad()

	repository.InitDB(cfg)
	validators.InitiateValidator()

	authService := service.NewAuthService(repository.NewAuthRepository(repository.DB))
	permissionService := service.NewPermissionService(repository.NewPermissionRepository(repository.DB))
	roleService := service.NewRoleService(repository.NewRoleRepository(repository.DB))

	r := chi.NewRouter()

	handler.RegisterAuthHandlers(r, authService)
	handler.RegisterPermissionHandlers(r, permissionService)
	handler.RegisterRoleHandler(r, roleService, permissionService)

	log.Printf("Server running on %s", cfg.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), r)
	if err != nil {
		panic(err)
	}

}
