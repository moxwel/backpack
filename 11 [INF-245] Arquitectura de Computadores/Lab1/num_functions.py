# is_valid
# ----------
# Valida si un numero esta representado correctamente
# en la base indicada.
# ----------
# Entrada:
#     - base: int    -> Base del numero a validar
#     - num : string -> Numero a validar
# Retorno:
#     - boolean -> El numero es valido (True) o no (False)
def is_valid(base, num):
    valid = True

    digits = ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F']
    #print(f"Digitos validos para base {base}:", digits[:base])

    for char in num:
        if char not in digits[:base]:
            #print(f"Digito {char} no es valido.")
            valid = False
            break

    return valid



# base_to_dec
# ----------
# Convierte un numero desde alguna base
# numerica (entre 2-16) a un numero en base 10.
#
# NOTAS:
#     - SE ASUME QUE EL NUMERO ES VALIDO RESPECTO A LA BASE.
#     - NUMEROS BINARIOS RESULTANTES NO ESTAN EN COMPLEMENTO 2.
# ----------
# Entrada:
#     - base: int    -> Base del numero a convertir
#     - num : string -> Numero a convertir
# Retorno:
#     - string -> Numero convertido a base 10
def base_to_dec(base, num):
    # Convierte las posibles letras minusculas a mayusculas
    # para que pertenezcan al diccionario
    num = num.upper()

    # Diccionario con el valor de cada numero, hasta base 16
    digits = {"0":0, "1":1, "2":2, "3":3, "4":4,
              "5":5, "6":6, "7":7, "8":8, "9":9,
              "A":10, "B":11, "C":12, "D":13, "E":14, "F":15}

    # Exponente del digito mas significativo
    power = len(num) - 1

    # Sumatoria que calcula el número en base decimal
    dec_num = 0
    for i in range(len(num)):
        #print(f"{num[i]} * {base}^{power} ---> {digits[num[i]]} * {base}^{power}")
        dec_num += digits[num[i]] * (base**power)

        # Se le resta 1 al exponente, hasta llegar a 0
        power -= 1

    return str(dec_num)



# dec_to_bin
# ----------
# Convierte un numero decimal a binario.
#
# NOTAS:
#     - SE ASUME QUE EL NUMERO ES VALIDO EN BASE 10.
# ----------
# Entrada:
#     - dec_num: string -> Numero en decimal
# Retorno:
#     - string -> Numero en binario correspondiente
def dec_to_bin(dec_num):
    num = int(dec_num)

    # Se implementa cola para mas eficiencia al ingresar numeros
    # por la izquierda del arreglo (bit mas significativo).
    bin_num = ""

    is_zero = False
    while not is_zero:
        rem = num % 2
        num = num // 2
        # print(f"{num} {rem}\n---")
        bin_num = str(rem) + bin_num

        if num == 0:
            is_zero = True

    return bin_num



# bin_sum
# ----------
# Suma binaria basica.
#
# NOTAS:
#     - SE ASUME QUE LOS NUMEROS SON VALIDOS EN BASE 2.
# ----------
# Entrada:
#     - bin_num1: string -> Numero binario
#     - bin_num2: string -> Numero binario
# Retorno:
#     - string -> Suma binaria de los dos numeros ingresados
def bin_sum(bin_num1, bin_num2):
    # Normalizar tamanos de registro (extension de ceros)
    max_reg = max(len(bin_num1), len(bin_num2))
    bin_num1 = bin_num1.zfill(max_reg)
    bin_num2 = bin_num2.zfill(max_reg)

    #print(f"Numero 1: {bin_num1}\nNumero 2: {bin_num2}")

    # Recorrer los numeros desde el menos significativo al mas significativo
    carry = 0
    result = ""
    for i in range(max_reg-1, -1, -1):
        suma = int(bin_num1[i]) + int(bin_num2[i]) + carry
        result = str(suma % 2) + result

        if suma >= 2:
            carry = 1
        else:
            carry = 0

    if carry == 1:
        result = "1" + result

    return result



# invert_bin
# ----------
# Invierte los bits de un numero binario.
#
# NOTAS:
#     - SE ASUME QUE EL NUMERO ES VALIDO EN BASE 2.
# ----------
# Entrada:
#     - bin_num: string -> Numero en binario a invertir
# Retorno:
#     - string -> Numero binario invertido
def invert_bin(bin_num):
    new_num = ""

    # Cambiar los 0 por 1 y los 1 por 0
    for digit in bin_num:
        if digit == "0":
            new_num = new_num + "1"
        else:
            new_num = new_num + "0"

    return new_num



