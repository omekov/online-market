package model

import (
	"testing"
	"time"
)

// TestCategory - тестовая категория для теста
func TestCategory(t *testing.T) *Category {
	now := time.Now()
	return &Category{
		Name:     "Овощи",
		RusName:  "Овощи",
		Color:    "green",
		CreateAt: now,
		UpdateAt: &now,
		OriginID: 1,
	}
}

// TestCategories ...
func TestCategories(t *testing.T) *[]Category {
	now := time.Now()
	return &[]Category{
		Category{
			Name:     "Овощи",
			RusName:  "Овощи",
			Color:    "green",
			CreateAt: now,
			UpdateAt: &now,
			OriginID: 1,
		},
	}
}
