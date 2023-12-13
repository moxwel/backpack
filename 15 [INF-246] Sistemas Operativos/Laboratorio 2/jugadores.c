#include "jugadores.h"
#include "mapa.h"

t_Jugador* crearJugador(int numero, int rol, int x, int y) {
    t_Jugador* nuevo_jugador = malloc(sizeof(t_Jugador));

    nuevo_jugador->numero = numero;
    nuevo_jugador->rol = rol;
    nuevo_jugador->map_coord.x = 5;
    nuevo_jugador->map_coord.y = 5;
    nuevo_jugador->lab_coord.x = x;
    nuevo_jugador->lab_coord.y = y;
    nuevo_jugador->seMovio = 0;
    nuevo_jugador->tieneTesoro = 0;

    return nuevo_jugador;
}

void liberarJugador(t_Jugador* jugador) {
    free(jugador);
}

void imprimirJugador(t_Jugador* jugador) {
    printf(" --- Jugador %d ---\n Rol %d (%s)\n Posicion en mapa: (%d, %d)\n Posicion en laberinto: (%d, %d)\n Se movio: %d\n Tiene tesoro: %d\n ------------- \n",
        jugador->numero+1,
        jugador->rol,
        jugador->rol == 0 ? "Buscar" : "Llave",
        jugador->map_coord.x,
        jugador->map_coord.y,
        jugador->lab_coord.x,
        jugador->lab_coord.y,
        jugador->seMovio,
        jugador->tieneTesoro
    );
}
