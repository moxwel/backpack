#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "./tda-clienteBanco.c"
#include "./file_replace.c"
#include "./actualizarSaldos.c"

/**
 * ./main <1> <2>
 ******
 * Argumentos de programa:
 *     1 : Archivo que contiene los clientes
 *     2 : Archivo que contiene las transacciones
 ******
 * Ejemplo:
 *     ./main clientes.dat transacciones.txt
 **/
int main(int argc, char **argv) {

    if (argc != 3) {
        printf("No estan los argumentos requeridos.\n\nUso: ./main <1> <2>\n\t<1> : Archivo binario de clientes\n\t<2> : Archivo de transacciones\n\nEjemplo: ./main clientes.dat transacciones.txt\n");
        return 1;
    }

    actualizarSaldos(argv[1],argv[2]);

    return 0;
}
