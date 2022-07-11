import re
'''
Modulo de Expresiones Regulares, utilizado
para buscar y eliminar caracteres especiales
de puntuacion de las letras de las canciones.

Informacion para este caso sacada de:

- https://docs.python.org/2/library/re.html

- https://www.tutorialspoint.com/How-to-remove-specific-characters-from-a-string-in-Python

- https://stackoverflow.com/questions/3939361/remove-specific-characters-from-a-string-in-python
'''

def reducer(dicc):
    '''
    Genera el archivo con los datos agrupados.
    '''

    print("\n\nGenerando archivo...")

    outfile = open("output.txt", "w", encoding="utf-8")
    # Como las palabras pueden contener archivos, de la misma forma
    # en como la abrimos, mencionamos el codificado a UTF-8. (Notese
    # al abrir archivos mas abajo).

    for key in dicc:
        outfile.write(("{}\t{}\n".format(key, dicc[key])))

    outfile.close()
    print("Done.")

def shuffle(lst):
    '''
    Cuenta las palabras de cada cancion
    y las pone en un diccionario.
    '''

    print("Reduciendo...")

    dicc = {}
    for itm in lst:
        dicc[itm] = dicc.get(itm, 0) + 1

    print("Conteo final:\n" + str(dicc))
    return dicc

def mapred(listarch):
    '''
    Lista todas las palabras que hay en cada cancion.
    Todo el resto de funciones se ejecutan desde aqui.
    '''

    # Aqui se juntaran todos los mapas al final de todo.
    todo = []

    # Para multiples archivos
    for arch in listarch:

        print("Mapeando archivo: " + arch)

        arch = open(arch,"r",encoding="utf-8")
        '''
        Como los archivos pueden contener acentos,
        se especifica que se utilice la codificacion
        UTF-8 al abrir el archivo.

        Informacion sacada de:
        - https://stackoverflow.com/questions/491921/unicode-utf-8-reading-and-writing-to-files-in-python

        - http://python-notes.curiousefficiency.org/en/latest/python3/text_file_processing.html
        '''

        for line in arch:

            line = line.strip()
            line = line.lower()
            # Estandarizar mayusculas y signos de puntuacion,
            # de esta forma no interfieren con el conteo.

            line = re.sub('[()?!¡¿.:;,\"-_]', '', line)
            # Aqui se utiliza el modulo "re", lo que hay entre
            # los corchetes son los caracteres que se buscan
            # en el string de la linea. Cuando los encuentra,
            # los reemplaza con un string vacio (eliminar).

            line = line.split()
            print(line)
            todo += line

        print("=====\nMapa final de archivo:\n" + str(todo) + "\n\n")
        arch.close()

    print("Fin de mapeo. Mapa total:\n" + str(todo) + "\n\n")
    return reducer(shuffle(todo))

# Toma la cantidad de archivos que se van a contar,
# deja de pedir archivos cuando se ingresa un cero.
listaArchivos = []
while True:
    inp = input("Archivos: ")
    if inp == "0":
        break
    listaArchivos.append(inp)
print("=====\nArchivos que se van a leer: " + str(listaArchivos) + "\n\n")

# La lista con todos los archivos a contar se pasa a la funcion.
mapred(listaArchivos)

'''
Para explicar el algoritmo creado, es necesario detallar
las funciones usadas durante el proceso.

Primero se define la funcion "reducer" que basicamente es el
archivo que nos da como salida, que cuando al abrirlo y a la vez crearlo,
tomamos la consideracion de UTF-8 que es un formato de codificación de
caracteres Unicode, el cual nos permite reconocer o tomar en cuenta
ciertos caracteres que el codigo ASCII no lo hace (las vocales con tilde)
para que el archivo creado al final sea visible.

Luego ahi mismo tenemos un ciclo for el cual recorre la llave del diccionario
que le entrego la funcion. Luego tenemos la funcion "shuffle" la cual como se
comenta el codigo, cuenta las palabras de cada cancion y
las coloca en un diccionario, y retorna ese mismo.

Finalmente tenemos la funcion "mapred", la cual lista las palabras que hay en cada
cancion (en el caso que se desee considerar varias canciones), donde ademas
se ejecutaran las funciones antes mencionadas.

En esa misma funcion al recorrer cada linea de cada archivo, para evitar
ciertas complicaciones, todas las palabras se pasan a minusculas, y si alguna
de ellas esta acompañada de un caracter especial como por ejemplo ()?!¡¿.:;,"-_
se remplaza con un string vacio, para evitar contar esa palabra como
diferente a otra simplemente por el hecho de tener ese signo.
==================
En el estudio de 3 albumes de reggaton (Oasis de Bad Bunny, Talento de Barrio
de Daddy Yankee y Prestige del mismo cantante), nos hemos fijado en que la mayoria
de palabras que se repiten, son monosilabos, como una forma distinta de pronunciar
palabras completas ("mas" como "ma'").

Ademas que la palabra "que" se repite una cantidad curiosa de veces (en total unas 709),
y a comparacion de albumes antiguos de Daddy Yankee, el uso de palabras
"obscenas" aumentó ligeramente.
'''