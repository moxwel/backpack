#ifndef __LABERINTO__
#define __LABERINTO__

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <dirent.h>
#include <sys/types.h>
#include <sys/stat.h>

/**
 * Tipos de celda:
 * 0: Nada
 * 1: Pared
 * 2: Puerta
 * 3: Tesoro
 * 4: Pasaje
*/
typedef struct {
    int tipo;
    int puertaEstaAbierta;
    int tesoroNumero;
    int pasajeDireccion;
} t_Celda;

/**
 * Tipos de salidas:
 * 0: N O S E
 * 1: N O E
 * 2: N O S
 * 3: O S E
 * 4: N S E
 * 5: N O
 * 6: O S
 * 7: S E
 * 8: N E
 *
 * El laberinto es de 5x5
*/
typedef struct Laberinto {
    t_Celda** mtrx_laberinto;
    int esInicial;

    int tipoSalidas;
    struct Laberinto* Nor;
    struct Laberinto* Sur;
    struct Laberinto* Est;
    struct Laberinto* Oes;

    int estaUsado;
} t_Laberinto;

/**
 * t_Laberinto* crearLaberinto
 *
 * Recibe la direccion de un archivo que contiene la informacion de un laberinto
 * y retorna un puntero a un elemento t_Laberinto con la informacion del laberinto.
 *
 * Input:
 *     - char* nombre_archivo : Direccion del archivo que contiene la informacion del laberinto.
 *
 * Return:
 *     - t_Laberinto* : Puntero a un elemento t_Laberinto con la informacion del laberinto.
*/
t_Laberinto* crearLaberinto(char* nombre_archivo);

/**
 * void liberarLaberinto
 *
 * Recibe un puntero a un elemento t_Laberinto y libera la memoria asignada a este.
 *
 * Input:
 *     - t_Laberinto* laberinto : Puntero a un elemento t_Laberinto.
*/
void liberarLaberinto(t_Laberinto* laberinto);

/**
 * void imprimirLaberinto
 *
 * Recibe un puntero a un elemento t_Laberinto e imprime en consola la informacion del laberinto.
 *
 * Input:
 *     - t_Laberinto* laberinto : Puntero a un elemento t_Laberinto.
*/
void imprimirLaberinto(t_Laberinto* laberinto);

#endif
