import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;

public class Sopa {
    // Objeto Sopa.

    // Atributos:
    public String direccionArchivo;
    public int dimension;
    public String palabraBuscar;
    public char[][] matrizSopa;

    // Estos atributos definen el rango valido de la matriz.
    // Hecho para definir subcuadrantes
    public int filaDesde;
    public int filaHasta;
    public int colDesde;
    public int colHasta;

    // Constructor: Leer el archivo y guardar los datos en los atributos.
    public Sopa(String fp) {
        this.direccionArchivo = fp;

        // Leer el archivo
        try {
            BufferedReader reader = new BufferedReader(new FileReader(this.direccionArchivo));

            this.dimension = Integer.parseInt(reader.readLine());
            this.palabraBuscar = reader.readLine();
            this.matrizSopa = new char[this.dimension][this.dimension];

            this.filaDesde = 0;
            this.filaHasta = this.dimension;
            this.colDesde = 0;
            this.colHasta = this.dimension;

            // Leer cada linea y guardar cada caracter en la matriz
            String linea;
            for (int fila = 0; fila < this.dimension; fila++) {
                linea = reader.readLine();

                String[] arregloLinea = linea.split(" ");
                for (int col = 0; col < this.dimension; col++) {
                    this.matrizSopa[fila][col] = arregloLinea[col].charAt(0);
                }
            }

            reader.close();

        } catch (IOException e) {
            e.printStackTrace();
            System.out.println("Error al leer el archivo.");
        }
    }



    // Metodo para definir el rango de la matriz.
    // Hecho para definir subcuadrantes. Modifica los atributos.
    public boolean setLimites(int fDesde, int fHasta, int cDesde, int cHasta) {
        // Verificar que no se salga de los limites de la matriz
        if (fDesde < 0 || fHasta > this.dimension || cDesde < 0 || cHasta > this.dimension) {
            return false;
        }

        // 'desde' no puede ser mayor o igual que 'hasta'
        if (fDesde >= fHasta || cDesde >= cHasta) {
            return false;
        }

        this.filaDesde = fDesde;
        this.filaHasta = fHasta;
        this.colDesde = cDesde;
        this.colHasta = cHasta;
        return true;
    }



    // ===========
    // Metodo para imprimir la sopa de letras.

    public void printSopa() {
        System.out.printf("\nDimension: %d\nPalabra: %s (%d)\nDimension subcuadrante: %dx%d\nRango subcuadrante: [F:%d-%d, C:%d-%d]\n",
                          this.dimension, this.palabraBuscar, this.palabraBuscar.length(),
                          (this.filaHasta - this.filaDesde), (this.colHasta - this.colDesde),
                          this.filaDesde, this.filaHasta, this.colDesde, this.colHasta);

        for (int fila = this.filaDesde; fila < this.filaHasta; fila++) {
            for (int col = this.colDesde; col < this.colHasta; col++) {
                System.out.print(this.matrizSopa[fila][col] + " ");
            }
            System.out.println();
        }
    }

    // Metodo para imprimir la sopa de letras pero con un rango especifico (sin modificar los atributos).
    public void printSopa(int fDesde, int fHasta, int cDesde, int cHasta) {
        System.out.printf("\nDimension: %d\nPalabra: %s (%d)\nDimension subcuadrante: %dx%d\nRango subcuadrante: [F:%d-%d, C:%d-%d]\n",
                          this.dimension, this.palabraBuscar, this.palabraBuscar.length(),
                          (fHasta - fDesde), (cHasta - cDesde),
                          fDesde, fHasta, cDesde, cHasta);

        for (int fila = fDesde; fila < fHasta; fila++) {
            for (int col = cDesde; col < cHasta; col++) {
                System.out.print(this.matrizSopa[fila][col] + " ");
            }
            System.out.println();
        }
    }
    // ===========





    // ===========
    // Metodos para buscar la palabra en la sopa de letras.

    public void buscarPalabra() {
        System.out.println("\nBuscando la palabra: \"" + this.palabraBuscar + "\" en el rango: [F:" + this.filaDesde + "-" + this.filaHasta + ", C:" + this.colDesde + "-" + this.colHasta + "]");
        int largoPalabra = this.palabraBuscar.length();

        for (int fila = this.filaDesde; fila < this.filaHasta; fila++) {
            for (int col = this.colDesde; col < this.colHasta; col++) {
                if (buscarHorizontal(fila, col, largoPalabra)) {
                    System.out.println("Palabra horizontal en la fila " + (fila + 1) + ", columna " + (col + 1));
                    return;
                }

                if (buscarVertical(fila, col, largoPalabra)) {
                    System.out.println("Palabra vertical en la fila " + (fila + 1) + ", columna " + (col + 1));
                    return;
                }
            }
        }

        System.out.println("No se encontro la palabra.");
    }

