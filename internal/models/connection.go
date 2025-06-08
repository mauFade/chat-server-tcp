package models

import "net"

type Connection interface {
	Write([]byte) (int, error)
	Read([]byte) (int, error)
	Close() error
	RemoteAddr() net.Addr
}
