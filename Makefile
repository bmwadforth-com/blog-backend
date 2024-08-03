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

docker_build:
	docker build --build-arg NODE_AUTH_TOKEN=$(NODE_AUTH_TOKEN) . --file Dockerfile \
	--tag ghcr.io/$(REPO)/$(REPO_NAME):$(COMMIT_SHA) \
	--tag $(GAR_LOCATION)-docker.pkg.dev/$(PROJECT_ID)/$(REPO)/$(REPO_NAME):$(COMMIT_SHA) \
	--tag ghcr.io/$(REPO)/$(REPO_NAME):latest \
	--tag $(GAR_LOCATION)-docker.pkg.dev/$(PROJECT_ID)/$(REPO)/$(REPO_NAME):latest

build_all: frontend_build swagger backend_build

build_all_start: frontend_build swagger backend_build backend_start