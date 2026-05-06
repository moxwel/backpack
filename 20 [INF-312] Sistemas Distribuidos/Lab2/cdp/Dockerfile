FROM golang:1.23.9

# Dependencias
WORKDIR /app/types
COPY /types ./

WORKDIR /app/cdp
COPY /cdp/go.mod ./
COPY /cdp/go.sum ./
RUN go mod download

# Codigo fuente
COPY /cdp ./
RUN go build -o cdp main.go productorRabbit.go consumerRabbitGym.go decypherGym.go clienteLCP.go

CMD ["./cdp"]