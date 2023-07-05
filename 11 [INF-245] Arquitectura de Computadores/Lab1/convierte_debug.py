from num_functions import *

A_global  = 0
B_global  = 0
C_global  = 0
D_global  = 0
total_err = 0

# Limpiar resultados.txt
x = open("resultados.txt","w")
x.close()

running =  True
while running:
    valid_reg = False
    while not valid_reg:
        tamano_registro = int(input("Ingrese tamano de registro: "))
        if tamano_registro in range(0,33):
            valid_reg = True
        else:
            print("Tamano no valido.")


    print("======")

    if tamano_registro != 0:
        A = 0
        B = 0
        C = 0
        D = 0

        # Comienza a recorrerse el archivo...
        cont = 1
        archivo = open("numeros.txt","r")
        for line in archivo:
            print(f"LINEA {cont}\n")

            # Hay dos numeros por linea
            A += 2

            # Sanitizar string. Ej: 10;12-16;F ---> ['10','12','16','F']
            line = line.strip().replace(";"," ").replace("-"," ").split(" ")

            # Verificaciones al primer numero
            base1      = int(line[0])
            num1       = str(line[1])
            valid_num1 = is_valid(base1, num1)
            print(f"Numero 1: {num1} en base {base1} (Valido: {valid_num1})")

            if valid_num1:
                dec_num1 = base_to_dec(base1, num1)
                bin_num1 = dec_to_bin(dec_num1)
                fit_num1 = fits_reg(bin_num1, tamano_registro)
                print(f"  Decimal: {dec_num1}\n  Binario: {bin_num1} (Cabe en {tamano_registro} bits: {fit_num1})")

                if not fit_num1:
                    valid_num1 = False
                    C += 1
                    print("[!] Error de registro. +1 en C")
            else:
                B += 1
                print("[X] Error de representacion. +1 en B")

            print("")

            # Verificaciones al segundo numero
            base2      = int(line[2])
            num2       = str(line[3])
            valid_num2 = is_valid(base2, num2)
            print(f"Numero 2: {num2} en base {base2} (Valido: {valid_num2})")

            if valid_num2:
                dec_num2 = base_to_dec(base2, num2)
                bin_num2 = dec_to_bin(dec_num2)
                fit_num2 = fits_reg(bin_num2, tamano_registro)
                print(f"  Decimal: {dec_num2}\n  Binario: {bin_num2} (Cabe en {tamano_registro} bits: {fit_num2})")

                if not fit_num2:
                    valid_num2 = False
                    C += 1
                    print("[!] Error de registro. +1 en C")
            else:
                B += 1
                print("[X] Error de representacion. +1 en B")

            print("")

            # Si ambos numeros no tienen problemas, se procede a la suma y se verifica overflow
            if valid_num1 and valid_num2:
                if not has_overflow(bin_num1, bin_num2, tamano_registro):
                    print("La suma no produce overflow :)")
                else:
                    D += 1
                    print("[*] La suma produce overflow. +1 en D")
            else:
                print("[-] Uno de los numeros tiene error. No se suma.")

            print("---")
            cont += 1

        # Resultados obtenidos (se escribe en archivo resultados.txt)
        print(f"Resultados: A={A} B={B} C={C} D={D}")

        # Se guardan resultados para la siguiente iteracion
        A_global = A
        B_global += B
        C_global += C
        D_global += D

        # El total de errores se guarda en el caso que supere la cantidad de numeros
        total_err = B_global + C_global + D_global
        print(f"Errores totales hasta ahora: B={B_global} C={C_global} D={D_global} (Total: {total_err})\n======")

        archivo.close()

        # Escribir resultados en resultados.txt
        archivo_res = open("resultados.txt","a")
        archivo_res.write(f"{A};{B};{C};{D}\n")
        archivo_res.close()

    else:
        # Se ingreso 0, se verifica la cantidad de errores hasta ahora.
        print(f"Total numeros en archivo: {A_global}\nTotal errores hasta ahora: {total_err}\n======")
        if total_err > A_global:
            print("Total de errores supera a cantidad de numeros. Terminar.")
            running = False
