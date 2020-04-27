
Шпаргалки
create database titan
create user author with login password 'qwerty';
\password author

drop database ddd
drop user author

Создание таблицы
migrate create-ext sql -dir migrations <table>

миграция таблиц
migrate -path database/migrations -database "postgres://localhost:5432/marketdb?sslmode=disable&user=postgres&password=123" down

чтобы ошибку устранить
migrate -path database/migrations -database "postgres://localhost:5432/marketdb?sslmode=disable&user=postgres&password=Welcome01" force 20200417090142