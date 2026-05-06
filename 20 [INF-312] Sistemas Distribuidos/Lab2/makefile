PROTO_FILES=tipos.proto mensajes.proto
PROTO_DIR=./proto
PROTO_OUT_DIR=./proto/mensajes
SERVICE_DIR=PokemonProto
TARGET_DIRS=./cdp ./entrenadores ./lcp ./snp ./gym

protoFiles:
	cd $(PROTO_DIR) && \
	protoc --go_out=. --go-grpc_out=. $(PROTO_FILES) && \
	cd .. && \
	for dir in $(TARGET_DIRS); do \
		cp -rf $(PROTO_OUT_DIR)/$(SERVICE_DIR) $$dir; \
	done

docker-lcp:
	docker-compose up --no-deps lcp

docker-cdp-gym:
	docker-compose up --no-deps cdp gym

docker-snp-rabbitmq:
	docker-compose up --no-deps snp rabbitmq

docker-entrenadores:
	docker-compose run --no-deps entrenadores

build:
	docker-compose build

rename-docker-compose:
	mv docker-compose.yml docker-compose_old.yml
	mv docker-compose_real.yml docker-compose.yml
