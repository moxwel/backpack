def concat_str(lista):
    s = ""
    for x in lista:
        s += x
    return s

def get_palabra(palabra, clave):
    l_clave = len(clave)
    string_lista = list(palabra)

    if clave in palabra:
        pos_init = palabra.index(clave)

        for x in range(l_clave):
            del string_lista[pos_init]

        print(concat_str(string_lista))
    else:
        print("-1")

get_palabra("foniwida","iwi")