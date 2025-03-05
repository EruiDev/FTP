package commons

import "net"

type Info struct {
	CurrentDir     string
	Username       string
	IsLogged       bool
	Conn           net.Conn
	DataConnection net.Conn
}

type Command struct {
	Description string
	Handler     func(args []string, user *Info) error
}