package handlers

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/omekov/online-market/backend/models"
)

func categoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strID := vars["id"]
	id, err := strconv.ParseInt(strID, 32, 32)
	if err != nil {
		ERROR(w, http.StatusBadRequest, err)
		return
	}
	category := models.Category{}
	err = models.GetCategory(int32(id), &category)
	if err != nil {
		if err == sql.ErrNoRows {
			ERROR(w, http.StatusNotFound, err)
			return
		}
		ERROR(w, http.StatusInternalServerError, err)
		return
	}
	JSON(w, http.StatusOK, category)
}
func categoriesHandler(w http.ResponseWriter, r *http.Request) {
	var categories []models.Category
	err := models.GetCategories(&categories)
	if err != nil {
		ERROR(w, http.StatusInternalServerError, err)
		return
	}
	JSON(w, http.StatusOK, categories)
}
func createCategoryHandler(w http.ResponseWriter, r *http.Request) {
	category := models.Category{}
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		if err == io.EOF {
			ERROR(w, http.StatusNoContent, err)
			return
		}
		ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = models.SaveCategory(&category)
	if err != nil {
		ERROR(w, http.StatusInternalServerError, err)
		return
	}
	JSON(w, http.StatusCreated, category)
}
func updateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		ERROR(w, http.StatusBadRequest, err)
		return
	}
	category := models.Category{}
	err = json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		if err == io.EOF {
			ERROR(w, http.StatusNoContent, err)
			return
		}
		ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = models.UpdateCategory(int32(id), &category)
	if err != nil {
		if err == sql.ErrNoRows {
			ERROR(w, http.StatusNotFound, err)
			return
		}
		ERROR(w, http.StatusInternalServerError, err)
		return
	}
	JSON(w, http.StatusOK, category)
}
func deleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = models.DeleteCategory(int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			ERROR(w, http.StatusNotFound, err)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	JSON(w, http.StatusOK, id)
}
