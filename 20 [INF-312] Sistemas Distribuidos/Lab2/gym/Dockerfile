FROM golang:1.23.9

# Dependencias
WORKDIR /app/types
COPY /types ./

WORKDIR /app/gym
COPY /gym/go.mod ./
COPY /gym/go.sum ./
RUN go mod download

# Codigo fuente
COPY /gym ./
RUN go build -o gym main.go productorRabbit.go servidorGym.go cypherGym.go

CMD ["./gym"]