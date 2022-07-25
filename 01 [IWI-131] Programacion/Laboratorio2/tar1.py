encuesta = [[12,5,1999,"A"],[29,5,1984,"B"],[11,11,2000,"A"],[1,3,1993,"B"],
[27,9,1988,"A"],[1,3,1963,"A"],[1,3,1972,"B"],[1,3,1948,"A"]]

encuesta = [[12,5,1999,"A"],[29,5,1984,"C"],[11,11,2000,"A"],
            [1,3,1993,"B"],[27,9,1988,"B"],[1,3,1963,"B"],
            [1,3,1972,"C"],[1,3,1948,"A"]]

def proporcion_votos(encuesta, candidato):
    lis = []
    for x in encuesta:
        lis.append(x[3])

    #print(lis,len(lis),lis.count(candidato))
    prop = lis.count(candidato)/len(lis)
    return(round(prop,3))


def distribucion_x_edad(encuesta, fecha_actual):
    fa = fecha_actual[0] + 30*fecha_actual[1] + 365*fecha_actual[2]
    #print(fa)
    me = 0
    aj = 0
    ad = 0
    am = 0
    for x in encuesta:
        edad = abs((x[0] + 30*x[1] + 365*x[2]-fa)//365)
        #print(edad)

        if edad < 18:
            me += 1
        elif 29>= edad >= 18:
            aj += 1
        elif edad >= 30 and edad <= 64:
            ad += 1
        elif edad >= 65:
            am += 1

    return((me,aj,ad,am))

def pronostico(encuesta):
    dic = {}

    for x in encuesta:
        if x[3] not in dic:
            dic[x[3]] = 0
        dic[x[3]] += 1

    lis = []
    for x in dic:
        lis.append((dic[x],x))

    lis.sort()
    lis.reverse()
    mayor = lis[0][0]
    #print(lis,mayor)

    lis2 = []
    for x in lis:
        if x[0] == mayor:
            lis2.append(x[1])

    lis2.sort()
    return(lis2)



        
    
print(distribucion_x_edad(encuesta, (26,7,2019)))
