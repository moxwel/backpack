package objetos;

import personajes.*;
import java.util.Random;

public class Activo implements Objeto {

    private String  nombre;
    private int     corazonesMod;
    private int     danoMod;
    private int     velocidadMod;
    private int     criticoMod;
    private boolean usable; // Los objetos activos solo se pueden usar una vez

    /**
     * Constrctor para objetos Activos
     * @param id Tipo de objeto. 1=Cham Rec, 2=Shine
     **/
    public Activo(int id) {
        nombre       = "Item";
        corazonesMod = 0;
        danoMod      = 0;
        velocidadMod = 0;
        criticoMod   = 0;
        usable       = true;

        if (id <= 1) {
            nombre       = "Champinon Recursivo";
            danoMod      = 3;
            corazonesMod = -2;
        } else if (id >= 2) {
            nombre   = "Shine";
        }
    }

    // Funcion de la interfaz 'Objeto'
    public void efecto(Personaje chr) {
        // Solo se podra utilizar si no se ha usado antes
        if (usable) {
            // System.out.println("Aplicar efecto");

            if (nombre == "Shine") {
                // System.out.println("[SHINE] Usando Shine...");
                // Caso especial para Shine
                Random rand = new Random();

                // Hay un 10 porciento de que se use con exito
                if (rand.nextInt(101) <= 10) {
                    System.out.println("    FUNCIONO! +5 dano, -2 corazones");
                    chr.setCorazones(chr.getCorazones() - 2);
                    chr.setDano(chr.getDano() + 5);

                } else {
                    System.out.println("    Lamentablemente se rompio...");
                }

            } else {
                chr.setCorazones(chr.getCorazones() + corazonesMod);
                chr.setDano(chr.getDano() + danoMod);
                chr.setVelocidad(chr.getVelocidad() + velocidadMod);
                chr.setCritico(chr.getCritico() + criticoMod);
            }

            // No se podra usar mas
            usable = false;
        }
    }

    // Funcion de la interfaz 'Objeto'
    public void quitarEfecto(Personaje chr) {
        System.out.println("Quitar Efecto");
    }

    // Funcion de la interfaz 'Objeto'
    public String getNombre() { return nombre; }

    // Funcion de la interfaz 'Objeto'
    public boolean isPassive() { return false; }

    // Funcion de la interfaz 'Objeto'
    public boolean isUsable() { return usable; }

    // Funcion de la interfaz 'Objeto'
    public void print() {
        System.out.println("\n===== ITEM ACTIVO =====");
        System.out.println("Nom: " + nombre);
        System.out.println("Crz: " + corazonesMod);
        System.out.println("Dno: " + danoMod);
        System.out.println("Vzd: " + velocidadMod);
        System.out.println("Crt: " + criticoMod);
        System.out.print("Use: ");

        if (usable) {
            System.out.println("Usable");
        } else {
            System.out.println("No usable");
        }
        System.out.println("=======================\n");
    }
}
