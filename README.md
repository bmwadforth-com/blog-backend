# Template for web services
This template allows you to quickly scaffold a go repo to create a generic RESTFUL API.

It comes with a variety of boilerplate code to connect to and consume commonly used services in a web service.

## Pre-requisites
1. The application authenticates with Google cloud platform via the `GOOGLE_APPLICATION_CREDENTIALS` environment variable. When deployed, the cloud run instance automatically has access to the credentials. When developing locally you will need to set the variable. See https://cloud.google.com/docs/authentication/application-default-credentials
2. When changes are pushed to the main branch, [Cloud build](https://cloud.google.com/build?hl=en) will automatically build the code and deploy it into [Cloud run](https://cloud.google.com/run?hl=en).

## Documentation
1. Gin - https://github.com/gin-gonic/gin

## Guide
Simply run `go run main.go` from the project directory, and it will start a web service on port `8080` with a couple of endpoints.

When the application is deployed in production, it reads its configuration from environment variables. All environment variables that start with `WEB_TEMPLATE__` are loaded. When the application is in development mode, it will read the configuration from `config.json`.

## Swagger
To generate a swagger specification, ensure under `controllers` your documentation is correct and then run the following commands from the root directory:

```
go install github.com/swaggo/swag/cmd/swag@latest
swag init
```

You should then have your swagger files accessible under `docs`. Additionally, in non-production (when APP_ENV=DEVELOPMENT) when starting the application - you can access the swagger documentation at `localhost:8080/swagger/index.html`.