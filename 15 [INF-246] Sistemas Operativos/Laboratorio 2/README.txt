    README

Maximiliano Sepúlveda
Antonia Zuñiga

====================

    ARCHIVOS

- main.c         : programa principal que ejecuta el juego
- archivos.c     : contiene las funciones para manipular archivos
- jugadores.c    : contiene las funciones para manipular jugadores
- laberinto.c    : contiene las funciones para manipular el laberinto
- mapa.c         : contiene las funciones para manipular el mapa
- herramientas.c : contiene las funciones para manipular el juego

    ARCHIVOS DE PRUEBA

- main2.c
- moverJugador.c

====================

    INSTRUCCIONES PARA COMPILAR

El codigo 'main.c' contiene el codigo principal.

Para compilar, ejecutar en terminal:

$ make

Se generará un ejecutable llamado 'magic_maze.out'. Para ejecutar:

$ ./magic_maze.out

====================

    FUNCIONAMIENTO

Al iniciar el programa, comenzará la fase de inicialización. Aqui se crean los procesos y objetos que
se utilizarán durante el juego.

Cada mensaje importante tiene incluido el PID del proceso desde donde se envía. Ejemplo:

            [27896] ====== INICIALIZANDO ======
            [27896] INICIALIZANDO JUGADORES
            [27896] INICIALIZANDO jugador 1 en posicion (1, 1)
            [27896] INICIALIZANDO jugador 2 en posicion (3, 1)
            [27896] INICIALIZANDO jugador 3 en posicion (1, 3)
            [27896] INICIALIZANDO jugador 4 en posicion (3, 3)
            [27896] Creando proceso 1
            [27896] Creando proceso 2
            [27896] Creando proceso 3
            [27896] Creando proceso 4
            [27896] INICIALIZANDO PIPES de padre (player_number = -1)
              [27897][J1] INICIALIZANDO PIPES de hijo (player_number = 0)
              [27897][J1] Cerrando pipe 0 modo 1
              [27897][J1] Cerrando pipe 1 modo 0
              [27897][J1] Esperando orden...
              [27898][J2] INICIALIZANDO PIPES de hijo (player_number = 1)
              [27898][J2] Cerrando pipe 2 modo 1
              [27898][J2] Cerrando pipe 3 modo 0
              [27898][J2] Esperando orden...
            . . .

Cada proceso tiene asociado un numero de jugador, que va desde 0 a 3. El proceso padre tiene player_number = -1
y es quien actua de anfitrion del juego.

Cuando los procesos estan listos, el padre envia una "señal" al hijo para que comience su turno. El hijo recibe
la señal y comienza a jugar. El padre espera a que el hijo termine su turno para enviar la señal al siguiente jugador.

            [27896] ============ RONDA 1 =============
            [27896] ======== TURNO JUGADOR 1 ========
            [27896] Enviando orden de turno al jugador 1
            [27896] Jugador 1, ingresa una orden:
              0: avanzar
              1: buscar/escalera
            [27896] Esperando input del jugador 1...
              [27897][J1] Recibida orden 1 desde el padre
              [27897][J1] Orden 1: comenzar turno...
              [27897][J1] Ingrese un numero:
            >>> _
            . . .

Cuando aparezca el simbolo '>>>' en la terminal, el jugador debe ingresar un numero.




    SOBRE LAS REGLAS DEL JUEGO

Consideramos las 4 cartas (2 de buscar y 2 de escalera) como "roles" que se asignan a los jugadores.

Un "mapa" es una matriz de "laberintos".
Cada laberinto es una matriz de "casillas".
Cada casilla tiene un "tipo" (pared, pasillo, escalera, etc).

Cada jugador tiene un "rol", una posicion (x, y) en el mapa y una posicion (x, y) en el laberinto correspondiente.




    COMO LEER LA INFORMACION

Al inicio y termino de cada ronda, se imprime el mapa y el laberinto donde se encuentra el jugador.

