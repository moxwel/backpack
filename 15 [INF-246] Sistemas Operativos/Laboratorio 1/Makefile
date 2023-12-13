CC = gcc
CFLAGS = -Wall -g
SRCS = main.c sopa.c carpetas.c
TARGET = sopa_letras.out

all: $(TARGET)

$(TARGET): $(SRCS)
	$(CC) $(CFLAGS) $(SRCS) -o $(TARGET)

clean:
	rm -f $(TARGET) *.o
	rm -rf ./archivos/horizontal
	rm -rf ./archivos/vertical
