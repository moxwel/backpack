#include <iostream>
#include <fstream>
#include <sstream>
#include <string>
#include <vector>
#include <array>
#include <algorithm>
#include <iomanip>
#include <ctime>

using namespace std;

// Clase Bloque: cada bloque tiene un ID, un ancho y un alto.
class Bloque {

private:
    int  ID;    // Identificador del bloque
    int  W;     // Ancho del bloque
    int  H;     // Alto del bloque
    int  A;     // Area del bloque
    bool USED;  // Indica si el bloque ya fue usado en el acomodo.
    bool ROTD;  // Indica si el bloque fue rotado.
    int x;      // Posicion en X del bloque (esquina inferior izquierda).
    int y;      // Posicion en Y del bloque (esquina inferior izquierda).

public:
    Bloque(int pID, int pW, int pH) {
        ID   = pID;      // Identificador del bloque
        W    = pW;       // Ancho del bloque
        H    = pH;       // Alto del bloque
        A    = pW * pH;  // Area del bloque
        USED = false;    // Si el bloque ya fue usado en el acomodo.
        ROTD = false;    // Si el bloque fue rotado.
        x    = 0;        // Posicion en X del bloque (esquina inferior izquierda).
        y    = 0;        // Posicion en Y del bloque (esquina inferior izquierda).
    }
    int getID()    { return ID; }
    int getW()     { return W; }
    int getH()     { return H; }
    int getA()     { return A; }
    int getX()     { return x; }
    int getY()     { return y; }
    bool isUsed()  { return USED; }
    void use()     { USED = true; }
    bool isRoted() { return ROTD; }
    void setCoords(int pX, int pY) { x = pX; y = pY; }
    void setX(int pX) { x = pX; }
    void setY(int pY) { y = pY; }
    void rotate() {
        int temp = W;
        W        = H;
        H        = temp;
        ROTD     = !ROTD;
    }
};

void printBloques(vector<Bloque> L) {
    cout << "======LISTA BLOQUES======" << endl;
    cout << left << setw(4) << "ID" << setw(4) << "R" << setw(4) << "W" << setw(4) << "H" << setw(4) << "A" << setw(4) << "U" << setw(4) << "X" << setw(4) << "Y" << endl;
    for (int i = 0; i < L.size(); i++) {
        cout << left << setw(4) << L[i].getID() << setw(4) << L[i].isRoted() << setw(4) << L[i].getW() << setw(4) << L[i].getH() << setw(4) << L[i].getA() << setw(4) << L[i].isUsed() << setw(4) << L[i].getX() << setw(4) << L[i].getY() << endl;
    }
    cout << "====FIN LISTA BLOQUES====" << endl;
}

bool compararHDesc(Bloque B1, Bloque B2) {
    return B1.getH() > B2.getH();
}

bool compararWDesc(Bloque B1, Bloque B2) {
    return B1.getW() > B2.getW();
}

bool compararIDAsc(Bloque B1, Bloque B2) {
    return B1.getID() < B2.getID();
}

bool compararYDesc(Bloque B1, Bloque B2) {
    return B1.getY() > B2.getY();
}





class Instancia {
private:
    vector<int> listaRotados;  // Lista de rotaciones de cada bloque.

    /*
    Ejemplos:
    [0,0,1,0,0] : El bloque 3 esta rotado.
    [1,0,1,0,0] : El bloque 1 y 3 estan rotados.
    [1,1,1,0,0] : El bloque 1, 2 y 3 estan rotados.
    [1,1,1,1,0] : El bloque 1, 2, 3 y 4 estan rotados.
    [1,1,1,1,1] : Todos los bloques estan rotados.
    */

public:
    Instancia(int totalRectangulos) {
        listaRotados = vector<int>(totalRectangulos, 0);
    }
    string toString() {
        string resultado;
        for (int i = 0; i < listaRotados.size(); i++) {
            resultado += to_string(listaRotados[i]);
        }
        return resultado;
    }
    int estaRotado(int indexB) { return listaRotados[indexB]; }
    void marcarRotacion(int indexB) { listaRotados[indexB] = !listaRotados[indexB]; }
    void imprimirInstancia() {
        // cout << "=========INSTANCIA========" << endl;
        cout << "[";
        for (int i = 0; i < listaRotados.size(); i++) {
            cout << listaRotados[i];
            if (i != listaRotados.size() - 1) {
                cout << ", ";
            }
        }
        cout << "]" << endl;
        // cout << "=======FIN INSTANCIA=======" << endl;
    }
};

