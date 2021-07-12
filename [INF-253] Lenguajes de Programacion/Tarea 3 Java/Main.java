import java.util.Random;
import java.util.Scanner;

import personajes.*;
import objetos.*;
import salas.*;

public class Main {
    public static void main(String[] args) {
        System.out.println("\n\n     THE BINDING OF DON FONDON     ");
        System.out.println("===================================\n");

        // Seleccionar personaje
        Scanner entrada = new Scanner(System.in);
        System.out.println("Elige aquel quien va a estar a tu merced...");
        System.out.print("[1] Don Mario\n[2] Don Sora\n[3] Don NoMuerto\n\n>>>");
        int jug = entrada.nextInt();
        entrada.nextLine();
        Jugador yomismo = new Jugador(jug);
        yomismo.print();
        entrada.reset();

        boolean playerAlive = true;
        boolean escaped = false;
        int puntaje = 0;
        // Loop principal del juego
        while (true) {
            // Cada vez que se termine un mapa, se genera otro
            Laberinto mapa = new Laberinto();

            // Mientras el jugador este vivo
            while (playerAlive) {

                // Si es una sala de enemigos, hay que pelear primero
                if ((!escaped) && (mapa.getRoomType() == 2)) {
                    mapa.print();

                    // Hay que pelear hasta que no hayan mas enemigos o no hayas escapado
                    while ((!escaped) && (((Enemigos)mapa.getSala()).getEnemys() != 0)) {

                        System.out.println("\nHay " + ((Enemigos)mapa.getSala()).getEnemys() + " enemigo/s!");
                        Enemigo actualEnemy = ((Enemigos)mapa.getSala()).devolverEnemigo();
                        System.out.println("Te encuentras con '" + actualEnemy.getNombre() + "'");

                        // Solo se puede pelear mientras sigan vivos
                        while ((yomismo.getCorazones() > 0) && (actualEnemy.getCorazones() > 0)) {
                            yomismo.print();
                            actualEnemy.print();
                            System.out.print("Que hacer? [atacar | usar | escapar]\n>>>");
                            String resp = entrada.nextLine();

                            if (resp.equals("atacar")) {
                                yomismo.pelearConEnemigo(actualEnemy);

                            } else if (resp.equals("usar")) {
                                if (yomismo.getItem().isUsable()) {
                                    System.out.println("Usando item '" + yomismo.getItem().getNombre() + "'");
                                    yomismo.activarObjeto();
                                } else {
                                    System.out.println("El item no se puede usar.");
                                }
                                actualEnemy.ataque(yomismo);

                            } else if (resp.equals("escapar")) {
                                Random rand = new Random();
                                int escape = rand.nextInt(101);

                                // Hay un 10 poriciento de probabilidad de escapar de un enemigo
                                if (escape <= 10) {
                                    System.out.println("Lograste escapar del enemigo!");
                                    escaped = true;
                                    break;

                                } else {
                                    System.out.println("Intentaste correr, pero te tropezaste.");
                                    actualEnemy.ataque(yomismo);
                                }
                            }
                        }
                    }
                    yomismo.print();
                }

                // En cada accion, se verifica que el personaje siga vivo
                if (yomismo.getCorazones() <= 0) {
                    System.out.println("\nTe moristes wey.\n/// Fin del juego ///");
                    playerAlive = false;
                    break;
                }

                mapa.print();
                mapa.whereIAm();
                System.out.print("Que hacer? [help]\n>>>");
                String inp = entrada.nextLine();

                if (inp.equals("arriba") || inp.equals("subir")) {
                    escaped = false;
                    mapa.mover(0);

                } else if (inp.equals("abajo") || inp.equals("bajar")) {
                    escaped = false;
                    mapa.mover(2);

                } else if (inp.equals("izq") || inp.equals("izquierda")) {
                    escaped = false;
                    mapa.mover(3);

                } else if (inp.equals("der") || inp.equals("derecha")) {
                    escaped = false;
                    mapa.mover(1);

                } else if (inp.equals("morir") || inp.equals("estudiar")) {
                    System.out.println("Por alguna razon decidiste estudiar en la USM. Y falleciste...\n/// Fin del juego ///");
                    playerAlive = false;
                    break;

                } else if (inp.equals("stats") || inp.equals("yo")) {
                    yomismo.print();

                } else if (inp.equals("usar")) {
                    if (yomismo.getItem().isUsable()) {
                        System.out.println("Usando item '" + yomismo.getItem().getNombre() + "'");
                        yomismo.activarObjeto();
                    } else {
                        System.out.println("El item no se puede usar.");
                    }

                } else if (inp.equals("meta")) {
                    if (mapa.getRoomType() == 4) {
                        System.out.println("Llegaste a la meta!\n/// +1 PUNTO ///");
                        break;
                    } else {
                        System.out.println("Ya quisieras, pero no, no hay escape aqui...");
                    }

                } else if (inp.equals("cofre")) {
                    if (mapa.getRoomType() == 1) {
                        System.out.println("Abriendo cofre...");
                        yomismo.abrirCofre(((Cofre)mapa.getSala()));

                    } else {
                        System.out.println("No hay ningun cofre aqui.");
                    }

                } else if (inp.equals("tesoro")) {
                    if (mapa.getRoomType() == 0) {
                        System.out.println("Viendo tesoro...");

                        Objeto piso = ((Tesoro)mapa.getSala()).getItemEnElPiso();
                        piso.print();

                        yomismo.print();

                        System.out.print("Tomar? [si | no]\n>>>");
                        String resp = entrada.nextLine();

                        if (resp.equals("si") || resp.equals("y")) {
                            ((Tesoro)mapa.getSala()).setItemEnElPiso(yomismo.agarrarObjeto(piso));
                            yomismo.print();
                        }

                    } else {
                        System.out.println("No hay ningun tesoro aqui.");
                    }

                } else if (inp.equals("atacar")) {
                    System.out.println("Intentaste golpear al aire, pero no paso nada.");

                } else if (inp.equals("escapar")) {
                    System.out.println("De quien intentas escapar?");

                } else if (inp.equals("help")) {
                    System.out.println("\n===== HELP =====");
                    System.out.println("- arriba | abajo | izq | der : Moverse por el mapa");
                    System.out.println("- yo | stats                 : Revisar el estado del jugador\n");
                    System.out.println("- atacar                     : Ataca a un enemigo");
                    System.out.println("- escapar                    : Escapa de un enemigo\n");
                    System.out.println("- usar                       : Utilizar el item en la mano");
                    System.out.println("- cofre                      : Abrir el cofre de la sala");
                    System.out.println("- tesoro                     : Tomar el tesoro de la sala");
                    System.out.println("- meta                       : Escapar (si es posible)\n");
                    System.out.println("- morir                      : Se explica solo...?");
                    System.out.println("================\n");
                }
            }

            if (!playerAlive) {
                System.out.println("\nPuntaje total: " + puntaje + "\n");
                break;
            }

            puntaje = puntaje + 1;
            System.out.println("\nPuntaje: " + puntaje + "\n");
        }

        entrada.close();
    }
}
