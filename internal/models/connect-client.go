package models

import (
	"fmt"
	"net"
	"sync"
)

type Client struct {
	mu      sync.RWMutex
	Clients map[string]net.Conn
}

func (c *Client) AddClient(conn net.Conn) {
	c.mu.Lock()
	defer c.mu.Unlock()
	fmt.Println("Adding client: ", conn.RemoteAddr().String())
	c.Clients[conn.RemoteAddr().String()] = conn
}

func (c *Client) RemoveClient(strConn string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	fmt.Println("Removed client: ", strConn)
	delete(c.Clients, strConn)
}

func (c *Client) ListClients() {
	c.mu.RLock()
	defer c.mu.RUnlock()

	for _, conn := range c.Clients {
		conn.Write([]byte("conn: " + conn.RemoteAddr().String() + "\n"))
	}
}

func (c *Client) Broadcast(senderConn, message string) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	for _, conn := range c.Clients {
		if conn.RemoteAddr().String() != senderConn {
			conn.Write([]byte("[" + senderConn + "]: " + message + "\n"))
		}
	}
}
