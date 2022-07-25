package salas;

public class Vacia extends Sala {
    private int vaciaOMeta;

    /**
     * Constructor para sala vacia
     * @param id Tipo de sala. 0=Vacia, 1=Meta
     **/
    public Vacia(int id) {
        visited = false;
        if (id == 1) {
            nombreDeLaSala = "Meta";
            roomType = 4;
            vaciaOMeta = 1;
        } else {
            nombreDeLaSala = "Vacia";
            roomType = 3;
            vaciaOMeta = 0;
        }
    }

    public int getVaciaOMeta() {
        return vaciaOMeta;
    }
}
