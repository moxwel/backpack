from math import ceil as techo

nLibros = int(input())
libros = input()
presup = int(input())
libros = list(map(int, libros.split(" ")))

# print(nLibros)
# print(libros)
# print(presup)

izq = 0
der = nLibros-1
mid = techo((izq+der)/2)

while ((libros[izq]+libros[der]) > presup):
	# print(str(libros[der]) +" + "+ str(libros[izq]) + " es mayor a "+str(presup)+"\n"+str(libros[der])+" es muy grande.\n")
	der = techo((der+mid)/2)


if ((libros[izq]+libros[der]) == presup):
	# print(str(libros[der]) +" + "+ str(libros[izq]) + " es igual a "+str(presup))
	# print("Se deben comprar los libros "+str(izq+1)+" y "+str(der+1)+".")
	print(str(izq+1) + " " + str(der+1))
	quit()

#print(str(libros[der]) +" + "+ str(libros[izq]) + " es menor a "+str(presup)+"\n"+str(libros[izq])+" es muy bajo.\n")
izq = techo((izq+mid)/2)

while ((libros[izq]+libros[der]) > presup):
	# print(str(libros[der]) +" + "+ str(libros[izq]) + " es mayor a "+str(presup)+"\n"+str(libros[der])+" es muy grande.\n")
	izq = techo(izq/2)

if ((libros[izq]+libros[der]) == presup):
	# print(str(libros[der]) +" + "+ str(libros[izq]) + " es igual a "+str(presup))
	# print("Se deben comprar los libros "+str(izq+1)+" y "+str(der+1)+".")
	print(str(izq+1) + " " + str(der+1))
	quit()
