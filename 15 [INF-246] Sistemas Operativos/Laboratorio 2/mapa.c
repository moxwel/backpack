#include "mapa.h"

t_Mapa* crearMapa(t_Laberinto** mazo_laberintos, int n_laberintos) {
    // Crear el mapa y asignar la matriz de laberintos
    t_Mapa* nuevo_mapa = malloc(sizeof(t_Mapa));

    // Crear matriz de punteros a laberintos
    t_Laberinto*** matriz_laberintos = malloc(11 * sizeof(t_Laberinto**));
    for (int i = 0; i < 11; i++) {
        matriz_laberintos[i] = malloc(11 * sizeof(t_Laberinto*));

        for (int j = 0; j < 11; j++) {
            matriz_laberintos[i][j] = NULL;
        }
    }

    // Desde el mazo de laberintos, buscar el inicial y asignarlo al medio del mapa
    for (int i = 0; i < n_laberintos; i++) {
        if (mazo_laberintos[i]->esInicial) {
            matriz_laberintos[5][5] = mazo_laberintos[i];
            mazo_laberintos[i]->estaUsado = 1;
            break;
        }
    }

    // Guardar la matriz de laberintos en el mapa
    nuevo_mapa->mtrx_mapa = matriz_laberintos;

    return nuevo_mapa;
}


void imprimirMapa(t_Mapa* mapa) {
    printf(" --- Mapa ---\n");
    for (int i = 0; i < 11; i++) {
        for (int j = 0; j < 11; j++) {
            if (mapa->mtrx_mapa[i][j] == NULL) {
                printf("- ");
            } else {
                printf("%d ", mapa->mtrx_mapa[i][j]->tipoSalidas);
            }
        }
        printf("\n");
    }
    printf(" ------------\n");
}

void liberarMapa(t_Mapa* mapa) {
    for (int i = 0; i < 11; i++) {
        free(mapa->mtrx_mapa[i]);
    }
    free(mapa->mtrx_mapa);
    free(mapa);
}


