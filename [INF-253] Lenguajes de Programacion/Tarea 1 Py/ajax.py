from os import write
import re # Expresiones regulares
import sys # Argumentos de programa

# Es necesario que haya un argumento de programa para poder abrir el archivo de las consultas
if len(sys.argv) == 1:
    print("ERROR: Sin archivo para leer.")
    print("")
    print("    Ejemplo de uso:")
    print("    python ajax.py entradas.txt")
    print("")
    sys.exit()

# Abriendo archivo de consultas
fileName = sys.argv[1]
try:
    queryFile = open(fileName, "r", encoding="utf-8")
except FileNotFoundError:
    print("ERROR: Archivo " + fileName + " no existe.\n")
    sys.exit()




# --------------------
# Expresiones regulares que se van a utilizar
ajaxRE = re.compile(r"\{\n\tdata: (?P<data>{.*}),\n\ttype: \"(?P<type>GET|POST|DELETE)\",\n\turl: \"(?P<url>.+)\"\n\}\n?")

# Denuevo fallé, perdoname Dios u.u
dictFullRE = re.compile(r"{(?:\"\w+\": (?:{(?:\"\w+\": (?:{(?:\"\w+\": (?:{(?:\"\w+\": (?:{(?:\"\w+\": (?:{(?:\"\w+\": (?:{(?:\"\w+\": (?:{(?:\"\w+\": (?:{(?:\"\w+\": (?:{(?:\"\w+\": (?:{(?:\"\w+\": (?:{(?:\"\w+\": (?:{(?:\"\w+\": (?:{(?:\"\w+\": (?:{(?:\"\w+\": (?:{(?:\"\w+\": (?:{(?:\"\w+\": (?:{(?:\"\w+\": (?:{(?:\"\w+\": (?:{(?:\"\w+\": (?:{(?:\"\w+\": (?:\[(?:\d*|\d+(?:, \d+)*)\]|\d+|\"[\w -]*\")(?:, )?)*}|\[(?:\d*|\d+(?:, \d+)*)\]|\d+|\"[\w -]*\")(?:, )?)*}|\[(?:\d*|\d+(?:, \d+)*)\]|\d+|\"[\w -]*\")(?:, )?)*}|\[(?:\d*|\d+(?:, \d+)*)\]|\d+|\"[\w -]*\")(?:, )?)*}|\[(?:\d*|\d+(?:, \d+)*)\]|\d+|\"[\w -]*\")(?:, )?)*}|\[(?:\d*|\d+(?:, \d+)*)\]|\d+|\"[\w -]*\")(?:, )?)*}|\[(?:\d*|\d+(?:, \d+)*)\]|\d+|\"[\w -]*\")(?:, )?)*}|\[(?:\d*|\d+(?:, \d+)*)\]|\d+|\"[\w -]*\")(?:, )?)*}|\[(?:\d*|\d+(?:, \d+)*)\]|\d+|\"[\w -]*\")(?:, )?)*}|\[(?:\d*|\d+(?:, \d+)*)\]|\d+|\"[\w -]*\")(?:, )?)*}|\[(?:\d*|\d+(?:, \d+)*)\]|\d+|\"[\w -]*\")(?:, )?)*}|\[(?:\d*|\d+(?:, \d+)*)\]|\d+|\"[\w -]*\")(?:, )?)*}|\[(?:\d*|\d+(?:, \d+)*)\]|\d+|\"[\w -]*\")(?:, )?)*}|\[(?:\d*|\d+(?:, \d+)*)\]|\d+|\"[\w -]*\")(?:, )?)*}|\[(?:\d*|\d+(?:, \d+)*)\]|\d+|\"[\w -]*\")(?:, )?)*}|\[(?:\d*|\d+(?:, \d+)*)\]|\d+|\"[\w -]*\")(?:, )?)*}|\[(?:\d*|\d+(?:, \d+)*)\]|\d+|\"[\w -]*\")(?:, )?)*}|\[(?:\d*|\d+(?:, \d+)*)\]|\d+|\"[\w -]*\")(?:, )?)*}|\[(?:\d*|\d+(?:, \d+)*)\]|\d+|\"[\w -]*\")(?:, )?)*}|\[(?:\d*|\d+(?:, \d+)*)\]|\d+|\"[\w -]*\")(?:, )?)*}|\[(?:\d*|\d+(?:, \d+)*)\]|\d+|\"[\w -]*\")(?:, )?)*}")

