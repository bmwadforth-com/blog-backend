FROM node:18 as node-env

WORKDIR /app
COPY ./web .

RUN npm run artifactregistry-login
RUN npm install
RUN npm run build



FROM golang:1.21.3 as go-env
ENV PORT=8080
ENV APP_ENV=PRODUCTION

WORKDIR /app
COPY --from=node-env /app/build ./web/build
COPY go.mod go.sum ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /blog-backend

COPY ./blog-backend ./

EXPOSE 8080
CMD ["/blog-backend"]