
import random
import time

#3
# 1 2     1 2    7 4
# 3 9     3 1   30 15
#[[1,2,3],[3,9,2],[3,1,2]]  [[1,2,6],[3,1,1],[7,6,3]]
#[[1,2,3,1],[3,1,3,9],[2,9,8,1],[7,3,4,5]]
def suma_de_vectores(a,b):
    vacia= []
    for x in range(len(a)):
        valor = a[x] + b[x]
        vacia.append(valor)
    return vacia
def resta_de_vectores(a,b):
    vacia= []
    for x in range(len(a)):
        valor = a[x] - b[x]
        vacia.append(valor)
    return vacia
def suma_de_matrices(a,b):
    vacia = []
    for x in range(len(a)):
        suma= suma_de_vectores(a[x], b[x])
        vacia.append(suma)
    return vacia
def resta_de_matrices(a,b):
    vacia = []
    for x in range(len(a)):
        resta= resta_de_vectores(a[x], b[x])
        vacia.append(resta)
    return vacia
def division_de_matrices(a):
    largo = len(a)
    if largo >= 2:
        medio=int(largo/2)
        matriz_1=[]
        matriz_2=[]
        matriz_3=[]
        matriz_4=[]
        for x in range(largo):
            if x < medio:
                aux1 = a[x][:medio]
                aux2= a[x][medio:]
                matriz_1.append(aux1)
                matriz_2.append(aux2)
            else :
                aux1 = a[x][:medio]
                aux2= a[x][medio:]
                matriz_3.append(aux1)
                matriz_4.append(aux2)
    return matriz_1, matriz_2,matriz_3,matriz_4
def union_de_matrices(matriz_1,matriz_2,matriz_3,matriz_4):
    matriz_final = []
    for x in range(len(matriz_1)):
        aux = matriz_1[x] + matriz_2[x]
        matriz_final.append(aux)
    for x in range(len(matriz_3)):
        aux = matriz_3[x] + matriz_4[x]
        matriz_final.append(aux)
    return matriz_final
def printear(a):
    for x in a:
        print(str(x)+"\n")
    print("\n")
def producto_matrices(a, b):
    filas_a = len(a)
    filas_b = len(b)
    columnas_a = len(a[0])
    columnas_b = len(b[0])
    if columnas_a != filas_b:
        return None
    # Asignar espacio al producto. Es decir, rellenar con "espacios vacíos"
    producto = []
    for i in range(filas_b):
        producto.append([])
        for j in range(columnas_b):
            producto[i].append(None)
    # Rellenar el producto
    for c in range(columnas_b):
        for i in range(filas_a):
            suma = 0
            for j in range(columnas_a):
                suma += a[i][j]*b[j][c]
            producto[i][c] = suma
    return producto
def calculo2(a, b):
    if len(a) == 2 or len(a) % 2 != 0:
        matriz = producto_matrices(a, b)
        return matriz
    else :
        m1, m2, m3, m4 = division_de_matrices(a)
        m5, m6, m7, m8 = division_de_matrices(b)
        p1_a= resta_de_matrices(m6, m8)
        p1 = calculo2(m1,p1_a)

        p2_a = suma_de_matrices(m1, m2)
        p2 = calculo2(p2_a, m8)

        p3_a = suma_de_matrices(m3, m4)
        p3 = calculo2(p3_a, m5)

        p4_a= resta_de_matrices(m7, m5)
        p4 = calculo2(m4 , p4_a)

        p5_a= suma_de_matrices(m1,m4)
        p5_b=suma_de_matrices(m5,m8)
        p5= calculo2(p5_a, p5_b)

        p6_a= resta_de_matrices(m2 , m4)
        p6_b=suma_de_matrices(m7,m8)
        p6= calculo2(p6_a, p6_b)

        p7_a= resta_de_matrices(m1,m3)
        p7_b=suma_de_matrices(m5,m6)
        p7= calculo2(p7_a, p7_b)

        parte_1 = suma_de_matrices(p5, p4)
        parte_1 = resta_de_matrices(parte_1, p2)
        parte_1 = suma_de_matrices(parte_1,p6)

        parte_2 = suma_de_matrices(p1, p2)

        parte_3 = suma_de_matrices(p3, p4)

        parte_4 = suma_de_matrices(p1, p5)
        parte_4 = resta_de_matrices(parte_4, p3)
        parte_4 = resta_de_matrices(parte_4, p7)
        matriz = union_de_matrices(parte_1, parte_2, parte_3, parte_4)
        return matriz
def calculo(a, b):
    if len(a) == 2 or len(a) % 2 != 0:
        matriz = producto_matrices(a, b)
        return matriz
    else :
        m1, m2, m3, m4 = division_de_matrices(a)
        m5, m6, m7, m8 = division_de_matrices(b)
        printear(m8)       #Si tu matriz es muy grande, comenta esta linea.
        p1_1 = calculo(m1 , m5)
        p1_2 = calculo(m2 , m7)
        p1 = suma_de_matrices(p1_1, p1_2)

        p2_1 = calculo(m1 , m6)
        p2_2 = calculo(m2 , m8)
        p2 = suma_de_matrices(p2_1, p2_2)
        p3_1 = calculo(m3 , m5)
        p3_2 = calculo(m4 , m7)
        p3 = suma_de_matrices(p3_1, p3_2)
        p4_1 = calculo(m3 , m6)
        p4_2 = calculo(m4 , m8)
        p4 = suma_de_matrices(p4_1, p4_2)
        matriz = union_de_matrices(p1,p2,p3,p4)
        return matriz

dimen = int(input())

matriz_1 = []
matriz_2 = []

for n in range(dimen):
    fila = input()
    fila = fila.strip().split(" ")
    fila = list(map(int, fila))

    if (len(fila) != dimen):
        print("ERROR: Las dimensiones no coinciden")
        exit()

    matriz_1.append(fila)
for n in range(dimen):
    fila = input()
    fila = fila.strip().split(" ")
    fila = list(map(int, fila))

    if (len(fila) != dimen):
        print("ERROR: Las dimensiones no coinciden")
        exit()

    matriz_2.append(fila)


print("Matrices iniciales: ")   #Si tu matriz es muy grande, comenta esta linea.
printear(matriz_1)              #Si tu matriz es muy grande, comenta esta linea.
printear(matriz_2)              #Si tu matriz es muy grande, comenta esta linea.
v1=time.time()
final = calculo2(matriz_1, matriz_2)
v2 = time.time()
print("Dimension de la matriz resultante: ", dimen)
print("Codigo reducido (Strassen) en un tiempo de :" ,v2-v1)

print("Matriz resultante: ")
printear(final)     #Si tu matriz es muy grande, comenta esta linea.

x1 = time.time()
final2 = calculo(matriz_1, matriz_2)
x2 = time.time()
print("Dimension de la matriz resultante: ", dimen)
print("Codigo Tradicional en un tiempo de: ", x2-x1)

print("Matriz resultante: ")
printear(final2)    #Si tu matriz es muy grande, comenta esta linea.
