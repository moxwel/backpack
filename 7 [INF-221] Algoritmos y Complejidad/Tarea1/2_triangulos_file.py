import sys



def comparar(tupla_a, tupla_b, tupla_c, tupla_d):   #compara 2 tuplas c y d a partir de las tuplas de indices a y b
    tupla_iguales = set(tupla_a) & set(tupla_b)    #interseccion de la tuplas
    for x in tupla_iguales:    #verifica que los elementos pueden hacer join, en caso contrario retorna 0
        indice_1 = tupla_a.index(x)
        indice_2 = tupla_b.index(x)
        if tupla_c[indice_1] != tupla_d[indice_2]:
            return 0
    indices = []
    for x in tupla_a:    #almacena el indice del elemento de los elementos repetidos
        if x in tupla_iguales:
            indice_1 = tupla_a.index(x)
            indices.append(indice_1)
    lista1 = list(tupla_c)
    lista2 = list(tupla_d)
    union= set(tupla_a+tupla_b)    # une los indices
    cont  = 0
    lista_final =[]
    for x in union:    #recorre la union de indices
        if x in tupla_a:     #si el indice pertenece a la tupla a, entonces toma el elemento de la tupla c
            indice = tupla_a.index(x)
            lista_final.append(tupla_c[indice])
        else :      #si el indice pertenece a la tupla b, entonces toma el elemento de la tupla d
            indice = tupla_b.index(x)
            lista_final.append(tupla_d[indice])
    return lista_final   # retorna la lista final




# LEER ARCHIVO

cant_arg = len(sys.argv)
if cant_arg == 1:
	print("\nERROR: No se ingreso ningun archivo.")
	print("    Uso: python 2_triangulos_file.py \'archivo.txt\'\n")
	exit()

n_archivo = sys.argv[1]
try:
	archivo = open(n_archivo, "r")
except FileNotFoundError:
	print("\nERROR: El archivo no existe.\n")
	exit()



lista_tuplas = []
for line in archivo: # Guardar tuplas
    lista_tuplas.append(tuple(map(int, line.split(" "))))
#print(str(lista_tuplas))



lista_triangulos = []
for x in lista_tuplas: # Buscar triangulos
    for j in lista_tuplas:    # compara una tupla con el resto de tuplas
        join = comparar((0,1),(1,2), x, j)    # verifica si alguna hace join
        if join != 0:     # si hace join , compara la segunda tupla con el resto de tuplas
            for k in lista_tuplas:
                join = comparar((0,1),(1,2), j, k)
                join2 = comparar((0,1),(1,2),k , x)    # compara el join de la tupla final con la inicial para verificar el triangulo
                if join != 0 and join2 != 0 :    #verifica los join
                    lista= [x , j , k]    # crea la lista con las 3 tuplas ordenadas
                    lista.sort()
                    if lista not in lista_triangulos:    # en caso de que no exista el triangulo, lo agrega a la lista de triangulos
                        lista_triangulos.append(lista)
                        #print("encontre un triangulo")

print(str(len(lista_triangulos)))

archivo.close()
