package main

import (
	"auth_app/internal/app/auth"
	"auth_app/internal/app/permission"
	"auth_app/internal/app/role"
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

	authService := auth.NewAuthService(repository.DB)
	roleService := role.NewRoleService(repository.DB)
	permissionService := permission.NewPermissionService(repository.DB)

	r := chi.NewRouter()

	handler.RegisterAuthHandlers(r, authService)
	handler.RegisterRoleHandlers(r, roleService)
	handler.RegisterPermissionHandlers(r, permissionService)

	log.Printf("Server running on %s", cfg.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), r)
	if err != nil {
		panic(err)
	}

}
