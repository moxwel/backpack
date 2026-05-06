FROM golang:1.23.9

# Dependencias
WORKDIR /app/types
COPY /types ./

WORKDIR /app/snp
COPY /snp/go.mod ./
COPY /snp/go.sum ./
RUN go mod download

# Codigo fuente
COPY /snp ./
RUN go build -o snp main.go productorRabbit.go consumerRabbitCDP.go consumerRabbitLCP.go

CMD ["./snp"]