dictSingleRE = re.compile(r"(?:\"\w+\": (?:{(?:\"\w+\": (?:{(?:\"\w+\": (?:{(?:\"\w+\": (?:{(?:\"\w+\": (?:{(?:\"\w+\": (?:{(?:\"\w+\": (?:{(?:\"\w+\": (?:{(?:\"\w+\": (?:{(?:\"\w+\": (?:{(?:\"\w+\": (?:{(?:\"\w+\": (?:{(?:\"\w+\": (?:{(?:\"\w+\": (?:{(?:\"\w+\": (?:{(?:\"\w+\": (?:{(?:\"\w+\": (?:{(?:\"\w+\": (?:{(?:\"\w+\": (?:{(?:\"\w+\": (?:{(?:\"\w+\": (?:\[(?:\d*|\d+(?:, \d+)*)\]|\d+|\"[\w -]*\")(?:, )?)*}|\[(?:\d*|\d+(?:, \d+)*)\]|\d+|\"[\w -]*\")(?:, )?)*}|\[(?:\d*|\d+(?:, \d+)*)\]|\d+|\"[\w -]*\")(?:, )?)*}|\[(?:\d*|\d+(?:, \d+)*)\]|\d+|\"[\w -]*\")(?:, )?)*}|\[(?:\d*|\d+(?:, \d+)*)\]|\d+|\"[\w -]*\")(?:, )?)*}|\[(?:\d*|\d+(?:, \d+)*)\]|\d+|\"[\w -]*\")(?:, )?)*}|\[(?:\d*|\d+(?:, \d+)*)\]|\d+|\"[\w -]*\")(?:, )?)*}|\[(?:\d*|\d+(?:, \d+)*)\]|\d+|\"[\w -]*\")(?:, )?)*}|\[(?:\d*|\d+(?:, \d+)*)\]|\d+|\"[\w -]*\")(?:, )?)*}|\[(?:\d*|\d+(?:, \d+)*)\]|\d+|\"[\w -]*\")(?:, )?)*}|\[(?:\d*|\d+(?:, \d+)*)\]|\d+|\"[\w -]*\")(?:, )?)*}|\[(?:\d*|\d+(?:, \d+)*)\]|\d+|\"[\w -]*\")(?:, )?)*}|\[(?:\d*|\d+(?:, \d+)*)\]|\d+|\"[\w -]*\")(?:, )?)*}|\[(?:\d*|\d+(?:, \d+)*)\]|\d+|\"[\w -]*\")(?:, )?)*}|\[(?:\d*|\d+(?:, \d+)*)\]|\d+|\"[\w -]*\")(?:, )?)*}|\[(?:\d*|\d+(?:, \d+)*)\]|\d+|\"[\w -]*\")(?:, )?)*}|\[(?:\d*|\d+(?:, \d+)*)\]|\d+|\"[\w -]*\")(?:, )?)*}|\[(?:\d*|\d+(?:, \d+)*)\]|\d+|\"[\w -]*\")(?:, )?)*}|\[(?:\d*|\d+(?:, \d+)*)\]|\d+|\"[\w -]*\")(?:, )?)*}|\[(?:\d*|\d+(?:, \d+)*)\]|\d+|\"[\w -]*\")(?:, )?)*}|\[(?:\d*|\d+(?:, \d+)*)\]|\d+|\"[\w -]*\"))")

