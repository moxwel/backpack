/*
    Implementacion de lista enlazada (adaptada)
    https://github.com/moxwel/basics/tree/master/C/11-edd/5-reimplementation
*/

#ifndef __LINKED_LIST__
#define __LINKED_LIST__

// ====================

#include <stdio.h>
#include <stdlib.h>
#include "tools.h"

/**
 * Tipo de dato: Elementos de la lista.
 *
 * Definido como: void*
 **/
typedef void* tListElem;

/**
 * Tipo de dato: Nodo de lista.
 *
 * Celdas:
 * - info : Elemento a guardar (caracteristica)
 * - tipo : Define el tipo de dato de 'info'
 * ---- 'i' (105) : int
 * ---- 'c' (99) : char
 * ---- 'h' (104) : char* (arreglo de char)
 * ---- 'a' (97) : int* (arreglo de int)
 * - sig : Puntero al nodo siguiente
 **/
typedef struct ListNode{
    tListElem info; // Caracteristica
    char tipo;
    struct ListNode *sig;
} nodo;

/**
 * Tipo de dato: EDD Lista.
 *
 * Celdas:
 * - head : Puntero al primer nodo
 * - tail : Puntero al ultimo nodo
 * - curr : Puntero al nodo actual
 * - pos : Posicion del cursor
 * - len : tamano de la lista
 **/
typedef struct {
    nodo* head;
    nodo* tail;
    nodo* curr;
    int pos;
    int len;
} Cuerpo;

/**
 * Union: Dato de caracteristica
 *
 * Puede usar cualquiera de los siguientes tipos:
 * - 'i' (105) : int
 * - 'c' (99) : char
 * - 'h' (104) : char* (arreglo de char)
 * - 'a' (97) : int* (arreglo de int)
 **/
typedef union {
    int vacio;
    int i;
    char c;
    int* a;
    char* h;
} dataInfo;

// Operaciones de inicializacion

/**
 * Crear lista nueva.
 * Debe asignarse a una variable Cuerpo*.
 **/
Cuerpo* newList();

/**
 * Eliminar lista.
 * Libera toda la memoria.
 **/
void deleteList(Cuerpo* L);

// Operaciones de elemento

/**
 * Insertar elemento en lista.
 * Se va a insertar despues del cursor.
 *
 * Se debe especificar el tipo de dato a insertar:
 * - 'i' (105) : int
 * - 'c' (99) : char
 * - 'h' (104) : char* (arreglo de char)
 * - 'a' (97) : int* (arreglo de int)
 **/
void insert(Cuerpo* L, char tipoItem, tListElem item);

/**
 * Adjuntar elemento en la lista.
 * Se va a insertar al final de la lista.
 *
 * Se debe especificar el tipo de dato a adjuntar:
 * - 'i' (105) : int
 * - 'c' (99) : char
 * - 'h' (104) : char* (arreglo de char)
 * - 'a' (97) : int* (arreglo de int)
 **/
void append(Cuerpo* L, char tipoItem, tListElem item);

/**
 * Eliminar elemento de la lista.
 * Se va a eliminar el elemento despues del cursor.
 *
 * NOTA: Tambien libera la memoria del puntero de 'info'
 *
 * Retornos:
 * - 1 : Exito
 * - 0 : Sin exito (lista vacia)
 **/
int erase(Cuerpo* L);

/**
 * Eliminar elemento de la lista.
 * Se va a eliminar el elemento despues del cursor.
 *
 * NOTA: Esta funcion no libera la memoria del puntero de 'info'
 *
 * Retornos:
 * - 1 : Exito
 * - 0 : Sin exito (lista vacia)
 **/
int removeNode(Cuerpo* L);

/**
 * Limpiar lista.
 * Todos los elementos se borran.
 **/
void clearList(Cuerpo* L);

// Operaciones de cursor

/**
 * Mover cursor hacia el siguiente elemento.
 *
 * Retornos:
 * - 1 : Exito
 * - 0 : Sin exito (final de la lista)
 **/
int next(Cuerpo* L);

/**
 * Mover cursor al elemento anterior.
 *
 * Retornos:
 * - 1 : Exito
 * - 0 : Sin exito (comienzo de la lista)
 **/
int prev(Cuerpo* L);

/**
 * Mover cursor a posicion determinada.
 *
 * Retornos:
 * - 1 : Exito
 * - 0 : Sin exito (invalido)
 **/
int moveToPos(Cuerpo* L, int target);

/**
 * Mover cursor al inicio de la lista.
 **/
void moveToStart(Cuerpo* L);

/**
 * Mover cursor al final de la lista.
 **/
void moveToEnd(Cuerpo* L);

// Operaciones de obtencion

/**
 * Obtiene el valor del elemento siguiente al cursor.
 * Si el cursor esta al final, se obtiene el valor actual.
 *
 * Retorno:
 * Una union 'dataInfo'.
 * Se debe acceder a la celda dependiendo del tipo con el operador punto.
 * - 'i' (105) : int
 * - 'c' (99) : char
 * - 'h' (104) : char* (arreglo de char)
 * - 'a' (97) : int* (arreglo de int)
 **/
dataInfo getValue(Cuerpo* L);

/**
 * Retorna el puntero 'info' del elemento siguiente al cursor.
 * Si el cursor esta al final, se obtiene el valor actual.
 **/
void* getInfoPointer(Cuerpo* L);

/**
 * Imprime el valor del elemento siguiente al cursor.
 * Si el cursor esta al final, se obtiene el valor actual.
 **/
void printValue(Cuerpo* L);

/**
 * Obitne el tipo del elemento siguiente al cursor.
 * Si el cursor esta al final, se obtiene el valor actual.
 *
 * Retorno:
 * Un char.
 * - 'i' (105) : int
 * - 'c' (99) : char
 * - 'h' (104) : char* (arreglo de char)
 * - 'a' (97) : int* (arreglo de int)
 **/
char getType(Cuerpo* L);

/**
 * Obtiene el valor del tamano de la lista.
 **/
int getLength(Cuerpo* L);

/**
 * Obtiene el valor de la posicion del cursor.
 **/
int getPos(Cuerpo* L);



/**
 * Imprimir lista.
 **/
void printList(Cuerpo* L);

// ====================

#endif
