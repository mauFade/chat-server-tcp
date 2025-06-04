package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"slices"
	"strings"

	"github.com/mauFade/chat-server-tcp/internal/models"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal("Error listening on port 8080: ", err.Error())
	}

	defer listen.Close()
	CLIENTS := models.Client{
		Clients: make(map[string]net.Conn),
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

func handleClient(c net.Conn, clients *models.Client) {
	defer c.Close()
	clients.AddClient(c)

	reader := bufio.NewReader(c)

	commands := []string{
		"/list\n", // List connections
	}

	for {
		message, err := reader.ReadString('\n')

		if err != nil {
			// fmt.Println("Error reading message: ", err.Error())
			clients.RemoveClient(c.RemoteAddr().String())
			return
		}

		if strings.HasPrefix(message, "/") && slices.Contains(commands, message) {
			switch message {
			case "/list\n":
				clients.ListClients()
			}
		} else {
			clients.Broadcast(c.RemoteAddr().String(), message)
		}

	}
}
