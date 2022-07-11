typedef struct {
    int ventaCliente; // Unidades de producto vendido por cliente
} tipoInfoVenta;

typedef struct {
    tipoClave clave; // Codigo del producto
    tipoInfoVenta info;
} ranuraVenta;

/**
 * ranuraVenta* hashInitVenta
 ******
 * Inicializa una tabla de hash para VENTAS con todas sus claves vacias
 ******
 * Input:
 *     ranuraVenta HT[] : Arreglo donde estara la informacion
 *     int hSize       : Tamaño de la tabla 'M'
 **/
ranuraVenta* hashInitVenta(int hSize) {
    tipoInfoVenta VALORINVALIDO;
    VALORINVALIDO.ventaCliente = 0;

    ranuraVenta *hashVenta = malloc(hSize*sizeof(ranuraVenta));

    for (int i = 0; i < hSize; i++){
        hashVenta[i].clave = VACIA;
        hashVenta[i].info = VALORINVALIDO;
    }

    return hashVenta;
}

/**
 * int regVenta
 ******
 * Registra una venta para un producto.
 ******
 * Input:
 *     ranuraVenta HT[] : Tabla de hash en donde insertar
 *     tipoClave k     : Clave del producto
 *     int M           : Tamaño de la tabla de hash (se usa para la funcion
 *                       de hashing)
 ******
 * Return:
 *     int : 0 = Repetido - 1 = Nuevo
 **/
int regVenta(ranuraVenta HT[], tipoClave k, int M) {
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
        HT[pos].info.ventaCliente ++;
        return 0; // inserción no exitosa: clave repetida
    } else {
        HT[pos].clave = k;
        HT[pos].info.ventaCliente ++;
        return 1; // inserción exitosa
    }
}

/**
 * void resetVenta
 ******
 * Reinicia la tabla de ventas a 0.
 ******
 * Input:
 *     ranuraVenta HT[] : Tabla de hashing de Ventas
 *     int M            : Tamaño de la tabla de hashing
 **/
void resetVenta(ranuraVenta HT[], int M) {
    for (int i = 0; i < M; i++) {
        HT[i].info.ventaCliente = 0;
    }
}


/**
 * void printArrVent
 ******
 * Imprime la tabla de hashing en bruto (es decir, el arreglo).
 * En un formato legible
 ******
 * Input:
 *     ranuraVenta HT[] : Tabla de hash de VENTAS a leer
 *     int M           : Tamaño de la tabla de hashing
 **/
void printArrVent(ranuraVenta HT[], int M) {
    printf("\n------Imprimiendo tabla VENTAS-------\n");
    printf("Ind\tCod\tVenta\n");
    for (int i = 0; i < M; i++) {
        printf("%i\t%i\t%i\n",
               i,HT[i].clave,HT[i].info.ventaCliente);
    }
    printf("------------------------\n");
}





/**
 * int getClaveArr
 ******
 * Retorna la clave asociada a un indice en la tabla de hashing
 ******
 * Input:
 *     ranuraVenta HT[] : Tabla de hashing
 *     int i           : Indice a revisar
 ******
 * Return:
 *     int : Clave asociada al indice 'i'
 **/
int getClaveArr(ranuraVenta HT[], int i) {
    return HT[i].clave;
}

/**
 * int getVentaClienteArr
 ******
 * Retorna la cantidad de vendidos asociada a un indice en la tabla de hashing
 ******
 * Input:
 *     ranuraVenta HT[] : Tabla de hashing
 *     int i           : Indice a revisar
 ******
 * Return:
 *     int : Cantidad de vendidos asociada al indice 'i'
 **/
int getVentaClienteArr(ranuraVenta HT[], int i) {
    return HT[i].info.ventaCliente;
}
