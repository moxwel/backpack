def cortarStr(texto):                   # Para cortar strings segun el largo de este
    n = len(texto) // 3                 # Calculo "n"
    return list(texto[:n])              # Lista como string cortado

def ganador(texto):
    # Contadores Armada, Ejercito, aViacion
    cA = 0
    cE = 0
    cV = 0

    # Contar las veces que sale cada uno, en el string cortado
    for compet in cortarStr(texto):
        if compet == "A":
            cA += 1
        elif compet == "E":
            cE += 1
        elif compet == "V":
            cV += 1

    #print(cA,cE,cV)

    # El ganador es el que aparece mas veces
    ganador = max(cA,cE,cV)
    # Se busca cual de los 3 es el mayor
    if ganador == cA:
        return "A"
    elif ganador == cE:
        return "E"
    elif ganador == cV:
        return "V"
    

cantComp = int(input("Numero de competencias? "))

cont = 1
while cont <= cantComp:
    resultado = input("Resultados Competencia " + str(cont) + ": ")
    #print(cortarStr(resultado))
    #print(ganador(resultado))

    ronda = ganador(resultado)          # Ganador de la ronda
    
    cantA = 0
    cantE = 0
    cantV = 0
    
    if ronda == "A":
        cantA += 1
    elif ronda == "E":
        cantE += 1
    elif ronda == "V":
        cantV += 1
        
    cont += 1

elMejor = max(cantA,cantE,cantV)

if elMejor == cantA:
    print("Gano Armada")
elif elMejor == cantE:
    print("Gano Ejercito")
elif elMejor == cantV:
    print("Gano Aviacion")
