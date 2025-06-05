# Um chat em TCP escrito em Golang

## É necessário ter Docker rodando na maquina

### Inicie o MongoDB pelo Docker:

```bash
docker compose up

```

### Em outro terminal, rode o projeto:

```bash
go run cmd/main.go

```

## Abra um último terminal e se conecte ao TCP:

```bash
nc localhost 8080

```
