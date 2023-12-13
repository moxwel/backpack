#ifndef FUNCIONES_H
#define FUNCIONES_H

void quitarSaltoLinea(char* string);
char** generarMatrizSopa(char* file_path, int* dimm);
void liberarMatrizSopa(char** matriz, int dimm);
void extaerNombreArchivo(char* file_path, char* file_name);
int buscarPalabraSopa(char** sopa_letras, int dimm, char* palabra, int es_horizontal);
int buscarPalabraSopaNoPrint(char** sopa_letras, int dimm, char* palabra, int es_horizontal);
int buscarPalabraSopaLento(char** sopa_letras, int dimm, char* palabra, int es_horizontal);
int buscarPalabraSopaRapido(char** sopa_letras, int dimm, char* palabra, int es_horizontal);
int orientacionSopa(char* ruta_archivo);

#endif
