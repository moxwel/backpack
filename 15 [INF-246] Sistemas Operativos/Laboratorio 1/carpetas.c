#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <dirent.h>
#include <sys/types.h>
#include <sys/stat.h>

/**
 * char** arregloArchivos
 *
 * Recibe la direccion de un directorio y retorna un arreglo con las direcciones
 * de todos los archivos que se encuentran en el directorio.
 *
 * Input:
 *     - char* directorio : Direccion del directorio.
 *     - int* n_archivos : Puntero a entero donde se guardará el numero de archivos
 *
 * Return:
 *     - char** : Puntero a un arreglo de strings con las direcciones de los archivos
*/
char** arregloArchivos(char* directorio, int* n_archivos) {
    DIR* dir_stream = opendir(directorio);

    if (dir_stream == NULL) {
        printf("Error al abrir el directorio\n");
        return NULL;
    }

    struct dirent* dir_entry;

    // Crear arreglo de nombres de archivos
    char** lista_archivos = (char**)malloc(15 * sizeof(char*));

    int i = 0;
    while ((dir_entry = readdir(dir_stream)) != NULL) {
        if (strcmp(dir_entry->d_name, ".") != 0 && strcmp(dir_entry->d_name, "..") != 0) {
            // Se guarda el nombre del archivo en el arreglo. Se asume que el nombre del archivo no supera los 250 caracteres
            lista_archivos[i] = (char*)malloc(250 * sizeof(char));
            char file_path[250];
            strcpy(file_path, directorio);
            strcat(file_path, "/");
            strcat(file_path, dir_entry->d_name);
            strcpy(lista_archivos[i], file_path);
            i++;
        }
    }

    closedir(dir_stream);

    *n_archivos = i;
    return lista_archivos;
}



/**
 * void liberarArregloArchivos
 *
 * Recibe un arreglo de strings y libera la memoria asignada a cada string y al arreglo.
 *
 * Input:
 *     - char** arreglo : Arreglo de strings a liberar.
 *     - int n_archivos : Numero de archivos en el arreglo.
*/
void liberarArregloArchivos(char** arreglo, int n_archivos) {
    for (int i = 0; i < n_archivos; i++) {
        free(arreglo[i]);
    }
    free(arreglo);
}



/**
 * void crearCarpetas
 *
 * Crea las carpetas necesarias para guardar los archivos de salida.
*/
void crearCarpetas() {
    printf("Creando carpetas...\n");
    mkdir("./archivos/horizontal", 0755);
    mkdir("./archivos/vertical", 0755);
    mkdir("./archivos/horizontal/50x50", 0755);
    mkdir("./archivos/horizontal/100x100", 0755);
    mkdir("./archivos/horizontal/200x200", 0755);
    mkdir("./archivos/vertical/50x50", 0755);
    mkdir("./archivos/vertical/100x100", 0755);
    mkdir("./archivos/vertical/200x200", 0755);
}



/**
 * void moverArchivo
 *
 * Recibe la direccion de un archivo, si es horizontal o vertical y su dimension
 * y dependiendo de estos parametros mueve el archivo a la carpeta correspondiente.
 *
 * Input:
 *     - char* archivo : Direccion del archivo a mover.
 *     - int es_horizontal : 1 si el archivo es horizontal, 0 si es vertical.
 *     - int dimm : Dimension de la sopa de letras.
 *
 * Return:
 *     - char** : Puntero a un arreglo de strings con las direcciones de los archivos
*/
void moverArchivo(char* archivo, int es_horizontal, int dimm) {
    char* slash = strrchr(archivo, '/');

    char* destino = (char*)malloc(250 * sizeof(char));
    strcpy(destino, "./archivos/");
    if (es_horizontal) {
        strcat(destino, "horizontal/");
    } else {
        strcat(destino, "vertical/");
    }


    char dimm_str[10];
    sprintf(dimm_str, "%dx%d", dimm, dimm);
    strcat(destino, dimm_str);
    strcat(destino, "/");
    strcat(destino, slash + 1);

    printf("Moviendo archivo %s a %s\n", archivo, destino);

    rename(archivo, destino);
    free(destino);
}