int unirLaberinto(t_Mapa* mapa, t_Laberinto** mazo_laberintos, int n_laberintos, int x, int y, int direccion) {
    // Verificar que en la coordenada (x, y) haya un laberinto
    if (mapa->mtrx_mapa[y][x] == NULL) {
        printf("No hay un laberinto en la coordenada (%d, %d)\n", x, y);
        return 0;
    }

    /**
     * Direcciones:
     * 0: N
     * 1: O
     * 2: S
     * 3: E
    */

    /**
     * Tipos de salidas de un laberinto:
     * 0: N O S E
     * 1: N O E
     * 2: N O S
     * 3: O S E
     * 4: N S E
     * 5: N O
     * 6: O S
     * 7: S E
     * 8: N E
     *
     * El laberinto es de 5x5
    */

    // Marcas de uniones compatibles en el laberinto indicado
    int N = 0, O = 0, S = 0, E = 0;

    // Verificar que el laberinto en la coordenada (x, y) tenga la salida en la direccion indicada
    if      (mapa->mtrx_mapa[y][x]->tipoSalidas == 0) { N = 1; O = 1; S = 1; E = 1; }
    else if (mapa->mtrx_mapa[y][x]->tipoSalidas == 1) { N = 1; O = 1; E = 1; }
    else if (mapa->mtrx_mapa[y][x]->tipoSalidas == 2) { N = 1; O = 1; S = 1; }
    else if (mapa->mtrx_mapa[y][x]->tipoSalidas == 3) { O = 1; S = 1; E = 1; }
    else if (mapa->mtrx_mapa[y][x]->tipoSalidas == 4) { N = 1; S = 1; E = 1; }
    else if (mapa->mtrx_mapa[y][x]->tipoSalidas == 5) { N = 1; O = 1; }
    else if (mapa->mtrx_mapa[y][x]->tipoSalidas == 6) { O = 1; S = 1; }
    else if (mapa->mtrx_mapa[y][x]->tipoSalidas == 7) { S = 1; E = 1; }
    else if (mapa->mtrx_mapa[y][x]->tipoSalidas == 8) { N = 1; E = 1; }
    else {
        printf("Error al leer el laberinto en la coordenada (%d, %d)\n", x, y);
        return 0;
    }

    printf("Salidas del laberinto en la coordenada (%d, %d): N:%d O:%d S:%d E:%d\n", x, y, N, O, S, E);

    // Verificar que el laberinto en la coordenada (x, y) tenga la salida en la direccion indicada
    switch (direccion) {
        case 0:
            if (!N) {
                printf("El laberinto en la coordenada (%d, %d) no tiene salida al norte\n", x, y);
                return 0;
            }
            break;
        case 1:
            if (!O) {
                printf("El laberinto en la coordenada (%d, %d) no tiene salida al oeste\n", x, y);
                return 0;
            }
            break;
        case 2:
            if (!S) {
                printf("El laberinto en la coordenada (%d, %d) no tiene salida al sur\n", x, y);
                return 0;
            }
            break;
        case 3:
            if (!E) {
                printf("El laberinto en la coordenada (%d, %d) no tiene salida al este\n", x, y);
                return 0;
            }
            break;
        default:
            printf("Direccion invalida\n");
            return 0;
    }

    printf("El laberinto en la coordenada (%d, %d) tiene salida en la direccion %d\n", x, y, direccion);

    // Verificar que en la coordenada siguiente no haya un laberinto
    switch (direccion) {
        case 0: // Verificar que no haya un laberinto al norte
            if (mapa->mtrx_mapa[y - 1][x] != NULL) {
                printf("Ya hay un laberinto en la coordenada siguiente (%d, %d)\n", x, y-1);
                return 0;
            }
            break;
        case 1: // Verificar que no haya un laberinto al oeste
            if (mapa->mtrx_mapa[y][x - 1] != NULL) {
                printf("Ya hay un laberinto en la coordenada siguiente (%d, %d)\n", x-1, y);
                return 0;
            }
            break;
        case 2: // Verificar que no haya un laberinto al sur
            if (mapa->mtrx_mapa[y + 1][x] != NULL) {
                printf("Ya hay un laberinto en la coordenada siguiente (%d, %d)\n", x, y+1);
                return 0;
            }
            break;
        case 3: // Verificar que no haya un laberinto al este
            if (mapa->mtrx_mapa[y][x + 1] != NULL) {
                printf("Ya hay un laberinto en la coordenada siguiente (%d, %d)\n", x+1, y);
                return 0;
            }
            break;
        default:
            printf("Direccion invalida\n");
            return 0;
    }

    printf("No hay un laberinto en la coordenada siguiente\n");

    // Dependiendo de la direccion, buscar un laberinto compatible en el mazo
    switch (direccion) {
        case 0: // Añadir laberinto al norte
            for (int i = 0; i < n_laberintos; i++) {
                int tipoSalidas = mazo_laberintos[i]->tipoSalidas;

                // 'tipoSalidas' debe ser 0,2,3,4,6,7
                if (tipoSalidas == 0 || tipoSalidas == 2 || tipoSalidas == 3 || tipoSalidas == 4 || tipoSalidas == 6 || tipoSalidas == 7) {
                    if (!mazo_laberintos[i]->estaUsado) {
                        printf("Se añadio el laberinto %d al norte en la coordenada (%d, %d)\n", i, x, y-1);
                        mapa->mtrx_mapa[y - 1][x] = mazo_laberintos[i];
                        mazo_laberintos[i]->estaUsado = 1;

                        // Conectar laberintos
                        mapa->mtrx_mapa[y][x]->Nor = mapa->mtrx_mapa[y - 1][x];
                        mapa->mtrx_mapa[y - 1][x]->Sur = mapa->mtrx_mapa[y][x];

                        return 1;
                    }
                }
            }
            break;
        case 1: // Añadir laberinto al oeste
            for (int i = 0; i < n_laberintos; i++) {
                int tipoSalidas = mazo_laberintos[i]->tipoSalidas;

                // 'tipoSalidas' debe ser 0,1,3,4,7,8
                if (tipoSalidas == 0 || tipoSalidas == 1 || tipoSalidas == 3 || tipoSalidas == 4 || tipoSalidas == 7 || tipoSalidas == 8) {
                    if (!mazo_laberintos[i]->estaUsado) {
                        printf("Se añadio el laberinto %d al oeste en la coordenada (%d, %d)\n", i, x-1, y);
                        mapa->mtrx_mapa[y][x - 1] = mazo_laberintos[i];
                        mazo_laberintos[i]->estaUsado = 1;

                        // Conectar laberintos
                        mapa->mtrx_mapa[y][x]->Oes = mapa->mtrx_mapa[y][x - 1];
                        mapa->mtrx_mapa[y][x - 1]->Est = mapa->mtrx_mapa[y][x];

                        return 1;
                    }
                }
            }
            break;
        case 2: // Añadir laberinto al sur
            for (int i = 0; i < n_laberintos; i++) {
                int tipoSalidas = mazo_laberintos[i]->tipoSalidas;

                // 'tipoSalidas' debe ser 0,1,2,4,5,8
                if (tipoSalidas == 0 || tipoSalidas == 1 || tipoSalidas == 2 || tipoSalidas == 4 || tipoSalidas == 5 || tipoSalidas == 8) {
                    if (!mazo_laberintos[i]->estaUsado) {
                        printf("Se añadio el laberinto %d al sur en la coordenada (%d, %d)\n", i, x, y+1);
                        mapa->mtrx_mapa[y + 1][x] = mazo_laberintos[i];
                        mazo_laberintos[i]->estaUsado = 1;

                        // Conectar laberintos
                        mapa->mtrx_mapa[y][x]->Sur = mapa->mtrx_mapa[y + 1][x];
                        mapa->mtrx_mapa[y + 1][x]->Nor = mapa->mtrx_mapa[y][x];

                        return 1;
                    }
                }
            }
            break;
        case 3: // Añadir laberinto al este
            for (int i = 0; i < n_laberintos; i++) {
                int tipoSalidas = mazo_laberintos[i]->tipoSalidas;

                // 'tipoSalidas' debe ser 0,1,2,3,5,6
                if (tipoSalidas == 0 || tipoSalidas == 1 || tipoSalidas == 2 || tipoSalidas == 3 || tipoSalidas == 5 || tipoSalidas == 6) {
                    if (!mazo_laberintos[i]->estaUsado) {
                        printf("Se añadio el laberinto %d al este en la coordenada (%d, %d)\n", i, x+1, y);
                        mapa->mtrx_mapa[y][x + 1] = mazo_laberintos[i];
                        mazo_laberintos[i]->estaUsado = 1;

                        // Conectar laberintos
                        mapa->mtrx_mapa[y][x]->Est = mapa->mtrx_mapa[y][x + 1];
                        mapa->mtrx_mapa[y][x + 1]->Oes = mapa->mtrx_mapa[y][x];

                        return 1;
                    }
                }
            }
            break;
        default:
            printf("Direccion invalida\n");
            return 0;
    }
    return 0;
}
