#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "./ABB_TDA.c"
#include "./sucesor.c"
#include "./fileGen.c"

/**
 * ./main <1>
 ******
 * Argumentos de programa:
 *     <1> : Archivo que contiene las operaciones
 ******
 * Ejemplo:
 *     ./main input1.dat
 **/
int main(int argc, char **argv) {

    // Error holder
    if (argc != 2) {
        printf("No estan los argumentos requeridos.\n\nUso: ./main <1>\n\t<1> : Archivo que contiene las operaciones\n\nEjemplo: ./main input.txt\n");
        return 1;
    }


    tABB mi_arbol;
    tABB *pa = &mi_arbol;

    initABB(pa);

    FILE *fp = fopen(argv[1], "r");
    if (fp == NULL) {
        printf("Error. No se encuentra el archivo.\n");
        exit(1);
    }
    FILE *fo = fopen("output.txt","w");

    int u;
    fscanf(fp,"%i",&u);

    char *act = (char*)malloc(sizeof(char)*10);
    int elem;

    while (fscanf(fp,"%s %i", act, &elem) > 0) {
        if (strcmp(act,"PREORDEN") == 0){
            //printf(">>> imprimir preorden\n");
            preOrdenFile(pa,fo);
        }
        if (strcmp(act,"SUCESOR") == 0){
            //printf("->-> sucesor del numero %i\n",elem);
            sucesorFile(pa,elem,u,fo);
        }
        if (strcmp(act,"INSERTAR") == 0){
            //printf("+++ ingresar numero %i\n",elem);
            insert(pa,elem);
        }
        if (strcmp(act,"BORRAR") == 0){
            //printf("--- borrar numero %i\n",elem);
            removeNode(pa,elem);
        }
    }

    fclose(fp);
    fclose(fo);
    clear(pa);
    return 0;
}
