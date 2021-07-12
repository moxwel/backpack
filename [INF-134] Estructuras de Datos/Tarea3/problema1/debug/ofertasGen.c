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
    int cantidad_descuento;
    int monto_descuento;
} oferta;

int main() {

    FILE *fp = fopen("ofertas.dat","w");

    int cantidad_ofertas = 3;

    oferta p1,p2,p3;

    // Leche
    p1.codigo_producto = 5809;
    p1.cantidad_descuento = 3;
    p1.monto_descuento = 300;

    // Huevo
    p2.codigo_producto = 2645;
    p2.cantidad_descuento = 5;
    p2.monto_descuento = 200;

    // Sopa
    p3.codigo_producto = 666;
    p3.cantidad_descuento = 2;
    p3.monto_descuento = 500;

    fwrite(&cantidad_ofertas, sizeof(int), 1, fp);

    fwrite(&p1, sizeof(oferta), 1, fp);
    fwrite(&p2, sizeof(oferta), 1, fp);
    fwrite(&p3, sizeof(oferta), 1, fp);

    fclose(fp);

    return 0;
}
