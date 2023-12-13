    README

Maximiliano Sepúlveda
Antonia Zuñiga

====================

    ARCHIVOS

- MainClasico.java     : programa principal que ejecuta el algoritmo de búsqueda de palabras de forma clásica
- MainFork.java        : programa principal que ejecuta el algoritmo de búsqueda de palabras usando Fork/Join
- MainThread.java      : programa principal que ejecuta el algoritmo de búsqueda de palabras usando hebras
- Sopa.java            : contiene las funciones para manipular la sopa de letras
- TiempoAlgoritmo.java : contiene las funciones para medir el tiempo de ejecución de los algoritmos

====================

    INSTRUCCIONES PARA COMPILAR

El código 'MainClasico.java', 'MainFork.java' y 'MainThread.java' contienen el código principal.

Para compilar, ejecutar en terminal:

$ make

Se generará un ejecutable para cada uno de los archivos principales.

Cada programa recibe la direccion del archivo que contiene la sopa de letras:

$ java MainClasico <archivo_sopa>
$ java MainFork <archivo_sopa>
$ java MainThread <archivo_sopa>

    Por ejemplo: $ java MainClasico "./prueba/CHECOESLOVAQUIA/sopa_de_letras.txt"

Para ejecutar con todos los archivos dentro de la carpeta 'prueba', ejecutar:

$ make runClasico
$ make runFork
$ make runThread

====================
