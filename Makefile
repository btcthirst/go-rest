include .env
exe_path := bin/app
work_path := cmd/main.go

.PHONY: go-build
go-build:
	go build -o ./${exe_path} ./${work_path}

.PHONY: go-run
go-run: go-build
	./${exe_path}

.PHONY: test
test:
	go test -v ./...

.PHONY: dep
dep:
	go mod download

.PHONY: docker-build
docker-build:
	docker build -t go-app .

.PHONY: up
up:
	docker-compose up --force-recreate --build

.PHONY: down
down:
	docker-compose down --rmi all

.PHONY: db-migrate
db-migrate:
	goose -dir internal/database/migrations up

.PHONY: status
status:
	goose status

.DEFAULT_GOAL=docker-up