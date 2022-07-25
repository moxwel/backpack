/**
 * char** buscar_str
 ******
 * Retorna un arreglo de strings, las cuales contienen
 * el prefijo deseado.
 ******
 * Input:
 *     char **s  : Arreglo de strings en donde se van a buscar prefijos.
 *     int n     : La cantidad de elementos del arreglo 's'.
 *     char *p   : String que contiene el prefijo que se va a buscar en 's'.
 *     int *nout : Cantidad de elementos del arreglo que retorna la funcion.
 *                 Se debe usar una variable externa en donde guardar el valor.
 ******
 * Returns:
 *     char** : Arreglo de 'nout' strings en heap que contiene los strings de 's' que si
 *              poseen el prefijo 'p' que se busca.
 **/
char** buscar_str(char **s, int n, char *p, int *nout) {
    int plen = strlen(p);
    // printf("Prefijo entrante: %s\nTamaño del prefijo: %i\n---\n",p,plen);

    // En 'cad' se va a guardar un string individual de hasta 200 caracteres (un buffer)
    char *cad = (char*)malloc(200 * sizeof(char));

    /* Un arreglo de strings que va a almacenar los strings que
     * contengan el prefijo, para despues retornarlo.
     */
    char **salida = (char**)malloc(n * sizeof(char*));

    /* Indice del arreglo 'salida', va a comenzar guardando
     * las palabras desde el indice 0. Mas adelante se muestra su accion.
     */
    int j = 0;

    for (int i = 0; i < n; i++) {
        /* La busqueda de strings dependen del tamaño del prefijo:
         * no vale la pena buscar los prefijos que tienen menor cantidad
         * de caracteres de lo que se busca. Tampoco vale la pena buscar
         * los prefijos que tienen mas letras, ya que suponiendo que
         * si poseen el prefijo, todos lo demas tambien lo tendran.
         */

        // Se copia el string del parametro a 'cad' para poder modificarlo en el heap
        strcpy(cad,s[i]);

        // printf("String: %s\n",cad);

        // Se introduce un caracter nulo en el string para "truncarlo" al tamaño de
        // caracteres del prefijo.
        cad[plen] = '\0';

        // printf("String: %s\n",cad);

        // printf("Comparar %s con %s... ",cad,p);

        if (strcmp(cad,p) == 0) {
            // printf("Son iguales. Estamos en la casilla %i.",i);

            /* En cada posicion del arreglo 'salida', se va a guardar un string
             * (arreglo de char) de hasta 200 caracteres.
             */
            salida[j] = (char*)malloc(200 * sizeof(char));

            strcpy(salida[j],s[i]);

            j++;
        }

        // printf("\n---\n");
    }

    free(cad);

    *nout = j;
    return salida;
}
