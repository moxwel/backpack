#include "Transimperativer.h"

Transimperativer* CrearTransimperativer(
                           void (*mezclarF)(Transimperativer* T1, Transimperativer* T2),
                           void (*separarF)(Transimperativer* padre, Transimperativer* hijo1, Transimperativer* hijo2)){

    // Asignar memoria al transimperativer junto a su cuerpo
    Transimperativer* T = (Transimperativer*)malloc(sizeof(Transimperativer));
    T->cuerpo = newList();

    // Asociar funciones de mezcla y separacion
    T->mezclar = mezclarF;
    T->separar = separarF;

    return T;
}

void BorrarTransimperativer(Transimperativer* T) {
    // Liberando memoria del cuerpo primero
    deleteList(T->cuerpo);

    // Liberando memeoria del Transimperativer completo
    free(T);
}

void MostrarTransimperativer(Transimperativer* T){
    int tam = getLength(T->cuerpo);

    printf("-- TRANSIMPERATIVER --\nLen: %d\nLista de caracteristicas:\n", tam);
    if (tam == 0) {
        printf("[ LISTA VACIA ]\n======================\n");
        return;
    }

    moveToStart(T->cuerpo);
    for (int i = 0; i < tam; i++){
        printValue(T->cuerpo);
        next(T->cuerpo);
    }
    printf("======================\n");
}

void AnadirAtributo(Transimperativer* T, char tipo, void* caracteristica){
    // Las caractersitcias nuevas se agregan al final
    append(T->cuerpo, tipo, caracteristica);
}

void QuitarAtributo(Transimperativer* T, int indice){
    moveToPos(T->cuerpo, indice);
    erase(T->cuerpo);
}

void Pelear(Transimperativer* T1, Transimperativer* T2){
    int tam1 = getLength(T1->cuerpo);
    int tam2 = getLength(T2->cuerpo);
    int puntaje1 = 0;
    int puntaje2 = 0;

    printf("------------ PELEA ------------\n");
    // Calcular Transimperativer 1
    if (tam1 != 0) {
        moveToStart(T1->cuerpo);
        for (int i = 0; i < tam1; i++){

            if (getType(T1->cuerpo) == 'i') {
                puntaje1 += getValue(T1->cuerpo).i * 2;

            } else if (getType(T1->cuerpo) == 'c') {
                puntaje1 += getValue(T1->cuerpo).c / 2;

            } else if (getType(T1->cuerpo) == 'h') {
                puntaje1 += averageArrayChar(getValue(T1->cuerpo).h);

            } else if (getType(T1->cuerpo) == 'a') {
                puntaje1 += averageArrayInt(getValue(T1->cuerpo).a);
            }

            next(T1->cuerpo);
        }
    }

    // Calcular Transimperativer 2
    if (tam2 != 0) {
        moveToStart(T2->cuerpo);
        for (int i = 0; i < tam2; i++){

            if (getType(T2->cuerpo) == 'i') {
                puntaje2 += getValue(T2->cuerpo).i * 2;

            } else if (getType(T2->cuerpo) == 'c') {
                puntaje2 += getValue(T2->cuerpo).c / 2;

            } else if (getType(T2->cuerpo) == 'h') {
                puntaje2 += averageArrayChar(getValue(T2->cuerpo).h);

            } else if (getType(T2->cuerpo) == 'a') {
                puntaje2 += averageArrayInt(getValue(T2->cuerpo).a);
            }

            next(T2->cuerpo);
        }
    }

    printf("Puntaje de Transimperativer 1: %d\nPuntaje de Transimperativer 2: %d\n", puntaje1, puntaje2);

    if (puntaje1 > puntaje2) {
        printf("=== TRANSIMPERATIVER 1 GANA ===\n");
    } else if (puntaje1 < puntaje2) {
        printf("=== TRANSIMPERATIVER 2 GANA ===\n");
    } else if (puntaje1 == puntaje2) {
        printf("=== EMPATE ===\n");
    }
}

void MezclarTransimperativer(Transimperativer* T1, Transimperativer* T2){
    T1->mezclar(T1, T2);
}

void SepararTransimperativer(Transimperativer* padre, Transimperativer* hijo1, Transimperativer* hijo2){
    padre->separar(padre, hijo1, hijo2);

    // printf("Borrando padre\n");
    BorrarTransimperativer(padre);
}

void ReemplazarCuerpo(Transimperativer* T, Cuerpo* L){
    deleteList(T->cuerpo);
    T->cuerpo = L;
}

int TamanoTransimperativer(Transimperativer* T){
    return T->cuerpo->len;
}

void removerTransimpLista(Transimperativer* array[], int indice, int size){
    for (int i = indice; i < size; i++) {
        array[i] = array[i + 1];
    }
}
