# Otras tareas
up:
	docker-compose up -d

down:
	docker-compose down

build:
	docker-compose build

restart:
	docker-compose down
	docker-compose up -d

pull:
	docker pull cassandra:latest

bash1:
	docker exec -it nodo1 bash

bash2:
	docker exec -it nodo2 bash

bash3:
	docker exec -it nodo3 bash

run: build up

create-keyspace:
	docker exec -it nodo1 cqlsh -e "CREATE KEYSPACE IF NOT EXISTS universia_keyspace WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 3};"

create-tables:
	docker exec -it nodo1 cqlsh -e "USE universia_keyspace; CREATE TABLE IF NOT EXISTS postulantes_por_carrera1 (periodo INT,carrera TEXT, puntaje_psu INT, region TEXT, sexo TEXT, cedula INT, preferencia INT, matriculado TEXT, facultad TEXT, puntaje INT, grupo_depen TEXT, latitud FLOAT, longitud FLOAT, puntaje_nem INT, pace VARCHAR, gratuidad TEXT, PRIMARY KEY ((carrera), matriculado, periodo, region, cedula));"

	docker exec -it nodo1 cqlsh -e "USE universia_keyspace; CREATE TABLE IF NOT EXISTS postulantes_por_carrera2 (periodo INT,carrera TEXT, puntaje_psu INT, region TEXT, sexo TEXT, cedula INT, preferencia INT, matriculado TEXT, facultad TEXT, puntaje INT, grupo_depen TEXT, latitud FLOAT, longitud FLOAT, puntaje_nem INT, pace VARCHAR, gratuidad TEXT, PRIMARY KEY ((facultad), matriculado,puntaje_psu, cedula));"

query1:
	docker exec -it nodo1 cqlsh -e "USE universia_keyspace;SELECT periodo, carrera, cedula , matriculado FROM postulantes_por_carrera1 WHERE facultad = 'MEDICINA' AND periodo IN (2015, 2016, 2017)AND carrera = 'MEDICINA'AND matriculado = 'SI'ALLOW FILTERING"

query2:
	docker exec -it nodo1 cqlsh -e "USE universia_keyspace; SELECT periodo, carrera, cedula , matriculado , region FROM postulantes_por_carrera1 WHERE matriculado = 'SI' AND region = 'MAULE' AND carrera = 'INGENIERÍA CIVIL INFORMÁTICA'ALLOW FILTERING"

query3:
	docker exec -it nodo1 cqlsh -e "USE universia_keyspace;SELECT periodo, facultad, carrera, cedula , matriculado, puntaje_psu FROM postulantes_por_carrera2 WHERE matriculado = 'SI' AND facultad = 'CIENCIAS DE LA SALUD'ALLOW FILTERING"

