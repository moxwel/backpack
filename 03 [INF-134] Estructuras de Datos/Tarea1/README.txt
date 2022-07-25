==========
README
==========

Nota: Los programas se crearon en un entorno de Linux, los archivos
      que se utilizaron para testeo tienen terminaciones
      de linea de tipo LF (\n).

==========

Problema 1:

- El archivo 'main.c' es el que posee el codigo principal, y el que deberia compilarse.

    - El programa se ha compilado con el siguiente comando: gcc main.c -o main -Wall

    - 'main' acepta 2 parametros opcionales:

        - El primero corresponde al archivo donde se ubican las palabras

        - El segundo corresponde al archivo donde se ubican los prefijos

        - Si no se especifican parametros, se utilizaran los valores predeterminados,
          que son "S.txt" y "P.txt" respectivamente.

          Ejemplo: ./main mis_palabras.txt mis_prefijos.txt

- El resto de archivos '.c' son dependencias de 'main.c', poseen el codigo de las funciones que se utilizaran.

Nota: Para testeo propio, se utilizó el comando ./main rae_completa.txt pref.txt

      'rae_completa.txt' es un archivo de 700.000 lineas aproximadamente, que contiene (casi) todas las palabras de la Real Academia Española.
      Sentimos que fue una buena forma de poner al limite nuestro programa.

      Fuente: https://github.com/JorgeDuenasLerin/diccionario-espanol-txt

==========

Problema 2:

- El archivo 'main.c' es el que posee el codigo principal, y el que deberia compilarse.

    - El programa se ha compilado con el siguiente comando: gcc main.c -o main -Wall

    - 'main' necesita 2 parametros, vease en el codigo.

- En la carpeta 'debug' hay codigos auxiliares que nos ayudó a testear la tarea:

    - 'generar_clientes.c' genera un archivo binario con datos de tipo 'clienteBanco'.

    - 'leer_clientes.c' puede leer ese archivo binario.

        - Este programa necesita 1 parametro: el archivo binario con los clientes a ser leidos.

- El resto de archivos '.c' son dependencias de 'main.c', poseen el codigo de las funciones que se utilizaran.

Nota: Para testeo propio, se utilizó el archivo binario generado por 'generar_clientes.c', y para los datos que se contienen, se utilizo
      'leer_clientes.c'.
