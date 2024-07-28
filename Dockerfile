FROM golang:1.22.5-alpine

WORKDIR /app

# Copia os arquivos de configuração e dependências
COPY go.mod go.sum ./
RUN go mod tidy

# Copia todos os arquivos do projeto para o contêiner
COPY . .

# Compila o aplicativo
RUN go build -o main ./api/cmd/main.go

EXPOSE 8080

CMD ["./main"]
