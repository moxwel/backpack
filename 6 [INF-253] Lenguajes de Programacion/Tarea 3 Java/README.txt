README
Maximiliano Sepúlveda
ROL #########-#
====================

NOTAS AL MARGEN
---------------

No se si lo entendi bien, pero los objetos "activos" los interpreté como "consumibles", es decir, el
objeto solo se podra utilizar una vez, y los efectos son permanentes.

Algunas personas creeran que esta mecanica esta rota, pero la verdad, no. El que los ataques sean
siempre 50% de acertar puede arruinar una run completa sin importar cuanto daño tengas (viva el RNG).

El programa lo hice en un entorno de Windows. Makefile utiliza comandos de Windows para 'clean'.
No es que afecte mucho la tarea, pero puede ser algo a tener en consideracion.





COMPILACION Y EJECUCION
-----------------------

Para compilar el programa, se debe ejecutar el comando make.

    $ make
    Compila el programa, los archivos '.class' se guardan en una carpeta 'build'.

    $ make jar
    Compila el programa, y lo construye a un '.jar'.

    $ make run
    Ejecuta el '.jar'.





JUGABILIDAD
-----------

Al iniciar el juego, lo primero que pregunta es el personaje a utilizar. Ingresa el numero
correspondiente al jugador que vas a usar.

        Elige aquel quien va a estar a tu merced...
        [1] Don Mario
        [2] Don Sora
        [3] Don NoMuerto

        >>>1

        ===== JUGADOR =====
        Nom: Don Mario
        Crz: 9
        Dno: 3
        Vzd: 3
        Crt: 0
        Itm: Champinon Recursivo [Usable]
        ===================

        Generando mapa...

        |===== MAP =====|
        |               |
        | .  .          |
        | .  .          |
        | .  .          |
        | .  .  .  @    |
        |===============|

        Ubicacion: [4][3] - Tipo: Vacia
        Que hacer? [help]
        >>>_

Cuando veas los simbolos '>>>', significa que esta esperando una entrada. Para saber los comandos
disponibles, ingresa el comando 'help'.

Existe una mecanica de puntajes: cuando se llega a la meta, el juego se reinicia y vuelve a
generar un nuevo mapa, aumentando un contador de puntaje. Intenta pasar todas las salas que puedas.





MAPA
----

        |===== MAP =====|
        |               |
        | .  .          |
        | .  .          |
        | .  .          |
        | .  .  .  @    |
        |===============|

Los simbolos del mapa son los siguientes:

    @ = Tu ubicacion
    . = Sala sin visitar aun
    O = Sala vacia
    C = Cofre
    T = Tesoro
    E = Enemigos
    # = Meta

Cuando se llega a una sala especial, debes ejecutar el comando correspondiente a la sala. Es decir,
si estas en una sala de cofre, debes ejecutar el comando 'cofre' para poder abrirlo; si estas en
la sala de la meta, debes ejecutar el comando 'meta' para superar el nivel; y asi.





BATALLAS
--------

        Hay 3 enemigo/s!
        Te encuentras con 'Leakitu'

        ===== JUGADOR =====
        Nom: Don Mario
        Crz: 9
        Dno: 3
        Vzd: 3
        Crt: 0
        Itm: Champinon Recursivo [Usable]
        ===================


        ===== ENEMIGO =====
        Nom: Leakitu
        Crz: 7
        Dno: 1
        Vzd: 2
        Crt: 0
        ===================

        Que hacer? [atacar | usar | escapar]
        >>>_

En las batallas, existe la opcion de poder "escapar" de un enemigo, sin embargo, solo tienes
un 10% de exito. Si logras escapar de un enemigo, el enemigo actual "muere", pero los restantes
no te atacaran. Con esto, tienes la oportunidad de moverte a otra sala, sin embargo, si vuelves a entrar
a la sala, vas a tener que pelear con los enemigos que quedaron (o volver a escapar).

En las batallas, cualquier accion que no sea "atacar", causara que pierdas un turno. Es decir,
si en vez de atacar, decides "usar" tu item en la mano, entonces perderas tu turno de atacar y
el enemigo atacara. Esto tambien aplica a la opcion de "escape": si tu escape falla, perderas el
turno de atacar, y el enemigo atacara.
