# Cuantos terminos se utlizaran en la suma infinita
n           = int(input())
suma        = 0

for x in range(n):
    denominador = ((-1)**x) * (2*x+1)
    suma = suma + (1/denominador)
    
print(4*suma)