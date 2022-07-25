package objetos;

import personajes.*;

public interface Objeto {

    /**
     * Aplica el efecto del objeto a un personaje dado
     * @param chr Personaje a aplicar efectos
     **/
    void efecto(Personaje chr);

    /**
     * Quita los efectos del objeto a un personaje dado
     * @param chr Personaje a quitar efectos
     **/
    void quitarEfecto(Personaje chr);

    /**
     * Obtiene el nombre del objeto
     * @return 'String' con el nombre del objeto
     **/
    String getNombre();

    /**
     * Consulta si el Objeto es de tipo Pasivo
     * @return True=Pasivo, False=Activo
     **/
    boolean isPassive();

    /**
     * Consulta si el Objeto es usable
     * @return True=Si, False=No
     **/
    boolean isUsable();

    /**
     * Imprime informacion del objeto
     **/
    void print();
}
