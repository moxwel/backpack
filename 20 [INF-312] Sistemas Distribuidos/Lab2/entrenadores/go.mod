module entrenadores

go 1.23

toolchain go1.23.9

replace l2/types => ../types

require (
	github.com/streadway/amqp v1.1.0
	google.golang.org/grpc v1.72.1
	google.golang.org/protobuf v1.36.6
	l2/types v0.0.0-00010101000000-000000000000
)

require (
	golang.org/x/net v0.35.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250218202821-56aae31c358a // indirect
)
