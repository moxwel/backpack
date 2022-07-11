package salas;

import java.util.Random;
import objetos.*;

public class Tesoro extends Sala {

    private Objeto itemEnElPiso;
    private Objeto generado;

    /**
     * Constructor para Sala de Tesoro
     **/
    public Tesoro() {
        nombreDeLaSala = "Tesoro";
        roomType = 0;
        visited = false;
        itemEnElPiso = generarItem();
    }

    /**
     * Genera un objeto de forma aleatoria
     * @return 'Objeto' aleatorio
     **/
    Objeto generarItem() {
        Random rand = new Random();

        // Generar un objeto pasivo o activo
        if (rand.nextBoolean()) {
            // Pasivo
            int itemID = rand.nextInt(3);
            generado = new Pasivo(itemID + 1);

        } else {
            // Activo
            int itemID = rand.nextInt(2);
            generado = new Activo(itemID + 1);

        }

        return generado;
    }

    public Objeto getItemEnElPiso() {
        return itemEnElPiso;
    }

    public void setItemEnElPiso(Objeto itemEnElPiso) {
        this.itemEnElPiso = itemEnElPiso;
    }
}
