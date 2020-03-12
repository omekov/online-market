package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Router ...
func Router() *mux.Router {
	mux := mux.NewRouter()
	mux.HandleFunc("/origins", originHandler)
	mux.HandleFunc("/categories", categoriesHandler).Methods("GET")
	mux.HandleFunc("/categories", createCategoryHandler).Methods("POST")
	mux.HandleFunc("/categories/{id:[0-9]+}", categoryHandler).Methods("GET")
	mux.HandleFunc("/categories/{id:[0-9]+}", updateCategoryHandler).Methods("PUT")
	mux.HandleFunc("/categories/{id:[0-9]+}", deleteCategoryHandler).Methods("DELETE")
	return mux
}

// JSON ...
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

// ERROR ...
func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}
