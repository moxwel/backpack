#include <stdio.h>

float promedio(int *arr, int n) {

    int suma = 0;

    for (int i = 0; i < n; i++) {
        suma += arr[i];
    }

    return (float)suma / n;
}


int main() {
    int A[10] = {54,3,21,5,7,2,5,7,9,3};

    printf("El promedio de los valores es: %f\n", promedio(A,10));

    return 0;
}
