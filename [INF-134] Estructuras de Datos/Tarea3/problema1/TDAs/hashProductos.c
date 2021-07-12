typedef struct {
    char nombre[31]; // Nombre del producto
    int precio; // Precio del producto
    int totalVenta; // Dinero total recaudado
    int cantVendidos; // Unidades de producto vendido
} tipoInfoProd;

typedef struct {
    tipoClave clave; // Codigo del producto
    tipoInfoProd info;
} ranuraProd;

/**
 * void hashInitProd
 ******
 * Inicializa una tabla de hash para PRODUCTOS con todas sus claves vacias
 ******
 * Input:
 *     ranuraprod HT[] : Arreglo donde estara la informacion
 *     int hSize       : Tamaño de la tabla 'M'
 **/
void hashInitProd(ranuraProd HT[], int hSize) {
    tipoInfoProd VALORINVALIDO;
    strcpy(VALORINVALIDO.nombre,"");
    VALORINVALIDO.precio = 0;
    VALORINVALIDO.cantVendidos = 0;
    VALORINVALIDO.totalVenta = 0;

    for (int i = 0; i < hSize; i++){
        HT[i].clave = VACIA;
        HT[i].info = VALORINVALIDO;
    }
}

/**
 * int hashInsertProd
 ******
 * Inserta un elemento en la tabla de hashing para PRODUCTOS
 ******
 * Input:
 *     ranuraProd HT[] : Tabla de hash en donde insertar
 *     tipoClave k     : Clave a usar
 *     char* nomb      : Nombre del producto
 *     int money       : Precio del producto
 *     int M           : Tamaño de la tabla de hash (se usa para la funcion
 *                       de hashing)
 ******
 * Return:
 *     int : 0 = Fallo - 1 = Exito
 **/
int hashInsertProd(ranuraProd HT[], tipoClave k, char* nomb, int money, int M) {
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
        strcpy(HT[pos].info.nombre,nomb);
        HT[pos].info.precio = money;
        HT[pos].info.cantVendidos = 0;
        HT[pos].info.totalVenta = 0;
        return 1; // inserción exitosa
    }
}

/**
 * tipoInfoProd hashSearchProd
 ******
 * Busca un elemento en la tabla de hashing de PRODUCTOS
 ******
 * Input:
 *     ranuraProd HT[] : Tabla de hash en donde buscar
 *     tipoClave k     : Clave a buscar
 *     int M           : Tamaño de la tabla de hash (se usa para la funcion
 *                       de hashing)
 ******
 * Return:
 *     tipoInfoProd : Informacion del producto dentro de la tabla
 **/
tipoInfoProd hashSearchProd(ranuraProd HT[], tipoClave k, int M) {
    tipoInfoProd VALORINVALIDO;
    strcpy(VALORINVALIDO.nombre,"");
    VALORINVALIDO.cantVendidos = 0;
    VALORINVALIDO.totalVenta = 0;
    VALORINVALIDO.precio = 0;

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
    //     printf("\n============\nhashSearchProd - Sistema Anti-Loop Infinito: BUSQUEDA ABORTADA\n============\n");
    // }

    if (HT[pos].clave == k){
        // printf("Encontrado.\n");
        return HT[pos].info; // registro encontrado, búsqueda exitosa
    } else {
        // printf("No Encontrado.\n");
        return VALORINVALIDO;
    }
}



/**
 * int venderProd
 ******
 * Incrementa en 1 la celda 'cantVendidos' de un producto
 * dentro de una tabla de hashing de productos.
 ******
 * Input:
 *     ranuraProd HT[] : Tabla de hashing de productos
 *     tipoClave k     : Codigo del producto a vender
 *     int M           : Tamaño de la tabla de hashing
 ******
 * Returns:
 *     int : 1 Exito - 0 Fallo
 **/
int venderProd(ranuraProd HT[], tipoClave k, int M) {
    int inicio, i;
    int pos = h(k,M);
    inicio = h(k,M);

    for (i = 1; HT[pos].clave != VACIA && HT[pos].clave != k && i < 10000000; i++){
        // printf("No encontrado. Siguiente.\n");
        pos = (inicio + p(k, i, M)) % M; // próxima ranura en la secuencia
        // printf("Resultado intento %i: %i\n",i,pos);
        // printf("Clave actual tabla: %i - Clave dada: %i\n\n",HT[pos].clave,k);
    }

    // if (i == 10000000) {
    //     printf("\n============\nvenderProd - Sistema Anti-Loop Infinito: BUSQUEDA ABORTADA\n============\n");
    // }

    if (HT[pos].clave == k){
        // printf("Encontrado.\n");
        HT[pos].info.cantVendidos++; // registro encontrado, incrementando vendidos
        return 1;
    } else {
        // printf("No Encontrado.\n");
        return 0;
    }

}






