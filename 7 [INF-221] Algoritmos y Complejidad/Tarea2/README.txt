==========
README TAREA 2 ALGORITMOS Y COMPLEJIDAD
==========


==========
NOMBRES
==========

- Tomás Vasquez


- Maximiliano Sepúlveda


- Cristopher Fuentes




==========
INFORMACION
==========

Para la ejecucion de los programas, se utilizó Python 3.8.6.

1-. StrassenVSTradicional_auto.py

    Este programa ".py" no requiere de algo externo a Python, por lo que, solo te pedirá una entrada
    al momento de ejecutarlo. Esta entrada será una insersión manual, la cual será la dimensión "N"
    de las matrices que se multiplicaran, estas matrices serán completamente aleatorias, pero al
    momento de multiplicarlas, se usarán las mismas para el algoritmo de Strassen y el Tradicional.

    Lo que entrega este programa, será:

        a-. Matrices iniciales: estas serán las matrices que se utilizarán para los calculos
            posteriores, como hemos dicho, son totalmente aleatorias, pero se utilizan LAS MISMAS
            para ambos algoritmos.

        b-. Dimensión de la matriz resultante: esta será la dimensión de la matriz resultante
            del algoritmo de Strassen.

        c-. Codigo reducido (Strassen): esto indicará el tiempo en el cual fue hecha la multiplicación
            mediante el algoritmo de Strassen.

        d-. Matriz resultante: está será la matriz resultante que se generó mediante el algoritmo de
            Strassen

        e-. Dimensión de la matriz resultante: esta será la dimensión de la matriz resultante
            del algoritmo Tradicional.

        f-. Codigo Tradicional: esto indicará el tiempo en el cual fue hecha la multiplicación
            mediante el algoritmo de Tradicional.

        g-. Matriz resultante: está será la matriz resultante que se generó mediante el algoritmo
            Tradicional.

2-. StrassenVSTradicional_man.py

    Este programa es identico al anterior, con la diferencia que la entrada en este caso es manual.
    Se puede realizar tanto "a mano" como por redirección de entrada.

    Por ejemplo:

        python StrassenVSTradicional_man.py < entrada.txt



==========
NOTAS
==========

1-. Para dimensiones 'N' de matrices demasiado grandes, al momento de mostrar ("printear") las
    matrices con las cuales se harán las multiplicaciones y las matrices resultantes, al ser de
    un tamaño muy grande, hará que no se muestre gran parte de las matrices multiplicando y
    multiplicador (por llamarlos de alguna manera) ni tampoco el tiempo de ejecución de las operaciones.

    Para esto, hay que ir al código y comentar algunas lineas, las cuales estan marcadas con un comentario
    a su lado que dice lo siguiente: "Si tu matriz es muy grande, comenta esta linea."

2-. No tenemos hecho un "MAKEFILE" ya que, no lo encontramos necesario debido a que solo utilizamos un
    programa ".py"

3-. Dentro de la carpeta "tools", hicimos un programa llamado 'generadorMatrices.py', con el fin de
    ayudarnos a generar una entrada valida para los programas. Allí tambien se encuentran archivos de prueba.
