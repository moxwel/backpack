estafados = [('12.234.567-8', 2000000,'pelados_furiosos', (25, 5, 2015)),
             ('9.111.567-k', 5500000,'multibank', (1, 10, 2014)),
             ('14.987.007-1', 100000000,'inversiones_seguras',(30, 11, 2016)),
             ('12.234.567-8', 10000000,'multibank', (2, 7, 2015)),
             ('12.234.567-8', 2500000,'multibank', (18, 1, 2016))]

def estafado_por(lista, rut):
    l = []
    for x in lista:
        if rut in x:
            if x[2] not in l:
                l.append(x[2])
                
    if l == []:
        return "Rut no estafado"
    else:
        return l
    
def ranking(estafados):
    l_e = []
    for x in estafados:
        if x[2] not in l_e:
            l_e.append(x[2])
    l_sumas = []
    for empresa in l_e:
        suma = 0
        for x in estafados:
            if empresa == x[2]:
                suma += x[1]
        l_sumas.append((suma,empresa))
        
    l_sumas.sort()
    l_sumas.reverse()
    
    l_final = []
    for x in l_sumas:
        l_final.append((x[1],x[0]))
    
    return l_final
        
def estafados_por_fecha(estafados, inicial, final):
    def fecha(f):
        return f[0] + f[1]*30 + f[2]*365
    
    ini = fecha(inicial)
    fin = fecha(final)
    
    l = []
    for x in estafados:
        if fecha(x[3]) < fin and fecha(x[3]) > ini:
            l.append((x[1],x[0]))
            
    l.sort()
    l.reverse()
    
    l_f = []
    for x in l:
        l_f.append((x[1],x[0]))
    
    return l_f

print(estafados_por_fecha(estafados,(3,1,2016),(1,1,2017)))