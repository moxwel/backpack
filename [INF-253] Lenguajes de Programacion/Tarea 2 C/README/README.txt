README
Maximiliano Sepulveda Alvear
ROL 201973536-5
====================

Este es el Readme principal del programa. La documentacion de la lista enlazada
y de los Transimperativers estan en otros archivos aparte en esta misma carpeta.
Esto con el fin de ser mas directo y que no sea solo un archio TXT con mas de 500 lineas xddd





NOTAS AL MARGEN
---------------

Mi implementacion del arreglo de Transimperativer fue hecho a ultima hora, y no pude hacerlo
de forma dinamica :c, tiene un maximo de 100000 Transimperativer que puede guardar. Aunque
supongo que haran explotar mi programa añadiendo 1 gugol de Transimperativers T-T

Para las funciones de mezcla y separacion, se tienen placeholders dentro de main, como 'separar2',
'mezclar3', etc. Esas funciones estan vacias y no hacen nada. En teoria, se pueden rellenar con
cualquier cosa; y seran guardados en los arreglos de funciones correspondientes.

El arreglo de funciones y el arreglo de transimperativers, ambos comienzan con el indice 0,
es decir, al elegir el Transimperativer 0, se esta eligiendo el primero en la lista.

¡¡¡CUDADO CON EL DANGLING!!!
Cuando estaba probando la funcion de borrar Transimperativer, esta utiliza la funcion de 'deleteList()',
quien a su vez usa la funcion 'erase()', y esta ultima borra un nodo de la lista enlazada,
como es de esperarse, pero tambien libera la memoria que esta contenida en el puntero 'ínfo' (caracteristica).
Por esta razon, cree otra funcion llamada 'removeNode()', que hace lo mismo que 'erase()',
pero sin liberar la memoria del puntero 'info', justamente para evitar el dangling.
Notese la funciones de 'MezclarMitad' y 'SepararMitad' para ver el caso de uso de 'removeNode()'.
PD: Tambien aplica a la funcion 'ReemplazarCuerpo()'.

Los comentarios de las funciones del TDA se encuentran en los archivos ".h", en IDEs como VSCode
se reconoce el header, y vincula los comentarios, entonces al pasar el mouse encima del nombre
de una funcion, se puede ver la descripcion de este. De todas formas, para una explicacion mas
exhaustiva de las funciones, vease la documentacion de lista enlazada.

El archivo 'tools.c' y 'tools.h' contienen funciones utiles y auxiliares que se utilizan
en el resto del programa, por ejemplo, calcular promedios del contenido de un arreglo.

La dependencia de los archivos es la siguiente:

        main.c
        ├ stdio.h
        └ Transimperativer.h
          └ listaEnalzada.h
            ├ stdio.h
            ├ stdlib.h
            └ tools.h
              └ stdio.h





COMPILACION Y EJECUCION
-----------------------

Para compilar el programa, se debe ejecutar el comando make.

    $ make
    Compila el programa, generando un ejecutable 'programa.out'.

    $ make start
    Compila el programa, y lo ejecuta.

    $ make test
    Compila el programa, y lo ejecuta en Valgrind.





MENU
----

Al iniciar el programa, se mostrara el menu principal:

        MAIN MENU
        =========
        Transimperativers creados: 0/100000

        Acciones:
            1. Crear Transimperativer
            2. Mostrar Transimperativer
            3. Agregar atributo
            4. Quitar atributo
            5. Pelear
            6. Mezclar
            7. Separar
            8. Borrar Transimperativer
            0. Detener programa

Se debe elegir la accion a realizar a continuacion.

        Seleccionar accion: 1

        [CREAR TRANSIMPERATIVER]
        Elegir funcion de mezcla [0-3]: 0
        Elegir funcion de separacion [0-3]: 0
        Listo.

Los elementos entre corchetes pueden ser los elemntos que puede recibir.
Despues de agregar un atributo, se muestra el Transimperativer modificado
con sus cambios.

        Seleccionar accion: 3

        [AGREGAR ATRIBUTO]
        Elegir transimperativer [0-0]: 0
        Tipo del atributo [i, c, h, a]: a
        Añadiendo Arreglo de Ints. Ingrese tamaño del arreglo: 4
        Ingrese elemento en indice 1: 23
        Ingrese elemento en indice 2: 124
        Ingrese elemento en indice 3: 23
        Ingrese elemento en indice 4: 64
        -- TRANSIMPERATIVER --
        Len: 1
        Lista de caracteristicas:
        [0   ][a] Arreglo de Int  | Size: 4 = [ 23, 124, 23, 64 ] = Avrg: 59
        ======================
        Listo.

Para el formato del cuerpo del Transimperativer, vease la documentacion de
lista enlazada.

ADVETENCIA: No me dio el tiempo de hacer verificacion de tipo al ingresar datos.
Si pones un numero en el lugar donde debes elegir el tipo a ingresar por ejemplo,
es muy probable que explote.

        Tipo del atributo [i, c, h, a]: 5       // Por favor, no hagas eso :c
