typedef struct {
    int codigo_producto;
    char nombre_producto[31];
    int precio;
} producto;

/**
 * int getCodigoProducto
 ******
 * Retorna el codigo de un producto
 ******
 * Input:
 *     producto *P : Producto a revisar
 ******
 * Return:
 *     int : Codigo del producto
 **/
int getCodigoProducto(producto *P) {
    return P->codigo_producto;
}

/**
 * char* getNombreProducto
 ******
 * Retorna el nombre de un producto
 ******
 * Input:
 *     producto *P : Producto a revisar
 ******
 * Return:
 *     char* : Nombre del producto
 **/
char* getNombreProducto(producto *P) {
    return P->nombre_producto;
}

/**
 * int getPrecioProducto
 ******
 * Retorna el precio de un producto
 ******
 * Input:
 *     producto *P : Producto a revisar
 ******
 * Return:
 *     int : Precio del producto
 **/
int getPrecioProducto(producto *P) {
    return P->precio;
}



typedef struct {
    int codigo_producto;
    int cantidad_descuento;
    int monto_descuento;
} oferta;

/**
 * int getCodigoOferta
 ******
 * Retorna el codigo del producto en oferta
 ******
 * Input:
 *     oferta *P : Oferta del producto a revisar
 ******
 * Return:
 *     int : Codigo del producto en oferta
 **/
int getCodigoOferta(oferta *P) {
    return P->codigo_producto;
}

/**
 * int getCantDesc
 ******
 * Retorna la cantidad de unidades necesarias de un
 * producto para aplicar el descuento
 ******
 * Input:
 *     oferta *P : Oferta del producto a revisar
 ******
 * Return:
 *     int : Cantidad de unidades necesarias para aplicar a oferta
 **/
int getCantDesc(oferta *P) {
    return P->cantidad_descuento;
}

/**
 * int getMontDesc
 ******
 * Retorna el monto de descuento del producto en oferta
 ******
 * Input:
 *     oferta *P : Oferta del producto a revisar
 ******
 * Return:
 *     int : Monto de descuento
 **/
int getMontDesc(oferta *P) {
    return P->monto_descuento;
}
