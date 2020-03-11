package handlers

import (
	"github.com/gorilla/mux"
)

// Router ...
func Router() *mux.Router {
	mux := mux.NewRouter()
	mux.HandleFunc("/origins", originHandler)
	mux.HandleFunc("/categories/{id:[0-9]+}", categoryHandler).Methods("GET")
	mux.HandleFunc("/categories", saveCategoryHandler).Methods("POST")
	mux.HandleFunc("/categories", categoriesHandler).Methods("GET")
	mux.HandleFunc("/categories/{id:[0-9]+}", updateCategoryHandler).Methods("PUT")
	mux.HandleFunc("/categories/{id:[0-9]+}", deleteCategoryHandler).Methods("DELETE")
	return mux
}