# sign_extension
# ----------
# Extiende por signo un numero en binario segun cantidad de registros.
#
# NOTAS:
#     - SE ASUME QUE EL NUMERO ES VALIDO EN BASE 2.
#     - SE ASUME QUE EL NUMERO ESTA EN COMPLEMENTO 2.
# ----------
# Entrada:
#     - bin_num: string -> Numero en binario
#     - reg    : int    -> Tamano del registro
# Retorno:
#     - string -> Numero extendido en binario.
#                 Si la cantidad de registros es menor que
#                 el tamano del numero, retorna string vacio.
def sign_extension(bin_num, reg):
    diff = reg - len(bin_num)
    final_number = ""

    if diff >= 0:
        # bin_num[0] tiene el bit mas significativo.
        # A patir de ese se va a extender el numero.
        final_number = bin_num[0]*diff + bin_num

    return final_number



# c2_to_dec
# ----------
# Convierte un numero binario en complemento 2 a decimal.
#
# NOTAS:
#     - SE ASUME QUE EL NUMERO ES VALIDO EN BASE 2.
#     - SE ASUME QUE EL NUMERO ESTA EN COMPLEMENTO 2.
# ----------
# Entrada:
#     - bin_num: string -> Numero en binario en complemento 2
# Retorno:
#     - string -> Numero en decimal correspondiente.
def c2_to_dec(bin_num):
    result = ""

    # Si el binario es negativo, se debe hacer el algoritmo, si no, solo retornar.
    if bin_num[0] == "1":
        # Invertir y sumar 1, luego se obtiene el numero positivo correspondiente.
        invertido = invert_bin(bin_num)
        sumado    = bin_sum(invertido, "1")
        positivo  = base_to_dec(2,sumado)
        negativo  = "-" + positivo
        #print(f"og:  {bin_num}\ninv: {invertido}\nsum: {sumado}\npos: {positivo}\nneg: {negativo}")

        result = negativo
    else:
        result = base_to_dec(2,bin_num)

    return result



# fits_reg
# ----------
# Comprueba si el numero binario cabe en el registro indicado.
#
# NOTAS:
#     - SE ASUME QUE EL NUMERO ES VALIDO EN BASE 2.
# ----------
# Entrada:
#     - bin_num: string -> Numero en binario
#     - reg    : int    -> Tamano del registro
# Retorno:
#     - boolean -> Si el numero cabe (True) o no (False) en el registro.
def fits_reg(bin_num, reg):
    fit_status = True

    if reg < len(bin_num):
        fit_status = False

    return fit_status



# has_overflow
# ----------
# Comprueba si la suma de dos numeros binarios en complemento 2 tienen
# overflow o no.
#
# NOTAS:
#     - SE ASUME QUE LOS NUMEROS SON VALIDOS EN BASE 2.
#     - SE ASUME QUE LOS NUMEROS ESTAN EN COMPLEMENTO 2.
# ----------
# Entrada:
#     - bin_num1: string -> Numero 1 en binario representado en C2
#     - bin_num2: string -> Numero 1 en binario representado en C2
#     - reg     : int    -> Tamano del registro
# Retorno:
#     - boolean -> Si la suma en C2 produce overflow (True) o no (False).
def has_overflow(bin_num1, bin_num2, reg):
    of_status = False

    # Segun el tamaño de registro, calcular rango valido de valores en complemento 2
    max_range   = 2**(reg-1)-1
    min_range   = -2**(reg-1)
    valid_range = range(min_range, max_range+1)
    #print(f"Rango valido de valores: [{min_range}, {max_range}]\n{list(valid_range)}")

    # Se transforman a decimal y se suman normalmente
    dec_num1 = int(c2_to_dec(bin_num1))
    dec_num2 = int(c2_to_dec(bin_num2))
    suma     = dec_num1 + dec_num2
    #print(f"Suma: {dec_num1} + {dec_num2} = {suma}")

    # Se verifica si la suma se excede de los valores validos en complemento 2
    if suma not in valid_range:
        of_status = True
        #print("Hay overflow")

    return of_status
