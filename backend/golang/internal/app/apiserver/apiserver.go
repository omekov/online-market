package apiserver

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gomodule/redigo/redis"
	"github.com/joho/godotenv"
	"github.com/omekov/online-market/backend/golang/internal/app/store/sqlstore"
	"github.com/sirupsen/logrus"
)

// Start - соединение с базой подключение Роута
func Start() error {
	db, err := newDB()
	if err != nil {
		return err
	}
	defer db.Close()
	c, err := newRedis()
	if err != nil {
		return err
	}
	defer c.Close()
	store := sqlstore.New(db)
	r := newServer(store, c)
	return http.ListenAndServe(":5053", r)
}

// newDB - обработка соединение с базой
func newDB() (*sql.DB, error) {
	config, err := dbConfig()
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("postgres", config)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

// dbConfig - изъятие переменных окружений
func dbConfig() (string, error) {
	if _, err := os.Stat(".env"); !os.IsNotExist(err) {
		err = godotenv.Load()
		if err != nil {
			return "", err
		}
	}
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_NAME"),
	)
	return connStr, nil
}

func newRedis() (redis.Conn, error) {
	c, err := redis.DialURL("redis://shop_redis_container")
	if err != nil {
		return nil, err
	}
	logrus.Info("success connection to Redis")
	return c, nil
}
