package apiserver

import (
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
	"github.com/omekov/online-market/backend/golang/internal/app/store"
	"github.com/sirupsen/logrus"
)

type Server struct {
	router  *mux.Router
	logger  *logrus.Logger
	store   store.Store
	session redis.Conn
}

// newServer - метод старта api handles
func newServer(store store.Store, cache redis.Conn) *Server {
	s := &Server{
		router:  mux.NewRouter(),
		logger:  logrus.New(),
		store:   store,
		session: cache,
	}
	s.Handlers()
	return s
}
