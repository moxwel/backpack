FROM golang:1.23.9

# Dependencias
WORKDIR /app/types
COPY /types ./

WORKDIR /app/entrenadores
COPY /entrenadores/go.mod ./
COPY /entrenadores/go.sum ./
RUN go mod download

# Codigo fuente
COPY /entrenadores ./
RUN go build -o entrenadores main.go consumerRabbitSNP.go clienteLCP.go

CMD ["./entrenadores"]