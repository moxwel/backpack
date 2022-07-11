def valor(inicio,final,objeto,demanda):
    min_inicio = inicio%100
    min_final = final%100

    h_inicio = inicio//100
    h_final = final//100

    min_inicio += h_inicio*60
    min_final += h_final*60

    diff = min_final - min_inicio

    if diff <= 0:
        return 0
    else:
        if objeto == "libro":
            if demanda == "alta":
                return diff*30
            else:
                return diff*10
        elif objeto == "calculadora":
            return diff*50
        else:
            if demanda == "alta":
                return diff*5
            else:
                return diff

elem = int(input("Ingrese numero de elementos a cobrar: "))

suma = 0
cont = 1
while cont <= elem:
    obj = input("Ingrese objeto " + str(cont) + ": ")

    if obj == "calculadora":
        dem = " "
    else:
        dem = input("Ingrese demanda [alta/baja]: ")
    
    hora_norm = int(input("A que hora debera devolverlo? [HHMM]: "))
    hora_real = int(input("A que hora lo devolvio? [HHMM]: "))

    suma += valor(hora_norm,hora_real,obj,dem)
    cont += 1

print("El total de multas fue de",suma)
