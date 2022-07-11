"""
ROL #########-#_#########-#_#########-#
PAR 205
NOM CristopherFuentes_FelipeMellado_MaximilianoSepulveda
"""

def precio_entradas(dias,adultos,menores):
    precioTotal = 0

    i = 0
    while i < dias: # se va a repetir "dias" veces

        sem = input("Dia de la semana: ") # dia de la semana
        if sem == "Lunes" or sem == "Martes" or sem == "Miercoles" or sem == "Jueves" or sem == "Viernes":

            if menores > 0: # si hay niños...
                if adultos + menores >= 4: # ...y si hay 4 o mas personas...
                    precioTotal = precioTotal + (adultos*20) + (menores*8) # ...se aplica descuento.
                else:
                    precioTotal = precioTotal + (adultos*20) + (menores*10) #...no se aplica descuento.
            else: # si no hay niños...
                precioTotal = precioTotal + (adultos*20)

            precioTotal += precioMenuDelDia(adultos,menores)

        elif sem == "Sabado" or sem == "Domingo": # en fin de semana el precio es el doble

            if menores > 0: # si hay niños...
                if adultos + menores >= 4: # ...y si hay 4 o mas personas...
                    precioTotal = precioTotal + (adultos*40) + (menores*16) # ...se aplica descuento.
                else:
                    precioTotal = precioTotal + (adultos*40) + (menores*20) #...no se aplica descuento.
            else: # si no hay niños...
                precioTotal = precioTotal + (adultos*40)

            precioTotal += precioMenuDelDia(adultos,menores)

        i += 1

    return precioTotal


def precioMenuDelDia(adultos,menores):
    precioMenuMenor = menores * 11 # se calcula el precio de los niños
    precioMenuAdulto = 0

    i = 1
    while i <= adultos: # se repite por cantidad de adultos
        tipoMenu = input("Menu Adulto " + str(i) + " :")
        if tipoMenu == "A":
            precioMenuAdulto += 15
        elif tipoMenu == "B":
            precioMenuAdulto += 18
        i += 1

    precioTotal = precioMenuMenor + precioMenuAdulto
    return precioTotal

dias   = int(input("Cantidad de dias: "))
c_adul = int(input("Cantidad de adultos: "))
c_meno = int(input("Cantidad de menores: "))

print(precio_entradas(dias,c_adul,c_meno))
