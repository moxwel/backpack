nint = int(input())
mint = int(input())

def euclides(n, m):
    resto = n % m

    while resto != 0:
        n = m
        m = resto
        resto = n % m
        
    return m

print(euclides(nint, mint))