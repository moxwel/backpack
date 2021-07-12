def sumaDados(cant):
    i = 0
    suma = 0
    atragantamiento = False # Flag
    
    while i < cant: # Se repite tantas veces como dados hayan
    
        #print("Dado",i+1)
        dado = int(input())
        
        if dado == 1: # Si es 1...
            atragantamiento = True # ... entonces se atraganta
            #print("Atr!")
            i += 1
        else:
            suma += dado
            i += 1
            
    if atragantamiento: # Si se atraganta...
        suma = 1 # ...el puntaje final es 1
        
    #print(suma)
    return(suma)


cant1 = int(input()) # Cantidad de dados jugador 1
dados1 = sumaDados(cant1)
    
cant2 = int(input()) # Cantidad de dados jugador 2
dados2 = sumaDados(cant2)

print(dados1,dados2)

#print("Primeros digitos",dados1 % 10,dados2 % 10)
#print("Segundos digitos",dados1 // 10,dados2 // 10)

if dados1 % 10 == dados2 % 10:
    print("Canallada")
    temp = dados1
    dados1 = dados2
    dados2 = temp
    print(dados1,dados2)
    
if dados1 // 10 != 0 and dados2 // 10 != 0:
    if dados1 // 10 == dados2 // 10:
        print("Canallada")
        temp = dados1
        dados1 = dados2
        dados2 = temp
        print(dados1,dados2)
    