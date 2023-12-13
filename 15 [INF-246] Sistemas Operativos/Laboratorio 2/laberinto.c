#include "laberinto.h"

t_Laberinto* crearLaberinto(char* nombre_archivo) {
    // Abrir el archivo
    FILE* archivo = fopen(nombre_archivo, "r");
    if (archivo == NULL) {
        printf("Error al abrir el archivo %s\n", nombre_archivo);
        return NULL;
    }





    // Crear el laberinto y asignar la matriz de celdas
    t_Laberinto* nuevo_laberinto = malloc(sizeof(t_Laberinto));

    nuevo_laberinto->estaUsado = 0;

    if (strcmp(nombre_archivo, "./laberinto/Inicio.txt") == 0) {
        nuevo_laberinto->esInicial = 1;
    } else {
        nuevo_laberinto->esInicial = 0;
    }






    // Crear la matriz de celdas de 5x5
    t_Celda** matriz_celdas = malloc(5 * sizeof(t_Celda*));
    for (int i = 0; i < 5; i++) {
        matriz_celdas[i] = malloc(5 * sizeof(t_Celda));
    }



    // Marcas para las salidas
    int N = 0, S = 0, E = 0, O = 0;



    // Leer las celdas del archivo y crear los elementos t_Celda correspondientes
    for (int i = 0; i < 5; i++) {

        // Leer una linea del archivo
        char linea[20];
        fgets(linea, 20, archivo);

        // Quitar el salto de linea
        linea[strcspn(linea, "\n")] = '\0';
        linea[strcspn(linea, "\r")] = '\0';

        // printf("i = %d - Linea: \"%s\"\n", i, linea);





        // Tokenizar la linea

        /**
         * Caracteres del archivo:
         * 0: Nada
         * /: Pared
         * B: Pasaje
         * E: Puerta
         * J1, J2, J3, J4: nada
        */

        char* token = strtok(linea, " ");
        int j = 0;
        while (token != NULL) {

            // printf("  j = %d - Token: %s\n", j, token);

            if      (strcmp(token, "0") == 0) { matriz_celdas[i][j].tipo = 0; } // Nada
            else if (strcmp(token, "/") == 0) { matriz_celdas[i][j].tipo = 1; } // Pared



            // Pasaje
            else if (strcmp(token, "B") == 0) {
                matriz_celdas[i][j].tipo = 4;

                // Hay que ver en que bordes se encuentra el pasaje
                if      (i == 0) { matriz_celdas[i][j].pasajeDireccion = 0; N = 1; }
                else if (j == 0) { matriz_celdas[i][j].pasajeDireccion = 1; O = 1; }
                else if (i == 4) { matriz_celdas[i][j].pasajeDireccion = 2; S = 1; }
                else if (j == 4) { matriz_celdas[i][j].pasajeDireccion = 3; E = 1; }
                else {
                    printf("    Error al identificar posicion del pasaje.\n");
                    return NULL;
                }
                // printf("    Pasaje en direccion %d\n", matriz_celdas[i][j].pasajeDireccion);



            // Puerta
            } else if (strcmp(token, "E") == 0) {
                matriz_celdas[i][j].tipo = 2;
                matriz_celdas[i][j].puertaEstaAbierta = 0;
            }



            // Tesoros
            else if (strcmp(token, "1") == 0) {
                matriz_celdas[i][j].tipo = 3;
                matriz_celdas[i][j].tesoroNumero = 1;
            } else if (strcmp(token, "2") == 0) {
                matriz_celdas[i][j].tipo = 3;
                matriz_celdas[i][j].tesoroNumero = 2;
            } else if (strcmp(token, "3") == 0) {
                matriz_celdas[i][j].tipo = 3;
                matriz_celdas[i][j].tesoroNumero = 3;
            } else if (strcmp(token, "4") == 0) {
                matriz_celdas[i][j].tipo = 3;
                matriz_celdas[i][j].tesoroNumero = 4;
            }



            // Jugadores (no se hace nada con el)
            else if (strcmp(token, "J1") == 0) { matriz_celdas[i][j].tipo = 0; }
            else if (strcmp(token, "J2") == 0) { matriz_celdas[i][j].tipo = 0; }
            else if (strcmp(token, "J3") == 0) { matriz_celdas[i][j].tipo = 0; }
            else if (strcmp(token, "J4") == 0) { matriz_celdas[i][j].tipo = 0; }
            else {
                printf("    Error al parsear el token.\n");
                return NULL;
            }


            // Leer el siguiente token
            token = strtok(NULL, " ");
            j++;
        }

        // printf("  Fin de la linea\n");
    }



    // Identificar las salidas del laberinto
    if      (N && O && S && E) { nuevo_laberinto->tipoSalidas = 0; }
    else if (N && O && E)      { nuevo_laberinto->tipoSalidas = 1; }
    else if (N && O && S)      { nuevo_laberinto->tipoSalidas = 2; }
    else if (O && S && E)      { nuevo_laberinto->tipoSalidas = 3; }
    else if (N && S && E)      { nuevo_laberinto->tipoSalidas = 4; }
    else if (N && O)           { nuevo_laberinto->tipoSalidas = 5; }
    else if (O && S)           { nuevo_laberinto->tipoSalidas = 6; }
    else if (S && E)           { nuevo_laberinto->tipoSalidas = 7; }
    else if (N && E)           { nuevo_laberinto->tipoSalidas = 8; }
    else {
        printf("Error al leer el archivo %s\n", nombre_archivo);
        return NULL;
    }



    // Cerrar el archivo
    fclose(archivo);

    // Guardar la matriz de celdas en el laberinto
    nuevo_laberinto->mtrx_laberinto = matriz_celdas;

    return nuevo_laberinto;
}




void liberarLaberinto(t_Laberinto* laberinto) {
    for (int i = 0; i < 5; i++) {
        free(laberinto->mtrx_laberinto[i]);
    }
    free(laberinto->mtrx_laberinto);
    free(laberinto);
}





void imprimirLaberinto(t_Laberinto* laberinto) {

    // Verificar que el laberinto no sea NULL
    if (laberinto == NULL) {
        printf("No hay laberinto.\n");
        return;
    }

    printf(" --- Laberinto ---\n");
    for (int i = 0; i < 5; i++) {
        printf("  ");
        for (int j = 0; j < 5; j++) {
            printf("%d ", laberinto->mtrx_laberinto[i][j].tipo);
        }
        printf("\n");
    }
    printf("Es inicial: %d\n", laberinto->esInicial);
    printf("Tipo de salida: %d\n", laberinto->tipoSalidas);
    printf("Esta usado: %d\n", laberinto->estaUsado);

    printf(" ----------------\n");
}
