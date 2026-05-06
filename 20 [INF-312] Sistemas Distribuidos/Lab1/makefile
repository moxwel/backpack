PROTO_FILE=./proto/mensajes.proto
PROTO_OUT_DIR=./proto/mensajes
SERVICE_DIR=servicioMensajes
TARGET_DIRS=./gobmun ./marine ./submun ./cazarec

protoFiles:
	protoc --go_out=. --go-grpc_out=. $(PROTO_FILE)
	@for dir in $(TARGET_DIRS); do \
		cp -rf $(PROTO_OUT_DIR)/$(SERVICE_DIR) $$dir; \
	done

goMod:
	@for dir in $(TARGET_DIRS); do \
		rm -f $$dir/go.mod $$dir/go.sum; \
		(cd $$dir && go mod init $$(basename $$dir) && go mod tidy); \
	done

docker-submun:
	docker-compose up submun

docker-marine:
	docker-compose up marine

docker-gobmun:
	docker-compose up gobmun

docker-cazarec:
	docker-compose up cazarec

docker:
	docker-compose up

docker-submun-detach:
	docker-compose up -d submun

docker-marine-detach:
	docker-compose up -d marine

docker-gobmun-detach:
	docker-compose up -d gobmun

docker-cazarec-detach:
	docker-compose up -d cazarec

docker-detach:
	docker-compose up -d 

build-gobmun:
	docker-compose build gobmun

build-submun:
	docker-compose build submun

build-marine:
	docker-compose build marine

build-cazarec:
	docker-compose build cazarec

build:
	docker-compose build