package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/omekov/online-market/backend/models"
)

func originHandler(w http.ResponseWriter, r *http.Request) {
	product := models.FoodProduct{}
	err := models.GetOrigins(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	out, err := json.Marshal(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(out)
}
