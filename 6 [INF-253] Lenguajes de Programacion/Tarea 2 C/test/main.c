#include <stdio.h>
#include "Transimperativer.h"

void MezclarMitad(Transimperativer* T1, Transimperativer* T2){
    int tamT1 = getLength(T1->cuerpo);
    int mitadT1 = techo(tamT1, 2);
    int tamT2 = getLength(T2->cuerpo);
    int mitadT2 = techo(tamT2, 2);

    Cuerpo* aux1 = newList();
    Cuerpo* aux2 = newList();


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
        next(T1->cuerpo);
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
        next(T2->cuerpo);
    }

    ReemplazarCuerpo(T1, aux1);
    ReemplazarCuerpo(T2, aux2);
}

void SepararMitad(Transimperativer* padre, Transimperativer* hijo1, Transimperativer* hijo2){
    int tamP = getLength(padre->cuerpo);
    int mitad = techo(tamP, 2);

    printf("--- SEPARANDO ---\nTamaño del padre: %d - Mitad: %d\n", tamP, mitad);

    moveToStart(padre->cuerpo);
    for (int i = 0; i < tamP; i++) {
        if (i < mitad) {
            printf("Insertando padre[%d] en hijo1....\n", i);
            AnadirAtributo(hijo1, getType(padre->cuerpo), getInfoPointer(padre->cuerpo));
        } else {
            printf("Insertando padre[%d] en hijo2....\n", i);
            AnadirAtributo(hijo2, getType(padre->cuerpo), getInfoPointer(padre->cuerpo));
        }
        next(padre->cuerpo);
    }

    printf("=================\n");
}

int main() {
    int num1 = 1;
    int num2 = 2;
    int num3 = 3;
    int num4 = 4;
    int num5 = 5;
    int num6 = 6;
    int num7 = 7;
    int num8 = 8;
    int num9 = 9;
    char char1 = 'a';
    int a[] = {5, 21, 321, 1432, 3214, 43214};
    char b[] = {7, 'c', 'a', 'c', 'a', 'p', 'i', 's'};
    int c[] = {4, 1, 2, 3, 4};
    char char2 = 'z';

    Transimperativer* mi_robot = CrearTransimperativer(MezclarMitad, SepararMitad);



    Transimperativer* otro_robot = CrearTransimperativer(MezclarMitad, SepararMitad);

    Transimperativer* robobo = CrearTransimperativer(MezclarMitad, SepararMitad);




    AnadirAtributo(mi_robot, 'a', a);
    AnadirAtributo(mi_robot, 'i', &num1);
    AnadirAtributo(mi_robot, 'c', &char1);
    AnadirAtributo(mi_robot, 'i', &num2);
    AnadirAtributo(mi_robot, 'i', &num3);
    AnadirAtributo(mi_robot, 'h', b);

    // AnadirAtributo(otro_robot, 'i', &num4);
    // AnadirAtributo(otro_robot, 'i', &num5);
    // AnadirAtributo(otro_robot, 'i', &num6);
    // AnadirAtributo(otro_robot, 'i', &num7);
    // AnadirAtributo(otro_robot, 'i', &num8);
    // AnadirAtributo(otro_robot, 'i', &num9);

    printList(mi_robot->cuerpo);
    printValue(mi_robot->cuerpo);



    printf("mi_robot\n");
    MostrarTransimperativer(mi_robot);
    printf("otro_robot\n");
    MostrarTransimperativer(otro_robot);
    printf("robobo\n");
    MostrarTransimperativer(robobo);



    MezclarTransimperativer(mi_robot, otro_robot);





    printf("mi_robot\n");
    MostrarTransimperativer(mi_robot);
    printf("otro_robot\n");
    MostrarTransimperativer(otro_robot);








    BorrarTransimperativer(mi_robot);
    BorrarTransimperativer(otro_robot);
    BorrarTransimperativer(robobo);



    printf("%d\n", techo(1,2));

    return 0;
}
