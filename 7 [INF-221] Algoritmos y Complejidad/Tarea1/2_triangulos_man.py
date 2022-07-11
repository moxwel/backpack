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



x = input("[[\'EOF\' (literal) para terminar]]\nIngrese arco: ")
lista_tuplas = []
while x not in ("EOF", "eof"):    #lee las tuplas ingresadas mientras entrada  no sea EOF
    lista = x.split(" ")
    lista= map(int, lista)
    tupla = tuple(lista)
    lista_tuplas.append(tupla)
    #print(tupla)
    x = input("Ingrese arco: ")



lista_triangulos= []
for x in lista_tuplas:
    for j in lista_tuplas:    # compara una tupla con el resto de tuplas
        join = comparar((0,1),(1,2), x, j)    # verifica si alguna hace join
        if join != 0:     # si hace join , compara la segunda tupla con el resto de tuplas
            for k in lista_tuplas:
                join = comparar((0,1),(1,2), j, k)
                join2 = comparar((0,1),(1,2),k , x)    # compara el join de la tupla final con la inicial para verificar el triangulo
                if join != 0 and join2 != 0 :    #verifica los join
                    lista= [x , j , k]    # crea la lista con las 3 tuplas ordenadas
                    lista.sort()
                    if lista not in  lista_triangulos:    # en caso de que no exista el triangulo, lo agrega a la lista de triangulos
                        lista_triangulos.append(lista)
                        #print("encontre un triangulo")

print("Cantidad de triangulos: " + str(len(lista_triangulos)))
