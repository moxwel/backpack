/*
    Implementacion de lista enlazada (adaptada)
    https://github.com/moxwel/basics/tree/master/C/11-edd/5-reimplementation
*/

#include "listaEnlazada.h"

Cuerpo* newList(){
    // Asignar memoria para la lista
    Cuerpo* L = (Cuerpo*)malloc(sizeof(Cuerpo));

    // Asignar memoria para el primer nodo
    L->head = L->tail = L->curr = (nodo*)malloc(sizeof(nodo));

    L->curr->info = NULL; L->curr->sig = NULL; L->curr->tipo = 'i';

    L->pos = 0; L->len = 0;

    return L;
}

void insert(Cuerpo* L, char tipoItem, tListElem item){
    nodo* aux = L->curr->sig;

    L->curr->sig = (nodo*)malloc(sizeof(nodo));
    L->curr->sig->info = item;
    L->curr->sig->tipo = tipoItem;
    L->curr->sig->sig = aux;

    // Si el cursor esta al final, el tail debe moverse.
    if (L->curr == L->tail) {
        L->tail = L->curr->sig;
    }

    L->len++;
}

void append(Cuerpo* L, char tipoItem, tListElem item){
    nodo* aux = L->tail;

    L->tail->sig = (nodo*)malloc(sizeof(nodo));
    L->tail->sig->info = item;
    L->tail->sig->tipo = tipoItem;
    L->tail->sig->sig = NULL;

    L->tail = aux->sig;

    L->len++;
}

void printList(Cuerpo* L){
    nodo* view = L->head;

    printf("-----LIST-----\nLen: %d - Pos: %d \n", getLength(L), getPos(L));
    for (int i = -1; i < getLength(L); i++) {
        printf("[%-4d][%c] %-15p", i, view->tipo, view->info);
        if (L->head == view) {
            printf("| H ");
        }
        if (L->curr == view) {
            printf("| C ");
        }
        if (L->tail == view) {
            printf("| T ");
        }
        printf("\n");
        view = view->sig;
    }
    printf("==============\n");
}

int erase(Cuerpo* L){
    // Si la lista esta vacia, entonces no hacer nada
    if (getLength(L) == 0) {
        return 0;
    }

    // Si el cursor esta al final, mover cursor hacia atras
    if (L->curr == L->tail) {
        prev(L);
    }

    nodo* aux = L->curr->sig;

    L->curr->sig = L->curr->sig->sig;

    // Si el tail esta en el nodo a eliminarse, el tail debe moverse.
    if (L->tail == aux) {
        L->tail = L->curr;
    }

    free(aux->info);
    free(aux);

    L->len--;

    return 1;
}

int removeNode(Cuerpo* L){
    // Si la lista esta vacia, entonces no hacer nada
    if (getLength(L) == 0) {
        return 0;
    }

    // Si el cursor esta al final, mover cursor hacia atras
    if (L->curr == L->tail) {
        prev(L);
    }

    nodo* aux = L->curr->sig;

    L->curr->sig = L->curr->sig->sig;

    // Si el tail esta en el nodo a eliminarse, el tail debe moverse.
    if (L->tail == aux) {
        L->tail = L->curr;
    }

    free(aux);

    L->len--;

    return 1;
}

int next(Cuerpo* L){
    // Si la lista esta vacia, entonces no hacer nada
    if (getLength(L) == 0) {
        return 0;
    }

    // Se puede mover hacia adelante hasta que el cursor llegue al tail
    if (L->curr != L->tail) {
        L->curr = L->curr->sig;
        L->pos++;
        return 1;
    }

    return 0;
}

int prev(Cuerpo* L){
    // Si la lista esta vacia, entonces no hacer nada
    if (getLength(L) == 0) {
        return 0;
    }

    // Se puede mover hacia atras hasta que el cursor llegue al head
    if (L->curr != L->head) {
        nodo* aux = L->head;

        while (aux->sig != L->curr) {
            aux = aux->sig;
        }

        L->curr = aux;

        L->pos--;

        return 1;
    }

    return 0;
}

void clearList(Cuerpo* L){
    while(erase(L)){
        continue;
    }
}

void deleteList(Cuerpo* L){
    clearList(L);
    nodo* aux = L->curr;

    free(aux);

    free(L);
}

