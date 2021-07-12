#include <stdio.h>
#include <math.h>


float tempconv(float t, char i, char o) {

    if (i != 'C' && i != 'F') {
        printf("Error: Unidad no valida.\n");
        return 1;
    }
    if (o != 'C' && o != 'F') {
        printf("Error: Unidad no valida.\n");
        return 1;
    }
    if (i == o) {
        return t;
    }
    if (i == 'C') {
        return t*(9.0/5.0) + 32.0;
    }
    if (i == 'F') {
        return ((t - 32.0)*5.0)/9.0;
    }
}


int main() {
    float res, temp;
    char inp, out;

    temp = 1003;
    inp = 'C';
    out = 'F';
    res = tempconv(temp, inp, out);

    printf("Resultado de %c a %c: %f\n", inp, out, res);

    return 0;
}
