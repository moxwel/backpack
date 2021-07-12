package salas;

import java.util.Random;
import personajes.*;

public class Enemigos extends Sala {

    private Enemigo[] enemigos = {null, null, null};
    private int nEnemy; // Cantidad de enemigos

    /**
     * Constructor para Sala de Enemigos
     **/
    public Enemigos() {
        nombreDeLaSala = "Enemigos";
        nEnemy = 0;
        roomType = 2;
        visited = false;
        generarEnemigos();
    }

    /**
     * Generar enemigos aleatorios
     **/
    void generarEnemigos() {
        Random rand = new Random();

        int randomCant = rand.nextInt(3);
        int enemigo1 = rand.nextInt(3) + 1;
        int enemigo2 = rand.nextInt(3) + 1;
        int enemigo3 = rand.nextInt(3) + 1;

        // Pueden ser 1, 2 o 3 enemigos
        if (randomCant == 0) {
            enemigos[0] = new Enemigo(enemigo1);
            nEnemy = 1;
        } else if (randomCant == 1) {
            enemigos[0] = new Enemigo(enemigo1);
            enemigos[1] = new Enemigo(enemigo2);
            nEnemy = 2;
        } else if (randomCant == 2) {
            enemigos[0] = new Enemigo(enemigo1);
            enemigos[1] = new Enemigo(enemigo2);
            enemigos[2] = new Enemigo(enemigo3);
            nEnemy = 3;
        }
    }

    /**
     * Devuelve el primer enemigo de la lista, y lo borra.
     * @return Objeto de tipo 'Enemigo'
     **/
    public Enemigo devolverEnemigo() {
        Enemigo ultimo = null;

        if (nEnemy == 1) {
            ultimo = enemigos[0];
            enemigos[0] = null;
            nEnemy = 0;
        } else if (nEnemy == 2) {
            ultimo = enemigos[0];
            enemigos[0] = enemigos[1];
            enemigos[1] = null;
            nEnemy = 1;
        } else if (nEnemy == 3) {
            ultimo = enemigos[0];
            enemigos[0] = enemigos[1];
            enemigos[1] = enemigos[2];
            enemigos[2] = null;
            nEnemy = 2;
        }

        return ultimo;
    }

    /**
     * Retorna la cantidad de enemigos dentro de la sala
     * @return 'int' con la cantidad de enemigos
     **/
    public int getEnemys() {
        return nEnemy;
    }
}
