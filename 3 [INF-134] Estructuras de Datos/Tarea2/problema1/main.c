#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "./linked_list_TDA.c"
#include "./memTools.c"

/**
 * ./main <1>
 ******
 * Argumentos de programa:
 *     1 : Archivo que contiene las operaciones
 ******
 * Ejemplo:
 *     ./main input1.dat
 **/
int main(int argc, char **argv) {

    if (argc != 2) {
        printf("No estan los argumentos requeridos.\n\nUso: ./main <1>\n\t<1> : Archivo que contiene las operaciones\n\nEjemplo: ./main input1.dat\n");
        return 1;
    }

    FILE *fp = fopen(argv[1], "r");
    if (fp == NULL) {
        printf("Error. No se encuentra el archivo.\n");
        exit(1);
    }

    FILE *fo = fopen("output1.dat","w");

    int maxMem, nOps;
    fscanf(fp,"%i",&maxMem);
    fscanf(fp,"%i",&nOps);
    //printf("Memoria maxima: %i\nOperaciones a realizar: %i\n\n",maxMem,nOps);

    tLista libres, ocupados;
    tLista *pl = &libres;
    tLista *po = &ocupados;
    initListMem(pl, maxMem);
    initList(po);

    //printList(pl);
    //printList(po);

    char *operacion = (char*)malloc(sizeof(char)*8);
    int b;

    // Hacer cada operacion
    for (int i = 0; i < nOps; i++) {
        fscanf(fp,"%s %i",operacion,&b);
        //printf("::::::::::::::::::OPERACION %i:::::::::::::::::::::\n\n",i+1);

        if (strcmp(operacion,"malloc") == 0) {

            //printf(">>> Alocando %i bytes\n",b);
            //printf("Posicion en lista: %i\n",memInsert(b, pl));

            if (memInsert(b,pl) != 0) {
                memMalloc(pl,po,memInsert(b,pl),b);

                fprintf(fo, "Bloque de %i asignado a partir del byte %i\n",b,po->tail->bStrt);
            } else {
                //printf("No se puede alocar esa memoria\n");

                fprintf(fo, "Bloque de %i bytes NO puede ser asignado\n",b);
            }

        } else if (strcmp(operacion,"free") == 0) {

            //printf("--- Liberando bloque que comienza en el byte %i\n",b);
            //printf("Posicion en lista: %i\n", memRemove(b,po));

            if (memRemove(b,po) != 0) {
                memFree(pl,po,memRemove(b,po),b);

                fprintf(fo, "El bloque de %i bytes liberado\n",bloqSize(pl));

                //printList(pl);
                //printList(po);

                // Despues de liberar memoria ocupada, hay que revisar si esa memoria
                // son contiguas.

                //printf("Revisar memoria contigua...\n");
                //printf("Final del bloque cursor: %i - Comienzo del bloque siguiente: %i\n",getbEndOfCurr(pl),getbStrt(pl));

                memFusion(pl);

                //printf("Revisar adelante...\n");
                next(pl);
                //printList(pl);
                //printList(po);

                //printf("Final del bloque cursor: %i - Comienzo del bloque siguiente: %i\n",getbEndOfCurr(pl),getbStrt(pl));

                memFusion(pl);

                prev(pl);
            }
            //else {
            //    printf("No hay ningun byte ocupado que comienza con %i.\n",b);
            //}
        }
        //printList(pl);
        //printList(po);
    }

    int bloquesOcupados = bloqOcup(po);
    int bytesOcupados = memOcup(po);


    //printf("Bytes ocupados: %i - Bloques totales: %i\n",bytesOcupados,bloquesOcupados);

    if (bytesOcupados == 0 && bloquesOcupados == 0) {
        //printf("Toda la memoria dinamica esta libre.\n");

        fprintf(fo, "Toda la memoria dinamica pedida fue liberada\n");
    }
    else {
        //printf("Quedaron bloques de memoria sin liberar.\n");

        fprintf(fo, "Quedaron %i bloques sin liberar (%i bytes)\n",bloquesOcupados,bytesOcupados);
    }


    free(operacion);
    fclose(fp);
    fclose(fo);

    return 0;
}
