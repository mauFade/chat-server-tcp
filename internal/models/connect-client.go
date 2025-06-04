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

func (c *Client) AddClient(conn net.Conn) {
	c.mu.Lock()
	defer c.mu.Unlock()
	fmt.Println("Adding client: ", conn.RemoteAddr().String())
	c.Clients[conn] = conn.RemoteAddr().String()
}

func (c *Client) RemoveClient(conn net.Conn) {
	c.mu.Lock()
	defer c.mu.Unlock()
	fmt.Println("Removed client: ", conn.RemoteAddr().String())
	delete(c.Clients, conn)
}

func (c *Client) ListClients() {
	c.mu.RLock()
	defer c.mu.RUnlock()

	for conn, name := range c.Clients {
		conn.Write([]byte("user: " + name + "\n"))
	}
}

func (c *Client) Broadcast(senderConn net.Conn, message string) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	senderName := c.Clients[senderConn]
	for conn := range c.Clients {
		if conn.RemoteAddr().String() != senderConn.RemoteAddr().String() {
			conn.Write([]byte("[" + senderName + "]: " + message + "\n"))
		}
	}
}
