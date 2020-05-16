package apiserver

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/omekov/online-market/backend/golang/internal/app/model"
	"github.com/omekov/online-market/backend/golang/internal/app/store"
)

const route string = "/api"

// Router - метод для обработки api запросов
func (s *Server) Handlers() {
	s.router.Use(s.setRequestID)
	s.router.Use(s.logRequest)
	s.router.Use(s.setHeader)
	s.router.Use(s.setCookie)
	s.router.HandleFunc("/", s.handlerHome()).Methods(http.MethodGet)
	s.router.HandleFunc(route+"/products", s.handlerProducts()).Methods(http.MethodGet)
	s.router.HandleFunc(route+"/products/{id:[0-9]+}", s.handlerProductGetByID()).Methods(http.MethodGet)
	s.router.HandleFunc(route+"/products", s.handlerProductCreate()).Methods(http.MethodPost)
	s.router.HandleFunc(route+"/categories", s.handlerCategories()).Methods(http.MethodGet)
	s.router.HandleFunc(route+"/categories", s.handlerCategoryCreate()).Methods(http.MethodPost)
	s.router.HandleFunc(route+"/stocks", s.handlerStockCreate()).Methods(http.MethodPost)
	s.router.HandleFunc(route+"/products/{id:[0-9]+}", s.handlerProductUpdate()).Methods(http.MethodPut)
	s.router.HandleFunc(route+"/products/{id:[0-9]+}", s.handlerProductDelete()).Methods(http.MethodDelete)
	s.router.HandleFunc(route+"/uploadproductphoto", s.handlerUploadPhotoProduct())
}

// ServeHTTP - Нужен для роута
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// handlerHome - начальный handler для проверки сервера
func (s *Server) handlerHome() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server golang starting"))
	})
}

// handlerProducts - обработчик категорий продуктов
func (s *Server) handlerProducts() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		products, err := s.store.Product().GetAll()
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		s.respond(w, r, http.StatusOK, products)
	})
}

// handlerproductGetByID - получить одну категорию
func (s *Server) handlerProductGetByID() http.HandlerFunc {
	product := model.Product{Category: &model.Category{}, Stock: &model.Stock{}}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		err = s.store.Product().GetByID(id, &product)
		if err != nil {
			if err == sql.ErrNoRows {
				s.error(w, r, http.StatusNotFound, store.ErrRecordNotFound)
				return
			}
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		s.respond(w, r, http.StatusOK, product)
	})
}

// handlerProductCreate - создания категорий
func (s *Server) handlerProductCreate() http.HandlerFunc {
	product := model.Product{}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		if err := s.store.Product().Create(&product); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		s.respond(w, r, http.StatusCreated, product)
	})
}

func (s *Server) handlerCategoryCreate() http.HandlerFunc {
	category := model.Category{}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := json.NewDecoder(r.Body).Decode(&category)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		err = s.store.Product().CreateCategory(&category)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		s.respond(w, r, http.StatusCreated, category)
	})
}

func (s *Server) handlerCategories() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		categories, err := s.store.Product().GetAllCategory()
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		s.respond(w, r, http.StatusOK, categories)
	})
}

func (s *Server) handlerStockCreate() http.HandlerFunc {
	stock := model.Stock{}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := json.NewDecoder(r.Body).Decode(&stock)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		err = s.store.Product().CreateStock(&stock)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		s.respond(w, r, http.StatusCreated, stock)
	})
}

// handlerProductUpdate - обновить категорию
func (s *Server) handlerProductUpdate() http.HandlerFunc {
	product := model.Product{}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		err = json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
		}
		err = s.store.Product().Update(id, &product)
		if err != nil {
			if err == sql.ErrNoRows {
				s.error(w, r, http.StatusNotFound, store.ErrRecordNotFound)
				return
			}
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		s.respond(w, r, http.StatusOK, product)
	})
}

// handlerProductDelete - удалить категорию
func (s *Server) handlerProductDelete() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		err = s.store.Product().Delete(id)
		if err == sql.ErrNoRows {
			s.error(w, r, http.StatusNotFound, store.ErrRecordNotFound)
			return
		}
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		s.respond(w, r, http.StatusOK, nil)
	})
}

func (s *Server) handlerUploadPhotoProduct() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(10 << 20)
		file, handler, err := r.FormFile("myfile")
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		defer file.Close()
		fmt.Printf("Uploaded file: %+v\n", handler.Filename)
		fmt.Printf("File size: %+v\n", handler.Size)
		fmt.Printf("MIME Header: %+v\n", handler.Header)
		tempFile, err := ioutil.TempFile("temp-images", "product-photo-*.png")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer tempFile.Close()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		tempFile.Write(fileBytes)
		s.respond(w, r, http.StatusOK, map[string]string{"status": "Successfully Uploaded File"})
	})
}
