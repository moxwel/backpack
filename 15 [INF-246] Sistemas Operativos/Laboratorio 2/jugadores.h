#ifndef __JUGADOR__
#define __JUGADOR__

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <dirent.h>
#include <sys/types.h>
#include <sys/stat.h>

typedef struct {
    int x;
    int y;
} t_Posicion;

/**
 * Roles de los jugadores:
 * 0: Buscar
 * 1: Llave
*/
typedef struct {
    int seMovio;
    int numero;
    int rol;
    int tieneTesoro;
    t_Posicion map_coord;
    t_Posicion lab_coord;
} t_Jugador;

/**
 * t_jugador* crearJugador
 *
 * Crea un jugador con los datos entregados.
 * Se asume que el jugador comienza siempre en el laberinto central del mapa. (5,5)
 *
 * Input:
 *     int numero : Numero del jugador
 *     int rol : Rol del jugador
 *     int x : Coordenada x del jugador en el laberinto
 *     int y : Coordenada y del jugador en el laberinto
 *
 * Return:
 *     t_Jugador* : Puntero al jugador creado
*/
t_Jugador* crearJugador(int numero, int rol, int x, int y);

/**
 * void liberarJugador
 *
 * Libera la memoria asociada a un jugador.
 *
 * Input:
 *     t_Jugador* jugador : Puntero al jugador a liberar
*/
void liberarJugador(t_Jugador* jugador);

/**
 * void imprimirJugador
 *
 * Imprime en consola la informacion de un jugador.
 *
 * Input:
 *     t_Jugador* jugador : Puntero al jugador a imprimir
*/
void imprimirJugador(t_Jugador* jugador);


#endif
