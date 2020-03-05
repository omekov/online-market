package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/omekov/online-market/backend/db"
	m "github.com/omekov/online-market/backend/models"
)

// productsHandler ...
func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	name := ""
	foodProduct := m.FoodProduct{Origins: []m.Origins{
		m.Origins{
			ID:          1,
			Name:        "Plant",
			RussianName: "Продукты растительного происхождения",
			CreateAt:    time.Now(),
			UpdateAt:    time.Now().AddDate(1, 02, 01),
			Categories: &[]m.Category{
				m.Category{
					ID:          1,
					Name:        &name,
					RussianName: "Овощи",
					Color:       "blue",
					CreateAt:    time.Now(),
					UpdateAt:    time.Now().AddDate(0, 02, 02),
				},
				m.Category{
					ID:          2,
					Name:        nil,
					RussianName: "Фрукты",
					Color:       "orange",
					CreateAt:    time.Now(),
					UpdateAt:    time.Now().AddDate(1, 02, 02),
				},
			},
		},
		m.Origins{
			ID:          1,
			Name:        "Plant",
			RussianName: "Продукты растительного происхождения",
			CreateAt:    time.Now(),
			UpdateAt:    time.Now().AddDate(1, 02, 01),
			Categories:  nil,
		},
	}}
	products, err := json.Marshal(foodProduct)
	if err != nil {
		fmt.Fprintf(w, "JSON marshall", err)
		return
	}
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "applications/json")
	w.Write(products)
}
func originHandler(w http.ResponseWriter, r *http.Request) {
	product := m.FoodProduct{}
	err := db.SelectOrigin(&product)
	if err != nil {
		http.Error(w, err.Error(), 500)
		fmt.Fprintf(w, "Db SelectOrigin", err)
		return
	}
	out, err := json.Marshal(product)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "applications/json")
	w.Write(out)
}
