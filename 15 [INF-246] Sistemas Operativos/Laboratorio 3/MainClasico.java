public class MainClasico {
    public static void main(String[] args) {
        // Verificar si hay argumentos
        if (args.length == 0) {
            System.out.println("No se ha especificado el archivo de entrada.");
            return;
        }

        Sopa mi_sopa = new Sopa(args[0]);
        TiempoAlgoritmo cronometro = new TiempoAlgoritmo();

        cronometro.iniciar();
        mi_sopa.buscarPalabra(0, mi_sopa.filaHasta, 0, mi_sopa.colHasta, false);
        cronometro.terminar();

        System.out.println("Palarbra: " + mi_sopa.palabraBuscar + ". Tiempo de ejecucion: " + cronometro.obtenerTiempoMilisegundos() + " milisegundos.");
    }
}
