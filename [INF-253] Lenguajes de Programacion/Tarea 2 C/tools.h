#ifndef __TOOLS_C__
#define __TOOLS_C__

// ====================

#include <stdio.h>

/**
 * Imprimir arreglo de ints.
 *
 * NOTA: El primer elemento del arreglo es el tamano de este
 **/
void printArrayInt(int* array);

/**
 * Imprimir arreglo de char.
 *
 * NOTA: El primer elemento del arreglo es el tamano de este
 **/
void printArrayChar(char* array);

/**
 * Obtiene el promedio de los elementos de un arreglo de ints.
 *
 * NOTA: El primer elemento del arreglo es el tamano de este.
 * No se toma en consideracion para el calculo del promedio.
 **/
int averageArrayInt(int* array);

/**
 * Obtiene el promedio de los elementos de un arreglo de char.
 *
 * NOTA: El primer elemento del arreglo es el tamano de este.
 * No se toma en consideracion para el calculo del promedio.
 **/
int averageArrayChar(char* array);

/**
 * Divide dos numeros, redondeando hacia arriba.
 **/
int techo(int a, int b);

#endif
