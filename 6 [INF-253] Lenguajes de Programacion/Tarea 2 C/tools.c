#include "tools.h"

void printArrayInt(int* array) {
    int arraySize = array[0];

    printf("Size: %d = [ ", arraySize);
    for (int i = 1; i < arraySize; i++) {
        printf("%d, ", array[i]);
    }
    printf("%d ] = Avrg: %d\n", array[arraySize], averageArrayInt(array));
}

void printArrayChar(char* array) {
    int arraySize = array[0];

    printf("Size: %d = [ ", arraySize);
    for (int i = 1; i < arraySize; i++) {
        printf("%c, ", array[i]);
    }
    printf("%c ] = Avrg: %d\n", array[arraySize], averageArrayChar(array));
}

int averageArrayInt(int* array){
    int arraySize = array[0];
    int suma = 0;

    for (int i = 1; i < arraySize + 1; i++) {
        suma += array[i];
        // printf("Sumando: %d - Total: %d\n", array[i], suma);
    }

    return techo(suma, arraySize);
}

int averageArrayChar(char* array){
    int arraySize = array[0];
    int suma = 0;

    for (int i = 1; i < arraySize + 1; i++) {
        suma += (int)array[i];
        // printf("Sumando: %-3d (%c) - Total: %d\n", (int)array[i], array[i], suma);
    }

    return techo(suma, arraySize);
}

int techo(int a, int b){
    if ((b == 0) || (a == 0)) {
        return 0;
    }

    if (b == 1) {
        return a;
    }

    if (a <= b) {
        return 1;
    }

    return (a / b) + ((a % b) != 0);
}
