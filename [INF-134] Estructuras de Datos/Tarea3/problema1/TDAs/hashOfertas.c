typedef struct {
    int cantDesc; // Unidades de producto para aplicar descuento
    int descuento; // Dinero a descontar
} tipoInfoOfer;

typedef struct {
    tipoClave clave; // Codigo del producto
    tipoInfoOfer info;
} ranuraOfer;

/**
 * void hashInitOfer
 ******
 * Inicializa una tabla de hash para OFERTAS con todas sus claves vacias
 ******
 * Input:
 *     ranuraOfer HT[] : Arreglo donde estara la informacion
 *     int hSize       : Tamaño de la tabla 'M'
 **/
void hashInitOfer(ranuraOfer HT[], int hSize) {
    tipoInfoOfer VALORINVALIDO;
    VALORINVALIDO.cantDesc = 0;
    VALORINVALIDO.descuento = 0;

    for (int i = 0; i < hSize; i++){
        HT[i].clave = VACIA;
        HT[i].info = VALORINVALIDO;
    }
}

/**
 * int hashInsertOfer
 ******
 * Inserta un elemento en la tabla de hashing para OFERTAS
 ******
 * Input:
 *     ranuraOfer HT[] : Tabla de hash en donde insertar
 *     tipoClave k     : Clave a usar
 *     char* nomb      : Nombre del producto
 *     int M           : Tamaño de la tabla de hash (se usa para la funcion
 *                       de hashing)
 ******
 * Return:
 *     int : 0 = Fallo - 1 = Exito
 **/
int hashInsertOfer(ranuraOfer HT[], tipoClave k, int cant, int desc, int M) {
    int inicio, i;
    int pos = h(k,M);
    inicio = h(k,M);

    // printf("Resultado hash h: %i\n",h(k,M));

    for (i = 1; HT[pos].clave != VACIA && HT[pos].clave != k; i++){
        // printf("Colision!\n");
        pos = (inicio + p(k, i, M)) % M; // próxima ranura en la secuencia
        // printf("Resultado intento %i: %i\n",i,pos);
    }

    if (HT[pos].clave == k){
        return 0; // inserción no exitosa: clave repetida
    } else {
        HT[pos].clave = k;
        HT[pos].info.cantDesc = cant;
        HT[pos].info.descuento = desc;
        return 1; // inserción exitosa
    }
}

/**
 * tipoInfoProd hashSearchOfer
 ******
 * Busca un elemento en la tabla de hashing de OFERTAS
 ******
 * Input:
 *     ranuraProd HT[] : Tabla de hash en donde buscar
 *     tipoClave k     : Clave a buscar
 *     int M           : Tamaño de la tabla de hash (se usa para la funcion
 *                       de hashing)
 ******
 * Return:
 *     tipoInfoOfer : Informacion del producto dentro de la tabla
 **/
tipoInfoOfer hashSearchOfer(ranuraOfer HT[], tipoClave k, int M) {
    tipoInfoOfer VALORINVALIDO;
    VALORINVALIDO.cantDesc = 0;
    VALORINVALIDO.descuento = 0;

    int inicio, i;
    int pos = h(k,M);
    inicio = h(k,M);

    // printf("Resultado hash h: %i\n\n",h(k,M));

    // printf("Antes de for\nClave actual tabla: %i - Clave dada: %i\n\n",HT[pos].clave,k);

    for (i = 1; HT[pos].clave != VACIA && HT[pos].clave != k && i < 10000000; i++){
        // printf("No encontrado. Siguiente.\n");
        pos = (inicio + p(k, i, M)) % M; // próxima ranura en la secuencia
        // printf("Resultado intento %i: %i\n",i,pos);
        // printf("Clave actual tabla: %i - Clave dada: %i\n\n",HT[pos].clave,k);
    }

    // if (i == 10000000) {
    //     printf("\n============\nhashSearchOfer - Sistema Anti-Loop Infinito: BUSQUEDA ABORTADA\n============\n");
    // }

    if (HT[pos].clave == k){
        printf("Encontrado.\n");
        return HT[pos].info; // registro encontrado, búsqueda exitosa
    } else {
        printf("No Encontrado.\n");
        return VALORINVALIDO;
    }
}

