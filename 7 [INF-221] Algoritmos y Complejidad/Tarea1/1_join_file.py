import sys
from typing import Union



def tupla (texto):   # transforma un string de 2 numeros a una tupla de numeros enteros
    texto = texto.split(" ")
    int_list = list(map(int, texto))
    texto = tuple(int_list)
    return texto

def comparar(tupla_a, tupla_b, tupla_c, tupla_d):   #compara 2 tuplas c y d a partir de las tuplas de indices a y b
    tupla_iguales = set(tupla_a) & set(tupla_b)    #interseccion de la tuplas
    for x in tupla_iguales:     #verifica que los elementos pueden hacer join, en caso contrario retorna 0
        indice_1 = tupla_a.index(x)
        indice_2 = tupla_b.index(x)
        if tupla_c[indice_1] != tupla_d[indice_2]:
            return 0
    indices = []
    for x in tupla_a:    #      almacena el indice del elemento de los elementos repetidos
        if x in tupla_iguales:
            indice_1 = tupla_a.index(x)
            indices.append(indice_1)
    lista1 = list(tupla_c)
    lista2 = list(tupla_d)
    union= set(tupla_a+tupla_b)    # une los indices
    lista_final =[]
    for x in union:    #recorre la union de indices
        if x in tupla_a:
            indice = tupla_a.index(x)
            lista_final.append(tupla_c[indice])   #si el indice pertenece a la tupla a, entonces toma el elemento de la tupla c
        else :
            indice = tupla_b.index(x)
            lista_final.append(tupla_d[indice])     #si el indice pertenece a la tupla b, entonces toma el elemento de la tupla d
    return lista_final



# LEER ARCHIVO

cant_arg = len(sys.argv)
if cant_arg == 1:
	print("\nERROR: No se ingreso ningun archivo.")
	print("    Uso: python 1_join_file.py \'archivo.txt\'\n")
	exit()

n_archivo = sys.argv[1]
try:
	archivo = open(n_archivo, "r")
except FileNotFoundError:
	print("\nERROR: El archivo no existe.\n")
	exit()



# TABLA 1

x = int(archivo.readline().strip()) # Cantidad atributos
#print("Cantidad de atributos: " + str(x))

lista_indices = list(tupla(archivo.readline().strip())) # Lista de atributos
#print("Atributos: "+ str(lista_indices))

x = int(archivo.readline().strip()) # Cantidad tuplas
#print("Cantidad de tuplas: " + str(x))

lista_tuplas = []
for line in range(x): # Leer y almacenar tuplas
	t = tupla(archivo.readline().strip())
	#print("Tupla " + str(line) + ": " + str(t))
	lista_tuplas.append(t)

#print(str(lista_tuplas))



# TABLA 2

x = int(archivo.readline().strip()) # Cantidad atributos
#print("Cantidad de atributos: " + str(x))

lista_indices2 = list(tupla(archivo.readline().strip())) # Lista de atributos
#print("Atributos: "+ str(lista_indices2))

x = int(archivo.readline().strip()) # Cantidad tuplas
#print("Cantidad de tuplas: " + str(x))

lista_tuplas2 = []
for line in range(x): # Leer y almacenar tuplas
	t = tupla(archivo.readline().strip())
	#print("Tupla " + str(line) + ": " + str(t))
	lista_tuplas2.append(t)

#print(str(lista_tuplas2))



# JOIN

lista_final=[]
for x in lista_tuplas:    #con cada tupla, recorre el resto de tuplas
    for j in lista_tuplas2:
        valor = comparar(lista_indices, lista_indices2, x, j)    # verifica si el par de tuplas hace join
        if valor != 0:
            lista_final.append(valor)  #añade la tupla a la lista final

#print(str(lista_final))



# SALIDA

union_indices = set(lista_indices).union(set(lista_indices2))

print(len(union_indices))
for n in union_indices:
    print(str(n) + " ", end="")
print()

print(len(lista_final))
for l in lista_final:
    for n in l:
        print(str(n) + " ", end="")
    print()

archivo.close()
