SHELL = /bin/sh
run:
	go run main.go

build:
	go build


tools/golang-migrate:
	$(call print-target)
	GOBIN=$(CURDIR)/tools go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.14.1

#.PHONY: migrateup
#migrateup: $(call print-target)
#migrate -path migrations -database "postgres://sahan:142963@localhost:5432/sahan1?sslmode=disable" -verbose up


#.PHONY: migratedown
#migratedown: $(call print-target)
#migrate -path migrations -database "postgres://sahan:142963@localhost:5432/sahan1?sslmode=disable" -verbose down

PHONY: migrate_up
migrate_up: tools/migrate
	$(call print-target)
	./tools/migrate -path migrations/todo_api  -database ${POSTGRESQL_URL}  -verbose up 

PHONY: migrate_down
migrate_down: tools/migrate
	./tools/migrate -path migrations/todo_api -database ${POSTGRESQL_URL} -verbose down 