package handler

import (
	"auth_app/internal/dto"
	"auth_app/internal/packages/pagination"
	"auth_app/internal/packages/response"
	"auth_app/internal/service"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
)

func registerPermissionHandlers(r chi.Router, permissionService *service.PermissionService) {
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

		permissions, totalRecords, err := permissionService.GetPermissionsByFilter(filters)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		permPagination := &response.Pagination{
			Page:       int64(page),
			PerPage:    int64(perPage),
			TotalPages: pagination.GetTotalPages(totalRecords, perPage),
		}

		err = response.PaginatedJSON(w, 200, permissions, "", permPagination)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	r.Get("/permissions/{id}", func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")

		permID, err := strconv.Atoi(idParam)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		permission, err := permissionService.GetPermissionById(int64(permID))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				return
			}
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		if err := response.JSON(w, 200, permission, ""); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	r.Delete("/permissions/{id}", func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		permID, err := strconv.Atoi(idParam)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		permission, err := permissionService.GetPermissionById(int64(permID))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				return
			}
		}

		err = permissionService.DeletePermissionById(permission)
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		err = response.JSON(w, 200, "", fmt.Sprintf("Permission with ID: %v deleted successfully", idParam))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	r.Put("/permissions/{id}", func(w http.ResponseWriter, r *http.Request) {
		var reqBody dto.PermissionUpdateRequest
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		idParam := chi.URLParam(r, "id")
		permID, err := strconv.Atoi(idParam)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		permission, err := permissionService.GetPermissionById(int64(permID))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				return
			}
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		permission, err = permissionService.UpdatePermission(permission, &reqBody)

		if err != nil {
			if strings.Contains(err.Error(), "already exists") {
				http.Error(w, err.Error(), http.StatusConflict)
				return
			}
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err = response.JSON(w, 200, permission, "")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
