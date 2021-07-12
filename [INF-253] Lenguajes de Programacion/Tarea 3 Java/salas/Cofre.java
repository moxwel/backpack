package salas;

import java.util.Random;

public class Cofre extends Sala {

    private int recomp;
    private int r;
    private boolean abierto; // El cofre solo puede abrirse una vez

    /**
     * Constructor para Sala de Cofre
     **/
    public Cofre() {
        nombreDeLaSala = "Cofre";
        roomType = 1;
        visited = false;
        abierto = false;
        recomp = recompensa();
    }

    /**
     * Generar una recompensa aleatoria
     * @return 'int' con la recompensa. -1=Bomba, 1-2-3=Corazones
     **/
    int recompensa() {
        Random rand = new Random();

        int type = rand.nextInt(4);
        if (type == 0) {
            r = -1;
        } else if (type == 1) {
            r = 1;
        } else if (type == 2) {
            r = 2;
        } else if (type == 3) {
            r = 3;
        }

        return r;
    }

    /**
     * Obtiene la recompensa contenida en el cofre.
     * @return -1=Bomba, 1-2-3=Corazones
     **/
    public int getRecompensa() {
        abierto = true;
        return recomp;
    }

    /**
     * Consulta si el cofre ya esta abierto
     * @return True=Si, False=No
     **/
    public boolean isAbierto() {
        return abierto;
    }
}
