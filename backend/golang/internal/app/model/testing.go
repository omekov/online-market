package model

import (
	"testing"
	"time"
)

// TestCategory - тестовая категория для теста
func TestCategory(t *testing.T) *Category {
	now := time.Now()
	return &Category{
		Name:        "Овощи",
		Description: "Описание овощи",
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}
}

// TestCategories ...
func TestCategories(t *testing.T) *[]Category {
	now := time.Now()
	return &[]Category{
		Category{
			Name:        "Овощи",
			CreatedAt:   &now,
			UpdatedAt:   &now,
			Description: "Описание овощи",
		},
	}
}

func TestProduct(t *testing.T) *Product {
	now := time.Now()
	return &Product{
		Name:        "Овощи",
		Description: "Описание овощи",
		CreatedAt:   &now,
		UpdatedAt:   &now,
		Price:       999.99,
		Category: &Category{
			Name:        "Овощной",
			Description: "Овощной отдел",
			CreatedAt:   &now,
			UpdatedAt:   &now,
		},
		Stock: &Stock{
			Name:        "Скидка года",
			Description: "Бери не глядя",
			Precent:     0.01,
			CreatedAt:   &now,
			UpdatedAt:   &now,
		},
	}
}
