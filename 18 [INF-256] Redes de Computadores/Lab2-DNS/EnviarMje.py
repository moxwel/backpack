# Importo la libreria socket para manejar conexiones
import socket as skt

# Se declara dirección del servidor y puerto
serverAddr = 'localhost'
serverPort = 63420

# Se crea un socket para hacer el manejo de la conexión
clientSocket = skt.socket(skt.AF_INET, skt.SOCK_DGRAM)

# Solicito mensajes y luego ejecuto las funciones para conectar al servidor y realizar el intercambio de mensajes.
toSend = input("Texto a enviar:\n>")
if toSend == "STOP":
	clientSocket.sendto(toSend.encode(), (serverAddr,serverPort))
	print("Programa Finalizado")
else:
	clientSocket.sendto(toSend.encode(), (serverAddr,serverPort))

	msg, addr = clientSocket.recvfrom(1024)
	print(msg.decode()) 