// Generar archivo de salida
void generarOutput(vector<Bloque> L, Instancia instancia, int areaDesper, string testFileName) {

    // Crear carpeta de salida si no existe.
    system("mkdir -p ./output");
    string nombreArchivo = "./output/" + testFileName + "_" + instancia.toString() + ".txt";

    // Crear archivo de salida.
    ofstream archivoSalida(nombreArchivo);
    if (!archivoSalida.is_open()) {
        cout << "Error creando archivo de salida." << endl;
        return;
    }

    // Calcular la altura con el bloque que esta mas arriba.
    sort(L.begin(), L.end(), compararYDesc);
    int alturaUsada = L[0].getY() + L[0].getH();
    archivoSalida << alturaUsada << endl;
    archivoSalida << areaDesper << endl;

    // El archivo debe estar ordenado por ID.
    sort(L.begin(), L.end(), compararIDAsc);

    // Formato de linea: X Y R
    for (int i = 0; i < L.size(); i++) {
        archivoSalida << L[i].getX() << " " << L[i].getY() << " " << L[i].isRoted() << endl;
    }

    archivoSalida.close();
}





/*
Algoritmo de colocacion:
Ir ingresando cada bloque en el orden en como estan en la lista.
Si no cabe, continuar con el siguiente bloque.
Si no cabe ninguno, crear un nuevo nivel y volver a intentar.
Repetir hasta que todos los bloques hayan sido ingresados.
Al final, retorna el area sin usar.

NOTA: Esta funcion modifica la lista de bloques original.
*/
int colocarBloques(int maxW, vector<Bloque>& L, int totB) {

    int totH   = 0;
    int totA   = 0;
    int levelN = 0;
    int usedB  = 0;

    // cout << "=======COLOCACION=======" << endl;

    // Ir agregando bloques hasta que se hayan usado todos.
    while (usedB != totB) {

        levelN++;
        int levelW = 0;
        int levelH = 0;
        int levelB = 0;

        // cout << "=====CREANDO NIVEL " << levelN << "=====" << endl;

        // Ingresar cada bloque en orden mientras quepan en el nivel actual.
        for (int i = 0; i < totB; i++) {

            // cout << left << "B:" << setw(4) << L[i].getID() << "R:" << setw(4) << L[i].isRoted() << "W:" << setw (4) << L[i].getW() << "H:" << setw(4) << L[i].getH();

            // Si el bloque no ha sido usado, intentar agregarlo.
            if (!L[i].isUsed()) {

                // cout << "TEST...";

                // Si el bloque cabe en el nivel actual, agregarlo. Si no, continuar con el siguiente bloque.
                if (levelW + L[i].getW() <= maxW) {

                    L[i].setCoords(levelW, totH);

                    levelW += L[i].getW();
                    levelH =  max(levelH, L[i].getH());
                    totA   += L[i].getA();

                    L[i].use();
                    usedB++;
                    levelB++;

                    // cout << "ADD TO " << levelN << " (levelW:" << levelW << " , levelH:" << levelH << ")" << endl;
                }
                // else
                // {
                //     // El bloque no cabe en el nivel actual.
                //     cout << "FAIL!" << endl;
                // }
            }
            // else
            // {
            //     // El bloque ya fue usado.
            //     cout << "USED" << endl;
            // }
        }

        totH += levelH;

        // cout << usedB << " de " << totB << " bloques usados. (levelB:" << levelB << " , levelW:" << levelW << " , levelH:" << levelH << ")" << endl;
    }

    // cout << "==========STATS=========" << endl;
    // cout << "Niveles creados : " << levelN << endl;
    // cout << "Altura usada    : " << totH << endl;
    // cout << "Area usada      : " << totA << " de " << maxW * totH << endl;
    // cout << "Area sin usar   : " << maxW * totH - totA << endl;
    // cout << "Bloques despues de colocarBloques:" << endl;
    // printBloques(L);

    // cout << "=====FIN COLOCACION=====" << endl;

    return maxW * totH - totA;
}



/*
Probar instancia:
Primero rota los bloques que esten marcados en la instancia.
Luego ordena los bloques por altura.
Finalmente, llama a la funcion de colocacion.

NOTA: Esta funcion modifica la lista de bloques original.
*/
int testInstRotOrd(vector<Bloque>& L, Instancia rots, int maxW, int totB) {
    // cout << "=======PROBANDO INSTANCIA=======" << endl;
    // cout << "Instancia: ";
    // rots.imprimirInstancia();

    // cout << "Bloques antes de rotar:" << endl;
    // printBloques(L);

    // cout << "Rotando bloques..." << endl;
    for (int i = 0; i < totB; i++) {
        // cout << left << "B" << setw(4) << L[i].getID() << "...";

        if (rots.estaRotado(i)) {
            // cout << "ROTATE" << endl;
            L[i].rotate();
        }
        // else
        // {
        //     cout << "NO ROTATE" << endl;
        // }
    }

    // cout << "Bloques despues de rotar:" << endl;
    // printBloques(L);

    // cout << "Ordenando bloques por altura..." << endl;
    sort(L.begin(), L.end(), compararHDesc);

    // cout << "Colocando bloques..." << endl;
    int result = colocarBloques(maxW, L, totB);

    // cout << "=====FIN PROBANDO INSTANCIA=====" << endl;

    // cout << "Bloques dentro de probarInstanciaConOrden:" << endl;
    // printBloques(L);

    return result;
}


