    README

Maximiliano
Felipe

=================================================================

    INSTRUCCIONES DE USO

1. Ejecutar el servidor primero.

    $ go run server.go

2. Ejecutar el cliente.

    $ python EnviarMje.py

3. Ejecutar comandos en el cliente.

================================================================

    DESARROLLO

El servidor escuchara conexiones UDP en el puerto 63420.

El cliente esperará la entrada del usuario para enviarselos al servidor.

Hay dos tipos de comando:

    - "ADD" : Agrega un nuevo registro al servidor.
              Debe tener el formato: "ADD <dominio> <ip> <TTL> <tipo>"
              Ejemplo: "ADD www.google.com 8.8.8.8 3600 A"

    - "<dominio>" : Consulta la direccion IP del dominio indicado.
                    Ejemplo: "www.google.com"
