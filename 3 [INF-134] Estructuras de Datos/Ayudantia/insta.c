#include <string.h>
#include <stdio.h>
#include <stdlib.h>

// Cada foto tiene tags
typedef struct{
    char descripcion[100];
    char** tags;
}Foto;

// Cada perfil tiene fotos y amigos
// Cada amigo es un perfil
typedef struct Perfil{
    int cantFotos;
    int cantAmigos;
    char nombre[100];
    struct Perfil** amigos;
    Foto** fotos;
}Perfil;

int getCantFotos(Perfil* p);
int getCantAmigos(Perfil* p);
int getCantTags(Foto* p);
char* getNombre(Perfil* p);

Perfil** getAmigos(Perfil* p);
Foto** getFotos(Perfil* p);

char* getDescripcion(Foto* f);

char** getTags(Foto* f);
