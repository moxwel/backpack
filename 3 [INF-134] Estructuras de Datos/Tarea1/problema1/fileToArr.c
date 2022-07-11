/**
 * char** fileToArr
 ******
 * Transforma un archivo a un arreglo de strings.
 ******
 * Input:
 *     char *nombre : String del nombre del archivo a leer.
 *     int *nout    : Numero de elementos que tendra el arreglo retornado.
 *                    Se debe usar una variable externa en donde guardar el valor.
 ******
 * Returns:
 *     char** : Arreglo de 'nout' strings en heap equivalente a la cantidad de
 *              strings que hay en el archivo 'file' por cada linea.
 **/
char** fileToArr (char *file, int *nout) {
    FILE *fp = fopen(file,"r");
    if (fp == NULL) {
        printf("- Error al abrir archivo %s -\n",file);
        exit(1);
    }

    // Aqui se guardara el arreglo de strings, que despues se va a retornar
    char **arr = (char**)malloc(1000000 * sizeof(char*));

    // En este buffer se guardara de forma temporal el string que se va a
    // obtener al leer el archivo.
    char buf[201];

    // El indice del arreglo 'arr', comenzara desde 0. Al final se utilizara
    // para poder devolver el tamaño del arreglo de strings.
    int i = 0;

    while (fgets(buf,200,fp) != NULL) {
        // printf("%i,%s|",(int)strlen(buf),buf);

        arr[i] = (char*)malloc(200 * sizeof(char));
        strcpy(arr[i],buf);

        /* Al obtener strings desde un archivo, estos van a venir incluidos
         * con el caracter de salto de linea '\n' (ASCII 10), entonces, para
         * sanitizar los strings, se revisa el ultimo caracter de este.
         * Si es un caracter de salto de linea, se reemplazara con un
         * caracter nulo '\0'.
         */

        // printf("\nUltimo caracter de indice %i: %i, %c\n",i,arr[i][strlen(buf)-1],arr[i][strlen(buf)-1]);

        if (arr[i][strlen(buf)-1] == '\n') {
            // printf("Tiene salto de linea\n");
            arr[i][strlen(buf)-1] = '\0';
        }

        i++;
    }

    fclose(fp);

    *nout = i;
    return arr;
}
