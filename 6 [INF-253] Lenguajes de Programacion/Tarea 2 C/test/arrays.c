#include <stdio.h>


void eliminarLista(int array[], int indice, int size) {

    for (int j = indice; j < size; j++) {
        array[j] = array[j+1];
    }

}

int main() {
    int numeros[5] = {23, 45, 67, 32, 12};
    int tamano = 5;


    for (int i = 0; i < tamano; i++) {
        printf("[%d] %d\n", i, numeros[i]);
    }


    eliminarLista(numeros, 2, tamano);
    tamano--;

    for (int i = 0; i < tamano; i++) {
        printf("[%d] %d\n", i, numeros[i]);
    }

    return 0;
}