    private boolean buscarHorizontal(int fila, int col, int largoPalabra) {
        if (col + largoPalabra > this.colHasta) {
            return false;
        }

        for (int car = 0; car < largoPalabra; car++) {
            if (this.matrizSopa[fila][col + car] != this.palabraBuscar.charAt(car)) {
                return false;
            }
        }

        return true;
    }

    private boolean buscarVertical(int fila, int col, int largoPalabra) {
        if (fila + largoPalabra > this.filaHasta) {
            return false;
        }

        for (int car = 0; car < largoPalabra; car++) {
            if (this.matrizSopa[fila + car][col] != this.palabraBuscar.charAt(car)) {
                return false;
            }
        }

        return true;
    }
    // ===========




    // ===========
    // Metodos para buscar la palabra en la sopa de letras pero con un rango especifico (sin modificar los atributos).

    public boolean buscarPalabra(int fDesde, int fHasta, int cDesde, int cHasta) {
        System.out.println("\nBuscando la palabra: \"" + this.palabraBuscar + "\" en el rango: [F:" + fDesde + "-" + fHasta + ", C:" + cDesde + "-" + cHasta + "]");

        // Verificar que no se salga de los limites de la matriz
        if (fDesde < 0 || fHasta > this.dimension || cDesde < 0 || cHasta > this.dimension) {
            System.out.println("Error: Rango invalido.");
            return false;
        }

        // 'desde' no puede ser mayor o igual que 'hasta'
        if (fDesde >= fHasta || cDesde >= cHasta) {
            System.out.println("Error: Rango invalido.");
            return false;
        }

        int largoPalabra = this.palabraBuscar.length();

        for (int fila = fDesde; fila < fHasta; fila++) {
            for (int col = cDesde; col < cHasta; col++) {
                if (buscarHorizontal(fila, col, largoPalabra, cHasta)) {
                    System.out.println("Palabra horizontal en la fila " + (fila + 1) + ", columna " + (col + 1));
                    return true;
                }

                if (buscarVertical(fila, col, largoPalabra, fHasta)) {
                    System.out.println("Palabra vertical en la fila " + (fila + 1) + ", columna " + (col + 1));
                    return true;
                }
            }
        }

        System.out.println("No se encontro la palabra.");
        return false;
    }

    private boolean buscarHorizontal(int fila, int col, int largoPalabra, int cHasta) {
        if (col + largoPalabra > cHasta) {
            return false;
        }

        for (int car = 0; car < largoPalabra; car++) {
            if (this.matrizSopa[fila][col + car] != this.palabraBuscar.charAt(car)) {
                return false;
            }
        }

        return true;
    }

    private boolean buscarVertical(int fila, int col, int largoPalabra, int fHasta) {
        if (fila + largoPalabra > fHasta) {
            return false;
        }

        for (int car = 0; car < largoPalabra; car++) {
            if (this.matrizSopa[fila + car][col] != this.palabraBuscar.charAt(car)) {
                return false;
            }
        }

        return true;
    }

    // Sobrecarga de 'buscarPalabra' sin prints.
    public boolean buscarPalabra(int fDesde, int fHasta, int cDesde, int cHasta, boolean debug) {
        // System.out.println("\nBuscando la palabra: \"" + this.palabraBuscar + "\" en el rango: [F:" + fDesde + "-" + fHasta + ", C:" + cDesde + "-" + cHasta + "]");

        // Verificar que no se salga de los limites de la matriz
        if (fDesde < 0 || fHasta > this.dimension || cDesde < 0 || cHasta > this.dimension) {
            // System.out.println("Error: Rango invalido.");
            return false;
        }

        // 'desde' no puede ser mayor o igual que 'hasta'
        if (fDesde >= fHasta || cDesde >= cHasta) {
            // System.out.println("Error: Rango invalido.");
            return false;
        }

        int largoPalabra = this.palabraBuscar.length();

        for (int fila = fDesde; fila < fHasta; fila++) {
            for (int col = cDesde; col < cHasta; col++) {
                if (buscarHorizontal(fila, col, largoPalabra, cHasta)) {
                    System.out.println("Palabra horizontal en la fila " + (fila + 1) + ", columna " + (col + 1));
                    return true;
                }

                if (buscarVertical(fila, col, largoPalabra, fHasta)) {
                    System.out.println("Palabra vertical en la fila " + (fila + 1) + ", columna " + (col + 1));
                    return true;
                }
            }
        }

        // System.out.println("No se encontro la palabra.");
        return false;
    }
    // ===========
}