dataInfo getValue(Cuerpo* L){
    dataInfo result;

    // Si la lista esta vacia, retornar valor vacio
    if (getLength(L) == 0) {
        result.vacio = -9999;
        return result;
    }

    // Si el cursor esta al final, retornar valor actual
    if (L->curr == L->tail) {

        if (L->curr->tipo == 'i') {
            result.i = *(int*)L->curr->info;

        } else if (L->curr->tipo == 'c') {
            result.c = *(char*)L->curr->info;

        } else if (L->curr->tipo == 'h') {
            result.h = (char*)L->curr->info;

        } else if (L->curr->tipo == 'a') {
            result.a = (int*)L->curr->info;
        }

        return result;
    }

    if (L->curr->sig->tipo == 'i') {
        result.i = *(int*)L->curr->sig->info;

    } else if (L->curr->sig->tipo == 'c') {
        result.c = *(char*)L->curr->sig->info;

    } else if (L->curr->sig->tipo == 'h') {
        result.h = (char*)L->curr->sig->info;

    } else if (L->curr->sig->tipo == 'a') {
        result.a = (int*)L->curr->sig->info;
    }

    return result;
}

void printValue(Cuerpo* L){
    dataInfo result;
    printf("[%-4d][", getPos(L));

    // Si la lista esta vacia, retornar valor vacio
    if (getLength(L) == 0) {
        printf(" LISTA VACIA ]\n");
        result.vacio = -9999;
        return;
    }

    // Si el cursor esta al final, retornar valor actual
    if (L->curr == L->tail) {

        if (L->curr->tipo == 'i') {
            result.i = *(int*)L->curr->info;
            printf("i] %-15s | Valor: %d\n", "Int", result.i);

        } else if (L->curr->tipo == 'c') {
            result.c = *(char*)L->curr->info;
            printf("c] %-15s | Valor: %c (%d)\n", "Char", result.c, result.c);

        } else if (L->curr->tipo == 'h') {
            result.h = (char*)L->curr->info;
            printf("h] %-15s | ", "Arreglo de Char");
            printArrayChar(result.h);


        } else if (L->curr->tipo == 'a') {
            result.a = (int*)L->curr->info;
            printf("a] %-15s | ", "Arreglo de Int");
            printArrayInt(result.a);
        }

        return;
    }

    if (L->curr->sig->tipo == 'i') {
        result.i = *(int*)L->curr->sig->info;
        printf("i] %-15s | Valor: %d\n", "Int", result.i);

    } else if (L->curr->sig->tipo == 'c') {
        result.c = *(char*)L->curr->sig->info;
        printf("c] %-15s | Valor: %c (%d)\n", "Char", result.c, result.c);

    } else if (L->curr->sig->tipo == 'h') {
        result.h = (char*)L->curr->sig->info;
        printf("h] %-15s | ", "Arreglo de Char");
        printArrayChar(result.h);

    } else if (L->curr->sig->tipo == 'a') {
        result.a = (int*)L->curr->sig->info;
        printf("a] %-15s | ", "Arreglo de Int");
        printArrayInt(result.a);
    }

    return;
}

int getLength(Cuerpo* L){
    return L->len;
}

int getPos(Cuerpo* L){
    return L->pos;
}

int moveToPos(Cuerpo* L, int target){
    // Si la posicion es negativa, es invalido
    if (target < 0) {
        return 0;
    }

    // Si la posicion sale del rango, es invalido
    if (target >= getLength(L)) {
        return 0;
    }

    moveToStart(L);

    for (int i = 0; i < target; i++) {
        next(L);
    }

    return 1;

}

void moveToStart(Cuerpo* L){
    L->curr = L->head;
    L->pos = 0;
}

void moveToEnd(Cuerpo* L){
    L->curr = L->tail;
    L->pos = getLength(L);

    prev(L);
}

char getType(Cuerpo* L){
    // Si la lista esta vacia
    if (getLength(L) == 0) {
        return 'i';
    }

    // Si el cursor esta al final, retornar valor actual
    if (L->curr == L->tail) {
        return L->curr->tipo;
    }

    return L->curr->sig->tipo;
}

void* getInfoPointer(Cuerpo* L){
    // Si la lista esta vacia
    if (getLength(L) == 0) {
        return NULL;
    }

    // Si el cursor esta al final, retornar valor actual
    if (L->curr == L->tail) {
        return L->curr->info;
    }

    return L->curr->sig->info;
}
