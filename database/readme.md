CREATE TABLE categories (
	id serial primary key,
	name varchar(25) not null,
	russianName varchar(50) not null,
	color varchar(7) not null,
	createAt timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updateAt timestamp,
	originId serial not null
)

CREATE TABLE categories (
	id serial primary key,
	name varchar(25) not null,
	russianName varchar(50) not null,
	color varchar(7) not null,
	createAt timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updateAt timestamp,
	originId serial not null
)