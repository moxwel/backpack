def todos_d(lista):
    
    ant = []
    dist = True

    for x in lista:
        
        if x in ant:
            dist = False
        ant.append(x)

    return dist
        
        
print(todos_d([2,4,2]))
