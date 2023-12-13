public class TiempoAlgoritmo {
    private double inicio;
    private double fin;

    public void iniciar() {
        inicio = System.nanoTime();
    }

    public void terminar() {
        fin = System.nanoTime();
    }

    public double obtenerTiempo() {
        return fin - inicio;
    }

    public double obtenerTiempoMilisegundos() {
        return (fin - inicio) / 1000000;
    }
}
