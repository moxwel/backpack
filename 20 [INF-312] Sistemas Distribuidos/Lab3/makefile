PROTO_FILES=tipos.proto mensajes.proto
PROTO_DIR=./proto
PROTO_OUT_DIR=./proto/mensajes
SERVICE_DIR=MatchmakingProto
TARGET_DIRS=./admin ./matchmaker ./jugador ./servidor

protoFiles:
	cd $(PROTO_DIR) && \
	protoc --go_out=. --go-grpc_out=. $(PROTO_FILES) && \
	cd .. && \
	for dir in $(TARGET_DIRS); do \
		cp -rf $(PROTO_OUT_DIR)/$(SERVICE_DIR) $$dir; \
	done

docker-matchmaking:
	docker-compose up --no-deps matchmaker

docker-admin:
	docker-compose run --no-deps admin_client

docker-jugadores:
	docker-compose run --no-deps player_client1 player_client2

docker-jugador1:
	docker-compose run --no-deps player_client1
	
docker-jugador2:
	docker-compose run --no-deps player_client2

docker-servidores:
	docker-compose up --no-deps gameserver1 gameserver2 gameserver3

docker-servidor1:
	docker-compose up --no-deps gameserver1

docker-servidor2:
	docker-compose up --no-deps gameserver2

docker-servidor3:
	docker-compose up --no-deps gameserver3

build:
	docker-compose build

build-matchmaker:
	docker-compose build matchmaker

build-admin:
	docker-compose build admin_client

build-jugadores:
	docker-compose build player_client1 player_client2

build-jugador1:
	docker-compose build player_client1

build-jugador2:
	docker-compose build player_client2	

build-servidores:
	docker-compose build gameserver1 gameserver2 gameserver3

build-servidor1:
	docker-compose build gameserver1

build-servidor2:
	docker-compose build gameserver2

build-servidor3:
	docker-compose build gameserver3

build:
	docker-compose build

rename-docker-compose:
	mv docker-compose.yml docker-compose_old.yml
	mv docker-compose_real.yml docker-compose.yml
