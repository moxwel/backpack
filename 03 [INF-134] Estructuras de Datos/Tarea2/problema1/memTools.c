/**
 * int memInsert
 ******
 * Retorna la posicion EN LA LISTA del primer bloque donde quepan
 * los bytes que se quieren ingresar. Si es 0, significa que no se puede
 * alocar ese tamaño de memoria.
 ******
 * Input:
 *     int bytes  : Cantidad de bytes que se quieren ingresar
 *     tLista  *L : Lista de bloques de memoria disponible donde buscar
 ******
 * Return:
 *     int : Posicion de la lista 'L' donde esta el bloque de memoria libre. Si
 *           es 0, significa que no se puede alocar.
 **/
int memInsert(int bytes, tLista *L) {
    tNodo* aux = L->head->sig;

    for (int i = 1; i < length(L); i++) {
        if (bytes <= memSize(aux)) {
            return i;
        }
        aux = aux->sig;
    }

    return 0;
}

/**
 * void memMalloc
 ******
 * Extrae cierta cantidad de memoria libre para ser utilizada, moviendo
 * este bloque a la lista de memoria ocupada.
 ******
 * Input:
 *     tLista *l1  : Lista de memoria disponible
 *     tLista *t2  : Lista de memoria ocupada
 *     int nBloque : Posicion en la lista 'l1' donde se encuentra el bloque
 *                   con espacio libre.
 *     int tam     : Cantidad de bytes a alocar
 **/
void memMalloc(tLista *l1, tLista *l2, int nBloque, int tam) {
    //printf("Moviendo memoria...\n");
    moveToPos(l1,nBloque-1);

    int byteComienzo = getbStrt(l1);
    int byteFinal = tam + byteComienzo - 1;

    append(l2, byteComienzo, byteFinal);

    setbStrt(l1, byteFinal + 1);

    //printList(l1);
    //printList(l2);

    //printf("La memoria restante es 0? ");
    if (memSize(l1->curr->sig) == 0) {
        //printf("si. borrar.\n");
        erase(l1);
    }
    //else {
    //    printf("no\n");
    //}

}


/**
 * int memRemove
 ******
 * Retorna la posicion EN LA LISTA del bloque que comienza con el byte dado. Si
 * es 0, significa que no hay ninguno.
 ******
 * Input:
 *     int byte   : Byte inicial del bloque que se quiere liberar
 *     tLista  *L : Lista de bloques de memoria ocupada donde buscar
 ******
 * Return:
 *     int : Posicion de la lista 'L' donde esta el bloque que coincida con el
 *           byte de inicio en la lista de ocupados. Si es 0, significa que no
 *           esta.
 **/
int memRemove(int byte, tLista *L) {
    tNodo* aux = L->head->sig;

    for (int i = 1; i < length(L); i++) {
        if (byte == aux->bStrt) {
            return i;
        }
        aux = aux->sig;
    }

    return 0;
}

/**
 * void memFree
 ******
 * Libera un bloque de memoria ocupada, moviendo este bloque a la lista de memoria libre.
 * Teniendo en cuenta el byte de inicio, se inserta en orden.
 ******
 * Input:
 *     tLista *l1  : Lista de memoria disponible
 *     tLista *t2  : Lista de memoria ocupada
 *     int nBloque : Posicion en la lista 'l2' donde se encuentra el bloque a liberar.
 *     int tam     : Byte de inicio del bloque a liberar
 **/
void memFree(tLista *l1, tLista *l2, int nBloque, int byte) {
    //printf("Moviendo memoria...\n");
    moveToPos(l2,nBloque-1);

    moveToStart(l1);

    // Si la lista de disponibles esta vacia, entonces se agrega simplemente.
   //printf("Lista de disponibles estaba vacia? ");
    if (length(l1) == 1) {
        //printf("si. Append simple\n");
        append(l1, getbStrt(l2),getbEnd(l2));
        erase(l2);
        return;
    }
    //else {
    //    printf("no. Revisar orden.\n");
    //}

    // Nodos auxiliares, nos van a servir para comparar los bytes de inicio de los bloques.
    // De esa forma podemos ingresar bloques entremedio si es necesario.
    tNodo* aux1 = l1->head;
    tNodo* aux2 = l1->head->sig;


    for (int i = 0; i < length(l1); i++) {
        //printList(l1);
        //printList(l2);
        //printf("-><- Primer bloque: %i (Pos: %i) - Segundo bloque: %i (Pos: %i)\n",aux1->bStrt,i,aux2->bStrt,i+1);
        //printf("%i esta entremedio de %i y %i? ", byte, aux1->bStrt, aux2->bStrt);
        if (aux1->bStrt < byte && aux2->bStrt > byte) {
            //printf("si. Insertar aqui.\n");
            insert(l1,getbStrt(l2),getbEnd(l2));
            erase(l2);
            return;
        }
        //printf("no. Avanzar 1 bloque mas.\n");

        next(l1);
        aux1 = aux1->sig;
        aux2 = aux2->sig;
    }
}



/**
 * void memFusion
 ******
 * Fusiona dos bloques de memoria. El bloque que posee el cursor, y
 * el bloque siguiente al cursor.
 ******
 * Input:
 *     tLista *L : Lista de los bloques de memoria
 **/