Las coordenadas (x, y) se cuentan de izquierda a derecha y desde arriba hacia abajo respectivamente,
comenzando en (0, 0) y terminando en (10,10) para el mapa y (4, 4) para el laberinto.

             --- Mapa ---
            - - - - - - - - - - -
            - - - - - - - - - - -
            - - - - - - - - - - -
            - - - - - - - - - - -
            - - - - - - - - - - -
            - - - - - 0 - - - - -
            - - - - - - - - - - -
            - - - - - - - - - - -
            - - - - - - - - - - -
            - - - - - - - - - - -
            - - - - - - - - - - -
             ------------
             --- Laberinto ---
              0 0 4 0 0
              0 0 0 0 0
              4 0 0 0 4
              0 0 0 0 0
              0 0 4 0 0
            Es inicial: 1
            Tipo de salida: 0
            Esta usado: 1
             ----------------
             --- Jugador 1 ---
             Rol 0 (Buscar)
             Posicion en mapa: (5, 5)
             Posicion en laberinto: (1, 1)
             Se movio: 0
             Tiene tesoro: 0
             -------------
            Ubicado sobre casilla tipo: 0




    SINCRONIZACION DE PROCESOS

Se decidio por un diseño donde el proceso padre toma control de la ejecucion del juego, y los procesos hijos
solo ejecutan sus turnos cuando el padre se los indica. Esto se lleva a cabo mediante el uso de pipes.

Se crean 8 pipes, 2 para cada proceso. Uno para enviar ordenes desde el padre al hijo, y otro para enviar
respuestas desde el hijo al padre.

Cuando se crean los procesos, los hijos quedan en espera de una orden del padre.

Una vez que el padre este listo, envia una orden al primer hijo para comenzar su turno. Luego de eso
el padre queda en espera de la respuesta del hijo.

El hijo recibe la orden del turno y comienza a jugar. Cuando termina, envia su respuesta al padre y queda
nuevamente en espera de una nueva orden.

El padre recibe la respuesta del hijo y la procesa. Luego envia la siguiente orden al hijo que corresponda y
queda nuevamente en espera de la respuesta del hijo.

Este ciclo se repite tantas veces como rondas haya en el juego.




    FINALIZACION DE PROCESOS

Los procesos hijos se ejecutaran indefinidamente hasta que reciban la orden especial de terminar desde el padre.

Cuando los procesos hijos terminan, liberan memoria y detienen su ejecucion.

Cuando se acaban las rondas, el proceso padre envia la señal de termino a todos los hijos, y luego queda en espera
a que estos terminen. Cuando todos los hijos terminan, el padre libera memoria y detiene su ejecucion.

====================

    CONSIDERACIONES

- Para evitar prints impredecibles, se utilizó la funcion sleep() para pausar la ejecucion del programa.

- En vez de haber 4 cartas con uno solo por jugador, se decidio por un diseño mas simple en donde las cartas
son "roles". Estos son asignadas a cada jugador al momento de inicializar el juego.

- Las funcion de "escaleras" se renombró a "puerta/llave". Aquel jugador que tiene la llave puede abrir puertas.

- En vez de que el jugador 1 controle a todos los bots, se decidio por un diseño mas simple en donde cada jugador
se puede controlar de manera individual, y estos se comunican con el padre.

- Se crean 4 procesos, cada uno siendo un jugador. El proceso padre actua de anfitrion del juego y es quien maneja
las reglas y define los turnos de cada jugador.

- El rol de cada jugador se muestra en cada turno, cuando se imprime al jugador por pantalla.

- La posicion de los jugadores se muestra como coordenadas (x, y), no se muestra en la representacion del mapa.

- Se asume que el usuario siempre ingresa datos validos.

- Se asume que el laberinto nombrado 'Inicio.txt' es el laberinto inicial.

- Se asume que todos los laberintos estan ubicados en la carpeta 'laberinto'.
