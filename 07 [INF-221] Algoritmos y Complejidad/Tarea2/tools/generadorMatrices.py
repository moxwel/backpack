import random

dimen = int(input("Dimension de matrices a generar: "))

nombre = "matriz"+str(dimen)+"_"+str(random.randint(1,1000))+".txt"
archivo = open(nombre, "w")
archivo.write(str(dimen) + "\n")

for n in range(dimen*2):

	valores = ""
	for n in range(dimen):
		valores += str(random.randint(1,10)) + " "
	archivo.write(valores + "\n")

archivo.close()
