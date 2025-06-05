package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/mauFade/chat-server-tcp/internal/models"
	"github.com/mauFade/chat-server-tcp/internal/repository"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error connecting to .env file: " + err.Error())
	}
}

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
	c.Write([]byte("Enter your nickname: "))

	nick, _ := reader.ReadString('\n')
	nick = strings.TrimSpace(nick)

	for nick == "" || len(nick) > 20 || len(nick) < 2 {
		if nick == "" {
			c.Write([]byte("Nickname cannot be empty. "))
		} else if len(nick) > 20 {
			c.Write([]byte("Nickname too long (max 20 chars). "))
		} else {
			c.Write([]byte("Nickname too short (min 2 chars). "))
		}
		c.Write([]byte("Please enter a valid nickname: "))
		nick, _ = reader.ReadString('\n')
		nick = strings.TrimSpace(nick)
	}

	client, err := mongo.Connect(options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if err != nil {
		log.Fatal("Erro ao conectar MongoDB:", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	defer client.Disconnect(ctx)

	userRepo := repository.NewUserRepository(client)

	userr := userRepo.FindByNickname(nick)

	fmt.Println(userr)

	u := models.User{
		ID:        bson.NewObjectID(),
		Nickname:  nick,
		Room:      "",
		LastIP:    c.RemoteAddr().String(),
		CreatedAt: time.Now(),
	}

	err = userRepo.CreateUser(u)
	if err != nil {
		c.Write([]byte("Error creating user: " + err.Error()))
	}
	c.Write([]byte("Sucessfully created user! Welcome!\n"))

	clients.AddClient(c, nick)

	commandsMap := map[string]string{
		"/list": "List connections",
		"/quit": "Quit chat",
		"/help": "Show all available commands",
	}

	for {
		message, err := reader.ReadString('\n')
		message = strings.TrimSpace(message)
		if err != nil {
			clients.RemoveClient(c)
			return
		}

		_, isCommand := commandsMap[message]

		if strings.HasPrefix(message, "/") && isCommand {
			switch message {
			case "/list":
				clients.ListClients(c)
			case "/quit":
				clients.RemoveClient(c)
			case "/help":
				clients.ShowHelp(c, commandsMap)
			}

		} else {
			clients.Broadcast(c, message)
		}

	}
}
