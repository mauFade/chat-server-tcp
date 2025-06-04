package models

import (
	"fmt"
	"net"
	"sync"
)

type Client struct {
	mu      sync.RWMutex
	Clients map[net.Conn]string
}

func (c *Client) AddClient(conn net.Conn, nickname string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	fmt.Println("Adding client: ", nickname)
	c.Clients[conn] = nickname
}

func (c *Client) RemoveClient(conn net.Conn) {
	c.mu.Lock()
	defer c.mu.Unlock()

	senderName, ok := c.Clients[conn]

	if !ok {
		senderName = conn.RemoteAddr().String()
	}

	fmt.Println("Removed client: ", senderName)
	delete(c.Clients, conn)
}

func (c *Client) ListClients(senderConn net.Conn) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	for conn, name := range c.Clients {
		if senderConn.RemoteAddr().String() != conn.RemoteAddr().Network() {
			senderConn.Write([]byte("user: " + name + ", ip: " + conn.RemoteAddr().String() + "\n"))
		}
	}
}

func (c *Client) Broadcast(senderConn net.Conn, message string) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	senderName := c.Clients[senderConn]
	for conn := range c.Clients {
		conn.Write([]byte("[" + senderName + "]: " + message + "\n"))
	}
}

func (c *Client) ShowHelp(conn net.Conn, commands map[string]string) {
	conn.Write([]byte("\nCommands list:" + "\n\n"))
	for comm, desc := range commands {
		conn.Write([]byte(comm + ": " + desc + "\n"))
	}
}
