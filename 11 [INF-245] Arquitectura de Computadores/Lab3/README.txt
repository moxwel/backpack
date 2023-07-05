Maximiliano Sepúlveda
Cristopher Fuentes
====================

README
------

MODO DE USO:

El programa son dos archivos .sv (testbench.sv y design.sv) los cuales deben
ser abiertos en la pagina web https://edaplayground.com/.

Una vez abiertos ambos archivos, se deberán ingresar las operaciones que se
deseen en formato HEXADECIMAL en el archivo design.sv (es posible agregar
hasta 64 operaciones).

Luego, se deberá 'inicializar' el programa mediante el botón 'Run'.

Una vez hecha la acción anterior, el programa mostrará por pantalla las
soluciones correspondientes a dichas 'entradas' (mediante el formato que
se definirá en el siguiente apartado).

El programa se detiene cuando encuentra una dirección sin instrucción definida.

EXPLICACION:

Este es un diseño de un circuito y un testbench el cual simula la ejecución
de las operaciones Suma, Resta, Bitshift Izquierda, Bitshift Derecha,
Composición, Disyunción, Exclusión y Negación mediante lenguaje de máquina,
utilizando Contador, Memoria, Splitter y ALU para finalmente entregar un
resultado en BINARIO.

La operación se hace en HEXADECIMAL, la cual se decodifica a BINARIO
descartando el primer bit, para luego desglosar en 3 partes: siguientes 3
bits son la operacion, los siguientes 16 son DOS números de 8 bits cada uno.

Las operaciones pueden ser:
	- 000 : Suma
	- 001 : Resta
	- 010 : Bitshift Izquierda
	- 011 : Bitshift Derecha
	- 100 : Composición
	- 101 : Disyunción
	- 110 : Exclusión
	- 111 : Negación.

La suma y la resta se hace mediante el uso de not's, and's y or's, y las
siguientes 6 se hicieron mediante las funcionalidades del Lenguaje Verilog.

El formato de salida es:
	Addr   Instr     Op      Num A        Num B        Res
	FF     FFFFF     000     00000000     00000000     00000000

Donde:
	Addr:	Es la dirección de memoria (0 a 64).
	Instr:	Es la instrucción que está dicha en la memoria.
	Op:		Es la operacion a realizar (Suma, resta, etc.)
	Num A:	Es el primer numero de la operación ingresada.
	Num B: 	Es el segundo numero de la operación ingresada.

COMPONENTES DEL CIRCUITO:

Contador: 	Circuito que se encarga de indexar la memoria usango 6
			bits para cubrir 64 espacios de las 64 posibles operaciones.


Memoria:	Almacena las instrucciones en cada espacio disponible.

Splitter:	Es quien separa las instrucciones en sus componentes fundamentales
			(Operacion y dos numeros en BINARIO).

ALU:		Circuito que realiza operaciones aritméticas.

Resultado:	Resultado de la operacion aritmetica a los numeros correspondientes.
			(Que luego se muestra en el formato de salida definido anteriormente)


SUPUESTOS:

1- Las operaciones son ingresadas en formato HEXADECIMAL.
2- Las operaciones deben ser ingresadas antes de ejecutar el codigo.
3- El programa se detiene cuando encuentra una dirección sin instrucción definida.
====================
