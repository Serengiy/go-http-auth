package handler

import (
	"auth_app/internal/dto"
	"auth_app/internal/packages/response"
	"auth_app/internal/service"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func registerRoleHandler(r chi.Router, roleService *service.RoleService, permissionService *service.PermissionService) {
	r.Post("/roles", func(w http.ResponseWriter, r *http.Request) {
		const op = "create new role"
		var reqBody dto.RoleRequest

		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			http.Error(w, "Unsuccessful decoding", http.StatusBadRequest)
			return
		}

		role, err := roleService.CreateNewRole(reqBody)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = response.JSON(w, 201, role, "")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	r.Delete("/roles/{role_id}", func(w http.ResponseWriter, r *http.Request) {
		const op = "delete role"

		idParam := chi.URLParam(r, "role_id")
		roleId, err := strconv.Atoi(idParam)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		role, err := roleService.FindRoleByID(int64(roleId))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
		}

		err = roleService.DeleteRole(role)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = response.JSON(w, 201, "", "Role deleted")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	r.Post("/roles/{role_id}/add-permission/{permission_id}", func(w http.ResponseWriter, r *http.Request) {
		const op = "add permission"
		roleIdParam := chi.URLParam(r, "role_id")
		roleId, err := strconv.Atoi(roleIdParam)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		permissionParam := chi.URLParam(r, "permission_id")
		permissionId, err := strconv.Atoi(permissionParam)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		role, err := roleService.FindRoleByID(int64(roleId))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			http.Error(w, "Internal error", http.StatusInternalServerError)
		}

		permission, err := permissionService.GetPermissionById(int64(permissionId))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			http.Error(w, "Internal error", http.StatusInternalServerError)
		}

		err = roleService.AttachPermission(role, permission)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = response.JSON(w, 201, role, "")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
