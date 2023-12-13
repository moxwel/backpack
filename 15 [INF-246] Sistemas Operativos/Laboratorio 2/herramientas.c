#include "herramientas.h"

int chequearCelda(t_Mapa* mapa, int map_x, int map_y, int lab_x, int lab_y) {
    // Verificar que en la coordenada (x, y) haya un laberinto
    if (mapa->mtrx_mapa[map_y][map_x] == NULL) {
        printf("No hay un laberinto en la coordenada (%d, %d)\n", map_x, map_y);
        return -1;
    }

    return mapa->mtrx_mapa[map_y][map_x]->mtrx_laberinto[lab_y][lab_x].tipo;
}

int puertaEstaAbierta(t_Mapa* mapa, int map_x, int map_y, int lab_x, int lab_y) {
    // Verificar que en la coordenada (x, y) haya un laberinto
    if (mapa->mtrx_mapa[map_y][map_x] == NULL) {
        printf("No hay un laberinto en la coordenada (%d, %d)\n", map_x, map_y);
        return 0;
    };

    // Verificar que en la coordenada (x, y) haya una puerta
    if (mapa->mtrx_mapa[map_y][map_x]->mtrx_laberinto[lab_y][lab_x].tipo != 2) {
        printf("No hay una puerta en la coordenada (%d, %d)\n", lab_x, lab_y);
        return 0;
    }

    return mapa->mtrx_mapa[map_y][map_x]->mtrx_laberinto[lab_y][lab_x].puertaEstaAbierta;
}

int jugadorTieneLlave(t_Jugador* jugador) {
    return jugador->rol == 1;
}

int abrirPuerta(t_Mapa* mapa, int map_x, int map_y, int lab_x, int lab_y) {
    // Verificar que en la coordenada (x, y) haya un laberinto
    if (mapa->mtrx_mapa[map_y][map_x] == NULL) {
        printf("No hay un laberinto en la coordenada (%d, %d)\n", map_x, map_y);
        return 0;
    };

    // Verificar que en la coordenada (x, y) haya una puerta
    if (mapa->mtrx_mapa[map_y][map_x]->mtrx_laberinto[lab_y][lab_x].tipo != 2) {
        printf("No hay una puerta en la coordenada (%d, %d)\n", lab_x, lab_y);
        return 0;
    }

    // Abrir puerta
    mapa->mtrx_mapa[map_y][map_x]->mtrx_laberinto[lab_y][lab_x].puertaEstaAbierta = 1;
    return 1;
}

int direccionPasaje(t_Mapa* mapa, int map_x, int map_y, int lab_x, int lab_y) {
    // Verificar que en la coordenada (x, y) haya un laberinto
    if (mapa->mtrx_mapa[map_y][map_x] == NULL) {
        printf("No hay un laberinto en la coordenada (%d, %d)\n", map_x, map_y);
        return -1;
    };

    // Verificar que en la coordenada (x, y) haya un pasaje
    if (mapa->mtrx_mapa[map_y][map_x]->mtrx_laberinto[lab_y][lab_x].tipo != 4) {
        printf("No hay un pasaje en la coordenada (%d, %d)\n", lab_x, lab_y);
        return -1;
    }

    return mapa->mtrx_mapa[map_y][map_x]->mtrx_laberinto[lab_y][lab_x].pasajeDireccion;
}

int coinicideTesoro(t_Mapa* mapa, int map_x, int map_y, int lab_x, int lab_y, t_Jugador* jugador) {
    // Verificar que en la coordenada (x, y) haya un laberinto
    if (mapa->mtrx_mapa[map_y][map_x] == NULL) {
        printf("No hay un laberinto en la coordenada (%d, %d)\n", map_x, map_y);
        return 0;
    };

    // Verificar que en la coordenada (x, y) haya un tesoro
    if (mapa->mtrx_mapa[map_y][map_x]->mtrx_laberinto[lab_y][lab_x].tipo != 3) {
        printf("No hay un tesoro en la coordenada (%d, %d)\n", lab_x, lab_y);
        return 0;
    }

    // Verificar que el tesoro coincida con el rol del jugador
    if (mapa->mtrx_mapa[map_y][map_x]->mtrx_laberinto[lab_y][lab_x].tesoroNumero != jugador->numero) {
        printf("El tesoro en la coordenada (%d, %d) no coincide con el rol del jugador\n", lab_x, lab_y);
        return 0;
    }

    return 1;
}

void tomarTesoro(t_Mapa* mapa, int map_x, int map_y, int lab_x, int lab_y, t_Jugador* jugador) {
    // Verificar que en la coordenada (x, y) haya un laberinto
    if (mapa->mtrx_mapa[map_y][map_x] == NULL) {
        printf("No hay un laberinto en la coordenada (%d, %d)\n", map_x, map_y);
        return;
    };

    // Verificar que en la coordenada (x, y) haya un tesoro
    if (mapa->mtrx_mapa[map_y][map_x]->mtrx_laberinto[lab_y][lab_x].tipo != 3) {
        printf("No hay un tesoro en la coordenada (%d, %d)\n", lab_x, lab_y);
        return;
    }

    // Tomar tesoro
    mapa->mtrx_mapa[map_y][map_x]->mtrx_laberinto[lab_y][lab_x].tipo = 0;
    mapa->mtrx_mapa[map_y][map_x]->mtrx_laberinto[lab_y][lab_x].tesoroNumero = 0;
    jugador->tieneTesoro = 1;
}

