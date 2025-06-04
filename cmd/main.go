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
		Clients: make(map[net.Conn]string),
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
	reader := bufio.NewReader(c)
	c.Write([]byte("Enter your nickname:"))
	nick, _ := reader.ReadString('\n')
	nick = strings.TrimSpace(nick)
	clients.AddClient(c, nick)

	commands := []string{
		"/list", // List connections
		"/quit", // Quit chat
		"/help", // Show all available commands
	}

	for {
		message, err := reader.ReadString('\n')
		message = strings.TrimSpace(message)
		if err != nil {
			clients.RemoveClient(c)
			return
		}

		if strings.HasPrefix(message, "/") && slices.Contains(commands, message) {
			switch message {
			case "/list":
				clients.ListClients(c)
			case "/quit":
				clients.RemoveClient(c)
			case "/help":
				clients.ShowHelp(c, commands)
			}

		} else {
			clients.Broadcast(c, message)
		}

	}
}
