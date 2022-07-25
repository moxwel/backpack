// ESTE ES UN PROGRAMA AUXILIAR Y ANALOGO A 'ofertasGen.C', SIRVE PARA LEER ARCHIVOS BINARIOS DE OFERTAS

#include <stdio.h>
#include <stdlib.h>

typedef struct {
    int codigo_producto;
    int cantidad_descuento;
    int monto_descuento;
} oferta;

/**
 * ./leerOfertasDat <1>
 ******
 * Argumentos de programa:
 *     1 : Archivo binario que contiene las ofertas
 ******
 * Ejemplo:
 *     ./leerOfertas ofertas.dat
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

    oferta leer;

    int total_ofertas;
    fread(&total_ofertas,sizeof(int),1,fp);

    while (fread(&leer,sizeof(oferta),1,fp) != 0) {
        printf("Codigo: %i - Cantidad: %i - Descuento: %i\n\n",leer.codigo_producto,leer.cantidad_descuento,leer.monto_descuento);
    }

    fclose(fp);

    return 0;
}