int jugadorTieneTesoro(t_Jugador* jugador) {
    return jugador->tieneTesoro;
}

int avanzarPasaje(t_Mapa* mapa, int map_x, int map_y, int lab_x, int lab_y, t_Jugador* jugador, t_Laberinto** mazo_laberintos, int n_archivos) {
    // Verificar que en la coordenada (x, y) haya un laberinto
    if (mapa->mtrx_mapa[map_y][map_x] == NULL) {
        printf("No hay un laberinto en la coordenada (%d, %d)\n", map_x, map_y);
        return 0;
    };

    // Verificar que en la coordenada (x, y) haya un pasaje
    if (mapa->mtrx_mapa[map_y][map_x]->mtrx_laberinto[lab_y][lab_x].tipo != 4) {
        printf("No hay un pasaje en la coordenada (%d, %d)\n", lab_x, lab_y);
        return 0;
    }

    int direccionPasaje = mapa->mtrx_mapa[map_y][map_x]->mtrx_laberinto[lab_y][lab_x].pasajeDireccion;

    // Verificar que en la siguiente coordenada no haya un laberinto
    if (direccionPasaje == 0) {
        if (mapa->mtrx_mapa[map_y][map_x]->Nor == NULL) {
            if (unirLaberinto(mapa, mazo_laberintos, n_archivos, map_x, map_y, 0)) {
                printf("Se unio el laberinto NORTE\n");
            } else {
                printf("No se pudo unir el laberinto NORTE\n");
                return 0;
            }
        }
        jugador->map_coord.x = map_x;
        jugador->map_coord.y = map_y - 1;
        jugador->lab_coord.x = 2;
        jugador->lab_coord.y = 4;
        return 1;
    } else if (direccionPasaje == 1) {
        if (mapa->mtrx_mapa[map_y][map_x]->Oes == NULL) {
            if (unirLaberinto(mapa, mazo_laberintos, n_archivos, map_x, map_y, 1)) {
                printf("Se unio el laberinto OESTE\n");
            } else {
                printf("No se pudo unir el laberinto OESTE\n");
                return 0;
            }
        }
        jugador->map_coord.x = map_x - 1;
        jugador->map_coord.y = map_y;
        jugador->lab_coord.x = 4;
        jugador->lab_coord.y = 2;
        return 1;
    } else if (direccionPasaje == 2) {
        if (mapa->mtrx_mapa[map_y][map_x]->Sur == NULL) {
            if (unirLaberinto(mapa, mazo_laberintos, n_archivos, map_x, map_y, 2)) {
                printf("Se unio el laberinto SUR\n");
            } else {
                printf("No se pudo unir el laberinto SUR\n");
                return 0;
            }
        }
        jugador->map_coord.x = map_x;
        jugador->map_coord.y = map_y + 1;
        jugador->lab_coord.x = 2;
        jugador->lab_coord.y = 0;
        return 1;
    } else if (direccionPasaje == 3) {
        if (mapa->mtrx_mapa[map_y][map_x]->Est == NULL) {
            if (unirLaberinto(mapa, mazo_laberintos, n_archivos, map_x, map_y, 3)) {
                printf("Se unio el laberinto ESTE\n");
            } else {
                printf("No se pudo unir el laberinto ESTE\n");
                return 0;
            }
        }
        jugador->map_coord.x = map_x + 1;
        jugador->map_coord.y = map_y;
        jugador->lab_coord.x = 0;
        jugador->lab_coord.y = 2;
        return 1;
    }

    return 0;
}

int moverJugador(t_Mapa* mapa, int map_x, int map_y, int lab_x, int lab_y, t_Jugador* jugador, int direccion) {
    // Verificar que en la coordenada (x, y) haya un laberinto
    if (mapa->mtrx_mapa[map_y][map_x] == NULL) {
        printf("No hay un laberinto en la coordenada (%d, %d)\n", map_x, map_y);
        return 0;
    };

    // Verificar que la direccion sea valida
    if (direccion < 0 || direccion > 3) {
        printf("Direccion invalida: %d\n", direccion);
        return 0;
    }

    // Mover jugador verificando que no se salga del laberinto
    if (direccion == 0) {
        if (lab_y == 0) {
            printf("No se puede mover hacia el NORTE\n");
            return 0;
        }
        jugador->lab_coord.y--;
    } else if (direccion == 1) {
        if (lab_x == 0) {
            printf("No se puede mover hacia el OESTE\n");
            return 0;
        }
        jugador->lab_coord.x--;
    } else if (direccion == 2) {
        if (lab_y == 4) {
            printf("No se puede mover hacia el SUR\n");
            return 0;
        }
        jugador->lab_coord.y++;
    } else if (direccion == 3) {
        if (lab_x == 4) {
            printf("No se puede mover hacia el ESTE\n");
            return 0;
        }
        jugador->lab_coord.x++;
    }

    jugador->seMovio = 1;
    return 1;
}
