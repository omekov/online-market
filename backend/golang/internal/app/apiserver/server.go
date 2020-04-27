package apiserver

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/omekov/online-market/backend/golang/internal/app/model"
	"github.com/omekov/online-market/backend/golang/internal/app/store"
	"github.com/sirupsen/logrus"
)

type ctxKey int

const (
	ctxKeyRequestID ctxKey = iota
	route                  = "/api"
)

type server struct {
	router  *mux.Router
	logger  *logrus.Logger
	store   store.Store
	session redis.Conn
}

// newServer - метод старта api handles
func newServer(store store.Store, cache redis.Conn) *server {
	s := &server{
		router:  mux.NewRouter(),
		logger:  logrus.New(),
		store:   store,
		session: cache,
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
	s.router.HandleFunc(route+"/products", s.handlerProducts()).Methods(http.MethodGet)
	s.router.HandleFunc(route+"/products/{id:[0-9]+}", s.handlerProductGetByID()).Methods(http.MethodGet)
	s.router.HandleFunc(route+"/products", s.handlerProductCreate()).Methods(http.MethodPost)
	s.router.HandleFunc(route+"/categories", s.handlerCategoryCreate()).Methods(http.MethodPost)
	s.router.HandleFunc(route+"/stocks", s.handlerStockCreate()).Methods(http.MethodPost)
	s.router.HandleFunc(route+"/products/{id:[0-9]+}", s.handlerProductUpdate()).Methods(http.MethodPut)
	s.router.HandleFunc(route+"/products/{id:[0-9]+}", s.handlerProductDelete()).Methods(http.MethodDelete)
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
		cookie, err := r.Cookie("session_token")
		if err != nil {
			sessionToken := uuid.New().String()
			_, err := s.session.Do("SETEX", sessionToken, "120", "unknown")
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			cookie = &http.Cookie{
				Name:    "session_token",
				Value:   sessionToken,
				Expires: time.Now().Add(120 * time.Second),
			}
			http.SetCookie(w, cookie)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server golang starting"))
	})
}

// handlerProducts - обработчик категорий продуктов
func (s *server) handlerProducts() http.HandlerFunc {
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
func (s *server) handlerProductGetByID() http.HandlerFunc {
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
func (s *server) handlerProductCreate() http.HandlerFunc {
	product := model.Product{}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		err = s.store.Product().Create(&product)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		s.respond(w, r, http.StatusCreated, product)
	})
}

func (s *server) handlerCategoryCreate() http.HandlerFunc {
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

func (s *server) handlerStockCreate() http.HandlerFunc {
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
func (s *server) handlerProductUpdate() http.HandlerFunc {
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
func (s *server) handlerProductDelete() http.HandlerFunc {
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
