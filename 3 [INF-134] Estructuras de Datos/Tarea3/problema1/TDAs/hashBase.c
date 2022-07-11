#define VACIA -1

typedef int tipoClave;

/**
 * NOTAS:
 *
 * - Para aumentar la "aleatoriedad", usamos el tamaño de la tabla
 *   de hashing para generar los indices. Por esa razon se tiene
 *   un parametro adicional 'M'.
 *
 **/

int h(int k, int M) {
    return k % M;
}
int h2(int k, int M){
    return 3 - k % 3;
}
int p(int k, int i, int M){
    // Doble hashing
    return ((i * h2(k,M)) + h2(k,M)) % M;
}

#include "./hashProductos.c"
#include "./hashOfertas.c"
#include "./hashVenta.c"
