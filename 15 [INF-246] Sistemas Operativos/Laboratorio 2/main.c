#include <stdio.h>
#include <sys/types.h>
#include <unistd.h>
#include <stdlib.h>
#include <sys/wait.h>

#include "herramientas.h"
#include "jugadores.h"
#include "laberinto.h"
#include "mapa.h"
#include "archivos.h"

int main(int argc, char *argv[]) {

    // Variables configurables
    int round_count = 15;

    // Valores constantes
    int player_count = 4;




    // ====== CODIGO DE INICIALIZACION ======

    printf("[%d] ====== INICIALIZANDO ======\n", getpid());

    // =========== CREACION DE OBJETOS ============

    // Inicializacion de jugadores
    t_Jugador** arr_jugadores;

    printf("[%d] INICIALIZANDO JUGADORES\n", getpid());

    // Crear jugadores
    arr_jugadores = malloc(player_count * sizeof(t_Jugador*));

    for (int i = 0; i < player_count; i++){

        // Si es jugador 0, entonces comienza en (1,1)
        // Si es jugador 1, entonces comienza en (3,1)
        // Si es jugador 2, entonces comienza en (1,3)
        // Si es jugador 3, entonces comienza en (3,3)

        int x = 1 + (i%2)*2;
        int y = 1 + (i/2)*2;

        printf("[%d] INICIALIZANDO jugador %d en posicion (%d, %d)\n", getpid(), i+1, x, y);

        arr_jugadores[i] = crearJugador(i, i%2, x, y);
        imprimirJugador(arr_jugadores[i]);
    }

    // Inicializacion de laberintos
    int n_archivos;
    char** lista_archivos = arregloArchivos("./laberinto", &n_archivos);
    printf("\nHay %d archivos\n", n_archivos);
    t_Laberinto** mazo_laberintos = malloc(n_archivos * sizeof(t_Laberinto*));
    for (int i = 0; i < n_archivos; i++) {
        printf("\nArchivo: %s\n", lista_archivos[i]);
        mazo_laberintos[i] = crearLaberinto(lista_archivos[i]);
        imprimirLaberinto(mazo_laberintos[i]);
    }

    // Crear mapa con los laberintos
    t_Mapa* mapa_laberintos = crearMapa(mazo_laberintos, n_archivos);
    imprimirMapa(mapa_laberintos);

    // =========== CREACION DE PROCESOS ============

    // Creando pipes
    int arr_pipe[player_count*2][2];

    for (int i = 0; i < player_count*2; i++){
        // printf("[%d] Creando pipe %d, para proceso %d\n", getpid(), i, i/2+1);

        if (pipe(arr_pipe[i]) < 0){
            perror("pipe");
            exit(EXIT_FAILURE);
        }
    }

    /*
    Estructura de los pipes:
    pipe 0: padre  (escritura) -> hijo 1 (lectura)
    pipe 1: hijo 1 (escritura) -> padre  (lectura)

    pipe 2: padre  (escritura) -> hijo 2 (lectura)
    pipe 3: hijo 2 (escritura) -> padre  (lectura)

    pipe 4: padre  (escritura) -> hijo 3 (lectura)
    pipe 5: hijo 3 (escritura) -> padre  (lectura)

    pipe 6: padre  (escritura) -> hijo 4 (lectura)
    pipe 7: hijo 4 (escritura) -> padre  (lectura)
    */





    // Creando procesos hijos (jugadores)
    int player_number = -1; // El padre no tiene numero de jugador, sera el anfitrion del juego
    pid_t pid_child[player_count]; // Arreglo de los pid de los hijos

    for (int i = 0; i < player_count; i++){

        printf("[%d] Creando proceso %d\n", getpid(), i+1);

        pid_child[i] = fork(); // Guardar el pid del hijo creado

        if (pid_child[i] < 0){
            perror("fork");
            exit(EXIT_FAILURE);
        }

        if (pid_child[i] == 0){ // El hijo creado ejecuta esto
            player_number = i; // El hijo estara marcado con su numero de jugador comenzando desde 0
            break; // El hijo no debe crear mas procesos
        }
    }





    // Inicializacion de pipes
    if (player_number == -1) { // El padre ejecuta esto

        printf("[%d] INICIALIZANDO PIPES de padre (player_number = %d)\n", getpid(), player_number);

        // Cerrando pipes que no se usaran
        for (int i = 0; i < player_count; i++){
            printf("[%d] Cerrando pipe %d modo %d\n", getpid(), i*2, 0);
            close(arr_pipe[i*2][0]);   // Si es par, cerrar pipe de lectura
            printf("[%d] Cerrando pipe %d modo %d\n", getpid(), i*2+1, 1);
            close(arr_pipe[i*2+1][1]); // Si es impar, cerrar pipe de escritura
        }

    } else { // Todos los hijos ejecutaran esto

        // Esperar entre 1 y 4 segundos
        sleep(player_number+1);

        printf("  [%d][J%d] INICIALIZANDO PIPES de hijo (player_number = %d)\n", getpid(), player_number+1, player_number);

        // Cerrando pipes que no se usaran
        printf("  [%d][J%d] Cerrando pipe %d modo %d\n", getpid(), player_number+1, player_number*2, 1);
        close(arr_pipe[player_number*2][1]);   // Si es par, cerrar pipe de escritura
        printf("  [%d][J%d] Cerrando pipe %d modo %d\n", getpid(), player_number+1, player_number*2+1, 0);
        close(arr_pipe[player_number*2+1][0]); // Si es impar, cerrar pipe de lectura
    }

    printf("[%d] ====== INICIALIZACION TERMINADA ======\n", getpid());

    // ====== INICIALIZACION LISTA ======






    // ====== CODIGO DEL JUEGO ======

    if (player_number == -1){ // El padre ejecuta esto

        // Esperar 5 segundos
        sleep(player_count+1);

        // Ordenes que se enviaran a los hijos
        int order_send_newTurn = 1;
        int order_send_finalize = -1;

        // Loop principal de rondas
        int resp[player_count];

        for (int round_actual = round_count; round_actual > 0; round_actual--){

            printf("[%d] ============ RONDA %d =============\n", getpid(), round_actual);



            for (int i = 0; i < player_count; i++) {
                printf("[%d] ======== TURNO JUGADOR %d ========\n", getpid(), i+1);

                // Enviar orden de turno al jugador i+1
                printf("[%d] Enviando orden de turno al jugador %d\n", getpid(), i+1);
                write(arr_pipe[i*2][1], &order_send_newTurn, sizeof(int));


                imprimirMapa(mapa_laberintos);
                imprimirLaberinto(mapa_laberintos->mtrx_mapa[arr_jugadores[i]->map_coord.y][arr_jugadores[i]->map_coord.x]);
                imprimirJugador(arr_jugadores[i]);

                printf("Ubicado sobre casilla tipo: %d\n\n", chequearCelda(mapa_laberintos, arr_jugadores[i]->map_coord.x, arr_jugadores[i]->map_coord.y, arr_jugadores[i]->lab_coord.x, arr_jugadores[i]->lab_coord.y));

                printf("[%d] Jugador %d, ingresa una orden:\n  0: avanzar\n  1: buscar/escalera\n", getpid(), i+1);

                // Esperar input del jugador i+1
                printf("[%d] Esperando input del jugador %d...\n", getpid(), i+1);
                read(arr_pipe[i*2+1][0], &resp[i], sizeof(int));

                sleep(1);

                printf("[%d] Recibido %d desde el jugador %d\n", getpid(), resp[i], i+1);

                // Dependiendo del input recibido, hacer algo
                if (resp[i] == 0) {
                    while (resp[i] == 0 || arr_jugadores[i]->seMovio == 1) { // Avanzar
                        printf("[%d] Orden 0: avanzar\n", getpid());
                        printf("[%d] Elije direccion:\n  0: N\n  1: O\n  2: S\n  3: E\n  4: Pasar turno\n", getpid());
                        write(arr_pipe[i*2][1], &order_send_newTurn, sizeof(int));

                        // Esperar input del jugador i+1
                        printf("[%d] Esperando input del jugador %d...\n", getpid(), i+1);
                        read(arr_pipe[i*2+1][0], &resp[i], sizeof(int));

                        // Dependiendo del input recibido, hacer algo
                        moverJugador(mapa_laberintos, arr_jugadores[i]->map_coord.x, arr_jugadores[i]->map_coord.y, arr_jugadores[i]->lab_coord.x, arr_jugadores[i]->lab_coord.y, arr_jugadores[i], resp[i]);


                        imprimirLaberinto(mapa_laberintos->mtrx_mapa[arr_jugadores[i]->map_coord.y][arr_jugadores[i]->map_coord.x]);
                        imprimirJugador(arr_jugadores[i]);
                        printf("Ubicado sobre casilla tipo: %d\n\n", chequearCelda(mapa_laberintos, arr_jugadores[i]->map_coord.x, arr_jugadores[i]->map_coord.y, arr_jugadores[i]->lab_coord.x, arr_jugadores[i]->lab_coord.y));

                        if (resp[i] == 4) {
                            printf("[%d] Orden 4: pasar turno\n", getpid());
                            break;
                        }
                    }
                } else if (resp[i] == 1) { // Buscar/puerta
                    printf("[%d] Orden 1: buscar/puerta\n", getpid());

                    int tipoCeldaActual = chequearCelda(mapa_laberintos, arr_jugadores[i]->map_coord.x, arr_jugadores[i]->map_coord.y, arr_jugadores[i]->lab_coord.x, arr_jugadores[i]->lab_coord.y);

                    if (tipoCeldaActual == 4) { // Pasaje
                        if (arr_jugadores[i]->rol == 0) { // El jugador que busca es un explorador, y solo el puede pasar por el pasaje
                            if (avanzarPasaje(mapa_laberintos, arr_jugadores[i]->map_coord.x, arr_jugadores[i]->map_coord.y, arr_jugadores[i]->lab_coord.x, arr_jugadores[i]->lab_coord.y, arr_jugadores[i], mazo_laberintos, n_archivos)) {
                                printf("[%d] Se avanzo por el pasaje\n", getpid());
                            }
                        } else {
                        printf("[%d] No se pudo avanzar por el pasaje\n", getpid());
                        }
                    } else if (tipoCeldaActual == 2) { // Puerta
                        if (arr_jugadores[i]->rol = 1) { // El jugador que tiene la llave puede abrir la puerta
                            if (abrirPuerta(mapa_laberintos, arr_jugadores[i]->map_coord.x, arr_jugadores[i]->map_coord.y, arr_jugadores[i]->lab_coord.x, arr_jugadores[i]->lab_coord.y)) {
                                printf("[%d] Se abrio la puerta\n", getpid());
                            }
                        } else {
                            printf("[%d] No se pudo abrir la puerta\n", getpid());
                        }
                    } else {
                        printf("[%d] No se encontro nada\n", getpid());
                    }

                    imprimirLaberinto(mapa_laberintos->mtrx_mapa[arr_jugadores[i]->map_coord.y][arr_jugadores[i]->map_coord.x]);
                    imprimirJugador(arr_jugadores[i]);
                    printf("Ubicado sobre casilla tipo: %d\n\n", chequearCelda(mapa_laberintos, arr_jugadores[i]->map_coord.x, arr_jugadores[i]->map_coord.y, arr_jugadores[i]->lab_coord.x, arr_jugadores[i]->lab_coord.y));
                }
            }





            printf("[%d] Las respuestas han sido %d, %d, %d, %d\n", getpid(), resp[0], resp[1], resp[2], resp[3]);
            printf("[%d] Terminando ronda...\n", getpid());
        }




        // El loop termina cuando se acaban las rondas
        printf("============ JUEGO TERMINADO ============\n");
        printf("[%d] Terminando juego...\n", getpid());

        // Enviar orden de finalizar proceso a los hijos
        for (int i = 0; i < player_count; i++){
            printf("[%d] Enviando orden de termino (-1) al hijo %d\n", getpid(), pid_child[i]);
            write(arr_pipe[i*2][1], &order_send_finalize, sizeof(int));
        }

        // Esperar 1 segundo
        sleep(1);

        // Esperar a que los hijos terminen
        for (int i = 0; i < player_count; i++){
            printf("[%d] Esperando a %d\n", getpid(), pid_child[i]);
            waitpid(pid_child[i], NULL, 0);
        }





    } else { // Todos los hijos ejecutaran esto

        int order_receive, input;

        while (1) { // Los hijos deben funcionar siempre, hasta que reciban la orden de terminar

            // Esperar orden desde el padre
            printf("  [%d][J%d] Esperando orden...\n", getpid(), player_number+1);
            read(arr_pipe[player_number*2][0], &order_receive, sizeof(int));

            sleep(1);


            // Dependiendo de la orden recibida, hacer algo
            printf("  [%d][J%d] Recibida orden %d desde el padre\n", getpid(), player_number+1, order_receive);

            if (order_receive == -1) {
                printf("  [%d][J%d] Orden -1: terminar proceso...\n", getpid(), player_number+1);
                break;

            } else if (order_receive == 1) {
                printf("  [%d][J%d] Orden 1: comenzar turno...\n", getpid(), player_number+1);

                // Recibir input del usuario
                printf("  [%d][J%d] Ingrese un numero:\n>>> ", getpid(), player_number+1);
                scanf("%d", &input);

                // Enviar input al padre
                printf("  [%d][J%d] Enviando %d al padre\n", getpid(), player_number+1, input);
                write(arr_pipe[player_number*2+1][1], &input, sizeof(int));

                printf("  [%d][J%d] Terminando turno...\n", getpid(), player_number+1);

            } else {
                printf("  [%d][J%d] Orden %d: desconocida. No hacer nada.\n", getpid(), player_number+1, order_receive);
            }
        }
    }

    // Liberar memoria
    for (int i = 0; i < player_count; i++){
        liberarJugador(arr_jugadores[i]);
    }
    free(arr_jugadores);

    for (int i = 0; i < n_archivos; i++){
        liberarLaberinto(mazo_laberintos[i]);
    }
    free(mazo_laberintos);

    liberarArregloArchivos(lista_archivos, n_archivos);

    liberarMapa(mapa_laberintos);

	return 0;
}
