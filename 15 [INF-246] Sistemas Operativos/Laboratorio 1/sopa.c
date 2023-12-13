#include <stdio.h>
#include <stdlib.h>
#include <string.h>

/**
 * void quitarSaltoLinea
 *
 * Recibe un string y elimina el salto de linea
 * (si existe) al final del string.
 *
 * Input:
 *     - char* string : String al que se le quiere eliminar el salto de linea.
 **/
void quitarSaltoLinea(char* string){
    int largo_str = strlen(string);
    if (string[largo_str - 1] == '\n') {
        if (string[largo_str - 2] == '\r') {
            string[largo_str - 2] = '\0';
        }
        string[largo_str - 1] = '\0';
    }
}



/**
 * char** generarMatrizSopa
 *
 * Recibe la direccion del archivo de la sopa de letras y
 * retorna una matriz con las letras de la sopa
 * de letras y su dimension.
 *
 * Input:
 *     - char* file_path : String con la direccion de la sopa de letras.
 *     - int* dimm : Puntero a entero donde se guardará la dimension de la sopa de letras.
 *
 * Return:
 *     - char** : Puntero a una matriz cuadrada de dimension 'dim' con las letras
 *                de la sopa de letras.
 **/
char** generarMatrizSopa(char* file_path, int* dimm){
    FILE* archivo = fopen(file_path, "r");

    if (archivo == NULL) {
        // printf("Error al abrir el archivo\n");
        return NULL;
    }

    // Las lineas del archivo no van a ser mas largas que 400 caracteres
    char buff_linea[415];

    // La primera linea se salta (dice "horizontal" o "vertical")
    fgets(buff_linea, sizeof(buff_linea), archivo);

    // Crear la matriz. Un arreglo de arreglos de caracteres.
    // Se asume que ninguna sopa sera mayor a 200x200
    char** sopa_letras = (char**)malloc(215 * sizeof(char*));

    int i = 0;
    while(fgets(buff_linea, sizeof(buff_linea), archivo) != NULL) {

        // Eliminar el salto de linea (si existe)
        quitarSaltoLinea(buff_linea);

        // printf("Linea %d: %s\n", i, buff_linea);

        // Tokenizar la linea
        char* token = strtok(buff_linea, " ");

        // Asignamos memoria para la fila. Un arreglo de caracteres.
        char* fila = (char*)malloc(215 * sizeof(char));

        // Se guarda la fila en la matriz
        sopa_letras[i] = fila;

        int j = 0;
        while (token != NULL) {
            // printf("Token: %s\n", token);

            // Guardar el token en la fila
            fila[j] = *token;

            token = strtok(NULL, " ");
            j++;
        }
        // printf("j: %d\n", j);

        i++;
    }
    // printf("i: %d\n", i);

    fclose(archivo);

    *dimm = i;
    return sopa_letras;
}



/**
 * void liberarMatrizSopa
 *
 * Recibe una matriz y su dimension y libera la memoria
 * asignada a la matriz. Se asume que la matriz es cuadrada.
 *
 * Input:
 *     - char** matriz : Matriz a liberar.
 *     - int dimm : Dimension de la matriz.
 **/
void liberarMatrizSopa(char** matriz, int dimm){
    for (int i = 0; i < dimm; i++) {
        free(matriz[i]);
    }
    free(matriz);
}



/**
 * void extaerNombreArchivo
 *
 * Recibe una direccion de archivo y retorna el nombre del archivo sin extension.
 *
 * Input:
 *     - char* file_path : String con la direccion del archivo.
 *     - char* file_name : Puntero a string donde se guardará el nombre del archivo.
 **/
void extaerNombreArchivo(char* file_path, char* file_name) {
    // De un path como "../archivos/Carne.txt" se extrae el nombre del archivo sin extension "Carne"
    // Se asume que el nombre del archivo no tiene mas de 100 caracteres

    char* slash = strrchr(file_path, '/');

    // Copiar el nombre del archivo con extension
    if (slash != NULL){
        // Hay un slash en el path, la variable 'lastOccurrence' apunta a dicho slash
        strcpy(file_name, slash + 1);
    } else {
        strcpy(file_name, file_path);
    }

    // Eliminar la extension
    char* punto = strrchr(file_name, '.');
    // La variable 'extension' apunta al punto de la extension

    // Reemplazar el punto por un caracter nulo
    if (punto != NULL) {
        *punto = '\0';
    }
}



/**
 * int orientacionSopa
 *
 * Recibe la direccion de un archivo de sopa de letras y retorna
 * si la sopa de letras es horizontal (1) o vertical (0).
 *
 * Input:
 *     char* ruta_archivo : String con la direccion del archivo de la sopa de letras.
 *
 * Return:
 *     - int : 1 si la sopa de letras es horizontal.
 *             0 si la sopa de letras es vertical.
 *             -1 si hubo un error al abrir el archivo.
 **/
int orientacionSopa(char* ruta_archivo) {
    // Verificar si la primera linea dice "horizontal" o "vertical"
    FILE* archivo = fopen(ruta_archivo, "r");
    if (archivo == NULL) {
        printf("[!] Error al abrir el archivo\n");
        return -1;
    }

    char buff_orient[15];
    fgets(buff_orient, sizeof(buff_orient), archivo);
    quitarSaltoLinea(buff_orient);

    int es_horizontal = 0;
    if (strcmp(buff_orient, "horizontal") == 0) {
        es_horizontal = 1;
    }

    fclose(archivo);

    return es_horizontal;
}



