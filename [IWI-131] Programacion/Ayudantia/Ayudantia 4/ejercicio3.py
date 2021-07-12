def todos_distintos(lista):
    
    ant = []
    dist = True

    for x in lista:
        
        if x in ant:
            dist = False
        ant.append(x)

    return dist

def todos_iguales(lista):
    
    modelo = lista[0]
    iguales = True
    
    for x in lista:
        if modelo != x:
            iguales = False
            
    return iguales
        
        
print(todos_iguales([7,7,7,7,7,7,7,7,7,7,7,7]))
print(todos_distintos(list(range(1000))))
print(todos_iguales([12]))
print(todos_distintos(list("hiperblanduzcos")))
