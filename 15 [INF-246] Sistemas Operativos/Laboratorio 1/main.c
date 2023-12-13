#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>
#include <ctype.h>

#include "sopa.h"
#include "carpetas.h"

int main() {

    // Obtener lista de archivos
    int n_archivos = 0;
    char** ruta_archivos = arregloArchivos("./archivos", &n_archivos);
    printf("Numero de archivos: %d\n", n_archivos);

    crearCarpetas();

    for (int i = 0; i < n_archivos; i++) {
        printf("\nlista_archivos[%d]:\n",i);
        printf("- Ruta archivo: %s\n", ruta_archivos[i]);

        // Extraer palabra a buscar desde la direccion del archivo
        char palabra[100];
        extaerNombreArchivo(ruta_archivos[i], palabra);
        int palabra_largo = strlen(palabra);

        // Transformar palabra a mayuscula
        for (int i = 0; i < palabra_largo; i++) {
            palabra[i] = toupper(palabra[i]);
        }
        printf("- Palabra a buscar: %s\n", palabra);

        // Obtener orientacion de la sopa de letras
        int es_horizontal = orientacionSopa(ruta_archivos[i]);
        if (es_horizontal == -1) {
            printf("Error al leer el archivo\n");
            return 1;
        }
        printf("- Es horizontal (1=si, 0=no): %d \n", es_horizontal);

        // Guardar sopa de letras en una matriz
        int sopa_dimension = 0;
        char** sopa_letras = generarMatrizSopa(ruta_archivos[i], &sopa_dimension);
        if(sopa_letras == NULL){
            printf("Error al leer el archivo\n");
            return 1;
        }
        printf("- Dimension: %d\n", sopa_dimension);

        // Buscar palabra
        clock_t inicio = clock();
        int resultado = buscarPalabraSopaNoPrint(sopa_letras, sopa_dimension, palabra, es_horizontal);
        clock_t fin = clock();

        clock_t diff = fin - inicio;
        double tiempo = (double)diff / CLOCKS_PER_SEC;

        printf("- Ticks: %ld [ticks]\n", diff);
        printf("- Tiempo: %f [s]\n", tiempo);

        if (resultado) {
            printf("- Palabra encontrada\n");
        } else {
            printf("- Palabra no encontrada\n");
        }

        liberarMatrizSopa(sopa_letras, sopa_dimension);
        moverArchivo(ruta_archivos[i], es_horizontal, sopa_dimension);
    }

    liberarArregloArchivos(ruta_archivos, n_archivos);

    return 0;
}
