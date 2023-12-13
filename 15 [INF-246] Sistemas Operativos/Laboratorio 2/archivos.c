#include "archivos.h"

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

void liberarArregloArchivos(char** arreglo, int n_archivos) {
    for (int i = 0; i < n_archivos; i++) {
        free(arreglo[i]);
    }
    free(arreglo);
}