/*
Probar instancia:
Primero ordena los bloques por altura.
Luego rota los bloques que estén marcados en la instancia.
Finalmente, llama a la función de colocación.

NOTA: Esta función modifica la lista de bloques original.
*/
int testInstOrdRot(vector<Bloque>& L, Instancia rots, int maxW, int totB) {
    // cout << "=======PROBANDO INSTANCIA=======" << endl;
    // cout << "Instancia: ";
    // rots.imprimirInstancia();

    // cout << "Ordenando bloques por altura..." << endl;
    sort(L.begin(), L.end(), compararHDesc);

    // cout << "Bloques antes de rotar:" << endl;
    // printBloques(L);

    // cout << "Rotando bloques..." << endl;
    for (int i = 0; i < totB; i++) {
        // cout << left << "B" << setw(4) << L[i].getID() << "...";

        if (rots.estaRotado(i)) {
            // cout << "ROTATE" << endl;
            L[i].rotate();
        }
        // else
        // {
        //     cout << "NO ROTATE" << endl;
        // }
    }

    // cout << "Bloques despues de rotar:" << endl;
    // printBloques(L);

    // cout << "Colocando bloques..." << endl;
    int result = colocarBloques(maxW, L, totB);

    // cout << "=====FIN PROBANDO INSTANCIA=====" << endl;

    // cout << "Bloques dentro de probarInstanciaConOrden:" << endl;
    // printBloques(L);

    return result;
}








