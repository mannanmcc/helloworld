1. Install golang-migrate command using brew
```
brew install golang-migrate
```

2. Create a migration files using following command. It will add a migration file under migration folder, then fill up the sql in both up and down files 
```
migrate create -ext sql -dir db/migrations -seq create_items_table
```
3. The export database connection with command and exexute up/down sql
```
export POSTGRESQL_URL="postgres://[user]:[password]@[host]:[post]/[database_name]?sslmode=disable"
(export POSTGRESQL_URL="postgres://test:password@localhost:5432/fullstack_api?sslmode=disable")

migrate -database ${POSTGRESQL_URL} -path db/migrations up