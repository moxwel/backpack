/**
 * ranuraProd* leerProductos
 ******
 * Lee los productos disponibles de un supermercado desde un archivo,
 * y retorna una tabla de hashing con todos los productos de este.
 *
 * Las llaves de la tabla de hashing son los codigos de cada producto.
 ******
 * Input:
 *     char* fileName : Nombre del archivo a leer
 *     int *N         : Tamaño de la tabla de hashing (usado como return).
 ******
 * Return:
 *     ranuraProd* : Tabla de hashing con los productos
 *     int* N      : Tamaño de la tabla de hashing (se requiere para buscar en la tabla)
 **/
ranuraProd* leerProductos(char* fileName, int *N){

    // printf("Leyendo productos.dat ...\n\n");

    FILE *fprod = fopen(fileName,"r");

    // Solucion error archivo inexistente
    if (fprod == NULL) {
        printf("Error, no existe el archivo '%s'\n",fileName);
        exit(1);
    }



    int total_productos;
    fread(&total_productos, sizeof(int),1,fprod);

    int hashSizeProd = total_productos/0.7;
    // printf("Cant productos: %i - Tamaño hash: %i\n",total_productos, hashSizeProd);


    // printf("\n==========\nInicializando tabla...\n\n");

    ranuraProd *hashProducto = malloc(hashSizeProd * sizeof(ranuraProd));

    hashInitProd(hashProducto, hashSizeProd);
    // for (int i = 0; i < hashSizeProd; i++) {
    //     printf("Indice: %i - Clave: %i\n",i,hashProducto[i].clave);
    //     printf("Info: Nombre: %s - Precio: %i\n\n",hashProducto[i].info.nombre,hashProducto[i].info.precio);
    // }

    // printf("==========\nLeyendo productos...\n\n");

    producto leerP;
    for(int i = 0; i < total_productos; i++) {
        fread(&leerP,sizeof(producto),1,fprod);
        // printf("Codigo: %i - Nombre: %s - Precio: %i\n",getCodigoProducto(&leerP),getNombreProducto(&leerP),getPrecioProducto(&leerP));

        hashInsertProd(hashProducto, getCodigoProducto(&leerP), getNombreProducto(&leerP), getPrecioProducto(&leerP), hashSizeProd);

        // printf("\n");
    }

    // printf("Creacion de hashProducto terminado.\n==========\n");

    fclose(fprod);

    // printf("Tabla final productos:\n\n");
    // for (int i = 0; i < hashSizeProd; i++) {
    //     printf("Indice: %i - Clave: %i\n",i,hashProducto[i].clave);
    //     printf("Info: Nombre: %s - Precio: %i\n\n",hashProducto[i].info.nombre,hashProducto[i].info.precio);
    // }
    // printf("================\n");

    *N = hashSizeProd;
    return hashProducto;
}



/**
 * ranuraOfer* leerOfertas
 ******
 * Lee las ofertas disponibles de un supermercado desde un archivo,
 * y retorna una tabla de hashing con todos las ofertas de este.
 *
 * Las llaves de la tabla de hashing son los codigos de cada producto.
 ******
 * Input:
 *     char* fileName : Nombre del archivo a leer
 *     int *N         : Tamaño de la tabla de hashing (usado como return).
 ******
 * Return:
 *     ranuraProd* : Tabla de hashing con las ofertas
 *     int* N      : Tamaño de la tabla de hashing (se requiere para buscar en la tabla)
 **/
