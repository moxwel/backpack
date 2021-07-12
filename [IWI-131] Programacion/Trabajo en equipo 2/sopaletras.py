from string import ascii_letters
from random import choice

def iniciar_tablero(dimensiones):
    n_filas, n_colum = dimensiones
    tablero = []

    for fila in range(n_filas):
        fila = []
        for colum in range(n_colum):
            letra = choice(ascii_letters)
            fila.append(letra.lower())
        # print(fila)
        tablero.append(fila)

    # print(tablero)
    return tablero

def imprimir_tablero(tablero):
    n_filas = len(tablero)
    n_colum = len(tablero[0])

    for f in range(n_filas):
        s = "| "
        for c in range(n_colum):
            s += tablero[f][c] + " | "
        print(s)

def agregar_palabra(tablero, argum):
    n_filas = len(tablero)
    n_colum = len(tablero[0])

    palabra, coordenada, orientacion = argum
    f, c = coordenada
    palabra = palabra.upper()

    if orientacion == "H":
        if c + len(palabra) <= n_colum:
            for i in range(len(palabra)):
                tablero[f][c + i] = palabra[i]
        else:
            return tablero
    elif orientacion == "V":
        if f + len(palabra) <= n_filas:
            for i in range(len(palabra)):
                tablero[f + i][c] = palabra[i]
        else:
            return tablero
    
    return tablero


tab = iniciar_tablero((10,10))
print(tab)
# imprimir_tablero(tab)

agregar_palabra(tab,("perro",(0,6),"H"))
imprimir_tablero(tab)