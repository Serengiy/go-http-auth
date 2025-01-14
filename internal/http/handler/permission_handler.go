package handler

import (
	"auth_app/internal/dto"
	"auth_app/internal/packages/pagination"
	"auth_app/internal/service"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func RegisterPermissionHandlers(r chi.Router, permissionService *service.PermissionService) {
	r.Post("/permissions", func(w http.ResponseWriter, r *http.Request) {
		var reqBody dto.PermissionStoreRequest

		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		permission, err := permissionService.CreateNewPermission(&reqBody)
		if err != nil {
			var vErr service.ValidationError
			if errors.As(err, &vErr) {
				http.Error(w, vErr.Error(), http.StatusBadRequest)
				return
			}
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(permission); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	r.Get("/permissions", func(w http.ResponseWriter, r *http.Request) {

		pageParam := r.URL.Query().Get("page")
		perPageParam := r.URL.Query().Get("perPage")
		nameParam := r.URL.Query().Get("name")

		page, err := pagination.Paginate(pageParam)
		if err != nil {
			page = 1
		}

		perPage, err := pagination.PerPageNumber(perPageParam)
		if err != nil {
			perPage = 10
		}

		filters := dto.PermissionFilterStruct{
			Name:    nameParam,
			PerPage: perPage,
			Page:    page,
		}
		fmt.Printf("name: %s, perpage: %v, page: %v \n", filters.Name, filters.PerPage, filters.Page)

		permissions, err := permissionService.GetPermissionsByFilter(filters)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(permissions); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
