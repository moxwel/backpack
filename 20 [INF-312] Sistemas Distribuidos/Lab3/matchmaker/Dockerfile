FROM golang:1.24.1

# Dependencias
WORKDIR /app/types
COPY /types ./

WORKDIR /app/matchmaker
COPY /matchmaker/go.mod ./
COPY /matchmaker/go.sum ./
RUN go mod download

# Codigo fuente
COPY /matchmaker ./
RUN go build -o matchmaker main.go servidorMatchmaking.go clienteServer.go

CMD ["./matchmaker"]