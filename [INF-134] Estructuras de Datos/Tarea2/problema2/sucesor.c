int maxValueAux(tNodoArbolBin *nodo,int temp);

/**
 * int maxValue
 ******
 * Permite encontrar el mayor elemento que se encuentra
 * en el arbol binario.
 ******
 * Input:
 *     tABB *T : Puntero a un arbol binario clasico
 ******
 * Return:
 *     int : El maximo valor que se encuentra en 'T'
 **/
int maxValue(tABB *T) {

    return maxValueAux(T->raiz,0);
}
int maxValueAux(tNodoArbolBin *nodo,int temp) {
    if (nodo == NULL) {
        return temp;
    }
    temp = nodo->info;
    return maxValueAux(nodo->der,temp);
}



/**
 * int sucesor
 ******
 * Dado un numero cualquiera y un numero que represente un maximo,
 * retorna el siguiente elemento mas grande que se encuentre
 * en un arbol binario. Si no, retorna el numero maximo.
 ******
 * Input:
 *     tABB *T   : Puntero al arbol binario donde buscar
 *      int item : Numero con el cual se comienza a buscar su sucesor en 'T'
 *      int max  : Numero maximo
 ******
 * Return:
 *     int : El numero que sigue a 'item', que ademas se encuentra en 'T'
 **/
int sucesor(tABB *T, int item, int max) {

    if (item > maxValue(T)) {
        return max;
    }

    for (int i = item + 1; i <= maxValue(T);i++) {
        if (find(T,i)) {
            return i;
        }
    }

    return max;
}

/**
 * void sucesorFile
 ******
 * Funcion analoga a sucesor(), pero su salida se escribe en un archivo.
 ******
 * Input:
 *     tABB *T     : Puntero al arbol binario donde buscar
 *      int item   : Numero con el cual se comienza a buscar su sucesor en 'T'
 *      int max    : Numero maximo
 *      FILE *file : Archivo donde se escribirá la salida
 **/
void sucesorFile(tABB *T, int item, int max, FILE* file) {

    if (item > maxValue(T)) {
        fprintf(file,"%i\n",max);
        return;
    }

    for (int i = item + 1; i <= maxValue(T);i++) {
        if (find(T,i)) {
            fprintf(file,"%i\n",i);
            return;
        }
    }

    fprintf(file,"%i\n",max);
    return;
}
