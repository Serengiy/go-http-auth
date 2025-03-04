package handler

import (
	"auth_app/internal/repository"
	"auth_app/internal/service"

	"github.com/go-chi/chi/v5"
)

func RegisterAppHandlers(r chi.Router) {
	authService := service.NewAuthService(repository.NewAuthRepository(repository.DB))
	permissionService := service.NewPermissionService(repository.NewPermissionRepository(repository.DB))
	roleService := service.NewRoleService(repository.NewRoleRepository(repository.DB))
	userService := service.NewUserService(repository.NewUserRepository(repository.DB))

	registerAuthHandlers(r, authService)
	registerPermissionHandlers(r, permissionService)
	registerRoleHandler(r, roleService, permissionService)
	registerUserHandler(r, userService)
}
