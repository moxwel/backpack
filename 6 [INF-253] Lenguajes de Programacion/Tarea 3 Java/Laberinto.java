import java.util.Random;
import salas.*;

public class Laberinto {
    private int x;
    private int y;
    private Sala[][] mapa;
    private Sala creada;

    /**
     * Constructor para Laberinto.
     **/
    public Laberinto() {
        // System.out.println("Creando mapa...");
        mapa = new Sala[5][5];

        System.out.println("Generando mapa...");
        generarMapa();
    }

    /**
     * Genera salas aleatorias para el laberinto.
     **/
    private void generarMapa() {
        Random rand = new Random();

        // elegir ubicacion aleatoria de la meta
        int randM = rand.nextInt(4);
        if (randM == 0) {
            x = 0; y = rand.nextInt(5);
        } else if (randM == 1) {
            x = rand.nextInt(5); y = 0;
        } else if (randM == 2) {
            x = 4; y = rand.nextInt(5);
        } else if (randM == 3) {
            x = rand.nextInt(5); y = 4;
        }

        mapa[x][y] = new Vacia(1);
        // mapa[x][y].makeVisited();
        // System.out.println("Creando sala N: 1 en [0][4] (Meta)");

        // Crear 9 salas mas...
        for (int i = 1; i < 10; i++) {
            // Elegir tipo de sala aleatoria
            int randT = rand.nextInt(4);
            // System.out.println("Creando sala N: " + (i+1) + " - Tipo: " + randT);

            // Seleccionar el tipo de sala de forma aleatoria
            if (i == 9) {
                // Si es la ultima sala en generarse, sera la de inicio

                creada = new Vacia(0);
                creada.makeVisited();
                // System.out.println("Crando ultima sala... Sala vacia.");
            } else {
                if (randT == 0) {
                    creada = new Tesoro();
                } else if (randT == 1) {
                    creada = new Cofre();
                } else if (randT == 2) {
                    creada = new Enemigos();
                } else if (randT == 3) {
                    creada = new Vacia(0);
                }
            }
            // creada.makeVisited();

            boolean roomGenerated = false;
            while (!roomGenerated) {
                // Elegir ubicacion aleatoria
                int randX = rand.nextInt(5);
                int randY = rand.nextInt(5);

                // System.out.println("Revisando coord [" + randX + "][" + randY + "]");

                // Solo generar sala si es que la sala elegida esta vacia
                if (mapa[randX][randY] == null) {
                    if ((randX > 0) && (randX < 4)) {
                        if ((randY > 0) && (randY < 4)) {
                            if ((mapa[randX + 1][randY] != null) || (mapa[randX - 1][randY] != null) || (mapa[randX][randY + 1] != null) || (mapa[randX][randY - 1] != null)) {
                                x = randX;
                                y = randY;
                                mapa[x][y] = creada;
                                roomGenerated = true;
                            }
                        } else {
                            if ((mapa[randX + 1][randY] != null) || (mapa[randX - 1][randY] != null)) {
                                x = randX;
                                y = randY;
                                mapa[x][y] = creada;
                                roomGenerated = true;
                            }
                        }
                    } else {
                        if ((randY > 0) && (randY < 4)) {
                            if ((mapa[randX][randY + 1] != null) || (mapa[randX][randY - 1] != null)) {
                                x = randX;
                                y = randY;
                                mapa[x][y] = creada;
                                roomGenerated = true;
                            }
                        }
                    }
                }
            }
        }
    }

    /**
     * Imprime el laberinto actual por pantalla
     **/
    public void print(){
        System.out.print("\n|===== MAP =====|\n|");
        for (int x = 0; x < 5; x++) {
            for (int y = 0; y < 5; y++) {
                if (mapa[x][y] != null) {
                    if ((x == this.x) && (y == this.y)) {
                        System.out.print(" @ ");
                    } else if (mapa[x][y].isVisited() == false) {
                        System.out.print(" . ");
                    } else if (mapa[x][y].getRoomType() == 0) {
                        System.out.print(" T ");
                    } else if (mapa[x][y].getRoomType() == 1) {
                        System.out.print(" C ");
                    } else if (mapa[x][y].getRoomType() == 2) {
                        System.out.print(" E ");
                    } else if (mapa[x][y].getRoomType() == 3) {
                        System.out.print(" O ");
                    } else if (mapa[x][y].getRoomType() == 4) {
                        System.out.print(" # ");
                    }
                } else {
                    System.out.print("   ");
                }
            }
            System.out.print("|\n|");
        }
        System.out.print("===============|\n\n");
    }

    /**
     * Imprime por pantalla informacion acerca de la sala actual
     **/
    public void whereIAm() {
        System.out.println("Ubicacion: [" + x + "][" + y + "] - Tipo: " + mapa[x][y].getNombreDeLaSala());
    }

    /**
     * Consulta el tipo de sala actual
     * @return 0=Tesoro, 1=Cofre, 2=Enem, 3=Vacia, 4=Meta
     **/
    public int getRoomType() {
        return mapa[x][y].getRoomType();
    }

    /**
     * Retorna la 'Sala' actual
     * @return Objeto de tipo 'Sala'
     **/
    public Sala getSala() {
        return mapa[x][y];
    }

    /**
     * Moverse a una sala en alguna direccion.
     * @param select Direccion a moverse. 0 = arriba, 1 = derecha, 2 = abajo, 3 = izquierda
     **/
    public Sala mover(int select) {
        if (select == 0) {
            // Mover hacia arriba
            if ((x == 0) || (mapa[x - 1][y] == null)) {
                System.out.println("Te golpeaste con la pared, bruh.");
            } else {
                x = x - 1;
            }

        } else if (select == 1) {
            // Mover hacia la derecha
            if ((y == 4) || (mapa[x][y + 1] == null)) {
                System.out.println("Te golpeaste con la pared, bruh.");
            } else {
                y = y + 1;
            }

        } else if (select == 2) {
            // Mover hacia abajo
            if ((x == 4) || (mapa[x+1][y] == null)) {
                System.out.println("Te golpeaste con la pared, bruh.");
            } else {
                x = x + 1;
            }

        } else if (select == 3) {
            // Mover hacia la izquierda
            if ((y == 0) || (mapa[x][y - 1] == null)) {
                System.out.println("Te golpeaste con la pared, bruh.");
            } else {
                y = y - 1;
            }

        } else {
            System.out.println("No puedes hacer eso!");
        }
        mapa[x][y].makeVisited();
        return mapa[x][y];
    }
}