/**
 * int buscarPalabraSopa
 *
 * Busca una palabra en la sopa de letras.
 *
 * Input:
 *     - char** sopa_letras : Matriz con la sopa de letras.
 *     - int dimm : Dimension de la sopa de letras.
 *     - char* palabra : Palabra que se quiere buscar en la sopa de letras.
 *     - int es_horizontal : Entero que indica si la palabra se debe buscar en
 *                           horizontal (1) o en vertical (0).
 *
 * Return:
 *     - int : 1 si la palabra se encuentra en la sopa de letras.
 *             0 si la palabra no se encuentra en la sopa de letras.
 **/
int buscarPalabraSopa(char** sopa_letras, int dimm, char* palabra, int es_horizontal) {
    int largo_palabra = strlen(palabra);

    // Buscar palabra
    if(es_horizontal){
        printf("Buscando en horizontal...\n");

        // Buscar palabra en la matriz
        for (int i = 0; i < dimm; i++) {
            int k = 0;
            for (int j = 0; j < dimm; j++) {
                printf("%c", sopa_letras[i][j]);

                if (sopa_letras[i][j] == palabra[k]) {
                    // Letra encontrada
                    k++;
                } else {
                    // Letra no encontrada
                    k = 0;
                }

                if (k == largo_palabra) {
                    // Palabra encontrada
                    printf(" <- Palabra encontrada en la linea %d, columna %d\n", i+2, (j-largo_palabra+1)*2+2);
                    return 1;
                }
            }
            printf("\n");
        }

    } else {
        printf("Buscando en vertical...\n");

        // Buscar palabra en la matriz
        for (int j = 0; j < dimm; j++) {
            int k = 0;
            for (int i = 0; i < dimm; i++) {
                printf("%c", sopa_letras[i][j]);

                if (sopa_letras[i][j] == palabra[k]) {
                    // Letra encontrada
                    k++;
                } else {
                    // Letra no encontrada
                    k = 0;
                }

                if (k == largo_palabra) {
                    // Palabra encontrada
                    printf(" <- Palabra encontrada en la linea %d, columna %d\n", (i-largo_palabra)+3, j*2+2);
                    return 1;
                }
            }
            printf("\n");
        }
    }
    return 0;
}



/**
 * Equivalente a 'buscarPalabraSopa' pero sin imprimir nada
 **/
int buscarPalabraSopaNoPrint(char** sopa_letras, int dimm, char* palabra, int es_horizontal) {
    int largo_palabra = strlen(palabra);
    if(es_horizontal){
        for (int i = 0; i < dimm; i++) {
            int k = 0;
            for (int j = 0; j < dimm; j++) {
                if (sopa_letras[i][j] == palabra[k]) {
                    k++;
                } else {
                    k = 0;
                }
                if (k == largo_palabra) {
                    printf("    [!!!] Encontrado en %d:%d\n", i+2, (j-largo_palabra+1)*2+2);
                    return 1;
                }
            }
        }
    } else {
        for (int j = 0; j < dimm; j++) {
            int k = 0;
            for (int i = 0; i < dimm; i++) {
                if (sopa_letras[i][j] == palabra[k]) {
                    k++;
                } else {
                    k = 0;
                }
                if (k == largo_palabra) {
                    printf("    [!!!] Encontrado en %d:%d\n", (i-largo_palabra)+3, j*2+2);
                    return 1;
                }
            }
        }
    }
    return 0;
}



/**
 * Equivalente a 'buscarPalabraSopa' pero se busca en toda la sopa antes de retornar
 **/
int buscarPalabraSopaLento(char** sopa_letras, int dimm, char* palabra, int es_horizontal) {
    int resultado = 0;
    int largo_palabra = strlen(palabra);
    if(es_horizontal){
        for (int i = 0; i < dimm; i++) {
            int k = 0;
            for (int j = 0; j < dimm; j++) {
                if (sopa_letras[i][j] == palabra[k]) {
                    k++;
                } else {
                    k = 0;
                }
                if (k == largo_palabra) {
                    printf("    [!!!] Encontrado en %d:%d\n", i+2, (j-largo_palabra+1)*2+2);
                    resultado = 1;
                }
            }
        }
    } else {
        for (int j = 0; j < dimm; j++) {
            int k = 0;
            for (int i = 0; i < dimm; i++) {
                if (sopa_letras[i][j] == palabra[k]) {
                    k++;
                } else {
                    k = 0;
                }
                if (k == largo_palabra) {
                    printf("    [!!!] Encontrado en %d:%d\n", (i-largo_palabra)+3, j*2+2);
                    resultado = 1;
                }
            }
        }
    }
    return resultado;
}




/**
 * Equivalente a 'buscarPalabraSopa' pero se usan operadores ternarios
 **/
int buscarPalabraSopaRapido(char** sopa_letras, int dimm, char* palabra, int es_horizontal) {
    int largo_palabra = strlen(palabra);
    if(es_horizontal){
        for (int i = 0; i < dimm; i++) {
            int k = 0;
            for (int j = 0; j < dimm; j++) {
                sopa_letras[i][j] == palabra[k] ? k++ : (k = 0);
                if (k == largo_palabra) return 1;
            }
        }
    } else {
        for (int j = 0; j < dimm; j++) {
            int k = 0;
            for (int i = 0; i < dimm; i++) {
                sopa_letras[i][j] == palabra[k] ? k++ : (k = 0);
                if (k == largo_palabra) return 1;
            }
        }
    }
    return 0;
}
