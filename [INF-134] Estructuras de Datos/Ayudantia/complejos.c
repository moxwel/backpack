#include <stdio.h>
#include <stdlib.h>
#include <math.h>

// Creacion de TDA "complejo"
typedef struct {
    int re;
    int im;
} complejo;

// ==================================
//     Creacion de interfaz de TDA
// Estos si pueden acceder a los datos dentro
// de una TDA.

// Constructor: asignar memoria
// Retorna la direcicon donde esta ubicada el complejo
// (recordar usar malloc ya que no queremos que
// se borre cuando termine la funcion)
complejo* crearComplejo(int real, int imag) {
    complejo *n = (complejo*)malloc(sizeof(complejo));
    // Hay que decirle a malloc a que tipo
    // de dato esta apuntando la memoria que se designa

    n->re = real;
    n->im = imag;
    return n;
}

// Destructor: liberar memoria
void borrarComplejo(complejo *n) {
    free(n);
}

// Valor absoluto de un numero complejo
// Pitagoras (uso de math.h)
double valorAbsoluto(complejo *n) {
    return sqrt(pow(n->re, 2) + pow(n->im, 2));
}

// Conjugado de complejo
// Se crea un nuevo complejo, pero con la parte
// imaginaria con el signo cambiado
complejo* conjugado(complejo *n) {
    complejo *cnjg = crearComplejo(n->re, (n->im)*-1);
    return cnjg;
}

// Suma de complejos
// Se suma la parte real con la parte imaginaria
// de cada numero complejo
complejo* sumaCompl(complejo *n1, complejo *n2) {
    complejo *suma = crearComplejo((n1->re + n2->re), (n1->im + n2->im));
    return suma;
}

int parteRe(complejo *n) {
    return n->re;
}
int parteIm(complejo *n) {
    return n->im;
}
void printCompl(complejo *n) {
    printf("%i + %ii\n",parteRe(n),parteIm(n));
}

// ==============================


int main() {

    complejo *c1 = crearComplejo(4,4);
    complejo *c2 = crearComplejo(8,8);
    complejo *c3 = crearComplejo(2,-3);

    printCompl(sumaCompl(c2,c3));

    borrarComplejo(c1);
    borrarComplejo(c2);
    borrarComplejo(c3);

    return 0;
}
