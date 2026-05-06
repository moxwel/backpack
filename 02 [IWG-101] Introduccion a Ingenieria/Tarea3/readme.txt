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

===

En el estudio de 3 albumes de reggaton (Oasis de Bad Bunny, Talento de Barrio
de Daddy Yankee y Prestige del mismo cantante), nos hemos fijado en que la mayoria
de palabras que se repiten, son monosilabos, como una forma distinta de pronunciar
palabras completas ("mas" como "ma'").

Ademas que la palabra "que" se repite una cantidad curiosa de veces (en total unas 709),
y a comparacion de albumes antiguos de Daddy Yankee, el uso de palabras
"obscenas" aumentó ligeramente.
