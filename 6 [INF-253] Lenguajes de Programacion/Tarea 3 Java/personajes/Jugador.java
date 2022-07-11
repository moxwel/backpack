package personajes;

import objetos.*;
import salas.*;
import java.util.Random;

public class Jugador extends Personaje {

    private Objeto item;

    /**
     * Constrctor para Jugador
     * @param id Personaje. 1=Mario, 2=Sora, 3=NoMuerto
     **/
    public Jugador(int id) {
        if (id <= 1) {
            nombre    = "Don Mario";
            corazones = 9;
            dano      = 3;
            velocidad = 3;
            critico   = 0;
            item      = new Activo(1);
        } else if (id == 2) {
            nombre    = "Don Sora";
            corazones = 8;
            dano      = 4;
            velocidad = 3;
            critico   = 0;
            item      = new Pasivo(2);
            item.efecto(this);
        } else if (id >= 3) {
            nombre    = "Don NoMuerto";
            corazones = 11;
            dano      = 2;
            velocidad = 2;
            critico   = 0;
            item      = new Pasivo(1);
            item.efecto(this);
        }
    }

    // Funcion de la superclase 'Personaje'
    public void ataque(Personaje chr){
        System.out.println("[ATAQUE] " + this.getNombre() + " --> " + chr.getNombre());

        Random rand = new Random();

        // Hay un 50 porciento de acertar el ataque
        if (rand.nextBoolean()) {

            // Dependiendo del stat critico, es posible hacer mas dano
            if (rand.nextInt(101) < critico) {
                System.out.println("    SMAAAAAASH!!! " + (int)Math.ceil(dano * 1.5) + " puntos de dano.");
                chr.setCorazones(chr.getCorazones() - (int)Math.ceil(dano * 1.5));

            } else {
                System.out.println("    " + dano + " puntos de dano.");
                chr.setCorazones(chr.getCorazones() - dano);
            }

        } else {
            System.out.println("    FALLO!");
        }
    }

    /**
     * Pelea con un enemigo. Comienza quien tiene mayor velocidad.
     * @param enmy Enemigo con quien pelear
     **/
    public void pelearConEnemigo(Enemigo enmy) {
        System.out.println("[PELEA] '" + this.getNombre() + "' pelea contra '" + enmy.getNombre() + "'");

        // Si el jugador es mas rapido que el enemigo, comienza primero
        if (this.getVelocidad() >= enmy.getVelocidad()) {
            System.out.println("    '" + this.getNombre() + "' comienza primero!");

            this.ataque(enmy);

            // Si el enemigo ya no tiene vida, detener la pelea
            if (enmy.corazones <= 0) {
                System.out.println("    Enemigo '" + enmy.getNombre() + "' derrotado!");
                return;
            }

            enmy.ataque(this);

            // Si el jugador ya no tiene vida, detener la pelea
            if (this.corazones <= 0) {
                System.out.println("    Moriste!");
                return;
            }
        } else {
            System.out.println("    '" + enmy.getNombre() + "' comienza primero!");

            enmy.ataque(this);

            // Si el jugador ya no tiene vida, detener la pelea
            if (this.corazones <= 0) {
                System.out.println("    Moriste!");
                return;
            }

            this.ataque(enmy);

            // Si el enemigo ya no tiene vida, detener la pelea
            if (enmy.corazones <= 0) {
                System.out.println("    Enemigo '" + enmy.getNombre() + "' derrotado!");
                return;
            }
        }
    }

    /**
     * Intercambia el objeto en la mano por uno nuevo
     * @param newItem Nuevo 'Objeto' a tomar
     * @return 'Objeto' antiguo
     **/
    public Objeto agarrarObjeto(Objeto newItem) {
        System.out.println("[OBJETO] Tirando '" + item.getNombre() + "'. Tomando: '" + newItem.getNombre() + "'");
        Objeto oldItem = item;

        // Si el item actual es pasivo, al quitarlo, deberian cancelarse sus efectos
        if (item.isPassive()) {
            // System.out.println("Quitando efectos de item pasivo: " + item.getNombre());
            item.quitarEfecto(this);
        }

        // Poner el nuevo item
        setItem(newItem);

        // Si el NUEVO item es pasivo, entonces deberian aplicarse sus efectos
        if (item.isPassive()) {
            // System.out.println("Anadiendo efectos de item pasivo: " + item.getNombre());
            item.efecto(this);
        }

        return oldItem;
    }

    /**
     * Activa el objeto que tiene en la mano (solo si se puede)
     **/
    public void activarObjeto() {
        // Solo se puede activar si es usable
        if (item.isUsable()) {
            // System.out.println("[OBJETO] Activando '" + item.getNombre() + "'...");
            item.efecto(this);
        } else {
            System.out.println("[OBJETO] '" + item.getNombre() + "' no se puede usar.");
        }
    }

    /**
     * Abre un cofre dado (si se puede)
     * @param chst 'Cofre' a abrir
     **/
    public void abrirCofre(Cofre chst){
        // Solo se puede abrir el cofre si es que no se ha abierto antes
        if (!chst.isAbierto()) {
            int recomp = chst.getRecompensa();

            if (recomp == -1) {
                System.out.println("BOMBA! -1 corazon");
                this.setCorazones(this.getCorazones() - 1);
            } else {
                System.out.println("+" + recomp + " corazones!");
                this.setCorazones(this.getCorazones() + recomp);
            }

        } else {
            System.out.println("El cofre ya se abrio.");
        }

    }

    public Objeto getItem() { return item; }
    public void setItem(Objeto item) { this.item = item; }

    /**
     * Imprime informacion acerca del Personaje Jugador
     **/
    public void print() {
        System.out.println("\n===== JUGADOR =====");
        System.out.println("Nom: " + nombre);
        System.out.println("Crz: " + corazones);
        System.out.println("Dno: " + dano);
        System.out.println("Vzd: " + velocidad);
        System.out.println("Crt: " + critico);
        System.out.print("Itm: " + item.getNombre());

        if (item.isUsable()) {
            System.out.print(" [Usable]");
        } else {
            System.out.print(" [No Usable]");
        }

        System.out.println("\n===================\n");
    }
}
