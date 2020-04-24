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
		CreateAt:    &now,
		UpdateAt:    &now,
	}
}

// TestCategories ...
func TestCategories(t *testing.T) *[]Category {
	now := time.Now()
	return &[]Category{
		Category{
			Name:        "Овощи",
			CreateAt:    &now,
			UpdateAt:    &now,
			Description: "Описание овощи",
		},
	}
}
