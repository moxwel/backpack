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
    int nroCuenta;
    int saldo;
    char nbre[51];
    char direccion[51];
} clienteBanco;

int main() {
    FILE *fp = fopen("clientes.dat","w");

    clienteBanco cris, yuyo, max;

    cris.nroCuenta = 1;
    cris.saldo = 100;
    strcpy(cris.nbre,"Cristopher");
    strcpy(cris.direccion, "Tu corazon");

    yuyo.nroCuenta = 2;
    yuyo.saldo = 100;
    strcpy(yuyo.nbre,"Giuliana");
    strcpy(yuyo.direccion, "No se");

    max.nroCuenta = 3;
    max.saldo = 100;
    strcpy(max.nbre,"Maximiliano");
    strcpy(max.direccion, "En la durma");

    fwrite(&cris,sizeof(clienteBanco),1,fp);
    fwrite(&yuyo,sizeof(clienteBanco),1,fp);
    fwrite(&max,sizeof(clienteBanco),1,fp);

    fclose(fp);

    return 0;
}
