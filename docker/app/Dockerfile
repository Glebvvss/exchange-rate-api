FROM golang:1.21

WORKDIR /app

COPY ../.. .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /app.bin ./cmd/app
RUN CGO_ENABLED=0 GOOS=linux go build -o /migrate.bin ./cmd/migrate

EXPOSE 8080

CMD ["/app.bin"]