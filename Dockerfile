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
COPY data/swif_codes.csv /app/data/intern.csv
COPY ./makefile /app/makefile

ENV PORT=8081
ENV DATA_PATH=/app/data/intern.csv
EXPOSE 8081
CMD ["./main"]
