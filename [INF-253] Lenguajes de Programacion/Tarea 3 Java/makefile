all: build

# Crear Jar (v=verbose, e=clase de entrada).
# TBoDF.jar - Nombre del archivo
# Main - Clase de entrada
# -C build . - Añadir archivos de la carpeta build
jar: build
	jar cfve TBoDF.jar Main -C build .

# Compilar Main y dejar los archivos en la carpeta build
build: Main.java
	javac -d ./build Main.java

run: jar
	java -jar TBoDF.jar

clean:
	rmdir /s /q build
