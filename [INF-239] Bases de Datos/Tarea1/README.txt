README
----------------------
Maximiliano Sepúlveda
ROL: 201973536-5
----------------------

Por errores de sintaxis que Oracle me alegaba y no supe como lograr solucionarlos a tiempo,
no pude hacer todos los triggers y procedimientos que se pideron u.u ...

Un archivo llamado 'tests.sql' contiene codigo incompleto como intentos para solucionar esos
problemas.





IMPORTACION DE DATOS
--------------------
Se utilizó Oracle Database 18c XE y SQL Developer.

No es necesario crear ningun usuario en especifico. En teoria deberia ser "llegar y ejecutar".

Se utilizó la cuenta 'system' de la base de datos, asegurese de que la conexion a la BD en
SQL Developer esté utilizando dicha cuenta.

Utilizando SQL Developer, se deben ABRIR (Ctrl + O) los diferentes archivos SQL,
y EJECUTAR EL SCRIPT COMPLETO (F5) de cada uno de ellos.

Por amor al orden, distribui los comandos en 3 archivos:

- crear_datos.sql
  Posee las tablas, procedimientos y triggers a crear.

- borrar_datos.sql
  Borra las tablas, procedimientos y triggers.
  
- datos_prueba.sql
  Datos de prueba que utilicé en las tablas.

Recomiendo ejecutar primero 'borrar_datos.sql', luego 'crear_datos.sql', y finalmente
'datos_prueba.sql'. Si todo sale bien, todo deberia funcionar debidamente (al menos para mi
si funciono D:)





CARPETA ETC
-----------
Es la carpeta de pruebas y herramientas. Creé una plantilla de Excel para poder organizar
las tablas de mejor manera, y automatizar un poco la creacion de las consultas INSERT.
Decidi dejarlo en la entrega por si desea ver los datos que posee.

El archivo 'datos_prueba.sql' posee todas las consultas que ese archivo Excel tiene,
estan puestas en un orden especifico para respetar las dependencias de las claves foraneas.





NOTAS AL MARGEN
---------------
Algunas suposiciones:

- Los alimentos no diferencian por especie de mascota. Por ejemplo en los casos de prueba: existe un
  alimento llamado 'Dog Chau' y 'Cat Chau', pero el modelo de datos permite que un alimento
  llamado 'Cat Chau' sea recomendado a una mascota de especie 'perro'.
  
  - Supongo que esto se puede solucionar añadiendo un campo mas en la tabla 'alimentacion', para
    especificar "a que especie de mascota esta destinado tal alimento".



- Un alimento para una mascota es recomendable si la "edad recomendada" de tal alimento es menor
  que la edad de la mascota. Para mi tiene mas sentido, aunque podria darse el caso que una
  mascota adulta se le recomiende un alimento recomendada para una edad mucho mas temprana.
  
  - Por esta razon, siempre se va a recomendar el alimento con mayor edad recomendada que a la vez
    sea menor a la edad de la mascota. Esto causará que un mismo alimento se recomiende siempre
    a partir de cierta edad, a causa del 'trigger_alimentacion'.



- En 'vista_consumo' en vez de considerar los alimentos usados por mas de 5 dietas, decidi dejarlo
  con "mayor o igual" a 5 dietas. Esto a causa de la poca cantidad de datos que tengo, y para ver
  si realmente funcionaba.



- En 'procedure_entrenamiento', asumí que "dada una habilidad" en realdiad se refiere a la ID de
  la habilidad (es decir, al numero de 4 digitos de la columna 'id_entrenamiento'). Se me hace
  mas sentido y mas seguro que utilizar un string.