/**
 * void printArrProd
 ******
 * Imprime la tabla de hashing en bruto (es decir, el arreglo).
 * En un formato legible
 ******
 * Input:
 *     ranuraProd HT[] : Tabla de hash de productos a leer
 *     int M           : Tamaño de la tabla de hashing
 **/
void printArrProd(ranuraProd HT[], int M) {
    printf("\n------Imprimiendo tabla productos-------\n");
    printf("Ind\tCod\t%-20s\tPrec\tTotVen\tCantVen\n","Nombre");
    for (int i = 0; i < M; i++) {
        printf("%i\t%i\t%-20s\t%i\t%i\t%i\n",
               i,HT[i].clave,HT[i].info.nombre,HT[i].info.precio,HT[i].info.totalVenta,HT[i].info.cantVendidos);
    }
    printf("------------------------\n");
}



/**
 * int getPrecioArr
 ******
 * Retorna el precio asociado a un indice en la tabla de hashing
 ******
 * Input:
 *     ranuraProd HT[] : Tabla de hashing
 *     int i           : Indice a revisar
 ******
 * Return:
 *     int : Precio asociado al indice 'i'
 **/
int getPrecioArr(ranuraProd HT[], int i) {
    return HT[i].info.precio;
}

/**
 * int getTotVentaArr
 ******
 * Retorna el Total de venta asociado a un indice en la tabla de hashing
 ******
 * Input:
 *     ranuraProd HT[] : Tabla de hashing
 *     int i           : Indice a revisar
 ******
 * Return:
 *     int : Total de venta asociado al indice 'i'
 **/
int getTotVentaArr(ranuraProd HT[], int i) {
    return HT[i].info.totalVenta;
}

/**
 * int getCantVentaArr
 ******
 * Retorna la Cantidad de venta asociado a un indice en la tabla de hashing
 ******
 * Input:
 *     ranuraProd HT[] : Tabla de hashing
 *     int i           : Indice a revisar
 ******
 * Return:
 *     int : Cantidad de venta asociado al indice 'i'
 **/
int getCantVentaArr(ranuraProd HT[], int i) {
    return HT[i].info.cantVendidos;
}

/**
 * int getCodProdArr
 ******
 * Retorna el Codigo del producto asociado a un indice en la tabla de hashing
 ******
 * Input:
 *     ranuraProd HT[] : Tabla de hashing
 *     int i           : Indice a revisar
 ******
 * Return:
 *     int : Codigo del producto asociado al indice 'i'
 **/
int getCodProdArr(ranuraProd HT[], int i) {
    return HT[i].clave;
}

/**
 * char* getNombProdArr
 ******
 * Retorna el Nombre del producto asociado a un indice en la tabla de hashing
 ******
 * Input:
 *     ranuraProd HT[] : Tabla de hashing
 *     int i           : Indice a revisar
 ******
 * Return:
 *     char* : Nombre del producto asociado al indice 'i'
 **/
char* getNombProdArr(ranuraProd HT[], int i) {
    return HT[i].info.nombre;
}

/**
 * void setTotVenArr
 ******
 * Incrementa la ganancia total de un producto asociado
 * a un indice en la tabla de hashing.
 ******
 * Input:
 *     ranuraProd HT[] : Tabla de hashing
 *     int i           : Indice a revisar
 *     int v           : Total de venta a incrementar
 **/
void setTotVenArr(ranuraProd HT[], int i, int v) {
    HT[i].info.totalVenta += v;
}

/**
 * void resetTotVenArr
 ******
 * Resetea la ganancia total de un producto asociado
 * a un indice en la tabla de hashing a 0.
 ******
 * Input:
 *     ranuraProd HT[] : Tabla de hashing
 *     int i           : Indice a revisar
 **/
void resetTotVenArr(ranuraProd HT[], int i) {
    HT[i].info.totalVenta = 0;
}
