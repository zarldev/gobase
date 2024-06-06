FROM golang:1.22.3 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN --mount=type=bind,source=.env,target=/app/.env \
    export $(cat .env | xargs) && \
    make build

FROM debian:latest
WORKDIR /app
COPY --from=builder /app/go-base .

ENTRYPOINT ["./go-base"]