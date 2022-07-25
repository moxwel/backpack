/**
 * void actualizarSaldos
 ******
 * Realiza cambios en los saldos de las cuentas de banco segun las transacciones que se hayan hecho
 ******
 * Input:
 *     char* cfile : Archivo binario conteniendo los clientes
 *     char* tfile : Archivo ASCII conteniendo las transacciones
 **/
void actualizarSaldos(char *cfile, char *tfile) {

    FILE *fp = fopen(cfile,"r");
    if (fp == NULL) {
        printf("No se pudo abrir el archivo %s.\n",cfile);
        exit(1);
    }

    // Archivo que va a contener los cambios (temporal)
    FILE *fout = fopen("post-cliente.dat","w");

    clienteBanco buffer;

    // Leer 1 cliente a la vez
    int sal = 0;
    int cuentaActual = 0;
    while (fread(&buffer,sizeof(clienteBanco),1,fp) != 0) {
        sal = getSaldo(&buffer);
        cuentaActual = getCuenta(&buffer);

        // Ver las transacciones que afecten al cliente actual
        FILE *transacc_file = fopen(tfile,"r");
        if (transacc_file == NULL) {
            printf("No se pudo abrir el archivo %s.\n",tfile);
            exit(1);
        }

        char act; // Accion
        int corig, cdest, dinero; // Cuenta de origen, Cuenta de destino, Dinero transado
        while (fscanf(transacc_file, "%c ", &act) == 1) {
            cdest = 0;

            if (act == '+') {
                fscanf(transacc_file, "%i %i\n", &corig, &dinero);

                if (corig == cuentaActual) {
                    // printf("- La cuenta %i gana %i pesos -\n", cuentaActual,dinero);
                    sal += dinero;
                }

            } else if (act == '-') {
                fscanf(transacc_file, "%i %i\n", &corig, &dinero);

                if (corig == cuentaActual) {
                    // printf("- La cuenta %i pierde %i pesos -\n", cuentaActual,dinero);
                    sal -= dinero;
                }

            } else if (act == '>') {
                fscanf(transacc_file, "%i %i %i\n", &corig, &cdest, &dinero);

                if (corig == cuentaActual) {
                    // printf("- La cuenta %i pierde %i pesos -\n", cuentaActual,dinero);
                    sal -= dinero;
                } else if (cdest == cuentaActual) {
                    // printf("- La cuenta %i gana %i pesos -\n", cuentaActual,dinero);
                    sal += dinero;
                }
            }
        }

        fclose(transacc_file);

        // Actualizar el saldo final
        cambiarSaldo(&buffer, sal);

        fwrite(&buffer,sizeof(clienteBanco),1,fout);

        // printf("Cambio de cliente...\n");
    }

    fclose(fout);
    fclose(fp);

    file_replace(cfile, "post-cliente.dat");
}
