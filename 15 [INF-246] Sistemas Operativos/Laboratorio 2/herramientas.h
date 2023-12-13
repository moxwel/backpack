#ifndef __HERRAMIENTAS__
#define __HERRAMIENTAS__

#include <stdlib.h>
#include <stdio.h>
#include <unistd.h>

#include "mapa.h"
#include "jugadores.h"

/**
 * int chequearCelda
 *
 * Retorna el tipo de celda encontrada en las coordenadas entregadas.
 *
 * Input:
 *     t_Mapa* mapa : Puntero al mapa
 *     int map_x : Coordenada x en el mapa
 *     int map_y : Coordenada y en el mapa
 *     int lab_x : Coordenada x en el laberinto
 *     int lab_y : Coordenada y en el laberinto
 *
 * Return:
 *     int : Tipo de celda encontrada
*/
int chequearCelda(t_Mapa* mapa, int map_x, int map_y, int lab_x, int lab_y);

/**
 * int puertaEstaAbierta
 *
 * Retorna 1 si la puerta en las coordenadas entregadas esta abierta, 0 en caso contrario.
 *
 * Input:
 *     t_Mapa* mapa : Puntero al mapa
 *     int map_x : Coordenada x en el mapa
 *     int map_y : Coordenada y en el mapa
 *     int lab_x : Coordenada x en el laberinto
 *     int lab_y : Coordenada y en el laberinto
 *
 * Return:
 *     int : 1 si la puerta esta abierta, 0 en caso contrario.
*/
int puertaEstaAbierta(t_Mapa* mapa, int map_x, int map_y, int lab_x, int lab_y);

/**
 * int jugadorTieneLlave
 *
 * Retorna 1 si el jugador tiene la llave, 0 en caso contrario.
 *
 * Input:
 *     t_Jugador* jugador : Puntero al jugador
 *
 * Return:
 *     int : 1 si el jugador tiene la llave, 0 en caso contrario.
*/
int jugadorTieneLlave(t_Jugador* jugador);

/**
 * int abrirPuerta
 *
 * Abre la puerta en las coordenadas entregadas.
 *
 * Input:
 *     t_Mapa* mapa : Puntero al mapa
 *     int map_x : Coordenada x en el mapa
 *     int map_y : Coordenada y en el mapa
 *     int lab_x : Coordenada x en el laberinto
 *     int lab_y : Coordenada y en el laberinto
 *
 * Return:
 *     int : 1 si se pudo abrir la puerta, 0 en caso contrario.
*/
int abrirPuerta(t_Mapa* mapa, int map_x, int map_y, int lab_x, int lab_y);

/**
 * int direccionPasaje
 *
 * Retorna la direccion del pasaje en las coordenadas entregadas.
 *
 * Input:
 *     t_Mapa* mapa : Puntero al mapa
 *     int map_x : Coordenada x en el mapa
 *     int map_y : Coordenada y en el mapa
 *     int lab_x : Coordenada x en el laberinto
 *     int lab_y : Coordenada y en el laberinto
 *
 * Return:
 *     int : Direccion del pasaje
*/
int direccionPasaje(t_Mapa* mapa, int map_x, int map_y, int lab_x, int lab_y);

/**
 * int coincideTesoro
 *
 * Retorna 1 si el tesoro en las coordenadas entregadas coincide con el numero
 * del jugador, 0 en caso contrario.
 *
 * Input:
 *     t_Mapa* mapa : Puntero al mapa
 *     int map_x : Coordenada x en el mapa
 *     int map_y : Coordenada y en el mapa
 *     int lab_x : Coordenada x en el laberinto
 *     int lab_y : Coordenada y en el laberinto
 *     t_Jugador* jugador : Puntero al jugador
 *
 * Return:
 *     int : 1 si el tesoro coincide, 0 en caso contrario.
*/
int coinicideTesoro(t_Mapa* mapa, int map_x, int map_y, int lab_x, int lab_y, t_Jugador* jugador);

/**
 * void tomarTesoro
 *
 * Hace que el jugador tome el tesoro en las coordenadas entregadas.
 *
 * Input:
 *     t_Mapa* mapa : Puntero al mapa
 *     int map_x : Coordenada x en el mapa
 *     int map_y : Coordenada y en el mapa
 *     int lab_x : Coordenada x en el laberinto
 *     int lab_y : Coordenada y en el laberinto
 *     t_Jugador* jugador : Puntero al jugador
*/
void tomarTesoro(t_Mapa* mapa, int map_x, int map_y, int lab_x, int lab_y, t_Jugador* jugador);

/**
 * int jugadorTieneTesoro
 *
 * Retorna 1 si el jugador tiene el tesoro, 0 en caso contrario.
 *
 * Input:
 *     t_Jugador* jugador : Puntero al jugador
 *
 * Return:
 *     int : 1 si el jugador tiene el tesoro, 0 en caso contrario.
*/
int jugadorTieneTesoro(t_Jugador* jugador);

/**
 * int avanzarPasaje
 *
 * Avanza al jugador por el pasaje en las coordenadas entregadas.
 *
 * Input:
 *     t_Mapa* mapa : Puntero al mapa
 *     int map_x : Coordenada inicial x en el mapa
 *     int map_y : Coordenada inicial y en el mapa
 *     int lab_x : Coordenada inicial x en el laberinto
 *     int lab_y : Coordenada inicial y en el laberinto
 *     t_Jugador* jugador : Puntero al jugador
 *     t_Laberinto** mazo_laberintos : Mazo de laberintos
 *     int n_archivos : Cantidad de laberintos en el mazo
 *
 * Return:
 *     int : 1 si se pudo avanzar, 0 en caso contrario.
*/
int avanzarPasaje(t_Mapa* mapa, int map_x, int map_y, int lab_x, int lab_y, t_Jugador* jugador, t_Laberinto** mazo_laberintos, int n_archivos);

/**
 * int moverJugador
 *
 * Mueve al jugador en la direccion indicada.
 *
 * Input:
 *     t_Mapa* mapa : Puntero al mapa
 *     int map_x : Coordenada inicial x en el mapa
 *     int map_y : Coordenada inicial y en el mapa
 *     int lab_x : Coordenada inicial x en el laberinto
 *     int lab_y : Coordenada inicial y en el laberinto
 *     t_Jugador* jugador : Puntero al jugador
 *     int direccion : Direccion en la que se movera el jugador
 *
 * Return:
 *     int : 1 si se pudo mover, 0 en caso contrario.
*/
int moverJugador(t_Mapa* mapa, int map_x, int map_y, int lab_x, int lab_y, t_Jugador* jugador, int direccion);

#endif
