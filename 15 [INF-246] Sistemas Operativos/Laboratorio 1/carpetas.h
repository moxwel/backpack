#ifndef CARPETAS_H
#define CARPETAS_H

char** arregloArchivos(char* directorio, int* n_archivos);
void liberarArregloArchivos(char** arreglo, int n_archivos);
void crearCarpetas();
void moverArchivo(char* archivo, int es_horizontal, int dimm);

#endif
