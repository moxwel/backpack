==========
README
==========

Nota: Los programas se crearon en un entorno de Linux, los archivos
      que se utilizaron para testeo tienen terminaciones
      de linea de tipo LF (\n).

Nota: Por temas de tiempo (y porque nos aniquilaron en un certamen
      de fisica y mate a la vez (porfavor saquenme de esta pesadilla
      llamada USM)), admitimos que no alcanzamos a implementar una
      estructura de tipo HEAP u_u En vez de eso, usamos otro algoritmo (tal vez
      muy ineficiente) que al final funciona igual.

==========

Problema 1:

- El archivo 'main.c' es el que posee el codigo principal, y el que deberia compilarse.

    - El programa se ha compilado con el siguiente comando: gcc main.c -o main -Wall

    - 'main' necesita 3 parametros:

        - El nombre de un archivo binario con los productos a registrar.

        - El nombre de un archivo binario con las ofertas a registrar.

        - El nombre de un archivo de texto con las compras a realizar de cada cliente.

          Ejemplo: ./main productos.dat ofertas.dat compras.txt

- El resto de archivos '.c' son dependencias de 'main.c', poseen el codigo de las funciones que se utilizaran.

(PD 1)
La carpeta "debug" es un lugar donde dejamos algunos programas de utilidad para generar
los archivos binarios que se nos pedian, y usarlos como casos de prueba (aunque son chikitos).
Lo consideramos como un "borrador", no se toma en cuenta y "main.c" no los usa.

(PD 2)
Agradecimientos a un pibe que nos compartio sus casos de prueba uwu

(PD 3)
Esos momentos cuando el Heap se saca uno de esos que te dejan terrible ordenado kieee pero como