# Re de solo una iteracion:
# {(?:\"\w+\": (?:---|\[(?:\d*|\d+(?:, \d+)*)\]|\d+|\"[\w -]*\")(?:, (?!}))?)*}
# Donde esta "---" supuestamente esta el diccionario
# --------------------




# --------------------
# Funciones
def ajaxAction(action, data, url):
    """
    ajaxAction
    ----------
    Realiza acciones Ajax segun el tipo de consulta que se reciba.

    ----------
    Inputs:
    - action : String - Accion de la consulta Ajax. Puede ser POST, GET o DELETE.
    - data : String - Datos que se van a procesar.
    - url : String - Nombre del archivo a realizar las acciones.
    """

    if action == "POST":
        # print("[ajaxAction] Se va a realizar POST... Analizando estructura...")
        if data == "{}":
            print("Consulta POST no realizada... | RAZON: Data vacio.")
            return
        else:
            if validQuery(data):
                # print("[ajaxAction] Estructura de data correcta! Realizando accion")
                actionPOST(data, url)
                return
            else:
                print("Consulta POST no realizada... | RAZON: Data mal escrita.")
                return

    elif action == "GET":
        # print("[ajaxAction] Se va a realizar GET... Analizando estructura...")
        if data == "{}":
            # print("[ajaxAction] Data vacio, retornando todo el archivo...")
            printFile(url)
            return
        else:
            if validQuery(data):
                # print("[ajaxAction] Estructura de data correcta! Realizando accion")
                actionGET(data, url)
                return
            else:
                print("Consulta GET no realizada... | RAZON: Data mal escrita.")
                return

    elif action == "DELETE":
        # print("[ajaxAction] Se va a realizar DELETE...")
        if data == "{}":
            # print("[ajaxAction] Data vacio, borrando todo el archivo...")
            deleteFile(url)
            print("Consulta DELETE realizada!")
            return
        else:
            if validQuery(data):
                # print("[ajaxAction] Estructura de data correcta! Realizando accion")
                actionDELETE(data, url)
                print("Consulta DELETE realizada!")
                return
            else:
                print("Consulta DELETE no realizada... | RAZON: Data mal escrita.")
                return




def validQuery(data):
    """
    validDict
    ----------
    Revisa si el data de una consulta esta bien estructurada.

    ----------
    Inputs:
    - data : String - Data a verificar

    ----------
    Return:
    - True : Exito
    - False : No coinicide la estructura
    """

    # print("[validQuery] Analizando estrucutra de data...")
    # print("    " + data)
    result = dictFullRE.fullmatch(data)
    # print(result)
    if result != None:
        return True
    else:
        return False




def actionPOST(data, url):
    """
    actionPOST
    ----------
    Realiza la accion POST en un archivo.
    Añade la data de una consulta Ajax al final de un archivo dado por url.

    ----------
    Inputs:
    - data : String - Dato a usar para realizar POST.
    - url : String - Nombre de archivo.
    """

    actionFile = open(url, "a", encoding="utf-8")
    actionFile.write(data + "\n")
    print("Consulta POST realizada!")
    actionFile.close()




def actionGET(data, url):
    """
    actionGET
    ----------
    Realiza la accion GET en un archivo.
    Busca en el archivo de "url" datas que coincidan con el dado e imprime por pantalla.
    Si ningun data coincide, no se imprimirá nada.

    ----------
    Inputs:
    - data : String - Dato a usar para realizar GET.
    - url : String - Nombre de archivo.
    """

    # print(data)
    result = dictSingleRE.findall(data)
    # print(result)

    # Abrir archivo para leerlo
    try:
        openFile = open(url, "r", encoding="utf-8")
    except FileNotFoundError:
        print("[actionGET] ERROR: No existe el archivo " + url)
        return

    # Hay que ver si una linea contiene todos los matches necesarios
    goal = len(result)
    counter = 0
    for line in openFile:
        # print(line, end="")
        for item in result:
            result2 = re.search(item, line)
            if result2 != None:
                counter = counter + 1

        # Si los matches necesarios coinciden, entonces hay que imprimirlo
        if counter == goal:
            print(line, end="")

        counter = 0

    openFile.close()



