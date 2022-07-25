numinput = int(input("Ingrese largo maximo: "))

def pasos_collatz(num):
    cont = 1
    while num != 1:
        if num % 2 == 0:
            num = num/2
            cont += 1
        else:
            num = (3*num) + 1
            cont += 1
    return cont

i = 1
while pasos_collatz(i) <= numinput:
    print(i,"*"*pasos_collatz(i))
    i += 1