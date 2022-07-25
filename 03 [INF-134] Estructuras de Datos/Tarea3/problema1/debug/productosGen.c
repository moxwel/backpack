/* ESTE CODIGO C SOLO SE UTILIZO PARA FINES DE DEPURACION Y TESTEO DEL PROGRAMA
 *
 * SABEMOS QUE ACCEDIMOS DE FORMA "MALULA" AL TDA, NO NOS MATE PLS
 *
 * *****************************************
 * *          ESTO ES UN BORRADOR          *
 * *****************************************
 *
 */

#include <stdio.h>
#include <string.h>

typedef struct {
    int codigo_producto;
    char nombre_producto[31];
    int precio;
} producto;

int main() {

    FILE *fp = fopen("productos.dat","w");

    int cantidad_productos = 7;

    producto p1,p2,p3,p4,p5,p6,p7;

    p1.codigo_producto = 5809;
    strcpy(p1.nombre_producto,"Leche");
    p1.precio = 500;

    p2.codigo_producto = 6758;
    strcpy(p2.nombre_producto,"Pan");
    p2.precio = 300;

    p3.codigo_producto = 2645;
    strcpy(p3.nombre_producto,"Huevo");
    p3.precio = 150;

    p4.codigo_producto = 9999;
    strcpy(p4.nombre_producto,"Galletas");
    p4.precio = 1000;

    p5.codigo_producto = 666;
    strcpy(p5.nombre_producto,"Sopa");
    p5.precio = 1500;

    p6.codigo_producto = 1234;
    strcpy(p6.nombre_producto,"Super 8");
    p6.precio = 50;

    p7.codigo_producto = 8547;
    strcpy(p7.nombre_producto,"Crema");
    p7.precio = 2000;

    fwrite(&cantidad_productos, sizeof(int), 1, fp);

    fwrite(&p1, sizeof(producto), 1, fp);
    fwrite(&p2, sizeof(producto), 1, fp);
    fwrite(&p3, sizeof(producto), 1, fp);
    fwrite(&p4, sizeof(producto), 1, fp);
    fwrite(&p5, sizeof(producto), 1, fp);
    fwrite(&p6, sizeof(producto), 1, fp);
    fwrite(&p7, sizeof(producto), 1, fp);

    fclose(fp);

    return 0;
}
