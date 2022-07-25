#include <stdio.h>
#include <stdlib.h>

// La funcion retorna un arreglo de enteros
int* filtro_menor(int *arr, int n, int x, int *o) {

    // Asi el arreglo creado no se destruye
    int *B = (int*)malloc(sizeof(int) * n);

    int i = 0, j = 0;
    for (i = 0; i < n; i++) {
        if (arr[i] < x) {
            B[j] = arr[i];
            j++;
        }
    }

    *o = j;
    return B;
}


int main() {

    int A[10] = {1,2,3,4,5,6,7,8,9,10};

    int elem;
    int *C = filtro_menor(A, 10, 5, &elem);

    printf("Indice:\tValor:\n");
    for (int i = 0; i < elem; i++) {
        printf("%i\t%i\n", i, C[i]);
    }

    free((void*)C);

    return 0;
}
