==========
README TAREA 1 ALGORITMOS Y COMPLEJIDAD
==========



==========
NOMBRES
==========

- Tomás Vasquez
  ROL: #########-#

- Maximiliano Sepúlveda
  ROL: #########-#

- Cristopher Fuentes
  ROL: #########-#



==========
INFORMACION
==========

Para la ejecucion de los programas, se utilizó Python 3.8.6.

1-. Joins en Bases de Datos
    En este apartado, contamos con dos opciones a tener en cuenta:

    1a-. 1_join_file.py
         Este programa ".py" requiere de un archivo ".txt" como entrada (si no ingresamos ningun
         archivo, el programa explicará como es que se debe insertar un archivo).
         En ese archivo de texto, debe existir dos relacionales (una encima de otra)
         escritas de la siguiente forma:

             N             (Número entero entre 1 y 100, que indica la cantidad de atributos de la relación)
             0 1 2 ... N   (N numeros enteros, separados por un unico espacio en blanco, los cuales indican los atributos de la relación)
             NR            (Numero entero entre 1 y 10.000, que indica la cantidad de tuplas de la relación)
             0             (Las siguientes líneas (desde aquí hacia abajo) son NR líneas de texto,
             1             cada una con N valores enteros separados por un espacio
             2             que corresponden a las tuplas de la relación R).
             ... NR
             N             (Aquí comienza la segunda relación)
             0 1 2 ... N
             NR
             0
             1
             2
             ... NR

         De esta forma, el programa mostrará el resultado de la operación "join" sobre las dos
         relaciones dadas anteriormente (ese resultado tiene el mismo formato que el archivo de texto).



    1b-. 1_join_man.py
         Este programa ".py" requiere una inserción manual, es decir, entregar los datos uno por
         uno con el mismo formato mostrado anteriormente en el punto "1a". De igual forma,
         el programa entregará un resultado con el mismo formato entregado y será el efecto de
         hacer la funcion "join" a ambas relaciones.



2-. Contando Triángulos en Grafos
    Al igual que en la parte anterior, tendremos dos opciones:

    2a-. 2_triangulos_file.py
         Este es el programa el cual se requiere un archivo ".txt" el cual en su interior
         solo debe estar representado un grafo, la forma en la cual se debe escribir es
         representando un arco del grafo por linea, con dos valores enteros separados por un
         unico espacio (vease las notas al final de este documento).

         El resultado del programa es una unica linea que indica la cantidad de triángulos existentes en el grafo.



    2b-. 2_triangulos_man.py
         En este programa, se requiere ingresar los arcos del grafo que se quiera representar,
         de la misma forma explicada anteriormente. Pero, para poder finalizar la representacion
         del grafo hay que escribir al final "EOF" de manera literal (esto terminaria el proceso
         de ingreso de datos para que el programa te indique el resultado).

         Al igual que en la parte anterior, el resultado es una linea unica la cual indica la
         cantidad de triángulos existentes en el grafo



==========
NOTAS
==========

1-. Agregamos también una carpeta llamada "testing", la cual contiene algunos casos de prueba para
    los programas que requieran ".txt", estos tambien pueden ser utilizados para las versiones manuales,
    pero como su nombre lo indica, deben ser agregados manualmente por el usuario que corra el programa.
    Puede guiarse con esos archivos respecto al formato de entrada.

2-. Además de Python 3.8.6, tambien se probó en Python 3.9.4.
