#ifndef __ARCHIVOS__
#define __ARCHIVOS__

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <dirent.h>
#include <sys/types.h>
#include <sys/stat.h>

/**
 * char** arregloArchivos
 *
 * Recibe la direccion de un directorio y retorna un arreglo con las direcciones
 * de todos los archivos que se encuentran en el directorio.
 *
 * Input:
 *     - char* directorio : Direccion del directorio.
 *     - int* n_archivos : Puntero a entero donde se guardará el numero de archivos
 *
 * Return:
 *     - char** : Puntero a un arreglo de strings con las direcciones de los archivos
*/
char** arregloArchivos(char* directorio, int* n_archivos);

/**
 * void liberarArregloArchivos
 *
 * Recibe un arreglo de strings y libera la memoria asignada a cada string y al arreglo.
 *
 * Input:
 *     - char** arreglo : Arreglo de strings a liberar.
 *     - int n_archivos : Numero de archivos en el arreglo.
*/
void liberarArregloArchivos(char** arreglo, int n_archivos);

#endif
