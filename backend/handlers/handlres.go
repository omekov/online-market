package handlers

import "net/http"

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/products", ProductsHandler)
	mux.HandleFunc("/origins", originHandler)
	return mux
}