int main(int argc, char* argv[]) {

    // ======== INICIALIZACION =========
    // Verificar argumentos
    if (argc != 2) {
        cout << "Modo de uso: " << argv[0] << " <filename>" << endl;
        return 1;
    }

    // Abrir archivo
    string direccionArchivo = argv[1];
    ifstream archivoAbierto(direccionArchivo);
    if (!archivoAbierto.is_open()) {
        cout << "Error abriendo archivo " << direccionArchivo << endl;
        return 1;
    }

    // Crear lista de bloques.
    vector<Bloque> listaBloquesOriginal;

    // ========= LECTURA ARCHIVO =========
    string lineaLeida;

    getline(archivoAbierto, lineaLeida);             // Primera linea: total de rectangulos
    int totalBloques = stoi(lineaLeida);

    getline(archivoAbierto, lineaLeida);             // Segunda linea: ancho maximo
    int anchoMaximo = stoi(lineaLeida);

    for (int i = 0; i < totalBloques; i++) {         // Lineas restantes: datos de cada rectangulo
        getline(archivoAbierto, lineaLeida);
        stringstream stringLinea(lineaLeida);
        string stringNumero;

        getline(stringLinea, stringNumero, ' ');     // Primer numero: identificador
        int idntf = stoi(stringNumero);

        getline(stringLinea, stringNumero, ' ');     // Segundo numero: ancho
        int ancho = stoi(stringNumero);

        getline(stringLinea, stringNumero, ' ');     // Tercer numero: alto
        int alto = stoi(stringNumero);

        Bloque bloqueActual(idntf, ancho, alto);     // Crear bloque con los datos leidos y agregarlo a la lista de bloques.
        listaBloquesOriginal.push_back(bloqueActual);
    }

    archivoAbierto.close();
    // ========= FIN LECTURA ARCHIVO =========

    // Extraer nombre del archivo
    string nombreArchivo;
    size_t pos = direccionArchivo.find_last_of("/\\");
    if (pos != string::npos) {
        nombreArchivo = direccionArchivo.substr(pos + 1);
    }
    pos = nombreArchivo.find_last_of(".");
    if (pos != string::npos) {
        nombreArchivo = nombreArchivo.substr(0, pos);
    }

    // ========= FIN INICIALIZACION =========

    // Variables disponibles: totalBloques, anchoMaximo, listaBloques




    // ========== INICIANDO VARIABLES ==========
    vector<Bloque> listaBloques = listaBloquesOriginal;
    vector<Bloque> mejorListaBloques = listaBloquesOriginal;

    Instancia instanciaActual(totalBloques);
    Instancia mejorInstancia = instanciaActual;

    int mejorResultado = 99999999;
    int mejorIteracion = 0;

    cout << "Total de rectangulos: " << totalBloques << ". Ancho maximo: " << anchoMaximo << endl;
    cout << "Lista de rectangulos:" << endl;
    printBloques(listaBloques);
    // ========== FIN INICIANDO VARIABLES ==========


    // ========== INICIO HILL CLIMBING ==========
    // cout << "=====INICIO HILL CLIMBING=====" << endl;

    int maxIteraciones = 10;
    int evals = 0;
    int iter;
    clock_t start = clock();
    for (iter = 1; iter <= maxIteraciones; iter++) {
        // cout << "=====ITERACION " << iter << "=====" << endl;

        // ========= REVISION DE VECINOS =========
        bool algunaMejora = false;
        for (int indexB = 0; indexB < totalBloques; indexB++) {
            // Rotar bloque (si es que no es la primera iteracion
            if (iter != 1) {
                instanciaActual.marcarRotacion(indexB);
                // cout << "Probando rotacion de bloque " << indexB+1 << "..." << endl;
            }
            // else
            // {
            //     cout << "Probando solucion inicial..." << endl;
            // }

            listaBloques = listaBloquesOriginal; // Reiniciar lista de bloques.


            // int resultadoActual = testInstRotOrd(listaBloques, instanciaActual, anchoMaximo, totalBloques);
            int resultadoActual = testInstOrdRot(listaBloques, instanciaActual, anchoMaximo, totalBloques);


            evals++;
            // La funcion 'probarInstanciaConOrden' modifica la lista de bloques, por lo que se debe volver a copiar.

            // Mientras menor sea el resultado, se considera mejor.
            if (resultadoActual < mejorResultado) {
                // cout << "Mejora encontrada!" << endl;
                mejorResultado = resultadoActual;
                mejorInstancia = instanciaActual;
                mejorIteracion = iter;
                algunaMejora = true;
                mejorListaBloques = listaBloques; // Guardar lista de bloques en una copia.
            }
            else
            {
                // cout << "No hubo mejora." << endl;
                instanciaActual.marcarRotacion(indexB); // Deshacer rotacion.
            }

            // cout << "Resultado actual: " << resultadoActual << endl;
            // cout << "Mejor resultado: " << mejorResultado << endl;
            // cout << "Mejor instancia: "; mejorInstancia.imprimirInstancia();
            // cout << "Mejor iteracion: " << mejorIteracion << endl;

            if (algunaMejora) { break; }
        }
        // cout << "=====FIN ITERACION " << iter << "=====" << endl;
        // ========= FIN REVISION DE VECINOS =========

        if (!algunaMejora) {
            // cout << "No hubo ninguna mejora en esta iteracion.\nMINIMO LOCAL." << endl;
            break;
        }
    }
    clock_t end = clock();
    // cout << "=====FIN HILL CLIMBING=====" << endl;
    // ========== FIN HILL CLIMBING ==========





    clock_t nCiclos = end - start;
    double milisegundos = double(nCiclos) / (CLOCKS_PER_SEC / 1000);

    cout << "=====RESULTADO FINAL=====" << endl;
    cout << "Mejor resultado: " << mejorResultado << endl;
    cout << "Mejor instancia: " << mejorInstancia.toString() << endl;
    cout << "Mejor iteracion: " << mejorIteracion << endl;
    cout << "Total de iteraciones: " << iter << endl;
    cout << "Llamadas a la funcion: " << evals << endl;
    cout << "Tiempo de ejecucion: " << milisegundos << " ms" << endl;
    cout << "Bloques en main:" << endl;
    printBloques(mejorListaBloques);
    generarOutput(mejorListaBloques, mejorInstancia, mejorResultado, nombreArchivo);
    cout << "=====FIN RESULTADO FINAL=====" << endl;






    // ========== GENERACION ARCHIVO SALIDA ==========
    // Guardar resultados en archivo 'results.csv'
    ofstream archivoResultados("./output/results.tsv", ios::app);
    if (!archivoResultados.is_open()) {
        cout << "Error creando archivo de resultados." << endl;
        return 1;
    }

    // Formato separado por punto y comas:
    // nombreArchivo\ttotalBloques\tanchoMaximo\ttotalIteraciones\tmejorIteracion\tcantidadEvaluaciones\tmejorResultado\tmejorInstancia\tmilisegundos
    archivoResultados << nombreArchivo << "\t" << totalBloques << "\t" << anchoMaximo << "\t" << iter << "\t" << mejorIteracion << "\t" << evals << "\t" << mejorResultado << "\t" << mejorInstancia.toString() << "\t" << milisegundos << endl;

    archivoResultados.close();
    // ========== FIN GENERACION ARCHIVO SALIDA ==========





    return 0;
}
