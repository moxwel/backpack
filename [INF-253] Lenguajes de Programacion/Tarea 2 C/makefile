all: programa.out

programa.out: listaEnlazada.o Transimperativer.o tools.o main.o
	gcc listaEnlazada.o Transimperativer.o tools.o main.o -o programa.out

listaEnlazada.o: listaEnlazada.c listaEnlazada.h tools.h
	gcc -Wall -c listaEnlazada.c

Transimperativer.o: Transimperativer.c Transimperativer.h listaEnlazada.h
	gcc -Wall -c Transimperativer.c

tools.o: tools.c tools.h
	gcc -Wall -c tools.c

main.o: main.c Transimperativer.h
	gcc -Wall -c main.c

clean:
	rm -f *.o
	rm -f *.out

start: programa.out
	./programa.out

test: programa.out
	valgrind ./programa.out
