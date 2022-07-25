def costo_transito(min_comb):
    total = 0
    if min_comb > 480:
        difer = min_comb - 480
        hor = difer // 60
        if difer % 60 != 0:
            total += 5000
        total = 5000 * hor + total
        return total
    else:
        return total

def costo_pasajero(hrs_viaje,costo_base,min_comb):
    return hrs_viaje*30000 + costo_base + costo_transito(min_comb)

vuelo = input("Vuelo: ")
mini_cost = float("inf")
mini_nomb = ""

while vuelo != "FIN":
    horas = int(input("Horas: "))
    minut = int(input("Minutos de combinacion: "))
    base = int(input("Costo base: "))
    costo = costo_pasajero(horas,base,minut)

    if costo < mini_cost:
        mini_cost = costo
        mini_nomb = vuelo
    vuelo = input("Vuelo: ")

print("---")
print("Mejor vuelo:",mini_nomb)
print("Costo pasajero:",mini_cost)
