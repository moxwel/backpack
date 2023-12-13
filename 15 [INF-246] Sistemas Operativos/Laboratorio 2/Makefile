CC = gcc
CFLAGS = -g
SRCS = main.c herramientas.c laberinto.c jugadores.c mapa.c archivos.c
TARGET = magic_maze.out

all: $(TARGET)

$(TARGET): $(SRCS)
	$(CC) $(CFLAGS) $(SRCS) -o $(TARGET)

clean:
	rm -f $(TARGET) *.o

run: $(TARGET)
	./$(TARGET)
