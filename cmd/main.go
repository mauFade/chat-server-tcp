package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/mauFade/chat-server-tcp/internal/handlers"
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

	cs := models.Client{
		Clients: make(map[net.Conn]string),
	}

	commandHandler := handlers.NewCommandHandler(&cs)

	fmt.Println("Server is running on port 8080")

	for {
		conn, err := listen.Accept()

		if err != nil {
			log.Println("Error accepting connection: ", err.Error())
			continue
		}

		go handleClient(conn, &cs, commandHandler)
	}
}

func handleClient(c net.Conn, clients *models.Client, commandHandler *handlers.CommandHandler) {
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
	messageRepo := repository.NewMessageRepository(client)

	existingUser := userRepo.FindByNickname(nick)

	if existingUser == nil {
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
		existingUser = &u
	}

	clients.AddClient(c, nick)

	if existingUser.Room == "" {
		roomsRepo := repository.NewRoomRepository(client)

		rooms := roomsRepo.FindManyRooms()
		c.Write([]byte("select one of the following rooms to be part of: \n\n"))
		for i, room := range rooms {
			c.Write([]byte(strconv.Itoa(i+1) + " - " + room.Name + "\n"))
		}
		c.Write([]byte("\n[option]: "))
		roomOpt, _ := reader.ReadString('\n')
		roomOpt = strings.TrimSpace(roomOpt)

		index, err := strconv.Atoi(roomOpt)

		for err != nil || index < 1 || index > len(rooms) {
			c.Write([]byte("please enter a valid option: "))
			roomOpt, _ = reader.ReadString('\n')
			roomOpt = strings.TrimSpace(roomOpt)
			index, err = strconv.Atoi(roomOpt)
		}

		r := rooms[(index - 1)]

		existingUser.Room = r.Name
		userRepo.UpdateUserRoom(existingUser.ID, r.Name)
	}

	ms, err := messageRepo.FindByRoom(existingUser.Room)
	if err != nil {
		fmt.Println(err.Error())
	}

	if len(ms) > 0 {
		for _, m := range ms {
			c.Write([]byte("[" + m.UserNickname + " at room: (" + m.Room + ")]: " + m.Content + "\n"))
		}
	}

	for {
		c.Write([]byte("[" + existingUser.Nickname + " at room: (" + existingUser.Room + ")]: "))
		message, err := reader.ReadString('\n')
		message = strings.TrimSpace(message)
		if err != nil {
			clients.RemoveClient(c)
			return
		}

		// Verifica se é um comando
		isCommand, err := commandHandler.HandleCommand(c, message)
		if err != nil {
			c.Write([]byte("Error: " + err.Error() + "\n"))
			continue
		}

		// Se não for um comando, processa como mensagem normal
		if !isCommand {
			m := models.Message{
				ID:           bson.NewObjectID(),
				Content:      message,
				Room:         existingUser.Room,
				UserNickname: existingUser.Nickname,
				OriginIP:     c.RemoteAddr().String(),
				CreatedAt:    time.Now(),
				Type:         models.MessageTypeText,
			}

			messageRepo.CreateMessage(m)
			clients.Broadcast(c, message)
		}
	}
}
