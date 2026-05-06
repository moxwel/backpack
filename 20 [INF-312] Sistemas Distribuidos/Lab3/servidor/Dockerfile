FROM golang:1.24.1

# Dependencias
WORKDIR /app/types
COPY /types ./

WORKDIR /app/servidor
COPY /servidor/go.mod ./
COPY /servidor/go.sum ./
RUN go mod download

# Codigo fuente
COPY /servidor ./
RUN go build -o servidor main.go servidorGameServer.go clienteMatchmaking.go

CMD ["./servidor"]