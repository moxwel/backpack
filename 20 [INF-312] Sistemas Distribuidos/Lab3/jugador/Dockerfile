FROM golang:1.24.1

# Dependencias
WORKDIR /app/types
COPY /types ./

WORKDIR /app/jugador
COPY /jugador/go.mod ./
COPY /jugador/go.sum ./
RUN go mod download

# Codigo fuente
COPY /jugador ./
RUN go build -o jugador main.go clienteMatchmaking.go clienteGameServer.go

CMD ["./jugador"]