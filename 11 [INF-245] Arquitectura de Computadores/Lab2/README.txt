Maximiliano Sepúlveda
Cristopher Fuentes
====================

README
------

MODO DE USO:

El programa es un archivo .circ el cual deberá ser abierto por el ejecutable
'Logisim'.

Una vez abierto, se deberá 'inicializar' el reloj.

Luego, es necesario asignar valores a los puntos H y V (esto es posible
haciendo clic con el mouse en los '0' o '1' de los mismos), para utilizar
nuevamente el reloj. Este nos dirá si estamos o no en un punto perteneciente
al laberinto (si nos arroja un 'Estado sig.' distinto de '00000' es porque
estamos en un punto perteneciente).

Una vez se encuentre un punto perteneciente, sólo será necesario hacer clic
tantas veces como sea necesario hasta llegar al punto final.

EXPLICACION:

Este es un diseño de un circuito que simula un laberinto, el cual tiene 31
posibles posiciones más un punto '0' que sería un estado 'comodín' que no
cambiaría hasta tener un punto perteneciente a un punto válido del laberinto.
Todos los puntos estan compuestos por H (punto en horizontal) y V (punto en
vertical), los cuales son representados por 3 bits cada uno (0 a 7), y las
posibles posiciones son representados por 5 bits (ya que son 32 estados).
La máquina recibe 6 bits (de los puntos H y V) y es capaz de identificar en
el punto que se ubica, ya sea en el estado 0 (punto fuera del laverinto) o
un punto que corresponde al mismo. Una vez que se encuentre un punto perteneciente
al laberinto, se pasa a utilizar unicamente los estados ya que con estos se
conoce el siguiente estado al que se 'avanzará' para finalmente llegar al
'punto final' del laberinto.

COMPONENTES DEL CIRCUITO:

Lógica primer estado: 	Recibe un punto perteneciente al estado 'q_0' en logica
			de H y N para entregar un estado 'q_N' siempre y cuando
			el punto (H,N) pertenezca a un estado valido del laberinto

Lógica sgte estado:	Recibe un estado 'q_n' perteneciente al laberinto y
			entrega un estado 'q_(n+1)' el cual acerca o es el
			'punto final' del laberinto.

MUX:			Elige las salidas de los bloques 'Lógica primer estado' o
			'Lógica siguiente estado' dependiendo si estamos en un punto
			'validado' (estado perteneciente al laberinto) o si estamos
			en un punto que necesita ser validado (estado q_0).

CLOCK:			Señal de reloj que hace actuar a los Flip flops en la
			transición de 0 a 1.

Lógica de salida:	Dado un estado, la salida corresponde a la dirección del
			siguiente estado para encontrar la meta del laberinto.



SUPUESTOS:

La entrada inicial siempre es un punto válido dentro del laberinto.
====================
