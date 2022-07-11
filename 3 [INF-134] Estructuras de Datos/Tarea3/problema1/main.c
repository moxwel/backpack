#include <stdio.h>
#include <string.h>
#include <stdlib.h>


#include "./TDAs/TDAsueprmercado.c"
#include "./TDAs/hashBase.c"
#include "./fileTools.c"

int main(int argc, char **argv) {

    if (argc != 4) {
        printf("No estan los argumentos requeridos.\n\n"
               "Uso: ./main <1> <2> <3>\n"
               "\t<1> : Archivo binario de productos\n"
               "\t<2> : Archivo binario de ofertas\n"
               "\t<3> : Archivo ASCII de compras de clientes\n\n"
               "Ejemplo: ./main productos.dat ofertas.dat compras.txt\n");

        exit(1);
    }

    int MProductos;
    ranuraProd *tablaProductos = leerProductos(argv[1],&MProductos);

    int MOfertas;
    ranuraOfer *tablaOfertas = leerOfertas(argv[2],&MOfertas);

    // printArrProd(tablaProductos,MProductos);
    // printArrOfer(tablaOfertas,MOfertas);

    int cantRanking = 0;
    leerCompras(argv[3],tablaProductos,MProductos,tablaOfertas,MOfertas,&cantRanking);

    // printArrProd(tablaProductos,MProductos);
    // printf("cantRanking: %i\n",cantRanking);

    rankGen(cantRanking,tablaProductos,MProductos);

    free(tablaProductos);
    free(tablaOfertas);

    return 0;
}
