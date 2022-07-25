package personajes;

public abstract class Personaje {

    protected String  nombre;
    protected int     corazones;
    protected int     dano;
    protected int     velocidad;
    protected int     critico; // Porcentaje de probabilidad para hacer mas dano

    /**
     * Ataca a un 'Personaje'. 50 porciento de acertar.
     * @param chr 'Personaje' a atacar
     **/
    public abstract void ataque(Personaje chr);

    public String getNombre()    { return nombre; }
    public int    getCorazones() { return corazones; }
    public int    getDano()      { return dano; }
    public int    getVelocidad() { return velocidad; }
    public int    getCritico()   { return critico; }

    public void setCorazones(int corazones) { this.corazones = corazones; }
    public void setDano(int dano)           { this.dano = dano; }
    public void setVelocidad(int velocidad) { this.velocidad = velocidad; }
    public void setCritico(int critico)     { this.critico = critico; }
}
