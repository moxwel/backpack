/**
 * void printArrStr
 ******
 * Imprime un arreglo de strings en una tabla.
 ******
 * Input:
 *     char **arr : Arreglo de strings a leerse.
 *     int n      : Cantidad de elementos del arreglo 'arr'.
 **/
void printArrStr(char **arr, int n) {
    printf("==========\nInd:\tLen:\tVal:\n");
    for (int i = 0; i < n; i++) {
        printf("%i\t%i\t%s\n",i,(int)strlen(arr[i]),arr[i]);
    }
    printf("==========\n");
}
