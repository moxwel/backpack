package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	// Crear registro para DNS para guardar la informacion
	// de los dominios y sus direcciones IP
	dns_table := make(map[string]string)

	// El servidor DNS debe escuchar en el puerto 63420

	// Obtener direccion UDP
	addr, err := net.ResolveUDPAddr("udp", ":63420")
	if err != nil {
		fmt.Println("Error al resolver la dirección:", err)
		return
	}
	fmt.Println("Dirección UDP resuelta:", addr)

	// Crear servidor UDP
	conn_udp, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
		return
	}
	defer conn_udp.Close()
	buffer := make([]byte, 1024)
	fmt.Println("Servidor escuchando en UDP 63420...")

	// Esperar peticiones de los clientes
	for {
		n, addr, err := conn_udp.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error al leer los datos:", err)
			return
		}
		recibido := string(buffer[:n])

		// Los datos recibidos estaran en el siguiente formato:
		// "ADD dominio ip ttl tipo" - Agregar un registro a la tabla DNS
		// "dominio" - Obtener la direccion IP asociada al dominio

		datos := strings.Split(recibido, " ")
		if datos[0] == "ADD" {
			// Se asumen 4 datos
			// "dominio ip ttl tipo"
			dominio := datos[1]
			ip := datos[2]
			ttl := datos[3]
			tipo := datos[4]

			// Agregar un registro a la tabla DNS
			agregarRegistro(dns_table, dominio, ip, ttl, tipo)

			// Enviar respuesta
			respuesta := "Direccion agregada"
			_, err = conn_udp.WriteToUDP([]byte(respuesta), addr)
			if err != nil {
				fmt.Println("Error al enviar la respuesta:", err)
				return
			}
			fmt.Println("Direccion agregada:", dominio, ip, ttl, tipo)
		} else {
			// Obtener la direccion IP asociada al dominio
			dominio := datos[0]
			ip := obtenerRegistro(dns_table, dominio)

			// Enviar respuesta
			_, err = conn_udp.WriteToUDP([]byte(ip), addr)
			if err != nil {
				fmt.Println("Error al enviar la respuesta:", err)
				return
			}
			fmt.Println("Direccion IP de", dominio, ":", ip)
		}
	}
}

func agregarRegistro(registroDNS map[string]string, dominio string, ip string, ttl string, tipo string) {
	datos := ip + " " + ttl + " " + tipo
	registroDNS[dominio] = datos
}

func obtenerRegistro(registroDNS map[string]string, dominio string) string {
	// Obtener el registro del dominio, separa la ip del resto de los datos
	datos := registroDNS[dominio]

	if datos == "" {
		return "No hay registro."
	}

	ip := datos[:strings.Index(datos, " ")]
	return ip
}
