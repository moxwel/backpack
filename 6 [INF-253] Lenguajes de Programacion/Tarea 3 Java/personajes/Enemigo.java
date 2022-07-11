package personajes;

import java.util.Random;

public class Enemigo extends Personaje {

    /**
     * Constrctor para Enemigo
     * @param id Personaje. 1=Leak, 2=Goomb, 3=SiMuerto
     **/
    public Enemigo(int id) {
        if (id <= 1) {
            nombre    = "Leakitu";
            corazones = 7;
            dano      = 1;
            velocidad = 2;
        } else if (id == 2) {
            nombre    = "Goombation Fault";
            corazones = 8;
            dano      = 1;
            velocidad = 1;
        } else if (id >= 3) {
            nombre    = "SiMuerto";
            corazones = 6;
            dano      = 2;
            velocidad = 3;
        }
    }

    // Funcion de la superclase 'Personaje'
    public void ataque(Personaje chr) {
        System.out.println("[ATAQUE] " + this.getNombre() + " --> " + chr.getNombre());

        Random rand = new Random();

        // Hay un 50 porciento de acertar el ataque
        if (rand.nextBoolean()) {
            System.out.println("    " + dano + " puntos de dano.");

            chr.setCorazones(chr.getCorazones() - dano);
        } else {
            System.out.println("    FALLO!");
        }
    }

    /**
     * Imprime informacion acerca del Personaje Enemigo
     **/
    public void print() {
        System.out.println("\n===== ENEMIGO =====");
        System.out.println("Nom: " + nombre);
        System.out.println("Crz: " + corazones);
        System.out.println("Dno: " + dano);
        System.out.println("Vzd: " + velocidad);
        System.out.println("Crt: " + critico);
        System.out.println("===================\n");
    }
}
