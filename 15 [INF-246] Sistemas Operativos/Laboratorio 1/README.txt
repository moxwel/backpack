    README

Maximiliano Sepúlveda
Antonia Zuñiga

====================

    ARCHIVOS

- main.c      : programa principal que ejecuta las tareas.

- sopa.c      : funciones utiles para resolver sopa de letras.

- carpetas.c  : funciones utiles para manipular archivos y carpetas.

- Makefile    : instrucciones para compilar usado por 'make'.

- ./archivos/ : carpeta donde estarán los archivos de sopa de letras.

- ./testing/  : carpeta con codigo de prueba (no se utiliza).

====================

    INSTRUCCIONES

El codigo 'main.c' es quien contiene el codigo principal.

Para compilar, ejecutar en terminal:

$ make

Se generará un ejecutable llamado 'sopa_letras.out'. Para ejecutar:

$ ./sopa_letras.out

====================

    FUNCIONAMIENTO

El programa listará todos los archivos que se encuentren en la carpeta 'archivos' e ira
resolviendo cada una.

Primero se lee el archivo, y la sopa de letras será guardado en memoria en una matriz. Esta
matriz se utilizará para resolver la sopa de letras.

En cada resolucion, se toma el tiempo de ejecucion solamente del algoritmo de busqueda.

Dependiendo de las caracteristicas de este, dicho archivo se moverá a una nueva carpeta.

Cada archivo provoca una salida en consola. Ejemplo:

        Numero de archivos: <12>
        Creando carpetas...

        lista_archivos[0]:
        - Ruta archivo: ./archivos/banco.txt
        - Palabra a buscar: BANCO
        - Es horizontal (1=si, 0=no): 1
        - Dimension: 200
            [!!!] Encontrado en 177:24
        - Ticks: 141 [ticks]
        - Tiempo: 0.000141 [s]
        - Palabra encontrada
        Moviendo archivo ./archivos/banco.txt a ./archivos/horizontal/200x200/banco.txt

        lista_archivos[1]:
        - Ruta archivo: ./archivos/Carne.txt
        - Palabra a buscar: CARNE
        - Es horizontal (1=si, 0=no): 0
        - Dimension: 200
            [!!!] Encontrado en 182:100
        - Ticks: 67 [ticks]
        - Tiempo: 0.000067 [s]
        - Palabra encontrada
        Moviendo archivo ./archivos/Carne.txt a ./archivos/vertical/200x200/Carne.txt

====================

    CONSIDERACIONES

La carpeta 'archivos' es donde ocurre toda la manipulacion de archivos. Los nuevos textos
con sopas de letras deben ser añadidos a esa carpeta.

Se asume que los nombres de los archivos no seran muy largos (menos de 100 caracteres).

Se asume que todas las sopas de letras son cuadradas.

Se asume que las palabras solo pueden estar en horizontal o vertical.

Se asume que las palabras no van a estar al reves.

El limite maximo de las dimensiones de la sopa de letras es de 200x200.
