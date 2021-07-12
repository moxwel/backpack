package objetos;

import personajes.*;

public class Pasivo implements Objeto {

    private String nombre;
    private int    corazonesMod;
    private int    danoMod;
    private int    velocidadMod;
    private int    criticoMod;

    /**
     * Constrctor para objetos Pasivos
     * @param id Tipo de objeto. 1=Esp Her, 2=Fir Func, 3=Cap Comp
     **/
    public Pasivo(int id) {
        nombre       = "Item";
        corazonesMod = 0;
        danoMod      = 0;
        velocidadMod = 0;
        criticoMod   = 0;

        if (id <= 1) {
            nombre       = "Espada Heredada";
            corazonesMod = -1;
            danoMod      = 1;
        } else if (id == 2) {
            nombre     = "Firaga Funcional";
            criticoMod = 10;
        } else if (id >= 3) {
            nombre       = "Caparazon de Compilacion";
            corazonesMod = 2;
            danoMod      = -1;
        }
    }

    // Funcion de la interfaz 'Objeto'
    public void efecto(Personaje chr) {
        // System.out.println("Aplicar efecto");

        chr.setCorazones(chr.getCorazones() + corazonesMod);
        chr.setDano(chr.getDano() + danoMod);
        chr.setVelocidad(chr.getVelocidad() + velocidadMod);
        chr.setCritico(chr.getCritico() + criticoMod);
    }

    // Funcion de la interfaz 'Objeto'
    public void quitarEfecto(Personaje chr) {
        // System.out.println("Quitar Efecto");

        chr.setCorazones(chr.getCorazones() - corazonesMod);
        chr.setDano(chr.getDano() - danoMod);
        chr.setVelocidad(chr.getVelocidad() - velocidadMod);
        chr.setCritico(chr.getCritico() - criticoMod);
    }

    // Funcion de la interfaz 'Objeto'
    public String  getNombre() { return nombre; }

    // Funcion de la interfaz 'Objeto'
    public boolean isPassive() { return true; }

    // Funcion de la interfaz 'Objeto'
    public boolean isUsable()  { return false; }

    // Funcion de la interfaz 'Objeto'
    public void print() {
        System.out.println("\n===== ITEM PASIVO =====");
        System.out.println("Nom: " + nombre);
        System.out.println("Crz: " + corazonesMod);
        System.out.println("Dno: " + danoMod);
        System.out.println("Vzd: " + velocidadMod);
        System.out.println("Crt: " + criticoMod);
        System.out.println("=======================\n");
    }
}
