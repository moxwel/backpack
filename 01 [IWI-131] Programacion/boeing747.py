def costo(peso):
    if peso >= 0 and peso <= 25:
        return peso * 1000
    elif peso >= 26 and peso <= 300:
        return peso * 1500
    elif peso >= 301 and peso <= 500:
        return peso * 2000

hayEspacio = True
pesoTotal = 0
precio = 0
cant_bultos = 0
pesoMayor = 0

while hayEspacio:
    pesoBulto = int(input("Peso del bulto:"))
    if pesoBulto <= 500:
        if 18000 < pesoTotal + pesoBulto: # Si se supera el espacio disponible
            hayEspacio =  False
        else: # Aun queda espacio
            pesoTotal += pesoBulto
            precio += costo(pesoBulto)
            cant_bultos += 1
            if pesoBulto > pesoMayor:
                pesoMayor = pesoBulto
    else:
        print("Rechazado")

    #print("debug",pesoTotal,precio,cant_bultos,pesoMayor,)

print("---")
print("No hay mas espacio")
print("Numero de bultos:",cant_bultos)
print("Peso del bulto mas pesado:",pesoMayor)
print("Peso promedio:",round(pesoTotal/cant_bultos))
print("Ingreso total:",precio)
