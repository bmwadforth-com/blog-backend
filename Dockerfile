# Build frontend
FROM node:18 as frontend-build

ARG NODE_AUTH_TOKEN
ENV NODE_AUTH_TOKEN=$NODE_AUTH_TOKEN
WORKDIR /app
COPY ./web .
RUN npm install
RUN npm run build

# Build backend
FROM golang:1.22.1 as backend-build

ENV PORT=8080
ENV APP_ENV=PRODUCTION
WORKDIR /app
COPY --from=frontend-build /app/build ./web/build
COPY ./ ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o ./blog-backend

# Build production
FROM golang:1.22.1 as production

ENV PORT=8080
ENV APP_ENV=PRODUCTION
WORKDIR /app
COPY --from=backend-build /app/blog-backend ./
EXPOSE 8080
CMD ["/app/blog-backend"]