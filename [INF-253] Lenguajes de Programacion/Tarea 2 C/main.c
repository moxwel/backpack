#include <stdio.h>
#include "Transimperativer.h"

// ====================
// FUNCIONES DE MEZCLA
// ====================

void MezclarMitad(Transimperativer* T1, Transimperativer* T2){
    int tamT1 = getLength(T1->cuerpo);
    int mitadT1 = techo(tamT1, 2);
    int tamT2 = getLength(T2->cuerpo);
    int mitadT2 = techo(tamT2, 2);

    Cuerpo* aux1 = newList();
    Cuerpo* aux2 = newList();
    moveToStart(T1->cuerpo);
    moveToStart(T2->cuerpo);

    printf("Mitad de T1: %d\n", mitadT1);
    // Añadir las mitades
    for(int i = 0; i < tamT1; i++) {
        if (i < mitadT1) {
            printf("Insertando T1[%d] en aux1....\n", i);
            append(aux1, getType(T1->cuerpo), getInfoPointer(T1->cuerpo));
        } else {
            printf("Insertando T1[%d] en aux2....\n", i);
            append(aux2, getType(T1->cuerpo), getInfoPointer(T1->cuerpo));
        }

        // Eliminar el nodo, pero sin liberar el contenido para evitar dangling
        removeNode(T1->cuerpo);
    }

    printf("Mitad de T2: %d\n", mitadT2);
    for(int i = 0; i < tamT2; i++) {
        if (i < mitadT2) {
            printf("Insertando T2[%d] en aux1....\n", i);
            insert(aux2, getType(T2->cuerpo), getInfoPointer(T2->cuerpo));
            next(aux2);
        } else {
            printf("Insertando T2[%d] en aux2....\n", i);
            append(aux1, getType(T2->cuerpo), getInfoPointer(T2->cuerpo));
        }

        // Eliminar el nodo, pero sin liberar el contenido para evitar dangling
        removeNode(T2->cuerpo);
    }

    // Al final, los dos Transimperativers no van a tener nodos, asi que sera
    // seguro utilizar 'ReemplazarCuerpo()' sin sufrir dangling.

    ReemplazarCuerpo(T1, aux1);
    ReemplazarCuerpo(T2, aux2);
}

void mezclar2(Transimperativer* T1, Transimperativer* T2){
    return;
}
void mezclar3(Transimperativer* T1, Transimperativer* T2){
    return;
}
void mezclar4(Transimperativer* T1, Transimperativer* T2){
    return;
}




// =======================
// FUNCIONES DE SEPARACION
// =======================

void SepararMitad(Transimperativer* padre, Transimperativer* hijo1, Transimperativer* hijo2){
    int tamP = getLength(padre->cuerpo);
    int mitad = techo(tamP, 2);

    moveToStart(padre->cuerpo);

    printf("Mitad del padre: %d\n", mitad);

    moveToStart(padre->cuerpo);
    for (int i = 0; i < tamP; i++) {
        if (i < mitad) {
            printf("Insertando padre[%d] en hijo1....\n", i);
            AnadirAtributo(hijo1, getType(padre->cuerpo), getInfoPointer(padre->cuerpo));
        } else {
            printf("Insertando padre[%d] en hijo2....\n", i);
            AnadirAtributo(hijo2, getType(padre->cuerpo), getInfoPointer(padre->cuerpo));
        }

        // Eliminar el nodo, pero sin liberar el contenido para evitar dangling
        removeNode(padre->cuerpo);
    }

    // Al final, el Transimperativer padre no va a tener nodos, por lo tanto,
    // sera seguro utilizar 'BorrarTransimperativer()' sin sufrir dangling.
}

void separar2(Transimperativer* padre, Transimperativer* hijo1, Transimperativer* hijo2){
    return;
}
void separar3(Transimperativer* padre, Transimperativer* hijo1, Transimperativer* hijo2){
    return;
}
void separar4(Transimperativer* padre, Transimperativer* hijo1, Transimperativer* hijo2){
    return;
}




