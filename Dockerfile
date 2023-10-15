FROM node:18 as node-env

WORKDIR /app

COPY ./web .

RUN npm install
RUN npm run build


FROM golang:1.21.3 as go-env
ENV PORT=8080
ENV APP_ENV=PRODUCTION

WORKDIR /app

COPY --from=node-env /app/web/build ./web/build

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
#COPY *.go ./

COPY ./ ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /blog-backend

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 8080

# Run
CMD ["/blog-backend"]