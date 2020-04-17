package apiserver

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/omekov/online-market/back-api/internal/app/model"
	"github.com/omekov/online-market/back-api/internal/app/store"
	"github.com/sirupsen/logrus"
)

type ctxKey int

const (
	ctxKeyRequestID ctxKey = iota
	route                  = "/api"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

// newServer - метод старта api handles
func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
	}
	s.Handlers()
	return s
}

// Handlers - метод для обработки api запросов
func (s *server) Handlers() {
	s.router.Use(s.setRequestID)
	s.router.Use(s.logRequest)
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	s.router.Use(s.setHeader)
	s.router.HandleFunc("/", s.handlerHome()).Methods(http.MethodGet)
	s.router.HandleFunc(route+"/categories", s.handlerCategories()).Methods(http.MethodGet)
	s.router.HandleFunc(route+"/categories/{id:[0-9]+}", s.handlerCategoryGetByID()).Methods(http.MethodGet)
	s.router.HandleFunc(route+"/categories", s.handlerCategoryCreate()).Methods(http.MethodPost)
	s.router.HandleFunc(route+"/categories/{id:[0-9]+}", s.handlerCategoryUpdate()).Methods(http.MethodPut)
	s.router.HandleFunc(route+"/categories/{id:[0-9]+}", s.handlerCategoryDelete()).Methods(http.MethodDelete)
}

// ServeHTTP - Нужен для роута
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// setHeader - добавления стандартных заголовок
func (s *server) setHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// setRequestID - middleware для уникальности каждого запроса X-Request-ID
func (s *server) setRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(
			context.WithValue(
				r.Context(),
				ctxKeyRequestID,
				id,
			)))
	})
}

// logRequest - middleware для логирование запросов
func (s *server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
			"request_id":  r.Context().Value(ctxKeyRequestID),
		})
		logger.Infof("started %s %s", r.Method, r.RequestURI)
		start := time.Now()
		// для обработки response до и после
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)
		logger.Infof(
			"completed with %d %s in %v",
			rw.code,
			http.StatusText(rw.code),
			time.Now().Sub(start),
		)
	})
}

// handlerHome - начальный handler для проверки сервера
func (s *server) handlerHome() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server back-api starting"))
	})
}

// handlerCategories - обработчик категорий продуктов
func (s *server) handlerCategories() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		categories, err := s.store.Category().GetAll()
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		s.respond(w, r, http.StatusOK, categories)
	})
}

// handlerCategoryGetByID - получить одну категорию
func (s *server) handlerCategoryGetByID() http.HandlerFunc {
	category := model.Category{}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		err = s.store.Category().GetByID(id, &category)
		if err != nil {
			if err == sql.ErrNoRows {
				s.error(w, r, http.StatusNotFound, store.ErrRecordNotFound)
				return
			}
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		s.respond(w, r, http.StatusOK, category)
	})
}

// handlerCategoryCreate - создания категорий
func (s *server) handlerCategoryCreate() http.HandlerFunc {
	category := model.Category{}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := json.NewDecoder(r.Body).Decode(&category)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		err = s.store.Category().Create(&category)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		s.respond(w, r, http.StatusCreated, category)
	})
}

// handlerCategoryCreate - обновить категорию
func (s *server) handlerCategoryUpdate() http.HandlerFunc {
	category := model.Category{}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		err = json.NewDecoder(r.Body).Decode(&category)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
		}
		err = s.store.Category().Update(id, &category)
		if err != nil {
			if err == sql.ErrNoRows {
				s.error(w, r, http.StatusNotFound, store.ErrRecordNotFound)
				return
			}
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		s.respond(w, r, http.StatusOK, category)
	})
}

// handlerCategoryCreate - удалить категорию
func (s *server) handlerCategoryDelete() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		err = s.store.Category().Delete(id)
		if err != nil {
			if err == sql.ErrNoRows {
				s.error(w, r, http.StatusNotFound, store.ErrRecordNotFound)
				return
			}
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		s.respond(w, r, http.StatusOK, nil)
	})
}

// respond - Обработка успешного ответа
func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

// respond - Обработка ошибочного ответа
func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}