ranuraOfer* leerOfertas(char* fileName, int *N) {

    // printf("Leyendo ofertas.dat ...\n\n");

    FILE *fofer = fopen(fileName,"r");

    // Solucion error archivo inexistente
    if (fofer == NULL) {
        printf("Error, no existe el archivo '%s'\n",fileName);
        exit(1);
    }




    int total_ofertas;
    fread(&total_ofertas, sizeof(int),1,fofer);

    int hashSizeOfer = total_ofertas/0.7;
    // printf("Cant ofertas: %i - Tamaño hash: %i\n",total_ofertas, hashSizeOfer);


    // printf("\n==========\nInicializando tabla...\n\n");

    ranuraOfer *hashOferta = malloc(hashSizeOfer * sizeof(ranuraOfer));

    hashInitOfer(hashOferta, hashSizeOfer);
    // for (int i = 0; i < hashSizeOfer; i++) {
    //     printf("Indice: %i - Clave: %i\n",i,hashOferta[i].clave);
    //     printf("Info: Descuento: %i - Cantidad para aplicar: %i\n\n",hashOferta[i].info.descuento,hashOferta[i].info.cantDesc);
    // }

    // printf("==========\nLeyendo ofertas...\n\n");

    oferta leerO;
    for(int i = 0; i < total_ofertas; i++) {
        fread(&leerO,sizeof(oferta),1,fofer);
        // printf("Codigo: %i - Cantidad para aplicar: %i - Descuento: %i\n",getCodigoOferta(&leerO),getCantDesc(&leerO),getMontDesc(&leerO));

        hashInsertOfer(hashOferta, getCodigoOferta(&leerO), getCantDesc(&leerO), getMontDesc(&leerO), hashSizeOfer);

        // printf("\n");
    }

    // printf("Creacion de hashOferta terminado.\n==========\n");

    fclose(fofer);

    // printf("Tabla final ofertas:\n\n");
    // for (int i = 0; i < hashSizeOfer; i++) {
    //     printf("Indice: %i - Clave: %i\n",i,hashOferta[i].clave);
    //     printf("Info: Descuento: %i - Cantidad para aplicar: %i\n\n",hashOferta[i].info.descuento,hashOferta[i].info.cantDesc);
    // }
    // printf("================\n");

    *N = hashSizeOfer;
    return hashOferta;
}




/**
 * void leerCompras
 ******
 * Modifica la tabla de hash de productos ingresando las compras y las ventas de cada producto.
 * Todo eso, dependiendo de un archivo donde se mantiene registro de las compras.
 ******
 * Input:
 *     char* fileName   : Nombre del archivo de compras
 *     ranuraProd HTP[] : Tabla de hashing de productos
 *     int M_HTP        : Tamaño de la tabla de hashing de productos
 *     ranuraOfer HTO[] : Tabla de hashing de ofertas
 *     int M_HTO        : Tamaño de la tabla de hashing de ofertas
 *     int *r           : Variable externa donde guardar la cantidad de productos
 *                        a incluir en el ranking
 **/
