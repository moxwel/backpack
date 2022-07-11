// ESTE ES UN PROGRAMA AUXILIAR Y ANALOGO A 'productosGen.C', SIRVE PARA LEER ARCHIVOS BINARIOS DE PRODUCTOS

#include <stdio.h>
#include <stdlib.h>

typedef struct {
    int codigo_producto;
    char nombre_producto[31];
    int precio;
} producto;

/**
 * ./leerProductosDat <1>
 ******
 * Argumentos de programa:
 *     1 : Archivo binario que contiene los productos
 ******
 * Ejemplo:
 *     ./leerProductos productos.dat
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

    producto leer;

    int total_productos;
    fread(&total_productos,sizeof(int),1,fp);

    while (fread(&leer,sizeof(producto),1,fp) != 0) {
        printf("Codigo: %i - Nombre: %s - Precio: %i\n\n",leer.codigo_producto,leer.nombre_producto,leer.precio);
    }

    fclose(fp);

    return 0;
}
