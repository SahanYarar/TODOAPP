SHELL = /bin/sh
run:
	go run main.go

build:
	go build

.PHONY: install_tools
install_tools:  tools/golang-migrate

tools/golang-migrate:
	$(call print-target)
	GOBIN=$(CURDIR)/tools go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.14.1

PHONY: migrate_up
migrate_up: tools/migrate
	$(call print-target)
	echo ${POSTGRESQL_URL} 
	./tools/migrate -path migrations  -database "postgres://sahan:142963@localhost:5432/sahan?sslmode=disable"  -verbose up 

PHONY: migrate_down
migrate_down: tools/migrate
	./tools/migrate -path migrations -database "postgres://sahan:142963@localhost:5432/sahan?sslmode=disable" -verbose down 

lint:
 golangci-lint run