void leerCompras(char* fileName, ranuraProd HTP[], int M_HTP, ranuraOfer HTO[], int M_HTO, int *r) {
    // printf("\n===========\nLeyendo compras.txt\n\n");

    FILE *fcomp = fopen(fileName,"r");

    // Solucion error archivo inexistente
    if (fcomp == NULL) {
        printf("Error, no existe el archivo '%s'\n",fileName);
        exit(1);
    }

    int nRank,nClientes;
    fscanf(fcomp,"%i",&nRank);
    fscanf(fcomp,"%i",&nClientes);

    *r = nRank;

    // printf("Cantidad a incluir en ranking: %i\n",nRank);
    // printf("Cantidad de clientes a atender: %i\n\n",nClientes);

    ranuraVenta *HTC = hashInitVenta(M_HTP);

    for (int i = 1; i <= nClientes; i++) {
        // printf("---CLIENTE %i---\n",i);

        int nProductos = 0;
        fscanf(fcomp,"%i",&nProductos);

        // printf("Cantidad de productos a comprar: %i\n",nProductos);

        for (int j = 0; j < nProductos; j++) {
            int codigoProducto;
            fscanf(fcomp,"%i",&codigoProducto);

            // printf("- Comprar producto: %i\n",codigoProducto);

            if (venderProd(HTP, codigoProducto, M_HTP)) {
                regVenta(HTC,codigoProducto,M_HTP);
            }

        }

        // printArrVent(HTC,M_HTP);

        // printf("Revisando ofertas de los productos comprados:\n\n");
        int producto, cTotVentaC, precioProd,ventaTotal;
        for (int i = 0; i < M_HTP; i++) {
            producto = getClaveArr(HTC,i);
            cTotVentaC = getVentaClienteArr(HTC,i);
            precioProd = getPrecioArr(HTP,i);
            ventaTotal = cTotVentaC*precioProd;

            // printf("Indice: %i - Codigo: %i - Precio: %i\n",i,producto,precioProd);

            // printf("Total vendidos: %i - Ganancia total: %i\n",cTotVentaC,ventaTotal);

            // Tiene que al menos haber comprado 1 para aplicar a la oferta
            if (producto != VACIA && cTotVentaC > 0 && claveInOfer(HTO,producto,M_HTO)) {
                // printf("SI Tiene oferta. Aplicando descuento...\n");

                // printf("Indice de producto en oferta: %i\n",getIndexArr(HTO, producto, M_HTO));

                int cantidadParaDesc = getCantDescArr(HTO,getIndexArr(HTO, producto, M_HTO));
                int montoDesc = getDescArr(HTO,getIndexArr(HTO, producto, M_HTO));

                // printf("Cantidad para aplicar: %i - Monto a descontar: %i\n",cantidadParaDesc,montoDesc);

                // printf("Descontando %i pesos...\n",(cTotVentaC / cantidadParaDesc)*montoDesc);

                ventaTotal = ventaTotal - (cTotVentaC / cantidadParaDesc)*montoDesc;

            }
            // else {
            //     printf("NO tiene oferta\n");
            // }

            // printf("Ganancia final: %i\n",ventaTotal);

            setTotVenArr(HTP,i,ventaTotal);

            // printf("\n\n");

        }

        resetVenta(HTC,M_HTP);
        // printf("---------------\n\n");
    }

    fclose(fcomp);
    free(HTC);
}



/**
 * void rankGen
 ******
 * Genera un archivo "ranking.txt" de una cierta cantidad de productos ordenados
 * segun la ganancia monetaria de cada uno.
 ******
 * Input:
 *     int cantRanking  : Cantidad de productos a incluir en el ranking
 *     ranuraProd HTP[] : Tabla de productos a ver
 *     int M_HTP        : Tamaño de la tabla de hashing de productos
 **/
void rankGen(int cantRanking, ranuraProd HTP[], int M_HTP) {
    char** arrRank = malloc(cantRanking*sizeof(char*));

    FILE *fr = fopen("ranking.txt","w");

    for (int i = 0; i < cantRanking; i++) {

        int mayorGanancia = 0;
        int codigoMayor = 0;
        int indiceMayor = 0;
        int unidadesMayor = 0;
        char nombProd[31] = "---";
        for (int j = 0; j < M_HTP; j++) {


            // printf("%i es mayor a %i?\n",getTotVentaArr(HTP,j),mayorGanancia);
            if (getTotVentaArr(HTP,j) > mayorGanancia) {
                // printf("Si, ahora %i es el mayor.\n",getTotVentaArr(HTP,j));

                indiceMayor = j;
                mayorGanancia = getTotVentaArr(HTP,j);
                codigoMayor = getCodProdArr(HTP,j);
                unidadesMayor = getCantVentaArr(HTP,j);
                strcpy(nombProd,getNombProdArr(HTP,j));

            }

        }

        resetTotVenArr(HTP,indiceMayor);

        // printf("Puesto %i, %i %s %i %i\n==========\n",i+1, codigoMayor, nombProd, unidadesMayor, mayorGanancia);


        fprintf(fr,"%i %s %i %i\n", codigoMayor, nombProd, unidadesMayor, mayorGanancia);

    }

    free(arrRank);
    fclose(fr);
}
