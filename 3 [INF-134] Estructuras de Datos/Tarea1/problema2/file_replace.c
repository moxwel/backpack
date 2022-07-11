/**
 * void file_replace
 ******
 * Reemplaza un archivo
 ******
 * Input:
 *     char* out : Archivo a ser reemplazado
 *     char* in : Archivo con los contenidos a reemplazar
 **/
void file_replace(char *out, char *in) {

    FILE *leer = fopen(in, "r");
    FILE *reemp = fopen(out, "w");

    clienteBanco temp;

    while (fread(&temp,sizeof(clienteBanco),1,leer) != 0) {
        fwrite(&temp,sizeof(clienteBanco),1,reemp);
    }

    fclose(leer);
    fclose(reemp);
    remove(in);
}
