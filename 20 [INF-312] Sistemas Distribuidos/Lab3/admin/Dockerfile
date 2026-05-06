FROM golang:1.24.1

# Dependencias
WORKDIR /app/types
COPY /types ./

WORKDIR /app/admin
COPY /admin/go.mod ./
COPY /admin/go.sum ./
RUN go mod download

# Codigo fuente
COPY /admin ./
RUN go build -o admin main.go clienteMatchmaking.go

CMD ["./admin"]