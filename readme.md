# Chat Server TCP

Um servidor de chat em tempo real implementado em Golang usando TCP/IP e MongoDB para persistência de dados.

## Características

- Comunicação em tempo real via TCP/IP
- Sistema de salas de chat
- Persistência de mensagens em MongoDB
- Sistema de usuários com nicknames
- Comandos interativos
- Interface via terminal

## Pré-requisitos

- Docker e Docker Compose
- Netcat (para testar a conexão)

## Configuração

1. Clone o repositório:

```bash
git clone https://github.com/mauFade/chat-server-tcp.git
cd chat-server-tcp
```

2. Inicie a aplicação usando Docker Compose:

```bash
docker compose up -d
```

## Como Usar

1. Com o servidor rodando, abra um novo terminal e conecte-se usando netcat:

```bash
nc localhost 8080
```

2. Ao conectar, você será solicitado a:

   - Escolher um nickname (entre 2 e 20 caracteres)
   - Selecionar uma sala de chat disponível

3. Comandos disponíveis:

   - `/list` - Lista todos os usuários conectados
   - `/quit` - Sai do chat
   - `/help` - Mostra todos os comandos disponíveis

4. Para enviar mensagens, basta digitar e pressionar Enter.

## Estrutura do Projeto

```
.
├── cmd/
│   └── main.go         # Ponto de entrada da aplicação
├── internal/
│   ├── models/         # Definições de estruturas de dados
│   └── repository/     # Camada de persistência com MongoDB
├── docker-compose.yml  # Configuração dos serviços Docker
├── Dockerfile         # Configuração do build da aplicação
└── go.mod             # Dependências do projeto
```

## Tecnologias Utilizadas

- Go (Golang)
- MongoDB
- Docker
- TCP/IP

## Contribuindo

Sinta-se à vontade para abrir issues ou enviar pull requests com melhorias.

## Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.
