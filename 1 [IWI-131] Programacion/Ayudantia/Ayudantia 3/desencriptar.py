def codigo_palabra(codigo):
    return codigo[-1::-2]

def codigo_hora(t):
    sep = t.index(":")
    horas = t[:sep]
    minutos = t[sep+1:]

    sumaH = 0
    for x in horas:
        sumaH += int(x)
    sumaM = 0
    for x in minutos:
        sumaM += int(x)
        
    print(str(sumaH % 24) + ":" + str(sumaM % 60))

codigo_hora("776199:68556")