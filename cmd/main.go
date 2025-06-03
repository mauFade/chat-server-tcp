package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"slices"
	"strings"
	"sync"
)

type clients struct {
	mu      sync.RWMutex
	clients map[string]net.Conn
}

func (c *clients) addClient(conn net.Conn) {
	c.mu.Lock()
	defer c.mu.Unlock()
	fmt.Println("Adding client: ", conn.RemoteAddr().String())
	c.clients[conn.RemoteAddr().String()] = conn
}

func (c *clients) removeClient(strConn string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	fmt.Println("Removed client: ", strConn)
	delete(c.clients, strConn)
}

func (c *clients) listClients() {
	c.mu.RLock()
	defer c.mu.RUnlock()

	for _, c := range c.clients {
		c.Write([]byte("conn: " + c.RemoteAddr().String() + "\n"))
	}
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

	commands := []string{
		"/list\n", // List connections
	}

	for {
		message, err := reader.ReadString('\n')

		if err != nil {
			// fmt.Println("Error reading message: ", err.Error())
			clients.removeClient(c.RemoteAddr().String())
			return
		}

		if strings.HasPrefix(message, "/") && slices.Contains(commands, message) {
			switch message {
			case "/list\n":
				clients.listClients()
			default:
				c.Write([]byte("Received message: " + message))
			}
		}

	}
}
