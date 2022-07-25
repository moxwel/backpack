#include <stdio.h>

typedef struct {
    int numerador;
    int denominador;
} racional;

racional crearRacional(int num, int den) {
    racional r;
    r.numerador = num;
    r.denominador = den;
    return r;
}
int numerador(racional *r) {
    return r->numerador;
}
int denominador(racional *r) {
    return r->denominador;
}

int main() {
    racional r1, r2, r3;
    r1 = crearRacional(8, 33);
    r2 = crearRacional(9, 13);
    r3 = crearRacional(numerador(&r1)*numerador(&r2), denominador(&r1)*denominador(&r2));

    printf("El número resultante es %d/%d\n",numerador(&r3),denominador(&r3));

    return 0;
}
