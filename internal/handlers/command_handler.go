package handlers

import (
	"fmt"
	"net"
	"strings"

	"github.com/mauFade/chat-server-tcp/internal/models"
)

type CommandHandler struct {
	clients *models.Client
}

func NewCommandHandler(clients *models.Client) *CommandHandler {
	return &CommandHandler{
		clients: clients,
	}
}

func (h *CommandHandler) HandleCommand(conn net.Conn, message string) (bool, error) {
	if !strings.HasPrefix(message, "/") {
		return false, nil
	}

	parts := strings.Fields(message)
	command := parts[0]
	args := parts[1:]

	_, exists := models.Commands[command]
	if !exists {
		return true, fmt.Errorf("comando desconhecido: %s", command)
	}

	switch command {
	case "/help":
		h.showHelp(conn)
	case "/list":
		h.clients.ListClients(conn)
	case "/quit":
		h.clients.RemoveClient(conn)
	case "/room":
		return true, h.handleRoomCommand(conn, args)
	case "/whisper":
		return true, h.handleWhisperCommand(conn, args)
	case "/me":
		return true, h.handleMeCommand(conn, args)
	case "/clear":
		h.clearScreen(conn)
	case "/status":
		return true, h.handleStatusCommand(conn, args)
	}

	return true, nil
}

func (h *CommandHandler) showHelp(conn net.Conn) {
	conn.Write([]byte("\nComandos disponíveis:\n"))
	for _, cmd := range models.Commands {
		conn.Write([]byte(fmt.Sprintf("%s - %s\n", cmd.Usage, cmd.Description)))
	}
	conn.Write([]byte("\n"))
}

func (h *CommandHandler) handleRoomCommand(conn net.Conn, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("uso: /room [create|join|leave|list] [room_name]")
	}

	action := args[0]
	switch action {
	case "create":
		if len(args) < 2 {
			return fmt.Errorf("uso: /room create [room_name]")
		}
		roomName := args[1]
		// TODO: Implementar criação de sala
		conn.Write([]byte(fmt.Sprintf("Sala '%s' criada com sucesso!\n", roomName)))
	case "join":
		if len(args) < 2 {
			return fmt.Errorf("uso: /room join [room_name]")
		}
		roomName := args[1]
		// TODO: Implementar entrada em sala
		conn.Write([]byte(fmt.Sprintf("Entrou na sala '%s'\n", roomName)))
	case "leave":
		// TODO: Implementar saída de sala
		conn.Write([]byte("Saiu da sala atual\n"))
	case "list":
		// TODO: Implementar listagem de salas
		conn.Write([]byte("Lista de salas disponíveis:\n"))
	default:
		return fmt.Errorf("ação desconhecida: %s", action)
	}

	return nil
}

func (h *CommandHandler) handleWhisperCommand(conn net.Conn, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("uso: /whisper [user] [message]")
	}

	targetUser := args[0]
	message := strings.Join(args[1:], " ")

	// TODO: Implementar envio de mensagem privada
	conn.Write([]byte(fmt.Sprintf("Mensagem privada para %s: %s\n", targetUser, message)))
	return nil
}

func (h *CommandHandler) handleMeCommand(conn net.Conn, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("uso: /me [action]")
	}

	action := strings.Join(args, " ")
	nickname := h.clients.Clients[conn]

	// Broadcast da ação para todos os usuários
	h.clients.Broadcast(conn, fmt.Sprintf("* %s %s", nickname, action))
	return nil
}

func (h *CommandHandler) clearScreen(conn models.Connection) {
	// Envia sequência ANSI para limpar a tela
	conn.Write([]byte("\033[2J\033[H"))
}

func (h *CommandHandler) handleStatusCommand(conn models.Connection, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("uso: /status [online|away|busy]")
	}

	status := models.UserStatus(strings.ToLower(args[0]))
	switch status {
	case models.StatusOnline, models.StatusAway, models.StatusBusy:
		// TODO: Implementar mudança de status
		conn.Write([]byte(fmt.Sprintf("Status alterado para: %s\n", status)))
	default:
		return fmt.Errorf("status inválido. Use: online, away ou busy")
	}

	return nil
}
