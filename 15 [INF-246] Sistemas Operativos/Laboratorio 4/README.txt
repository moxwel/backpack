    README

Maximiliano Sepúlveda
Antonia Zuñiga

====================

    ARCHIVOS

- depto.py : programa principal que ejecuta la simulación

====================

    INSTRUCCIONES PARA EJECUTAR

EL archivo 'depto.py' posee el codigo principal.

Para ejecutar, escribir en terminal:

$ python3 depto.py

====================

    SALIDA > TERMINAL

En la terminal se imprime las acciones que estan ocurriendo en el momento.

Cuando una persona recien entra a la universidad, se imprime junto a los departamentos a los
que debe ir:

    ==> Persona 1 INGRESA a UNIVERSIDAD [ Informatica, DEFIDER ]

En este punto, se encuentra en el Patio de las Lamparas. Luego, ingresa a la fila del patio (aún
no se encuentra dentro del departamento).

    Persona 1 LLEGA al PATIO de las lámparas para Informatica (500)

En este punto, la persona 1 esta esperando a entrar a la fila dentro del departamento. El numero
en parentesis representa a las personas actuale que aun no han sido atendidas.

Si la fila dentro del departamento tiene capacidad, la persona puede entrar.

    Persona 1 ENTRA a FILA de Informatica. (500)

La fila no avanza hasta que el departamento tenga las personas necesarias para llenar la capacidad.
Cuando esto ocurren, todas las personas necesarias ingresan al departamento, en orden.

    ==> Persona 22 INGRESA a UNIVERSIDAD [ Informatica, Quimica ]
    Persona 22 LLEGA al PATIO de las lámparas para Informatica (500)
    Persona 22 ENTRA a FILA de Informatica. (500)
        Persona 1 NO ATENDIDA. Iniciando proceso en Informatica. (500)
            =0=Persona 1 ENTRO al DEPARTAMENTO en Informatica. (500)
            =0=Persona 22 ENTRO al DEPARTAMENTO en Informatica. (500)
            ---Departamento Informatica ATENDIENDO... (500)

El departamento demora en atender. Luego de esto, las personas atendidas van al siguiente departamento
siguiendo el mismo proceso. Cuando la persona es atendida en ambas, se va de la universidad.

            @@@Departamento DEFIDER TERMINA de ATENDER. (495)
            =X=Persona 34 YA ha sido ATENDIDO en DEFIDER. (495)
            =X=Persona 1 YA ha sido ATENDIDO en DEFIDER. (494)
    <== Persona 1 SALE de UNIVERSIDAD [ Informatica, DEFIDER ]

====================

    SALIDA > ARCHIVOS

Los archivos de salida con los tiempos de atencion se generarán en la carpeta 'salida'.
