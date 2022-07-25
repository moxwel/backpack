void preOrdenFileAux(tNodoArbolBin *nodo, FILE *file);

/**
 * void preOrdenFile
 ******
 * Recorre un arbol binario en pre-orden, y su salida
 * se escribe a un archivo
 ******
 * Input:
 *     tABB *T    : Arbol binario donde recorrer en pre-orden
 *     FILE *file : Archivo donde se escribirá la salida
 **/
void preOrdenFile(tABB *T, FILE* file) {
    preOrdenFileAux(T->raiz, file);
    fprintf(file,"\n");
}
void preOrdenFileAux(tNodoArbolBin *nodo, FILE *file) {
    if (nodo == NULL) {
        return;
    }
    fprintf(file,"%i ",getValue(nodo));
    preOrdenFileAux(nodo->izq,file);
    preOrdenFileAux(nodo->der,file);
}
