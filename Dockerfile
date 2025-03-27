#Building State
FROM golang:1.24-alpine AS builder
WORKDIR /app
RUN apk add --no-cache git
COPY go.mod go.sum  ./
COPY . .
RUN go mod download
RUN go build -o main main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY data/swift_codes.csv /app/data/intern.csv
COPY ./makefile /app/makefile

EXPOSE 8080
CMD ["./main"]
