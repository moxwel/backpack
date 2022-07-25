cant1 = int(input()) # Cantidad de dados del jugador 1

i = 0
sum1 = 0
while i < cant1:
    dado = int(input())
    sum1 += dado
    i += 1
    
cant2 = int(input()) # Cantidad de dados del jugador 2

i = 0
sum2 = 0
while i < cant2:
    dado = int(input())
    sum2 += dado
    i += 1
    
print(sum1,sum2)
