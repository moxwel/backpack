Maximiliano Sepúlveda
Cristopher Fuentes
====================

README
------

MODO DE USO Y PASOS:
El programa principal "convierte.py" no necesita parametros adicionales,
sin embargo, se necesita un archivo de entrada: "numeros.txt" el cual debe
estar en la misma carpeta que el programa.

1- Correr el programa:
    - Para correr el programa, se puede utilizar el método que estime conveniente
    el usuario. (Ej: Abrir una terminal en la carpeta en la cual esté el programa
    junto al archivo "numeros.txt").

        Ejemplo: python convierte.py

    NOTA: Una vez ejecutado el programa, el archivo "resultados.py" se limpiará.

2- Agregar un tamaño de registro:
    - Una vez el programa esté corriendo, este le pedira al usuario multiples tamaños
    de registro para poder hacer los cálculos correspondientes.

    - Una vez se ingrese el tamaño "0", si el total de errores supera a la cantidad de
    números, finaliza el programa. En caso contrario, pide nuevamente un tamaño de
    registro para continuar con su ejecución.

3- Salidas:
    - A medida que el programa funcione, se va generando un archivo "resultados.txt"
    el cual tiene por cada linea, el resultado de cada tamaño de registro ingresado.




EXPLICACION:
El programa esta descompuesto en funciones que realizan operaciones básicas
sobre números, tanto para transformarlos en diferentes bases o realizar
comprobaciones en ellos.

Dichas funciones se encuentran en el archivo "num_functions.py", que es
importado al programa principal para su uso.




ALGORITMOS:
Se han diseñado los siguientes algoritmos (num_functions.py):

    Conversion:
    - Base X a base 10.
    - Base 10 a base 2.
    - Binario C2 a base 10.

    Comprobacion:
    - Numero valido segun base.
    - Numero cabe en registro indicado.
    - Overflow con suma en C2.

    Manipulacion:
    - Extension de signo.
    - Invertir binario.
    - Suma binaria basica.

La comprobación de overflow se diseño usando un enfoque mas
analitico, considerando los rangos de valores segun el tamaño de registro:

    Sea un registro de N digitos, los valores válidos en C2 son: [ -2^(N-1) , 2^(N-1)-1 ]

Es decir, si se tiene un registro de tamaño 4, el rango de valores válidos en C2 seria [-8,7].

Podemos transformar ambos numeros binarios C2 a decimal, y luego sumarlos normalmente.
Si dicha suma esta fuera del rango válido de valores, entonces significa que ocurrio overflow.




SUPUESTOS:
    - Por cada linea del archivo "numeros.txt", se tienen 2 números con sus respectivas bases.
    - Las bases en "numeros.txt" estan en el rango entre 2 hasta 16.
    - Los numeros en "numeros.txt" no tienen representacion negativa.

====================
