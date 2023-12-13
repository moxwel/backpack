#ifndef __MAPA__
#define __MAPA__

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <dirent.h>
#include <sys/types.h>
#include <sys/stat.h>

#include "laberinto.h"

/**
 * Un mapa es una matriz de laberintos.
 *
 * En este caso, es una matriz de punteros a laberintos, ya que estos ultimos
 * estan creados en memoria con la funcion crearLaberinto() de la libreia 'laberinto.h'.
*/
typedef struct {
    t_Laberinto*** mtrx_mapa;
} t_Mapa;

/**
 * t_Mapa* crearMapa
 *
 * Crea un mapa usando los laberintos del mazo de laberintos.
 * Inicialmente, el mapa solo tiene el laberinto inicial en el centro.
 *
 * Input:
 *     t_Laberinto** mazo_laberintos : Mazo de laberintos
 *     int n_laberintos : Cantidad de laberintos en el mazo
 *
 * Return:
 *     t_Mapa* : Puntero al mapa creado en memoria.
*/
t_Mapa* crearMapa(t_Laberinto** mazo_laberintos, int n_laberintos);

/**
 * void imprimirMapa
 *
 * Imprime el mapa en consola.
 *
 * Input:
 *     t_Mapa* mapa : Puntero al mapa a imprimir.
 *
 * Return:
 *     void : Esta funcion no retorna nada.
*/
void imprimirMapa(t_Mapa* mapa);

/**
 * void liberarMapa
 *
 * Libera la memoria del mapa.
 *
 * Input:
 *     t_Mapa* mapa : Puntero al mapa a liberar.
 *
 * Return:
 *     void : Esta funcion no retorna nada.
*/
void liberarMapa(t_Mapa* mapa);

/**
 * int unirLaberinto
 *
 * Une un laberinto al mapa en la coordenada (x, y) en la direccion indicada.
 * La coordenada (x, y) debe tener un pasaje en la direccion indicada.
 *
 * Input:
 *     t_Mapa* mapa : Puntero al mapa donde se unira el laberinto.
 *     t_Laberinto** mazo_laberintos : Mazo de laberintos.
 *     int n_laberintos : Cantidad de laberintos en el mazo.
 *     int x : Coordenada x del mapa donde se unira el laberinto.
 *     int y : Coordenada y del mapa donde se unira el laberinto.
 *     int direccion : Direccion en la que se unira el laberinto.
 *
 * Return:
 *     int : 0 si se pudo unir el laberinto, 1 si no se pudo.
*/
int unirLaberinto(t_Mapa* mapa, t_Laberinto** mazo_laberintos, int n_laberintos, int x, int y, int direccion);

#endif
