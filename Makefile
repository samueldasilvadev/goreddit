.PHONY: postgres adminer migrate migrate-down

postgres:
	docker run --rm -it \
	--network host \
	-p 5432:5432 \
	-e POSTGRES_PASSWORD=secret \
	postgres

adminer:
		docker run --rm -it \
	--network host \
	adminer

# Require https://github.com/golang-migrate/migrate install
migrate:
	migrate -source file://migrations \
	-database postgres://postgres:secret@localhost/postgres?sslmode=disable up

# Require https://github.com/golang-migrate/migrate install
migrate-down:
	migrate -source file://migrations \
	-database postgres://postgres:secret@localhost/postgres?sslmode=disable down