def actionDELETE(data, url):
    """
    actionDELETE
    ----------
    Realiza la accion DELETE en un archivo.
    Busca en el archivo de "url" datas que coincidan con el dado y los borra.

    ----------
    Inputs:
    - data : String - Dato a usar para realizar DELETE.
    - url : String - Nombre de archivo.
    """

    # print(data)
    result = dictSingleRE.findall(data)
    # print(result)

    # Abrir archivo para leerlo
    try:
        openFile = open(url, "r", encoding="utf-8")
    except FileNotFoundError:
        print("[actionDELETE] ERROR: No existe el archivo " + url)
        return

    # Hay que ver si una linea contiene todos los matches necesarios
    goal = len(result)
    rewrite = []
    counter = 0
    for line in openFile:
        for item in result:
            result2 = re.search(item, line)
            if result2 != None:
                counter = counter + 1

        # Hay que ir añadiendo las lineas que no calcen en el match
        if counter != goal:
            rewrite.append(line)

        counter = 0

    openFile.close()

    # Ahora se reescribe el archivo sin las lineas que calazaron en el match, pues, se borran
    openFile = open(url, "w", encoding="utf-8")
    for item in rewrite:
        openFile.write(item)
    openFile.close()




def deleteFile(fileName):
    """
    deleteFile
    ----------
    Vacia un archivo dado. Borra todo su contenido.

    ----------
    Inputs:
    - fileName : String - Nombre del archivo
    """

    try:
        openFile = open(fileName, "w", encoding="utf-8")
        openFile.write("")
        openFile.close()
    except FileNotFoundError:
        # print("[deleteFile] ERROR: No existe el archivo " + fileName)
        return




def printFile(fileName):
    """
    printFile
    ----------
    Imprime por pantalla todas las lineas del archivo dado

    ----------
    Inputs:
    - fileName : String - Nombre del archivo
    """

    try:
        openFile = open(fileName, "r", encoding="utf-8")
        for line in openFile:
            print(line, end="")
        openFile.close()
    except FileNotFoundError:
        # print("[printFile] ERROR: No existe el archivo " + fileName)
        return
# --------------------




# --------------------
# Main
query = ""
for line in queryFile:
    if line == "-----\n":
        # print("\n[main] Se termino de leer una consulta. Detectando estructura...")
        # print(query)

        # Hay que asegurarse que haya match, si no, continuar con la siguiente consulta
        result = ajaxRE.fullmatch(query)

        # print(result)
        if result == None:
            print("Consulta mal escrita!")
            query = "" # Reiniciar para otra consulta nueva
            continue

        # Se tienen los datos separados para poder analizaros de forma independiente.
        queryData = result.group("data")
        queryType = result.group("type")
        queryUrl = result.group("url")
        # print("[main] Datos leidos:\n    " + queryData + "\n    " + queryType + "\n    " + queryUrl)

        # Realizar las acciones correspondientes segun el tipo de consulta
        ajaxAction(queryType, queryData, queryUrl)


        query = "" # Reiniciar para otra consulta nueva
        # print("")
    else:
        # Ir uniendo las lineas con sus caracteres para separar cada query
        query = query + line

# print("\n\nUltima consulta:")
# Hay que asegurarse que haya match, si no, continuar con la siguiente consulta
result = ajaxRE.fullmatch(query)

# print(result)
if result == None:
    print("Consulta mal escrita!")
    sys.exit()


# Se tienen los datos separados para poder analizaros de forma independiente.
queryData = result.group("data")
queryType = result.group("type")
queryUrl = result.group("url")
# print("[main] Datos leidos:\n    " + queryData + "\n    " + queryType + "\n    " + queryUrl)

# Realizar las acciones correspondientes segun el tipo de consulta
ajaxAction(queryType, queryData, queryUrl)

queryFile.close()