void memFusion(tLista *L) {
    int finalCurr = getbEndOfCurr(L);
    int inicioNext = getbStrt(L);

    //printf("Es contiguo? ");
    if (inicioNext == (finalCurr + 1)) {
        //printf("Si. Juntar memoria del cursor con el siguiente.\n");

        // Hay que tomar el byte final del SIGUIENTE al cursor, y reemplazarlo en
        // el byte final DEL cursor.

        int finalNext = getbEnd(L);

        setbEndOfCurr(L, finalNext);

        erase(L);

    }
    //else {
    //    printf("no.\n");
    //}

}

/**
 * int memOcup
 ******
 * Retorna la canitdad de bytes totales de una lista de bloques de memoria
 ******
 * Input:
 *     tLista *L : Lista de bloques de memoria a revisar
 ******
 * Return:
 *     int : Cantidad de bytes totales de cada bloque de la lista
 **/
int memOcup(tLista *L) {
    int bytes = 0;
    tNodo *aux = L->head->sig;

    for (int i = 0; i < length(L)-1; i++) {
        bytes += memSize(aux);
        aux = aux->sig;
    }

    return bytes;
}

/**
 * int bloqOcup
 ******
 * Retorna la cantidad de bloques que hay en una lista.
 ******
 * Input:
 *     tLista *L : Lista de bloques de memoria a revisar
 ******
 * Return:
 *     int : Cantidad de bloques totales de la lista
 **/
int bloqOcup(tLista *L) {
    return length(L)-1;
}

// Retornar el tamaño del bloque siguiente al cursor.
int bloqSize(tLista *L) {
    return memSize(L->curr->sig);
}


// ====================
// CEMENTERIO
// ====================
//     _________________________
//    /                         -
//   |        RIP GRUPO          |
//   |                           |
//   |  "Intentaron implementar  |
//   |  memFusion y c murieron"  |
//   |                           |
//  ----------------------------------
//
// Si veo otro 'segmentation fault' haré la murision AAAAAAAAAAAAAAAAAAAAAAAAAA

// void memFusion(tLista *L) {
//     moveToStart(L);
//     tNodo *inicio = L->head;
//     tNodo *final = L->head->sig;

//     int temp, bInicialSig, bFinalActual;

//     for (int i = 0; i < length(L); i++) {
//         temp = final->bEnd;
//         bInicialSig = final->bStrt;
//         bFinalActual = inicio->bEnd;
//         printf("Temp: %i. InicialSig: %i. FinalAct: %i\n",temp,bInicialSig,bFinalActual);

//         printf("%i es igual a %i mas 1? ",bInicialSig,bFinalActual);
//         if (bInicialSig == bFinalActual + 1) {
//             printf("si. juntar memoria..\n");


//             inicio->bEnd = temp;

//             erase(L);

//             prev(L);


//         } else {
//             printf("no.\n");
//         }

//         next(L);
//         inicio = inicio->sig;

//     }
// }

// void memFusionBack(tLista *L) {
//     int temp = getbEnd(L);
//     int byteFinalCurr = getbEndOfCurr(L);
//     int byteInicial = getbStrt(L);

//     printf("Detras: Temp: %i. Final: %i. Inicial del otro: %i\n",temp,byteFinalCurr,byteInicial);

//     printf("%i es igual a %i mas 1? ",byteInicial,byteFinalCurr);
//     if (byteInicial == (byteFinalCurr + 1)) {

//         printf("si\n");

//         erase(L);

//         setbEndOfCurr(L, temp);
//     }else{
//         printf("no\n");
//     }
// }

// void memFusionForward(tLista *L) {
//     int temp = getbEndOfNext(L);
//     int byteFinal = getbEnd(L);
//     int byteInicialSig = getbStrtOfNext(L);

//     printf("Adelante: Temp: %i. Final: %i. Inicial del otro: %i\n",temp,byteFinal,byteInicialSig);

//     printf("%i es igual a %i mas 1? ",byteInicialSig,byteFinal);
//     if (byteInicialSig == (byteFinal + 1)) {

//         printf("si\n");

//         next(L);
//         erase(L);

//         setbEndOfCurr(L, temp);
//     }else{
//         printf("no\n");
//     }
// }

// void memFusion(tLista *L) {

//     int temp = L->curr->sig->bEnd;
//     int byteFinal = L->curr->bEnd;
//     int byteInicial = L->curr->sig->bStrt;

//     printf("Temp: %i. Final: %i. Inicial del otro: %i\n",temp,byteFinal,byteInicial);

//     printf("%i es igual a %i mas 1? ",byteInicial,byteFinal);
//     if (byteInicial == (byteFinal+1)) {
//         printf("si\n");
//         erase(L);
//         L->curr->bEnd = temp;
//     }
//     printf("no\n");
// }

// void memFusion(tLista *L) {

//     moveToStart(L);
//     tNodo* aux1 = L->head;
//     tNodo* aux2 = L->head->sig;

//     int temp, byteFinal, byteInicial;

//     for (int i = 0; i < length(L)-1; i++) {

//         temp = aux1->bStrt;
//         byteFinal = aux1->bEnd;
//         byteInicial = aux2->bStrt;

//         printf("Temp: %i. Final: %i. Inicial del otro: %i\n",temp,byteFinal,byteInicial);

//         printf("%i es igual a %i mas 1? ",byteInicial,byteFinal);
//         if (byteInicial == (byteFinal+1)) {
//             printf("si\n");
//             erase(L);
//             aux1->bStrt = temp;
//         }

//         printf("no\n");
//         next(L);
//         aux1 = aux1->sig;
//         aux2 = aux2->sig;
//     }
// }


// 'Taba de la perruna