/**
 * int claveInOfer
 ******
 * Retorna TRUE (1) si la clave dada esta dentro de la tabla de hashing.
 ******
 * Input:
 *     ranuraOfer HT[] : Tabla de hashing
 *     tipoClave k     : Clave a revisar
 *     int M           : Tamaño de la tabla de hashing
 ******
 * Return:
 *     int : TRUE 1 Encontrado - FALSE 0 No encontrado
 **/
int claveInOfer(ranuraOfer HT[], tipoClave k, int M) {
    int inicio, i;
    int pos = h(k,M);
    inicio = h(k,M);

    // printf("Resultado hash h: %i\n\n",h(k,M));

    // printf("Antes de for\nClave actual tabla: %i - Clave dada: %i\n\n",HT[pos].clave,k);

    for (i = 1; HT[pos].clave != VACIA && HT[pos].clave != k && i < 10000000; i++){
        // printf("No encontrado. Siguiente.\n");
        pos = (inicio + p(k, i, M)) % M; // próxima ranura en la secuencia
        // printf("Resultado intento %i: %i\n",i,pos);
        // printf("Clave actual tabla: %i - Clave dada: %i\n\n",HT[pos].clave,k);
    }

    // if (i == 10000000) {
    //     printf("\n============\nclaveInOfer - Sistema Anti-Loop Infinito: BUSQUEDA ABORTADA\n============\n");
    // }

    if (HT[pos].clave == k){
        // printf("Encontrado.\n");
        return 1; // registro encontrado, búsqueda exitosa
    } else {
        // printf("No Encontrado.\n");
        return 0;
    }
}







/**
 * void printArrOfer
 ******
 * Imprime la tabla de hashing en bruto (es decir, el arreglo).
 * En un formato legible
 ******
 * Input:
 *     ranuraOfer HT[] : Tabla de hash de ofertas a leer
 *     int M           : Tamaño de la tabla de hashing
 **/
void printArrOfer(ranuraOfer HT[], int M) {
    printf("\n------Imprimiendo tabla ofertas-------\n");
    printf("Ind\tCod\tCanDesc\tDesc\n");
    for (int i = 0; i < M; i++) {
        printf("%i\t%i\t%i\t%i\n",
               i,HT[i].clave,HT[i].info.cantDesc,HT[i].info.descuento);
    }
    printf("------------------------\n");
}



int getIndexArr(ranuraOfer HT[], tipoClave k, int M) {
    int VALORINVALIDO = -1;
    int inicio, i;
    int pos = h(k,M);
    inicio = h(k,M);

    // printf("Resultado hash h: %i\n\n",h(k,M));

    // printf("Antes de for\nClave actual tabla: %i - Clave dada: %i\n\n",HT[pos].clave,k);

    for (i = 1; HT[pos].clave != VACIA && HT[pos].clave != k && i < 10000000; i++){
        // printf("No encontrado. Siguiente.\n");
        pos = (inicio + p(k, i, M)) % M; // próxima ranura en la secuencia
        // printf("Resultado intento %i: %i\n",i,pos);
        // printf("Clave actual tabla: %i - Clave dada: %i\n\n",HT[pos].clave,k);
    }

    // if (i == 10000000) {
    //     printf("\n============\ngetIndexArr - Sistema Anti-Loop Infinito: BUSQUEDA ABORTADA\n============\n");
    // }

    if (HT[pos].clave == k){
        // printf("Encontrado.\n");
        return pos; // registro encontrado, búsqueda exitosa
    } else {
        // printf("No Encontrado.\n");
        return VALORINVALIDO;
    }
}



/**
 * int getCantDescArr
 ******
 * Retorna la cantidad de descuento asociada a un indice en la tabla de hashing
 ******
 * Input:
 *     ranuraOfer HT[] : Tabla de hashing
 *     int i           : Indice a revisar
 ******
 * Return:
 *     int : Cantidad de descuento asociada al indice 'i'
 **/
int getCantDescArr(ranuraOfer HT[], int i) {
    return HT[i].info.cantDesc;
}

/**
 * int getDescArr
 ******
 * Retorna el descuento asociado a un indice en la tabla de hashing
 ******
 * Input:
 *     ranuraOfer HT[] : Tabla de hashing
 *     int i           : Indice a revisar
 ******
 * Return:
 *     int : Descuento asociado al indice 'i'
 **/
int getDescArr(ranuraOfer HT[], int i) {
    return HT[i].info.descuento;
}
