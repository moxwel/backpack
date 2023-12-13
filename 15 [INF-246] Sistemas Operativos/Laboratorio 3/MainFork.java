import java.util.concurrent.ForkJoinPool;
import java.util.concurrent.RecursiveTask;

public class MainFork extends RecursiveTask<Boolean> {

    Sopa mi_sopa;
    TiempoAlgoritmo cronometro;
    int filaDesde, filaHasta, colDesde, colHasta, tamMin;


    public MainFork(Sopa sopa, int fDesde, int fHasta, int cDesde, int cHasta, TiempoAlgoritmo cron) {
        this.mi_sopa = sopa;
        this.cronometro = cron;
        this.filaDesde = fDesde;
        this.filaHasta = fHasta;
        this.colDesde = cDesde;
        this.colHasta = cHasta;
        this.tamMin = sopa.palabraBuscar.length();
    }

    @Override
    protected Boolean compute() {
        int iniFila = this.filaDesde;
        int finFila = this.filaHasta;
        int iniCol = this.colDesde;
        int finCol = this.colHasta;

        if (finFila - iniFila == tamMin && finCol - iniCol == tamMin) {
            // Si el tamaño de la matriz es igual al tamaño minimo, se busca la palabra
            if (this.mi_sopa.buscarPalabra(iniFila, finFila, iniCol, finCol, false)) {
                // Se encuentra la palabra, se termina el cronometro.
                this.cronometro.terminar();
                return true;
            } else {
                return false;
            }
        } else {
            // Si el tamaño no corresponde, se divide la matriz en 4 partes y asignar a cada una de ellas un proceso
            int mitadFila = (finFila - iniFila) / 2;
            int mitadCol = (finCol - iniCol) / 2;

            // Se crean los procesos
            MainFork submatrizNO = new MainFork(this.mi_sopa,
                                                iniFila             , iniFila + mitadFila,
                                                iniCol              , iniCol + mitadCol,
                                                this.cronometro);

            MainFork submatrizNE = new MainFork(this.mi_sopa,
                                                iniFila             , iniFila + mitadFila,
                                                iniCol + mitadCol   , finCol,
                                                this.cronometro);

            MainFork submatrizSO = new MainFork(this.mi_sopa,
                                                iniFila + mitadFila , finFila,
                                                iniCol              , iniCol + mitadCol,
                                                this.cronometro);

            MainFork submatrizSE = new MainFork(this.mi_sopa,
                                                iniFila + mitadFila , finFila,
                                                iniCol + mitadCol   , finCol,
                                                this.cronometro);

            // Se ejecutan los procesos
            submatrizNO.fork();
            submatrizNE.fork();
            submatrizSO.fork();
            submatrizSE.fork();

            // Se espera a que terminen los procesos. El valor se retorna al metodo 'join()' de la clase padre.
            return submatrizNO.join() || submatrizNE.join() || submatrizSO.join() || submatrizSE.join();
            // En el caso del metodo 'main', el valor lo recibe el metodo 'invoke()''.
        }
    }

    public static void main(String[] args) {
        // Verificar si hay argumentos
        if (args.length == 0) {
            System.out.println("No se ha especificado el archivo de entrada.");
            return;
        }

        Sopa mi_sopa = new Sopa(args[0]);
        TiempoAlgoritmo cronometro = new TiempoAlgoritmo();

        // Pool de procesos, donde se ejecutaran los procesos creados con fork()
        ForkJoinPool pool = new ForkJoinPool();
        MainFork tarea = new MainFork(mi_sopa, 0, mi_sopa.filaHasta, 0, mi_sopa.colHasta, cronometro);

        // Se inicia el cronometro y se ejecuta la tarea
        cronometro.iniciar();
        pool.invoke(tarea);

        System.out.println("Palarbra: " + mi_sopa.palabraBuscar + ". Tiempo de ejecución: " + cronometro.obtenerTiempoMilisegundos() + " milisegundos.");
    }

}
