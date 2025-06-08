package models

type Command struct {
	Name        string
	Description string
	Handler     func(args []string) error
	Usage       string
}

var Commands = map[string]Command{
	"/help": {
		Name:        "help",
		Description: "Mostra todos os comandos disponíveis",
		Usage:       "/help",
	},
	"/list": {
		Name:        "list",
		Description: "Lista todos os usuários conectados",
		Usage:       "/list",
	},
	"/quit": {
		Name:        "quit",
		Description: "Sai do chat",
		Usage:       "/quit",
	},
	"/room": {
		Name:        "room",
		Description: "Gerencia salas de chat",
		Usage:       "/room [create|join|leave|list] [room_name]",
	},
	"/whisper": {
		Name:        "whisper",
		Description: "Envia mensagem privada para um usuário",
		Usage:       "/whisper [user] [message]",
	},
	"/me": {
		Name:        "me",
		Description: "Envia uma ação (ex: /me está digitando...)",
		Usage:       "/me [action]",
	},
	"/clear": {
		Name:        "clear",
		Description: "Limpa a tela do chat",
		Usage:       "/clear",
	},
	"/status": {
		Name:        "status",
		Description: "Define seu status (online/away/busy)",
		Usage:       "/status [status]",
	},
}
