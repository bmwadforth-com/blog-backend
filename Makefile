swagger:
	swag init

backend_build:
	go build -o bin/main main.go

backend_start:
	go run main.go

frontend_build:
	cd ./web && npm run artifactregistry-login && npm install && npm run build && cd ..

build_all: swagger backend_build frontend_build backend_start