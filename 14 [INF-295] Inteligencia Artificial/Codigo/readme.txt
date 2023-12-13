    README




    COMPILACION

Para compilar, ejecutar "make" en la terminal.

    $ make

Para ejecutar, ejecutar "./main" en la terminal. Este recibe como argumento el nombre del archivo de entrada.

    $ ./main "./instancias/6-10/9.txt"

Para automatizar la ejecucion del programa con multiples instancias, se puede utilizar el script "run.sh".

    $ ./run.sh "./instancias/6-10/"

El script recibe como argumento la carpeta que contiene las instancias de entrada.
Además, creará una carpeta con el nombre del tamaño de la instancia (en este caso '6-10') y guardará los archivos de salida en ella.

*NOTA: Antes de ejecutar el script, ejecutar 'make' para compilar el programa.

*NOTA: Si el script no tiene permisos de ejecucion, ejecutar "chmod +x run.sh" en la terminal.






    SALIDA > ARCHIVOS

Se generará un archivo de texto en la carpeta 'output'.

El nombre corresponde al numero de la instancia de entrada y la solucion encontrada. Por ejemplo: 9_000100001.txt

El formato de salida es el siguiente:

    ALTURA_USADA
    AREA_DESPERDICIADA
    X_1 Y_1 R_1
    X_2 Y_2 R_2
    ...
    X_N Y_N R_N

Ademas del archivo de salida, se genera un archivo 'results.tsv' con estadisticas de la ejecucion del programa.







    SALIDA > TERMINAL

El programa imprime en la terminal el estado inicial y final de los bloques.

    ======LISTA BLOQUES======
    ID  R   W   H   A   U   X   Y
    1   0   2   1   2   0   0   0
    2   0   8   4   32  0   0   0
    3   0   9   2   18  0   0   0
    4   0   9   1   9   0   0   0
    5   0   3   6   18  0   0   0
    6   0   10  8   80  0   0   0
    7   0   8   4   32  0   0   0
    8   0   6   7   42  0   0   0
    9   0   1   5   5   0   0   0
    10  0   8   2   16  0   0   0
    ====FIN LISTA BLOQUES====

- ID: Identificador del bloque
- R: Rotacion del bloque
- W: Ancho del bloque
- H: Alto del bloque
- A: Area del bloque
- U: Uso del bloque
- X: Posicion X del bloque
- Y: Posicion Y del bloque

*NOTA: Las posiciones X e Y se consideran desde la esquina inferior izquierda del bloque.

Tambien imprime estadisticas de la ejecucion del programa:

    =====RESULTADO FINAL=====
    Mejor resultado: 34
    Mejor instancia: 0000100010
    Mejor iteracion: 3
    Total de iteraciones: 4
    Llamadas a la funcion: 25
    Tiempo de ejecucion: 0.058 ms