// ====================

int main() {

    void (*funcionesMezclar[4])(Transimperativer*,Transimperativer*) = {MezclarMitad, mezclar2, mezclar3, mezclar4};
    void (*funcionesSeparar[4])(Transimperativer*,Transimperativer*,Transimperativer*) = {SepararMitad, separar2, separar3, separar4};

    Transimperativer* transimpList[100000];
    int tCreated = 0;
    int running = 1;
    int input;
    int inputI;
    char inputC;
    int selection;
    int selection2;
    int selection3;
    char selectType;
    while(running) {
        printf("\nMAIN MENU\n=========\n");
        printf("Transimperativers creados: %d/100000\n\nAcciones: \n", tCreated);
        printf("    1. Crear Transimperativer\n");
        printf("    2. Mostrar Transimperativer\n");
        printf("    3. Agregar atributo\n");
        printf("    4. Quitar atributo\n");
        printf("    5. Pelear\n");
        printf("    6. Mezclar\n");
        printf("    7. Separar\n");
        printf("    8. Borrar Transimperativer\n");
        printf("    0. Detener programa\n");

        printf("\nSeleccionar accion: ");
        scanf(" %d", &input);

        // Crear transimperativer
        if (input == 1) {
            if (tCreated >= 100000) {
                printf("\n[CREAR TRANSIMPERATIVER]\n/// Numero maximo de Transimperativers creados ///\n");
                continue;
            }

            printf("\n[CREAR TRANSIMPERATIVER]\nElegir funcion de mezcla [0-3]: ");
            scanf(" %d", &selection);

            while ((selection < 0) || (selection > 3)) {
                printf("/// No valido /// Elegir funcion de mezcla [0-3]: ");
                scanf(" %d", &selection);
            }

            printf("Elegir funcion de separacion [0-3]: ");
            scanf(" %d", &selection2);
            while ((selection2 < 0) || (selection2 > 3)) {
                printf("/// No valido /// Elegir funcion de separacion [0-3]: ");
                scanf(" %d", &selection2);
            }

            transimpList[tCreated] = CrearTransimperativer(funcionesMezclar[selection], funcionesSeparar[selection2]);
            tCreated++;

            printf("Listo.\n");

        // Mostrar Transimperativer
        } else if (input == 2) {
            // No se puede hacer nada si no hay Transimperativers creados
            if (tCreated == 0) {
                printf("\n[MOSTRAR TRANSIMPERATIVER]\n/// No hay ningun transimperativer creado aun ///\n");
                continue;
            }

            printf("\n[MOSTRAR TRANSIMPERATIVER]\nElegir transimperativer [0-%d]: ", tCreated-1);
            scanf(" %d", &selection);
            while ((selection >= tCreated) || (selection < 0)) {
                printf("/// No valido /// Elegir transimperativer [0-%d]: ", tCreated-1);
                scanf(" %d", &selection);
            }
            printf("\n");
            MostrarTransimperativer(transimpList[selection]);


        // Agregar atributo
        } else if (input == 3) {
            // No se puede hacer nada si no hay Transimperativers creados
            if (tCreated == 0) {
                printf("\n[AGREGAR ATRIBUTO]\n/// No hay ningun transimperativer creado aun ///\n");
                continue;
            }

            printf("\n[AGREGAR ATRIBUTO]\nElegir transimperativer [0-%d]: ", tCreated-1);
            scanf(" %d", &selection);
            while ((selection >= tCreated) || (selection < 0)) {
                printf("/// No valido /// Elegir transimperativer [0-%d]: ", tCreated-1);
                scanf(" %d", &selection);
            }

            printf("Tipo del atributo [i, c, h, a]: ");
            scanf(" %c", &selectType);
            while ((selectType != 'i') && (selectType != 'c') && (selectType != 'h') && (selectType != 'a')) {
                printf("/// No valido /// Tipo del atributo [i, c, h, a]: ");
                scanf(" %c", &selectType);
            }

            if (selectType == 'i') {
                printf("Añadiendo Int. Ingrese valor: ");
                scanf(" %d", &input);
                int* aux = (int*)malloc(sizeof(int));
                *aux = input;
                AnadirAtributo(transimpList[selection], selectType, aux);

            } else if (selectType == 'c') {
                printf("Añadiendo Char. Ingrese valor: ");
                scanf(" %c", &inputC);
                char* aux = (char*)malloc(sizeof(char));
                *aux = inputC;
                AnadirAtributo(transimpList[selection], selectType, aux);

            } else if (selectType == 'h') {
                printf("Añadiendo Arreglo de Char. Ingrese tamaño del arreglo: ");
                scanf(" %d", &input);
                char* aux = (char*)malloc(sizeof(char) * (input+1));
                aux[0] = input;

                for (int i = 1; i < input+1; i++) {
                    printf("Ingrese elemento en indice %d: ", i);
                    scanf(" %c", &inputC);
                    aux[i] = inputC;
                }

                AnadirAtributo(transimpList[selection], selectType, aux);
            } else if (selectType == 'a') {
                printf("Añadiendo Arreglo de Ints. Ingrese tamaño del arreglo: ");
                scanf(" %d", &input);
                int* aux = (int*)malloc(sizeof(int) * (input+1));
                aux[0] = input;

                for (int i = 1; i < input+1; i++) {
                    printf("Ingrese elemento en indice %d: ", i);
                    scanf(" %d", &inputI);
                    aux[i] = inputI;
                }

                AnadirAtributo(transimpList[selection], selectType, aux);
            }

            MostrarTransimperativer(transimpList[selection]);
            printf("Listo.\n");

        // Quitar Atributo
        } else if (input == 4) {
            // No se puede hacer nada si no hay Transimperativers creados
            if (tCreated == 0) {
                printf("\n[QUITAR ATRIBUTO]\n/// No hay ningun transimperativer creado aun ///\n");
                continue;
            }

            printf("\n[QUITAR ATRIBUTO]\nElegir transimperativer [0-%d]: ", tCreated-1);
            scanf(" %d", &selection);
            while ((selection >= tCreated) || (selection < 0)) {
                printf("/// No valido /// Elegir transimperativer [0-%d]: ", tCreated-1);
                scanf(" %d", &selection);
            }
            printf("\n");
            MostrarTransimperativer(transimpList[selection]);

            // No se puede hacer nada si no hay caracteristicas
            if (TamanoTransimperativer(transimpList[selection]) == 0) {
                printf("/// No hay ninguna caracteristica ///\n");
                continue;
            }

            printf("Indice de caracteristica a eliminar: ");
            scanf(" %d", &inputI);
            while (inputI >= TamanoTransimperativer(transimpList[selection])) {
                printf("/// No valido /// Indice de caracteristica a eliminar: ");
                scanf(" %d", &inputI);
            }

            QuitarAtributo(transimpList[selection], inputI);
            MostrarTransimperativer(transimpList[selection]);
            printf("Listo.\n");

        // Pelear
        } else if (input == 5) {
            // No se puede hacer nada si no hay Transimperativers creados
            if (tCreated < 2) {
                printf("\n[PELEAR]\n/// No hay transimperativer suficientes ///\n");
                continue;
            }

            printf("\n[PELEAR]\nElegir Transimperativer 1 [0-%d]: ", tCreated-1);
            scanf(" %d", &selection);
            while ((selection >= tCreated) || (selection < 0)) {
                printf("/// No valido /// Elegir Transimperativer 1 [0-%d]: ", tCreated-1);
                scanf(" %d", &selection);
            }

            printf("Elegir Transimperativer 2 [0-%d]: ", tCreated-1);
            scanf(" %d", &selection2);
            while ((selection2 >= tCreated) || (selection2 < 0) || (selection == selection2)) {
                printf("/// No valido /// Elegir Transimperativer 2 [0-%d]: ", tCreated-1);
                scanf(" %d", &selection2);
            }

            printf("\n");
            Pelear(transimpList[selection], transimpList[selection2]);

        // Mezclar
        } else if (input == 6) {
            // No se puede hacer nada si no hay Transimperativers creados
            if (tCreated < 2) {
                printf("\n[MEZCLAR]\n/// No hay transimperativer suficientes ///\n");
                continue;
            }

            printf("\n[MEZCLAR]\nElegir Transimperativer 1 [0-%d]: ", tCreated-1);
            scanf(" %d", &selection);
            while ((selection >= tCreated) || (selection < 0)) {
                printf("/// No valido /// Elegir Transimperativer 1 [0-%d]: ", tCreated-1);
                scanf(" %d", &selection);
            }

            printf("Elegir Transimperativer 2 [0-%d]: ", tCreated-1);
            scanf(" %d", &selection2);
            while ((selection2 >= tCreated) || (selection2 < 0) || (selection == selection2)) {
                printf("/// No valido /// Elegir Transimperativer 2 [0-%d]: ", tCreated-1);
                scanf(" %d", &selection2);
            }

            MezclarTransimperativer(transimpList[selection], transimpList[selection2]);
            printf("Listo.\n");

        // Separar
        } else if (input == 7) {
            // No se puede hacer nada si no hay Transimperativers creados
            if (tCreated < 3) {
                printf("\n[SEPARAR]\n/// No hay transimperativer suficientes ///\n");
                continue;
            }

            printf("\n[SEPARAR]\nElegir Transimperativer a separar [0-%d]: ", tCreated-1);
            scanf(" %d", &selection);
            while ((selection >= tCreated) || (selection < 0)) {
                printf("/// No valido /// Elegir Transimperativer a separar [0-%d]: ", tCreated-1);
                scanf(" %d", &selection);
            }

            printf("Elegir Transimperativer hijo 1 [0-%d]: ", tCreated-1);
            scanf(" %d", &selection2);
            while ((selection2 >= tCreated) || (selection2 < 0) || (selection == selection2)) {
                printf("/// No valido /// Elegir Transimperativer hijo 1 [0-%d]: ", tCreated-1);
                scanf(" %d", &selection2);
            }

            printf("Elegir Transimperativer hijo 2 [0-%d]: ", tCreated-1);
            scanf(" %d", &selection3);
            while ((selection3 >= tCreated) || (selection3 < 0) || (selection == selection3) || (selection2 == selection3)) {
                printf("/// No valido /// Elegir Transimperativer hijo 2 [0-%d]: ", tCreated-1);
                scanf(" %d", &selection3);
            }

            SepararTransimperativer(transimpList[selection], transimpList[selection2], transimpList[selection3]);
            removerTransimpLista(transimpList, selection, tCreated);
            tCreated--;

            printf("Listo.\n");



        // Borrar transimperativer
        } else if (input == 8) {
            // No se puede hacer nada si no hay Transimperativers creados
            if (tCreated == 0) {
                printf("\n[BORRAR TRANSIMPERATIVER]\n/// No hay ningun transimperativer creado aun ///\n");
                continue;
            }

            printf("\n[BORRAR TRANSIMPERATIVER]\nElegir Transimperativer a borrar [0-%d]: ", tCreated-1);
            scanf(" %d", &selection);
            while ((selection >= tCreated) || (selection < 0)) {
                printf("/// No valido /// Elegir Transimperativer a borrar [0-%d]: ", tCreated-1);
                scanf(" %d", &selection);
            }

            BorrarTransimperativer(transimpList[selection]);
            removerTransimpLista(transimpList, selection, tCreated);
            tCreated--;

            printf("Listo.\n");



        // Detener programa
        } else if (input == 0) {
            printf("\n[DETENER PROGRAMA]\nLiberando memoria...\n");
            for (int i = 0; i < tCreated; i++) {
                BorrarTransimperativer(transimpList[i]);
            }
            running = 0;
        }
    }



    return 0;
}
