FROM golang:1.22.1 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o blackenshovel-service ./cmd/server/main.go

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=build /app/blackenshovel-service .

EXPOSE 8080

CMD ["./blackenshovel-service"]
