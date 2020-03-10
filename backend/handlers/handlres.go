package handlers

import (
	"github.com/gorilla/mux"
)

// Router ...
func Router() *mux.Router {
	mux := mux.NewRouter()
	mux.HandleFunc("/products", ProductsHandler)
	mux.HandleFunc("/origins", originHandler)
	return mux
}
