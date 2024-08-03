swagger:
	swag init

backend_build:
	go build -o bin/main main.go

backend_start:
	go run main.go

frontend_build:
	npm --prefix ./web install && npm --prefix ./web run build

frontend_start:
	npm --prefix ./web start

build_all: frontend_build swagger backend_build

build_all_start: frontend_build swagger backend_build backend_start