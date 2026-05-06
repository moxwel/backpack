FROM golang:1.23.9

# Dependencias
WORKDIR /app/types
COPY /types ./

WORKDIR /app/lcp
COPY /lcp/go.mod ./
COPY /lcp/go.sum ./
RUN go mod download

# Codigo fuente
COPY /lcp ./
RUN go build -o lcp main.go servidorLCP.go productorRabbit.go consumerRabbitCDP.go clienteGym.go

CMD ["./lcp"]