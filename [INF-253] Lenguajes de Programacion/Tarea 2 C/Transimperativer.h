#ifndef __TRANSIMPERATIVER__
#define __TRANSIMPERATIVER__

// ====================

#include "listaEnlazada.h"

/**
 * Tipo de dato: Transimperativer.
 *
 * Celdas:
 * - mezclar : Funcion de mezcla asociado
 * - separar : Funcion de separacion asociado
 * - cuerpo : Lista de caracteristicas asociado
 **/
typedef struct Transimp{
    void (*mezclar)(struct Transimp*, struct Transimp*);
    void (*separar)(struct Transimp*, struct Transimp*, struct Transimp*);
    Cuerpo* cuerpo;
} Transimperativer;

// Operaciones de inicializacion

/**
 * Crea un Transimperativer nuevo.
 * Debe asignarse a una variable Transipmerativer*.
 *
 * Ademas, tambien debe recibir punteros a las funciones
 * 'mezclar' y 'separar' que se le van a asociar al Transimperativer creado.
 **/
Transimperativer* CrearTransimperativer(
                           void (*mezclarF)(Transimperativer* T1, Transimperativer* T2),
                           void (*separarF)(Transimperativer* padre, Transimperativer* hijo1, Transimperativer* hijo2));

/**
 * Elimina el Transimperativer 'T' dado.
 *
 * Libera memoria.
 **/
void BorrarTransimperativer(Transimperativer* T);

// Operaciones de atributos

/**
 * Agrega una caracteristica al Transimperativer 'T' dado.
 *
 * Se debe especificar el tipo de la caracteristica a agregar.
 * - 'i' (105) : int
 * - 'c' (99) : char
 * - 'h' (104) : char* (arreglo de char)
 * - 'a' (97) : int* (arreglo de int)
 *
 * NOTA: Se agrega al final de la lista de caracteristicas.
 **/
void AnadirAtributo(Transimperativer* T, char tipo, void* caracteristica);

/**
 * Remueve una caracteristica del Transimperativer 'T' dado.
 *
 * Se debe especificar el indice donde esta ubicado la caracteristica en la lista.
 **/
void QuitarAtributo(Transimperativer* T, int indice);

// Operaciones de Transimperativer

/**
 * Mezcla dos Transimperativer.
 *
 * Se intercambian caracteristicas segun como diga la funcion 'mezclar'
 * del Transimperativer 'T1'
 **/
void MezclarTransimperativer(Transimperativer* T1, Transimperativer* T2);

/**
 * Separa un Transimperativer 'padre' dado segun su funcion 'separar' asociado,
 * en dos nuevos Transimperativer 'hijo1' e 'hijo2'.
 *
 * NOTA: Despues de separar, el Transimperativer 'padre' deja de existir.
 **/
void SepararTransimperativer(Transimperativer* padre, Transimperativer* hijo1, Transimperativer* hijo2);

/**
 * Se hace pelear a dos Transimperativer 'T1' y 'T2'.
 *
 * Segun las caracteristicas de cada uno, se van sumando puntos.
 * Quien mas puntos tenga, es el ganador.
 **/
void Pelear(Transimperativer* T1, Transimperativer* T2);

/**
 * Reemplaza el cuerpo de un Transimperativer 'T' dado por uno nuevo 'L'.
 *
 * El cuerpo antiguo sera eliminado de la memoria.
 **/
void ReemplazarCuerpo(Transimperativer* T, Cuerpo* L);

// Operaciones de obtencion

/**
 * Obitne el tamano de la lista de caracteristicas del Transimperativer 'T' dado.
 **/
int TamanoTransimperativer(Transimperativer* T);

/**
 * Muestra las caracteristicas del Transimperativer 'T' dado.
 **/
void MostrarTransimperativer(Transimperativer* T);

/**
 * Quita un Transimperativer de un arreglo.
 **/
void removerTransimpLista(Transimperativer* array[], int indice, int size);

#endif
