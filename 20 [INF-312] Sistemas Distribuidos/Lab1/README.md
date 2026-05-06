# Grupo-19

## Integrantes

* Maximiliano Sepúlveda Alvear.
* Felipe Mellado Olea.

## Construir imágenes

Para construir todas las imágenes de cada servicio se debe ejecutar el siguiente comando:

```bash
make build
```

Para construir una imagen en especifico se debe ejecutar el siguiente comando:

```bash
make build-SERVICIO
```

Donde `SERVICIO` puede ser `cazarec`, `marine`, `submun` o `gobmun`.

Ejemplo, construir la imagen de cazarrecompensas:

```bash
make build-cazarec
```

## Ejecutar contenedores

Para ejecutar todos los contenedores de cada servicio se debe ejecutar el siguiente comando:

```bash
make docker
```

Para ejecutar un contenedor en especifico se debe ejecutar el siguiente comando:

```bash
make docker-SERVICIO
```

Donde `SERVICIO` puede ser `cazarec`, `marine`, `submun` o `gobmun`.

Ejemplo, ejecutar el contenedor de cazarrecompensas:

```bash
make docker-cazarec
```

## Configuración de variables de entorno

En el archivo `docker-compose.yml` se encuentran las variables de entorno necesarias para cada servicio.

Cada servicio "servidor" posee dos variables:

* `IP_ADDR`: Corresponde a la direccion IP en donde va a escuchar el servicio. Por defecto es 0.0.0.0 (todas las interfaces de red).
* `PORT_NUMBR`: Corresponde al puerto en donde va a escuchar el servicio.

La configuración de cada servicio es la siguiente:

* Submundo: puerto 50051
* Marina: puerto 50052
* Gobierno Mundial: puerto 50053

## Maquinas virtuales

Cada maquina virtual tiene una IP diferente y un contenedor diferente:

* dist073: IP=10.35.168.83, Contenedor: cazarec
* dist074: IP=10.35.168.84, Contenedor: submun
* dist075: IP=10.35.168.85, Contenedor: marine
* dist076: IP=10.35.168.86, Contenedor: gobmun

Cada maquina deberia ejecutar su propio contenedor correspondiente. Deben utilizarse los comandos mencionados anteriormente para construir y ejecutar los contenedores.

Las maquinas virtuales estan configuradas con el siguiente docker-compose:

```yaml
services:
    submun:
        build: ./submun
        container_name: submun
        ports:
            - "50051:50051"
        environment:
            - IP_ADDR=0.0.0.0
            - PORT_NUMBR=50051
    marine:
        build: ./marine
        container_name: marine
        ports:
            - "50052:50052"
        environment:
            - IP_ADDR=0.0.0.0
            - PORT_NUMBR=50052

            - GOB_IP_ADDR=10.35.168.86
            - GOB_PORT_NUMBR=50053

            - SUB_IP_ADDR=10.35.168.84
            - SUB_PORT_NUMBR=50051
    gobmun:
        build: ./gobmun
        container_name: gobmun
        ports:
            - "50053:50053"
        environment:
            - IP_ADDR=0.0.0.0
            - PORT_NUMBR=50053

            - MAR_IP_ADDR=10.35.168.85
            - MAR_PORT_NUMBR=50052

            - PIRATAS_FILE=piratas_grande.csv

    cazarec:
        build: ./cazarec
        container_name: cazarec
        environment:
            - SUB_IP_ADDR=10.35.168.84
            - SUB_PORT_NUMBR=50051

            - MAR_IP_ADDR=10.35.168.85
            - MAR_PORT_NUMBR=50052

            - GOB_IP_ADDR=10.35.168.86
            - GOB_PORT_NUMBR=50053

            - ITER_NUMBR=10
```

## Lógica del Sistema

- El programa inicia con el cazarrecompensas y se ejecuta por un número de iteraciones definido en la variable `ITER_NUMBR` (por defecto: `10`).

- En cada iteración:
  - El cazarrecompensas solicita al Gobierno Mundial la lista actual de piratas buscados.
  - Se selecciona aleatoriamente un pirata que no esté capturado.
  - Luego, se decide al azar si el pirata será entregado a la Marina (entrega legal) o al Submundo (venta ilegal).

- Cada cazarrecompensas mantiene su propio balance económico, el cual refleja:
  - Qué tan frecuentemente opera legal o ilegalmente.
  - Si comercia demasiado en el Submundo, se le aplica una penalización en el valor de las recompensas legales.

- El Gobierno Mundial tiene un balance global del sistema:
  - Si este balance baja de un cierto umbral, se activan redadas automáticas (funcionan como un "interruptor").
  - Las redadas continúan activas hasta que el balance se restablezca.

- Los piratas pueden tener los siguientes estados:
  - `En Captura`: listo para transportar
  - `Buscado`: disponible para captura.
  - `Confiscado`: la marina lo capturó en una redada
  - `Capturado`: entregado con éxito.
  - `Perdido`: se escapó o fue rescatado por mercenarios.


- El costo de transporte por cada pirata entregado es del 5% de su recompensa original, independientemente del destino.
