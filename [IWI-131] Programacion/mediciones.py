registro = [
    (2016, 1, 25, "d2:00:10", 0.0112),
    (2015, 10, 24, "3e:15:c2", 0.0105),
    (2015, 11, 3, "28:0b:5c", 0.0009),
]

def agregar_medicion(registro, a, m, d, iden, val):
    tupla = (a,m,d,iden,val)

    indice = -1
    for valor in range(len(registro)):
        # print(valor, val, registro[valor][4], val >= registro[valor][4])
        if val >= registro[valor][4]:
            indice = valor
            break
    
    # print(indice)
    if indice == -1:
        registro.append(tupla)
    else:
        registro.insert(indice,tupla)


def corregir_valor(registro,a,m,d,iden,nuevo_val):
    t = (a,m,d,iden)
    indice = -1
    for i in range(len(registro)):
        t2 = registro[i]
        if t == t2[:-1]:
            indice = i
            break

    print(i)
    del registro[indice]

    agregar_medicion(registro,a,m,d,iden,nuevo_val)
    return registro

def imprimir(registro):
    print("[")
    for x in registro:
        print(x)
    print("]")

imprimir(registro)
agregar_medicion(registro,2016,3,3,"AAAAAA",0.0099)
imprimir(registro)
agregar_medicion(registro,2016,3,3,"BBBBBB",0.0100)
imprimir(registro)
agregar_medicion(registro,2016,3,3,"CCCCCC",1.0100)
imprimir(registro)
agregar_medicion(registro,2016,3,3,"DDDDDD",0.0001)
imprimir(registro)
corregir_valor(registro,2016,3,3,"DDDDDD",26)
imprimir(registro)
