exe_path := bin/app
work_path := cmd/main.go
build:
	go build -o ./${exe_path} ./${work_path}

run: build
	./${exe_path}

test:
	go test -v main.go
dep:
	go mod download

.DEFAULT_GOAL=run