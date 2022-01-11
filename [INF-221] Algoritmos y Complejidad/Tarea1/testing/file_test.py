import sys








def tupla (texto):   # transforma un string de 2 numeros a una tupla de numeros enteros
    texto = texto.split(" ")
    int_list = list(map(int, texto))
    texto = tuple(int_list)
    return texto





# LEER ARCHIVO

cant_arg = len(sys.argv)

if cant_arg == 1:
	print("\nERROR: No se ingreso ningun archivo.")
	print("    Uso: python file_test.py \'archivo.txt\'\n")
	exit()

n_archivo = sys.argv[1]

try:
	archivo = open(n_archivo, "r")
except FileNotFoundError:
	print("\nERROR: El archivo no existe.\n")
	exit()



# TABLA 1

x = int(archivo.readline().strip()) # Cantidad atributos
print("Cantidad de atributos: " + str(x))

lista_indices = list(tupla(archivo.readline().strip())) # Lista de atributos
print("Atributos: "+ str(lista_indices))

x = int(archivo.readline().strip()) # Cantidad tuplas
print("Cantidad de tuplas: " + str(x))

lista_tuplas = []
for line in range(x): # Leer y almacenar tuplas
	t = tupla(archivo.readline().strip())
	print("Tupla " + str(line) + ": " + str(t))
	lista_tuplas.append(t)

print(str(lista_tuplas))




# TABLA 2

x = int(archivo.readline().strip()) # Cantidad atributos
print("Cantidad de atributos: " + str(x))

lista_indices2 = list(tupla(archivo.readline().strip())) # Lista de atributos
print("Atributos: "+ str(lista_indices2))

x = int(archivo.readline().strip()) # Cantidad tuplas
print("Cantidad de tuplas: " + str(x))

lista_tuplas2 = []
for line in range(x): # Leer y almacenar tuplas
	t = tupla(archivo.readline().strip())
	print("Tupla " + str(line) + ": " + str(t))
	lista_tuplas2.append(t)

print(str(lista_tuplas2))
