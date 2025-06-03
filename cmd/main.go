package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

var clients = make(map[string]string)

func main() {
	listen, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal("Error listening on port 8080: ", err.Error())
	}

	defer listen.Close()

	fmt.Println("Server is running on port 8080")

	for {
		conn, err := listen.Accept()

		if err != nil {
			log.Println("Error accepting connection: ", err.Error())
			continue
		}

		go handleClient(conn)

	}
}

func handleClient(c net.Conn) {
	defer c.Close()
	clients[c.RemoteAddr().String()] = c.RemoteAddr().String()

	reader := bufio.NewReader(c)

	for {
		message, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error reading message: ", err.Error())
			return
		}

		c.Write([]byte("Received message: " + message))
	}
}
