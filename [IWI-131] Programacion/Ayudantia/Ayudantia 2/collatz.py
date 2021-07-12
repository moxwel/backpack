def collatz(num):
    list = [num]
    
    while num != 1:
        if num % 2 == 0:
            num = num/2
            list.append(int(num))
        else:
            num = (3*num) + 1
            list.append(int(num))
    
    return list
    
print(collatz(int(input())))