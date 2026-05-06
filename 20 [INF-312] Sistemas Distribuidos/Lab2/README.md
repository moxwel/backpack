# Grupo-19

## Integrantes

* Maximiliano Sepúlveda Alvear.
* Felipe Mellado Olea.

## Construir imágenes

Para construir todas las imágenes de cada servicio se debe ejecutar el siguiente comando:

```bash
make build
```

## Ejecutar contenedores

Para ejecutar un contenedor en específico se debe ejecutar el siguiente comando:

```bash
make docker-SERVICIO
```

Donde `SERVICIO` puede ser `docker-cdp-gym`, `docker-lcp`, `docker-snp-rabbitmq` o `docker-entrenadores`.

Ejemplo, ejecutar el contenedor de cdp:

```bash
make docker-cdp-gym
```

## Configuración de variables de entorno

En el archivo `docker-compose.yml` se encuentran las variables de entorno necesarias para cada servicio.

Cada servicio "servidor" posee dos variables:

* `IP`: Corresponde a la dirección IP en donde va a escuchar el servicio. Por defecto es 0.0.0.0
* `PORT`: Corresponde al puerto en donde va a escuchar el servicio

La configuración de cada servicio es la siguiente (puesta por nosotros):

* RABBIT: 15672/5672
* LCP: 50051
* GYM: 50052

También considerar las 3 llaves AES de cada gimnasio la cual debe coincidir tanto en CDP como en el Gimnasio correspondiente

Las llaves de ejemplo que usamos son:

* AES_KEY_KANTO=12345678901234567890123456789012
* AES_KEY_JOHTO=abcdefghijklmnopqrstuvwxzy123456
* AES_KEY_SINNOH=09876543210987654321098765432109

## Maquinas virtuales

Cada máquina virtual tiene una IP diferente y un contenedor diferente :

* dist073 10.35.168.83; snp rabbit
* dist074 10.35.168.84; cdp gym
* dist075 10.35.168.85; lcp
* dist076 10.35.168.86; entrenador

Un ejemplo de configuración en `docker-compose.yml` podría ser:

```yaml
services:
    rabbitmq:
        image: rabbitmq:4-management
        container_name: rabbitmq
        ports:
            - "15672:15672"
            - "5672:5672"
        environment:
            - RABBITMQ_DEFAULT_USER=guest
            - RABBITMQ_DEFAULT_PASS=guest

    entrenadores:
        build:
            context: .
            dockerfile: entrenadores/Dockerfile
        container_name: entrenadores
        environment:
            - LCP_IP=10.35.168.85
            - LCP_PORT=50051

            - RABBIT_IP=10.35.168.83
            - RABBIT_PORT=5672
            - RABBIT_USER=guest
            - RABBIT_PASS=guest
        depends_on:
            - lcp
            - rabbitmq

    snp:
        build:
            context: .
            dockerfile: snp/Dockerfile
        container_name: snp
        environment:
            - RABBIT_IP=10.35.168.83
            - RABBIT_PORT=5672
            - RABBIT_USER=guest
            - RABBIT_PASS=guest
        depends_on:
            - rabbitmq

    lcp:
        build:
            context: .
            dockerfile: lcp/Dockerfile
        container_name: lcp
        environment:
            - IP=0.0.0.0
            - PORT=50051

            - GYM_IP=10.35.168.84
            - GYM_PORT=50052

            - RABBIT_IP=10.35.168.83
            - RABBIT_PORT=5672
            - RABBIT_USER=guest
            - RABBIT_PASS=guest
        ports:
            - "50051:50051"
        depends_on:
            - gym
            - rabbitmq

    cdp:
        build:
            context: .
            dockerfile: cdp/Dockerfile
        container_name: cdp
        environment:
            - LCP_IP=10.35.168.85
            - LCP_PORT=50051

            - RABBIT_IP=10.35.168.83
            - RABBIT_PORT=5672
            - RABBIT_USER=guest
            - RABBIT_PASS=guest

            - AES_KEY_KANTO=12345678901234567890123456789012
            - AES_KEY_JOHTO=abcdefghijklmnopqrstuvwxzy123456
            - AES_KEY_SINNOH=09876543210987654321098765432109
        depends_on:
            - lcp
            - rabbitmq

    gym:
        build:
            context: .
            dockerfile: gym/Dockerfile
        container_name: gym
        environment:
            - IP=0.0.0.0
            - PORT=50052

            - RABBIT_IP=10.35.168.83
            - RABBIT_PORT=5672
            - RABBIT_USER=guest
            - RABBIT_PASS=guest

            - AES_KEY_KANTO=12345678901234567890123456789012
            - AES_KEY_JOHTO=abcdefghijklmnopqrstuvwxzy123456
            - AES_KEY_SINNOH=09876543210987654321098765432109
        ports:
            - "50052:50052"
        depends_on:
            - rabbitmq
```

## Lógica del Sistema

- El sistema inicia con la creación de torneos y entrenadores, y se ejecuta por un número de iteraciones definido en la configuración.
- Los entrenadores pueden inscribirse en torneos, y los combates se asignan y resuelven a través de los servicios.
- Los servicios se comunican mediante gRPC y colas RabbitMQ para coordinar eventos, resultados y notificaciones.
- El sistema maneja penalizaciones, suspensiones y expulsiones de entrenadores según las reglas de la competencia.
- Si el combate dura menos de 5 seg se penalizará al ganador
- Con 3 penalizaciones es suspención permanente

## Ejemplo de funcionamiento

Inicio de entrenadores:

```
    89. Stacy Allen (Región: sinnoh, Ranking: 1750, Estado: Suspendido, Suspension: 5)
    90. Jason Morris (Región: johto, Ranking: 1457, Estado: Activo, Suspension: 0)
    91. Anne Reynolds (Región: kanto, Ranking: 1475, Estado: Inactivo, Suspension: 0)
    92. Daniel Odonnell (Región: hoenn, Ranking: 1398, Estado: Activo, Suspension: 0)
    93. Brooke Curry (Región: hoenn, Ranking: 1130, Estado: Activo, Suspension: 0)
    94. Julia Harris (Región: unova, Ranking: 1272, Estado: Inactivo, Suspension: 0)
    95. Andrea Richards (Región: sinnoh, Ranking: 1701, Estado: Inactivo, Suspension: 0)
    96. Brent Schwartz (Región: sinnoh, Ranking: 1531, Estado: Suspendido, Suspension: 1)
    97. Sonya Martin (Región: unova, Ranking: 1662, Estado: Activo, Suspension: 0)
    98. Rodney Wells (Región: hoenn, Ranking: 1245, Estado: Activo, Suspension: 0)
    99. Brian Copeland (Región: johto, Ranking: 1466, Estado: Activo, Suspension: 0)
    100. Jermaine Rosales (Región: sinnoh, Ranking: 1691, Estado: Activo, Suspension: 0)
Ingrese el número del entrenador: _
```

Menu principal:

```
==== MENU PRINCIPAL - 096 Brent Schwartz (sinnoh) Activo (0) ====
    1. Mostrar estado de entrenadores
    2. Ver torneos disponibles
    3. Inscribirse a un torneo
    4. Ver historial
    5. Ver notificaciones
    6. Cambiar entrenador activo
    0. Salir
Ingrese una opción: _
```
