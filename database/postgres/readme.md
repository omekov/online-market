<!-- Starting bash for container -->
docker exec -it shop_postgres_container bash

<!-- reset password postgres -->
passwd postgres

<!-- connnect to user postgres -->
su - postgres

<!-- create database for UNICODE -->
createdb --encoding UNICODE marketdb --username postgres

<!-- sign in command control for postgres  -->
psql

<!-- create new user -->
create user azamat with password 'azamat';

<!-- update role -->
ALTER USER azamat CREATEDB;

<!-- privileges of marketdb for new user -->
grant all privileges on database marketdb to azamat;


<!-- MIGRATIONS  table -->
<!-- RUN Command is local machine -->
migrate -path database/migrations -database "postgres://localhost:5431/marketdb?sslmode=disable&user=azamat&password=azamat" up
<!-- DROP TABLE -->
migrate -path database/migrations -database "postgres://localhost:5431/marketdb?sslmode=disable&user=azamat&password=azamat" down














wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc|sudo apt-key add-;\
RELEASE=$(isb_release -cs);\
echo "deb http://apt.postgresql.org/pub/repos/apt/${RELEASE}"-pgdg main|sudo tee /etc/apt/sources.list.d/pgdg.list;\
sudo apt update;\
sudo apt -y install postgresql-12

sudo localedef ru_RU.UTF-8 -i ru_RU -fUTF-8;\
export LANGUAGE=ru_RU.UTF-8;\
export LANG=ru_RU.UTF-8;\
export LC_ALL=ru_RU.UTF-8;\
sudo locale-gen ru_RU.UTF-8;\
sudo dpkg-reconfigure locales


docker exec -it 05b3a3471f6f bash

<!-- reset password -->
sudo passwd postgres
<!-- path -->
su - postgres
export PATH=$PATH:/usr/lib/postgres/12/bin
<!-- createdb -->
createdb --encoding UNICODE salesbeat_db --username postgres;
<!-- sing postgres -->
sudo -u postgres psql
<!-- create user only salesbeat -->
create user salesbeat with password 'AQG9bVPG';
<!-- db in user -->
ALTER USER salesbeat CREATEDB;
<!-- premission in group -->
grant all privileges on database salesbeat_db to sales beat;

sudo -u postgres psql
postgres=# ...
create user dbms with password 'some_password';
ALTER USER dbms CREATEDB;
grant all privileges on database dbms_db to dbms;
\c dbms_db
GRANT ALL ON ALL TABLES IN SCHEMA public to dbms;
GRANT ALL ON ALL SEQUENCES IN SCHEMA public to dbms;
GRANT ALL ON ALL FUNCTIONS IN SCHEMA public to dbms;
CREATE EXTENSION pg_trgm;
ALTER EXTENSION pg_trgm SET SCHEMA public;
UPDATE pg_opclass SET opcdefault = true WHERE opcname='gin_trgm_ops';
\q
exit

vim ~/.pgpass
	localhost:5432:dbms_db:dbms:some_password
chmod 600 ~/.pgpass
psql -h localhost -U dbms dbms_db

psql -h localhost dbms_db dbms  < dump.sql