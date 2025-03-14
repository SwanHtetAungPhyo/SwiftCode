#Building State
FROM golang:1.24-alpine AS builder
WORKDIR /app
RUN apk add --no-cache git
COPY go.mod go.sum  ./
COPY . .
RUN go mod download
RUN go build -o main ./cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY ./data/intern.csv /app/data/intern.csv
COPY ./makefile /app/makefile

ENV PORT=8080
ENV DATA_PATH=/app/data/intern.csv
EXPOSE 8080
CMD ["./main"]
