FROM golang:1.22.5-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

EXPOSE 8080

CMD ["sh", "-c", "go run ./api/cmd/main.go"]
