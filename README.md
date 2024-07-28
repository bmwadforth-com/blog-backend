# Blog Website
This repository is for the application deployed at `bmwadforth.com`.

# Prerequisites
1. This application uses services from Google Cloud (Storage, Firestore, etc.). To authenticate follow this guide [here](https://cloud.google.com/docs/authentication/application-default-credentials).
2. This application also uses swagger which is available at `localhost:8080/swagger/index.html` when you start the app. Install swagger by running the following command `go install github.com/swaggo/swag/cmd/swag@latest`. Make sure `swag` is in your `PATH` otherwise the makefile commands will not work.
3. You will need to install `npm` as the frontend for this application is under [web](web).
4. You will also need Go installed. Once the above steps are complete simply run `go run main.go` from the command line and the backend will start.
5. Copy [config.json](config.json) into a new file called `config.local.json`. This will be your configuration when the app is running locally.
6. To start the frontend, navigate to the web directory and run `npm install` and `npm run start`.
7. If you want to use the [makefile](Makefile) to **generate the latest swagger specification**, **build the backend**, **build the frontend** and **start the application** you can simply run `make build_all`.
8. Otherwise, simply run `go start main.go` in one CLI shell to start the backend, and in another CLI shell navigate to the [web](web) directory and run `npm run start`.

If you run `make build_all_start` make will build the frontend which will output its artifacts [here](./web/build). The makefile will then embed those build artifacts into the go binary using [embed](https://pkg.go.dev/embed). This is how the application is deployed - the frontend react application is embedded into the Go binary on the web server in GCP and the go application statically serves these files. 

# Deployment
When you merge into the `main` branch - [Cloud build](https://cloud.google.com/build?hl=en) will pull the changes and build a docker image based on the dockerfile defined [here](./Dockerfile). Cloud build will then deploy the service into [Cloud run](https://cloud.google.com/run?hl=en).

# Documentation
1. Gin - https://github.com/gin-gonic/gin
2. Firestore - https://cloud.google.com/firestore?hl=en
3. Cloud storage - https://cloud.google.com/storage?hl=en
4. Cloud run - https://cloud.google.com/run?hl=en
5. Cloud build - https://cloud.google.com/build?hl=en
6. API Gateway - https://cloud.google.com/api-gateway