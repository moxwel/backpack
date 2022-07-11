package salas;

public class Sala {

    protected String nombreDeLaSala;
    protected int roomType;          // 0=Tesoro, 1=Cofre, 2=Enem, 3=Vacia, 4=Meta
    protected boolean visited;       // Sala visitada

    public String getNombreDeLaSala() {
        return nombreDeLaSala;
    }

    /**
     * Consulta el tipo de sala
     * @return 0=Tesoro, 1=Cofre, 2=Enem, 3=Vacia, 4=Meta
     **/
    public int getRoomType() {
        return roomType;
    }

    /**
     * Consulta si la sala ha sido visitada o no
     * @return True=Visitada, Fasle=No visitada
     **/
    public boolean isVisited() {
        return visited;
    }

    /**
     * Marca una sala como "visitada"
     **/
    public void makeVisited() {
        if (visited == false) {
            visited = true;
        }
    }
}
