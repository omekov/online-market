CREATE TABLE categories (
	id serial primary key,
	name varchar(25) not null,
	russianName varchar(50) not null,
	color varchar(7) not null,
	createAt timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updateAt timestamp,
	originId serial not null
)

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