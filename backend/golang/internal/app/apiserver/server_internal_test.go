package apiserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/omekov/online-market/backend/golang/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestServer_handlerProductID(t *testing.T) {
	s := newServer(teststore.New(), nil)
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name:         "valid",
			payload:      "1",
			expectedCode: http.StatusOK,
		},
		{
			name:         "invalid not found",
			payload:      "2992929292",
			expectedCode: http.StatusNotFound,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/products/%s", route, tc.payload), nil)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_handlerProductCreate(t *testing.T) {
	s := newServer(teststore.New(), nil)
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]interface{}{
				"name":      "test",
				"rus_name":  "тест",
				"color":     "#ffffff",
				"create_at": time.Now(),
				"update_at": time.Now(),
				"origin_id": 1,
			},
			expectedCode: http.StatusCreated,
		},
		{
			name:         "invalid payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			buf := &bytes.Buffer{}
			json.NewEncoder(buf).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, route+"/products", buf)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_handlerProductsUpdate(t *testing.T) {
	s := newServer(teststore.New(), nil)
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]interface{}{
				"name":        "test",
				"description": "тест",
				"price":       9.99,
				"created_at":  time.Now(),
				"updated_at":  time.Now(),
				"origin_id":   1,
			},
			expectedCode: http.StatusOK,
		},
		{
			name:         "invalid payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			buf := &bytes.Buffer{}
			json.NewEncoder(buf).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPut, route+"/products/1", buf)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_handlerProductsDelete(t *testing.T) {
	s := newServer(teststore.New(), nil)
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name:         "valid",
			payload:      nil,
			expectedCode: http.StatusOK,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			buf := &bytes.Buffer{}
			json.NewEncoder(buf).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodDelete, route+"/products/666", buf)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_handlerCategoryCreate(t *testing.T) {
	s := newServer(teststore.New(), nil)
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]interface{}{
				"name":        "test",
				"description": "тест",
				"created_at":  time.Now(),
				"updated_at":  time.Now(),
			},
			expectedCode: http.StatusCreated,
		},
		{
			name:         "invalid payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			buf := &bytes.Buffer{}
			json.NewEncoder(buf).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, route+"/categories", buf)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_handlerStockCreate(t *testing.T) {
	s := newServer(teststore.New(), nil)
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]interface{}{
				"name":        "test",
				"description": "тест",
				"precent":     9.99,
				"created_at":  time.Now(),
				"updated_at":  time.Now(),
			},
			expectedCode: http.StatusCreated,
		},
		{
			name:         "invalid payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			buf := &bytes.Buffer{}
			json.NewEncoder(buf).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, route+"/stocks", buf)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
