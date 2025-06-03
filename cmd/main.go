package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
)

type clients struct {
	mu      sync.RWMutex
	clients map[string]net.Conn
}

func (c *clients) addClient(conn net.Conn) {
	c.mu.Lock()
	defer c.mu.Unlock()
	fmt.Println("Adding client: ", conn.RemoteAddr().String(), conn)
	c.clients[conn.RemoteAddr().String()] = conn
}

func main() {
	listen, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal("Error listening on port 8080: ", err.Error())
	}

	defer listen.Close()
	CLIENTS := clients{
		clients: make(map[string]net.Conn),
	}
	fmt.Println("Server is running on port 8080")

	for {
		conn, err := listen.Accept()

		if err != nil {
			log.Println("Error accepting connection: ", err.Error())
			continue
		}

		go handleClient(conn, &CLIENTS)

	}
}

func handleClient(c net.Conn, clients *clients) {
	defer c.Close()
	clients.addClient(c)

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
