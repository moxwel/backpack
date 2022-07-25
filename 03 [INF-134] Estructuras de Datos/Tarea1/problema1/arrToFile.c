/**
 * void arrToFile
 ******
 * Genera un archivo '.out' a partir de un arreglo de strings
 ******
 * Input:
 *     char **arr : Arreglo de strings a transformar.
 *     int n      : Cantidad de elementos del arreglo 'arr'.
 *     char *nomb : Nombre del archivo a generar.
 **/
void arrToFile(char **arr, int n, char *nomb) {
    char *name = (char*)malloc(204 * sizeof(char));

    // Se copia el parametro al heap para poder modificarlo y
    // añadir la extension '.out'.
    strcpy(name,nomb);
    strcat(name,".out");

    FILE *fp = fopen(name,"w");

    for (int i = 0; i < n; i++) {
        // printf("%i,%s\n",i,arr[i]);

        fprintf(fp, "%s\n", arr[i]);
    }

    free(name);
    fclose(fp);
}
