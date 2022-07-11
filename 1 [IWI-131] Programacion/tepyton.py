def hub(dia, meta, actual):
    print("Opcion 1. Ingresar nueva donacion")
    print("Opcion 2. Estado de cuentas")
    print("Opcion 3. Mayor donacion")
    print("Opcion 4. Avanzar al siguiente dia")
    print("DIA",dia)
    print("Meta:",meta)
    print("Llevamos:",actual)

meta = int(input("Cual sera la meta de esta TePyTon?: "))
dias = int(input("Ingrese la cantidad de dias de la TePyTon: "))
actual = 0
dia = 1

hub(dia,meta,actual)
mayor_rut = 0
mayor_donacion_ = 0

while dia <= dias:
    opcion = int(input("Ingrese operacion: "))

    if opcion == 1: # Ingresar nueva donacion
        rut = input("Ingrese rut (sin digito verificador): ")
        veri_rut = input("Ingrese digito verificador: ")
        donacion = int(input("Ingrese donacion: "))
        actual += donacion
    elif opcion == 2: # Estado de cuentas
        print("Faltan",meta-actual,"Vamos chilenos!!")
    # elif opcion == 3: # Mayor donacion
    #     #
    # elif opcion == 4: #Avanzar al siguiente dia
    #     #