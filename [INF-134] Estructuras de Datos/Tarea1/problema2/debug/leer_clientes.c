// ESTE ES UN PROGRAMA AUXILIAR Y ANALOGO A 'GENERAR_CLIENTES.C', SIRVE PARA LEER ARCHIVOS BINARIOS DE CLIENTES

#include <stdio.h>
#include <stdlib.h>

#include "../tda-clienteBanco.c"

/**
 * ./leer_clientes <1>
 ******
 * Argumentos de programa:
 *     1 : Archivo binario que contiene los clientes
 ******
 * Ejemplo:
 *     ./leer_clientes clientes.dat
 **/
int main(int argc, char **argv) {

    if (argc != 2) {
        printf("No estan los argumentos requeridos.\n\nUso: ./leer_clientes <1>\n\t<1> : Archivo binario de clientes\n\nEjemplo: ./leer_clientes clientes.dat\n");
        return 1;
    }

    char *file = argv[1];

    FILE *fp = fopen(file,"r");
    if (fp == NULL) {
        printf("Error al abrir archivo.\n");
        exit(1);
    }

    clienteBanco leer;

    while (fread(&leer,sizeof(clienteBanco),1,fp) != 0) {
        printf("Numero de cuenta: %i\nSaldo: %i\nNombre: %s\nDicreecion: %s\n\n",
               getCuenta(&leer),getSaldo(&leer),getNombre(&leer),getDireccion(&leer));
    }

    fclose(fp);

    return 0;
}
