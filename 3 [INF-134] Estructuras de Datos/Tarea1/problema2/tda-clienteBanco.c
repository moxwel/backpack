typedef struct {
    int nroCuenta;
    int saldo;
    char nbre[51];
    char direccion[51];
} clienteBanco;

int getSaldo(clienteBanco *C) {
    return C->saldo;
}

int getCuenta(clienteBanco *C) {
    return C->nroCuenta;
}

char* getNombre(clienteBanco *C) {
    return C->nbre;
}

char* getDireccion(clienteBanco *C) {
    return C->direccion;
}

void cambiarSaldo(clienteBanco *C, int S) {
    C->saldo = S;
}
