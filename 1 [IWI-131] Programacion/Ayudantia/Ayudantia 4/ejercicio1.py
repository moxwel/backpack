def todos_iguales(lista):
    
    modelo = lista[0]
    iguales = True
    
    for x in lista:
        if modelo != x:
            iguales = False
            
    return iguales
        
print(todos_iguales([2,2,2]))
