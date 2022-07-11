// Programa principal

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "./tools.c"
#include "./fileToArr.c"
#include "./arrToFile.c"
#include "./buscar_str.c"

/**
 * ./main <1> <2>
 ******
 * Argumentos de programa:
 *     1 : Archivo que contiene las palabras
 *     2 : Archivo que contiene los parrijos
 ******
 * Ejemplo:
 *     ./main S.txt P.txt
 **/
int main(int argc, char **argv) {

    // printArrStr(argv,argc);
    // printf("cantidad de arg: %i\n",argc);

    char *sfile;
    char *pfile;

    // Esto es para usarse con argumentos de programa.
    // Si no se ingresa ningun argumento, se usaran los predeterminados.
    if (argc == 1) {
        // printf("Utilizando parametros predeterminados...\n");
        sfile = "S.txt";
        pfile = "P.txt";
    } else if (argc == 2) {
        printf("- Falta 1 parametro mas -\n");
        exit(1);
    } else {
        // printf("Utilizando parametros personalizados...\n");
        sfile = argv[1];
        pfile = argv[2];
    }

    int n_sarr = 0;
    char **sarr = fileToArr(sfile, &n_sarr);

    int n_parr = 0;
    char **parr = fileToArr(pfile, &n_parr);

    int n_out = 0;
    for (int i = 0; i < n_parr; i++) {
        char **outarr = buscar_str(sarr, n_sarr, parr[i], &n_out);
        arrToFile(outarr,n_out,parr[i]);
        free(outarr);
    }

    free(sarr);
    free(parr);
    return 0;
}
