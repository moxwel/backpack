import java.util.concurrent.CountDownLatch;

public class MainThread {
    public static void main(String[] args) {
        // Verificar si hay argumentos
        if (args.length == 0) {
            System.out.println("No se ha especificado el archivo de entrada.");
            return;
        }

        Sopa sopa = new Sopa(args[0]);

        int dimension = sopa.dimension;

        // Crear una CountDownLatch para sincronizar las hebras
        CountDownLatch latchPadre = new CountDownLatch(1);

        // Crear un cronometro para medir el tiempo de ejecucion
        TiempoAlgoritmo cronometro = new TiempoAlgoritmo();

        // Crear hebra principal
        cronometro.iniciar();
        SopaThread threadInicial = new SopaThread(sopa, 0, dimension, 0, dimension, latchPadre, cronometro);
        threadInicial.start();

        try {
            // Esperar a que todas las hebras terminen
            latchPadre.await();
        } catch (InterruptedException e) {
            e.printStackTrace();
        }

        // Imprimir el tiempo de ejecucion
        System.out.println("Palarbra: " + sopa.palabraBuscar + ". Tiempo de ejecucion: " + cronometro.obtenerTiempoMilisegundos() + " milisegundos.");
    }
}

class SopaThread extends Thread {
    private Sopa sopa;
    private int iniFila;
    private int finFila;
    private int iniCol;
    private int finCol;
    private int tamMin;
    private CountDownLatch latch;
    public TiempoAlgoritmo cronometro;

    public SopaThread(Sopa sopa,
                      int filaDesde, int filaHasta,
                      int colDesde, int colHasta,
                      CountDownLatch latch, TiempoAlgoritmo cronometro) {

        this.sopa = sopa;
        this.iniFila = filaDesde;
        this.finFila = filaHasta;
        this.iniCol = colDesde;
        this.finCol = colHasta;
        this.tamMin = sopa.palabraBuscar.length();
        this.latch = latch;
        this.cronometro = cronometro;
    }

    @Override
    public void run() {
        if (finFila - iniFila == tamMin && finCol - iniCol == tamMin) {
            // Si la dimension de la sopa es igual a la longitud de la palabra, entonces buscar en la sopa
            if (this.sopa.buscarPalabra(iniFila, finFila, iniCol, finCol, false)) {
                this.cronometro.terminar();
            }
            this.latch.countDown();
            return;
        } else {
            // Subdividir en 4 cuadrantes y llamar recursivamente
            int midFila = (finFila - iniFila) / 2;
            int midCol = (finCol - iniCol) / 2;

            // Crear 4 hebras para cada cuadrante
            CountDownLatch latchCuadrante = new CountDownLatch(4);
            SopaThread threadNO = new SopaThread(sopa, iniFila, iniFila + midFila, iniCol, iniCol + midCol, latchCuadrante, cronometro);
            SopaThread threadNE = new SopaThread(sopa, iniFila, iniFila + midFila, iniCol + midCol, finCol, latchCuadrante, cronometro);
            SopaThread threadSO = new SopaThread(sopa, iniFila + midFila, finFila, iniCol, iniCol + midCol, latchCuadrante, cronometro);
            SopaThread threadSE = new SopaThread(sopa, iniFila + midFila, finFila, iniCol + midCol, finCol, latchCuadrante, cronometro);

            // Iniciar las 4 hebras
            threadNO.start();
            threadNE.start();
            threadSO.start();
            threadSE.start();

            try {
                // Esperar a que todas las hebras terminen
                latchCuadrante.await();
            } catch (InterruptedException e) {
                e.printStackTrace();
            }

            // Contar la hebra actual
            this.latch.countDown();
        }
        return;
    }
}
